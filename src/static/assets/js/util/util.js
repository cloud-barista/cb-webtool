
// div id = Ajax_Loading 이 있어야 함.
// 요청 인터셉터
axios.interceptors.request.use(function (config) {
    // 요청 전에 로딩 오버레이 띄우기
    $('#Ajax_Loading').show();
    return config;
}, function (error) {
    // 에라 나면 로딩 끄기
    $('#Ajax_Loading').hide();
    return Promise.reject(error);
});

// 응답 인터셉터
axios.interceptors.response.use(function (response) {
    // 응답 받으면 로딩 끄기
    $('#Ajax_Loading').hide();
    return response;
}, function (error) {
    // 응답 에러 시에도 로딩 끄기
    $('#Ajax_Loading').hide();
    return Promise.reject(error);
});

function AjaxLoadingShow(isShow){
    try{
        if(isShow) {
            $('#Ajax_Loading').show();
        }else{
            $('#Ajax_Loading').hide();
        }
    }catch(e){
        alert(e);
    }
}
//========== 로딩 바 시작 =========    
// $(document).ready(function(){
//     $('#Ajax_Loading').hide(); //첫 시작시 로딩바를 숨겨준다.
//  })
//  .ajaxStart(function(){
//      $('#Ajax_Loading').show(); //모든 ajax 통신 시작시 로딩바를 보여준다.
//      //$('html').css("cursor", "wait"); //마우스 커서를 로딩 중 커서로 변경
//  })
//  .ajaxStop(function(){
//      $('#Ajax_Loading').hide(); //모든 ajax 통신 종료시 로딩바를 숨겨준다.
//      //$('html').css("cursor", "auto"); //마우스 커서를 원래대로 돌린다
//  });
//========== 로딩 바 종료 =========


// 문자열이 빈 경우 defaultString을 return
function nvl(str, defaultStr){         
    if(typeof str == "undefined" || str == null || str == "")
        str = defaultStr ;
     
    return str ;
}
function nvlDash(str){         
    if(typeof str == "undefined" || str == null || str == "" || str == "undefined")
        str = '-';
     
    return str ;
}

// message를 표현할 alert 창
function commonAlertOpen(alertMessage){
    console.log(alertMessage);
    $('#alertText').text(alertMessage);
    $("#alertArea").modal();
}
// alert창 닫기
function commonAlertClose(){
    $("#alertArea").modal("hide");
}

// confirm modal창 보이기 modal창이 열릴 때 해당 창의 text 지정, close될 때 action 지정
function commonConfirmOpen(targetAction){
    console.log("commonConfirmOpen : " + targetAction)
    // var targetText = "";
    // if( targetAction == "logout"){
    //     targetText = "Would you like to logout?";
    // }else if ( targetAction == "Config"){
    //     targetText = "Would you like to set Cloud config ?";
    // }else if ( targetAction == "SDK"){
    //     targetText = "Would you like to set Cloud Driver SDK ?";
    // }else if ( targetAction == "Credential"){
    //     targetText = "Would you like to set Credential ?";
    // }else if ( targetAction == "Region"){
    //     targetText = "Would you like to set Region ?";
    // }else if ( targetAction == "Provider"){
    //     targetText = "Would you like to set Cloud Provider ?";
    // }else if ( targetAction == "required"){//-- IdPassRequired
    //     targetText = "ID/Password required !";
    // }

    //  [ id , 문구]
    let confirmModalTextMap = new Map(
        [
            ["logout", "Would you like to logout?"],
            ["Config", "Would you like to set Cloud config ?"],
            ["SDK", "Would you like to set Cloud Driver SDK ?"],
            ["Credential", "Would you like to set Credential ?"],
            ["Region", "Would you like to set Region ?"],
            ["Provider", "Would you like to set Cloud Provider ?"],
            // ["IdPassRequired", "ID/Password required !"],    --. 이거는 confirm이 아니잖아
            ["idpwLost", "Illegal account / password 다시 입력 하시겠습니까?"],
            ["ManageNS", "Would you like to manage <br />Name Space?"],
            ["NewNS", "Would you like to add a new Name Space?"],
            ["AddNewNameSpace", "Would you like to register NameSpace <br />Resource ?"],
            ["NameSpace", "Would you like to move <br />selected NameSpace?"],
            ["ChangeNameSpace", "Would you like to move <br />selected NameSpace?"],
            ["DeleteNameSpace", "Would you like to delete <br />selected NameSpace?"],

            ["AddNewVpc", "Would you like to register Network <br />Resource ?"],
            ["DeleteVpc", "Are you sure to delete this Network <br />Resource ?"],

            ["AddNewSecurityGroup", "Would you like to register Security <br />Resource ?"],
            ["DeleteSecurityGroup", "Would you like to un-register Security <br />Resource ?"],
            
            ["AddNewSshKey", "Would you like to register SSH key <br />Resource ?"],
            ["DeleteSshKey", "Would you like to un-register SSH key <br />Resource ?"],     
            
            ["AddNewVirtualMachineImage", "Would you like to register Image <br />Resource ?"],
            ["DeleteVirtualMachineImage", "Would you like to un-register Image <br />Resource ?"],  

            ["AddNewInstanceSpec", "Would you like to register Spec <br />Resource ?"],
            ["DeleteInstanceSpec", "Would you like to un-register Spec <br />Resource ?"],  

            ["AddNewMCIS", "Would you like to create MCIS ?"],
            
        ]
    );
    console.log(confirmModalTextMap.get(targetAction));
    try{
    // $('#modalText').text(targetText);// text아니면 html로 해볼까? 태그있는 문구가 있어서
    //$('#modalText').text(confirmModalTextMap.get(targetAction));
    $('#confirmText').html(confirmModalTextMap.get(targetAction));
    $('#confirmOkAction').val(targetAction);
    
    if( targetAction == "Region"){
        // button에 target 지정
        // data-target="#Add_Region_Register"
        // TODO : confirm 으로 물어본 뒤 OK버튼 클릭 시 targetDIV 지정하도록
    }
    $('#confirmArea').modal(); 
    }catch(e){
        console.log(e);
        alert(e);
    }
}

