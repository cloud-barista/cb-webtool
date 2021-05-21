///////////// MCIS Handling //////////////

// 등록 form으로 이동
function createNewMcis(){// Manage_MCIS_Life_Cycle_popup.html
    var targetUrl = "/operation/manage" + "/mcismng/regform"
    // location.href = "/Manage/MCIS/reg"
    // $('#loadingContainer').show();
    // location.href = url;
    changePage(targetUrl)
}

// MCIS 제어 : 선택한 VM의 상태 변경 
// callMcisLifeCycle -> util.callMcisLifeCycle -> callbackMcisLifeCycle 순으로 호출 됨
function callMcisLifeCycle(type){
    var checked_nothing = 0;
    $("[id^='td_ch_']").each(function(){
       
        if($(this).is(":checked")){
            checked_nothing++;
            console.log("checked")
            var mcisID = $(this).val()
            mcisLifeCycle(mcisID, type);           
        }else{
            console.log("checked nothing")
           
        }
    })
    if(checked_nothing == 0){
        commonAlert("Please Select MCIS!!")
        return;
    }
}

// McisLifeCycle을 호출 한 뒤 return값 처리
function callbackMcisLifeCycle(resultStatus, resultData, type){
    var message = "MCIS "+type+ " complete!."
    if(resultStatus == 200 || resultStatus == 201){            
        commonAlert(message);
        location.reload();//완료 후 페이지를 reload -> 해당 mcis만 reload
        // 해당 mcis 조회
        // 상태 count 재설정
    }else{
        commonAlert("MCIS " + type + " failed!");
    }
}

// list에 선택된 MCIS 제거. 1개씩
function deleteMCIS(){
    var checkedCount = 0;
    var mcisID = "";
    $("[id^='td_ch_']").each(function(){
       
        if($(this).is(":checked")){
            checkedCount++;
            console.log("checked")
            mcisID = $(this).val();
            // 여러개를 지울 때 호출하는 함수를 만들어 여기에서 호출
        }else{
            console.log("checked nothing")
           
        }
    })

    if(checkedCount == 0){
        commonAlert("Please Select MCIS!!")
        return;
    }else if( checkedCount > 1){
        commonAlert("Please Select One MCIS at a time")
        return;
    }

    // TODO : 삭제 호출부분 function으로 뺼까?
    var url = "/operation/manages/mcismng/" + mcisID;               
    axios.delete(url,{})
        .then(result=>{
            console.log("get  Data : ",result.data);

            var statusCode = result.data.status;
            var message = result.data.message;
            
            if( statusCode != 200 && statusCode != 201) {
                commonAlert(message +"(" + statusCode + ")");
                return;
            }else{
                commonAlert(message);
                // TODO : MCIS List 조회
                location.reload();
            }
            
        }).catch((error) => {
            console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            commonErrorAlert(statusCode, errorMessage) 
        });

}
////////////// MCIS Handling end //////////////// 



////////////// VM Handling ///////////
function addNewVirtualMachine(){
    var mcis_id = $("#mcis_id").val()
    var mcis_name = $("#mcis_name").val()
    // location.href = "/Manage/MCIS/reg/"+mcis_id+"/"+mcis_name
    location.href = "/operation/manages/mcismng/regform/"+mcis_id+"/"+mcis_name;
}

