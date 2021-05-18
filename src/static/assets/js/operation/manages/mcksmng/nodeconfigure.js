// connection 정보가 바뀌었을 때, 변경 될 object : 원래는 각각 만들어야 하나, 가져오는게 spec만 있어서 plane, worker 같이 씀.
function changeConnectionInfo(configName, targetObjId){
    console.log("config name : ",configName)
    if( configName == ""){
        // 0번째면 selectbox들을 초기화한다.(vmInfo, sshKey, image 등)
    }
    
    getSpecInfo(configName, targetObjId);
}

// connection에 맞는 spec들 조회
function getSpecInfo(configName, targetObjId){
    var configName = configName;
    if(!configName){
        configName = $("#nodeConnectionName option:selected").val();
    }

    var url = "/setting/resources/vmspec/list"
    var html = "";
    axios.get(url,{
        // headers:{
        // 	'Authorization': apiInfo
        // }
    }).then(result=>{
        // console.log(result.data)
        var data = result.data.VmSpecList
        console.log("spec result : ",data)
        if(data){
            html +="<option value=''>Select SpecName</option>"
            data.filter(csp => csp.connectionName === configName).map(item =>(
                html += '<option value="'+item.id+'">'+item.name+'('+item.cspSpecName+')</option>'	
            ))

        }else{
            html +=""
        }       
      
        $("#" + targetObjId).empty();
        $("#" + targetObjId).append(html);        
    })
}

// mcks , node deploy
// 우선 mcks 부터
function deploy_btn(){
    var mcksName = $("#mcks_name").val();

    var kubernatesNetworkCni = $("#kubernatesNetworkCni").val();
    var kubernatesPodCidr = $("#kubernatesPodCidr").val();
    var kubernatesServiceCidr = $("#kubernatesServiceCidr").val();
    var kubernatesServiceDnsDomain = $("#kubernatesServiceDnsDomain").val();
    
    var controlPlaneLength = $("input[name='controlPlaneCount']").length;
    console.log("controlPlaneLength1 " + controlPlaneLength)
    var controlPlaneConnectionData = new Array(controlPlaneLength);
    var controlPlaneCountData = new Array(controlPlaneLength);
    var controlPlaneSpecIdData = new Array(controlPlaneLength);
    for(var i=0; i<controlPlaneLength; i++){                          
        controlPlaneConnectionData[i] = $("select[name='controlPlaneConnectionName']")[i].value;
        controlPlaneCountData[i] = $("input[name='controlPlaneCount']")[i].value;
        controlPlaneSpecIdData[i] = $("select[name='controlPlaneSpecId']")[i].value;
    }
    console.log(controlPlaneConnectionData)
    console.log(controlPlaneCountData)
    console.log(controlPlaneSpecIdData)
    
    var workerCountLength = $("input[name='workerCount']").length;
    console.log("workerCountLength1 " + workerCountLength)
    var workerConnectionData = new Array(workerCountLength);
    var workerCountData = new Array(workerCountLength);
    var workerSpecIdData = new Array(workerCountLength);
    for(var i=0; i<workerCountLength; i++){                          
        workerConnectionData[i] = $("select[name='workerConnectionName']")[i].value;
        workerCountData[i] = $("input[name='controlPlaneCount']")[i].value;
        workerSpecIdData[i] = $("select[name='workerSpecId']")[i].value;
    }
    console.log(workerConnectionData)
    console.log(workerCountData)
    console.log(workerSpecIdData)
    var new_obj = {}
    // mcis 생성이므로 mcisID가 없음
    new_obj['name'] = mcksName
    
    var new_kubernetes = {}
    new_kubernetes['networkCni'] = kubernatesNetworkCni;
    new_kubernetes['podCidr'] = kubernatesPodCidr;
    new_kubernetes['serviceCidr'] = kubernatesServiceCidr;
    new_kubernetes['serviceDnsDomain'] = kubernatesServiceDnsDomain;

    new_obj['config'] = new_kubernetes
    var controlPlanes = new Array(controlPlaneLength);
    console.log("controlPlaneConnectionLength " + controlPlaneLength)
    for(var i=0; i<controlPlaneLength; i++){
        console.log("controlPlane " + i)
        var new_controlPlane = {}
        new_controlPlane['connection'] = controlPlaneConnectionData[i];
        new_controlPlane['count'] = controlPlaneCountData[i]
        new_controlPlane['spec'] = controlPlaneSpecIdData[i]
        controlPlanes[i] = new_controlPlane
    }
    new_obj['controlPlane'] = controlPlanes;

    var workers = new Array(workerCountLength);
    for(var i=0; i<workerCountLength; i++){
        console.log("workerCountLength " + i)
        var new_worker = {}
        new_worker['connection'] = workerConnectionData[i];
        new_worker['count'] = workerCountData[i]
        new_worker['spec'] = workerSpecIdData[i]
        workers[i] = new_worker
    }
    new_obj['worker'] = workers;
//     $("input[name='workerCount']").each(function (i) {
//         var new_worker = {}
//         console.log($("select[name='workerConnectionName']").eq(i));
//         new_worker['connection'] = $("select[name='workerConnectionName']").eq(i).attr("value");
//         new_worker['count'] = $("input[name='workerCount']").eq(i).attr("value")
//         new_worker['spec'] = $("select[name='workerSpecId']").eq(i).attr("value")
//         workers[i] = new_worker
//         console.log( i + "번째  : " );
//         console.log( new_worker);
//    });
   
    console.log(new_obj);

    try{
        var url = "/operation/manages/mcksmng/" + mcis_id +"/vm/reg/proc"
        axios.post(url,new_obj,{
            headers :{
                },
        }).then(result=>{
            console.log("MCIR VM Register data : ",result);
            console.log("Result Status : ",result.status); 

            var statusCode = result.data.status;
            var message = result.data.message;
            console.log("Result Status : ",statusCode); 
            console.log("Result message : ",message); 

            if(result.status == 201 || result.status == 200){
                vmCreateCallback(new_obj.name, "Success")
            //     commonAlert("Register Success")
            //     // location.href = "/Manage/MCIS/list";
            //     // $('#loadingContainer').show();
            //     // location.href = "/operation/manages/mcis/mngform/"
            //     var targetUrl = "/operation/manages/mcis/mngform"
            //     changePage(targetUrl)
            }else{
                vmCreateCallback(new_obj.name, message)    
            //     commonAlert("Register Fail")
            //     //location.reload(true);
            }
        }).catch((error) => {
            // console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            // commonErrorAlert(statusCode, errorMessage) 
            vmCreateCallback(new_obj.name, errorMessage)
        })
    }finally{
        
    }
}