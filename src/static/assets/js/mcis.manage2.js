function life_cycle2(type){
    var mcis_id = $("#mcis_id").val();
    var mcis_name = $("#mcis_name").val();
    if(!mcis_id){
        alert("Please Select MCIS!!")
        return;
    }
    var nameSpace = NAMESPACE;
    console.log("Start LifeCycle method!!!")
  
    var url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"?action="+type
    var message = mcis_name+" "+type+ " complete!."
  

    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        var status = result.status
        
        console.log("life cycle result : ",result)
        var data = result.data
        console.log("result Message : ",data.message)
        if(status == 200 || status == 201){
            
            alert(message);
            location.reload();
            //show_mcis(mcis_url,"");
        }
    })
}
const test_arr = new Array()
function show_mcis_list(url){
    console.log("Show mcis Url : ",url)
    $("#vm_detail").hide();
    checkNS();
 
    var apiInfo = ApiInfo;
 
    console.log("apiInfo : ",apiInfo);
     axios.get(url,{
         headers:{
             'Authorization': apiInfo
         }
     }).then(result=>{
       
        console.log("Dashboard Data :",result.status);
        var data = result.data;
        console.log("func show_mcis result data : ",data)
        if(!data.mcis){
           location.href = "/Manage/MCIS/reg";
           return;
        }
        if(data.mcis.length == 0 ){
         location.href = "/Manage/MCIS/reg";
         return;
      }
        
         console.log("showmcis Data : ",data)
         var html = "";
         var mcis = data.mcis;
         var len = 0
         var mcis_cnt = 0
         
         if(mcis){
            len = mcis.length;
         }
         mcis_cnt = len;
         var count = 0;
         
         var server_cnt = 0;
         
         var html = "";
         var run_cnt = 0;
         var stop_cnt = 0;
         var mcis_run_cnt = 0;
         var mcis_stop_cnt = 0;
         var mcis_terminated_cnt = 0;

         var run_vm_cnt = 0;
         var stop_vm_cnt = 0;
         var terminated_vm_cnt = 0;
         
         for(var i in mcis){
            test_arr.push(mcis[i])
            count++;
            var vm_run_cnt = 0;
            var vm_stop_cnt = 0;
            var terminate_cnt = 0;
            var vm_len = 0
            var sta = mcis[i].status;
            var sl = sta.split("-");
            var mcis_badge = "";
            var vm_badge = "";
            var status = sl[0].toLowerCase()
            var vms = mcis[i].vm
            console.log("mcis status : ",status)
            var vm_status = "";
             if(vms){
                vm_len = vms.length
                server_cnt = server_cnt+vm_len;
             }
             //VM  상태 및 기타 생성하기
             var vm_cnt = 0
             var vm_html = "";
             var provider = new Array();
             for(var o in vms){
                 vm_cnt++;
                var vm_status = vms[o].status
                var lat = vms[o].location.latitude
                var long = vms[o].location.longitude
                provider.push(vms[o].location.cloudType)

                if(vm_status == "Running"){
                    vm_badge += "shot bgbox_b";
                    run_cnt++;
                    vm_run_cnt++;
                    run_vm_cnt++;
                 }else if(vm_status == "include" ){
                    vm_badge += "shot bgbox_y"
                 }else if(vm_status == "Suspended"){
                    vm_badge += "shot bgbox_y";
                    stop_cnt++;
                    vm_stop_cnt++;
                    stop_vm_cnt++;
                 }else if(vm_status == "Terminated"){
                    vm_badge += "shot bgbox_r"
                    terminate_cnt++;
                    terminated_vm_cnt++;
                 }else{
                    vm_badge += "shot bgbox_g"
                 }

             }
             
             html +='<tr onclick="click_view(\''+mcis[i].id+'\',\''+i+'\');" id="server_info_tr_'+i+'" item="'+mcis[i].id+'|'+i+'">'
             //MCIS name  / MCIS 상태
             if(status == "running"){
               html +='<td class="overlay hidden td_left" data-th="Status"><img src="/assets/img/contents/icon_running.png" class="icon" alt=""/> Running  <span class="ov off"></span></td>'
                mcis_run_cnt++;
             }else if(status == "include" ){
              
             }else if(status == "suspended"){
               html += '<td class="overlay hidden td_left" data-th="Status"><img src="/assets/img/contents/icon_stop.png" class="icon" alt=""/> Suspended <span class="ov off"></span></td>'
                mcis_stop_cnt++;
             }else if(status == "terminate"){
                html +='<td class="overlay hidden td_left" data-th="Status"><img src="/assets/img/contents/icon_terminate.png" class="icon" alt=""/> Terminate <span class="ov off"></span></td>'
                mcis_terminated_cnt;
             }else{
                
             }
           

            html +='<td class="btn_mtd ovm" data-th="Name">'+mcis[i].name+'<span class="ov"></span></td>'
            var csp = ""
            if(provider){
                if(provider.length > 1){
                    csp = provider.join(",")
                }else if(provider.length == 1){
                    csp = provider[0]
                }
            }
            html += '<td class="overlay hidden" data-th="Cloud Connection">'+csp+'</td>'
            html +='<td class="overlay hidden" data-th="Total Infras">'+vm_cnt+'</td>'
            html +='<td class="overlay hidden" data-th="# of Servers">'+vm_cnt+' <span class="bar">/</span> '+vm_run_cnt+' <span class="bar">/</span> '+vm_stop_cnt+' <span class="bar">/</span> '+terminate_cnt+'</td>'
            html +='<td class="overlay hidden" data-th="Description">'+mcis[i].description+'</td>'
            html +='<td class="overlay hidden" data-th=""><input type="checkbox" name="chk" value="'+mcis[i].id+'" id="td_ch_'+i+'" title="" /><label for="td_ch_'+i+'"></label></td>'
            html +='</tr>'


             
            
        }
        // 새로운 퍼블리싱에 넣을 값
        $("#total_mcis").text(mcis_cnt);
        // 각각의  MCIS의 상태 별 갯수
        var mcis_numbox = '<div class="num bgbox_b"><span>'+mcis_run_cnt+'</span></div>'
                         +'<div class="num bgbox_r"><span>'+mcis_stop_cnt+'</span></div>'
                         +'<div class="num bgbox_g"><span>'+mcis_terminated_cnt+'</span></div>';
        
        //  서버 갯수 및 상태 값 붙여 넣기
        $("#mcis_numbox").empty();
        $("#mcis_numbox").append(mcis_numbox);
        
        // vm cnt server_cnt
        $("#total_vm").text(server_cnt);
        var vm_numbox = '<div class="num bgbox_b boxrd cursor" onclick="location.href=\'../operation/Manage_Mcis.html\'"><span>'+run_vm_cnt+'</span></div>'
        +'<div class="num bgbox_r boxrd cursor" onclick="location.href=\'../operation/Manage_Mcis.html\'"><span>'+stop_vm_cnt+'</span></div>'
        +'<div class="num bgbox_g boxrd cursor" onclick="location.href=\'../operation/Manage_Mcis.html\'"><span>'+terminated_vm_cnt+'</span></div>';
        $("#vm_numbox").empty();
        $("#vm_numbox").append(vm_numbox);
        // mcis list add
        $("#table_1").empty();
        $("#table_1").append(html);


   
        //event 속성
       
        $("#th_chall").click(function() {
            if ($("#th_chall").prop("checked")) {
                $("input[name=chk]").prop("checked", true);
            } else {
                $("input[name=chk]").prop("checked", false);
            }
        })

        $(".dashboard .status_list tbody tr").each(function(){
            var $td_list = $(this),
                    $status = $(".server_status"),
                    $detail = $(".server_info");
          
            $td_list.off("click").click(function(){
                  $td_list.addClass("on");
                  $td_list.siblings().removeClass("on");
                  $status.addClass("view");
                  $status.siblings().removeClass("on");
                $(".dashboard.register_cont").removeClass("active");
                 $td_list.off("click").click(function(){
                      if( $(this).hasClass("on") ) {
                          console.log("1번에 걸린다.")
                          $td_list.removeClass("on");
                          $status.removeClass("view");
                          $detail.removeClass("active");
                  } else {
                    console.log("아니다 2번에 걸린다.")
                          $td_list.addClass("on");
                          $td_list.siblings().removeClass("on");
                          $status.addClass("view");
                          $status.siblings().removeClass("view");
                        $(".dashboard.register_cont").removeClass("active");
                  }
                  });
              });
          });
          
        $(window).on("load resize",function(){
            var vpwidth = $(window).width();
            if (vpwidth > 768 && vpwidth < 1800) {
                $(".dashboard_cont .dataTable").addClass("scrollbar-inner");
                    $(".dataTable.scrollbar-inner").scrollbar();
            } else {
                $(".dashboard_cont .dataTable").removeClass("scrollbar-inner");
            }
        });
    

    }).catch(function(error){
     console.log("show mcis error at dashboard js: ",error);
    });
 }

 function show_mcis_list2(url){
    console.log("Show mcis Url : ",url)
    $("#vm_detail").hide();
    checkNS();
 
    var apiInfo = ApiInfo;
 
    console.log("apiInfo : ",apiInfo);
     axios.get(url,{
         headers:{
             'Authorization': apiInfo
         }
     }).then(result=>{
       
        console.log("Dashboard Data :",result.status);
        var data = result.data;
        console.log("func show_mcis result data : ",data)
        if(!data.mcis){
           location.href = "/Manage/MCIS/reg";
           return;
        }
        if(data.mcis.length == 0 ){
         location.href = "/Manage/MCIS/reg";
         return;
      }
        
         console.log("showmcis Data : ",data)
         var html = "";
         var mcis = data.mcis;
         var len = 0
         var mcis_cnt = 0
         
         if(mcis){
            len = mcis.length;
         }
         mcis_cnt = len;
         var count = 0;
         
         var server_cnt = 0;
         
         var html = "";
         var run_cnt = 0;
         var stop_cnt = 0;
         var mcis_run_cnt = 0;
         var mcis_stop_cnt = 0;
         var mcis_terminated_cnt = 0;

         var run_vm_cnt = 0;
         var stop_vm_cnt = 0;
         var terminated_vm_cnt = 0;
         
         for(var i in mcis){
            test_arr.push(mcis[i])
            count++;
            var vm_run_cnt = 0;
            var vm_stop_cnt = 0;
            var terminate_cnt = 0;
            var vm_len = 0
            var sta = mcis[i].status;
            var sl = sta.split("-");
            var mcis_badge = "";
            var vm_badge = "";
            var status = sl[0].toLowerCase()
            var vms = mcis[i].vm
            console.log("mcis status : ",status)
            var vm_status = "";
             if(vms){
                vm_len = vms.length
                server_cnt = server_cnt+vm_len;
             }
             //VM  상태 및 기타 생성하기
             var vm_cnt = 0
             var vm_html = "";
             var provider = new Array();
             for(var o in vms){
                 vm_cnt++;
                var vm_status = vms[o].status
                var lat = vms[o].location.latitude
                var long = vms[o].location.longitude
                provider.push(vms[o].location.cloudType)

                if(vm_status == "Running"){
                    vm_badge += "shot bgbox_b";
                    run_cnt++;
                    vm_run_cnt++;
                    run_vm_cnt++;
                 }else if(vm_status == "include" ){
                    vm_badge += "shot bgbox_y"
                 }else if(vm_status == "Suspended"){
                    vm_badge += "shot bgbox_y";
                    stop_cnt++;
                    vm_stop_cnt++;
                    stop_vm_cnt++;
                 }else if(vm_status == "Terminated"){
                    vm_badge += "shot bgbox_r"
                    terminate_cnt++;
                    terminated_vm_cnt++;
                 }else{
                    vm_badge += "shot bgbox_g"
                 }

             }
             
             html +='<tr onclick="click_view(\''+mcis[i].id+'\',\''+i+'\');" id="server_info_tr_'+i+'" item="'+mcis[i].id+'|'+i+'">'
             //MCIS name  / MCIS 상태
             if(status == "running"){
               html +='<td class="overlay hidden td_left" data-th="Status"><img src="/assets/img/contents/icon_running.png" class="icon" alt=""/> Running  <span class="ov off"></span></td>'
                mcis_run_cnt++;
             }else if(status == "include" ){
              
             }else if(status == "suspended"){
               html += '<td class="overlay hidden td_left" data-th="Status"><img src="/assets/img/contents/icon_stop.png" class="icon" alt=""/> Suspended <span class="ov off"></span></td>'
                mcis_stop_cnt++;
             }else if(status == "terminate"){
                html +='<td class="overlay hidden td_left" data-th="Status"><img src="/assets/img/contents/icon_terminate.png" class="icon" alt=""/> Terminate <span class="ov off"></span></td>'
                mcis_terminated_cnt;
             }else{
                
             }
           

            html +='<td class="btn_mtd ovm" data-th="Name">'+mcis[i].name+'<span class="ov"></span></td>'
            var csp = ""
            if(provider){
                if(provider.length > 1){
                    csp = provider.join(",")
                }else if(provider.length == 1){
                    csp = provider[0]
                }
            }
            html += '<td class="overlay hidden" data-th="Cloud Connection">'+csp+'</td>'
            html +='<td class="overlay hidden" data-th="Total Infras">'+vm_cnt+'</td>'
            html +='<td class="overlay hidden" data-th="# of Servers">'+vm_cnt+' <span class="bar">/</span> '+vm_run_cnt+' <span class="bar">/</span> '+vm_stop_cnt+' <span class="bar">/</span> '+terminate_cnt+'</td>'
            html +='<td class="overlay hidden" data-th="Description">'+mcis[i].description+'</td>'
            html +='<td class="overlay hidden" data-th=""><input type="checkbox" name="chk" value="'+mcis[i].id+'" id="td_ch_'+i+'" title="" /><label for="td_ch_'+i+'"></label></td>'
            html +='</tr>'


             
            
        }
        // 새로운 퍼블리싱에 넣을 값
        $("#total_mcis").text(mcis_cnt);
        // 각각의  MCIS의 상태 별 갯수
        var mcis_numbox = '<div class="num bgbox_b"><span>'+mcis_run_cnt+'</span></div>'
                         +'<div class="num bgbox_r"><span>'+mcis_stop_cnt+'</span></div>'
                         +'<div class="num bgbox_g"><span>'+mcis_terminated_cnt+'</span></div>';
        
        //  서버 갯수 및 상태 값 붙여 넣기
        $("#mcis_numbox").empty();
        $("#mcis_numbox").append(mcis_numbox);
        
        // vm cnt server_cnt
        $("#total_vm").text(server_cnt);
        var vm_numbox = '<div class="num bgbox_b boxrd cursor" onclick="location.href=\'../operation/Manage_Mcis.html\'"><span>'+run_vm_cnt+'</span></div>'
        +'<div class="num bgbox_r boxrd cursor" onclick="location.href=\'../operation/Manage_Mcis.html\'"><span>'+stop_vm_cnt+'</span></div>'
        +'<div class="num bgbox_g boxrd cursor" onclick="location.href=\'../operation/Manage_Mcis.html\'"><span>'+terminated_vm_cnt+'</span></div>';
        $("#vm_numbox").empty();
        $("#vm_numbox").append(vm_numbox);
        // mcis list add
        $("#table_1").empty();
        $("#table_1").append(html);


   
        //event 속성
       
        $("#th_chall").click(function() {
            if ($("#th_chall").prop("checked")) {
                $("input[name=chk]").prop("checked", true);
            } else {
                $("input[name=chk]").prop("checked", false);
            }
        })

        $(".dashboard .status_list tbody tr").each(function(){
            var $td_list = $(this),
                    $status = $(".server_status"),
                    $detail = $(".server_info");
          
            $td_list.off("click").click(function(){
                  $td_list.addClass("on");
                  $td_list.siblings().removeClass("on");
                  $status.addClass("view");
                  $status.siblings().removeClass("on");
                $(".dashboard.register_cont").removeClass("active");
                 $td_list.off("click").click(function(){
                      if( $(this).hasClass("on") ) {
                          console.log("1번에 걸린다.")
                          $td_list.removeClass("on");
                          $status.removeClass("view");
                          $detail.removeClass("active");
                  } else {
                    console.log("아니다 2번에 걸린다.")
                          $td_list.addClass("on");
                          $td_list.siblings().removeClass("on");
                          $status.addClass("view");
                          $status.siblings().removeClass("view");
                        $(".dashboard.register_cont").removeClass("active");
                  }
                  });
              });
          });
          
        $(window).on("load resize",function(){
            var vpwidth = $(window).width();
            if (vpwidth > 768 && vpwidth < 1800) {
                $(".dashboard_cont .dataTable").addClass("scrollbar-inner");
                    $(".dataTable.scrollbar-inner").scrollbar();
            } else {
                $(".dashboard_cont .dataTable").removeClass("scrollbar-inner");
            }
        });
        var mcis_id = $("#mcis_id").val()
        var mcis_name = $("#mcis_name").val()
        if(mcis_id){
            console.log("여기에 걸려야 함")
            var select_index = "";
            $("[id^='server_info_tr_']").each(function(){
                var item = $(this).attr("item").split("|")
                console.log("get item :", item);
                if(mcis_id == item[0]){
                    select_index = item[1];
                    $(this).addClass("on")
                }else{
                    $(this).removeClass("on")
                }
            })
            $(".server_status").addClass("view")   
            show_mcis2(mcis_id,select_index)
        }

    }).catch(function(error){
     console.log("show mcis error at dashboard js: ",error);
    });
 }
 function click_view(id,index){
     console.log("click view mcis id :",id)
    console.log("test_arr : ",test_arr);
    $(".server_status").addClass("view")
    $("[id^='server_info_tr_']").each(function(){
        var item = $(this).attr("item").split("|")
        
        if(id == item[0]){
            
            $(this).addClass("on")
        }else{
            $(this).removeClass("on")
        }
    })
    show_mcis2(id,index);
    
 }
 function show_mcis2(mcis_id, index){
    $(".server_status").addClass("view")
    var mcis_arr = test_arr.filter(item => item.id === mcis_id)
    var mcis = mcis_arr[0];
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
                vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
            }else if(vm_status == "Suspended"){
                vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
                
            }else if(vm_status == "Terminated"){
                vm_badge += '<li class="sel_cr bgbox_r"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
                
            }else{
                vm_badge += "shot bgbox_g"
            }

        }
        $("#mcis_server_info_box").empty();
        $("#mcis_server_info_box").append(vm_badge);
    }

    var csp = ""
    if(provider){
        if(provider.length > 1){
            csp = provider.join(",")
        }else if(provider.length == 1){
            csp = provider[0]
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

 function click_view_vm(mcis_id,vm_id){
    var select_mcis = test_arr.filter(mcis => mcis.id === mcis_id);
    console.log("click_view_vm arr : ",select_mcis);
    
    var vm_arr = select_mcis[0].vm
    console.log("vm_arr : ",vm_arr);
    vm_arr = vm_arr.filter(item => item.id === vm_id)
    console.log("click_view_vm arr : ",vm_arr)
    var mcis_name = select_mcis[0].name
    var select_vm = vm_arr[0];
    var vm_detail = select_vm.cspViewVmDetail
    var vm_name = select_vm.name

    $("#server_info_text").text('['+vm_name+'/'+mcis_name+']')
    $("#server_detail_info_text").text('['+vm_name+'/'+mcis_name+']')

    var vm_status = select_vm.status
    var vm_badge =""
    if(vm_status == "Running"){
        vm_badge = '<img src="/assets/img/contents/icon_running_db.png" alt=""/> '
    }else if(vm_status == "include" ){
        vm_badge = '<img src="/assets/img/contents/icon_stop_db.png"  alt=""/> ';
    }else if(vm_status == "Suspended"){
        vm_badge = '<img src="/assets/img/contents/icon_stop_db.png" alt=""/>'
        
    }else if(vm_status == "Terminated"){
        vm_badge = '<img src="/assets/img/contents/icon_terminate_db.png" alt=""/>'
        
    }else{
        vm_badge = '<img src="/assets/img/contents/icon_stop_db.png" alt=""/>'
    
    }
    $("#server_detail_view_server_status").val(vm_status);
    $("#server_info_status_img").empty()
    $("#server_info_status_img").append(vm_badge)

    $("#server_info_name").val(vm_name +"/"+ select_vm.id)
    $("#server_info_desc").val(select_vm.description)

    // ip information
    $("#server_info_public_ip").val(select_vm.publicIP)
    $("#server_detail_info_public_ip_text").text("Public IP : "+select_vm.publicIP)
    $("#server_info_public_dns").val(select_vm.publicDNS)
    $("#server_info_private_ip").val(select_vm.privateIP)
    $("#server_info_private_dns").val(select_vm.privateDNS)


    $("#server_detail_view_public_ip").val(select_vm.publicIP)
    $("#server_detail_view_public_dns").val(select_vm.publicDNS)
    $("#server_detail_view_private_ip").val(select_vm.privateIP)
    $("#server_detail_view_private_dns").val(select_vm.privateDNS)

    //cspvmdetail
    var vm_detail_keyValue = vm_detail.KeyValueList
    var architecture = vm_detail_keyValue.filter(item => item.Key === "Architecture")[0].Value
    $("#server_info_archi").val(architecture)
    $("#server_detail_view_archi").val(architecture)

    // server spec
    var vm_spec_name = vm_detail.VMSpecName
    $("#server_info_vmspec_name").val(vm_spec_name)
    $("#server_detail_view_server_spec_text").text(vm_spec_name)

    // start time
    var start_time = vm_detail.StartTime
    $("#server_info_start_time").val(start_time)

    // cloud type
    var csp = select_vm.location.cloudType
    var csp_icon = ""
    if(csp == "aws"){
        csp_icon = '<img src="/assets/img/contents/img_logo_a.png" alt=""/>'
    }

    $("#server_info_csp_icon").empty()
    $("#server_info_csp_icon").append(csp_icon)
    $("#server_connection_view_csp").val(csp)


    // region zone locate
    var locate = select_vm.location.briefAddr
    var region = select_vm.region.Region
    var zone = select_vm.region.Zone

    $("#server_info_region").val(locate +":"+region)
    $("#server_info_zone").val(zone)
    $("#server_info_cspVMID").val("cspVMID : "+vm_detail.IId.NameId)

    $("#server_detail_view_region").val(locate +":"+region)
    $("#server_detail_view_zone").val(zone)

    $("#server_connection_view_region").val(locate +"("+region+")")
    $("#server_connection_view_zone").val(zone)

    // connection name
    var connection_name = select_vm.connectionName;
    $("#server_info_connection_name").val(connection_name)
    $("#server_connection_view_connection_name").val(connection_name)

    // credential and driver info
    console.log("config arr2 : ",config_arr)
    console.log("connection_name :",connection_name)
    var arr_config = config_arr
    console.log("arr_config : ",arr_config);
    if(arr_config){
        var config_info = arr_config.filter(cred => cred.ConfigName === connection_name)[0]
        console.log("inner config info : ",config_info)
        console.log("config_info : ",config_info)
        var credentialName = config_info.CredentialName
        var driverName = config_info.DriverName
        $("#server_connection_view_credential_name").val(credentialName)
        $("#server_connection_view_driver_name").val(driverName)
    }
   

    
    
    // server id / system id
    $("#server_detail_view_server_id").val(select_vm.id)
    // systemid 를 사용할 경우 아래 꺼 사용
    //$("#server_detail_view_server_id").val(vm_detail.IId.SystemId)
   
    // image id
    var imageIId = vm_detail.ImageIId.NameId
    var imageId = select_vm.imageId
    $("#server_detail_view_image_id_text").text(imageId+"("+imageIId+")")

    //vpc subnet
    var vpcId = vm_detail.VpcIID.NameId
    var vpcSystemId = vm_detail.VpcIID.SystemId
    var subnetId = vm_detail.SubnetIID.NameId
    var subnetSystemId = vm_detail.SubnetIID.SystemId
    var eth = vm_detail.NetworkInterface
    $("#server_detail_view_vpc_id_text").text(vpcId+"("+vpcSystemId+")")
    $("#server_detail_view_subnet_id_text").text(subnetId+"("+subnetSystemId+")")
    $("#server_detail_view_eth_text").text(eth)

    // device info
    var root_device_type = vm_detail.VMBootDisk
    var root_device = vm_detail.VMBootDisk
    var block_device = vm_detail.VMBlockDisk
    $("#server_detail_view_root_device_type").val(root_device_type)
    $("#server_detail_view_root_device").val(root_device)
    $("#server_detail_view_block_device").val(block_device)

     // key pair info
    
     $("#server_detail_view_keypair_name").val(vm_detail.KeyPairIId.NameId)
     $("#server_detail_view_access_id_pass").val(vm_detail.VMUserId +"/"+vm_detail.VMUserPasswd)
     $("#server_detail_view_user_id_pass").val(select_vm.vmUserAccount +"/"+select_vm.vmUserPassword)

     // security Gorup
    var append_sg = ''
    var sg_arr = vm_detail.SecurityGroupIIds
    if(sg_arr){
        sg_arr.map((item,index)=>(
            append_sg +='<a href="javascript:void(0);" title="'+item.NameId+'" >'+item.NameId+'</a>'
        ))
    }
    append_sg +='인바운드 규칙 보기. 아웃바운드 규칙 보기'
    console.log("append sg : ",append_sg)
    
    $("#server_detail_view_security_group").empty()
    $("#server_detail_view_security_group").append(append_sg);
 
 }
// MCIS Control 
function life_cycle(tag,type,mcis_id,mcis_name,vm_id,vm_name){
    var url = ""
    var nameSpace = NAMESPACE;
    var message = ""
    console.log("Start LifeCycle method!!!")
    
    if(tag == "mcis"){
        url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"?action="+type
        message = mcis_name+" "+type+ " complete!."
    }else{
        url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"/vm/"+vm_id+"?action="+type
        message = vm_name+" "+type+ " complete!."
    }

    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        var status = result.status
        
        console.log("life cycle result : ",result)
        var data = result.data
        console.log("result Message : ",data.message)
        if(status == 200){
            setTimeout(function(){
                alert(data.message);
                location.reload(true);
            },5000)
            // alert(data.message);
            // location.reload(true);
        }
    })
}

