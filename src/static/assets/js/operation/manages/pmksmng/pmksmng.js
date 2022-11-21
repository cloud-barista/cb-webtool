var TOTAL_PMKS_LIST = new Map();// 모든 PMKS 정보
$(document).ready(function () {
    setTableHeightForScroll("pmksListTable", 700);

    getCommonAllPmksList("onload");
});

// Cluster 목록
function displayClusterListArea() {
    $(".dashboard .status_list tbody tr").each(function () {
        var $td_list = $(this),
            $status = $(".server_status"),// 1단계
            $status_sub = $(".server_status_sub"),// 2 단계
            $detail = $(".server_info");
        $td_list.off("click").click(function () {
            $td_list.addClass("on");
            $td_list.siblings().removeClass("on");
            $status.addClass("view");
            $status.siblings().removeClass("on");
            $status_sub.siblings().removeClass("on");// 2 단계

            $(".dashboard.register_cont").removeClass("active");
            $td_list.off("click").click(function () {
                if ($(this).hasClass("on")) {
                    console.log("reg ok button click")
                    $td_list.removeClass("on");
                    $status.removeClass("view");
                    $status_sub.removeClass("view");
                    $detail.removeClass("active");
                } else {
                    $td_list.addClass("on");
                    $td_list.siblings().removeClass("on");
                    $status.addClass("view");

                    $status.siblings().removeClass("view");
                    $status_sub.siblings().removeClass("view");
                    $(".dashboard.register_cont").removeClass("active");
                }
            });
        });
    });
}

// NodeGroup 정보 표시, Node List 표시
function displayNodeGroupInfoArea() {
    $(".dashboard .ds_cont .area_cont .listbox li.sel_cr").each(function () {
        var $sel_list = $(this);// nodeGroup icon 목록
        //var $detail = $(".server_info");// node 목록
        $status_sub = $(".server_status_sub"),// 2 단계
            console.log(">>>>>");
        $sel_list.off("click").click(function () {
            $sel_list.addClass("active");
            $sel_list.siblings().removeClass("active");

            $status_sub.addClass("view");

            $sel_list.off("click").click(function () {
                if ($(this).hasClass("active")) {
                    $sel_list.removeClass("active");
                    $status_sub.removeClass("view");
                } else {
                    $sel_list.addClass("active");
                    $sel_list.siblings().removeClass("active");
                    $status_sub.addClass("view");
                }
            });
        });
    });
}

