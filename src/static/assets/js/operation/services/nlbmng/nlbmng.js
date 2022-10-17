$(document).ready(function () {
    setTableHeightForScroll('nlbList', 300)
});

// area 표시
function displayNlbInfo(targetAction) {
    if (targetAction == "REG") {
        $('#nlbCreateBox').toggleClass("active");
        $('#nlbInfoBox').removeClass("view");
        $('#nlbListTable').removeClass("on");
        var offset = $("#nlbCreateBox").offset();
        
        $("#TopWrap").animate({ scrollTop: offset.top }, 300);

        // form 초기화
        $("#regCspSshKeyName").val('');
        goFocus('nlbCreateBox');
        
    } else if (targetAction == "REG_SUCCESS") {
        $('#nlbCreateBox').removeClass("active");
        $('#nlbInfoBox').removeClass("view");
        $('#nlbListTable').addClass("on");

        var offset = $("#nlbCreateBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);

        // form 초기화
        $("#regCspSshKeyName").val('');
        $("#regProvider").val('');
        $("#regCregConnectionNameidrBlock").val('');

        getSshKeyList("name");
    } else if (targetAction == "DEL") {
        $('#nlbCreateBox').removeClass("active");
        $('#nlbInfoBox').addClass("view");
        $('#nlbListTable').removeClass("on");

        var offset = $("#sskKeyInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 300);

    } else if (targetAction == "DEL_SUCCESS") {
        $('#nlbCreateBox').removeClass("active");
        $('#nlbInfoBox').removeClass("view");
        $('#nlbListTable').addClass("on");

        var offset = $("#sskKeyInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);

        getSshKeyList("name");
    } else if (targetAction == "CLOSE") {
        $('#nlbCreateBox').removeClass("active");
        $('#nlbInfoBox').removeClass("view");
        $('#nlbListTable').addClass("on");

        var offset = $("#sskKeyInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);
    }
}

// SshKey 목록 조회
function getNlbList(sort_type) {
    var url = "/setting/resources" + "/nlb/list"
    axios.get(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get Nlb Data : ", result.data);
        var data = result.data.NlbList; // exception case : if null 
        var html = ""

        if (data == null) {
            html += '<tr><td class="overlay hidden" data-th="" colspan="5">No Data</td></tr>'

            $("#sList").empty();
            $("#sList").append(html);

            ModalDetail()
        } else {
            if (data.length) { // null exception if not exist
                if (sort_type) {
                    console.log("check : ", sort_type);
                    data.filter(list => list.name !== "").sort((a, b) => (a[sort_type] < b[sort_type] ? - 1 : a[sort_type] > b[sort_type] ? 1 : 0)).map((item, index) => (
                        html += '<tr onclick="showNlbInfo(\'' + item.id + '\');">'
                        + '<td class="overlay hidden column-50px" data-th="">'
                        + '<input type="hidden" id="nlb_info_' + index + '" value="' + item.name + '|' + item.connectionName + '|' + item.cspSshKeyName + '"/>'
                        + '<input type="checkbox" name="chk" value="' + item.name + '" id="raw_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>'
                        + '<td class="btn_mtd ovm" data-th="Name">' + item.id
                        + '</td>'
                        + '<td class="overlay hidden" data-th="connectionName">' + item.connectionName + '</td>'
                        + '<td class="overlay hidden" data-th="cspSshKeyName">' + item.cspSshKeyName + '</td>'
                        + '</tr>'
                    ))
                } else {
                    data.filter((list) => list.name !== "").map((item, index) => (
                        html += '<tr onclick="showNlbInfo(\'' + item.id + '\');">'
                        + '<td class="overlay hidden column-50px" data-th="">'
                        + '<input type="hidden" id="nlb_info_' + index + '" value="' + item.name + '"/>'
                        + '<input type="checkbox" name="chk" value="' + item.name + '" id="raw_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>'
                        + '<td class="btn_mtd ovm" data-th="id">' + item.id + '<span class="ov"></span></td>'
                        + '<td class="overlay hidden" data-th="connectionName">' + item.connectionName + '</td>'
                        + '<td class="overlay hidden" data-th="cspSshKeyName">' + item.cspSshKeyName + '</td>'
                        + '</tr>'
                    ))

                }

                $("#nlbList").empty();
                $("#nlbList").append(html);

                ModalDetail()

            }
        }

    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

function deleteNlb() {
    var selNlbId = "";
    var count = 0;

    $("input[name='chk']:checked").each(function () {
        count++;
        selNlbId = selNlbId + $(this).val() + ",";
    });
    selNlbId = selNlbId.substring(0, selNlbId.lastIndexOf(","));

    console.log("nlbId : ", selNlbId);
    console.log("count : ", count);

    if (selNlbId == '') {
        alert("삭제할 대상을 선택하세요.");
        return false;
    }

    if (count != 1) {
        alert("삭제할 대상을 하나만 선택하세요.");
        return false;
    }

    var url = "" + "" + selNlbId;
    var url = "/operation/services/nlb/list" + "/nlb/" + nlbID;
    axios.delete(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        var data = result.data;
        console.log(data);
        if (result.status == 200 || result.status == 201) {
            commonAlert(data.message);
            
            displayNlbInfo("DEL_SUCCESS");
            
            getNlbList("name");
        } else {
            commonAlert(data.error);
        }
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

function showNlbInfo(nlbId) {
    console.log("target showNlbInfo : ", nlbId);

    $(".stxt").html(nlbId);

    var url = "/operation/services/nlb/" + nlbId;
    console.log("nlb URL : ", url)

    return axios.get(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        var data = result.data.SshKeyInfo
        console.log("Show Data : ", data);

        var dtlCspSshKeyName = data.cspSshKeyName;
        var dtlDescription = data.description;
        var dtlUserID = data.userID;
        var dtlConnectionName = data.connectionName;
        var dtlPublicKey = data.publicKey;
        var dtlPrivateKey = data.privateKey;
        var dtlFingerprint = data.fingerprint;


        $('#dtlCspSshKeyName').empty();
        $('#dtlDescription').empty();
        $('#dtlUserID').empty();
        $('#dtlConnectionName').empty();
        $('#dtlPublicKey').empty();
        $('#dtlPrivateKey').empty();
        $('#dtlFingerprint').empty();

        $('#dtlCspSshKeyName').val(dtlCspSshKeyName);
        $('#dtlDescription').val(dtlDescription);
        $('#dtlUserID').val(dtlUserID);
        $('#dtlConnectionName').val(dtlConnectionName);
        $('#dtlPublicKey').val(dtlPublicKey);
        $('#dtlPrivateKey').val(dtlPrivateKey);
        $('#dtlFingerprint').val(dtlFingerprint);
        
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

function createNlb() {
    var cspNlbName = $("#regCspNlbName").val()
    var connectionName = $("#regConnectionName").val()

    console.log("info param cspNlbName : ", cspNlbName);
    console.log("info param connectionName : ", connectionName);

    if (!cspNlbName) {
        alert("Input New Nlb Name")
        $("#regCspNlbName").focus()
        return;
    }
    if (!connectionName) {
        alert("Input Connection Name")
        $("#regConnectionName").focus()
        return;
    }

    var url = "/operation/services" + "/nlb/reg"
    console.log("nlb URL : ", url)
    var obj = {
        name: cspNlbName,
        connectionName: connectionName
    }
    console.log("info connectionconfig obj Data : ", obj);
    if (cspNlbName) {
        axios.post(url, obj, {
            headers: {
                'Content-type': "application/json",                
            }
        }).then(result => {
            console.log(result);
            if (result.status == 200 || result.status == 201) {
                commonAlert("Success Create Nlb ")
                //등록하고 나서 화면을 그냥 고칠 것인가?
                displayNlbInfo("REG_SUCCESS");
                
            } else {
                commonAlert("Fail Create Nlb ")
            }
        }).catch((error) => {
            console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.statusText;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage);
        });
    } else {
        commonAlert("Input SSH Key Name")
        $("#regCspSshKeyName").focus()
        return;
    }
}

// updateNlb : TODO : update를 위한 form을 만들 것인가 ... 기존 detail에서 enable 시켜서 사용할 것인가
function updateNlb() {
    var sshKeyId = $("#dtlSshKeyId").val()
    var sshKeyName = $("#dtlSshKeyName").val()
    var cspSshKeyId = $("#dtlCspSshKeyId").val()
    var cspSshKeyName = $("#dtlCspSshKeyName").val()
    var description = $("#dtlDescription").val()
    var publicKey = $("#dtlPublicKey").val()
    var privateKey = $("#dtlPrivateKey").val()
    var fingerprint = $("#dtlFingerprint").val()
    var username = $("#dtlUsername").val()
    var verifiedUsername = $("#dtlVerifiedUsername").val()// TODO : 사용자 이름 입력 받아야 함.
    var connectionName = $("#dtlConnectionName").val()

    console.log("info param cspSshKeyName : ", cspSshKeyName);
    console.log("info param connectionName : ", connectionName);

    var url = "/setting/resources" + "/sshkey/del/" + sshKeyId
    console.log("ssh key URL : ", url)
    var obj = {
        connectionName : connectionName,
        id : sshKeyId,
        name : sshKeyName,
        cspSshKeyId	: cspSshKeyId,
        cspSshKeyName : cspSshKeyName,
        description	: description,
        privateKey : privateKey,
        publicKey :	publicKey,
        fingerprint	: fingerprint,
        username :	username,
        verifiedUsername :	verifiedUsername
    }
    console.log("info updateNlb obj Data : ", obj);
    if (cspSshKeyName) {
        axios.post(url, obj, {
            headers: {
                'Content-type': "application/json",
            }
        }).then(result => {
            console.log(result);
            if (result.status == 200 || result.status == 201) {
                commonAlert(" Nlb Modification Success")
                displayNlbInfo("REG_SUCCESS");// TODO : MODIFY 성공일 때 어떻게 처리 할지 정의해서 보완할 것.
                
            } else {
                commonAlert("Fail Create Nlb")
            }
        }).catch((error) => {
            console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.statusText;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage);
        });
    } else {
        commonAlert("Input SSH Key Name")
        $("#regCspSshKeyName").focus()
        return;
    }
}

function ModalDetail() {
    $(".dashboard .status_list tbody tr").each(function () {
        var $td_list = $(this),
            $status = $(".server_status"),
            $detail = $(".server_info");
        $td_list.off("click").click(function () {
            $td_list.addClass("on");
            $td_list.siblings().removeClass("on");
            $status.addClass("view");
            $status.siblings().removeClass("on");
            $(".dashboard.register_cont").removeClass("active");
            $td_list.off("click").click(function () {
                if ($(this).hasClass("on")) {
                    console.log("reg ok button click")
                    $td_list.removeClass("on");
                    $status.removeClass("view");
                    $detail.removeClass("active");
                } else {
                    $td_list.addClass("on");
                    $td_list.siblings().removeClass("on");
                    $status.addClass("view");

                    $status.siblings().removeClass("view");
                    $(".dashboard.register_cont").removeClass("active");
                }
            });
        });
    });
}
