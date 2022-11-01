// PMKS의 NodeGroup을 추가
// Connection, ClusterID는 넘어오며 변경 불가.
$(document).ready(function() {

    // nodegroup 추가영역 표시
    addNewNodeGroupForm();

    // cluster info 영역 보이게 - 공통으로 사용하기 때문에 hidden으로 설정되어 있어 보이도록 처리.
    $status = $(".server_status");
    $status.addClass("view");

    // PMKS 조회
    var clusterID = $("#pmks_cluster_id").val();
    var connectionName = $("#pmks_cluster_connection").val();
    getCommonPmksData("pmks_nodegroup_onload", clusterID, connectionName)
})

// PMKS Data 조회 결과
function getPmksDataSuccess(caller, clusterID, data){
    console.log("caller", caller)
    console.log("clusterID", clusterID)
    console.log("data", data)

    // Cluster Info Set
    setClusterInfo(data);
    // NodeGroup이 있으면 NodeGroup icon을 표시? List로 표시?
    if( data.NodeGroupList != null){
        setNodeGroupList(data.NodeGroupList)
    }    
}

// 해당 data에 Key가 있으면 해당 값을 return 없으면 ""
function getCommonStringValue(data, key){
    // ojbect의 parame 조회
}


function setClusterInfo(data){
    var clusterID = data.IId.NameId;
    var clusterVersion = data.Version;
    var clusterStatus = data.Status;
    var connectionName = data.ConnectionName;
    connectionName = "ali-test-conn";// for the test
    // PMKS Info
    $("#pmks_cluster_id").val(clusterID);// hidden        
    $("#pmks_cluster_name").val(clusterID);// hidden        
    $("#pmks_info_txt").text("[ " + clusterID + " ]");
    $("#pmks_cluster_connection").val(connectionName);

    // Network 영역    
    if ( data.Network != null ){
        networkInfo = data.Network;
        $("#pmks_info_name").val(clusterID);
        $("#pmks_info_version").val(clusterVersion);
        $("#pmks_info_cloud_connection").val(connectionName);
        $("#pmks_info_status").val(clusterStatus);
        // provider
        // region
        $("#pmks_info_vpc").val(networkInfo.VpcIID.NameId);
        $("#pmks_info_subnet").val(networkInfo.SubnetIIDs);
        $("#pmks_info_security_group").val(networkInfo.SecurityGroupIIDs);
    }

}

// 조회 된 
//data == nodeGroupList

function setNodeGroupList(data){
    var html = "";
    // Cluster만 있는 경우 NodeGroup이 없을 수 있음
    if (nodeGroupList != null ){
        if (nodeGroupList.length) {
            var nodeGroupID = "" 
            for (var i in nodeGroupList) {
                nodeGroupInfo = nodeGroupList[i];
                nodeGroupID = nodeGroupInfo.IId.NameId;
                html += addNodeGroupData(nodeGroupInfo, i, clusterID);
                
            }
            $("#pmks_nodegroup_list_info_box").empty();
            $("#pmks_nodegroup_list_info_box").append(html);

            if(nodeGroupList.length == 1){
                setNodeList(clusterID, nodeGroupID);
            }
            
        } else {
            html += CommonTableRowNodata(8);
            $("#pmks_nodegroup_list_info_box").empty();
            $("#pmks_nodegroup_list_info_box").append(html);
        }
    }

    displayNodeGroupListArea();
}

// 
function addNodeGroupData(item, nodeGroupIndex, clusterID){
    console.log("addNodeGroupData")
    console.log(item)
    
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
    var nodeGroupStatusIcon = getNodeGroupStatusIcon(status);
    var nodeGroupDispStatus = getNodeGroupStatusDisp(status);
    var nodeGroupDispClass = getNodeGroupStatusClass(status);

    html += '<ul>';
    html += '   <li>';
    html += '       <label>Name</label>';
    html += '       <input type="text" name="" value="' + nodeGroupID +  '" placeholder="" title="" readonly />';
    html += '   </li>';
    html += '   <li>';
    html += '       <label>Image ID</label>';
    html += '       <input type="text" name="" value="' + imageID +  '" placeholder="" title="" readonly />';
    html += '   </li>';
    html += '   <li>';
    html += '       <label>VM Spec</label>';
    html += '       <input type="text" name="" value="' + vmSpecName +  '" placeholder="" title="" readonly />';
    html += '   </li>';
    html += '   <li>';
    html += '   <label>KeyPair</label>';
    html += '   <input type="text" name="" value="' + keyPairID +  '" placeholder="" title="" readonly />';
    html += '   </li>';
    html += '   <li>';
    html += '   <label>Desired Node Size</label>';
    html += '   <input type="text" name="" value="' + desiredNodeSize +  '" placeholder="" title="" readonly />';
    html += '   </li>';
    html += '   <li>';
    html += '   <label>Max Node Size</label>';
    html += '   <input type="text" name="" value="' + maxNodeSize +  '" placeholder="" title="" readonly />';
    html += '   </li>';
    html += '<li>';
    html += '<label>Min Node Size</label>';
    html += '<input type="text" name="" value="' + minNodeSize +  '" placeholder="" title="" readonly />';
    html += '</li>';
    html += '<li>';
    html += '<label>On Auto Scaling</label>';
    html += '<input type="text" name="" value="' + onAutoScaling +  '" placeholder="" title="" readonly />';
    html += '</li>';
    html += '<li>';
    html += '<label>Root Disk</label>';
    html += '<input type="text" name="" value="' + rootDiskType + ' / ' + rootDiskSize +  '" placeholder="" title="" readonly />';
    html += '</li>';  								
    html += "</ul>";
    
    
    //var nodeIds = new Array();
    //for (var i = 0; i < item.Nodes.length; i++) {
    //    var node = item.Nodes[i];
    //    var nodeName = node.NameId
    //    nodeIds.push(nodeName)       
    //}
    //console.log("li attrt class " + nodeGroupStatusIcon)

    //var nodeGroupBadge = '<li class="_sel_cr_ ' + nodeGroupDispClass + '" id="nodeGroupOfCluster_' + nodeGroupIndex + '"><a href="javascript:void(0);" onclick="clickListOfNodeGroup(\'' + clusterID + '\',\'' + nodeGroupID + '\')"><span class="txt">' + nodeGroupID + '</span></a></li>';

    return html;
}