function short_desc(str){
    var len = str.length;
    var result = "";
    if(len > 15){
        result = str.substr(0,15)+"...";
    }else{
        result = str;
    }

    return result;
 }
 function show_mcis(mcis_id, index){
    var nameSpace = NAMESPACE;
    var url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id
    console.log("Show mcis Url : ",url)
    
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        var mcis = result.data;
         console.log("showmcis Data : ",mcis)
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
                        vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
                    }else if(vm_status == "Suspended"){
                        vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
                      
                    }else if(vm_status == "Terminated"){
                        vm_badge += '<li class="sel_cr bgbox_r"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcis.id+'\',\''+vms[o].id+'\')"><span class="txt">'+vms[o].name+'</span></a></li>';
                      
                    }else{
                       vm_badge += "shot bgbox_g"
                    }
   
                }
                $("#mcis_server_info_box").empty();
                $("#mcis_server_info_box").append(vm_badge);
             }

             var csp = ""
             if(provider){
                 if(provider.length > 1){
                     csp = provider.join(",")
                 }else if(provider.length == 1){
                     csp = provider[0]
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
        
      
    });
 }
 const config_arr = new Array();
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
        $("#dash_2").empty();
        $("#dash_2").append(str);
        $("#dash_3").empty();
        $("#dash_3").append(html);
    })
    
}
 function show_vmList(mcis_id){
   
    var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id;
    var apiInfo = ApiInfo;
    console.log("MCIS Mangement mcisID : ",mcis_id);
    if(mcis_id){
        $.ajax({
            type:'GET',
            url:url,
            beforeSend : function(xhr){
                xhr.setRequestHeader("Authorization", apiInfo);
                xhr.setRequestHeader("Content-type","application/json");
            },
        // async:false,
            success:function(data){
                var vm = data.vm
                var mcis_name = data.name 
                $("#mcis_id").val(mcis_id)
                $("#mcis_name").val(mcis_name)
                var html = "";
                console.log("VM DATA : ",vm)
                for(var i in vm){
                    var sta = vm[i].status;
                    
                
                    var status = sta.toLowerCase()
                    console.log("VM Status : ",status)
                    var configName = vm[i].connectionName
                    console.log("outer vm configName2 : ",configName)
                    var count = 0;
                    console.log("Spider URL : ",SpiderURL)
                    $.ajax({
                        url: SpiderURL+"/connectionconfig",
                        async:false,
                        type:'GET',
                        beforeSend : function(xhr){
                            xhr.setRequestHeader("Authorization", apiInfo);
                            xhr.setRequestHeader("Content-type","application/json");
                        },
                        success : function(data2){
                            var badge = "";
                           
                            res = data2.connectionconfig
                            for(var k in res){
                                // console.log(" i value is : ",i)
                                // console.log("outer config name : ",configName)
                                // console.log("Inner ConfigName : ",res[k].ConfigName)
                                if(res[k].ConfigName == vm[i].connectionName){
                                    var provider = res[k].ProviderName
                                    console.log("Provider : ",provider);
                                    var kv_list = vm[i].cspViewVmDetail.KeyValueList
                                    var archi = ""
                                    for(var p in kv_list){
                                        if(kv_list[p].Key == "Architecture"){
                                         archi = kv_list[p].Value 
                                        }
                                    }

                                    if(status == "running"){
                                        badge += '<span class="badge badge-pill badge-success">RUNNING</span>'
                                    }else if(status == "suspended"){
                                        badge += '<span class="badge badge-pill badge-warning">SUSPEND</span>'
                                    }else if(status == "terminate"){
                                        badge += '<span class="badge badge-pill badge-dark">TERMINATED</span>'
                                    }else{
                                        badge += '<span class="badge badge-pill badge-dark">'+status+'</span>'
                                    }
                                    count++;
                                    if(count == 1){
                        
                                    }
                                    html += '<tr id="tr_id_'+count+'" >'
                                    +'<td class="text-center">'
                                    +'<div class="form-input">'
                                    +'<span class="input">'
                                    +'<input type="checkbox" item="'+mcis_name+'"    mcisid="'+mcis_id+'" class="chk2" id="chk2_'+count+'" value="'+vm[i].id+'|'+mcis_id+'"><i></i></span></div>'
                                    +'</td>'
                                    +'<td>'
                                    +badge
                                    +'</td>'
                                    +'<td><a href="#!" onclick="show_vm(\''+mcis_id+'\',\''+vm[i].id+'\',\''+vm[i].imageId+'\');">'+vm[i].name+'</a></td>'
                        
                                    +'<td>'+provider+'</td>'
                                    +'<td>'+vm[i].region.Region+'</td>'
                                    +'<td>'+vm[i].connectionName+'</td>'
                                    +'<td>'+archi+'</td>'
                                    +'<td>'+vm[i].publicIP+'</td>'
                                    +'<td>'+short_desc(vm[i].description)+'</td>'
                                    +'<td>'
                                    +'<button type="button" class="btn btn-icon dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">'
                                    +'<i class="fas fa-edit"></i>'
                                    +'<div class="dropdown-menu dropdown-menu-right" aria-labelledby="btnGroupDrop1">'
                                    +'<h6 class="dropdown-header text-center" style="background-color:#F2F4F4;;cursor:default;"><i class="fas fa-recycle"></i> LifeCycle</h6>'
                                        +'<a class="dropdown-item text-right" href="#!" onclick="life_cycle(\'vm\',\'resume\',\''+mcis_id+'\',\''+mcis_name+'\',\''+vm[i].id+'\',\''+vm[i].name+'\')">Resume</a>'
                                        +'<a class="dropdown-item text-right" href="#!" onclick="life_cycle(\'vm\',\'suspend\',\''+mcis_id+'\',\''+mcis_name+'\',\''+vm[i].id+'\',\''+vm[i].name+'\')">Suspend</a>'
                                        +'<a class="dropdown-item text-right" href="#!" onclick="life_cycle(\'vm\',\'reboot\',\''+mcis_id+'\',\''+mcis_name+'\',\''+vm[i].id+'\',\''+vm[i].name+'\')">Reboot</a>'
                                        +'<a class="dropdown-item text-right" href="#!" onclick="life_cycle(\'vm\',\'terminate\',\''+mcis_id+'\',\''+mcis_name+'\',\''+vm[i].id+'\',\''+vm[i].name+'\')">Terminate</a>'
                                    +'</div>'
                                    +'</button>'
                                    +'</td>'
                                    +'</tr>';
                                }
                                
                                
                                }
                                $("#table_2").empty();
                                $("#table_2").append(html);
                                $("#vm_detail").hide();
                                fnMove("table_2");

                        }

                    })
                    
                    }
            }
        })
    }else{
        $("#table_2").empty();
        $("#table_2").append("<td colspan='9'>Does not Exist</td>");
    }
            
    
 }
 
 
 
 function show_vm(mcis_id,vm_id,image_id){
    show_vmDetailList(mcis_id, vm_id);
    show_vmSpecInfo(mcis_id, vm_id);
    show_vmNetworkInfo(mcis_id, vm_id);
    show_vmSecurityGroupInfo(mcis_id, vm_id);
    show_vmSSHInfo(mcis_id, vm_id);
    show_images(image_id);
    $("#vm_detail").show();
 }


 function deleteHandler(cl,target,){
    var url = SpiderURL+"/connectionconfig"
 }

 function mcis_delete(){
    
    var cnt = 0;
    var mcis_id = "";
    var apiInfo = ApiInfo;
    $(".chk").each(function(){
        if($(this).is(":checked")){
            //alert("chk");
            cnt++;
            mcis_id = $(this).val();        
        }
        if(cnt < 1 ){
            alert("삭제할 대상을 선택해 주세요.");
            return;
        }

        if(cnt == 1){
           console.log("mcis_id ; ",mcis_id)
            var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id
            
            if(confirm("삭제하시겠습니까?")){
             axios.delete(url,{
                headers :{
                    'Content-type': 'application/json',
                    'Authorization': apiInfo,
                    }
             }).then(result=>{
                 var data = result.data
                 if(result.status == 200){
                     alert(data.message)
                     location.reload(true)
                 }
             })
            }
        }

        if(cnt >1){
            alert("한개씩만 삭제 가능합니다.")
            return;
        }

    })
 }

 function mcis_reg(){
    
    var cnt = 0;
    var mcis_id = "";
    $(".chk").each(function(){
        if($(this).is(":checked")){
            //alert("chk");
            cnt++;
            mcis_id = $(this).val();
            mcis_name = $(this).attr("item");

        }


    })
    if(cnt < 1 ){
        alert("등록할 대상을 선택해 주세요.");
        return;
    }

    if(cnt == 1){
       console.log("mcis_id ; ",mcis_id)
        var url = "/MCIS/reg/"+mcis_id+"/"+mcis_name
        
        if(confirm("등록하시겠습니까?")){
            location.href = url;
        }
    }

    if(cnt >1){
        alert("한개씩만 등록 가능합니다.")
        return;
    }
 }
 function vm_reg(){
    
    var cnt = 0;
    var mcis_id = "";
    var mcis_name = "";
    
    mcis_id = $("#mcis_id").val()
    mcis_name = $("#mcis_name").val()
    var url = "/MCIS/reg/"+mcis_id+"/"+mcis_name
    console.log("vm reg url : ",url)
    if(confirm("Add Server?")){
        location.href = url;
    }

 }

 function vm_delete(){
    
    var cnt = 0;
    var vm_id = "";
    var mcis_id ="";
    $(".chk2").each(function(){
        if($(this).is(":checked")){
            //alert("chk");
            cnt++;
            id = $(this).val(); 
            idArr = id.split ("|")  
            vm_id = idArr[0]
            mcis_id = idArr[1]    
        }
    })
    if(cnt < 1 ){
        alert("삭제할 대상을 선택해 주세요.");
        return;
    }

    if(cnt == 1){
       console.log("mcis_id ; ",vm_id)
        var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id+"/vm/"+vm_id
        
        if(confirm("삭제하시겠습니까?")){
         axios.delete(url,{
            headers :{
                'Content-type': 'application/json',
                'Authorization': apiInfo,
                }
         }).then(result=>{
             var data = result.data
             console.log(result);
             if(result.status == 200){
                 alert(data.message)
                 location.reload(true)
             }
         })
        }
    }

    if(cnt >1){
        alert("한개씩만 삭제 가능합니다.")
        return;
    }
 }

 function getProvider(connectionInfo){
     url = SpiderURL+"/connectionconfig"
     var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
         var data = result.data.connectionconfig
         

         for(var i in data){
             if(connetionInfo == data[i].ConfigName){}
         }
     })
 }

 function show_vmDetailList(mcis_id, vm_id){
     url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id+"/vm/"+vm_id
     var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
         var data = result.data
         console.log("show vmDetail List data : ",data)
         var html = ""
         $.ajax({
            url: SpiderURL+"/connectionconfig",
            async:false,
            type:'GET',
            beforeSend : function(xhr){
                xhr.setRequestHeader("Authorization", apiInfo);
                xhr.setRequestHeader("Content-type","application/json");
            },
            success : function(data2){
                res = data2.connectionconfig
                var provider = "";
                for(var k in res){
                    if(res[k].ConfigName == data.connectionName){
                        provider = res[k].ProviderName
                        console.log("Inner Provider : ",provider)
                    }
                }
                html += '<tr>'
                    +'<th scope="colgroup"rowspan="10" class="text-center">Infra - Server</th>'

                    +'<th scope="colgroup" class="text-right">Server ID</th>'
                    +'<td  colspan="1">'+data.id+'</td>'
                    
                    
                    +'<th scope="colgroup" class="text-right">Cloud Provider</th>'
                    +'<td colspan="1">'+provider+'</td>'
                    +'</tr>'


                    +'<tr>'
                    // +'<th scope="colgroup" class="text-right">CP VMID</th>'
                    // +'<td  colspan="1">'+data.id+'</td>'
                   
                    +'<th scope="colgroup" class="text-right">Region</th>'
                    +'<td  colspan="1" >'+data.region.Region+'</td>'
                    +'<th scope="colgroup" class="text-right">Zone</th>'
                    +'<td  colspan="1">'+data.region.Zone+'</td>'
                    +'</tr>'

                    
                    +'<tr>'
                    +'<th scope="colgroup" class="text-right">Public IP</th>'
                    +'<td  colspan="1">'+data.publicIP+'</td>'
                    
                    +'<th scope="colgroup" class="text-right">Public DNS</th>'
                    +'<td  colspan="1">'+data.publicDNS+'</td>'
                    +'</tr>'

                    +'<tr>'
                    +'<th scope="colgroup" class="text-right">Private IP</th>'
                    +'<td colspan="1">'+data.privateIP+'</td>'
                    
                    +'<th scope="colgroup" class="text-right">Private DNS</th>'
                    +'<td colspan="1">'+data.privateDNS+'</td>'
                    +'</tr>'

                    +'<tr>'
                    +'<th scope="colgroup" class="text-right">Server Status</th>'
                    +'<td colspan="3">'+data.status+'</td>'
                    +'</tr>';
                  
                $("#vm").empty();
                $("#vm").append(html);
                fnMove("vm_detail");

            }

        })
       
            
         
     })

 }



