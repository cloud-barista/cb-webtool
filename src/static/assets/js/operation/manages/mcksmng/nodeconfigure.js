function changeConnectionInfo(configName){
    console.log("config name : ",configName)
    if( configName == ""){
        // 0번째면 selectbox들을 초기화한다.(vmInfo, sshKey, image 등)
    }
    
    getSpecInfo(configName);
}

function getSpecInfo(configName){
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
        
      
        $("#controlPlaneSpecId_0").empty();
        $("#controlPlaneSpecId_0").append(html);
        
        $("#workerSpecId_0").empty();
        $("#workerSpecId_0").append(html);
    })
}