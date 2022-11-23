$(document).ready(function() {

    // 생성 완료 시 List화면으로 page이동
    $('#alertResultArea').on('hidden.bs.modal', function () {// bootstrap 3 또는 4
        changePage("McksMngForm");
    })

})

function btn_deploy(){
// function deploy_btn(){
    var mcksName = $("#mcksreg_name").val();
    if( !validateCloudbaristaKeyName(mcksName, 11) ){
        commonAlert("first letter = small letter <br/> middle letter = small letter, number, hyphen(-) only <br/> last letter = small letter <br/> max length = 11 ");
        return;
    }
    var mcksDescription = $("#mcksreg_desc").val();
    
    // ClusterConfigReq
    var installMonAgent = $("#clusterconfig_installMonitoringAgent").val();

    // ClusterConfigReqKubernetesReq
    var kubernatesVersion = $("#clusterconfig_kubernates_version").val();
    var kubernatesEtcd = $("#clusterconfig_kubernates_etcd").val();
    var kubernatesLoadBalancer = $("#clusterconfig_kubernates_loadbalancer").val();
    var kubernatesNetworkCni = $("#clusterconfig_kubernates_networkCni").val();
    var kubernatesPodCidr = $("#clusterconfig_kubernates_podCidr").val();
    var kubernatesServiceCidr = $("#clusterconfig_kubernates_serviceCidr").val();
    var kubernatesServiceDnsDomain = $("#clusterconfig_kubernates_serviceDnsDomain").val();
    var kubernatesStorageClassPath = $("#clusterconfig_kubernates_storageclass_nfs_path").val();
    var kubernatesStorageClassServer = $("#clusterconfig_kubernates_storageclass_nfs_server").val();

	// ControlPlane
    var controlPlaneLength = $("input[name='controlPlaneCount']").length;
    console.log("controlPlaneLength1 " + controlPlaneLength)
    var controlPlaneConnectionData = new Array(controlPlaneLength);
    var controlPlaneCountData = new Array(controlPlaneLength);
    var controlPlaneSpecIdData = new Array(controlPlaneLength);
    var controlPlaneRootDiskTypeData = new Array(controlPlaneLength);
    var controlPlaneRootDiskSizeData = new Array(controlPlaneLength);
    for(var i=0; i<controlPlaneLength; i++){                          
        controlPlaneConnectionData[i] = $("select[name='controlPlaneConnectionName']")[i].value;
        controlPlaneCountData[i] = $("input[name='controlPlaneCount']")[i].value;
        controlPlaneSpecIdData[i] = $("select[name='controlPlaneSpecId']")[i].value;
        controlPlaneRootDiskTypeData[i] = $("select[name='controlPlaneRootDiskType']")[i].value;
        controlPlaneRootDiskSizeData[i] = $("input[name='controlPlaneRootDiskSize']")[i].value;
    }
    console.log(controlPlaneConnectionData)
    console.log(controlPlaneCountData)
    console.log(controlPlaneSpecIdData)
    console.log(controlPlaneRootDiskTypeData)
    console.log(controlPlaneRootDiskSizeData)
    
    // Worker
    var workerCountLength = $("input[name='workerCount']").length;
    console.log("workerCountLength1 " + workerCountLength)
    var workerConnectionData = new Array();
    var workerCountData = new Array();
    var workerSpecIdData = new Array();
    var workerRootDiskTypeData = new Array();
    var workerRootDiskSizeData = new Array();
    for(var i=0; i<workerCountLength; i++){      
        var workerId = $("input[name='workerCount']").eq(i).attr("id");
        console.log("workerId " + workerId)
        if( workerId.indexOf("hidden_worker") > -1) continue;// 복사를 위한 영역이 있으므로

        workerConnectionData.push($("select[name='workerConnectionName']")[i].value);
        workerCountData.push($("input[name='workerCount']")[i].value);
        workerSpecIdData.push($("select[name='workerSpecId']")[i].value);
        workerRootDiskTypeData.push($("select[name='workerRootDiskType']")[i].value);
        workerRootDiskSizeData.push($("input[name='workerRootDiskSize']")[i].value);
    }
    console.log(workerConnectionData)
    console.log(workerCountData)
    console.log(workerSpecIdData)
    console.log(workerRootDiskTypeData)
    console.log(workerRootDiskSizeData)
    
    
    var new_obj = {}
    // mcks 생성이므로 mcksID가 없음
    new_obj['name'] = mcksName
    new_obj['description'] = mcksDescription
    
    var new_mcksConfig = {}
    var new_kubernetes = {}
    
    var new_storageclass = {}
    var new_nfs = {}
    new_nfs['path'] = kubernatesStorageClassPath;
    new_nfs['server'] = kubernatesStorageClassServer;
    new_storageclass['nfs'] = new_nfs;
    
    new_kubernetes['version'] = kubernatesVersion;
    new_kubernetes['etcd'] = kubernatesEtcd;
    new_kubernetes['loadbalancer'] = kubernatesLoadBalancer;
       

    new_kubernetes['networkCni'] = kubernatesNetworkCni;
    new_kubernetes['podCidr'] = kubernatesPodCidr;
    new_kubernetes['serviceCidr'] = kubernatesServiceCidr;
    new_kubernetes['serviceDnsDomain'] = kubernatesServiceDnsDomain;
    new_kubernetes['storageclass'] = new_storageclass;

    
    new_mcksConfig['installMonAgent'] = installMonAgent;
    new_mcksConfig['kubernetes'] = new_kubernetes;

    new_obj['config'] = new_mcksConfig;
    var controlPlanes = new Array(controlPlaneLength);
    console.log("controlPlaneConnectionLength " + controlPlaneLength)
    for(var i=0; i<controlPlaneLength; i++){
        console.log("controlPlane " + i)
        var new_controlPlane = {}
        new_controlPlane['connection'] = controlPlaneConnectionData[i];
        new_controlPlane['count'] = Number(controlPlaneCountData[i]);
        new_controlPlane['spec'] = controlPlaneSpecIdData[i]

        var rootDisk = {}
        rootDisk['type'] = controlPlaneRootDiskTypeData[i];
        rootDisk['size'] = controlPlaneRootDiskSizeData[i];

        new_controlPlane['rootDisk'] = rootDisk
        controlPlanes[i] = new_controlPlane
    }
    new_obj['controlPlane'] = controlPlanes;

    var workers = new Array(workerCountData.length);
    for(var i=0; i<workerCountData.length; i++){
        console.log("workerCountLength " + i)
        var new_worker = {}
        new_worker['connection'] = workerConnectionData[i];
        new_worker['count'] = Number(workerCountData[i]);
        new_worker['spec'] = workerSpecIdData[i]

        var rootDisk = {}
        rootDisk['type'] = workerRootDiskTypeData[i];
        rootDisk['size'] = workerRootDiskSizeData[i];
        new_worker['rootDisk'] = rootDisk

        workers[i] = new_worker
    }
    new_obj['worker'] = workers;
  
    console.log(new_obj);
    
    try{
        // configurer 는 mcks 선택하고 들어옴. : TODO : MCKS create 와 node create는 버튼 액션을 달리해야
        var url = "/operation/manages/mcksmng/reg/proc"
        axios.post(url,new_obj,{
            headers :{
                },
        }).then(result=>{
            console.log("data : ",result);
            console.log("Result Status : ",result.status); 

            var statusCode = result.data.status;
            var message = result.data.message;
            console.log("Result Status : ",statusCode); 
            console.log("Result message : ",message);

            if(result.status == 201 || result.status == 200){
                commonResultAlert("MCKS create request success")            
            }else{
                commonErrorAlert(statusCode, message) 
            }
        }).catch((error) => {
            console.log(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage) 
        })
    }finally{
        
    }
}

