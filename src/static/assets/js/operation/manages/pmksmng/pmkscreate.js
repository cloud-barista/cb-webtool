$(document).ready(function() {


})

function btn_deploy(){
// function deploy_btn(){
    var pmksName = $("#pmksreg_name").val();
    if( !validateCloudbaristaKeyName(pmksName, 11) ){
        commonAlert("first letter = small letter <br/> middle letter = small letter, number, hyphen(-) only <br/> last letter = small letter <br/> max length = 11 ");
        return;
    }
    
    var kubernatesNetworkCni = $("#kubernatesNetworkCni").val();
    var kubernatesPodCidr = $("#kubernatesPodCidr").val();
    var kubernatesServiceCidr = $("#kubernatesServiceCidr").val();
    var kubernatesServiceDnsDomain = $("#kubernatesServiceDnsDomain").val();
    
    var controlPlaneLength = $("input[name='controlPlaneCount']").length;
    console.log("controlPlaneLength1 " + controlPlaneLength)
    var controlPlaneConnectionData = new Array(controlPlaneLength);
    var controlPlaneCountData = new Array(controlPlaneLength);
    var controlPlaneSpecIdData = new Array(controlPlaneLength);
    for(var i=0; i<controlPlaneLength; i++){                          
        controlPlaneConnectionData[i] = $("select[name='controlPlaneConnectionName']")[i].value;
        controlPlaneCountData[i] = $("input[name='controlPlaneCount']")[i].value;
        controlPlaneSpecIdData[i] = $("select[name='controlPlaneSpecId']")[i].value;
    }
    console.log(controlPlaneConnectionData)
    console.log(controlPlaneCountData)
    console.log(controlPlaneSpecIdData)
    
    var workerCountLength = $("input[name='workerCount']").length;
    console.log("workerCountLength1 " + workerCountLength)
    var workerConnectionData = new Array();
    var workerCountData = new Array();
    var workerSpecIdData = new Array();
    for(var i=0; i<workerCountLength; i++){      
        var workerId = $("input[name='workerCount']").eq(i).attr("id");
        console.log("workerId " + workerId)
        if( workerId.indexOf("hidden_worker") > -1) continue;// 복사를 위한 영역이 있으므로

        workerConnectionData.push($("select[name='workerConnectionName']")[i].value);
        workerCountData.push($("input[name='workerCount']")[i].value);
        workerSpecIdData.push($("select[name='workerSpecId']")[i].value);
    }
    console.log(workerConnectionData)
    console.log(workerCountData)
    console.log(workerSpecIdData)
    var new_obj = {}
    // mcis 생성이므로 mcisID가 없음
    new_obj['name'] = pmksName
    
    var new_pmksConfig = {}
    var new_kubernetes = {}
    new_kubernetes['networkCni'] = kubernatesNetworkCni;
    new_kubernetes['podCidr'] = kubernatesPodCidr;
    new_kubernetes['serviceCidr'] = kubernatesServiceCidr;
    new_kubernetes['serviceDnsDomain'] = kubernatesServiceDnsDomain;

    new_pmksConfig['kubernetes'] = new_kubernetes;
    new_obj['config'] = new_pmksConfig;
    var controlPlanes = new Array(controlPlaneLength);
    console.log("controlPlaneConnectionLength " + controlPlaneLength)
    for(var i=0; i<controlPlaneLength; i++){
        console.log("controlPlane " + i)
        var new_controlPlane = {}
        new_controlPlane['connection'] = controlPlaneConnectionData[i];
        new_controlPlane['count'] = Number(controlPlaneCountData[i]);
        new_controlPlane['spec'] = controlPlaneSpecIdData[i]
        controlPlanes[i] = new_controlPlane
    }
    new_obj['controlPlane'] = controlPlanes;

    var workers = new Array(workerCountData.length);
    for(var i=0; i<workerCountData.length; i++){
        console.log("workerCountLength " + i)
        var new_worker = {}
        new_worker['connection'] = workerConnectionData[i];
        new_worker['count'] = Number(workerCountData[i]);
        new_worker['spec'] = workerSpecIdData[i]
        workers[i] = new_worker
    }
    new_obj['worker'] = workers;
//     $("input[name='workerCount']").each(function (i) {
//         var new_worker = {}
//         console.log($("select[name='workerConnectionName']").eq(i));
//         new_worker['connection'] = $("select[name='workerConnectionName']").eq(i).attr("value");
//         new_worker['count'] = $("input[name='workerCount']").eq(i).attr("value")
//         new_worker['spec'] = $("select[name='workerSpecId']").eq(i).attr("value")
//         workers[i] = new_worker
//         console.log( i + "번째  : " );
//         console.log( new_worker);
//    });
   
    console.log(new_obj);

    try{
        var url = "/operation/manages/pmksmng/reg/proc"
        axios.post(url,new_obj,{
            headers :{
                },
        }).then(result=>{
            console.log("data : ",result);
            console.log("Result Status : ",result.status); 

            var statusCode = result.data.status;
            var message = result.data.message;
            console.log("Result Status : ",statusCode); 
            console.log("Result message : ",message); 

            if(result.status == 201 || result.status == 200){
                commonResultAlert("PMKS create request success")
            
            }else{
                commonErrorAlert(statusCode, message) 
            }
        }).catch((error) => {
            console.log(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage) 
        })
    }finally{
        
    }
}

