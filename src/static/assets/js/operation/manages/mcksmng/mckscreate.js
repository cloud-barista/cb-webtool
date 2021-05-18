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
    
    var new_mcksConfig = {}
    var new_kubernetes = {}
    new_kubernetes['networkCni'] = kubernatesNetworkCni;
    new_kubernetes['podCidr'] = kubernatesPodCidr;
    new_kubernetes['serviceCidr'] = kubernatesServiceCidr;
    new_kubernetes['serviceDnsDomain'] = kubernatesServiceDnsDomain;

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
        controlPlanes[i] = new_controlPlane
    }
    new_obj['controlPlane'] = controlPlanes;

    var workers = new Array(workerCountLength);
    for(var i=0; i<workerCountLength; i++){
        console.log("workerCountLength " + i)
        var new_worker = {}
        new_worker['connection'] = workerConnectionData[i];
        new_worker['count'] = Number(workerCountData[i]);
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
                commonAlert("Clust Create Success")
            
            }else{
                commonErrorAlert(statusCode, message) 
            }
        }).catch((error) => {
            console.log(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            commonErrorAlert(statusCode, errorMessage) 
        })
    }finally{
        
    }
}
