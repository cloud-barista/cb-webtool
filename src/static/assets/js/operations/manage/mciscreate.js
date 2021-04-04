$(document).ready(function(){
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
    if(!mcis_name){
        alert("Please Input MCIS Name!!!!!")
        return;
    }
    var mcis_desc = $("#mcis_desc").val();
    var placement_algo = $("#placement_algo").val();
    var installMonAgent = $("#installMonAgent").val();
    console.log(Simple_Server_Config_Arr)
    var new_obj = {}
    var apiInfo = ApiInfo;
    new_obj['name'] = mcis_name
    new_obj['description'] = mcis_desc
    new_obj['installMonAgent'] = installMonAgent

    if(Simple_Server_Config_Arr){
        vm_len = Simple_Server_Config_Arr.length;			
        console.log("Simple_Server_Config_Arr length: ",vm_len);
        new_obj['vm'] = Simple_Server_Config_Arr;
        console.log("new obj is : ",new_obj);
        var url = CommonURL+"/ns/"+NAMESPACE+"/mcis";
        try{
            AjaxLoadingShow(true);
            axios.post(url,new_obj,{
                headers :{
                    'Content-type': 'application/json',
                    'Authorization': apiInfo,
                    },
            }).then(result=>{
            console.log("MCIR Register data : ",result);
            console.log("Result Status : ",result.status); 
            if(result.status == 201 || result.status == 200){
                alert("Register Success")
                location.href = "/Manage/MCIS/list";
            }else{
                alert("Register Fail")
                //location.reload(true);
            }
            })
        }finally{
            AjaxLoadingShow(false);
        }  



    }else{
        alert("Please Input Servers");
        $(".simple_servers_config").addClass("active");
        $("#s_name").focus();
    }
}

$(document).ready(function() {
    //OS_HW popup table scrollbar
  $('#OS_HW .btn_spec').on('click', function() {
        $('#OS_HW_Spec .dtbox.scrollbar-inner').scrollbar();
    });
    //Security popup table scrollbar
  $('#Security .btn_edit').on('click', function() {
    $("#security_edit").modal();
        $('#security_edit .dtbox.scrollbar-inner').scrollbar();
    });
});