function show_vmSpecInfo(mcis_id, vm_id){
    var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id+"/vm/"+vm_id
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        var data = result.data
        console.log("show vmSpecInfo Data : ",data)
        var html = ""
        var url2 = CommonURL+"/ns/"+NAMESPACE+"/resources/spec"
        var spec_id = data.specId
        $.ajax({
           url: url2,
           async:false,
           type:'GET',
           beforeSend : function(xhr){
            xhr.setRequestHeader("Authorization", apiInfo);
            xhr.setRequestHeader("Content-type","application/json");
        },
           success : function(result){
               var res = result.spec
              console.log("spec data from tumble : ",res)
               for(var k in res){
                   if(res[k].id == spec_id){
                    html += '<tr>'
                          
                           +'<th scope="colgroup" rowspan="5"class="text-right"><i class="fas fa-server"></i>Server Spec</th>'
                           +'<th scope="colgroup" class="text-right">vCPU</th>'
                           +'<td colspan="1">'+res[k].num_vCPU+' vcpu</td>'
                         
                           +'<th scope="colgroup" class="text-right">Memory(Ghz)</th>'
                           +'<td  colspan="1">'+res[k].mem_GiB+' GiB</td>'
                           +'</tr>'

                           +'<tr>'
                           +'<th scope="colgroup" class="text-right">Disk (GB)</th>'
                           +'<td colspan="1">'+res[k].storage_GiB+' GiB</th>'
                          
                           +'<th scope="colgroup" class="text-right">Cost($) / Hour </th>'
                           +'<td colspan="1">'+res[k].cost_per_hour+'</td>'
                           +'</tr>'

                           +'<tr>'
                           +'<th scope="colgroup">OsType</th>'
                           +'<td  colspan="3">'+res[k].os_type+'</td>'
                           +'</tr>'
                   }
               } 
               $("#vm_spec").empty();
               $("#vm_spec").append(html);

           }

       })
      
           
        
    })

}

