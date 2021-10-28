$(document).ready(function () {
    getVmList()
    getCommonNetworkList('vmcreate')
    getCommonVirtualMachineImageList('vmcreate')
    getCommonVirtualMachineSpecList('vmcreate')
    getCommonSecurityGroupList('vmcreate')
    getCommonSshKeyList('vmcreate')
    // e_vNetListTbody

    $('#alertResultArea').on('hidden.bs.modal', function () {// bootstrap 3 또는 4
        //$('#alertResultArea').on('hidden', function () {// bootstrap 2.3 이전
        let targetUrl = "/operation/manages/mcismng/mngform"
        changePage(targetUrl)
    })

    //OS_HW popup table scrollbar
    $('#OS_HW .btn_spec').on('click', function () {
        console.log("os_hw bpn_spec clicked ** ")
        $('#OS_HW_Spec .dtbox.scrollbar-inner').scrollbar();

        // connection 정보 set
        var esSelectedProvider = $("#es_regProvider option:selected").val();
        var esSelectedRegion = $("#es_regRegion option:selected").val();
        var esSelectedConnectionName = $("#es_regConnectionName option:selected").val();

        console.log("OS_HW_Spec_Assist click");
        if (esSelectedProvider) {
            $("#assist_select_provider").val(esSelectedProvider);
        }
        if (esSelectedRegion) {
            $("#assist_select_resion").val(esSelectedRegion);
        }
        if (esSelectedConnectionName) {
            $("#assist_select_connectionName").val(esSelectedConnectionName);
        }

        console.log("esSelectedProvider = " + esSelectedProvider + " : " + $("#assist_select_provider").val());
        console.log("esSelectedRegion = " + esSelectedRegion + " : " + $("#assist_select_resion").val());
        console.log("esSelectedConnectionName = " + esSelectedConnectionName + " : " + $("#assist_select_connectionName").val());
    });
    //Security popup table scrollbar
    $('#Security .btn_edit').on('click', function () {
        $("#security_edit").modal();
        $('#security_edit .dtbox.scrollbar-inner').scrollbar();
    });

    // $("input[name='vmInfoType']:radio").change(function () {
    //     //라디오 버튼 값을 가져온다.
    //     var formType = this.value;

    // });


    // server add 버튼 클릭 시
    // $('.servers_box .server_add').click(function(){	

    //     //<div class="servers_config import_servers_config" id="importServerConfig">
    //     //<div class="servers_config new_servers_config" id="expertServerConfig">
    // });

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
    //       if($Servers.hasClass("active")) {
    //         $Servers.toggleClass("active");
    //     } else {
    //         $Servers.toggleClass("active");
    //     }
    //     });
    //     // Simple add
    //   $(".servers_box .switch").change(function() {
    //     if ($(".switch .ch").is(":checked")) {	
    //             $('.servers_box .server_add').click(function(){	

    //                 $NewServers.addClass("active");
    //                 $SimpleServers.removeClass("active");		
    //             });
    //     } else {
    //             $('.servers_box .server_add').click(function(){

    //                 $NewServers.removeClass("active");
    //                 $SimpleServers.addClass("active");


    //             });		
    //     }
    //   });

});


