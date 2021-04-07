///////////// MCIS ADD

// 등록 form으로 이동
function createNewMCIS(){// Manage_MCIS_Life_Cycle_popup.html
    var url = "/operation/manage" + "/mcis/regform/"
    // location.href = "/Manage/MCIS/reg"
    location.href = url;
}



////////////// MCIS ADD 



const config_arr = new Array();

// CP / connection area 조회 :: TODO : 필요 있을까?
function getConnection(){
   var apiInfo = ApiInfo;
   $.ajax({
       url: SpiderURL+"/connectionconfig",
       async:false,
       type:'GET',
       beforeSend : function(xhr){
           xhr.setRequestHeader("Authorization", apiInfo);
           xhr.setRequestHeader("Content-type","application/json");
       },
      

   }).done( function(data2){
       res = data2.connectionconfig
       
       console.log("connection info : ",res);
       var provider = "";
       var aws_cnt = 0;
       var gcp_cnt = 0;
       var azure_cnt = 0;
       var open_cnt = 0;
       var cloudIt_cnt = 0;
       var ali_cnt = 0;
       var cp_cnt = 0;
       var connection_cnt = 0;
       var html = "";
       for(var k in res){
           config_arr.push(res[k])
           provider = res[k].ProviderName 
           connection_cnt++;
           provider = provider.toLowerCase();
           console.log("provider lowercase : ",provider);
           
           if(provider == "aws"){
               aws_cnt++;  
            
           }
           if(provider == "azure"){
               azure_cnt++;
                
           }
           if(provider == "alibaba"){
               ali_cnt++;
             
                   
           }
           if(provider == "gcp"){
               gcp_cnt++;
           
           }
           if(provider == "cloudit"){
               cloudIt_cnt++;
             
           }
           if(provider == "openstack"){
               open_cnt++;
             
           }
       }
       
       
       if(aws_cnt > 0 ){
          
           html +='<li class="bg_b">'
                +'<a href="#!"><span>AWS('
                +aws_cnt
                +')</span></a></li>';          
       }
       if(azure_cnt > 0){
           html +='<li class="bg_y">'
                +'<a href="#!"><span>AZ('
                +azure_cnt
                +')</span></a></li>';       
       }
       if(ali_cnt > 0){
          
           html +='<li class="bg_r">'
                +'<a href="#!"><span>ALI('
                +ali_cnt
                +')</span></a></li>';       
               
       }
       if(gcp_cnt > 0){
         
           html +='<li class="bg_g">'
           +'<a href="#!"><span>GCP('
           +gcp_cnt
           +')</span></a></li>';     
       }
       if(cloudIt_cnt > 0){
         
           html +='<li class="bg_n">'
           +'<a href="#!"><span>CLIT('
           +cloudIt_cnt
           +')</span></a></li>';  
       }
       if(open_cnt > 0){
          
           html +='<li class="bg_b">'
           +'<a href="#!"><span>OPS('
           +open_cnt
           +')</span></a></li>';  
       }

       if(aws_cnt > 1){
           aws_cnt = 1
       }
       if(azure_cnt > 1){
           azure_cnt = 1
       }
       if(ali_cnt > 1){
           ali_cnt = 1
       }
       if(open_cnt > 1){
           open_cnt = 1
       }
       if(cloudIt_cnt > 1){
           cloudIt_cnt = 1
       }
       if(gcp_cnt > 1){
           gcp_cnt = 1
       }

       cp_cnt = aws_cnt+azure_cnt+ali_cnt+open_cnt+cloudIt_cnt+gcp_cnt;
       var str = '<strong>'+cp_cnt+'</strong><span>/</span>'+connection_cnt;
       $("#cpConnectionTotal").empty();
       $("#cpConnectionTotal").append(str);
       $("#cpConnectionDetail").empty();
       $("#cpConnectionDetail").append(html);
   })
   
}


// List Of MCIS 클릭 시 
// mcis 테이블의 선택한 row 외의 다른 row들은 안보이게
// 해당 mcis의 상세 정보를 표시하는 MCIS INFO area 보이게
function displayListOfMCIS(id,index){
    console.log("click view mcis id :",id)
    $(".server_status").addClass("view");
    
    // 다른 area가 활성화 되어 있으면 안보이게
    $("#dashboard_detailBox").removeClass("active")

    // List Of MCIS에서 선택한 row 외에는 안보이게
    $("[id^='server_info_tr_']").each(function(){
        var item = $(this).attr("item").split("|")
        console.log()
        if(id == item[0]){
           
            $(this).addClass("on")
        }else{
            $(this).removeClass("on")
        }
    })

    // MCIS Info 에 mcis id 표시
    $("#mcis_id").val(id);

    // MCIS Info area set
    showServerListAndStatusArea(id,index);
}

const mcisInfoData = new Array()// test_arr : mcisInfo 1개 전체, pageLoad될 때 조회
// MCIS Info area 안의 Server List / Status 내용 표시
// 해당 MCIS의 모든 VM 표시

