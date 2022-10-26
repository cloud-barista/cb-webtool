$(document).ready(function () {
    setTableHeightForScroll("pmksListTable", 700);

    getAllPmksList();
});

// 모든 PMKS 목록 조회
function getAllPmksList(){
    // connection 목록
}

function clickListOfPmks(uid, pmksIndex) {
    console.log("click view pmks id :", uid)
    $(".server_status").addClass("view");

    // List Of PMKS에서 선택한 row 외에는 안보이게
    $("[id^='server_info_tr_']").each(function () {
        var item = $(this).attr("item").split("|")
        console.log(item)
        if (id == item[0]) {
            $(this).addClass("on")
        } else {
            $(this).removeClass("on")
        }
    })

    $("#pmks_uid").val($("#pmksUID" + pmksIndex).val());
    $("#pmks_name").val($("#pmksName" + pmksIndex).val());

    // PMKS Info area set
    showServerListAndStatusArea(uid, pmksIndex);
}


// PMKS Info area 안의 Node List 내용 표시
// 해당 PMKS의 모든 Node 표시
// TODO : 클릭했을 때 서버에서 조회하는것으로 변경할 것.
function showServerListAndStatusArea(uid, pmksIndex) {

    var pmksUID = $("#pmksUID" + pmksIndex).val();
    var pmksName = $("#pmksName" + pmksIndex).val();
    var pmksStatus = $("#pmksStatus" + pmksIndex).val();
    var pmksConfig = $("#pmksConfig" + pmksIndex).val();
    var nodeTotalCountOfPmks = $("#pmksNodeTotalCount" + pmksIndex).val();

    $(".server_status").addClass("view")
    $("#pmks_info_txt").text("[ " + pmksName + " ]");
    $("#pmks_server_info_status").empty();
    $("#pmks_server_info_status").append('<strong>Node List </strong>  <span class="stxt">[ ' + pmksName + ' ]</span>  Node(' + nodeTotalCountOfPmks + ')')

    //
    $("#pmks_info_name").val(pmksName + " / " + pmksUID)
    $("#pmks_info_Status").val(pmksStatus)
    $("#pmks_info_cloud_connection").val(pmksConfig)

    $("#pmks_name").val(pmksName)

    var pmksNodes = "";
    //var pmksStatusIcon = "";
    $("[id^='pmksNodeUID_']").each(function () {
        var pmksNode = $(this).attr("id").split("_")
        thisPmksIndex = pmksNode[1]
        nodeIndexOfPmks = pmksNode[2]

        if (thisPmksIndex == pmksIndex) {
            var nodeID = $("#pmksNodeUID_" + thisPmksIndex + "_" + nodeIndexOfPmks).val();
            var nodeName = $("#pmksNodeName_" + thisPmksIndex + "_" + nodeIndexOfPmks).val();

            //nodeStatusIcon ="bgbox_g"
            nodeStatusIcon = "bgbox_b"
            // node 목록 표시
            pmksNodes += '<li class="sel_cr ' + nodeStatusIcon + '"><a href="javascript:void(0);" onclick="nodeDetailInfo(\'' + thisPmksIndex + '\',\'' + nodeIndexOfPmks + '\')"><span class="txt">' + nodeName + '</span></a></li>';
        }
    });
    $("#pmks_server_info_box").empty();
    $("#pmks_server_info_box").append(pmksNodes);


    //Manage PMKS Server List on/off : table을 클릭하면 해당 Row 에 active style로 보여주기
    $(".dashboard .ds_cont .area_cont .listbox li.sel_cr").each(function () {
        var $sel_list = $(this);
        var $detail = $(".server_info");
        console.log($sel_list);
        console.log($detail);
        console.log(">>>>>");
        $sel_list.off("click").click(function () {
            $sel_list.addClass("active");
            $sel_list.siblings().removeClass("active");
            $detail.addClass("active");
            $detail.siblings().removeClass("active");
            $sel_list.off("click").click(function () {
                if ($(this).hasClass("active")) {
                    $sel_list.removeClass("active");
                    $detail.removeClass("active");
                } else {
                    $sel_list.addClass("active");
                    $sel_list.siblings().removeClass("active");
                    $detail.addClass("active");
                    $detail.siblings().removeClass("active");
                }
            });
        });
    });
}

