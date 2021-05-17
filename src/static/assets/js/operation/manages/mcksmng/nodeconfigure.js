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
    
    var controlPlaneConnectionLength = $("input[name='controlPlaneConnectionName']").length;
    console.log("controlPlaneConnectionLength1 " + controlPlaneConnectionLength)
    var controlPlaneConnectionData = new Array(controlPlaneConnectionLength);
    var controlPlaneCountData = new Array(controlPlaneConnectionLength);
    var controlPlaneSpecIdData = new Array(controlPlaneConnectionLength);
    for(var i=0; i<controlPlaneConnectionLength; i++){                          
        controlPlaneConnectionData[i] = $("input[name='controlPlaneConnectionName']")[i].value;
        controlPlaneCountData[i] = $("input[name='controlPlaneCount']")[i].value;
        controlPlaneSpecIdData[i] = $("input[name='controlPlaneSpecId']")[i].value;
    }
    
    var workerConnectionNameLength = $("input[name='workerConnectionName']").length;
    console.log("workerConnectionNameLength1 " + workerConnectionNameLength)
    var workerConnectionData = new Array(workerConnectionNameLength);
    var workerCountData = new Array(workerConnectionNameLength);
    var workerSpecIdData = new Array(workerConnectionNameLength);
    for(var i=0; i<workerConnectionNameLength; i++){                          
        workerConnectionData[i] = $("input[name='workerConnectionName']")[i].value;
        workerCountData[i] = $("input[name='controlPlaneCount']")[i].value;
        workerSpecIdData[i] = $("input[name='workerSpecId']")[i].value;
    }
    
    var new_obj = {}
    // mcis 생성이므로 mcisID가 없음
    new_obj['name'] = mcksName
    
    var new_kubernetes = {}
    new_kubernetes['networkCni'] = kubernatesNetworkCni;
    new_kubernetes['podCidr'] = kubernatesPodCidr;
    new_kubernetes['serviceCidr'] = kubernatesServiceCidr;
    new_kubernetes['serviceDnsDomain'] = kubernatesServiceDnsDomain;

    new_obj['config'] = new_kubernetes
    var controlPlanes = new Array(controlPlaneConnectionLength);
    console.log("controlPlaneConnectionLength " + controlPlaneConnectionLength)
    for(var i=0; i<controlPlaneConnectionLength; i++){
        console.log("controlPlane " + i)
        var new_controlPlane = {}
        new_controlPlane['connection'] = controlPlaneConnectionData[i];
        new_controlPlane['count'] = controlPlaneCountData[i]
        new_controlPlane['spec'] = controlPlaneSpecIdData[i]
        controlPlanes[i] = new_controlPlane
    }
    new_obj['controlPlane'] = controlPlanes;

    var workers = new Array();
    $("input[name='workerCount']").each(function (i) {
        var new_worker = {}
        console.log($("select[name='workerConnectionName']").eq(i));
        new_worker['connection'] = $("select[name='workerConnectionName']").eq(i).attr("value");
        new_worker['count'] = $("input[name='workerCount']").eq(i).attr("value")
        new_worker['spec'] = $("select[name='workerSpecId']").eq(i).attr("value")
        workers[i] = new_worker
        console.log( i + "번째  : " );
        console.log( new_worker);
   });


    new_obj['worker'] = workers;
    console.log(new_obj);
}