function vmLifeCycle(type){
    var mcisID = $("#mcis_id").val();
    var vmID = $("#vm_id").val();
    var vmName = $("#vm_name").val();
    
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
    if(!mcisID){
        commonAlert("Please Select MCIS!!")
        return;
    }
    if(!vmID){
        commonAlert("Please Select VM!!")
        return;
    }
    
    // var nameSpace = NAMESPACE;
    console.log("Start LifeCycle method!!!")

    //url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"/vm/"+vm_id+"?action="+type
    //var url = "/operation/manage" + "/mcis/" + mcis + "/vm/" + vm_id + "/action/" + type
    var url = "/operation/manages" + "/mcismng/proc/vmlifecycle";
    // url = "http://54.248.3.145:1323/tumblebug/ns/ns-01/mcis/mz-azure-mcis/vm/mz-azure-ubuntu1804-5?action=suspend";
    // var apiInfo = "Basic ZGVmYXVsdDpkZWZhdWx0"
    
    // /////////
    // axios.get(url,{
    //         headers:{
    //                 'Authorization': apiInfo
    //             }
    //         }).then(result=>{
    //             var data = result.data
    //             console.log(data)
    //         });
    /////////
    // + mcis + "/vm/" + vm_id + "/action/" + type
    
    console.log("life cycle3 url : ",url);
   
    var message = vmName+" "+type+ " complete!."  
    var obj = {
        mcisID : mcisID,
        vmID : vmID,
        lifeCycleType : type
     }
    axios.post(url,obj,{
        headers: { 
            'Content-type': 'application/json',
            // 'Authorization': apiInfo, 
        }
    // })
    // axios.post(url,{
    //     headers: { },
        // mcisID:mcis_id,
        // vmID:vm_id,
        // vmLifeCycleType:type
    }).then(result=>{
        var status = result.status
        
        console.log("life cycle result : ",result)
        var data = result.data
        console.log("result Message : ",data.message)
        if(status == 200 || status == 201){            
            commonAlert(message);
            location.reload();// TODO 일단은 Reaoad : 해당 영역(MCIS의 VM들 status 조회)를 refresh할 수 있는 기능 필요
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
function clickListOfMcis(id,index){
    console.log("click view mcis id :",id)
    $(".server_status").addClass("view");
    
    // 다른 area가 활성화 되어 있으면 안보이게
    $("#dashboard_detailBox").removeClass("active")

    // List Of MCIS에서 선택한 row 외에는 안보이게
    $("[id^='server_info_tr_']").each(function(){
        var item = $(this).attr("item").split("|")
        console.log(item)
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
    var vmTotalCountOfMcis = $("#mcisVmTotalCount" + mcisIndex).val();
    var vms = $("#mcisVmStatusList" + mcisIndex).val();

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

    var vm_badge = "";
    $("[id^='mcisVmID_']").each(function(){		
        var mcisVm = $(this).attr("id").split("_")
        thisMcisIndex = mcisVm[1]
        vmIndexOfMcis = mcisVm[2]

        if( thisMcisIndex == mcisIndex){

            var vmID = $("#mcisVmID_" + mcisIndex + "_" + vmIndexOfMcis).val();
            var vmName = $("#mcisVmName_" + mcisIndex + "_" + vmIndexOfMcis).val();
            var vmStatus = $("#mcisVmStatus_" + mcisIndex + "_" + vmIndexOfMcis).val();
            vmStatus = vmStatus.toLowerCase();
            
            var vmStatusIcon ="bgbox_g";            
            if(vmStatus == "running"){ 
                vmStatusIcon ="bgbox_b"
            }else if(vmStatus == "include" ){
                vmStatusIcon ="bgbox_g"
            }else if(vmStatus == "suspended"){
                vmStatusIcon ="bgbox_g"
            }else if(vmStatus == "terminated"){
                vmStatusIcon ="bgbox_r"
            }else{
                vmStatusIcon ="bgbox_g"
            }
            vm_badge += '<li class="sel_cr ' + vmStatusIcon + '"><a href="javascript:void(0);" onclick="vmDetailInfo(\''+mcisID+'\',\''+mcisName+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
        }
    });
    // console.log(vm_badge);
    $("#mcis_server_info_box").empty();
    $("#mcis_server_info_box").append(vm_badge);

    // var sta = mcisStatus;
    // var sl = sta.split("-");
    // var status = sl[0].toLowerCase()
    // var vm_badge = "";
    
    // var vmList = vms.split("@") // vm목록은 @
    // console.log("vmList " + vmList);
    // // for(var x in vmList){
    // for( var x= 0; x < vmList.length; x++){
    //     var vmInfo = vmList[x].split("|") // 이름과 상태는 "|"로 구분
    //     console.log("x " + x);
    //     console.log("vmInfo " + vmInfo);

    //     vmID = vmInfo[0];
    //     vmName = vmInfo[1];

    //     vmStatus = vmInfo[1].toLowerCase();

    //     var vmStatusIcon ="bgbox_g";
        
    //     if(vmStatus == "running"){ 
    //         vmStatusIcon ="bgbox_b"
    //     }else if(vmStatus == "include" ){
    //         vmStatusIcon ="bgbox_g"
    //         // vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcisID+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
    //     }else if(vmStatus == "suspended"){
    //         vmStatusIcon ="bgbox_g"
    //         // vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcisID+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
            
    //     }else if(vmStatus == "terminated"){
    //         vmStatusIcon ="bgbox_r"
    //         // vm_badge += '<li class="sel_cr bgbox_r"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcisID+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
            
    //     }else{
    //         vmStatusIcon ="bgbox_g"
    //         // vm_badge += '<li class="sel_cr bgbox_g"><a href="javascript:void(0);" onclick="click_view_vm(\''+mcisID+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
    //     }
    //     vm_badge += '<li class="sel_cr ' + vmStatusIcon + '"><a href="javascript:void(0);" onclick="vmDetailInfo(\''+mcisID+'\',\''+mcisName+'\',\''+vmID+'\')"><span class="txt">'+vmName+'</span></a></li>';
    //     //console.log(vm_badge);
    //     $("#mcis_server_info_box").empty();
    //     $("#mcis_server_info_box").append(vm_badge);
    // }

    //Manage MCIS Server List on/off : table을 클릭하면 해당 Row 에 active style로 보여주기
    $(".dashboard .ds_cont .area_cont .listbox li.sel_cr").each(function(){
        var $sel_list = $(this);
        var $detail = $(".server_info");
        console.log($sel_list);
        console.log($detail);
        console.log(">>>>>");
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
    var url = "/operation/manages/mcismng/" + mcisID + "/vm/" + vmID
    axios.get(url,{})
        .then(result=>{
            console.log("get  Data : ",result.data);

            var statusCode = result.data.status;
            var message = result.data.message;
            
            if( statusCode != 200 && statusCode != 201) {
                commonAlert(message +"(" + statusCode + ")");
                return;
            }
            var data = result.data.VmInfo;
            
            var vmId = data.id;
            var vmName = data.name;
            var vmStatus = data.status;
            var vmDescription = data.description;
            var vmPublicIp = data.publicIP === undefined ? "" : data.publicIP;
            //vm info
            $("#vm_id").val(vmId);   
            $("#vm_name").val(vmName);
            console.log("vm_id " + vmId + ", vm_name " + vmName)

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
            $("#server_info_public_ip").val(vmPublicIp)
            $("#server_detail_info_public_ip_text").text("Public IP : "+vmPublicIp)
            $("#server_info_public_dns").val(data.publicDNS)
            $("#server_info_private_ip").val(data.privateIP)
            $("#server_info_private_dns").val(data.privateDNS)

            $("#server_detail_view_public_ip").val(vmPublicIp)
            $("#server_detail_view_public_dns").val(data.publicDNS)
            $("#server_detail_view_private_ip").val(data.privateIP)
            $("#server_detail_view_private_dns").val(data.privateDNS)
        
            $("#manage_mcis_popup_public_ip").val(vmPublicIp)


            //////vm detail tab////
            var vmDetail = data.cspViewVmDetail
            //    //cspvmdetail
            var vmDetailKeyValueList = vmDetail.KeyValueList
            var architecture = "";   
            if(vmDetailKeyValueList){
                for (var key in vmDetailKeyValueList) {
                    if( key == "Architecture"){// ?? 이게 뭐지?
                        architecture = architecture[key].Value  
                        break;
                    }
                }
                // architecture = vmDetailKeyValueList.filter(item => item.Key === "Architecture")
                // console.log("architecture : ",architecture.length)
                // if(architecture.length > 0){
                //     architecture = architecture[0].Value
                //     console.log("architecture2 : ",architecture)                    
                // }
                console.log("architecture = " + architecture)
                $("#server_info_archi").val(architecture)
                $("#server_detail_view_archi").val(architecture)
            }
            //    // server spec
            // var vmSecName = data.VmSpecName
            var vmSecName = data.vmspecName// TODO : 바로 return하는 경우인가?? 이름이 모두 소문자네
            $("#server_info_vmspec_name").val(vmSecName)
            $("#server_detail_view_server_spec_text").text(vmSecName)
            //var spec_id = data.specId
            var specId = data.specId
            // set_vmSpecInfo(spec_id);// memory + cpu  : TODO : spec정보는 자주변경되는것이 아닌데.. 매번 통신할 필요있나...
           
            var startTime = vmDetail.StartTime
            $("#server_info_start_time").val(startTime)

            //    // server spec
            var vmSpecName = vmDetail.VmSpecName
            $("#server_info_vmspec_name").val(vmSpecName)
            $("#server_detail_view_server_spec_text").text(vmSpecName)
            
            var cloudType = data.location.cloudType
            var cspIcon = ""
            if(cloudType == "aws"){
                cspIcon = "img_logo1"
            }else if(cloudType == "azure"){
                cspIcon = "img_logo5"
            }else if(cloudType == "gcp"){
                cspIcon = "img_logo7"
            }else if(cloudType == "cloudit"){
                cspIcon = "img_logo6"
            }else if(cloudType == "openstack"){
                cspIcon = "img_logo9"
            }else if(cloudType == "alibaba"){
                cspIcon = "img_logo4"
            }else{
                csp_icon = '<img src="/assets/img/contents/img_logo1.png" alt=""/>'
            }
            $("#server_info_csp_icon").empty()
            $("#server_info_csp_icon").append('<img src="/assets/img/contents/' + cspIcon + '.png" alt=""/>')
            $("#server_connection_view_csp").val(cloudType)
            $("#manage_mcis_popup_csp").val(cloudType)

            // region zone locate
            var locate = data.location.briefAddr
            var region = data.region.region
            var zone = data.region.zone
            console.log(vmDetail.iid);
            $("#server_info_region").val(locate +":"+region)
            $("#server_info_zone").val(zone)
            $("#server_info_cspVMID").val("cspVMID : "+vmDetail.iid.nameId)

            $("#server_detail_view_region").val(locate +":"+region)
            $("#server_detail_view_zone").val(zone)

            $("#server_connection_view_region").val(locate +"("+region+")")
            $("#server_connection_view_zone").val(zone)

            // connection name
            var connectionName = data.connectionName;
            $("#server_info_connection_name").val(connectionName)
            $("#server_connection_view_connection_name").val(connectionName)

            // credential and driver info
            console.log("config arr2 : ",config_arr)
            console.log("connection_name :",connectionName)
            // var arr_config = config_arr
            // console.log("arr_config : ",arr_config);
            // if(arr_config){
            //     var config_info = arr_config.filter(cred => cred.ConfigName === connection_name)[0]
            //     console.log("inner config info : ",config_info)
            //     console.log("config_info : ",config_info)
            //     var credentialName = config_info.CredentialName
            //     var driverName = config_info.DriverName
            //     $("#server_connection_view_credential_name").val(credentialName)
            //     $("#server_connection_view_driver_name").val(driverName)
            // }

            // server id / system id
            $("#server_detail_view_server_id").val(data.id)
            // systemid 를 사용할 경우 아래 꺼 사용
            $("#server_detail_view_server_id").val(vmDetail.iid.systemId)
            
            // image id
            var imageIId = vmDetail.imageIId.nameId
            var imageId = data.imageId
            // set_vmImageInfo(imageId) // 
            $("#server_detail_view_image_id_text").text(imageId+"("+imageIId+")")

            //vpc subnet
            var vpcId = vmDetail.vpcIID.nameId
            var vpcSystemId = vmDetail.vpcIID.systemId
            var subnetId = vmDetail.subnetIID.nameId
            var subnetSystemId = vmDetail.subnetIID.systemId
            var eth = vmDetail.networkInterface
            $("#server_detail_view_vpc_id_text").text(vpcId+"("+vpcSystemId+")")
            // set_vmVPCInfo(vpcId, subnetId);

            $("#server_detail_view_subnet_id_text").text(subnetId+"("+subnetSystemId+")")
            $("#server_detail_view_eth_text").val(eth)

            // ... TODO : 우선 제어명령부터 처리. 나중에 해당항목 mapping하여 확인 
            ////// vm connection tab //////

            ////// vm mornitoring tab //////
            // install Mon agent
            var installMonAgent = data.monAgentStatus
            showVmMonitoring(mcisID,vmID)
        }
    // ).catch(function(error){
    //     var statusCode = error.response.data.status;
    //     var message = error.response.data.message;
    //     commonErrorAlert(statusCode, message)        
    // });
    ).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        commonErrorAlert(statusCode, errorMessage) 
    });


    /////////////////////


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


// 조회 성공 시 Monitoring Tab 표시
function showVmMonitoring(mcisID, vmID){
    $("#mcis_detail_info_check_monitoring").prop("checked",true)
    $("#mcis_detail_info_check_monitoring").attr("disabled",true)
    $("#Monitoring_tab").show();
    var duration = "5m"
    var period_type = "m"
    var metric_arr = ["cpu","memory","disk","network"];
    var statisticsCriteria = "last";
    for(var i in metric_arr){
        getVmMetric("canvas_"+i,metric_arr[i],mcisID,vmID,metric_arr[i],period_type,statisticsCriteria,duration);
    }    
 }
 

// getVMMetric 는 mcis.chart.js로 이동 