function show_vmNetworkInfo(mcis_id, vm_id){
    var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id+"/vm/"+vm_id
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        var data = result.data
        var html = ""
        var url2 = CommonURL+"/ns/"+NAMESPACE+"/resources/vNet"
        var spec_id = data.vNetId
        $.ajax({
           url: url2,
           async:false,
           type:'GET',
           beforeSend : function(xhr){
            xhr.setRequestHeader("Authorization", apiInfo);
            xhr.setRequestHeader("Content-type","application/json");
        },
           success : function(result){
               var res = result.vNet
              console.log("Network Info : ",result)
               for(var k in res){
                   if(res[k].id == spec_id){
                    var subnetInfoList = res[k].subnetInfoList
                    var subnetArr = new Array()
                    var str = ""
                    if(subnetInfoList){
                        for(var o in subnetInfoList){
                             subnetArr.push(subnetInfoList[o].IPv4_CIDR)
                        }
                        str = subnetArr.join(",")
                    }
                    console.log("Subnet str : ",str)
                    html += '<tr>'
                           +'<th scope="colgroup" rowspan="5" class="text-right"><i class="fas fa-network-wired"></i>Network</th>'
                           +'<th scope="colgroup" class="text-right">Network Name</th>'
                           +'<td  colspan="1">'+res[k].cspVNetName+'</td>'
                           +'<th scope="colgroup" class="text-right">Network ID</th>'
                           +'<td colspan="1">'+res[k].cspVNetId+'</td>'
                          
                           +'</tr>'
                           +'<tr>'
                           +'<th scope="colgroup" class="text-right">Cidr Block</th>'
                           +'<td colspan="3">'+res[k].cidrBlock+'</th>'
                           +'</tr>'
                           +'<tr>'
                           +'<th scope="colgroup" class="text-right">Subnet</th>'
                           +'<td colspan="3">'+str+'</th>'
                           +'</tr>'
                        //    +'<tr>'
                        //    +'<th scope="colgroup">Interface</th>'
                        //    +'<td colspan="3">'+res[k].cidrBlock+'</th>'
                        //    +'</tr>'
                          
                   }
               } 
               console.log("vnetwork html : ",html)
               $("#vm_vnetwork").empty();
               $("#vm_vnetwork").append(html);

           }

       })
      
           
        
    })

}