// confirm modal창에서 ok버튼 클릭시 수행할 method 지정
function commonConfirmOk(){
    //modalArea
    var targetAction = $('#confirmOkAction').val();
    if( targetAction == "logout"){
        // Logout처리하고 index화면으로 간다. Logout ==> cookie expire
    }else if ( targetAction == "Config"){
        //id="Config"
        console.log("commonConfirmOk " + targetAction);
    }else if ( targetAction == "SDK"){
        //id="SDK"
        console.log("commonConfirmOk " + targetAction);
    }else if ( targetAction == "Credential"){
        //id="Credential"
        console.log("commonConfirmOk " + targetAction);
    }else if ( targetAction == "Region"){
        //id="Region"
        console.log("commonConfirmOk " + targetAction);
    }else if ( targetAction == "Provider"){
        //id="Provider"
        console.log("commonConfirmOk " + targetAction);
    }else if ( targetAction == "required"){//-- IdPassRequired
        console.log("commonConfirmOk " + targetAction);
    }else if ( targetAction == "idpwLost"){//-- 
        console.log("commonConfirmOk " + targetAction);
    }else if ( targetAction == "ManageNS"){//-- ManageNS
        console.log("commonConfirmOk " + targetAction);
    }else if ( targetAction == "NewNS"){//-- NewNS
        console.log("commonConfirmOk " + targetAction);
    }else if ( targetAction == "ChangeNameSpace"){//-- ChangeNameSpace
        var changeNameSpaceID = $("#tempSelectedNameSpaceID").val();
        setDefaultNameSpace(changeNameSpaceID)
    }else if ( targetAction == "AddNewNameSpace"){//-- AddNewNameSpace
        displayNameSpaceInfo("REG")
        goFocus('reg_name');
    }else if ( targetAction == "DeleteNameSpace"){
        deleteNameSpace ()
    }else if ( targetAction == "AddNewVpc"){
        displayVNetInfo("REG")
        goFocus('vnetCreateBox');
    }else if ( targetAction == "DeleteVpc"){
        deleteVPC()
    }else if ( targetAction == "AddNewSecurityGroup"){
        displaySecurityGroupInfo("REG")
        goFocus('securityGroupCreateBox');
    }else if ( targetAction == "DeleteSecurityGroup"){
        deleteSecurityGroup()
    }else if ( targetAction == "AddNewSshKey"){
        displaySshKeyInfo("REG")
        goFocus('sshKeyCreateBox');
    }else if ( targetAction == "DeleteSshKey"){
        deleteSshKey()
    }else if ( targetAction == "AddNewVirtualMachineImage"){
        displayVirtualMachineImageInfo("REG")
        goFocus('virtualMachineImageCreateBox');
    }else if ( targetAction == "DeleteVirtualMachineImage"){
        deleteVirtualMachineImage()
    }else if ( targetAction == "AddNewInstanceSpec"){
        displayInstanceSpecInfo("REG")
        goFocus('instanceSpecCreateBox');
    }else if ( targetAction == "DeleteInstanceSpec"){
        deleteInstanceSpec()
    }else if ( targetAction == "AddNewMCIS"){
        createNewMCIS()
    }

       
    
    console.log("commonConfirmOk " + targetAction);
    commonConfirmClose();
}

