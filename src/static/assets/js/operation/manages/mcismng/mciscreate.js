$(document).ready(function () {

    // 생성 완료 시 List화면으로 page이동
    $('#alertResultArea').on('hidden.bs.modal', function () {// bootstrap 3 또는 4
        //$('#alertResultArea').on('hidden', function () {// bootstrap 2.3 이전
        let targetUrl = "/operation/manages/mcismng/mngform"
        changePage(targetUrl)
    });

    console.log("mcisCreate.js ")
    getCommonCloudConnectionList('mciscreate', '', true)

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
function displayNewServerForm() {

    var simpleServerConfig = $("#simpleServerConfig");
    var expertServerConfig = $("#expertServerConfig");
    var importServerConfig = $("#importServerConfig");

    if ($("#isImport").is(":checked")) {
        simpleServerConfig.removeClass("active");
        expertServerConfig.removeClass("active");
        importServerConfig.addClass("active");
    } else if ($("#isExpert").is(":checked")) {
        simpleServerConfig.removeClass("active");
        expertServerConfig.addClass("active");
        importServerConfig.removeClass("active");
    } else {
        //simpleServerConfig        
        simpleServerConfig.addClass("active");
        expertServerConfig.removeClass("active");
        importServerConfig.removeClass("active");
    }
}



// 서버정보 입력 area에서 'DONE'버튼 클릭시 array에 담고 form을 초기화

var TotalServerConfigArr = new Array();
// deploy 버튼 클릭시 등록한 서버목록을 배포.
// function btn_deploy(){
function deployMcis() {
    var mcis_name = $("#mcis_name").val();
    if (!mcis_name) {
        commonAlert("Please Input MCIS Name!!!!!")
        return;
    }
    var mcis_desc = $("#mcis_desc").val();
    var placement_algo = $("#placement_algo").val();
    var installMonAgent = $("#installMonAgent").val();

    var new_obj = {}

    var vm_len = 0;

    if (IsImport) {
        // ImportedMcisScript.name = mcis_name;
        // ImportedMcisScript.description = mcis_desc;
        // ImportedMcisScript.installMonAgent = installMonAgent;
        // console.log(ImportedMcisScript);
        //var theJson = jQuery.parseJSON($(this).val())
        //$("#mcisImportScriptPretty").val(fmt);	
        new_obj = $("#mcisImportScriptPretty").val();
        new_obj.id = "";// id는 비워준다.
    } else {
        //         console.log(Simple_Server_Config_Arr)

        // mcis 생성이므로 mcisID가 없음
        new_obj['name'] = mcis_name
        new_obj['description'] = mcis_desc
        new_obj['installMonAgent'] = installMonAgent

        if (Simple_Server_Config_Arr) {
            vm_len = Simple_Server_Config_Arr.length;
            for (var i in Simple_Server_Config_Arr) {
                TotalServerConfigArr.push(Simple_Server_Config_Arr[i]);
            }
        }

        if (Expert_Server_Config_Arr) {
            vm_len = Expert_Server_Config_Arr.length;
            for (var i in Expert_Server_Config_Arr) {
                TotalServerConfigArr.push(Expert_Server_Config_Arr[i]);
            }
        }

        if (TotalServerConfigArr) {
            vm_len = TotalServerConfigArr.length;
            console.log("Server_Config_Arr length: ", vm_len);
            new_obj['vm'] = TotalServerConfigArr;
            console.log("new obj is : ", new_obj);
        } else {
            commonAlert("Please Input Servers");
            $(".simple_servers_config").addClass("active");
            $("#s_name").focus();
        }
    }

    // var url = CommonURL+"/ns/"+NAMESPACE+"/mcis";
    var url = "/operation/manages/mcismng/reg/proc"
    try {
        axios.post(url, new_obj, {
            headers: {
                'Content-type': 'application/json',
            },
        }).then(result => {
            console.log("MCIR Register data : ", result);
            console.log("Result Status : ", result.status);
            if (result.status == 201 || result.status == 200) {
                commonResultAlert("Register Success")
                // var targetUrl = "/operation/manages/mcismng/mngform"
                // changePage(targetUrl)
            } else {
                commonAlert("Register Fail")
                //location.reload(true);
            }
        }).catch((error) => {
            // console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage)

        })
    } catch (error) {
        commonAlert(error);
        console.log(error);
    }
}

// MCIS Create 와 VM Create의 function이름이 같음
function displayMcisImportServerFormByImport(importType) {
    // var $SimpleServers = $("#simpleServerConfig");
    // var $ExpertServers = $("#expertServerConfig");
    // var $ImportServers = $("#importServerConfig");
    // var check = $(".switch .ch").is(":checked");
    // console.log("check=" + check);
    // if( check){
    //     $SimpleServers.removeClass("active");
    //     $ExpertServers.removeClass("active");            
    //     $ImportServers.addClass("active");

    //     importMCISInfoFromFile();// import창 띄우기 
    // }

    var addMcisByScriptArea = $("#addMcisByScript");
    var addMcisByScriptBtn = $("#addMcisCancel");
    var addVmListArea = $("#addVmList");
    var mcisInfoboxArea = $("#mcisInfobox");

    if (importType) {
        addMcisByScriptArea.css("display", "block");
        addMcisByScriptBtn.css("display", "inline");
        addVmListArea.css("display", "none");
        mcisInfoboxArea.css("display", "none");
        // addMcisByScriptArea.addClass("active");
        // addVmListArea.removeClass("active");

        importMCISInfoFromFile();// import창 띄우기 
    } else {
        $("#mcisImportScriptPretty").val("");
        addMcisByScriptArea.css("display", "none");
        addMcisByScriptBtn.css("display", "none");
        addVmListArea.css("display", "block");
        mcisInfoboxArea.css("display", "block");
    }
    IsImport = importType;// 전역으로 set
}

// mcis export한 파일 선택하여 읽기
function importMCISInfoFromFile() {
    var input = document.createElement("input");
    input.type = "file";
    // input.accept = "text/plain"; // 확장자가 xxx, yyy 일때, ".xxx, .yyy"
    input.accept = ".json";
    input.onchange = function (event) {
        importMcisFileProcess(event.target.files[0]);
    };
    input.click();
}

// 선택한 MCIS 파일을 읽어 화면에 보여줌
var ImportedMcisScript = "";
var IsImport = false;
function importMcisFileProcess(file) {
    try {
        var reader = new FileReader();
        reader.onload = function () {
            console.log(reader.result);
            console.log("---1")
            // $("#fileContent").val(reader.result);

            var jsonStr = JSON.stringify(reader.result)
            console.log(JSON.stringify(jsonStr));

            ImportedMcisScript = $.parseJSON(reader.result);

            setMcisInfoToForm(ImportedMcisScript);
            mcisTojsonFormatter(ImportedMcisScript)

        };
        //reader.readAsText(file, /* optional */ "euc-kr");
        reader.readAsText(file);
    } catch (error) {
        commonAlert("File Load Failed");
        console.log(error);
    }
}

// json 객체를 textarea에 표시할 때 예쁘게
function mcisTojsonFormatter(mcisInfoObj) {
    var fmt = JSON.stringify(mcisInfoObj, null, 4);    // stringify with 4 spaces at each level
    $("#mcisImportScriptPretty").val(fmt);
}

// 화면의 mcis 의 값이 변경되면 적용
function setMcisValue(mcisObjId, mcisObjValue) {
    if (IsImport) {
        if (mcisObjId == "mcis_name") {
            ImportedMcisScript.name = mcisObjValue;
        } else if (mcisObjId == "mcis_desc") {
            ImportedMcisScript.description = mcisObjValue;
        } else if (mcisObjId == "installMonAgent") {
            ImportedMcisScript.installMonAgent = mcisObjValue;
        }
        mcisTojsonFormatter(ImportedMcisScript)
    }
}
// 선택한 파일을 읽어 form에 Set
function setMcisInfoToForm(mcisInfoObj) {
    //export form
    $("#mcis_name").val(mcisInfoObj.name);
    $("#mcis_desc").val(mcisInfoObj.description);
    // $("#label").val(mcisInfoObj.label);
    $("#installMonAgent").val(mcisInfoObj.installMonAgent);

    // // 수신한 obj를 바로 deploy로 던질까?
    // var url = "/operation/manages/mcismng/reg/proc"
    // try{        
    //     axios.post(url,mcisInfoObj,{
    //         headers :{
    //             'Content-type': 'application/json',
    //             // 'Authorization': apiInfo,
    //             },
    //     }).then(result=>{
    //         console.log("MCIR Register data : ",result);
    //         console.log("Result Status : ",result.status); 
    //         if(result.status == 201 || result.status == 200){
    //             commonAlert("Register Success")
    //             // location.href = "/Manage/MCIS/list";
    //             // $('#loadingContainer').show();
    //             // location.href = "/operation/manages/mcismng/mngform/"
    //             var targetUrl = "/operation/manages/mcismng/mngform"
    //             changePage(targetUrl)
    //         }else{
    //             commonAlert("Register Fail")
    //             //location.reload(true);
    //         }
    //     }).catch((error) => {
    //         // console.warn(error);
    //         console.log(error.response)
    //         var errorMessage = error.response.data.error;
    //         var statusCode = error.response.status;
    //         commonErrorAlert(statusCode, errorMessage) 

    //     })
    // }finally{
    //     // AjaxLoadingShow(false);
    // }  


}


var totalCloudConnectionList = new Array();
function getCloudConnectionListCallbackSuccess(caller, data, sortType) {
    totalCloudConnectionList = data;
}
// target Object(selectbox) 에 해당 provider목록만 표시
// mcisvmcreate.js에도 동일한 function이 있으므로 추후 통합해야 함.
function filterConnectionByProvider(provider, targetObjId) {
    $('#' + targetObjId).find('option').remove();
    $('#' + targetObjId).append('<option value="">Selected Connection</option>')
    for (var connIndex in totalCloudConnectionList) {
        var aConnection = totalCloudConnectionList[connIndex];
        if (provider == "" || provider == aConnection.ProviderName) {
            $('#' + targetObjId).append('<option value="' + aConnection.ConfigName + '">' + aConnection.ConfigName + '</option>')
        }
    }
}

// 등록 된 vm search 결과
function getCommonSearchVmImageListCallbackSuccess(caller, vmImageList) {
    console.log(vmImageList);
    var html = ""
    if (vmImageList.length > 0) {
        // if( caller == "imageAssist" ){
        // 조회 조건으로 provider, connection이 선택되어 있으면 조회 후 filter
        var assistProviderName = $("#assistImageProviderName option:selected").val();
        var assistConnectionName = $("#assistImageConnectionName option:selected").val();
        console.log("getCommonSearchVmImageListCallbackSuccess")
        vmImageList.forEach(function (vImageItem, vImageIndex) {
            //console.log(assistConnectionName + " : " + vImageItem.connectionName)
            console.log(vImageItem)
            if (assistConnectionName == "" || assistConnectionName == vImageItem.connectionName) {
                //connectionName
                //cspSpecName
                html += '<tr onclick="setAssistValue(' + vImageIndex + ');">'
                    + '     <input type="hidden" id="vmImageAssist_id_' + vImageIndex + '" value="' + vImageItem.id + '"/>'
                    + '     <input type="hidden" id="vmImageAssist_name_' + vImageIndex + '" value="' + vImageItem.name + '"/>'
                    + '     <input type="hidden" id="vmImageAssist_connectionName_' + vImageIndex + '" value="' + vImageItem.connectionName + '"/>'
                    + '     <input type="hidden" id="vmImageAssist_cspImageId_' + vImageIndex + '" value="' + vImageItem.cspImageId + '"/>'
                    + '     <input type="hidden" id="vmImageAssist_cspImageName_' + vImageIndex + '" value="' + vImageItem.cspImageName + '"/>'
                    + '     <input type="hidden" id="vmImageAssist_guestOS_' + vImageIndex + '" value="' + vImageItem.guestOS + '"/>'
                    + '     <input type="hidden" id="vmImageAssist_description_' + vImageIndex + '" value="' + vImageItem.description + '"/>'
                    + '<td class="overlay hidden" data-th="Name">' + vImageItem.name + '</td>'
                    + '<td class="overlay hidden" data-th="CspImageName">' + vImageItem.cspImageName + '</td>'
                    + '<td class="overlay hidden" data-th="CspImageId">' + vImageItem.cspImageId + '</td>'

                    // + '<td class="overlay hidden" data-th="GuestOS">' + vImageItem.guestOS + '</td>'
                    // + '<td class="overlay hidden" data-th="Description">' + vImageItem.description + '</td>'
                    + '</tr>'
            }
        });
        $("#assistVmImageList").empty()
        $("#assistVmImageList").append(html)

        $("#assistVmImageList tr").each(function () {
            $selector = $(this)

            $selector.on("click", function () {

                if ($(this).hasClass("on")) {
                    $(this).removeClass("on");
                } else {
                    $(this).addClass("on")
                    $(this).siblings().removeClass("on");
                }
            })
        })
    }
}

// 등록된 spec조회 성공 시 table에 뿌려주고, 클릭시 spec 내용 set.
function getCommonFilterVmSpecListCallbackSuccess(caller, vmSpecList) {
    // function getCommonFilterVmImageListCallbackSuccess(caller, vmSpecList){
    console.log(vmSpecList);
    if (vmSpecList == null || vmSpecList == undefined || vmSpecList == "null") {

    } else {// 아직 data가 1건도 없을 수 있음
        var html = ""
        if (vmSpecList.length > 0) {
            vmSpecList.forEach(function (vSpecItem, vSpecIndex) {
                //connectionName
                //cspSpecName
                html += '<tr onclick="setAssistValue(' + vSpecIndex + ');">'
                    + '     <input type="hidden" id="vmSpecAssist_id_' + vSpecIndex + '" value="' + vSpecItem.id + '"/>'
                    + '     <input type="hidden" id="vmSpecAssist_name_' + vSpecIndex + '" value="' + vSpecItem.name + '"/>'
                    + '     <input type="hidden" id="vmSpecAssist_cspSpecName_' + vSpecIndex + '" value="' + vSpecItem.cspSpecName + '"/>'
                    + '     <input type="hidden" id="vmSpecAssist_memGiB_' + vSpecIndex + '" value="' + vSpecItem.memGiB + '"/>'
                    + '     <input type="hidden" id="vmSpecAssist_numvCPU_' + vSpecIndex + '" value="' + vSpecItem.numvCPU + '"/>'
                    + '     <input type="hidden" id="vmSpecAssist_numGpu_' + vSpecIndex + '" value="' + vSpecItem.numGpu + '"/>'
                    + '     <input type="hidden" id="vmSpecAssist_connectionName_' + vSpecIndex + '" value="' + vSpecItem.connectionName + '"/>'
                    + '<td class="overlay hidden" data-th="Name">' + vSpecItem.name + '</td>'
                    + '<td class="overlay hidden" data-th="CspSpecName">' + vSpecItem.cspSpecName + '</td>'
                    + '<td class="btn_mtd ovm td_left" data-th="Memory">'
                    + vSpecItem.memGiB
                    + '</td>'
                    + '<td class="overlay hidden" data-th="VCPU">' + vSpecItem.numvCPU + '</td>'

                    + '<td class="overlay hidden" data-th="GPU">' + vSpecItem.numGpu + '</td>'
                    + '</tr>'
            })
            $("#assistSpecList").empty()
            $("#assistSpecList").append(html)

            $("#assistSpecList tr").each(function () {
                $selector = $(this)

                $selector.on("click", function () {

                    if ($(this).hasClass("on")) {
                        $(this).removeClass("on");
                    } else {
                        $(this).addClass("on")
                        $(this).siblings().removeClass("on");
                    }
                })
            })
        }

        // "associatedObjectList": null,
        // "connectionName": "aws-conn-osaka",
        // "costPerHour": 0,
        // "cspSpecName": "t3.small",
        // "description": "",
        // "ebsBwMbps": 0,
        // "evaluationScore01": 0,
        // "evaluationScore02": 0,
        // "evaluationScore03": 0,
        // "evaluationScore04": 0,
        // "evaluationScore05": 0,
        // "evaluationScore06": 0,
        // "evaluationScore07": 0,
        // "evaluationScore08": 0,
        // "evaluationScore09": 0,
        // "evaluationScore10": 0,
        // "evaluationStatus": "",
        // "gpuMemGiB": 0,
        // "gpuModel": "",
        // "gpuP2p": "",
        // "id": "osaka-t3small",
        // "isAutoGenerated": false,
        // "maxNumStorage": 0,
        // "maxTotalStorageTiB": 0,
        // "memGiB": 2,
        // "name": "osaka-t3small",
        // "namespace": "osaka-ns",
        // "netBwGbps": 0,
        // "numCore": 0,
        // "numGpu": 0,
        // "numStorage": 0,
        // "numvCPU": 2,
        // "orderInFilteredResult": 0,
        // "osType": "",
        // "storageGiB": 0


    }
}
// $(document).ready(function() {
//     //OS_HW popup table scrollbar
// //   $('#OS_HW .btn_spec').on('click', function() {
// //         $('#OS_HW_Spec .dtbox.scrollbar-inner').scrollbar();
// //     });
// //     //Security popup table scrollbar
// //   $('#Security .btn_edit').on('click', function() {
// //     $("#security_edit").modal();
// //         $('#security_edit .dtbox.scrollbar-inner').scrollbar();
// //     });
// });