function displayNodeGroupListArea(){
    $("#pmks_nodegroup_detail_info_box").addClass("view")

    $("[id^='server_info_tr_']").each(function () {
        var item = $(this).attr("item").split("|")
        // console.log(item)
        if (id == item[0]) {
            $(this).addClass("on")
        } else {
            $(this).removeClass("on")
        }
    })
    $(".server_status").addClass("view")

}


// 새로운 NodeGroup을 입력하는 Form 추가
function addNewNodeGroupForm(){
    // 기존 nodeGroup의 index + 1
    var newNodeGroupIndex = 0;
    $("[name='nodegroup_idx']").each(function (idx, ele) {
        // 중간에 이빨이 빠진게 생기면 index는 max를 가리키도록
        var addedId = $(this).attr("id")
        var toAddId = "nodegroup_idx_" + newNodeGroupIndex;
        if( addedId != toAddId){
            newNodeGroupIndex = addedId
        }
        newNodeGroupIndex++;
	});

    var html = ""
    //"Name": "Economy", 
    //"VMSpecName": "ecs.c6.xlarge", "RootDiskType": "cloud_essd", "RootDiskSize": "70", "KeyPairName": "keypair-01",
    //"OnAutoScaling": "true", "DesiredNodeSize": "2", "MinNodeSize": "2", "MaxNodeSize": "2" 
    html += '<div class="servers_box" name="nodegroup_idx" id="nodegroup_idx_' + newNodeGroupIndex + '">';
    html += '<div class="list">';
    html += '<ul>';
    html += '   <li>';
    html += '       <label>Name</label>';
    html += '       <input type="text" name="nodegroup_info_name" value="" id="nodegroup_info_name_' + newNodeGroupIndex + ' placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '       <label>Image ID</label>';
    html += '       <input type="text" name="nodegroup_info_imageid" value="" id="nodegroup_info_imageid_' + newNodeGroupIndex + ' placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '       <label>VM Spec</label>';
    html += '       <input type="text" name="nodegroup_info_vmspecid" value="" id="nodegroup_info_vmspecid_' + newNodeGroupIndex + ' placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '   <label>KeyPair</label>';
    html += '       <input type="text" name="nodegroup_info_keypairid" value="" id="nodegroup_info_keypairid_' + newNodeGroupIndex + ' placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '   <label>Desired Node Size</label>';
    html += '       <input type="text" name="nodegroup_info_desired_nodesize" value="" id="nodegroup_info_desired_nodesize_' + newNodeGroupIndex + ' placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '   <label>Max Node Size</label>';
    html += '       <input type="text" name="nodegroup_info_max_nodesize" value="" id="nodegroup_info_max_nodesize_' + newNodeGroupIndex + ' placeholder="" title="" />';
    html += '   </li>';
    html += '<li>';
    html += '<label>Min Node Size</label>';
    html += '       <input type="text" name="nodegroup_info_min_nodesize" value="" id="nodegroup_info_min_nodesize_' + newNodeGroupIndex + ' placeholder="" title="" />';
    html += '</li>';
    html += '<li>';
    html += '<label>On Auto Scaling</label>';
    html += '       <input type="text" name="nodegroup_info_onautoscaling" value="" id=nodegroup_info_onautoscaling_"' + newNodeGroupIndex + ' placeholder="" title="" />';
    html += '</li>';
    html += '<li>';
    html += '<label>Root Disk</label>';
    html += '       <input type="text" name="nodegroup_info_rootdisk_type" value="" id="nodegroup_info_rootdisk_type_' + newNodeGroupIndex + ' placeholder="type" title="" />';
    html += '       <input type="text" name="nodegroup_info_rootdisk_size" value="" id="nodegroup_info_rootdisk_size_' + newNodeGroupIndex + ' placeholder="size" title="" />';
    html += '</li>';  								
    html += '</ul>';
    html += '</div>';
    if( newNodeGroupIndex > 0 ){
    html += '<div class="btnbox">';
	html += '		<div class="btn_right">';
	html += '			<button type="button" id="nodegroup_info_remove_' + newNodeGroupIndex + '" value="" class="btn_done btn_del" onclick="removeNodeGroupForm(' + newNodeGroupIndex+');"></button>';
	html += '		</div>';
	html += '	</div>';
    }
    html += '</div>';

    //addNodeGroupArea 에 append
    $("#addNodeGroupArea").append(html)    
}

// NodeGroup 입력 form 제거
function removeNodeGroupForm(nodeGroupIndex){
    console.log($("#nodegroup_idx_" + nodeGroupIndex))
    
    $('div').remove("#nodegroup_idx_" + nodeGroupIndex)
    //$("#nodegroup_idx_" + nodeGroupIndex).remove();
    
    console.log($("#nodegroup_idx_" + nodeGroupIndex))
    //alert("7")
    //$("#nodegroup_idx_" + nodeGroupIndex).empty();
}