// MCIS Control
function life_cycle(tag,type,mcis_id,mcis_name,vm_id,vm_name){
    var url = ""
    var nameSpace = NAMESPACE;
    var message = ""
    console.log("Start LifeCycle method!!!")
    
    if(tag == "mcis"){
        url ="/ns/"+nameSpace+"/mcis/"+mcis_id+"?action="+type
        message = mcis_name+" "+type+ " complete!."
    }else{
        url ="/ns/"+nameSpace+"/mcis/"+mcis_id+"/vm/"+vm_id+"?action="+type
        message = vm_name+" "+type+ " complete!."
    }

    axios.get(url).then(result=>{
        var status = result.status
        
        console.log("life cycle result : ",result)
        var data = result.data
        console.log("result Message : ",data.message)
        if(status == 200){
            
            alert(message);
            location.reload();
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
 function show_mcis(url){
    var html = "";
    axios.get(url).then(result=>{
        var data = result.data;
         console.log("showmcis Data : ",data)
         var html = "";
         var mcis = data.mcis;
         var len = mcis.length;
         var count = 0;
         
          
         for(var i in mcis){
             var sta = mcis[i].status;
             var badge = "";
             var status = sta.toLowerCase()
             var vms = mcis[i].vm
             var vm_len = vms.length 

             if(status == "running"){
                badge += '<span class="badge badge-pill badge-success">RUNNING</span>'
             }
             count++;
             if(count == 1){

             }
             html += '<tr id="tr_id_'+count+'" onclick="show_card(\' '+mcis[i].id+'\' )">'
              +'<td class="text-center">'
              +'<div class="form-input">'
              +'<span class="input">'
              +'<input type="checkbox" class="chk" id="chk_'+count+'" value="'+mcis[i].id+'"><i></i></span></div>'
              +'</td>'
              +'<td><a href="">'+mcis[i].name+'</a></td>'
              +'<td>12:32:30</td>'
              +'<td>'+vm_len+'</td>'
              +'<td>'+short_desc(mcis[i].description)+'</td>'
              +'<td>'
              +badge
              +'</td>'
              +'<td>'
              +'<button type="button" class="btn btn-icon dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">'
              +'<i class="fas fa-edit"></i>'
              +'<div class="dropdown-menu dropdown-menu-right" aria-labelledby="btnGroupDrop1">'
                  +'<a class="dropdown-item" href="#!" onclick="life_cycle(\'mcis\',\'resume\',\''+mcis[i].id+'\',\''+mcis[i].name+'\')">Resume</a>'
                  +'<a class="dropdown-item" href="#!" onclick="life_cycle(\'mcis\',\'suspend\',\''+mcis[i].id+'\',\''+mcis[i].name+'\')">Suspend</a>'
                  +'<a class="dropdown-item" href="#!" onclick="life_cycle(\'mcis\',\'reboot\',\''+mcis[i].id+'\',\''+mcis[i].name+'\')">Reboot</a>'
                  +'<a class="dropdown-item" href="#!" onclick="life_cycle(\'mcis\',\'terminate\',\''+mcis[i].id+'\',\''+mcis[i].name+'\')">Terminate</a>'
              +'</div>'
              +'</button>'
             +'</td>'
             +'</tr>';
        }
        
        $("#table_1").empty();
        $("#table_1").append(html);
        show_card(mcis[0].id);
        show_vmList(mcis[0].id);
    });
 }
 function show_vmList(mcis_id){
   
    var url = "/ns/"+NAMESPACE+"/mcis/"+mcis_id;
    $.ajax({
        type:'GET',
        url:url,
       // async:false,
        success:function(data){
            var vm = data.vm
            var mcis_name = data.name 
            var html = "";
            
            for(var i in vm){
                var sta = vm[i].status;
                console.log(" i value is : ",i)
               
                var status = sta.toLowerCase()
                var configName = vm[i].config_name
                console.log("outer vm configName : ",configName)
                var count = 0;
                $.ajax({
                    url:"/connectionconfig",
                    async:false,
                    type:'GET',
                    success : function(res){
                        var badge = "";
                        for(var k in res){
                            console.log(" i value is : ",i)
                            console.log("outer config name : ",configName)
                            console.log("Inner ConfigName : ",res[k].ConfigName)
                            if(res[k].ConfigName == vm[i].config_name){
                                var provider = res[k].ProviderName
                                console.log("Inner Provider : ",provider)
                                if(status == "running"){
                                    badge += '<span class="badge badge-pill badge-success">RUNNING</span>'
                                 }
                                 count++;
                                 if(count == 1){
                    
                                 }
                                 html += '<tr id="tr_id_'+count+'" onclick="showMonitoring(\''+mcis_id+'\',\''+vm[i].id+'\');">'
                                  +'<td class="text-center">'
                                  +'<div class="form-input">'
                                  +'<span class="input">'
                                  +'<input type="checkbox" class="chk2" id="chk2_'+count+'" value="'+vm[i].id+'"><i></i></span></div>'
                                  +'</td>'
                                  +'<td><a href="">'+vm[i].name+'</a></td>'
                                  +'<td>12:32:30</td>'
                                  +'<td>'+provider+'</td>'
                                  +'<td>'+vm[i].region.Region+'</td>'
                                  +'<td>'+short_desc(vm[i].description)+'</td>'
                                  +'<td>'
                                  +badge
                                  +'</td>'
                                  +'<td>'
                                  +"3/8"
                                  +'</td>'
                                  +'<td>'
                                  +'<button type="button" class="btn btn-icon dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">'
                                  +'<i class="fas fa-edit"></i>'
                                  +'<div class="dropdown-menu dropdown-menu-right" aria-labelledby="btnGroupDrop1">'
                                      +'<a class="dropdown-item" href="#!" onclick="life_cycle(\'vm\',\'resume\',\''+mcis_id+'\',\''+mcis_name+'\',\''+vm[i].id+'\',\''+vm[i].name+'\')">Resume</a>'
                                      +'<a class="dropdown-item" href="#!" onclick="life_cycle(\'vm\',\'suspend\',\''+mcis_id+'\',\''+mcis_name+'\',\''+vm[i].id+'\',\''+vm[i].name+'\')">Suspend</a>'
                                      +'<a class="dropdown-item" href="#!" onclick="life_cycle(\'vm\',\'reboot\',\''+mcis_id+'\',\''+mcis_name+'\',\''+vm[i].id+'\',\''+vm[i].name+'\')">Reboot</a>'
                                      +'<a class="dropdown-item" href="#!" onclick="life_cycle(\'vm\',\'terminate\',\''+mcis_id+'\',\''+mcis_name+'\',\''+vm[i].id+'\',\''+vm[i].name+'\')">Terminate</a>'
                                  +'</div>'
                                  +'</button>'
                                 +'</td>'
                                 +'</tr>';
                            }
                            
                            
                            }
                            $("#table_2").empty();
                            $("#table_2").append(html);

                    }

                })
                
                }
        }
    })
            
    
 }
 
 function show_card(mcis_id){
     var url = "/ns/"+NAMESPACE+"/mcis/"+mcis_id;
     var html = "";
    axios.get(url).then(result=>{
        var data = result.data
        console.log("show card data : ",result)
        var vm_cnt = data.vm
        vm_cnt = vm_cnt.length
        
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

        show_vmList(mcis_id)
       
    })
 }
 function show_vm(mcis_id,vm_id){
     var url = "/ns/"+NAMESPACE+"/mcis/"+mcis_id+"/vm/"+vm_id;
     var html = "";
    axios.get(url).then(result=>{
        var data = result.data
        console.log("show card result : ",result)
                   
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
                    +"vm_cnt"
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
       
    })
 }
 function sel_table(targetNo,mcid){
     var $target = $("#card_"+targetNo+"");
     var html = "";
     url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcid
     axios.get(url).then(result=>{
         var data = result.data.vm
         for(var i in data){

         }
     })
     html += '<tr>'
             +'<td class="text-center">'
             +'<div class="form-input">'
             +'<span class="input">'
             +'<input type="checkbox" id=""><i></i>'
             +'</span></div>'
             +'</td>'
             +'<td>1</td>'
             +'<td><a href="">Baristar1</a></td>'
             +'<td>aws driver 1aws driver ver0.1</td>'
             +'<td>aws key 1</td>'
             +'<td>ap-northest-1</td>'
             +'<td>'
             +'<div class="custom-control custom-switch">'
             +'<input type="checkbox" class="custom-control-input" id="customSwitch1">'
             +'<label class="custom-control-label" for="customSwitch1"></label></div>'
             +'</td>'
             +'<td>'
             +'<span class="badge badge-pill badge-warning">stop</span>'
             +'</td>'
             +'<td>2019-05-05</td>'
             +'</tr>';
             
    $target.empty();         
    $target.append(html);

 }

 function deleteHandler(cl,target,){

 }

 function getProvider(connectionInfo){
     url ="/connectionconfig"
     axios.get(url).then(result=>{
         var data = result.data

         for(var i in data){
             if(connetionInfo == data[i].ConfigName){}
         }
     })
 }

 function mappingMetric(obj){
    var name = obj.name
    var columnArr = obj.columns
    var valuesArr = obj.values
    var valuesCnt = valuesArr.length
    var objArr = new Array();
    for(var i in  valuesArr){
       var newObject = {}
        for(var k in valuesArr[i]){
            var key = columnArr[k]
            var value = valuesArr[i][k]
            newObject[key] = value
        }
        objArr.push(newObject)
    }
    console.log("Mapping Metric : ",objArr);
    return objArr
}

