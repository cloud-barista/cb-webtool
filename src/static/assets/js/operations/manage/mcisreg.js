function getConnectionInfo(provider, cnt){
    var url = SpiderURL+"/connectionconfig";
    console.log("provider : ",provider)
    //var provider = $("#provider option:selected").val();
    var html = "";
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        var data = result.data.connectionconfig
        console.log("connection data : ",data);
        var count = 0; 
        var configName = "";
        var confArr = new Array();
        for(var i in data){
            if(provider == data[i].ProviderName){ 
                count++;
                html += '<option value="'+data[i].ConfigName+'" item="'+data[i].ProviderName+'">'+data[i].ConfigName+'</option>';
                configName = data[i].ConfigName
                confArr.push(data[i].ConfigName)
                
            }
        }
        if(count == 0){
            alert("해당 Provider에 등록된 Connection 정보가 없습니다.")
                html +='<option selected>Select Configname</option>';
        }
        if(confArr.length > 1){
            configName = confArr[0];
        }
        $("#configName_"+cnt+"").empty();
        $("#configName_"+cnt+"").append(html);
        getImageInfo(configName,cnt);
        getSecurityInfo(configName,cnt);
        getSSHKeyInfo(configName,cnt);
      //  getPublicIPInfo(configName,cnt);
        getVnetInfo(configName,cnt);

        
    })
}

function changeConnectionInfo(configName,cnt){
    getImageInfo(configName,cnt);
    getSecurityInfo(configName,cnt);
    getSSHKeyInfo(configName,cnt);
    //getPublicIPInfo(configName,cnt);
    getVnetInfo(configName,cnt);
}

function getImageInfo(configName,cnt){
   console.log("1 : ",configName);
    var configName = configName;
    if(!configName){
        configName = $("#configName_"+cnt+" option:selected").val();
    }
    console.log("2 : ",configName);

    var url = CommonURL+"/ns/"+NAMESPACE+"/resources/image";
    var html = "";
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        console.log("Image Info : ",result.data)
        data = result.data.image
        if(!data){
            alert("등록된 이미지 정보가 없습니다.")
            location.href = "/Image/reg"
            return;
        }
        for(var i in data){
            if(data[i].connectionName == configName){
                html += '<option value="'+data[i].id+'" >'+data[i].name+'('+data[i].id+')</option>'; 
            }
        }
        $("#image_"+cnt+"").empty();
        $("#image_"+cnt+"").append(html);
        
    })
}

function getSecurityInfo(configName,cnt){
    var configName = configName;
    if(!configName){
        configName = $("#configName_"+cnt+" option:selected").val();
    }
    var url = CommonURL+"/ns/"+NAMESPACE+"/resources/securityGroup";
    var html = "";
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        data = result.data.securityGroup
        for(var i in data){
            if(data[i].connectionName == configName){
                html += '<option value="'+data[i].id+'" >'+data[i].cspSecurityGroupName+'('+data[i].id+')</option>'; 
            }
        }
      
        $("#sg_"+cnt+"").empty();
        $("#sg_"+cnt+"").append(html);
        
    })
}
function getSSHKeyInfo(configName,cnt){
    var configName = configName;
    if(!configName){
        configName = $("#configName_"+cnt+" option:selected").val();
    }
    var url = CommonURL+"/ns/"+NAMESPACE+"/resources/sshKey";
    var html = "";
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        console.log("sshKeyInfo result :",result)
        data = result.data.sshKey
        for(var i in data){
            if(data[i].connectionName == configName){
                html += '<option value="'+data[i].id+'" >'+data[i].cspSshKeyName+'('+data[i].id+')</option>'; 
            }
        }
        $("#sshKey_"+cnt+"").empty();
        $("#sshKey_"+cnt+"").append(html);
        
    })
}