function showServerListAndStatusArea(mcis_id, index){
   $(".server_status").addClass("view")
   var mcis_arr = mcisInfoData.filter(item => item.id === mcis_id)
   var mcis = mcis_arr[0];
   console.log("showServerListAndStatusArea " )
   console.log(mcis_arr)
   $("#mcis_name").val(mcis.name)
   console.log("showmcis2 Data : ",mcis)
   var mcis_badge = "";
   var sta = mcis.status;
   var sl = sta.split("-");
   var status = sl[0].toLowerCase()
   var vms = mcis.vm
   var vm_len = 0
   var provider = new Array();
   var vm_badge = ""
   if(vms){
        vm_len = vms.length
       for(var o in vms){
           var vm_status = vms[o].status
           var lat = vms[o].location.latitude
           var long = vms[o].location.longitude
           provider.push(vms[o].location.cloudType)

           if(vm_status == "Running"){
               vm_badge += '<li class="sel_cr bgbox_b" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\')"><a href="javascript:void(0);" ><span class="txt">'+vms[o].name+'</span></a></li>';
               
           }else if(vm_status == "include" ){
               vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\',\''+vms[o].name+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
           }else if(vm_status == "Suspended"){
               vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\',\''+vms[o].name+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
               
           }else if(vm_status == "Terminated"){
               vm_badge += '<li class="sel_cr bgbox_r"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\',\''+vms[o].name+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
               
           }else{
               vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\',\''+vms[o].name+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
           }
           console.log("vm_status : ", vm_status)

       }
       $("#mcis_server_info_box").empty();
       $("#mcis_server_info_box").append(vm_badge);
   }

   var csp = ""
   var new_provider  = provider.filter((item,index, arr)=>(arr.indexOf(item) === index))
   if(new_provider){
       if(new_provider.length > 1){
           csp = new_provider.join(",")
       }else if(new_provider.length == 1){
           csp = new_provider[0]
       }
   }
   $("#mcis_info_cloud_connection").val(csp)

   console.log("mcis Status 1: ", mcis.status)
   console.log("mcis Status 2: ", status)
   if(status == "running"){
       mcis_badge = '<img src="/assets/img/contents/icon_running_db.png" alt=""/> '
   }else if(status == "include" ){
       mcis_badge = '<img src="/assets/img/contents/icon_stop_db.png"  alt=""/> '
   }else if(status == "suspended"){
       mcis_badge = '<img src="/assets/img/contents/icon_stop_db.png" alt=""/>'
   }else if(status == "terminate"){
       mcis_badge = '<img src="/assets/img/contents/icon_terminate_db.png" alt=""/>'
   }else{
       mcis_badge = '<img src="/assets/img/contents/icon_stop_db.png" alt=""/>'
   }
   $("#service_status_icon").empty();
   $("#service_status_icon").append(mcis_badge)

   var mcis_name = mcis.name
   var mcis_id = mcis.id
   var targetStatus = mcis.targetStatus
   var targetAction = mcis.targetAction
   var description = mcis.description

   $("#mcis_info_txt").text("[ "+mcis_name+" ]");
   $("#mcis_server_info_status").empty();
   $("#mcis_server_info_status").append('<strong>Server List / Status</strong>  <span class="stxt">[ '+mcis_name+' ]</span>  Server('+vm_len+')')

   $("#mcis_info_name").val(mcis_name+" / "+mcis_id)
   $("#mcis_info_description").val(description);
   $("#mcis_info_targetStatus").val(targetStatus);
   $("#mcis_info_targetAction").val(targetAction);

   var target = "server_info_tr_"+index
   $td_list = $("#"+target+"");
   $status = $(".server_status");
   console.log("click tr target : ",target)
   $td_list.off("click").click(function(){
       $td_list.addClass("on");
       $td_list.siblings().removeClass("on");
       $status.addClass("view");
       $status.siblings().removeClass("on");
       $(".dashboard.register_cont").removeClass("active");
       $td_list.off("click").click(function(){
               if( $(this).hasClass("on") ) {
                   
                   $td_list.removeClass("on");
                   $status.removeClass("view");
                   //$detail.removeClass("active");
           } else {
           
                   $td_list.addClass("on");
                   $td_list.siblings().removeClass("on");
                   $status.addClass("view");
                   $status.siblings().removeClass("view");
               $(".dashboard.register_cont").removeClass("active");
           }
       });
   });
       //Manage MCIS Server List on/off
   $(".dashboard .ds_cont .area_cont .listbox li.sel_cr").each(function(){
       var $sel_list = $(this),
               $detail = $(".server_info");
       $sel_list.off("click").click(function(){
             $sel_list.addClass("active");
             $sel_list.siblings().removeClass("active");
             $detail.addClass("active");
             $detail.siblings().removeClass("active");
            $sel_list.off("click").click(function(){
                 if( $(this).hasClass("active") ) {
                     $sel_list.removeClass("active");
                     $detail.removeClass("active");
             } else {
                     $sel_list.addClass("active");
                     $sel_list.siblings().removeClass("active");
                     $detail.addClass("active");
                     $detail.siblings().removeClass("active");
             }
             });
         });
     }); 
}