function displayNodeGroupListArea() {
    //<div class="dashboard dashboard_cont server_status_sub" id="pmks_nodegroup_detail">
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

// 해당 area가 나타나면서 set된 data표시
function displayNodeArea(clusterID, nodeGroupID, nodeID) {

    // var vmID = vmData.id;
    // var vmName = vmData.name;
    // var vmStatus = vmData.status;
    // var vmDispStatus = getVmStatusDisp(vmStatus);
    // var vmStatusIcon = getVmStatusIcon(vmDispStatus);

    $("#node_info_text").text('[' + clusterID + '/' + nodeGroupID + '/' + nodeID + ']')
    // $("#node_info_status_icon_img").attr("src", vmStatusIcon);


    // $("#server_detail_info_text").text('[' + vmName + '/' + mcisName + ']')
    // $("#server_detail_info_public_ip_text").text("Public IP : " + vmPublicIp)


    // $("#server_detail_view_server_status").val(vmStatus);// detail tab
    $(".dashboard .ds_cont .area_cont .listbox li.sel_cr").each(function () {
        var $sel_list = $(this);
        var $detail = $(".server_info");
        // console.log($sel_list);
        // console.log($detail);
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

// 모든 connection에 대한 PMKS 목록
function getCommonAllPmksListSuccess(caller, data) {
    if (caller == "onload") {

        TOTAL_PMKS_LIST = new Map();
        if (data.length) {
            for (var i in data) {
                clusterInfo = data[i]
                var clusterID = clusterInfo.IId.NameId;
                TOTAL_PMKS_LIST.set(clusterID, clusterInfo)
            }
        }
        setClusterList()
    }
}


// Cluster 단건조회 data = PmksInfo
function getPmksSuccess(clusterID, data) {
    //console.log(data);

}


// cluster Data를 매핑
function setClusterList() {
    var html = "";
    var idx = 0;
    for (const [clusterID, clusterInfo] of TOTAL_PMKS_LIST) {
        html += addClusterData(clusterInfo, idx);
        idx++;
    }
    $("#clusterList").empty();
    $("#clusterList").append(html);

    displayClusterListArea();// 화면에 표시
}
// Cluster 1개 Row 추가
function addClusterData(item, index) {
    // spider를 직접 호출하기 때문에 가져오는 data형태가 좀 다름.
    var html = "";

    // nodegroup은 IId(소문자), Vpc는 IID(대문자임...)
    var clusterID = item.IId.NameId;
    var clusterSystemID = item.IId.SystemId;
    var clusterVersion = item.Version;
    var clusterProviderName = item.ProviderName;
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
    console.log(item)
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

    // cluster 만 있을 때는 없을 수도 있음.
    if (item.NodeGroupList != null) {
        for (var i = 0; i < item.NodeGroupList.length; i++) {
            var nodeGroup = item.NodeGroupList[i];
            var nodeGroupName = nodeGroup.IId.NameId
            nodeGroupIds.push(nodeGroupName)
        }
    }

    html +=
        "<tr onclick=\"clickListOfCluster('" + clusterID + "', " + index + ");\">"
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


        + '<input type="checkbox" name="clusterchk" value="' + clusterID + '" id="cluster_ch_td_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>'

        + '<td class="btn_mtd ovm column-80px" data-th="status">' + clusterStatus + "</td>"
        + '<td class="overlay hidden" data-th="clusterName">' + clusterID + "</td>"
        + '<td class="btn_mtd ovm" data-th="provider">' + clusterProviderName + "</td>"

        + '<td class="overlay hidden" data-th="vpc">' + vpcId + "</td>"
        + '<td class="overlay hidden" data-th="subnet">' + subnetIds + "</td>"
        + '<td class="overlay hidden" data-th="securityGroup">' + securityGroupIds + "</td>"
        + '<td class="overlay hidden" data-th="nodeGroup">' + nodeGroupIds + "</td>"
        + "</tr>"
    return html
}

// Cluster의 NodeGroup 영역 표시
function setNodeGroupList(clusterID) {
    var clusterInfo = TOTAL_PMKS_LIST.get(clusterID);
    var nodeGroupList = clusterInfo.NodeGroupList;
    var html = "";

    // Cluster만 있는 경우 NodeGroup이 없을 수 있음
    if (nodeGroupList != null) {
        if (nodeGroupList.length) {
            var nodeGroupID = ""
            for (var i in nodeGroupList) {
                nodeGroupInfo = nodeGroupList[i];
                nodeGroupID = nodeGroupInfo.IId.NameId;
                html += addNodeGroupData(nodeGroupInfo, i, clusterID);

            }
            $("#pmks_nodegroup_list_info_box").empty();
            $("#pmks_nodegroup_list_info_box").append(html);

            if (nodeGroupList.length == 1) {
                setNodeList(clusterID, nodeGroupID);
            }

        } else {
            html += CommonTableRowNodata(8);
            $("#pmks_nodegroup_list_info_box").empty();
            $("#pmks_nodegroup_list_info_box").append(html);
        }
    }

    displayNodeGroupInfoArea();
}

// NodeGroup Table : nodegroupList
function addNodeGroupData(item, nodeGroupIndex, clusterID) {
    console.log("addNodeGroupData")
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
    var nodeGroupStatusIcon = getNodeGroupStatusIcon(status);
    var nodeGroupDispStatus = getNodeGroupStatusDisp(status);
    var nodeGroupDispClass = getNodeGroupStatusClass(status);

    // var vmDispStatus = getVmStatusDisp(vmStatus);
    // var vmStatusClass = getVmStatusClass(vmDispStatus)
    // vmLi += '<li id="server_status_icon_' + vmID + '" class="sel_cr ' + vmStatusClass + '"><a href="javascript:void(0);" onclick="vmDetailInfo(\'' + mcisID + '\',\'' + mcisName + '\',\'' + vmID + '\')"><span class="txt">' + vmName + '</span></a></li>';
    $("#nodegroup_info_name").val(nodeGroupID);
    $("#nodegroup_info_imageid").val(imageID);
    $("#nodegroup_info_spec").val(vmSpecName);
    $("#nodegroup_info_keypair").val(keyPairID);
    $("#nodegroup_info_desirednodesize").val(desiredNodeSize);
    $("#nodegroup_info_maxnodesize").val(maxNodeSize);
    $("#nodegroup_info_minnodesize").val(minNodeSize);
    $("#nodegroup_info_onautoscaling").val(onAutoScaling);
    $("#nodegroup_info_rootdisktype").val(rootDiskType);
    $("#nodegroup_info_rootdisksize").val(rootDiskSize);


    var nodeIds = new Array();
    for (var i = 0; i < item.Nodes.length; i++) {
        var node = item.Nodes[i];
        var nodeName = node.NameId
        nodeIds.push(nodeName)
    }
    console.log("li attrt class " + nodeGroupStatusIcon)

    var nodeGroupBadge = '<li class="sel_cr ' + nodeGroupDispClass + '" id="nodeGroupOfCluster_' + nodeGroupIndex + '"><a href="javascript:void(0);" onclick="clickListOfNodeGroup(\'' + clusterID + '\',\'' + nodeGroupID + '\')"><span class="txt">' + nodeGroupID + '</span></a></li>';

    return nodeGroupBadge
}

// cluster 선택 -> ( nodegroup 목록 표시) -> nodeGroup 선택 했을 때 해당 Node 목록 표시
function setNodeList(clusterID, nodeGroupID) {
    var clusterInfo = TOTAL_PMKS_LIST.get(clusterID);
    var nodeGroupList = clusterInfo.NodeGroupList;
    var nodeGroupInfo = "";

    // nodeGroup 정보 표시
    console.log("nodeGroupList", nodeGroupList)

    //$("#pmks_nodegroup_detail_info_box")

    // node 목록
    for (var i in nodeGroupList) {
        tempNodeGroupInfo = nodeGroupList[i];
        tempNodeGroupID = tempNodeGroupInfo.IId.NameId;

        if (nodeGroupID == tempNodeGroupID) {
            nodeGroupInfo = tempNodeGroupInfo;
            break
        }
    }

    $("#pmks_nodegroup_id").val(nodeGroupInfo.IId.NameId);// 선택한 nodeGroup Id. PmksMng.html에 정의.

    $("#nodegroup_info_name").val(nodeGroupInfo.IId.NameId);
    $("#nodegroup_info_imageid").val(nodeGroupInfo.ImageIID.NameId);
    $("#nodegroup_info_spec").val(nodeGroupInfo.VMSpecName);
    $("#nodegroup_info_keypair").val(nodeGroupInfo.KeyPairIID.NameId);
    // nodeGroupInfo.Status
    $("#nodegroup_info_desired_node_size").val(nodeGroupInfo.DesiredNodeSize);
    $("#nodegroup_info_max_node_size").val(nodeGroupInfo.MaxNodeSize);
    $("#nodegroup_info_min_node_size").val(nodeGroupInfo.MinNodeSize);

    $("#nodegroup_info_on_scaling_auto").val(nodeGroupInfo.OnAutoScaling);


    $("#nodegroup_info_root_disk_type").val(nodeGroupInfo.RootDiskType);
    $("#nodegroup_info_root_disk_size").val(nodeGroupInfo.RootDiskSize);







    var html = "";
    if (nodeGroupInfo.Nodes != null && nodeGroupInfo.Nodes.length > 0) {
        var nodeList = nodeGroupInfo.Nodes;

        console.log("nodeList size ", nodeGroupInfo.Nodes.length)
        console.log("nodeList", nodeList)
        var html = "";
        for (var i in nodeList) {
            html += addNodeRow(nodeList[i], i, clusterID, nodeGroupID);
        }

    } else {
        html = "<li><span>Node doesn't exists, Please check on auto scaling option</span></a></li>"
    }
    $("#pmks_node_list_info_box").empty();
    $("#pmks_node_list_info_box").append(html);

    // TODO : 각 Node의 상태정보 조회
    displayNodeGroupInfoArea();
}

// 선택한 Node의 상세정보 표시
function setNode(clusterID, nodeGroupID, nodeID) {

    // $("#pmks_node_txt").text(nodeName + " / " + nodeUID);

    // $("#pmks_node_name").val(nodeName);
    // $("#pmks_node_kind").val(nodeKind);
    // $("#pmks_node_role").val(nodeRole);

    // $("#pmks_node_detail").css("display", "block");
    displayNodeArea(clusterID, nodeGroupID, nodeID)
}

// NodeGroup Table : nodegroupList
function addNodeRow(item, nodeIndex, clusterID, nodeGroupID) {
    console.log("addNodeRow")
    console.log(item)

    var html = "";

    var nodeID = item.NameId;
    var nodeSystemID = item.SystemId;
    var nodeName = nodeID;
    if (nodeName == "") {
        nodeName = nodeSystemID;
    }

    // 현재 Node는 status 정보가 없음
    // var status = item.Status;
    // var nodeStatusIcon = getNodeStatusIcon(status);
    // var nodeDispStatus = getNodeStatusDisp(status);
    // var nodeDispClass = getNodeStatusClass(status);

    var nodeBadge = '<li class="sel_cr bgbox_r" id="node_' + nodeIndex + '"><a href="javascript:void(0);" onclick="clickListOfNode(\'' + clusterID + '\',\'' + nodeGroupID + '\',\'' + nodeID + '\')"><span class="txt">' + nodeName + '</span></a></li>';

    return nodeBadge
}

// Cluster 목록에서 Cluster 클릭 -> Cluster Info + NodeGroup List icon 
function clickListOfCluster(clusterID, clusterIndex) {
    var clusterInfo = TOTAL_PMKS_LIST.get(clusterID);
    // cluster Info 표시    
    $("#pmks_cluster_id").val(clusterID);// hidden        
    $("#pmks_cluster_name").val(clusterID);// hidden        
    $("#pmks_cluster_connection").val(clusterInfo.ConnectionName);// hidden

    $("#pmks_info_txt").text("[ " + clusterID + " ]");// title    

    // Cluster Info
    $("#pmks_info_name").val(clusterID);
    $("#pmks_info_version").val(clusterInfo.Version);
    $("#pmks_info_cloud_connection").val(clusterInfo.ConnectionName);
    $("#pmks_info_status").val(clusterInfo.Status);

    // NetworkInfo
    if (clusterInfo.Network != null) {
        var networkInfo = clusterInfo.Network;

        $("#pmks_info_vpc").val(networkInfo.VpcIID.NameId);
        $("#pmks_info_subnet").val(networkInfo.SubnetIIDs);
        $("#pmks_info_security_group").val(networkInfo.SecurityGroupIIDs);
    }

    // AccessInfo
    if (clusterInfo.AccessInfo != null) {
        var accessInfo = clusterInfo.AccessInfo;
        $("#pmks_info_endpoint").val(accessInfo.Endpoint);
        $("#pmks_info_kubeconfig").val(accessInfo.Kubeconfig);
    }

    // NodeGroup 이 있으면NodeGroup 목록 표시
    if (clusterInfo.NodeGroupList != null) {
        var html = "";
        for (var o in clusterInfo.NodeGroupList) {
            var nodeGroupStatus = clusterInfo.NodeGroupList[o].Status;
            var nodeGroupName = clusterInfo.NodeGroupList[o].IId.NameId

            var nodeGroupDispStatus = "";//getNodeGroupStatusDisp(nodeGroupStatus);
            var nodeGroupStatusClass = "bgbox_b";//getNodeGroupStatusClass(nodeGroupName)
            html += '<li id="nodegroup_status_icon_' + o + '" class="sel_cr ' + nodeGroupStatusClass + '"><a href="javascript:void(0);" onclick="clickListOfNodeGroup(\'' + clusterID + '\',\'' + nodeGroupName + '\')"><span class="txt">' + nodeGroupName + '</span></a></li>';
        }
        $("#cluster_nodegroup_list").empty();
        $("#cluster_nodegroup_list").append(html);
    }
}

// NodeGroup 목록에서 NodeGroup 클릭 -> Node 목록
function clickListOfNodeGroup(clusterID, nodeGrouID) {
    console.log("clickListOfNodeGroup")
    setNodeList(clusterID, nodeGrouID);
}

// Node 목록에서 Node 클릭
function clickListOfNode(clusterID, nodeGrouID, nodeID) {
    setNode(clusterID, nodeGrouID, nodeID);
}


// PMKS Info area 안의 NodeList 내용 표시
// 해당 PMKS를 조회하여 NodeGroup 상세정보 표시
function setNodeGroupList(clusterID, clusterIndex) {
    // Name, version
    $("#pmks_uid").val(clusterID);// hidden    

    // PMKS Info
    $("#pmks_info_txt").text("[ " + clusterID + " ]");

    // cluster 정보
    var clusterInfo = TOTAL_PMKS_LIST.get(clusterID);
    var connectionName = $("#cluster_connection_" + clusterIndex).val()

    $("#pmks_uid").val(clusterID);
    $("#pmks_info_version").val($("#cluster_version_" + clusterIndex).val());
    $("#pmks_info_name").val($("#cluster_info_" + clusterIndex).val() + " / " + $("#cluster_systemid_" + clusterIndex).val());
    $("#pmks_info_status").val($("#cluster_status_" + clusterIndex).val());
    $("#pmks_info_cloud_connection").val(connectionName);

    // Network
    $("#pmks_vpc").val($("#network_vpc_nameid_" + clusterIndex).val() + " / " + $("#network_vpc_systemid_" + clusterIndex).val());
    $("#pmks_subnet").val($("#network_subnet_nameid_" + clusterIndex).val());
    $("#pmks_info_security_group").val($("#network_securitygroup_nameid_" + clusterIndex).val());

    // NodeGroupList :     
    //getPmks(uid, connectionName)
    addNodeGroupData(clusterID)
}

// 해당 pmks에 nodeGroup 추가
// pmks가 경로에 들어가야 함. node 등록 form으로 이동
function addNewNodeGroup() {
    var clusterID = $("#pmks_cluster_id").val();
    var clusterName = $("#pmks_cluster_name").val();
    var connectionName = $("#pmks_cluster_connection").val();

    if (clusterID == "") {
        commonAlert("PMKS 정보가 올바르지 않습니다.");
        return;
    }

    var urlParamMap = new Map();
    urlParamMap.set(":clusterID", clusterID)
    //changePage('PmksNodeGroupRegForm', urlParamMap)
    url = setUrlByParam("PmksNodeGroupRegForm", urlParamMap)

    //onClick="changePage('PmksNodeGroupRegForm')"
    //pmksNodeGroupRegGroup := e.Group("/operation/manages/pmksmng/cluster/:clusterID/regform", pmksNodeGroupRegTemplate)
    //var url = "/operation/manages/pmksmng/regform/" + clusterId + "/" + clusterName;

    location.href = url + "?connectionName=" + connectionName;
}

// PMKS 삭제
function deleteCluster() {
    //function deletePMKS() {
    var checkedCount = 0;
    var clusterID = "";
    var clusterName = "";
    var connectionName = "";
    $("[id^='cluster_ch_td_']").each(function () {
        if ($(this).is(":checked")) {
            checkedCount++;
            console.log("checked")
            clusterID = $(this).val();
            var objId = $(this).attr("id")
            var lastIndexArr = objId.split("_")
            var lastIndex = lastIndexArr[lastIndexArr.length - 1];
            console.log(lastIndexArr)
            console.log("lastIndex " + lastIndex);
            connectionName = $("#cluster_connection_" + lastIndex).val();
            console.log("connectionName = " + connectionName)
        } else {
            console.log("checked nothing")
        }
    })

    if (checkedCount == 0) {
        commonAlert("Please Select a PMKS")
        return;
    } else if (checkedCount > 1) {
        commonAlert("Please Select One PMKS at a time")
        return;
    }

    var url = "/operation/manages/pmks/" + clusterID + "?connectionName=" + connectionName;
    axios.delete(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get  Data : ", result.data);
        var statusCode = result.data.status;
        var message = result.data.message;

        if (statusCode != 200 && statusCode != 201) {
            commonAlert(message + "(" + statusCode + ")");
            return;
        } else {
            commonAlert("PMKS Deletion Requested");
            getCommonAllPmksList("onload");//PMKS List 조회
        }

    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage)
    });

}

// nodeGroup 삭제
function deleteNodeGroupOfPmks() {
    // nodegroup 1개씩 삭제

    var selectedPmksId = $("#pmks_cluster_id").val();
    //var selectedPmksName = $("#pmks_cluster_name").val();
    var selectedNodeId = $("#pmks_nodegroup_id").val();
    //var selectedNodeName = $("#pmks_nodegroup_name").val();

    var urlParamMap = new Map();
    urlParamMap.set(":clusterID", selectedPmksId)
    //urlParamMap.set(":clusterName", selectedPmksName)
    urlParamMap.set(":nodeGroupID", selectedNodeId)
    //urlParamMap.set(":nodeGroupName", selectedNodeId)
    var url = setUrlByParam("PmksNodeGroupDelProc", urlParamMap)
    console.log("URL : ", url)
    axios.delete(url, {
        // headers: {
        //     'Content-Type': "application/json"
        // }
    }).then(result => {
        // var data = result.data;
        // if (result.status == 200 || result.status == 201) {
        var statusCode = result.data.status;
        if (statusCode == 200 || statusCode == 201) {
            commonAlert("NodeGroup Deleted");

        } else {
            var message = result.data.message;
            commonAlert("Fail to Delete NodeGroup : " + message + "(" + statusCode + ")");

        }
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

// Scale Size 변경
function UpdateNodeGroupScaleSize(caller) {
    var clusterID = $("#pmks_cluster_id").val();
    var clusterName = $("#pmks_cluster_name").val();
    var connectionName = $("#pmks_cluster_connection").val();
    var nodeGroupId = $("#pmks_nodegroup_id").val();
    var desiredSize = $("#nodegroup_info_desired_node_size").val();
    var maxSize = $("#nodegroup_info_max_node_size").val();
    var minSize = $("#nodegroup_info_min_node_size").val();
    var sizeValue = "";
    if (caller == "DesiredNodeSize") {
        if (desiredSize > maxSize) {
            commonAlert("too many");
            return
        }
        if (desiredSize < minSize) {
            commonAlert("too little");
            return
        }
        sizeValue = $("#nodegroup_info_desired_node_size").val();
    } else if (caller == "MaxNodeSize") {
        if (minSize > maxSize) {
            commonAlert("too many");
            return
        }

        sizeValue = $("#nodegroup_info_max_node_size").val();
    } else if (caller == "MinNodeSize") {
        if (minSize > maxSize) {
            commonAlert("too many");
            return
        }
        sizeValue = $("#nodegroup_info_min_node_size").val();
    }

    var urlParamMap = new Map();
    urlParamMap.set(":clusterID", clusterID)
    urlParamMap.set(":nodeGroupID", selectedNodeUid)

    var new_obj = {}
    new_obj['ConnectionName'] = connectionName


    var url = setUrlByParam("PmksNodeGroupDelProc", urlParamMap)
    try {
        axios.put(url, new_obj, {
            // headers: {
            //     'Content-type': "application/json",
            // },
        }).then(result => {
            console.log("update data : ", result);
            console.log("Result Status : ", result.status);
            if (result.status == 201 || result.status == 200) {
                commonResultAlert("Updated")
            } else {
                commonAlert("Update Failed")
            }
        }).catch((error) => {
            console.warn(error);
            commonAlert(error);
            // console.log(error.response)
            // var errorMessage = error.response.data.error;
            // var statusCode = error.response.status;
            // commonErrorAlert(statusCode, errorMessage)

        })
    } catch (error) {
        commonAlert(error);
        console.log(error);
    }
}

function UpdateNodeGroupAutoScalingOnOff() {
    //
    var scalingOnOff = $("#nodegroup_info_on_scaling_auto").val()
}




