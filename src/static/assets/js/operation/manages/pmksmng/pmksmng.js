$(document).ready(function () {
    setTableHeightForScroll("pmksListTable", 700);

    getPmksList("onload");
});

function ModalClusterDetail(){
    $(".dashboard .status_list tbody tr").each(function(){
    var $td_list = $(this),
            $status = $(".server_status"),
            $detail = $(".server_info");
    $td_list.off("click").click(function(){
            $td_list.addClass("on");
            $td_list.siblings().removeClass("on");
            $status.addClass("view");
            $status.siblings().removeClass("on");
            $(".dashboard.register_cont").removeClass("active");
        $td_list.off("click").click(function(){
                if( $(this).hasClass("on") ) {
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

// 모든 PMKS 목록 조회
function getPmksList(caller){
    var url = "/operation/manages/pmks/list"
    axios.get(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get Cluster List : ", result.data);

        var data = result.data.PmksList;
        getPmksListSuccess(caller, data)
    }).catch(error => {
        console.log(error);
    });
}
// e.GET("/operation/manages/pmks/list", controller.GetPmksList)
// 	e.GET("/operation/manages/pmks/:clusterID", controller.GetPmksInfoData)
// 	e.POST("/operation/manages/pmks/cluster", controller.PmksRegProc)
// 	e.DELETE("/operation/manages/pmks/:clusterID", controller.PmksDelProc)
// 	e.PUT("/operation/manages/pmks/:clusterID", controller.PmksClusterUpdateProc)

// 	e.POST("/operation/manages/pmks/:clusterID/nodegroup", controller.PmksNodeGroupRegProc)
// 	e.DELETE("/operation/manages/pmks/:clusterID/nodegroup/:nodeGroupID", controller.PmksNodeGroupDelProc)

// 모든 PMKS 목록 조회
function getPmks(clusterID, connectionName){
    var url = "/operation/manages/pmks/" + clusterID + "?connectionName=" + connectionName
    axios.get(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get Cluster  : ", result.data);

        var data = result.data.PmksInfo;
        getPmksSuccess(clusterID, data)
    }).catch(error => {
        console.log(error);
    });
}

function getPmksListSuccess(caller, data){
    if ( caller == "onload"){
        var html = "";

        if (data.length) {
            data.filter((list) => list.name !== "").map((item, index) => (
                html += addClusterTableRow(item, index))
            );                
            $("#clusterList").empty();
            $("#clusterList").append(html);
            ModalClusterDetail();
        } else {
            html += CommonTableRowNodata(8);
            $("#clusterList").empty();
            $("#clusterList").append(html);
        }
    }else if ( caller == "nodegrouplist"){
        var nodeGroupList = data.NodeGroupList;
        console.log(nodeGroupList);
        var clusterID = $("#pmks_uid").val();
        var html = "";

        if (nodeGroupList.length) {
            nodeGroupList.filter((list) => list.name !== "").map((item, index) => (
                html += addNodeGroupRow(item, clusterID, index))
            );                
            $("#nodegroupList").empty();
            $("#nodegroupList").append(html);
            ModalNodeGroupDetail();
        } else {
            html += CommonTableRowNodata(8);
            $("#nodegroupList").empty();
            $("#nodegroupList").append(html);
        }
    }
}

// Cluster 단건조회 data = PmksInfo
function getPmksSuccess(clusterID, data){
    //console.log(data);
    var nodeGroupList = data.NodeGroupList;
        console.log(nodeGroupList);
        //var clusterID = $("#pmks_uid").val();
        var html = "";

        if (nodeGroupList.length) {
            nodeGroupList.filter((list) => list.name !== "").map((item, index) => (
                html += addNodeGroupRow(item, index, clusterID))
            );                
            $("#nodegroupList").empty();
            $("#nodegroupList").append(html);
            ModalNodeGroupDetail();
        } else {
            html += CommonTableRowNodata(8);
            $("#nodegroupList").empty();
            $("#nodegroupList").append(html);
        }
}

// Cluster table 에 Row 추가
function addClusterTableRow(item, index){
    // spider를 직접 호출하기 때문에 가져오는 data형태가 좀 다름.
    var html = "";    

    // nodegroup은 IId(소문자), Vpc는 IID(대문자임...)
    var clusterID = item.IId.NameId;
    var clusterSystemID = item.IId.SystemId;
    var clusterVersion = item.Version;
    var clusterConnectionName = item.ConnectionName;
    var clusterStatus = item.Status;
    var clusterNetwork = item.Network;
    var vpcId = item.Network.VpcIID.NameId;
    var vpcSystemId = item.Network.VpcIID.SystemId;
    var subnetIds = new Array();
    var subnetSystemIds = new Array();
    var securityGroupIds = new Array();
    var securityGroupSystemIds = new Array();
    var nodeGroupIds = new Array();
    for (var i = 0; i < clusterNetwork.SubnetIIDs.length; i++) {
        var subnet = clusterNetwork.SubnetIIDs[i];
        subnetIds.push(subnet.NameId)
        subnetSystemIds.push(subnet.SystemId)
    }
    for (var i = 0; i < clusterNetwork.SecurityGroupIIDs.length; i++) {
        var securityGroup = clusterNetwork.SecurityGroupIIDs[i];
        securityGroupIds.push(securityGroup.NameId)
        securityGroupSystemIds.push(securityGroup.SystemId)
    }
    for (var i = 0; i < item.NodeGroupList.length; i++) {
        var nodeGroup = item.NodeGroupList[i];
        var nodeGroupName = nodeGroup.IId.NameId
        nodeGroupIds.push(nodeGroupName)

        //
        //nodeGroupName(nodeCount)
        //var nodeList = nodeGroup.Nodes;//SpIIDList
    }

    
     
    html +=
        "<tr onclick=\"clickListOfPmks('" + clusterID + "', " + index + ");\">" 
        + '<td class="overlay hidden column-50px" data-th="">' 
        + '<input type="hidden" id="cluster_info_' + index + '" value="' + clusterID + '"/>'
        + '<input type="hidden" id="cluster_systemid_' + index + '" value="' + clusterSystemID + '"/>'
        + '<input type="hidden" id="cluster_version_' + index + '" value="' + clusterVersion + '"/>'
        + '<input type="hidden" id="cluster_status_' + index + '" value="' + clusterStatus + '"/>'
        + '<input type="hidden" id="cluster_connection_' + index + '" value="' + clusterConnectionName + '"/>'
        

        + '<input type="hidden" id="network_vpc_nameid_' + index + '" value="' + vpcId + '"/>'
        + '<input type="hidden" id="network_vpc_systemid_' + index + '" value="' + vpcSystemId + '"/>'
        + '<input type="hidden" id="network_subnet_nameid_' + index + '" value="' + subnetIds + '"/>'
        + '<input type="hidden" id="network_subnet_systemid_' + index + '" value="' + subnetSystemIds + '"/>'
        + '<input type="hidden" id="network_securitygroup_nameid_' + index + '" value="' + securityGroupIds + '"/>'
        + '<input type="hidden" id="network_securitygroup_systemid_' + index + '" value="' + securityGroupSystemIds + '"/>'

        + '<input type="hidden" id="cluster_nodegrouplist_' + index + '" value="' + nodeGroupIds.join() + '"/>'


        + '<input type="checkbox" name="clusterchk" value="' + clusterID + '" id="raw_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' 

        + '<td class="btn_mtd ovm" data-th="status">' + clusterStatus + "</td>" 
        + '<td class="overlay hidden" data-th="clusterName">' + clusterID + "</td>" 
        + '<td class="btn_mtd ovm" data-th="provider">' + item.Provider + "</td>" 
        
        + '<td class="overlay hidden" data-th="vpc">' + vpcId + "</td>" 
        + '<td class="overlay hidden" data-th="subnet">' +  subnetIds  + "</td>"         
        + '<td class="overlay hidden" data-th="securityGroup">' +  securityGroupIds  + "</td>"         
        + '<td class="overlay hidden" data-th="nodeGroup">' +  nodeGroupIds  + "</td>"         
        +"</tr>"        
    return html
}

// NodeGroup Table : nodegroupList
function addNodeGroupRow(item, nodeIndex, clusterID){
    console.log("addNodeGroupRow")
    console.log(item)
    //cluster_nodegrouplist_' + index + '
    var html = "";    

    var nodeGroupID = item.IId.NameId;
    var vmSpecName = item.VMSpecName;// IID 형태가 아님.
    var imageID = item.ImageIID.NameId;
    
    var keyPairID = item.KeyPairIID.NameId;
    var desiredNodeSize = item.DesiredNodeSize;
    var maxNodeSize = item.MaxNodeSize;
    var minNodeSize = item.MinNodeSize;
    var nodes = item.Nodes;
    var onAutoScaling = item.OnAutoScaling;
    var rootDiskSize = item.RootDiskSize;
    var rootDiskType = item.RootDiskType;
    var status = item.Status;
    

    var nodeIds = new Array();
    for (var i = 0; i < item.Nodes.length; i++) {
        var node = item.Nodes[i];
        var nodeName = node.NameId
        nodeIds.push(nodeName)       
    }

     
    // html +=
    //     "<tr onclick=\"clickNodeGroup('" + clusterID + ",' + '" + nodeGroupID + ", " + nodeIndex + ");\">" 
    //     + '<td class="overlay hidden column-50px" data-th="">' 
    //     + '<input type="hidden" id="nodegroup_info_' + nodeIndex + '" value="' + nodeGroupID + '"/>'
    //     // + '<input type="hidden" id="nodegroup_systemid_' + nodeIndex + '" value="' + nodeGroupSystemID + '"/>'
    //     + '<input type="hidden" id="nodegroup_status_' + nodeIndex + '" value="' + status + '"/>'

    //     + '<input type="hidden" id="nodegroup_image_nameid_' + nodeIndex + '" value="' + imageID + '"/>'
    //     //+ '<input type="hidden" id="nodegroup_image_systemid_' + nodeIndex + '" value="' + imageSystemId + '"/>'
    //     + '<input type="hidden" id="nodegroup_vmspecname_' + nodeIndex + '" value="' + vmSpecName + '"/>'
    //     + '<input type="hidden" id="nodegroup_keypair_nameid_' + nodeIndex + '" value="' + keyPairID + '"/>'
        
        
    //     + '<input type="checkbox" name="nodegroupchk" value="' + nodeGroupID + '" id="raw_' + nodeIndex + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' 
    //     + '<td class="btn_mtd ovm" data-th="status">' + status + "</td>" 
    //     + '<td class="overlay hidden" data-th="nodeGroupName">' + nodeGroupID + "</td>" 
        
    //     + '<td class="overlay hidden" data-th="nodeGroupImage">' + imageID + "</td>" 
    //     + '<td class="overlay hidden" data-th="nodeGroupSpec">' +  vmSpecName  + "</td>"         
    //     + '<td class="overlay hidden" data-th="nodeGroupKeyPair">' +  keyPairID  + "</td>"         
    //     + '<td class="overlay hidden" data-th="nodeSize">' + minNodeSize + ' / ' +  desiredNodeSize + ' / ' +  maxNodeSize + "</td>"
    //     + '<td class="overlay hidden" data-th="nodes">' +  nodes  + "</td>"
    //     + '<td class="overlay hidden" data-th="onAutoScaling">' +  onAutoScaling  + "</td>"
    //     + '<td class="overlay hidden" data-th="rootdisk">' +  rootDiskType + ' / ' + rootDiskSize  + "</td>"
    //     +"</tr>"        
    // return html
}

function clickListOfPmks(uid, clusterIndex) {
    console.log("click view pmks id :", uid)
    $(".server_status").addClass("view");

    $("[id^='server_info_tr_']").each(function () {
        var item = $(this).attr("item").split("|")
        // console.log(item)
        if (id == item[0]) {
            $(this).addClass("on")
        } else {
            $(this).removeClass("on")
        }
    })


    // Name, version
    $("#pmks_uid").val(uid);// hidden    
    
    // 해당 cluster의 NodeGroup정보 표시
    showNodeGroupListAndStatusArea(uid, clusterIndex);
}


// PMKS Info area 안의 NodeList 내용 표시
// 해당 PMKS를 조회하여 NodeGroup 상세정보 표시
function showNodeGroupListAndStatusArea(uid, clusterIndex) {

    $(".server_status").addClass("view")

    // PMKS Info
    $("#pmks_info_txt").text("[ " + uid + " ]");

    // cluster 정보
    var connectionName = $("#cluster_connection_" + clusterIndex).val()
    connectionName = "ali-test-conn";// for the test

    $("#pmks_uid").val(uid);
    $("#pmks_info_version").val($("#cluster_version_" + clusterIndex).val());
    $("#pmks_info_name").val($("#cluster_info_" + clusterIndex).val() + " / " + $("#cluster_systemid_" + clusterIndex).val());
    $("#pmks_info_status").val($("#cluster_status_" + clusterIndex).val());
    $("#pmks_info_cloud_connection").val(connectionName);
    
    // Network
    $("#pmks_vpc").val($("#network_vpc_nameid_" + clusterIndex).val() + " / " + $("#network_vpc_systemid_" + clusterIndex).val());
    $("#pmks_subnet").val($("#network_subnet_nameid_" + clusterIndex).val());
    $("#pmks_info_security_group").val($("#network_securitygroup_nameid_" + clusterIndex).val());
    
    // NodeGroupList :     
    getPmks(uid, connectionName)
//     var nodeGroupList = "";
//     addNodeGroupRow(item, clusterID, clusterIndex, nodeIndex)

//     // Node Detail


// // $("#pmks_server_info_status").empty();
// // $("#pmks_server_info_status").append('<strong>NodeGroup List </strong>  <span class="stxt">[ ' + pmksName + ' ]</span>  Node(' + nodeTotalCountOfPmks + ')')


//     //$("#pmks_name").val(pmksName)

//     var pmksNodes = "";
//     //var pmksStatusIcon = "";
//     $("[id^='pmksNodeUID_']").each(function () {
//         var pmksNode = $(this).attr("id").split("_")
//         thisPmksIndex = pmksNode[1]
//         nodeIndexOfPmks = pmksNode[2]

//         if (thisPmksIndex == pmksIndex) {
//             var nodeID = $("#pmksNodeUID_" + thisPmksIndex + "_" + nodeIndexOfPmks).val();
//             var nodeName = $("#pmksNodeName_" + thisPmksIndex + "_" + nodeIndexOfPmks).val();

//             //nodeStatusIcon ="bgbox_g"
//             nodeStatusIcon = "bgbox_b"
//             // node 목록 표시
//             pmksNodes += '<li class="sel_cr ' + nodeStatusIcon + '"><a href="javascript:void(0);" onclick="nodeDetailInfo(\'' + thisPmksIndex + '\',\'' + nodeIndexOfPmks + '\')"><span class="txt">' + nodeName + '</span></a></li>';
//         }
//     });
//     $("#pmks_server_info_box").empty();
//     $("#pmks_server_info_box").append(pmksNodes);


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
