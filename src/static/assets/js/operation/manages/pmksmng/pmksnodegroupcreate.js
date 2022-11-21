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
    // if( data.NodeGroupList != null){
    //     setNodeGroupList(data.NodeGroupList)
    // }    
}

// 해당 data에 Key가 있으면 해당 값을 return 없으면 ""
function getCommonStringValue(data, key){
    // ojbect의 parame 조회
}

// NodeGroup 추가 : pmksmng에도 이름이 동일한 function 있음.
function setClusterInfo(data){
    var clusterID = data.IId.NameId;
    var clusterVersion = data.Version;
    var clusterStatus = data.Status;
    var connectionName = data.ConnectionName;

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

    // AccessInfo 영역
    if ( data.AccessInfo != null ){
        accessInfo = data.AccessInfo;
        $("#pmks_info_endpoint").val(accessInfo.Endpoint);
        $("#pmks_info_kubeconfig").val(accessInfo.Kubeconfig);
    }

    // NodeGroup 이 있으면NodeGroup 목록 표시
    if (data.NodeGroupList != null){
        var html = "";
        for (var o in data.NodeGroupList) {
            var nodeGroupStatus = data.NodeGroupList[o].Status;
            var nodeGroupName = data.NodeGroupList[o].IId.NameId

            var nodeGroupDispStatus = "";//getNodeGroupStatusDisp(nodeGroupStatus);
            var nodeGroupStatusClass = "bgbox_b";//getNodeGroupStatusClass(nodeGroupName)

            // NodeGroup 생성부분이라 Click Event 는 없음 pmksmng.js 에서는 clickListOfNodeGroup 있음.
            html += '<li id="nodegroup_status_icon_' + o + '" class="sel_cr ' + nodeGroupStatusClass + '"><span class="txt">' + nodeGroupName + '</span></li>';
        }
        $("#cluster_nodegroup_list").empty();
        $("#cluster_nodegroup_list").append(html);        
    }
    
    //connectionName
    getCommonLookupDiskInfo("pmksnodegroup", "", connectionName); // -> getCommonLookupDiskInfoSuccess
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
    html += '       <input type="text" name="nodegroup_info_name" value="" id="nodegroup_info_name_' + newNodeGroupIndex + '" placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '       <label>Image ID</label>';
    html += '       <input type="text" name="nodegroup_info_imageid" value="" id="nodegroup_info_imageid_' + newNodeGroupIndex + '" placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '       <label>VM Spec</label>';
    html += '       <input type="text" name="nodegroup_info_vmspecid" value="" id="nodegroup_info_vmspecid_' + newNodeGroupIndex + '" placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '   <label>KeyPair</label>';
    html += '       <input type="text" name="nodegroup_info_keypairid" value="" id="nodegroup_info_keypairid_' + newNodeGroupIndex + '" placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '   <label>Desired Node Size</label>';
    html += '       <input type="text" name="nodegroup_info_desired_nodesize" value="" id="nodegroup_info_desired_nodesize_' + newNodeGroupIndex + '" placeholder="" title="" />';
    html += '   </li>';
    html += '   <li>';
    html += '   <label>Max Node Size</label>';
    html += '       <input type="text" name="nodegroup_info_max_nodesize" value="" id="nodegroup_info_max_nodesize_' + newNodeGroupIndex + '" placeholder="" title="" />';
    html += '   </li>';
    html += '<li>';
    html += '<label>Min Node Size</label>';
    html += '       <input type="text" name="nodegroup_info_min_nodesize" value="" id="nodegroup_info_min_nodesize_' + newNodeGroupIndex + '" placeholder="" title="" />';
    html += '</li>';
    html += '<li>';
    html += '<label>On Auto Scaling</label>';
    html += '       <input type="text" name="nodegroup_info_onautoscaling" value="" id="nodegroup_info_onautoscaling_' + newNodeGroupIndex + '" placeholder="" title="" />';
    html += '</li>';
    html += '<li>';
    html += '<label>Root Disk</label>';    
    //html += '       <input type="text" name="nodegroup_info_rootdisk_type" value="" id="nodegroup_info_rootdisk_type_' + newNodeGroupIndex + '" placeholder="type" title="" />';
    html += '<select class="selectbox white pline sel_4" name="rootDiskType" id="nodegroup_info_rootdisk_type_' + newNodeGroupIndex + '" onchange="changeDiskSize(this.value);"></select>';    
    html += '       <input type="text" name="nodegroup_info_rootdisk_size" value="" id="nodegroup_info_rootdisk_size_' + newNodeGroupIndex + '" placeholder="size" title="" />';
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

// addButton 이 필요없어 주석처리.
//function nodegroupDone_btn(){
//}

// 추가 된 nodeGroup을 저장
function deployNodeGroup(){
    var clusterID = $("#pmks_cluster_id").val();
    var connectionName = $("#pmks_cluster_connection").val();

    var obj = {};// request Obj
    var nodeGroupObj = {};

    var addNodeGroupIndex = 0;

    // 추가한 nodeGroup들 
    $("[name='nodegroup_idx']").each(function (idx, ele) {
        // 중간에 이빨이 빠진게 있을 수 있으므로 id에서 index 추출하여 set
        var addNodeGroup = $(this).attr("id")
        // index 추출
        var tempNodeGroupIdxArr = addNodeGroup.split("_")
        addNodeGroupIndex = tempNodeGroupIdxArr[tempNodeGroupIdxArr.length -1]
	});

    // 1개만 추가하므로 루프 돌 필요 없음.
    var nodeGroupName = $("#nodegroup_info_name_" + addNodeGroupIndex).val();
    if( nodeGroupName == ""){
        commonAlert("Please Input Name");
        return
    }
    var nodeGroupImageId = $("#nodegroup_info_imageid_" + addNodeGroupIndex).val();
    var nodeGroupVmSpecId = $("#nodegroup_info_vmspecid_" + addNodeGroupIndex).val();
    if( nodeGroupVmSpecId == ""){
        commonAlert("Please Input Spec");
        return
    }
    var nodeGroupKeyPairId = $("#nodegroup_info_keypairid_" + addNodeGroupIndex).val();
    if( nodeGroupVmSpecId == ""){
        commonAlert("Please Input Key Pair");
        return
    }
    var nodeGroupDesiredNodeSize = $("#nodegroup_info_desired_nodesize_" + addNodeGroupIndex).val();
    if( nodeGroupVmSpecId == ""){
        commonAlert("Please Input Desired Node Size");
        return
    }
    var nodeGroupMaxNodeSize = $("#nodegroup_info_max_nodesize_" + addNodeGroupIndex).val();
    if( nodeGroupVmSpecId == ""){
        commonAlert("Please Input Max Node Size");
        return
    }
    var nodeGroupMinNodeSize = $("#nodegroup_info_min_nodesize_" + addNodeGroupIndex).val();
    if( nodeGroupVmSpecId == ""){
        commonAlert("Please Input Min Node Size");
        return
    }
    var nodeGroupOnAutoScaling = $("#nodegroup_info_onautoscaling_" + addNodeGroupIndex).val();
    
    var nodeGroupRootDiskType = $("#nodegroup_info_rootdisk_type_" + addNodeGroupIndex).val();
    if( nodeGroupVmSpecId == ""){
        commonAlert("Please Input Root Disk Type");
        return
    }
    var nodeGroupRootDiskSize = $("#nodegroup_info_rootdisk_size_" + addNodeGroupIndex).val();
    if( nodeGroupVmSpecId == ""){
        commonAlert("Please Input Root Disk Size");
        return
    }
    
    

    var url = "/operation/manages/pmks/" + clusterID + "/nodegroup"
    console.log("URL : ", url)

    nodeGroupObj = {
        Name: nodeGroupName,
        ImageName: nodeGroupImageId,
        VMSpecName: nodeGroupVmSpecId,
        KeyPairName: nodeGroupKeyPairId,
        DesiredNodeSize: nodeGroupDesiredNodeSize,
        MaxNodeSize: nodeGroupMaxNodeSize,
        MinNodeSize: nodeGroupMinNodeSize,
        OnAutoScaling: nodeGroupOnAutoScaling,
        RootDiskType: nodeGroupRootDiskType,
        RootDiskSize: nodeGroupRootDiskSize
    }
  
    obj['ConnectionName'] = connectionName
    obj['ReqInfo'] = nodeGroupObj;

    console.log("info image obj Data : ", obj);

    if (obj) {
        axios.post(url, obj, {
            headers: {
                //'Content-type': 'application/json',
                // 'Authorization': apiInfo,
            }
        }).then(result => {
            console.log("result : ", result);
            if (result.status == 200 || result.status == 201) {
                commonAlert("NodeGroup Creation Requested")
                // 목록 화면으로 전환
            } else {
                var message = result.data.message;
                commonAlert("Failed to Create NodeGroup : " + message + "(" + result.status + ")");
            }
        }).catch((error) => {
            console.warn(error);
            //commonErrorAlert(statusCode, errorMessage);
        });
    }
}


var DISK_TYPES_SIZES
// ConnectionName에 따른 DisType 목록
function getCommonLookupDiskInfoSuccess(caller, providerID, data){
    DISK_TYPES_SIZES = data;
    var root_disk_type = [];    
	var res_item = data
    console.log(res_item)
	res_item.forEach(item=>{
		//var temp_provider = item.provider
		//if(temp_provider == providerID){
			root_disk_type = item.rootdisktype
			DISK_TYPES_SIZES = item.disksize
		//}
	})

	var html = '<option value="">Select Root Disk Type</option>'
	console.log("root_disk_type : ",root_disk_type);
	root_disk_type.forEach(item=>{
		html += '<option value="'+item+'">'+item+'</option>'
	})

    
    var addNodeGroupIndex = 0;
    // 추가한 nodeGroup들 
    $("[name='nodegroup_idx']").each(function (idx, ele) {
        // 중간에 이빨이 빠진게 있을 수 있으므로 id에서 index 추출하여 set
        var addNodeGroup = $(this).attr("id")
        // index 추출
        var tempNodeGroupIdxArr = addNodeGroup.split("_")
        addNodeGroupIndex = tempNodeGroupIdxArr[tempNodeGroupIdxArr.length -1]

        var rootDiskType = $("#nodegroup_info_rootdisk_type_" + addNodeGroupIndex);
        rootDiskType.empty();
		rootDiskType.append(html);
        console.log(rootDiskType)
        console.log("html", html)
	});

    
}

function changeDiskSize(diskType){

}