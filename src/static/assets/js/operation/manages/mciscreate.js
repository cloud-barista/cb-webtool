$(document).ready(function(){
    //Servers Expert on/off
//     var check = $(".switch .ch");
//     var $Servers = $(".servers_config");
//     var $NewServers = $(".new_servers_config");
//     var $SimpleServers = $(".simple_servers_config");
//     var simple_config_cnt = 0;
//     var expert_config_cnt = 0;
    
//     check.click(function(){
//         $(".switch span.txt_c").toggle();
//         $NewServers.removeClass("active");
//     });
   
//   //Expert add
//     $('.servers_box .server_add').click(function(){
//         $NewServers.toggleClass("active");
//         if($Servers.hasClass("active")) {
//         $Servers.toggleClass("active");
//         } else {
//             $Servers.toggleClass("active");
//         }
//     });
//     // Simple add
//     $(".servers_box .switch").change(function() {
//         if ($(".switch .ch").is(":checked")) {	
//                 $('.servers_box .server_add').click(function(){	
                    
//                     $NewServers.addClass("active");
//                     $SimpleServers.removeClass("active");		
//                 });
//         } else {
//             $('.servers_box .server_add').click(function(){
            
//                 $NewServers.removeClass("active");
//                 $SimpleServers.addClass("active");
            
            
//             });		
//         }
//     });
});

// 서버 더하기버튼 클릭시 서버정보 입력area 보이기/숨기기
// isExpert의 체크 여부에 따라 바뀜.
// newServers 와 simpleServers가 있음.
function displayNewServerForm(){
    
    var expertServerConfig = $("#expertServerConfig");
    var simpleServerConfig = $("#simpleServerConfig");
    if ($("#isExpert").is(":checked")) {
        //expertServerConfig
        if(expertServerConfig.hasClass("active")) {
            expertServerConfig.removeClass("active");
        }else{
            expertServerConfig.addClass("active");
            simpleServerConfig.removeClass("active");
        }
    }else{
        //simpleServerConfig
        if(simpleServerConfig.hasClass("active")) {
            simpleServerConfig.removeClass("active");
        }else{
            expertServerConfig.removeClass("active");
            simpleServerConfig.addClass("active");
        }
    }

}

// 서버정보 입력 area에서 'DONE'버튼 클릭시 array에 담고 form을 초기화


// deploy 버튼 클릭시 등록한 서버목록을 배포.
// function btn_deploy(){
function deployMcis(){
    var mcis_name = $("#mcis_name").val();
    if(!mcis_name){
        commonAlert("Please Input MCIS Name!!!!!")
        return;
    }
    var mcis_desc = $("#mcis_desc").val();
    var placement_algo = $("#placement_algo").val();
    var installMonAgent = $("#installMonAgent").val();
    console.log(Simple_Server_Config_Arr)
    var new_obj = {}
    // mcis 생성이므로 mcisID가 없음
    new_obj['name'] = mcis_name
    new_obj['description'] = mcis_desc
    new_obj['installMonAgent'] = installMonAgent

    if(Simple_Server_Config_Arr){
        vm_len = Simple_Server_Config_Arr.length;			
        console.log("Simple_Server_Config_Arr length: ",vm_len);
        new_obj['vm'] = Simple_Server_Config_Arr;
        console.log("new obj is : ",new_obj);
        // var url = CommonURL+"/ns/"+NAMESPACE+"/mcis";
        var url = "/operation/manages/mcis/reg/proc"
        try{
            // AjaxLoadingShow(true);// interceptor 에서 loading 보여줌
            axios.post(url,new_obj,{
                headers :{
                    'Content-type': 'application/json',
                    // 'Authorization': apiInfo,
                    },
            }).then(result=>{
                console.log("MCIR Register data : ",result);
                console.log("Result Status : ",result.status); 
                if(result.status == 201 || result.status == 200){
                    commonAlert("Register Success")
                    // location.href = "/Manage/MCIS/list";
                    // $('#loadingContainer').show();
                    // location.href = "/operation/manages/mcis/mngform/"
                    var targetUrl = "/operation/manages/mcis/mngform"
					changePage(targetUrl)
                }else{
                    commonAlert("Register Fail")
                    //location.reload(true);
                }
            })
        }finally{
            // AjaxLoadingShow(false);
        }  
    }else{
        alert("Please Input Servers");
        $(".simple_servers_config").addClass("active");
        $("#s_name").focus();
    }
}

$(document).ready(function() {
    //OS_HW popup table scrollbar
//   $('#OS_HW .btn_spec').on('click', function() {
//         $('#OS_HW_Spec .dtbox.scrollbar-inner').scrollbar();
//     });
//     //Security popup table scrollbar
//   $('#Security .btn_edit').on('click', function() {
//     $("#security_edit").modal();
//         $('#security_edit .dtbox.scrollbar-inner').scrollbar();
//     });
});