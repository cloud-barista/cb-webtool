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
             
             html +='<tr onclick="click_view(\''+mcis[i].id+'\',\''+i+'\');" id="server_info_tr_'+i+'">'
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

        // $(".dashboard .status_list tbody tr").each(function(){
        //     var $td_list = $(this),
        //             $status = $(".server_status"),
        //             $detail = $(".server_info");
          
        //     $td_list.off("click").click(function(){
        //           $td_list.addClass("on");
        //           $td_list.siblings().removeClass("on");
        //           $status.addClass("view");
        //           $status.siblings().removeClass("on");
        //         $(".dashboard.register_cont").removeClass("active");
        //          $td_list.off("click").click(function(){
        //               if( $(this).hasClass("on") ) {
        //                   console.log("1번에 걸린다.")
        //                   $td_list.removeClass("on");
        //                   $status.removeClass("view");
        //                   $detail.removeClass("active");
        //           } else {
        //             console.log("아니다 2번에 걸린다.")
        //                   $td_list.addClass("on");
        //                   $td_list.siblings().removeClass("on");
        //                   $status.addClass("view");
        //                   $status.siblings().removeClass("view");
        //                 $(".dashboard.register_cont").removeClass("active");
        //           }
        //           });
        //       });
        //   });
          
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
 function click_view(id,index){
     console.log("click view mcis id :",id)
    
    show_mcis(id,index);

    
     
     

    
    
   
    
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
                       vm_badge += '<li class="sel_cr bgbox_b"><a href="javascript:void(0);"><span class="txt">'+vms[o].name+'</span></a></li>';
                     
                    }else if(vm_status == "include" ){
                        vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);"><span class="txt">'+vms[o].name+'</span></a></li>';
                    }else if(vm_status == "Suspended"){
                        vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);"><span class="txt">'+vms[o].name+'</span></a></li>';
                      
                    }else if(vm_status == "Terminated"){
                        vm_badge += '<li class="sel_cr bgbox_r"><a href="javascript:void(0);"><span class="txt">'+vms[o].name+'</span></a></li>';
                      
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
      
    });
 }
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
 
 function show_card(mcis_id){
    $("#vm_detail").hide();
     var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id;
     var html = "";
     var apiInfo = ApiInfo
     axios.get(url,{
         headers:{
             'Authorization': apiInfo
         }
     }).then(result=>{
        var data = result.data
        console.log("show card data : ",result)
        var vm_cnt = data.vm
        var mcis_name = data.name
        if(vm_cnt){
            vm_cnt = vm_cnt.length
        }else{
            vm_cnt = 0
        }
        
        
            html += '<div class="col-xl-12 col-lg-12">'
                    +'<div class="card card-stats mb-12 mb-xl-0">'
                    +'<div class="card-body">'
                    +'<div class="row">'
                    +'<div class="col">'
                    +'<h5 class="card-title text-uppercase text-muted mb-0">'+data.name+'</h5>'
                    +'<span class="h2 font-weight-bold mb-0">350,897</span>'
                    +'</div>'
                    +'<div class="col-auto">'
                    +'<div class="icon icon-shape bg-danger text-white rounded-circle shadow">'
                    //+'<i class="fas fa-chart-bar"></i>'
                    +vm_cnt
                    +'</div>'
                    +'</div>'
                    +'</div>'
                    +'<p class="mt-3 mb-0 text-muted text-sm">'
                    +'<span class="text-success mr-2"><i class="fa fa-arrow-up"></i> 3.48%</span>'
                    +'<span class="text-nowrap">Since last month</span>'
                    +'</p>'
                    +'</div>'
                    +'</div>'
                    +'</div>';
        
        $("#card").empty()
        $("#card").append(html)
        if(vm_cnt > 0){
            show_vmList(mcis_id)
        }else{
            show_vmList("")
        }
       
        $("#mcis_id").val(mcis_id)
        $("#mcis_name").val(mcis_name)
       
    })
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
//  function sel_table(targetNo,mcid){
//      var $target = $("#card_"+targetNo+"");
//      var html = "";
//      url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcid
//      var apiInfo = ApiInfo
    // axios.get(url,{
    //     headers:{
    //         'Authorization': apiInfo
    //     }
    // })then(result=>{
//          var data = result.data.vm
//          for(var i in data){

//          }
//      })
//      html += '<tr>'
//              +'<td class="text-center">'
//              +'<div class="form-input">'
//              +'<span class="input">'
//              +'<input type="checkbox" id=""><i></i>'
//              +'</span></div>'
//              +'</td>'
//              +'<td>1</td>'
//              +'<td><a href="">Baristar1</a></td>'
//              +'<td>aws driver 1aws driver ver0.1</td>'
//              +'<td>aws key 1</td>'
//              +'<td>ap-northest-1</td>'
//              +'<td>'
//              +'<div class="custom-control custom-switch">'
//              +'<input type="checkbox" class="custom-control-input" id="customSwitch1">'
//              +'<label class="custom-control-label" for="customSwitch1"></label></div>'
//              +'</td>'
//              +'<td>'
//              +'<span class="badge badge-pill badge-warning">stop</span>'
//              +'</td>'
//              +'<td>2019-05-05</td>'
//              +'</tr>';
             
//     $target.empty();         
//     $target.append(html);

//  }

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

//  function show_vmDetailInfo(mcis_id, vm_id){
//     var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcis_id+"/vm/"+vm_id
//     var apiInfo = ApiInfo
    // axios.get(url,{
    //     headers:{
    //         'Authorization': apiInfo
    //     }
    // })then(result=>{
//         var data = result.data
//         console.log("show vmDetailInfo data : ",data)
//         var html = ""
//         $.ajax({
//            url:SpiderURL+"/connectionconfig",
//            async:false,
//            type:'GET',
//            success : function(data2){
//             res = data2.connectionconfig
//                var provider = "";
//                for(var k in res){
//                    if(res[k].ConfigName == data.connectionName){
//                        provider = res[k].ProviderName
//                        console.log("Inner Provider : ",provider)
//                    }
//                }
//                html += '<tr>'
//                     +'<th scope="colgroup"rowspan="8">Infra - Server</th>'
//                     +'<th scope="colgroup">Cloud Provider</th>'
//                     +'<td colspan="3">'+provider+'</td>'
//                     +'</tr>'
//                     +'<tr>'

//                     +'<th scope="colgroup">Server ID</th>'
//                     +'<td  colspan="3">'+data.id+'</td>'
//                     +'</tr>'

//                     +'<th scope="colgroup">CP VMID</th>'
//                     +'<td  colspan="3">'+data.id+'</td>'
//                     +'</tr>'

//                     +'<tr>'
//                     +'<th scope="colgroup">Region</th>'
//                     +'<td  colspan="3">'+data.region.Region+'</td>'
//                     +'</tr>'

                    
//                     +'<tr>'
//                     +'<th scope="colgroup">Public IP</th>'
//                     +'<td  colspan="3">'+data.publicIP+'</td>'
//                     +'</tr>'

//                     +'<tr>'
//                     +'<th scope="colgroup">Public DNS</th>'
//                     +'<td  colspan="3">'+data.publicDNS+'</td>'
//                     +'</tr>'

//                     +'<tr>'
//                     +'<th scope="colgroup">Private IP</th>'
//                     +'<td colspan="3">'+data.privateIP+'</td>'
//                     +'</tr>';
//                     +'<tr>'
//                     +'<th scope="colgroup">Private DNS</th>'
//                     +'<td colspan="3">'+data.privateDNS+'</td>'
//                     +'</tr>';
//                     +'<tr>'
//                     +'<th scope="colgroup">Server Status</th>'
//                     +'<td colspan="3">'+data.status+'</td>'
//                     +'</tr>';
               
//                    +'</tbody>'
//                    +'<tbody>'
//                    +'<tr>'
//                    +'<th scope="colgroup" rowspan="3">VM Meta</th>'
//                    +'<th scope="colgroup">VM ID</th>'
//                    +'<td colspan="3">'+data.cspViewVmDetail.Id+'</td>'
//                    +'</tr>'
//                    +'<tr>'
//                    +'<th scope="colgroup">VM NAME</th>'
//                    +'<td  colspan="3">'+data.cspViewVmDetail.Name+'</td>'
//                    +'</tr>'
                   

                 
//                $("#vm").empty();
//                $("#vm").append(html);
//                fnMove("vm_detail");


//            }

//        })
      
           
        
//     })

// }

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