function getVnetInfo(configName,cnt){
    var configName = configName;
    if(!configName){
        configName = $("#configName_"+cnt+" option:selected").val();
    }
    var url = CommonURL+"/ns/"+NAMESPACE+"/resources/vNet";
    var html = "";
    var html2 = "";
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    }).then(result=>{
        data = result.data.vNet
        console.log("vNetwork Info : ",result)
        for(var i in data){
            if(data[i].connectionName == configName){
                html += '<option value="'+data[i].id+'" selected>'+data[i].cspVNetName+'('+data[i].id+')</option>'; 
                var subnetInfoList = data[i].subnetInfoList
                for(var k in subnetInfoList){
                    html2 += '<option value="'+subnetInfoList[k].IId.NameId+'" >'+subnetInfoList[k].IPv4_CIDR+'</option>'; 
                }
            }
        }
        $("#vnet_"+cnt+"").empty();
        $("#vnet_"+cnt+"").append(html);
        $("#subnet_"+cnt+"").empty();
        $("#subnet_"+cnt+"").append(html2);
        
    })
}


$(document).ready(function(){
vm_cnt = 1;
$("#server_cnt").val(vm_cnt);
show_vm(vm_cnt)

$("[class^='vm_add'").click(function(){
vm_cnt++;
show_vm(vm_cnt);
$("#server_cnt").val(vm_cnt);
$("#provider_"+vm_cnt).focus();
btn_click_cnt++;
});


})