// Connection이 선택 되면
// 하위에 해당하는 정보들을 filter해서 가져온다
function changeConnectionInfo(connectionName){
    console.log("connectionName name : ",connectionName)
    if( connectionName == ""){
        // 0번째면 selectbox들을 초기화한다.(vmInfo, sshKey, image 등)
        return
    }
    
    var caller = "pmkscreate";
    var sortType = "name";
    var optionParam = "";
    var filterKey = "connectionName"
    var filterVal = connectionName

    // vpc : filter
    getVnetInfoListForSelectbox(connectionName, 'regVNetId')

    // subnet

    // security group
    getSecurityGroupListForSelectbox(connectionName, 'regSecurityGroup')

    // image, spec, sshkey는 hiddenArea에 set한 다음 NodeGroup Add 버튼 클릭 시 사용
    getCommonVirtualMachineImageList(caller, sortType, optionParam, filterKey, filterVal)
    // image

    // spec
    //getSpecInfo(configName, targetObjId);
    //getCommonVirtualMachineSpecList(caller, sortType, optionParam, filterKey, filterVal)

    // sshkey
}

function changeVpcInfo(vpcID, targetSelectBoxID){
    getSubnetListForSelectbox(vpcID, targetSelectBoxID)
}

// NodeGroup은 Connection이 선택되면 보여준다.
// 0번째는 기본으로 추가한다.
// Connection이 바뀌면 Add 된 내역을 모두 초기화 하므로 confirm을 띄워 확인한다.
function getVmSpecVmImageByConnection(){
    var connectionName = $("#regConnectionName").val();
    var caller = "pmkscreate";
    var sortType = "name";
    var optionParam = "";
    var filterKey = "connectionName"
    var filterVal = connectionName
    // vm pulbic image목록
    getCommonVirtualMachineImageList(caller, sortType, optionParam, filterKey, filterVal)
    // vm private image목록 ??

    // vm spec 목록
    //getCommonVirtualMachineSpecList(caller, sortType, optionParam, filterKey, filterVal)
}

// 모든 커넥션 목록
var totalCloudConnectionList = new Array();
function getCloudConnectionListCallbackSuccess(caller, data, sortType) {
    console.log("connection result: ", data);
    totalCloudConnectionList = data;
}

// public Image 목록 결과
// addNodeGroup을 한 경우 name=nodeGroupPublicImageId가 여러개일 수 있으므로 해당 select를 찾아 모두 set 한다.
function getImageListCallbackSuccess(caller, data, sortType){
    if( caller == "pmkscreate"){
        console.log(data);
        if (data == null || data == undefined || data == "null") {

        } else {// 아직 data가 1건도 없을 수 있음
            var selectImageObjs = $("select[name='nodeGroupPublicImageId']");
            var objLength = selectImageObjs.length
            if( objLength == 1 ){
                objId = selectImageObjs.attr("id");
                // console.log(selectImageObjs)
                // console.log("OBJ ID = " + objId)
                // for(var vmImageIndex in data){
                //     console.log(data[vmImageIndex])
                //     addOptionToSelectObj(objId, data[vmImageIndex].id, data[vmImageIndex].name)
                // }
                setSelectOptions(objId, data, true)
            }else{
                console.log(selectImageObjs.length)
                selectImageObjs.forEach( function( selObj, index){
                    console.log(selObj)
                })
                var publicImageOptions = ""
                publicImageOptions = "<option>Select Image</option>"
                for(var imageIndex in data){
                    publicImageOptions +='<option value="'+data[imageIndex].id+'">'+data[imageIndex].name + '</option>'
                }
                $("#nodeGroupPublicImageId").empty()
                $("#nodeGroupPublicImageId").append(publicImageOptions)
            }            
        }
    }
}

// select 의 option 추가.   isClear = true면 모두 지우고 새로 set
// data object 는 id, name 이라는 field가 반드시 있어야 함.
function setSelectOptions(targetObjID, data, isClear){
    if(isClear){
        $("#" + targetObjID).empty()
        addOptionToSelectObj(targetObjID, "", "Select Image")  
    }
    for(var idx in data){
        console.log(data[idx])
        addOptionToSelectObj(targetObjID, data[idx].id, data[idx].name)
    }
}
function addOptionToSelectObj(targetObjID, optionVal, optionName){
    var addOption ='<option value="'+optionVal+'">'+optionName + '</option>'
    $("#" + targetObjID).append(addOption)
}

function getImageListCallbackFail(caller, error) {
    commonAlert("조회 오류 " + error);
}
// 현재 namespace에 등록된 모든 spec 목록
function getSpecListCallbackSuccess(caller, data, sortType) {
    if( caller == "pmkscreate"){
        console.log(data);
        if (data == null || data == undefined || data == "null") {

        } else {// 아직 data가 1건도 없을 수 있음
            var selectSpecObjs = $("select[name='nodeGroupSpecId']");
            console.log(selectSpecObjs.length)
            selectSpecObjs.forEach( function( selObj, index){
                console.log(selObj)
            })
            var vmSpecOptions = ""
            vmSpecOptions = "<option>Select Spec</option>"
            
            for(var specIndex in data){
                vmSpecOptions +='<option value="'+data[specIndex].id+'">'+data[specIndex].name + '</option>'
            }
            $("#nodeGroupSpecId").empty()
            $("#nodeGroupSpecId").append(vmSpecOptions)
        }
    }

}

function getSpecListCallbackFail(caller, error) {
    commonAlert("조회 오류 " + error);
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
        var addRowCount = 0;
        vmImageList.forEach(function (vImageItem, vImageIndex) {
            console.log(assistConnectionName + " : " + vImageItem.connectionName)
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
                addRowCount++
            }
        });
        if (addRowCount == 0) {
            html = '<tr>'
                + '<td class="overlay hidden" data-th="Name" rowspan="2">Nodata</td>'
                + '</tr>'
        }
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
    } else {
        html += '<tr>'

            + '<td class="overlay hidden" data-th="Name" rowspan="2">Nodata</td>'
            + '</tr>'
        $("#assistVmImageList").empty()
        $("#assistVmImageList").append(html)
    }
}