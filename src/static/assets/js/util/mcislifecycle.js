
// MCIS 제어 : 선택한 VM의 상태 변경 
function mcisLifeCycle(type){
    var checked_nothing = 0;
    $("[id^='td_ch_']").each(function(){
       
        if($(this).is(":checked")){
            checked_nothing++;
            console.log("checked")
            var mcisID = $(this).val()
            console.log("check td value : ",mcis_id);
            // var nameSpace = NAMESPACE;
            console.log("Start LifeCycle method!!!")
            // var url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"?action="+type
            var url = "/operation/manage" + "/mcis/proc/mcislifecycle";
            
            console.log("life cycle3 url : ",url);
            var message = "MCIS "+type+ " complete!."
            // var apiInfo = ApiInfo
            // axios.get(url,{
            //     headers:{
            //         'Authorization': apiInfo
            //     }
            axios.post(url,{
                headers: { },
                mcisID:mcisID,
                lifeCycleType:type
            }).then(result=>{
                var status = result.status
                
                console.log("life cycle result : ",result)
                var data = result.data
                console.log("result Message : ",data.message)
                if(status == 200 || status == 201){
                    
                    alert(message);
                    location.reload();
                    //show_mcis(mcis_url,"");
                }else{
                    alert(status)
                    return;
                }
            })
        }else{
            console.log("checked nothing")
           
        }
    })
    if(checked_nothing == 0){
        alert("Please Select MCIS!!")
        return;
    }
}
////////////// MCIS Handling end //////////////// 