var totalDeployServerCount = 0;
function btn_deploy() {
    var mcis_name = $("#mcis_name").val();
    var mcis_id = $("#mcis_id").val();
    if (!mcis_id) {
        commonAlert("Please Select MCIS !!!!!")
        return;
    }
    totalDeployServerCount = 0;// deploy vm 개수 초기화

    console.log(Simple_Server_Config_Arr);
    if (Simple_Server_Config_Arr) {// mcissimpleconfigure.js 에 const로 정의 됨.
        var vm_len = Simple_Server_Config_Arr.length;
        if (vm_len > 0) {
            totalDeployServerCount += vm_len
            console.log("Simple_Server_Config_Arr length: ", vm_len);
            // var new_obj = {}
            // new_obj['vm'] = Simple_Server_Config_Arr;
            // console.log("new obj is : ",new_obj);
            // var url = "/operation/manages/mcis/:mcisID/vm/reg/proc"
            var url = "/operation/manages/mcismng/" + mcis_id + "/vm/reg/proc"

            // 한개씩 for문으로 추가
            for (var i in Simple_Server_Config_Arr) {
                new_obj = Simple_Server_Config_Arr[i];
                console.log("new obj is : ", new_obj);
                try {
                    resultVmCreateMap.set(new_obj.name, "")
                    axios.post(url, new_obj, {
                        headers: {
                        },
                    }).then(result => {
                        console.log("MCIR VM Register data : ", result);
                        console.log("Result Status : ", result.status);

                        var statusCode = result.data.status;
                        var message = result.data.message;
                        console.log("Result Status : ", statusCode);
                        console.log("Result message : ", message);

                        if (result.status == 201 || result.status == 200) {
                            vmCreateCallback(new_obj.name, "Success")
                            //     commonAlert("Register Success")
                            //     // location.href = "/Manage/MCIS/list";
                            //     // $('#loadingContainer').show();
                            //     // location.href = "/operation/manages/mcis/mngform/"
                            //     var targetUrl = "/operation/manages/mcis/mngform"
                            //     changePage(targetUrl)
                        } else {
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
                } finally {

                }

                // post로 호출을 했으면 해당 VM의 정보는 비활성시킨 후(클릭 Evnet 안먹게)
                // 상태값을 모니터링 하여 결과 return 까지 대기.
                // return 받으면 해당 VM
            }
        }
    }

    ///////// export
    console.log(Expert_Server_Config_Arr);
    if (Expert_Server_Config_Arr) {
        var vm_len = Expert_Server_Config_Arr.length;
        console.log("Expert_Server_Config_Arr length: ", vm_len);
        if (vm_len > 0) {
            totalDeployServerCount += vm_len
            // var new_obj = {}
            // new_obj['vm'] = Simple_Server_Config_Arr;
            // console.log("new obj is : ",new_obj);
            // var url = "/operation/manages/mcis/:mcisID/vm/reg/proc"
            var url = "/operation/manages/mcismng/" + mcis_id + "/vm/reg/proc"

            // 한개씩 for문으로 추가
            for (var i in Expert_Server_Config_Arr) {
                new_obj = Expert_Server_Config_Arr[i];
                console.log("new obj is : ", new_obj);
                try {
                    resultVmCreateMap.set("Expert" + i, "")
                    axios.post(url, new_obj, {
                        headers: {
                        },
                    }).then(result => {
                        console.log("MCIR VM Register data : ", result);
                        console.log("Result Status : ", result.status);

                        var statusCode = result.data.status;
                        var message = result.data.message;
                        console.log("Result Status : ", statusCode);
                        console.log("Result message : ", message);

                        if (result.status == 201 || result.status == 200) {
                            vmCreateCallback("Expert" + i, "Success")
                        } else {
                            vmCreateCallback("Expert" + i, message)
                        }
                    }).catch((error) => {
                        // console.warn(error);
                        console.log(error.response)
                        var errorMessage = error.response.data.error;
                        // commonErrorAlert(statusCode, errorMessage) 
                        vmCreateCallback("Expert" + i, errorMessage)
                    })
                } finally {

                }

                // post로 호출을 했으면 해당 VM의 정보는 비활성시킨 후(클릭 Evnet 안먹게)
                // 상태값을 모니터링 하여 결과 return 까지 대기.
                // return 받으면 해당 VM
            }
        }
    }
    ///////// import
    if (Import_Server_Config_Arr) {// mcissimpleconfigure.js 에 const로 정의 됨.
        // TODO : 어차피 simple/expert와 로직이 다른데... 
        // json 그대로 넘기도록
        var vm_len = Import_Server_Config_Arr.length;
        if (vm_len > 0) {
            console.log("Import_Server_Config_Arr length: ", vm_len);
            totalDeployServerCount += vm_len
            // var new_obj = {}
            // new_obj['vm'] = Simple_Server_Config_Arr;
            // console.log("new obj is : ",new_obj);
            // var url = "/operation/manages/mcis/:mcisID/vm/reg/proc"
            var url = "/operation/manages/mcismng/" + mcis_id + "/vm/reg/proc"

            // 한개씩 for문으로 추가
            for (var i in Import_Server_Config_Arr) {
                new_obj = Import_Server_Config_Arr[i];
                console.log("new obj is : ", new_obj);
                try {
                    resultVmCreateMap.set("Import" + i, "")
                    axios.post(url, new_obj, {
                        headers: {
                        },
                    }).then(result => {
                        console.log("MCIR VM Register data : ", result);
                        console.log("Result Status : ", result.status);

                        var statusCode = result.data.status;
                        var message = result.data.message;
                        console.log("Result Status : ", statusCode);
                        console.log("Result message : ", message);

                        if (result.status == 201 || result.status == 200) {
                            vmCreateCallback("Import" + i, "Success")
                            //     commonAlert("Register Success")
                            //     // location.href = "/Manage/MCIS/list";
                            //     // $('#loadingContainer').show();
                            //     // location.href = "/operation/manages/mcis/mngform/"
                            //     var targetUrl = "/operation/manages/mcis/mngform"
                            //     changePage(targetUrl)
                        } else {
                            vmCreateCallback("Import" + i, message)
                            //     commonAlert("Register Fail")
                            //     //location.reload(true);
                        }
                    }).catch((error) => {
                        // console.warn(error);
                        console.log(error.response)
                        var errorMessage = error.response.data.error;
                        // commonErrorAlert(statusCode, errorMessage) 
                        vmCreateCallback("Import" + i, errorMessage)
                    })
                } finally {

                }

                // post로 호출을 했으면 해당 VM의 정보는 비활성시킨 후(클릭 Evnet 안먹게)
                // 상태값을 모니터링 하여 결과 return 까지 대기.
                // return 받으면 해당 VM
            }
        }
    }
}

// Import / Export Modal 표시
function btn_ImportExport() {
    // export할 VM을 선택한 후 export 버튼 누르라고...
    $("#VmImportExport").modal();
    $('#VmImportExport .dtbox.scrollbar-inner').scrollbar();
}

// vm 생성 결과 표시
// 여러개의 vm이 생성될 수 있으므로 각각 결과를 표시
var resultVmCreateMap = new Map();
function vmCreateCallback(resultVmKey, resultStatus) {
    resultVmCreateMap.set(resultVmKey, resultStatus)
    var resultText = "";
    var createdServer = 0;
    for (let key of resultVmCreateMap.keys()) {
        console.log("vmCreateresult " + key + " : " + resultVmCreateMap.get(resultVmKey));
        resultText += key + " = " + resultVmCreateMap.get(resultVmKey) + ","
        //totalDeployServerCount--
        createdServer++;
    }

    $("#serverRegistResult").text(resultText);

    if (resultStatus != "Success") {
        // add된 항목 제거 해야 함.

        // array는 초기화
        Simple_Server_Config_Arr.length = 0;
        simple_data_cnt = 0
        // TODO : expert 추가하면 주석 제거할 것
        Expert_Server_Config_Arr.length = 0;
        expert_data_cnt = 0
        Import_Server_Config_Arr.length = 0;
        import_data_cnt = 0
    }

    if (createdServer === totalDeployServerCount) { //모두 성공
        //getVmList();
        commonResultAlert($("#serverRegistResult").text());

    } else if (createdServer < totalDeployServerCount) { //일부 성공
        commonResultAlert($("#serverRegistResult").text());

    } else if (createdServer = 0) { //모두 실패
        commonResultAlert($("#serverRegistResult").text());

    }
}

// 현재 mcis의 vm 목록 조회 : deploy후 상태볼 때 사용
function getVmList() {

    console.log("getVmList()")
    var mcis_id = $("#mcis_id").val();


    // /operation/manages/mcis/:mcisID
    var url = "/operation/manages/mcismng/" + mcis_id
    axios.get(url, {})
        .then(result => {
            console.log("MCIR VM Register data : ", result);
            console.log("Result Status : ", result.status);

            var statusCode = result.data.status;
            var message = result.data.message;
            //
            console.log("Result Status : ", statusCode);
            console.log("Result message : ", message);


            if (result.status == 201 || result.status == 200) {
                var mcis = result.data.McisInfo
                console.log(mcis)


                var vms = mcis.vm
                if (vms) {
                    vm_len = vms.length

                    $("#mcis_server_list *").remove();
                    var appendLi = "";

                    for (var o in vms) {
                        var vm_status = vms[o].status
                        var vm_name = vms[o].name

                        console.log(o + "번째 " + vm_name + " : " + vm_status)
                        // mcis_server_list 밑의 li들을 1개빼고 삭제. 
                        // 가져온 vm list 를 add? (1개는 더하기 버튼이므로)                    


                        appendLi = appendLi + '<li>';
                        appendLi = appendLi + '<div class="server server_on bgbox_g">';
                        appendLi = appendLi + '<div class="icon"></div>';
                        appendLi = appendLi + '<div class="txt">' + vm_name + '</div>';
                        appendLi = appendLi + '</li>';

                        appendLi = appendLi + '</li>';

                    }
                    appendLi = appendLi + '<li>';
                    appendLi = appendLi + '<div class="server server_add" onClick="displayNewServerForm()">';
                    appendLi = appendLi + '</div>';
                    appendLi = appendLi + '</li>';

                    $("#mcis_server_list").append(appendLi);

                    // commonAlert("VM 목록 조회 완료")
                    //$("#serverRegistResult").text("VM 목록 조회 완료");
                }
            }
        }).catch((error) => {
            // console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
        })
}

// 화면 Load시 가져오나 굳이?
function getNetworkListCallbackSuccess(caller, data) {
    console.log(data);
    if (data == null || data == undefined || data == "null") {

    } else {// 아직 data가 1건도 없을 수 있음
        var html = ""
        if (data.length > 0) {
            data.forEach(function (vNetItem, vNetIndex) {
                // TODO : 생성 function으로 뺄 것. vnet에 subnet이 2개 이상 있을 수 있는데 그중 1개의 subnet을 선택해야 함.
                var subnetHtml = ""
                var subnetData = vNetItem.subnetInfoList
                var subnetIds = ""
                console.log(subnetData)
                subnetData.forEach(function (subnetItem, subnetIndex) {
                    // subnetHtml +='<input type="hidden" name="vNet_subnet_' + vNetItem.id + '" id="vNet_subnet_' + vNetItem.id + '_' + subnetIndex + '" value="' + subnetItem.iid.nameId + '"/>'
                    //             + subnetIndex + ' || ' + subnetItem.iid.nameId + ' <p>'
                    // console.log(subnetItem)
                    // console.log(subnetItem.iid)
                    subnetHtml += subnetIndex + ' || ' + subnetItem.id + '<p>'
                    if (subnetIndex > 0) { subnetIds += "," }
                    subnetIds += subnetItem.name

                })
                subnetIds += ""
                subnetHtml += '<input type="hidden" name="vNet_subnet_' + vNetItem.id + '" id="vNet_subnet_' + vNetItem.id + '_' + vNetIndex + '" value="' + subnetIds + '"/>'

                console.log("subnetIds = " + subnetIds)

                console.log(subnetHtml)
                html += '<tr onclick="setVnetValueToFormObj(\'es_vNetList\', \'tab_vNet\', \'vNetItem.ID\',\'vNet\',' + vNetIndex + ', \'e_vNetId\');">'

                    + '        <input type="hidden" id="vNet_id_' + vNetIndex + '" value="' + vNetItem.id + '"/>'
                    + '        <input type="hidden" name="vNet_connectionName" id="vNet_connectionName_' + vNetIndex + '" value="' + vNetItem.connectionName + '"/>'
                    + '        <input type="hidden" name="vNet_name" id="vNet_name_' + vNetIndex + '" value="' + vNetItem.name + '"/>'
                    + '        <input type="hidden" name="vNet_description" id="vNet_description_' + vNetIndex + '" value="' + vNetItem.description + '"/>'
                    + '        <input type="hidden" name="vNet_cidrBlock" id="vNet_cidrBlock_' + vNetIndex + '" value="' + vNetItem.cidrBlock + '"/>'
                    + '        <input type="hidden" name="vNet_cspVnetName" id="vNet_cspVnetName_' + vNetIndex + '" value="' + vNetItem.cspVNetName + '"/>'

                    + '        <input type="hidden" name="vNet_subnetInfos" id="vNet_subnetInfos_' + vNetIndex + '" value="' + subnetIds + '"/>'

                    //    사용하지 않는데 굳이 리스트를 할당할 필요가 있을까?
                    //+'        <input type="hidden" name="vNet_keyValueInfos" id="vNet_keyValueInfos_' + vNetIndex + '" value="' + vNetItem.keyValueInfos + '"/>'

                    + '        <input type="hidden" id="vNet_info_' + vNetIndex + '" value="' + vNetItem.id + '|' + vNetItem.name + ' |' + vNetItem.cspVNetName + '|' + vNetItem.cidrBlock + '|' + subnetIds + '"/>'

                    + '    <td class="overlay hidden" data-th="Name">' + vNetItem.name + '</td>'
                    + '    <td class="btn_mtd ovm td_left" data-th="CidrBlock">'
                    + '        ' + vNetItem.cidrBlock
                    + '    </td>'
                    + '    <td class="btn_mtd ovm td_left" data-th="SubnetInfo">' + subnetHtml
                    // +'        { {range $subnetIndex, $subnetItem := .SubnetInfos + ''
                    // +'        <input type="hidden" name="vNet_subnet_' + vNetItem.ID + '" id="vNet_subnet_' + vNetItem.ID + '_' + subnetIndex + '" value="' + subnetItem.IID.NameId + '"/>'
                    // +'        ' + subnetIndex + ' || ' + subnetItem.IID.NameId + ' <p>'
                    // +'        { { end  + ''
                    + '    </td>'
                    + '    <td class="overlay hidden" data-th="Description">' + vNetItem.description + '</td>'
                    + '</tr>'
            })
            $("#e_vNetListTbody").empty()
            $("#e_vNetListTbody").append(html)
        }
    }

}
function getNetworkListCallbackFail(caller, error) {
    // no data
    var html = ""
    html += '<tr>'
        + '<td class="overlay hidden" data-th="" colspan="4">No Data</td>'
        + '</tr>';
    $("#e_vNetListTbody").empty()
    $("#e_vNetListTbody").append(html)
}

function getSpecListCallbackSuccess(caller, data) {
    console.log(data);
    if (data == null || data == undefined || data == "null") {

    } else {// 아직 data가 1건도 없을 수 있음
        var html = ""
        if (data.length > 0) {
            data.forEach(function (vSpecItem, vSpecIndex) {

                html += '<tr onclick="setValueToFormObj(\'es_specList\', \'tab_vmSpec\', \'vmSpec\',' + vSpecIndex + ', \'e_specId\');">'
                    + '     <input type="hidden" id="vmSpec_id_' + vSpecIndex + '" value="' + vSpecItem.id + '"/>'
                    + '     <input type="hidden" name="vmSpec_connectionName" id="vmSpec_connectionName_' + vSpecIndex + '" value="' + vSpecItem.connectionName + '"/>'
                    + '     <input type="hidden" name="vmSpec_info" id="vmSpec_info_' + vSpecIndex + '" value="' + vSpecItem.id + '|' + vSpecItem.name + '|' + vSpecItem.connectionName + '|' + vSpecItem.cspImageId + '|' + vSpecItem.cspImageName + '|' + vSpecItem.guestOS + '|' + vSpecItem.description + '"/>'
                    + '<td class="overlay hidden" data-th="Name">' + vSpecItem.name + '</td>'
                    + '<td class="btn_mtd ovm td_left" data-th="ConnectionName">'
                    + vSpecItem.connectionName
                    + '</td>'
                    + '<td class="overlay hidden" data-th="CspSpecName">' + vSpecItem.cspSpecName + '</td>'

                    + '<td class="overlay hidden" data-th="Description">' + vSpecItem.description + '</td>'
                    + '</tr>'

            })
            $("#e_specListTbody").empty()
            $("#e_specListTbody").append(html)
        }
    }
}
function getSpecListCallbackFail(caller, error) {
    // no data
    var html = ""
    html += '<tr>'
        + '<td class="overlay hidden" data-th="" colspan="4">No Data</td>'
        + '</tr>';
    $("#e_specListTbody").empty()
    $("#e_specListTbody").append(html)
}

function getImageListCallbackSuccess(caller, data) {
    console.log(data);
    if (data == null || data == undefined || data == "null") {

    } else {// 아직 data가 1건도 없을 수 있음
        var html = ""
        if (data.length > 0) {
            data.forEach(function (vImageItem, vImageIndex) {

                html += '<tr onclick="setValueToFormObj(\'es_imageList\', \'tab_vmImage\', \'vmImage\',' + vImageIndex + ', \'e_imageId\');">'
                    + '     <input type="hidden" id="vmImage_id_' + vImageIndex + '" value="' + vImageItem.id + '"/>'
                    + '     <input type="hidden" name="vmImage_connectionName" id="vmImage_connectionName_' + vImageIndex + '" value="' + vImageItem.connectionName + '"/>'
                    + '     <input type="hidden" name="vmImage_info" id="vmImage_info_' + vImageIndex + '" value="' + vImageItem.id + '|' + vImageItem.name + '|' + vImageItem.connectionName + '|' + vImageItem.cspImageId + '|' + vImageItem.cspImageName + '|' + vImageItem.guestOS + '|' + vImageItem.description + '"/>'

                    + '<td class="overlay hidden" data-th="Name">' + vImageItem.name + '</td>'
                    + '<td class="btn_mtd ovm td_left" data-th="ConnectionName">'
                    + vImageItem.connectionName
                    + '</td>'
                    + '<td class="overlay hidden" data-th="CspImageId">' + vImageItem.cspImageId + '</td>'
                    + '<td class="overlay hidden" data-th="CspImageName">' + vImageItem.cspImageName + '</td>'
                    + '<td class="overlay hidden" data-th="GuestOS">' + vImageItem.guestOS + '</td>'
                    + '<td class="overlay hidden" data-th="Description">' + vImageItem.description + '</td>'
                    + '</tr>'
            })
            $("#es_imageListTbody").empty()
            $("#es_imageListTbody").append(html)
        }
    }
}
function getImageListCallbackFail(error) {
    // no data
    var html = ""
    html += '<tr>'
        + '<td class="overlay hidden" data-th="" colspan="4">No Data</td>'
        + '</tr>';
    $("#es_imageListTbody").empty()
    $("#es_imageListTbody").append(html)
}

function getSecurityGroupListCallbackSuccess(caller, data){
    // expert에서 사용할 securityGroup
    if (data == null || data == undefined || data == "null") {

    } else {// 아직 data가 1건도 없을 수 있음
        var html = ""
        if (data.length > 0) {
            data.forEach(function (vSecurityGroupItem, vSecurityGroupIndex) {

                html += '<tr>'

                    + '<td class="overlay hidden column-50px" data-th="">'
                    + '     <input type="checkbox" name="securityGroup_chk" id="securityGroup_Raw_' + vSecurityGroupIndex + '" title="" />'
                    + '     <input type="hidden" id="securityGroup_id_' + vSecurityGroupIndex + '" value="' + vSecurityGroupItem.id + '"/>'
                    + '     <input type="hidden" id="securityGroup_name_' + vSecurityGroupIndex + '" value="' + vSecurityGroupItem.name + '"/>'
                    + '     <input type="hidden" name="securityGroup_connectionName" id="securityGroup_connectionName_' + vSecurityGroupIndex +'" value="' + vSecurityGroupItem.connectionName + '"/>'
                    + '     <input type="hidden" name="securityGroup_info" id="securityGroup_info_' + vSecurityGroupIndex + '" value="'+ vSecurityGroupItem.name +'|' + vSecurityGroupItem.connectionName + '|' + vSecurityGroupItem.description + '"/>'
                    + '     <label for="td_ch1"></label> <span class="ov off"></span>'
                    + '</td>'
                    + '<td class="btn_mtd ovm td_left" data-th="Name">'
                    + vSecurityGroupItem.name
                    + '</td>'
                    + '<td class="btn_mtd ovm td_left" data-th="ConnectionName">'
                    + vSecurityGroupItem.connectionName
                    + '</td>'
                    + '<td class="overlay hidden" data-th="Description">' + vSecurityGroupItem.description + '</td>'

                    + '</tr>'
            })
            $("#e_securityGroupListTbody").empty()
            $("#e_securityGroupListTbody").append(html)

        }
    }
}

function getSecurityGroupListCallbackFail(error){

}

function getSshKeyListCallbackSuccess(caller, data){
    // expert에서 사용할 sshkey
    if (data == null || data == undefined || data == "null") {

    } else {// 아직 data가 1건도 없을 수 있음
        var html = ""
        if (data.length > 0) {
            data.forEach(function (vSshKeyItem, vSshKeyIndex) {

                html += '<tr onclick="setValueToFormObj(\'es_sshKeyList\', \'tab_sshKey\', \'sshKey\', ' + vSshKeyIndex + ', \'e_sshKeyId\');">'

                    + '<td class="overlay hidden" data-th="Name">' + vSshKeyItem.name + '</td>'

                    + '     <input type="hidden" name="sshKey_id" id="sshKey_id_' + vSshKeyIndex + '" value="' + vSshKeyItem.id + '"/>'
                    + '     <input type="hidden" name="sshKey_connectionName" id="sshKey_connectionName_' + vSshKeyIndex + '" value="' + vSshKeyItem.connectionName + '"/>'
                    + '     <input type="hidden" name="sshKey_info" id="sshKey_info_' + vSshKeyIndex + '" value="' + vSshKeyItem.name + '|' + vSshKeyItem.connectionName + '|' + vSshKeyItem.description + '"/>'
                    + '</td>'
                    + '<td class="btn_mtd ovm td_left" data-th="ConnectionName">'
                    + vSshKeyItem.connectionName
                    + '</td>'
                    + '<td class="overlay hidden" data-th="Description">' + vSshKeyItem.description + '</td>'

                    + '</tr>'
            })
            $("#e_sshKeyListTbody").empty()
            $("#e_sshKeyListTbody").append(html)

        }
    }
}

function getSshKeyListCallbackFail(caller, error){

}