function show_vm(vm_cnt){
//vm_cnt++;
console.log(vm_cnt);
var arrCnt = vm_cnt-1;
var html = "";
html += '<form id="form_'+vm_cnt+'">'
    +'<div class="card">'
    +'<div class="card-header">'
    +'<div class="d-flex justify-content-between align-items-center">'
    +'<div>'
    +'<strong>Server #'+vm_cnt+'Config</strong>'
    +'</div>'
    +'<div>'
    +'<!-- 좌우측 정렬되는 버튼이 필요할때 여기에 넣어주세요 -->'
    +'</div></div></div>'
    +'<div class="card-body">'
    +'<table class="table table-bordered table-horizontal mb-0">'
    +'<colgroup><col style="width: 20%"><col></colgroup>'
    +'<tbody>'
    +'<tr>'

    +'<th>Cloud Provider<strong class="text-danger">*</strong></th>'
    +'<td>'
    +'<select class="form-control form-control-sm" name="provider" id="provider_'+vm_cnt+'" onchange="getConnectionInfo(this.value,'+vm_cnt+');"  required>'
    // +'<option selected>Select Cloud Provider</option>'
    // +'<option value="AWS">AWS</option>'
    // +'<option value="AZURE">AZURE</option>'
    // +'<option value="ALIBABA">Alibaba</option>'
    // +'<option value="GCP">GCP</option>'
    // +'<option value="CLOUDIT">Cloudit</option>'
    // +'<option value="OPENSTACK">Openstack</option>'
    +'</select>'
    +'</td></tr>'

    +'<tr><th>Cloud Connection<strong class="text-danger">*</strong></th>'
    +'<td>'
    +'<select class="form-control form-control-sm" name="connectionName" id="configName_'+vm_cnt+'"  onchange="changeConnectionInfo(this.value,'+vm_cnt+')"  required>'
    +'<option selected>Select Cloud Connection</option>'
    +'</select></td></tr>'
    +'<tr><th>Server Name</th><td>'
    +'<input class="vm_name form-control form-control-sm" type="text" name="name" onkeyup="view_vm();" placeholder="input Name" id="vmName_'+vm_cnt+'"  required></td></tr>'
    
    +'<tr><th>Server Spec<strong class="text-danger">*</strong></th><td>'
    +'<div class="form-row">'
    +'<div class="col">'
    +'<input type="hidden" id="vmspec_'+vm_cnt+'" name="specId" class="form-control form-control-sm" placeholder="Select Server Spec" readonly>'
    
    +'<input type="text" id="vmspecName_'+vm_cnt+'" name="spec_gb" class="form-control form-control-sm" placeholder="Select Spec" readonly></div>'
    +'<div class="col">'
    +'<button type="button" class="btn btn-dark btn-sm" onclick="show_pop('+vm_cnt+');">Search Server Spec</button>'
    +'</div></div></td></tr>'

    +'<tr><th class="th-right">vCPU</th><td>'
    +'<input class="form-control form-control-sm" type="text" placeholder="CPU" id="vcpu_'+vm_cnt+'" name="vcpu" readonly></td></tr>'
    +'<tr><th class="th-right">Memory(Ghz)</th><td>'
    +'<input class="form-control form-control-sm" type="text" placeholder="Memory" id="mem_'+vm_cnt+'" name="mem" readonly></td></tr>'
    +'<tr><th class="th-right">Storage(GB)</th><td>'
    +'<input class="form-control form-control-sm" type="text" placeholder="Storage" id="storage_'+vm_cnt+'" name="storage" readonly></td></tr>'
    +'<tr><th class="th-right">OS Type</th><td>'
    +'<input class="form-control form-control-sm" type="text" placeholder="OS Type" id="osType_'+vm_cnt+'" name="osType" readonly></td></tr>'
    +'<tr><th class="th-right">Cost($) / Hour</th><td>'
    +'<input class="form-control form-control-sm" type="text" placeholder="Cost Per Hour" id="cPh'+vm_cnt+'" name="costPerHour" readonly></td></tr>'
    +'<tr><th>Select Image</th><td>'
    +'<select class="form-control form-control-sm" name="imageId" id="image_'+vm_cnt+'"   required>'
    +'<option selected>Select Image</option>'
    
    +'</select></td></tr>'
    +'<tr><th>InstallMonAgent<strong class="text-danger">*</strong></th>'
    +'<td>'
    +'<select class="form-control form-control-sm" name="installMonAgent" id="i_mon_agent_'+vm_cnt+'"  required>'
    +'<option value="no">NO</option>'
    +'<option value="yes">YES</option>'
    +'</select></td></tr>'
    +'<tr><th>Network<strong class="text-danger">*</strong></th>'
    +'<td>'
    +'<select class="form-control form-control-sm" name="vNetId" id="vnet_'+vm_cnt+'"  required >'
    +'<option selected>Select Network</option>'
    +'</select></td></tr>'

    +'<tr><th>Subnet<strong class="text-danger">*</strong></th>'
    +'<td>'
    +'<select class="form-control form-control-sm" name="subnetId" id="subnet_'+vm_cnt+'"  >'
    +'<option selected>Select SubNet</option>'
    +'</select></td></tr>'
    // +'<tr><th>PublicIP</th>'
    // +'<td>'
    // +'<select class="form-control form-control-sm" name="public_ip_id" id="publicIp_'+vm_cnt+'"  >'
    // +'<option selected>Select PublicIP</option>'
    // +'</select></td></tr>'
    
    // +'<tr><th>PublicIP</th><td>'
    // +'<div class="form-row">'
    // +'<div class="col">'
    // +'<input type="text" id="publicIP_'+vm_cnt+'" name="publicIP[]" class="form-control form-control-sm" placeholder="Public IP" readonly></div>'
    // +'<div class="col">'
    // +'<button type="button" class="btn btn-dark btn-sm" onclick="show_pop2('+vm_cnt+');">Search PublicIP</button>'
    // +'</div></div></td></tr>'
    // +'</td></tr>'
    +'<tr><th>SecurityGroup<strong class="text-danger">*</strong></th>'
    +'<td>'
    +'<select class="form-control form-control-sm" multiple name="securityGroupIds" id="sg_'+vm_cnt+'"  required>'
    +'<option>Select SecurityGroup</option>'
    +'</select></td></tr>'
   

    +'<tr><th>Access (SSH Key)<strong class="text-danger">*</strong></th>'
    +'<td>'
    +'<select class="form-control form-control-sm" name="sshKeyId" id="sshKey_'+vm_cnt+'"  required >'
    +'<option selected>Select SSH KEY</option>'
    +'</select></td></tr>'

    +'<tr><th>Access ID</th><td>'
    +'<input class="form-control form-control-sm" type="text" placeholder="input Access ID" id="vm_access_'+vm_cnt+'" name="vmUserAccount" ></td></tr>'

    +'<tr><th>Access Password</th><td>'
    +'<input class="form-control form-control-sm" type="text" placeholder="input Access Password" id="vm_access_passwd_'+vm_cnt+'" name="vmUserPassword" ></td></tr>'

    +'<tr><th>Description</th><td>'
    +'<textarea class="form-control" name="description" id="exampleFormControlTextarea1" rows="3"></textarea></td></tr>'  
    
    +'</tbody></table></div>'
    +'<div class="card-footer">'
    +'<div class="d-flex justify-content-end align-items-center">'
    //+'<button type="button" class="add_vm btn btn-dark btn-sm">Add Server</button>'
    +'추가할 서버 : '
    +'&nbsp;'
    +'<input class="vm_view " type="text" id="vm_view_'+vm_cnt+'" name="" >'
    +'&nbsp;'
    +'<button type="button" onclick="btn_click();" class="btn btn-dark btn-sm">Confirm</button>'
    +'&nbsp;'
    +'<button type="button" class="btn btn-sm btn-danger" onclick="cancel_btn();">Cancel</button>'
    +'</div></div></div>'
    +'</form>'

    $("#target").after(html);
   // $("#main").append(html)

   alert("MCISRegister.html");
    getCloudOS(ApiInfo,'provider_' + vm_cnt);
    console.log("MCISRegister finished");
}   
function show_pop2() {
    // alert("show pop")
    //$("body:eq(0)").append("<div id='transDiv' style='background-color: #ffffff; position: absolute;  overflow: hidden;'></div>");
    $("#main").append("<div id='transDiv' style='background-color: #ffffff; position: absolute;  overflow: hidden;'></div>");
    // $("#transDiv").load("/Pop/spec");
    // $("#transDiv").show();
}
function view_vm(){

    var str = ""
    var arr = new Array();

    $("[class^='vm_name']").each(function(k,v){
    console.log("Key : ",k);
    console.log("value : ",v); 


    arr.push(this.value)

    });
    console.log("view name : ",arr);

    if(arr.length > 1){
    str = arr.join(", ");
    }else{
    str = arr[0];
    }

    console.log(str)
    $("[class^='vm_view']").each(function(k,v){
    this.value = str
    });

}
function btn_click(){

    //폼의 필수 필드를 검증 함.
    if (!chkFormValidate($("#form"))) return false;

    var result = setObj();
    result.name = $("#mcis_name").val()
    result.description = $("#mcis_description").val()
    //result.headers = {'Authorization': apiInfo,'Content-type': 'application/json',}
    console.log("seObj result : ",result)
    var url = CommonURL+"/ns/"+NAMESPACE+"/mcis"
    console.log("request Create mcis url : ",url)

    try{
        AjaxLoadingShow(true);
        axios.post(url,result,{
        headers :{
            'Content-type': 'application/json',
            'Authorization': apiInfo,
            },
        }).then(result=>{
        console.log("MCIR Register data : ",result);
        console.log("Result Status : ",result.status); 
        if(result.status == 201 || result.status == 200){
        alert("Register Success")
        location.href = "/MCIS/list";
        }else{
        alert("Register Fail")
        //location.reload(true);
        }
        })
    }finally{
        AjaxLoadingShow(false);
    }    
}


function setObj(){
    var clickCnt = btn_click_cnt
    var ov = new Array();
    var objArr = new Array()
    var t = 0;

    for(var i = 0; i < clickCnt;i++){
    t++;
    ov[i] = $("#form_"+t+"").serializeArray();
    }

    for(var i in ov){
    var obj  = {}
    var security_arr = new Array()
    for(var k in ov[i]){

    console.log(ov[i][k].name)
    var n = ov[i][k].name
    var v = ov[i][k].value
    //res.vm_req[i].ov[i].name = ov[i].value
    var sCnt = 0
    //var obj = res.vm_req[i]
    if(n == "securityGroupIds"){
    
        security_arr.push(v)
        sCnt++
        obj[n] = security_arr
    }else{
        obj[n] =   v 
    }
    }
    objArr[i] = obj;
    }
    var apiInfo = ApiInfo
    var result = {

    vm : objArr,
    vm_num : clickCnt
    }
    console.log("vm num : ",clickCnt);
    console.log("Object result :",result)
    return result
}