// 해당 pmks에 nodeGroup 추가
// pmks가 경로에 들어가야 함. node 등록 form으로 이동
function addNewNodeGroup() {
    //var clusterId = $("#pmks_uid").val(); // pmks id 값이 없음
    var clusterId = $("#pmks_name").val();
    var clusterName = $("#pmks_name").val();

    if (clusterId == "") {
        commonAlert("PMKS 정보가 올바르지 않습니다.");
        return;
    }
    var url = "/operation/manages/pmksmng/regform/" + clusterId + "/" + clusterName;
    location.href = url;
}

// PMKS 삭제
function deletePMKS() {
    var checkedCount = 0;
    var pmksID = "";
    var pmksName = "";
    $("[id^='td_ch_']").each(function () {
        var checkedIndex = $(this).val();
        if ($(this).is(":checked")) {
            checkedCount++;
            console.log("checked")
            pmksID = $("#pmksUID" + checkedIndex).val();
            pmksName = $("#pmksName" + checkedIndex).val();
            // 여러개를 지울 때 호출하는 함수를 만들어 여기에서 호출
        } else {
            console.log("checked nothing")
        }
    })

    if (checkedCount == 0) {
        commonAlert("Please Select PMKS!!")
        return;
    } else if (checkedCount > 1) {
        commonAlert("Please Select One PMKS at a time")
        return;
    }

    // TODO : 삭제 호출부분 function으로 뺼까?
    //var url = "/ns/{namespace}/clusters/{cluster}"
    var url = "/operation/manages/pmksmng/" + pmksID + "/" + pmksName;
    axios.delete(url, {})
        .then(result => {
            console.log("get  Data : ", result.data);
            //StatusInfo.code
            //StatusInfo.kind
            //StatusInfo.message
            var statusCode = result.data.status;
            var message = result.data.message;

            if (statusCode != 200 && statusCode != 201) {
                commonAlert(message + "(" + statusCode + ")");
                return;
            } else {
                commonAlert(message);
                // TODO : PMKS List 조회
                //location.reload();
            }

        }).catch((error) => {
            console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage)
        });

}

function deleteNodeOfPmks() {
    // worker만 삭제
    // 1개씩 삭제

    var selectedPmksUid = $("#pmks_uid").val();
    var selectedPmksName = $("#pmks_name").val();
    var selectedNodeUid = $("#node_uid").val();
    var selectedNodeName = $("#node_name").val();
    var selectedNodeRole = $("#pmks_node_role").val();

    if (selectedNodeRole.toLowerCase() != "worker") {
        commonAlert("Only worker node can be deleted")
        return;
    }

    var urlParamMap = new Map();
    urlParamMap.set(":clusterUID", selectedPmksUid)
    urlParamMap.set(":clusterName", selectedPmksName)
    urlParamMap.set(":nodeID", selectedNodeUid)
    urlParamMap.set(":nodeName", selectedNodeName)
    var url = setUrlByParam("PmksClusterNodeData", urlParamMap)
    console.log("URL : ", url)
    axios.delete(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        // var data = result.data;
        // if (result.status == 200 || result.status == 201) {
        var statusCode = result.data.status;
        if (statusCode == 200 || statusCode == 201) {
            commonAlert("Success Delete Node.");

        } else {
            var message = result.data.message;
            commonAlert("Fail Delete Node : " + message + "(" + statusCode + ")");

        }
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

// 선택한 Node의 상세정보 표시
function nodeDetailInfo(pmksIndex, nodeIndex) {
    var nodeUID = $("#pmksNodeUID_" + pmksIndex + "_" + nodeIndex).val();
    var nodeName = $("#pmksNodeName_" + pmksIndex + "_" + nodeIndex).val();
    var nodeKind = $("#pmksNodeKind_" + pmksIndex + "_" + nodeIndex).val();
    var nodeRole = $("#pmksNodeRole_" + pmksIndex + "_" + nodeIndex).val();

    // hidden 값 setting. 삭제 등에서 사용
    $("#node_uid").val(nodeUID);
    $("#node_name").val(nodeName);

    $("#pmks_node_txt").text(nodeName + " / " + nodeUID);

    $("#pmks_node_name").val(nodeName);
    $("#pmks_node_kind").val(nodeKind);
    $("#pmks_node_role").val(nodeRole);

    $("#pmks_node_detail").css("display", "block");

}
