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