//confirm modal창에서 cancel 버튼 클릭시 수행할 method 지정. 그냥 창만 듣을 경우에는 commonModalClose() 호출
function commonConfirmCancel(targetAction){
    console.log("commonConfirmCancel : " + targetAction)
    //
    if( targetAction == ''){
        
    }
    commonConfirmClose();
}
// confirm modal창 닫기. setting값 초기화
function commonConfirmClose(){
    $('#confirmText').text('');
    $('#confirmOkAction').val('');
    // $('#modalArea').hide(); 
    $("#confirmArea").modal("hide");
}

// provider에 등록된 connection을 selectbox에 표시
function getConnectionListForSelectbox(provider, targetSelectBoxID){
    
    var data = new Array();
    var url = "/setting/connections/cloudconnectionconfig/" + "list"
    console.log("provider : ",provider)
    var html = "";
    axios.get(url,{
        headers:{
            // 'Authorization': apiInfo
        }
    }).then(result=>{
        console.log('getConnectionConfig result: ',result)
        data = result.data.ConnectionConfig
        console.log("set data array " + data.length);
        
        console.log("connection data : ",data);
        var count = 0; 
        var configName = "";
        var confArr = new Array();
        html +='<option selected>Select Configname</option>';//onChange Event발생을 위해 0번째를 넣음.
        for(var i in data){
            if(provider == data[i].ProviderName){ 
                count++;
                html += '<option value="'+data[i].ConfigName+'" item="'+data[i].ProviderName+'">'+data[i].ConfigName+'</option>';
                configName = data[i].ConfigName
                confArr.push(data[i].ConfigName)                
            }
        }
        if(count == 0){
            commonAlertOpen("해당 Provider에 등록된 Connection 정보가 없습니다.")
            html +='<option selected>Select Configname</option>';
        }
        
        $("#" + targetSelectBoxID).empty();
        $("#" + targetSelectBoxID).append(html);

        if(confArr.length > 1){
            configName = confArr[0];
            console.log("chage value")
            // 0번째 자동으로 선택하여 vNetID목록 갱신
            // $("#" + targetSelectBoxID + " option[value=" + configName + "]").prop('selected', 'selected').change();
            $("#" + targetSelectBoxID + " option[value=" + configName + "]").prop('selected', true).change();         
        }
        // getVnetInfoListForSelectbox(configName);
    }).catch(function(error){
        console.log("Network data error : ",error);        
    });   
}

// connection에 등록된 vnet List를 selectbox에 표시
function getVnetInfoListForSelectbox(configName, targetSelectBoxID){
    console.log("vnet : ", configName);
    
    var url = "/setting/resources" + "/network/list"
    var html = "";
    axios.get(url,{
        headers:{
            // 'Authorization': apiInfo
        }
    }).then(result=>{
        data = result.data.VNetList;
        console.log("vNetwork Info : ",result);
        console.log("vNetwork data : ",data);
        var count = 0; 
        for(var i in data){
            count++;
            if(data[i].connectionName == configName){
                html += '<option value="'+data[i].id+'" selected>'+data[i].cspVNetName+'('+data[i].id+')</option>'; 
            }
        }

        if( count == 0){
            commonAlertOpen("해당 Provider에 등록된 Connection 정보가 없습니다.")
                html +='<option selected>Select Configname</option>';
        }
    
        $("#" + targetSelectBoxID).empty();
        $("#" + targetSelectBoxID).append(html);  
    })
}

function getProviderNameByConnection(configName, targetObjID){
    console.log("configName : ", configName);
    
    var url = "/setting/connections" + "/cloudconnectionconfig/" + configName
    axios.get(url,{
        headers:{
            // 'Authorization': apiInfo
        }
    }).then(result=>{
        data = result.data.ConnectionConfig;
        console.log("connection data : ",data);
        var providerName = data.ProviderName
        console.log("providerName : ",providerName);
        $("#" + targetObjID).val(providerName);
        
    })
}

function isEmpty(str){
	if(typeof str == "undefined" || str == null || str == "")
		return true;
	else
		return false ;
}