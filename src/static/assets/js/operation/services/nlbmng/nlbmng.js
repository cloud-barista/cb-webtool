$(document).ready(function () {
    setTableHeightForScroll("nlbList", 300);
    
    getAllNlbList()    
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

        getAllNlbList();
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

        getAllNlbList();
    } else if (targetAction == "CLOSE") {
        $("#nlbCreateBox").removeClass("active");
        $("#nlbInfoBox").removeClass("view");
        $("#nlbListTable").addClass("on");

        var offset = $("#sskKeyInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);
    }
}

function getAllNlbList(){
    var url = "/operation/services/mcis/nlb/listall";
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

            if (data.length) {
                data.filter((list) => list.name !== "").map((item, index) => (
                    html += addNlbListTableRow(item, index))
                );                
                $("#nlbList").empty();
                $("#nlbList").append(html);
                ModalDetail();
            } else {
                html += nodataTableRow(8);
                $("#nlbList").empty();
                $("#nlbList").append(html);
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

// nodata일 때 row 표시 : colspanCount를 받아 col을 합친다.
function nodataTableRow(colspanCount){
    var html = "";
    html += "<tr>";
    html += '<td class="overlay hidden" data-th="" colspan="' + colspanCount + '">No Data</td>';
    html += "</tr>";
    return html
}
// list의 1개 row에 대한 html 
function addNlbListTableRow(item, index) {
    var html = "";
    html +=
        "<tr onclick=\"showNlbInfo('" + item.mcisId + "', '"+ item.id + "');\">" 
        + '<td class="overlay hidden column-50px" data-th="">' 
        + '<input type="hidden" id="nlb_info_' + index + '" value="' + item.mcisId + "|" + item.name + "|" + item.connectionName + '"/>'
        + '<input type="checkbox" name="vmchk" value="' + item.mcisId + "|" + item.name + '" id="raw_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' 
        + '<td class="btn_mtd ovm" data-th="provider">' + item.location.cloudType + "</td>" 
        + '<td class="btn_mtd ovm" data-th="mcis">' + item.mcisId + "</td>" 
        + '<td class="overlay hidden" data-th="nlbName">' + item.name + "</td>" 
        + '<td class="overlay hidden" data-th="scope">' + item.scope + "</td>" 
        + '<td class="overlay hidden" data-th="type">' +  item.type  + "</td>"         
        //+ '<td class="overlay hidden column-80px" data-th="health">' + item.status + "</td>" 
        + '<td class="overlay hidden" data-th="listenerInfo">' + item.listener.protocol + " | " + item.listener.port + "</td>" 
        + '<td class="overlay hidden" data-th="listenerInfo">' + item.targetGroup.protocol + " | " + item.targetGroup.port +"</td>" 
        + '<td class="overlay hidden" data-th="listenerInfo">' + item.targetGroup.subGroupId + "</td>" 
        +"</tr>"
        
    return html
}
// nlb 목록 조회
function getNLBList(mcisID) {
    console.log("mcisID : ", mcisID);
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

            if (data.length) {
                data.filter((list) => list.name !== "").map((item, index) => (
                    html += addNlbListTableRow(item, index, mcisID))
                );
                $("#nlbList").append(html);

                ModalDetail();
            } else {
                html += nodataTableRow(8);
                $("#nlbList").empty();
                $("#nlbList").append(html);
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
    console.log("deleteNlb!!!");
    
    var selValues = "";
    var selMcisId = "";
    var selNlbId = "";
    var count = 0;


    $("input[name='vmchk']:checked").each(function () {
        count++;
        selValues = selValues + $(this).val() + ",";
    });
    console.log("selValues : ", selValues)
    selValues = selValues.substring(0, selValues.lastIndexOf(","));
    
    var splitData = selValues.split("|");
    selMcisId = selMcisId + splitData[0];
    selNlbId = selNlbId + splitData[1];

    console.log("selMcisId : ", selMcisId);
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

    var url = "/operation/services/mcis/" + selMcisId + "/nlb/del/" + selNlbId;
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

function showNlbInfo(mcisId, nlbId) {
    console.log("target showNlbInfo : ", mcisId);

    $(".stxt").html(nlbId);

    var url = "/operation/services/mcis/" + mcisId + "/nlb/" + nlbId;
    console.log("n1lb URL : ", url);

    return axios
        .get(url, {
            headers: {
                "Content-Type": "application/json",
            },
        })
        .then((result) => {
            var data = result.data.NlbInfo;
            console.log("result DT : ", data);
            
            var dtlNlbName = data.name;
            var dtlDescription = data.description;

            var dtlHcThreshold = data.healthChecker.threshold;
            var dtlHcInterval = data.healthChecker.interval;
            var dtlHcTimeout = data.healthChecker.timeout;

            var dtlListenerProtocol = data.listener.protocol;
            var dtlListenerPort = data.listener.port;

            var dtlProvider = data.location.cloudType;
            var dtlConnectionName = data.connectionName;
            //var dtlVNetId = data.keyValueList.keyValueList;

            var dtlTgVms = data.targetGroup.vms;
            strVms = dtlTgVms.join();
            strVms = strVms.replaceAll(",", "\r\n");

            var dtlTgProtocol = data.targetGroup.protocol;
            var dtlTgPort = data.targetGroup.port;
            var dtlTgSubGroupId = data.targetGroup.subGroupId;
            
            $("#dtl_nlbName").empty();
            $("#dtl_description").empty();
            $("#dtl_hc_threshold").empty();
            $("#dtl_hc_interval").empty();
            $("#dtl_hc_timeout").empty();
            $("#dtl_ls_protocol").empty();
            $("#dtl_ls_port").empty();
            $("#dtl_provider").empty();
            $("#dtl_connectionName").empty();
            //$("#dtl_vNetId").empty();
            $("#dtl_tg_vms").empty();
            $("#dtl_tg_protocol").empty();
            $("#dtl_tg_port").empty();
            $("#dtl_tg_subGroupId").empty();

            $("#dtl_nlbName").val(dtlNlbName);
            $("#dtl_description").val(dtlDescription);
            $("#dtl_hc_threshold").val(dtlHcThreshold);
            $("#dtl_hc_interval").val(dtlHcInterval);
            $("#dtl_hc_timeout").val(dtlHcTimeout);
            $("#dtl_ls_protocol").val(dtlListenerProtocol);
            $("#dtl_ls_port").val(dtlListenerPort);
            $("#dtl_provider").val(dtlProvider);
            $("#dtl_connectionName").val(dtlConnectionName);
            //$("#dtl_vNetId").val(dtlVNetId);
            $("#dtl_tg_vms").val(strVms);
            $("#dtl_tg_protocol").val(dtlTgProtocol);
            $("#dtl_tg_port").val(dtlTgPort);
            $("#dtl_tg_subGroupId").val(dtlTgSubGroupId);
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

    var cloudProvider = $("#regProvider").val();
    var connectionName = $("#regConnectionName").val();

    // Health Check
    var hcThreshold = $("#hc_threshold").val();
    var hcInterval = $("#hc_interval").val();
    var hcTimeout = $("#hc_timeout").val();

    // Listener
    var lsProtocol = $("#ls_protocol").val();
    var lsPort = $("#ls_port").val();
    
    // Target Group
    var tgVms = $("#tg_vms").val();
    var tgProtocol = $("#tg_protocol").val();
    var tgPort = $("#tg_port").val();
    var tgSubGroupId = $("#tg_subGroupId").val();

    var vNetId = $("#regVNetId").val();
    var mcisId = $("#tg_mcisId").val();
    
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
        threshold: hcThreshold,
        interval: hcInterval,
        timeout: hcTimeout,
    };

    var listenerObj = {
        protocol: lsProtocol,
        port: lsPort
    };

    var targetGroupObj = {
        protocol: tgProtocol,
        port: tgPort,
        //vms: tgVms.split(","),
        subGroupId: tgSubGroupId
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
        commonAlert("Input NLB Name");
        $("#reg_name").focus();
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

    console.log(regVNetId);

    if(regVNetId == "" || regVNetId == null) {
        commonAlert("VPC ID를 선택해주세요.");
        return;
    } 
    //caller, isCallback, targetObjId, optionParam, filterKeyVal
    getCommonMcisList("nlbreg", true, "", "", "filterKey=vNetId&filterVal=" + regVNetId);
    
}

// MCIS 목록 조회 후 화면에 Set
function getMcisListCallbackSuccess(caller, mcisList) {
    totalMcisListObj = mcisList;
    console.log("caller:", caller);
    console.log("total mcis:", totalMcisListObj);

    // console.log("caller = " + caller)
    if (caller == "nlbreg") {
        console.log(mcisList);
        // 리스트 화면 구현 append

        // 팝업 show
        displayMcisVmsListTable(mcisList);
    } else if ((caller = "nlbOnload")) {
        $("#nlbList").empty();
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
            '<td class="overlay hidden" data-th="" colspan="4">No Data</td>';
        addMcis += "</tr>";
        $("#mcisList").empty();
        $("#mcisList").append(addMcis);
    }
}

function setMcisListTableRow(aMcisData, mcisIndex) {
    var html = "";
    try {
        console.log(aMcisData)
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
                '<input type="hidden" name="subGroupId" id="subGroupId_' +
                i +
                '" value="' +
                vmList[i].subGroupId +
                '"/>' +
                '<input type="hidden" name="vms_info" id="vms_info_' +
                i +
                '" value="' +
                vmList[i].id +
                '"/>' +
                '<input type="checkbox" name="vmchk" value="' +
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
    // 선택한 녀석의 VM LIST 정보를 VMS칸(tg_vms)에다가 입력
    var subgrouplistValue = $("input[name='vms_info']").length;
    var vmchk = $("input[name='vmchk']");
    var subgrouplistData = new Array(subgrouplistValue);
    var mcisId = "";
    var subGroupId = ""
    var vmIdList = ""
    // var infoshow = "";

    // check 된 항목 찾기 : 1개만 가능
    var checkedCount = 0;    
    for (var i = 0; i < subgrouplistValue; i++) {
        if( checkedCount > 1){
            commonAlert("1개만 선택 가능")
            return
        }

        //console.log(i+ " : ", $("input[name='vms_info']")[i].value, vmchk[i].checked);
        if (vmchk[i].checked ) {
            checkedCount++;

            mcisId = $("input[name='mcisId']")[i].value;
            subGroupId = $("input[name='subGroupId']")[i].value;
            
            // subgrouplistData[i] = $("input[name='vms_info']")[i].value;
            // console.log("subgrouplistData" + [i] + " : ", subgrouplistData[i]);
            // if (infoshow != "") {
            //     //infoshow += ", ";
            //     infoshow += "\r\n";
                
            // }
            // infoshow += subgrouplistData[i];
        }
    }

    for (var i = 0; i < subgrouplistValue; i++) {
        tempSubGroupId = $("input[name='subGroupId']")[i].value;
        if( subGroupId == tempSubGroupId){
            vmIdList += $("input[name='vms_info']")[i].value;
            vmIdList += "\r\n";
        }
    }


    $("#tg_vms").empty();
    $("#tg_mcisId").val(mcisId);

    // $("#tg_vms").val(infoshow);
    $("#tg_vms").val(vmIdList);
    $("#tg_subGroupId").val(subGroupId);
    $("#vmRegisterBox").modal("hide");
}

