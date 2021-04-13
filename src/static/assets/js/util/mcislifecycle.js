
// MCIS 제어 : 선택한 VM의 상태 변경 
function mcisLifeCycle(type){
    
    var url = "/operation/manage" + "/mcis/proc/mcislifecycle";
    
    console.log("life cycle3 url : ",url);
    var message = "MCIS "+type+ " complete!."
           
    axios.post(url,{
        headers: { },
        mcisID:mcisID,
        lifeCycleType:type
    }).then(result=>{
        var status = result.status
        var data = result.data
        callbackMcisLifeCycle(status, data, type)
        // console.log("life cycle result : ",result)
        // console.log("result Message : ",data.message)
        // if(status == 200 || status == 201){
            
        //     alert(message);
        //     location.reload();
        //     //show_mcis(mcis_url,"");
        // }else{
        //     alert(status)
        //     return;
        // }
    }).catch(function(error){
        // console.log(" display error : ",error);
        console.log(error.response.data);
        console.log(error.response.status);
        // console.log(error.response.headers); 
        var status = error.response.status;
        var data =  error.response.data

        callbackMcisLifeCycle(status, data, type)
    });
        
    
}
////////////// MCIS Handling end //////////////// 