$(document).ready(function() {


})

// Connection이 선택 되면
// 하위에 해당하는 정보들을 filter해서 가져온다
function changeConnectionInfo(connectionName){
    console.log("connectionName name : ",connectionName)
    if( connectionName == ""){
        // 0번째면 selectbox들을 초기화한다.(vmInfo, sshKey, image 등)
        return
    }
    
    var caller = "pmkscreate";
    var sortType = "name";
    var optionParam = "";
    var filterKey = "connectionName"
    var filterVal = connectionName


    // vpc : filter    
    getVnetInfoListForSelectbox(connectionName, 'regVNetId', 'regSubnetId')

    // subnet

    // security group
    getSecurityGroupListForSelectbox(connectionName, 'regSecurityGroupId')


    // sshkey
}

function validateCluster(){
    var clusterName = $("#cluster_info_name").val();

    var regProvider = $("#regProvider").val();
    var regConnectionName = $("#regConnectionName").val();
    var regVNetId = $("#regVNetId").val();
    var regSubnetId = $("#regSubnetId").val();
    var regSecurityGroupId = $("#regSecurityGroupId").val();
    
    if (!clusterName) {        
        commonAlert("Please Input Cluster Name");
        $("#cluster_info_name").focus()
        return false;
    }

    if (!regConnectionName) {
        commonAlert("Please Select Connection");
        return false;
    }
    if (!regVNetId) {
        commonAlert("Please Select VPC");
        return false;
    }
    if (!regSubnetId) {
        commonAlert("Please Select Subnet");
        return false;
    }
    if (!regSecurityGroupId) {
        commonAlert("Please Select Security Group");
        return false;
    }

    var selectedSubnets = [];
    var hasEmptyValue = false;
    $.each($("#regSubnetId option:selected"), function(){   
        if( $(this).val() == ""){            
            hasEmptyValue = true;
            return false;
        }
        selectedSubnets.push($(this).val());
    });
    if( hasEmptyValue){
        commonAlert("Please Select Subnet")
        return false;
    }
    if( selectedSubnets.length == 0){
        commonAlert("Please Select Subnet")
        return false;
    }
    
    var selectedSecurityGroups = [];
    $.each($("#regSecurityGroupId option:selected"), function(){            
        if( $(this).val() == ""){
            hasEmptyValue = true;
            return false;
        }
        selectedSecurityGroups.push($(this).val());
    });
    if( hasEmptyValue){
        commonAlert("Please Select Security Group")
        return false;
    }
    if( selectedSubnets.length == 0){
        commonAlert("Please Select Security Group")
        return false;
    }
    
    
    return true;
}

function createCluster(){
    if ( validateCluster() ){
        var clusterName = $("#cluster_info_name").val();
        var clusterVersion = $("#cluster_info_version").val();

        var regProvider = $("#regProvider").val();
        var regConnectionName = $("#regConnectionName").val();
        var regVNetId = $("#regVNetId").val();
        var regSubnetId = $("#regSubnetId").val();
        var regSecurityGroupId = $("#regSecurityGroupId").val();

        var selectedSubnets = [];
        $.each($("#regSubnetId option:selected"), function(){   
            if( $(this).val() == ""){
                commonAlert("Please Select Subnet")
                return false;
            }
            selectedSubnets.push($(this).val());
        });
        var selectedSecurityGroups = [];
        $.each($("#regSecurityGroupId option:selected"), function(){            
            if( $(this).val() == ""){
                commonAlert("Please Select Security Group")
                return false;
            }
            selectedSecurityGroups.push($(this).val());
        });

        var new_obj = {}
        var clusterReqInfo = {}
        clusterReqInfo["Name"] = clusterName;
        clusterReqInfo["Version"] = clusterVersion;
        clusterReqInfo["VPCName"] = regVNetId;
        clusterReqInfo["SubnetNames"] = selectedSubnets;
        clusterReqInfo["SecurityGroupNames"] = selectedSecurityGroups;

        new_obj['ConnectionName'] = regConnectionName
        new_obj['ReqInfo'] = clusterReqInfo
            
        console.log(new_obj)

        var url = getWebToolUrl("PmksClusterRegProc")
        console.log(url)
        try {
            axios.post(url, new_obj, {
                // headers: {
                //     'Content-type': "application/json",
                // },
            }).then(result => {
                console.log("PMKS Register data : ", result);
                console.log("Result Status : ", result.status);
                if (result.status == 201 || result.status == 200) {
                    commonResultAlert("Register Requested")

                    // 생성요청이 완료되었습니다. 관리화면으로 이동합니다.
                    //changePage("PmksMngForm")
                } else {
                    commonAlert("Register Fail")
                }
            }).catch((error) => {
                console.warn(error);
                //console.log(error.response)
                //var errorMessage = error.response.data.error;
                //var statusCode = error.response.status;
                //commonErrorAlert(statusCode, errorMessage)
    
            })
        } catch (error) {
            commonAlert(error);
            console.log(error);
        }
    }
}
