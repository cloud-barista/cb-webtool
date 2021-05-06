$(document).ready(function(){
    //OS_HW popup table scrollbar
    $('#OS_HW .btn_spec').on('click', function() {
        $('#OS_HW_Spec .dtbox.scrollbar-inner').scrollbar();
    });
    //Security popup table scrollbar
    $('#Security .btn_edit').on('click', function() {
    $("#security_edit").modal();
        $('#security_edit .dtbox.scrollbar-inner').scrollbar();
    });

    //Servers Expert on/off
    var check = $(".switch .ch");
    var $Servers = $(".servers_config");
    var $NewServers = $(".new_servers_config");
    var $SimpleServers = $(".simple_servers_config");
    var simple_config_cnt = 0;
    var expert_config_cnt = 0;
    
    check.click(function(){
        $(".switch span.txt_c").toggle();
        $NewServers.removeClass("active");
    });
   
  //Expert add
    $('.servers_box .server_add').click(function(){
        $NewServers.toggleClass("active");
      if($Servers.hasClass("active")) {
        $Servers.toggleClass("active");
    } else {
        $Servers.toggleClass("active");
    }
    });
    // Simple add
  $(".servers_box .switch").change(function() {
    if ($(".switch .ch").is(":checked")) {	
            $('.servers_box .server_add').click(function(){	
                
                $NewServers.addClass("active");
                $SimpleServers.removeClass("active");		
            });
    } else {
            $('.servers_box .server_add').click(function(){
            
                $NewServers.removeClass("active");
                $SimpleServers.addClass("active");
            
            
            });		
    }
  });
});

function btn_deploy(){
    var mcis_name = $("#mcis_name").val();
    var mcis_id = $("#mcis_id").val();
    if(!mcis_id){
        commonAlert("Please Select MCIS !!!!!")
        return;
    }
        
    if(Simple_Server_Config_Arr){// mcissimpleconfigure.js 에 const로 정의 됨.
        vm_len = Simple_Server_Config_Arr.length;			
        console.log("Simple_Server_Config_Arr length: ",vm_len);
        // var new_obj = {}
        // new_obj['vm'] = Simple_Server_Config_Arr;
        // console.log("new obj is : ",new_obj);
        // var url = "/operation/manages/mcis/:mcisID/vm/reg/proc"
        var url = "/operation/manages/mcis/" + mcis_id +"/vm/reg/proc"

        // 한개씩 for문으로 추가
        for(var i in Simple_Server_Config_Arr){
            new_obj = Simple_Server_Config_Arr[i];
            console.log("new obj is : ",new_obj);
            try{
                resultVmCreateMap.set(new_obj.name, "")
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

            // post로 호출을 했으면 해당 VM의 정보는 비활성시킨 후(클릭 Evnet 안먹게)
            // 상태값을 모니터링 하여 결과 return 까지 대기.
            // return 받으면 해당 VM
        } 
    }
}

// vm 생성 결과 표시
// 여러개의 vm이 생성될 수 있으므로 각각 결과를 표시
var resultVmCreateMap = new Map();
function vmCreateCallback(resultVmKey, resultStatus){
    resultVmCreateMap.set(resultVmKey, resultStatus)
    var resultText = "";
    for (let key of resultVmCreateMap.keys()) { 
        console.log("vmCreateresult " + key + " : " + resultVmCreateMap.get(resultVmKey) );
        resultText += key + " = " + resultVmCreateMap.get(resultVmKey) + ","
    }

    $("#serverRegistResult").text(resultText)
}