function show_vmSecurityGroupInfo(mcis_id, vm_id){
    var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id+"/vm/"+vm_id
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        var data = result.data
        var html = ""
        // var url2 = "/ns/"+NAMESPACE+"/resources/securityGroup"
        var spec_id = data.securityGroupIds
        var cnt = spec_id.length
        html += '<tr>'
             +'<th scope="colgroup" colspan="'+cnt+'" "class="text-right"><i class="fas fa-shield-alt"></i>SecurityGroup</th>'
             +'<th scope="colgroup" colspan="'+cnt+'" class="text-right">SecurityGroup ID</th>'
        for(var i in spec_id){
            if( i == 0){
                html +='<td colspan="3">'+spec_id[i]+'</td></tr>'
            }else{
                html +='<tr><td colspan="3">'+spec_id[i]+'</td></tr>'
            }
        }
        

        $("#vm_sg").empty();
        $("#vm_sg").append(html);

                
        
    })

}


function show_vmSSHInfo(mcis_id, vm_id){
    var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id+"/vm/"+vm_id
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        var data = result.data
        var html = ""
        var url2 = CommonURL+"/ns/"+NAMESPACE+"/resources/sshKey"
        var spec_id = data.sshKeyId
        $.ajax({
           url: url2,
           async:false,
           type:'GET',
           beforeSend : function(xhr){
            xhr.setRequestHeader("Authorization", apiInfo);
            xhr.setRequestHeader("Content-type","application/json");
        },
           success : function(result){
               var res = result.sshKey
              
               for(var k in res){
                   if(res[k].id == spec_id){
                    html += '<tr>'
                           +'<th scope="colgroup" rowspan="3" class="text-right"><i class="fas fa-key"></i>Access(SSH Key)</th>'
                           +'<th scope="colgroup" class="text-right">Key Name</th>'
                           +'<td  colspan="1">'+res[k].cspSshKeyName+'</td>'
                           +'<th scope="colgroup" class="text-right">SSH Key ID</th>'
                           +'<td colspan="1">'+res[k].id+'</td>'
                          
                           
                           +'</tr>'
                           +'<tr>'
                           +'<th scope="colgroup" class="text-right">Description</th>'
                           +'<td colspan="3">'+res[k].description+'</th>'
                           +'</tr>'
                          
                   }
               } 
               $("#sshKey").empty();
               $("#sshKey").append(html);

           }

       })
      
           
        
    })

}

function show_images(image_id){
    var url = CommonURL+"/ns/"+NAMESPACE+"/resources/image/"+image_id
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        var data = result.data
        console.log("Image Data : ",data);
        var html = ""
            
        html += '<tr>'
                +'<th scope="colgroup" rowspan="5" class="text-right"><i class="fas fa-compact-disc"></i>Image</th>'
                +'<th scope="colgroup" class="text-right">Image Name</th>'
                +'<td  colspan="1">'+data.name+'</td>'
                +'<th scope="colgroup" class="text-right">Image ID</th>'
                +'<td colspan="1">'+data.id+'</td>'
                
                +'</tr>'
                +'<tr>'
                +'<th scope="colgroup" class="text-right">Guest OS</th>'
                +'<td colspan="1">'+data.guestOS+'</th>'
                
                +'<th scope="colgroup" class="text-right">Description</th>'
                +'<td colspan="1">'+data.description+'</th>'
                +'</tr>'
            
                          
             
             
               $("#vm_image").empty();
               $("#vm_image").append(html);

           })

    

}