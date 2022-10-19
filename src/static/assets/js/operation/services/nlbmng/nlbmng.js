$(document).ready(function () {
    setTableHeightForScroll("nlbList", 300);
    //getNlbList("LB Name");
    getCommonMcisList("nlbOnload", true, "", "id");
});

// area 표시
function displayNlbInfo(targetAction) {
    if (targetAction == "REG") {
        $("#nlbCreateBox").toggleClass("active");
        $("#nlbInfoBox").removeClass("view");
        $("#nlbListTable").removeClass("on");
        var offset = $("#nlbCreateBox").offset();

        $("#TopWrap").animate({ scrollTop: offset.top }, 300);

        // form 초기화
        $("#regCspSshKeyName").val("");
        goFocus("nlbCreateBox");
    } else if (targetAction == "REG_SUCCESS") {
        $("#nlbCreateBox").removeClass("active");
        $("#nlbInfoBox").removeClass("view");
        $("#nlbListTable").addClass("on");

        var offset = $("#nlbCreateBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);

        // form 초기화
        $("#regCspSshKeyName").val("");
        $("#regProvider").val("");
        $("#regCregConnectionNameidrBlock").val("");

       // getNLBList("LB Name");
    } else if (targetAction == "DEL") {
        $("#nlbCreateBox").removeClass("active");
        $("#nlbInfoBox").addClass("view");
        $("#nlbListTable").removeClass("on");

        var offset = $("#sskKeyInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 300);
    } else if (targetAction == "DEL_SUCCESS") {
        $("#nlbCreateBox").removeClass("active");
        $("#nlbInfoBox").removeClass("view");
        $("#nlbListTable").addClass("on");

        var offset = $("#sskKeyInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);

        getNLBList("LB Name");
    } else if (targetAction == "CLOSE") {
        $("#nlbCreateBox").removeClass("active");
        $("#nlbInfoBox").removeClass("view");
        $("#nlbListTable").addClass("on");

        var offset = $("#sskKeyInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);
    }
}

// SshKey 목록 조회
function getNLBList(mcisID) {
    var url = "/operation/services/mcis/" + mcisID + "/nlb/list";
    axios
        .get(url, {
            headers: {
                "Content-Type": "application/json",
            },
        })
        .then((result) => {
            console.log("get Nlb Data : ", result.data);
            var data = result.data.NlbList; // exception case : if null
            var html = "";

            if (data == null) {
                html +=
                    '<tr><td class="overlay hidden" data-th="" colspan="9">No Data</td></tr>';
                //$("#sList").empty();
                $("#nlbList").empty();
                $("#nlbList").append(html);

                ModalDetail();
            } else {
                if (data.length) {
                    // null exception if not exist
                    if (sort_type) {
                        console.log("check : ", sort_type);
                        data.filter((list) => list.name !== "")
                            .sort((a, b) =>
                                a[sort_type] < b[sort_type]
                                    ? -1
                                    : a[sort_type] > b[sort_type]
                                    ? 1
                                    : 0
                            )
                            .map(
                                (item, index) =>
                                    (html +=
                                        // item 값 확인해서 출력 필요.
                                        "<tr onclick=\"showNlbInfo('" +
                                        item.id +
                                        "');\">" +
                                        '<td class="overlay hidden column-50px" data-th="">' +
                                        '<input type="hidden" id="nlb_info_' +
                                        index +
                                        '" value="' +
                                        item.name +
                                        "|" +
                                        item.connectionName +
                                        "|" +
                                        item.cspSshKeyName +
                                        '"/>' +
                                        '<input type="checkbox" name="chk" value="' +
                                        item.name +
                                        '" id="raw_' +
                                        index +
                                        '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' +
                                        '<td class="btn_mtd ovm" data-th="Name">' +
                                        item.id +
                                        "</td>" +
                                        '<td class="overlay hidden" data-th="connectionName">' +
                                        item.connectionName +
                                        "</td>" +
                                        '<td class="overlay hidden" data-th="cspSshKeyName">' +
                                        item.cspSshKeyName +
                                        "</td>" +
                                        "</tr>")
                            );
                    } else {
                        data.filter((list) => list.name !== "").map(
                            (item, index) =>
                                (html +=
                                    "<tr onclick=\"showNlbInfo('" +
                                    item.id +
                                    "');\">" +
                                    '<td class="overlay hidden column-50px" data-th="">' +
                                    '<input type="hidden" id="nlb_info_' +
                                    index +
                                    '" value="' +
                                    item.name +
                                    '"/>' +
                                    '<input type="checkbox" name="chk" value="' +
                                    item.name +
                                    '" id="raw_' +
                                    index +
                                    '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' +
                                    '<td class="btn_mtd ovm" data-th="id">' +
                                    item.id +
                                    '<span class="ov"></span></td>' +
                                    '<td class="overlay hidden" data-th="connectionName">' +
                                    item.connectionName +
                                    "</td>" +
                                    '<td class="overlay hidden" data-th="cspSshKeyName">' +
                                    item.cspSshKeyName +
                                    "</td>" +
                                    "</tr>")
                        );
                    }

                    $("#nlbList").empty();
                    $("#nlbList").append(html);

                    ModalDetail();
                }
            }
        })
        .catch((error) => {
            console.warn(error);
            console.log(error.response);
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

    if (selNlbId == "") {
        alert("삭제할 대상을 선택하세요.");
        return false;
    }

    if (count != 1) {
        alert("삭제할 대상을 하나만 선택하세요.");
        return false;
    }

    var url = "" + "" + selNlbId;
    var url = "/operation/services/nlb/list" + "/nlb/" + nlbID;
    axios
        .delete(url, {
            headers: {
                "Content-Type": "application/json",
            },
        })
        .then((result) => {
            var data = result.data;
            console.log(data);
            if (result.status == 200 || result.status == 201) {
                commonAlert(data.message);

                displayNlbInfo("DEL_SUCCESS");

                getNlbList("name");
            } else {
                commonAlert(data.error);
            }
        })
        .catch((error) => {
            console.warn(error);
            console.log(error.response);
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage);
        });
}

function showNlbInfo(nlbId) {
    console.log("target showNlbInfo : ", nlbId);

    $(".stxt").html(nlbId);

    var url = "/operation/services/nlb/" + nlbId;
    console.log("nlb URL : ", url);

    return axios
        .get(url, {
            headers: {
                "Content-Type": "application/json",
            },
        })
        .then((result) => {
            var data = result.data.SshKeyInfo;
            console.log("Show Data : ", data);

            var dtlCspSshKeyName = data.cspSshKeyName;
            var dtlDescription = data.description;
            var dtlUserID = data.userID;
            var dtlConnectionName = data.connectionName;
            var dtlPublicKey = data.publicKey;
            var dtlPrivateKey = data.privateKey;
            var dtlFingerprint = data.fingerprint;

            $("#dtlCspSshKeyName").empty();
            $("#dtlDescription").empty();
            $("#dtlUserID").empty();
            $("#dtlConnectionName").empty();
            $("#dtlPublicKey").empty();
            $("#dtlPrivateKey").empty();
            $("#dtlFingerprint").empty();

            $("#dtlCspSshKeyName").val(dtlCspSshKeyName);
            $("#dtlDescription").val(dtlDescription);
            $("#dtlUserID").val(dtlUserID);
            $("#dtlConnectionName").val(dtlConnectionName);
            $("#dtlPublicKey").val(dtlPublicKey);
            $("#dtlPrivateKey").val(dtlPrivateKey);
            $("#dtlFingerprint").val(dtlFingerprint);
        })
        .catch((error) => {
            console.warn(error);
            console.log(error.response);
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage);
        });
}

function createNlb() {
    console.log("enter createNlb");
    var nlbName = $("#reg_name").val();
    var description = $("#reg_description").val();

    var connectionName = $("#regConnectionName").val();

    // Health Check
    var hcProtocol = $("#hc_protocol").val();
    var hcport = $("#hc_port").val();
    var hcThreshold = $("#hc_threshold").val();
    var hcInterval = $("#hc_interval").val();
    var hcTimeout = $("#hc_timeout").val();

    // Listener
    var lsProtocol = $("#ls_protocol").val();
    var lsIp = $("#ls_ip").val();
    var lsPort = $("#ls_port").val();
    var lsDnsName = $("#ls_dnsName").val();

    // Target Group
    var tgVms = $("#tg_vms").val();
    var tgProtocol = $("#tg_protocol").val();
    var tgPort = $("#tg_port").val();

    var vNetId = $("#regVNetId").val();
    var mcisId = $("#tg_mcisId").val();
    /*
    console.log("info param cspNlbName : ", cspNlbName);
    console.log("info param description : ", description);

    console.log("info param connectionName : ", connectionName);

    // Health Check
    console.log("info param hcProtocol : ", hcProtocol);
    console.log("info param hcport : ", hcport);
    console.log("info param hcThreshold : ", hcThreshold);
    console.log("info param hcInterval : ", hcInterval);
    console.log("info param hcTimeout : ", hcTimeout);

    // Listener
    console.log("info param lsProtocol : ", lsProtocol);
    console.log("info param lsIp : ", lsIp);
    console.log("info param lsPort : ", lsPort);
    console.log("info param lsDnsName : ", lsDnsName);

    // Target Group
    console.log("info param tgVms : ", tgVms);
    console.log("info param tgProtocol : ", tgProtocol);
    console.log("info param tgPort : ", tgPort);

    console.log("info param vNetId : ", vNetId);
*/
    if (!nlbName) {
        alert("Input New NLB Name");
        $("#reg_name").focus();
        return;
    }
    if (!connectionName) {
        alert("Input Connection Name");
        $("#regConnectionName").focus();
        return;
    }

    var url = "/operation/services/mcis/" + mcisId + "/nlb/reg";
    console.log("nlb URL : ", url);

    var healthCheckerObj = {
        protocol: hcProtocol,
        port: hcport,
        threshold: hcThreshold,
        interval: hcInterval,
        timeout: hcTimeout,
    };
    var listenerObj = {
        protocol: lsProtocol,
        ip: lsIp,
        port: lsPort,
        dnsName: lsDnsName,
    };

    var targetGroupObj = {
        protocol: tgProtocol,
        port: tgPort,
        vms: tgVms.split(","),
    };

    var obj = {
        connectionName: connectionName,
        description: description,
        healthChecker: healthCheckerObj,
        listener: listenerObj,
        name: nlbName,
        scope: "REGION",
        targetGroup: targetGroupObj,
        type: "PUBLIC",
        vNetId: vNetId,
    };
    /*var nlbojb = {}
    var listenerobj = {

    }
    nlbojb.listenre = listenerobj
    var obj = {
        name: cspNlbName,
        connectionName: connectionName,
        targetGroup: {
            asfasf = afasfasf,
        }

    };
    */
    console.log("info connectionconfig obj Data : ", obj);
    if (nlbName) {
        axios
            .post(url, obj, {
                // headers: {
                //     'Content-type': "application/json",
                // },
            })
            .then((result) => {
                console.log(result);
                if (result.status == 200 || result.status == 201) {
                    commonAlert("Success Create Nlb ");
                    //등록하고 나서 화면을 그냥 고칠 것인가?
                    displayNlbInfo("REG_SUCCESS");
                } else {
                    commonAlert("Fail Create Nlb ");
                }
            })
            .catch((error) => {
                console.warn(error);
                console.log(error.response);
                var errorMessage = error.response.statusText;
                var statusCode = error.response.status;
                commonErrorAlert(statusCode, errorMessage);
            });
    } else {
        commonAlert("Input SSH Key Name");
        $("#regCspSshKeyName").focus();
        return;
    }
}

// updateNlb : TODO : update를 위한 form을 만들 것인가 ... 기존 detail에서 enable 시켜서 사용할 것인가
function updateNlb() {
    var sshKeyId = $("#dtlSshKeyId").val();
    var sshKeyName = $("#dtlSshKeyName").val();
    var cspSshKeyId = $("#dtlCspSshKeyId").val();
    var cspSshKeyName = $("#dtlCspSshKeyName").val();
    var description = $("#dtlDescription").val();
    var publicKey = $("#dtlPublicKey").val();
    var privateKey = $("#dtlPrivateKey").val();
    var fingerprint = $("#dtlFingerprint").val();
    var username = $("#dtlUsername").val();
    var verifiedUsername = $("#dtlVerifiedUsername").val(); // TODO : 사용자 이름 입력 받아야 함.
    var connectionName = $("#dtlConnectionName").val();

    console.log("info param cspSshKeyName : ", cspSshKeyName);
    console.log("info param connectionName : ", connectionName);

    var url = "/setting/resources" + "/sshkey/del/" + sshKeyId;
    console.log("ssh key URL : ", url);
    var obj = {
        connectionName: connectionName,
        id: sshKeyId,
        name: sshKeyName,
        cspSshKeyId: cspSshKeyId,
        cspSshKeyName: cspSshKeyName,
        description: description,
        privateKey: privateKey,
        publicKey: publicKey,
        fingerprint: fingerprint,
        username: username,
        verifiedUsername: verifiedUsername,
    };
    console.log("info updateNlb obj Data : ", obj);
    if (cspSshKeyName) {
        axios
            .post(url, obj, {
                headers: {
                    "Content-type": "application/json",
                },
            })
            .then((result) => {
                console.log(result);
                if (result.status == 200 || result.status == 201) {
                    commonAlert(" Nlb Modification Success");
                    displayNlbInfo("REG_SUCCESS"); // TODO : MODIFY 성공일 때 어떻게 처리 할지 정의해서 보완할 것.
                } else {
                    commonAlert("Fail Create Nlb");
                }
            })
            .catch((error) => {
                console.warn(error);
                console.log(error.response);
                var errorMessage = error.response.statusText;
                var statusCode = error.response.status;
                commonErrorAlert(statusCode, errorMessage);
            });
    } else {
        commonAlert("Input SSH Key Name");
        $("#regCspSshKeyName").focus();
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
                    console.log("reg ok button click");
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

function getMcisVmsPop() {
    var regVNetId = $("#regVNetId").val();
    getCommonMcisList("nlbreg", true, "", "vnet=" + regVNetId);
}

// MCIS 목록 조회 후 화면에 Set
function getMcisListCallbackSuccess(caller, mcisList) {
    totalMcisListObj = mcisList;
    console.log("caller:", caller);
    console.log("total mcis:", totalMcisListObj);

    // console.log("caller = " + caller)
    if (caller == "nlbreg") {
        console.log("VMS 체크");
        console.log(mcisList);
        // 리스트 화면 구현 append

        // 팝업 show
        displayMcisVmsListTable(mcisList);
    } else if ((caller = "nlbOnload")) {
        for (var i = 0; i < mcisList.length; i++) {
            mcisID = mcisList[i];
            getNLBList(mcisID);
        }
        // for문으로 getNLBList 가져오기, mcisList에 있는 mcis id만 추출해서 list 조회
    }
    AjaxLoadingShow(false);
}

function getMcisListCallbackFail(caller, error) {
    // List table에 no data 표시? 또는 조회 오류를 표시?
    var addMcis = "";
    addMcis += "<tr>";
    addMcis += '<td class="overlay hidden" data-th="" colspan="8">No Data</td>';
    addMcis += "</tr>";
    $("#mcisList").empty();
    $("#mcisList").append(addMcis);
}

function displayVmRegModal(isShow) {
    if (isShow) {
        $("#vmRegisterBox").modal();
        $(".dtbox.scrollbar-inner").scrollbar();
    } else {
        $("#nlbCreateBox").toggleClass("active");
    }
}

function displayMcisVmsListTable(mcisList) {
    if (!isEmpty(mcisList) && mcisList.length > 0) {
        //totalMcisCnt = mcisList.length;
        var addMcis = "";
        for (var mcisIndex in mcisList) {
            var aMcis = mcisList[mcisIndex];
            if (aMcis.id != "") {
                addMcis += setMcisListTableRow(aMcis, mcisIndex);
            }
        } // end of mcis loop
        $("#subGroupList").empty();
        $("#subGroupList").append(addMcis);

        displayVmRegModal(true);
    } else {
        var addMcis = "";
        addMcis += "<tr>";
        addMcis +=
            '<td class="overlay hidden" data-th="" colspan="8">No Data</td>';
        addMcis += "</tr>";
        $("#mcisList").empty();
        $("#mcisList").append(addMcis);
    }
}

function setMcisListTableRow(aMcisData, mcisIndex) {
    var html = "";
    try {
        var vmList = aMcisData.vm;
        for (var i = 0; i < vmList.length; i++) {
            html +=
                "<tr>" +
                '<td class="column-50px" class="overlay hidden" data-th="">' +
                '<input type="hidden" name="mcisId" id="mcisId_' +
                i +
                '" value="' +
                aMcisData.id +
                '"/>' +
                '<input type="hidden" name="vms_info" id="vms_info_' +
                i +
                '" value="' +
                vmList[i].id +
                '"/>' +
                '<input type="checkbox" name="chk" value="' +
                vmList[i].id +
                '" id="raw_' +
                i +
                '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' +
                '<td class="btn_mtd ovm" data-th="name">' +
                aMcisData.name +
                "</td>" +
                '<td class="overlay hidden" data-th="subGroupId">' +
                vmList[i].subGroupId +
                "</td>" +
                '<td class="overlay hidden" data-th="vmName">' +
                vmList[i].name +
                "</td>" +
                "</tr>";
        }
    } catch (e) {
        console.log("list of mcis error");
        console.log(e);

        html = "<tr>";
        html +=
            '<td class="overlay hidden" data-th="" colspan="4">No Data</td>';
        html += "</tr>";
    }
    return html;
}

function applySubGroup() {
    // 선택한 녀석들의 VM LIST 정보를 VMS칸(tg_vms)에다가 입력
    var subgrouplistValue = $("input[name='vms_info']").length;
    var chk = $("input[name='chk']");
    var subgrouplistData = new Array(subgrouplistValue);
    var mcisid = "";
    var infoshow = "";

    for (var i = 0; i < subgrouplistValue; i++) {
        if (chk[i].checked) {
            mcisid = $("input[name='mcisId']")[i].value;
            subgrouplistData[i] = $("input[name='vms_info']")[i].value;
            console.log("subgrouplistData" + [i] + " : ", subgrouplistData[i]);
            if (infoshow != "") {
                infoshow += ", ";
            }
            infoshow += subgrouplistData[i];
        }
    }

    $("#tg_vms").empty();
    $("#tg_mcisId").val(mcisid);

    $("#tg_vms").val(infoshow);
    $("#vmRegisterBox").modal("hide");
}
