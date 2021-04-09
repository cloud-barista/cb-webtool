///////////// MCIS Handling //////////////

// 등록 form으로 이동
function createNewMCIS(){// Manage_MCIS_Life_Cycle_popup.html
    var url = "/operation/manage" + "/mcis/regform/"
    // location.href = "/Manage/MCIS/reg"
    location.href = url;
}

// MCIS 제어
function mcisLifeCycle(type){
    var checked_nothing = 0;
    $("[id^='td_ch_']").each(function(){
       
        if($(this).is(":checked")){
            checked_nothing++;
            console.log("checked")
            var mcis_id = $(this).val()
            console.log("check td value : ",mcis_id);
            var nameSpace = NAMESPACE;
            console.log("Start LifeCycle method!!!")
            var url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"?action="+type
            
            console.log("life cycle3 url : ",url);
            var message = "MCIS "+type+ " complete!."
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



////////////// VM Handling ///////////
function addNewVirtualMachine(){
    var mcis_id = $("#mcis_id").val()
    var mcis_name = $("#mcis_name").val()
    //location.href = "/Manage/MCIS/reg/"+mcis_id+"/"+mcis_name
}

function vmLifeCycle(type){
    var mcis_id = $("#mcis_id").val();
    var vm_id = $("#vm_id").val();
    var vm_name = $("#vm_name").val();
    
    // var checked =""
    // $("[id^='td_ch_'").each(function(){
    //     if($(this).is(":checked")){
    //         var checked_value = $(this).val();
    //         console.log("checked value : ",checked_value)
    //     }else{
    //         console.log("체크된게 없어!!")
    //     }
    // })
    // return;
    if(!mcis_id){
        alert("Please Select MCIS!!")
        return;
    }
    if(!vm_id){
        alert("Please Select VM!!")
        return;
    }
    
    var nameSpace = NAMESPACE;
    console.log("Start LifeCycle method!!!")
    var url ="";
    if(vm_id){
        
        url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"/vm/"+vm_id+"?action="+type 
       
    }
    console.log("life cycle3 url : ",url);
   
    var message = vm_name+" "+type+ " complete!."
  

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
///////////// VM Handling end ///////////



const config_arr = new Array();

// refresh : mcis 및 vm정보조회
// 각 mcis 별 vmstatus 목록


// List Of MCIS 클릭 시 
// mcis 테이블의 선택한 row 강조( on )
// 해당 MCIS의 VM 상태목록 보여주는 함수 호출
function clickListOfMCIS(id,index){
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

const mcisInfoDataList = new Array()// test_arr : mcisInfo 1개 전체, pageLoad될 때, refresh 할때 data를 set. mcis클릭시 상세내용 보여주기용 조회

// MCIS Info area 안의 Server List / Status 내용 표시
// 해당 MCIS의 모든 VM 표시
// TODO : 클릭했을 때 서버에서 조회하는것으로 변경할 것.
function showServerListAndStatusArea(mcis_id, mcisIndex){
    
    var mcisID =  $("#mcisID" + mcisIndex).val();
    var mcisName =  $("#mcisName" + mcisIndex).val();
    var mcisDescription =  $("#mcisDescription" + mcisIndex).val();
    var mcisStatus =  $("#mcisStatus" + mcisIndex).val();
    var mcisCloudConnections = $("#mcisCloudConnections" + mcisIndex).val();
    var vmTotalCountOfMcis = $("#mcisVMTotalCount" + mcisIndex).val();
    var vms = $("#mcisVMStatusList" + mcisIndex).val();

    $(".server_status").addClass("view")
    $("#mcis_info_txt").text("[ "+ mcisName +" ]");
    $("#mcis_server_info_status").empty();
    $("#mcis_server_info_status").append('<strong>Server List / Status</strong>  <span class="stxt">[ '+mcisName+' ]</span>  Server('+vmTotalCountOfMcis+')')

    //
    $("#mcis_info_name").val(mcisName+" / "+mcisID)
    $("#mcis_info_description").val(mcisDescription);
    // $("#mcis_info_targetStatus").val(targetStatus);
    // $("#mcis_info_targetAction").val(targetAction);
    $("#mcis_info_cloud_connection").val(mcisCloudConnections)    //
    
    $(".server_status").addClass("view")
   
    $("#mcis_name").val(mcisName)

    var mcis_badge = "";
    var mcisStatusIcon = "";
    if(mcisStatus == "running"){ mcisStatusIcon = "icon_running_db.png"        
    }else if(mcisStatus == "include" ){ mcisStatusIcon = "icon_stop_db.png"
    }else if(mcisStatus == "suspended"){mcisStatusIcon = "icon_stop_db.png"
    }else if(mcisStatus == "terminate"){mcisStatusIcon = "icon_terminate_db.png"
    }else{
        mcisStatusIcon = "icon_stop_db.png"
    }
    mcis_badge = '<img src="/assets/img/contents/' + mcisStatusIcon +'" alt=""/> '
    $("#service_status_icon").empty();
    $("#service_status_icon").append(mcis_badge)


    var sta = mcisStatus;
    var sl = sta.split("-");
    var status = sl[0].toLowerCase()
    var vm_badge = "";
    
    var vmList = vms.split("@") // vm목록은 @
    console.log("vmList " + vmList);
    // for(var x in vmList){
    for( var x= 0; x < vmList.length; x++){
        var vmInfo = vmList[x].split("|") // 이름과 상태는 "|"로 구분
        console.log("x " + x);
        console.log("vmInfo " + vmInfo);

        var vmIdAndName = vmInfo[0].split("+"); // ID와 이름은 "-"로 구분
        vmID = vmIdAndName[0];
        vmName = vmIdAndName[1];

        vmStatus = vmInfo[1].toLowerCase();

        var vmStatusIcon ="bgbox_g";
        
        if(vmStatus == "running"){ 
            vmStatusIcon ="bgbox_g"
        }else if(vmStatus == "include" ){
            vmStatusIcon ="bgbox_g"
            // vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcisID+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
        }else if(vmStatus == "suspended"){
            vmStatusIcon ="bgbox_g"
            // vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcisID+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
            
        }else if(vmStatus == "terminated"){
            vmStatusIcon ="bgbox_r"
            // vm_badge += '<li class="sel_cr bgbox_r"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcisID+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
            
        }else{
            vmStatusIcon ="bgbox_g"
            // vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcisID+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
        }
        vm_badge += '<li class="sel_cr ' + vmStatusIcon + '"><a href="javascript:void(0);" onclick="vmDetailInfo(\''+mcisID+'\',\''+mcisName+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
            
        $("#mcis_server_info_box").empty();
        $("#mcis_server_info_box").append(vm_badge);
    }

    //Manage MCIS Server List on/off : table을 클릭하면 해당 Row 에 active style로 보여주기
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

// VM 목록에서 VM 클릭시 해당 VM의 상세정보 
function vmDetailInfo(mcisID, mcisName, vmID){
    var url = "/operation/manage/mcis/" + mcisID + "/vm/" + vmID
    axios.get(url,{})
        .then(result=>{
            console.log("get  Data : ",result.data);
            var data = result.data.VMInfo;

            var vmId = data.id;
            var vmName = data.name;
            var vmStatus = data.status;
            var vmDescription = data.description;
             
            //vm info
            $("#vm_id").val(vmId);   
            $("#vm_name").val(vmName);

            $("#manage_mcis_popup_vm_id").val(vmId)
            $("#manage_mcis_popup_mcis_id").val(mcisID)

            
            $("#server_info_text").text('['+vmName+'/'+mcisName+']')
            $("#server_detail_info_text").text('['+vmName+'/'+mcisName+']')

            
            var vmBadge =""
            var vmStatusIcon = "icon_running_db.png";
            if(vmStatus == "Running"){
                vmStatusIcon = "icon_running_db.png";
            }else if(vmStatus == "include"){
                vmStatusIcon = "icon_stop_db.png";
            }else if(vmStatus == "Suspended"){
                vmStatusIcon = "icon_stop_db.png";
            }else if(vmStatus == "Terminated"){
                vmStatusIcon = "icon_terminate_db.png";
            }else{
                vmStatusIcon = "icon_stop_db.png";
            }
            vmBadge = '<img src="/assets/img/contents/' + vmStatusIcon +'" alt="' + vmStatus + '"/>'

            $("#server_detail_view_server_status").val(vmStatus);
            $("#server_info_status_img").empty()
            $("#server_info_status_img").append(vmBadge)

            $("#server_info_name").val(vmName +"/"+ vmID)
            $("#server_info_desc").val(vmDescription)

            // ip information
            $("#server_info_public_ip").val(data.publicIP)
            $("#server_detail_info_public_ip_text").text("Public IP : "+data.publicIP)
            $("#server_info_public_dns").val(data.publicDNS)
            $("#server_info_private_ip").val(data.privateIP)
            $("#server_info_private_dns").val(data.privateDNS)


            //vm detail tab

            // vm connection tab

            // vm mornitoring tab

        }
    ).catch(function(error){
        console.log(" display error : ",error);        
    });


    /////////////////////
//     $("#vm_id").val(vm_id);
   
//     // $("#vm_name").val(vm_name);

//     // Popup install monitoring agent set value
//     $("#manage_mcis_popup_vm_id").val(vm_id)
//     $("#manage_mcis_popup_mcis_id").val(mcis_id)

//    var select_mcis = test_arr.filter(mcis => mcis.id === mcis_id);
//    console.log("click_view_vm arr : ",select_mcis);
   
//    var vm_arr = select_mcis[0].vm
//    console.log("vm_arr : ",vm_arr);
//    vm_arr = vm_arr.filter(item => item.id === vm_id)
//    console.log("click_view_vm arr : ",vm_arr)
//    var mcis_name = select_mcis[0].name
//    var select_vm = vm_arr[0];
//    var vm_detail = select_vm.cspViewVmDetail
//    var vm_name = select_vm.name

//    $("#server_info_text").text('['+vm_name+'/'+mcis_name+']')
//    $("#server_detail_info_text").text('['+vm_name+'/'+mcis_name+']')

//    var vm_status = select_vm.status
//    var vm_badge =""
//    if(vm_status == "Running"){
//        vm_badge = '<img src="/assets/img/contents/icon_running_db.png" alt=""/> '
//    }else if(vm_status == "include" ){
//        vm_badge = '<img src="/assets/img/contents/icon_stop_db.png"  alt=""/> ';
//    }else if(vm_status == "Suspended"){
//        vm_badge = '<img src="/assets/img/contents/icon_stop_db.png" alt=""/>'
       
//    }else if(vm_status == "Terminated"){
//        vm_badge = '<img src="/assets/img/contents/icon_terminate_db.png" alt=""/>'
       
//    }else{
//        vm_badge = '<img src="/assets/img/contents/icon_stop_db.png" alt=""/>'
   
//    }
//    $("#server_detail_view_server_status").val(vm_status);
//    $("#server_info_status_img").empty()
//    $("#server_info_status_img").append(vm_badge)

//    $("#server_info_name").val(vm_name +"/"+ select_vm.id)
//    $("#server_info_desc").val(select_vm.description)

//    // ip information
//    $("#server_info_public_ip").val(select_vm.publicIP)
//    $("#server_detail_info_public_ip_text").text("Public IP : "+select_vm.publicIP)
//    $("#server_info_public_dns").val(select_vm.publicDNS)
//    $("#server_info_private_ip").val(select_vm.privateIP)
//    $("#server_info_private_dns").val(select_vm.privateDNS)


//    $("#server_detail_view_public_ip").val(select_vm.publicIP)
//    $("#server_detail_view_public_dns").val(select_vm.publicDNS)
//    $("#server_detail_view_private_ip").val(select_vm.privateIP)
//    $("#server_detail_view_private_dns").val(select_vm.privateDNS)

//    $("#manage_mcis_popup_public_ip").val(select_vm.publicIP)

//    //cspvmdetail
//    var vm_detail_keyValue = vm_detail.KeyValueList
//    var architecture = "";   
//    if(vm_detail_keyValue){

//        architecture = vm_detail_keyValue.filter(item => item.Key === "Architecture")
//        console.log("architecture : ",architecture.length)
//        if(architecture.length > 0){
//            architecture = architecture[0].Value
//            console.log("architecture2 : ",architecture)
           
//        }
//    }
  
   
//    $("#server_info_archi").val(architecture)
//    $("#server_detail_view_archi").val(architecture)

//    // server spec
//    var vm_spec_name = vm_detail.VMSpecName
//    $("#server_info_vmspec_name").val(vm_spec_name)
//    $("#server_detail_view_server_spec_text").text(vm_spec_name)
//    var spec_id = select_vm.specId
//    set_vmSpecInfo(spec_id);
   
//    // start time
//    var start_time = vm_detail.StartTime
//    $("#server_info_start_time").val(start_time)

//    // cloud type
//    var csp = select_vm.location.cloudType
//    var csp_icon = ""
//    if(csp == "aws"){
//        csp_icon = '<img src="/assets/img/contents/img_logo1.png" alt=""/>'
//    }
//    if(csp == "azure"){
//        csp_icon = '<img src="/assets/img/contents/img_logo5.png" alt=""/>'
//    }
//    if(csp == "gcp"){
//        csp_icon = '<img src="/assets/img/contents/img_logo7.png" alt=""/>'
//    }
//    if(csp == "cloudit"){
//        csp_icon = '<img src="/assets/img/contents/img_logo6.png" alt=""/>'
//    }
//    if(csp == "openstack"){
//        csp_icon = '<img src="/assets/img/contents/img_logo9.png" alt=""/>'
//    }
//    if(csp == "alibaba"){
//        csp_icon = '<img src="/assets/img/contents/img_logo4.png" alt=""/>'
//    }

//    $("#server_info_csp_icon").empty()
//    $("#server_info_csp_icon").append(csp_icon)
//    $("#server_connection_view_csp").val(csp)
//    $("#manage_mcis_popup_csp").val(csp)


//    // region zone locate
//    var locate = select_vm.location.briefAddr
//    var region = select_vm.region.Region
//    var zone = select_vm.region.Zone

//    $("#server_info_region").val(locate +":"+region)
//    $("#server_info_zone").val(zone)
//    $("#server_info_cspVMID").val("cspVMID : "+vm_detail.IId.NameId)

//    $("#server_detail_view_region").val(locate +":"+region)
//    $("#server_detail_view_zone").val(zone)

//    $("#server_connection_view_region").val(locate +"("+region+")")
//    $("#server_connection_view_zone").val(zone)

//    // connection name
//    var connection_name = select_vm.connectionName;
//    $("#server_info_connection_name").val(connection_name)
//    $("#server_connection_view_connection_name").val(connection_name)

//    // credential and driver info
//    console.log("config arr2 : ",config_arr)
//    console.log("connection_name :",connection_name)
//    var arr_config = config_arr
//    console.log("arr_config : ",arr_config);
//    if(arr_config){
//        var config_info = arr_config.filter(cred => cred.ConfigName === connection_name)[0]
//        console.log("inner config info : ",config_info)
//        console.log("config_info : ",config_info)
//        var credentialName = config_info.CredentialName
//        var driverName = config_info.DriverName
//        $("#server_connection_view_credential_name").val(credentialName)
//        $("#server_connection_view_driver_name").val(driverName)
//    }
  

   
   
//    // server id / system id
//    $("#server_detail_view_server_id").val(select_vm.id)
//    // systemid 를 사용할 경우 아래 꺼 사용
//    //$("#server_detail_view_server_id").val(vm_detail.IId.SystemId)
  
//    // image id
//    var imageIId = vm_detail.ImageIId.NameId
//    var imageId = select_vm.imageId
//    set_vmImageInfo(imageId)
//    $("#server_detail_view_image_id_text").text(imageId+"("+imageIId+")")

//    //vpc subnet
//    var vpcId = vm_detail.VpcIID.NameId
//    var vpcSystemId = vm_detail.VpcIID.SystemId
//    var subnetId = vm_detail.SubnetIID.NameId
//    var subnetSystemId = vm_detail.SubnetIID.SystemId
//    var eth = vm_detail.NetworkInterface
//    $("#server_detail_view_vpc_id_text").text(vpcId+"("+vpcSystemId+")")
//    set_vmVPCInfo(vpcId, subnetId);

//    $("#server_detail_view_subnet_id_text").text(subnetId+"("+subnetSystemId+")")
//    $("#server_detail_view_eth_text").val(eth)

//    // install Mon agent
//    var installMonAgent = select_vm.monAgentStatus
//    checkDragonFly(mcis_id,vm_id)

//    // device info
//    var root_device_type = vm_detail.VMBootDisk
//    var root_device = vm_detail.VMBootDisk
//    var block_device = vm_detail.VMBlockDisk
//    $("#server_detail_view_root_device_type").val(root_device_type)
//    $("#server_detail_view_root_device").val(root_device)
//    $("#server_detail_view_block_device").val(block_device)

//     // key pair info
   
//     $("#server_detail_view_keypair_name").val(vm_detail.KeyPairIId.NameId)
//     var sshkey = vm_detail.KeyPairIId.NameId
//     if(sshkey){
//        set_vmSSHInfo(sshkey)
//     }
//     // user account
//     $("#server_detail_view_access_id_pass").val(vm_detail.VMUserId +"/"+vm_detail.VMUserPasswd)
//     $("#server_detail_view_user_id_pass").val(select_vm.vmUserAccount +"/"+select_vm.vmUserPassword)
//     $("#manage_mcis_popup_user_name").val(select_vm.vmUserAccount)
    
//     // namespace 
//     var ns_id = NAMESPACE
//     $("#manage_mcis_popup_ns_id").val(ns_id)
    

//     // security Gorup
//    var append_sg = ''

//    var sg_arr = vm_detail.SecurityGroupIIds
//    if(sg_arr){
//        //여기서 호출해서 세부 값을 가져 오자
       
//        sg_arr.map((item,index)=>{
           
//            append_sg +='<a href="javascript:void(0);" onclick="set_vmSecurityGroupInfo(\''+item.NameId+'\');"title="'+item.NameId+'" >'+item.NameId+'</a> '
//        })
//    }
  
//    console.log("append sg : ",append_sg)
   
//    $("#server_detail_view_security_group").empty()
//    $("#server_detail_view_security_group").append(append_sg);

}