// 원래는 confirm창이었으나 입력form에 물어볼 필요 없으므로 config Area만 show
function addMonitoringAlertPolicy(){
    console.log("##########AddMonitoringAlertPolicy")
    $(".dashboard.register_cont").toggleClass("active");
    $(".dashboard.server_status").removeClass("view");
    $(".dashboard .status_list tbody tr").removeClass("on");
    //ok 위치이동
    // $('#RegistBox').on('hidden.bs.modal', function () {
    var offset = $("#CreateBox").offset();
    $("#wrap").animate({scrollTop : offset.top}, 300);        
    // })

    // 등록 폼 초기화
    $("#regMonitoringAlertName").val('');				 
    $("#regMonitoringAlertMeasure").val('');            
    $("#regMonitoringAlertTargetType").val('');         
    $("#regMonitoringAlertTargetID").val('');           
    $("#regMonitoringAlertEventDuration").val('');      
    $("#regMonitoringAlertMetric").val('');             
    $("#regMonitoringAlertAlertMathExpression").val('');
    $("#regMonitoringAlertAlertThreshold").val('');     
    $("#regMonitoringAlertWarnEventCount").val('');     
    $("#regMonitoringAlertCriticEventCount").val('');  
}


function deleteMonitoringAlertPolicy(){
    console.log("##########DeleteMonitoringAlertPolicy")

    var monitoringAlertId = "";
    var count = 0;

    $( "input[name='chk']:checked" ).each( function () {
        count++;
        monitoringAlertId = monitoringAlertId + $(this).val()+",";
    });

    monitoringAlertId = monitoringAlertId.substring(0,monitoringAlertId.lastIndexOf( ","));
    
    console.log("monitoringAlertId : ", monitoringAlertId);
    console.log("count : ", count);

    if(monitoringAlertId == ''){
        commonAlert("삭제할 대상을 선택하세요.");
        return false;
    }

    if(count != 1){
        commonAlert("삭제할 대상을 하나만 선택하세요.");
        return false;
    }

    var url = "/operation/policies/monitoringalertpolicy/del/" + monitoringAlertId
    console.log("del monitoringAlertPolicy url : ", url);

    axios.delete(url, {
        headers: {
            // 'Authorization': "{{ .apiInfo}}",
            'Content-Type': "application/json"
        }
    }).then(result => {
        var data = result.data;
        console.log(data);
        if (result.status == 200 || result.status == 201) {
           commonAlert("Success Delete Threshold")
           displayMonitoringAlertPolicyInfo("DEL_SUCCESS")
        }else{
            commonAlert(data)
        }
    }).catch(function(error){
        commonAlert(error)
        console.log("Threshold delete error : ",error);        
    });
}


function createMonitoringAlertPolicy(){
    console.log("##########CreateMonitoringAlertPolicy")

    var monitoringAlertName = $("#regMonitoringAlertName").val();				 
    var monitoringAlertMeasure = $("#regMonitoringAlertMeasure").val();            
    var monitoringAlertTargetType = $("#regMonitoringAlertTargetType").val();         
    var monitoringAlertTargetID = $("#regMonitoringAlertTargetID").val();           
    var monitoringAlertEventDuration = $("#regMonitoringAlertEventDuration").val();      
    var monitoringAlertMetric = $("#regMonitoringAlertMetric").val();             
    var monitoringAlertAlertMathExpression = $("#regMonitoringAlertAlertMathExpression").val();
    var monitoringAlertAlertThreshold = $("#regMonitoringAlertAlertThreshold").val();     
    var monitoringAlertWarnEventCount = $("#regMonitoringAlertWarnEventCount").val();     
    var monitoringAlertCriticEventCount = $("#regMonitoringAlertCriticEventCount").val(); 

    if(!monitoringAlertName) {
        commonAlert("Input New Threshold Name")
        $("#regMonitoringAlertName").focus()
        return;
    }

    var url = "/operation/policies/monitoringalertpolicy/reg/proc"
    console.log("Threshold Reg URL : ", url)
    var obj = {
        name                    : monitoringAlertName,				 
        measurement             : monitoringAlertMeasure,            
        target_type             : monitoringAlertTargetType,         
        target_id               : monitoringAlertTargetID,         
        event_duration          : monitoringAlertEventDuration,      
        metric                  : monitoringAlertMetric,             
        alert_math_expression   : monitoringAlertAlertMathExpression,
        alert_threshold         : Number(monitoringAlertAlertThreshold),     
        warn_event_cnt          : Number(monitoringAlertWarnEventCount),     
        critic_event_cnt        : Number(monitoringAlertCriticEventCount),
        alert_event_type        : "slack",
        alert_event_name        : "slackHandler",
        alert_event_message     : monitoringAlertName
    }

    console.log("info Threshold obj Data : ", obj);

    if(monitoringAlertName) {
        axios.post(url, obj, {
            headers: {
                'Content-type': 'application/json',
                // 'Authorization': apiInfo,
            }
        }).then(result => {
            console.log("result Threshold : ", result);
            var data = result.data;
            console.log(data);

            if (data.status == 200 || data.status == 201) {
                commonAlert("Success Create Threshold!!")
                
                displayMonitoringAlertPolicyInfo("REG_SUCCESS")
            } else {
                commonAlert("Fail Create Threshold " + data.message)
            }
        }).catch(function(error){
            var data = error.data;
            console.log(data);
            console.log(error);        
            commonAlert("Threshold create error : ",error)            
        });
    } else {
        commonAlert("Input Threshold Name")
        $("#regMonitoringAlertName").focus()
        return;
    }
}


function showMonitoringAlertPolicyInfo(monitoringAlertName) {
    console.log("showMonitoringAlertPolicyInfo : ", monitoringAlertName);

    $('#thresholdName').text(monitoringAlertName)

    var url = "/operation/policies/monitoringalertpolicy/" + encodeURIComponent(monitoringAlertName);
    console.log("Threshold detail URL : ",url)

    return axios.get(url,{
        // headers:{
        //     'Authorization': apiInfo
        // }
    }).then(result=>{
        console.log(result);
        console.log(result.data);
        var data = result.data.MonitoringAlertPolicyInfo
        console.log("Show Data : ",data);
        
        var dtlMonitoringAlertName				   = data.name
        var dtlMonitoringAlertMeasure              = data.measurement
        var dtlMonitoringAlertTargetType           = data.target_type
        var dtlMonitoringAlertTargetID             = data.target_id
        var dtlMonitoringAlertEventDuration        = data.event_duration
        var dtlMonitoringAlertMetric               = data.metric
        var dtlMonitoringAlertAlertMathExpression  = data.alert_math_expression
        var dtlMonitoringAlertAlertThreshold       = data.alert_threshold
        var dtlMonitoringAlertWarnEventCount       = data.warn_event_cnt
        var dtlMonitoringAlertCriticEventCount     = data.critic_event_cnt

        $("#dtlMonitoringAlertName").empty();				 
        $("#dtlMonitoringAlertMeasure").empty();            
        $("#dtlMonitoringAlertTargetType").empty();         
        $("#dtlMonitoringAlertTargetID").empty();           
        $("#dtlMonitoringAlertEventDuration").empty();      
        $("#dtlMonitoringAlertMetric").empty();             
        $("#dtlMonitoringAlertAlertMathExpression").empty();
        $("#dtlMonitoringAlertAlertThreshold").empty();     
        $("#dtlMonitoringAlertWarnEventCount").empty();     
        $("#dtlMonitoringAlertCriticEventCount").empty();

        $("#dtlMonitoringAlertName").val(dtlMonitoringAlertName);				 
        $("#dtlMonitoringAlertMeasure").val(dtlMonitoringAlertMeasure);               
        $("#dtlMonitoringAlertTargetType").val(dtlMonitoringAlertTargetType);         
        $("#dtlMonitoringAlertTargetID").val(dtlMonitoringAlertTargetID);           
        $("#dtlMonitoringAlertEventDuration").val(dtlMonitoringAlertEventDuration);      
        $("#dtlMonitoringAlertMetric").val(dtlMonitoringAlertMetric);             
        $("#dtlMonitoringAlertAlertMathExpression").val(dtlMonitoringAlertAlertMathExpression);
        $("#dtlMonitoringAlertAlertThreshold").val(dtlMonitoringAlertAlertThreshold);     
        $("#dtlMonitoringAlertWarnEventCount").val(dtlMonitoringAlertWarnEventCount);     
        $("#dtlMonitoringAlertCriticEventCount").val(dtlMonitoringAlertCriticEventCount);
        
    }) .catch(function(error){
        console.log("Threshold detail error : ",error);        
    });
    
}


function displayMonitoringAlertPolicyInfo(targetAction) {
    if( targetAction == "REG_SUCCESS" ) {
        console.log("##########AddMonitoringAlertPolicy REG_SUCCESS")
        $(".dashboard.register_cont").removeClass("active");
        $(".dashboard.server_status").removeClass("view");
        $(".dashboard .status_list tbody tr").addClass("on");
        
        var offset = $("#CreateBox").offset();
        $("#wrap").animate({scrollTop : offset.top}, 0);        
        
        // 등록 폼 초기화
        $("#regMonitoringAlertName").val('');				 
        $("#regMonitoringAlertMeasure").val('');            
        $("#regMonitoringAlertTargetType").val('');         
        $("#regMonitoringAlertTargetID").val('');           
        $("#regMonitoringAlertEventDuration").val('');      
        $("#regMonitoringAlertMetric").val('');             
        $("#regMonitoringAlertAlertMathExpression").val('');
        $("#regMonitoringAlertAlertThreshold").val('');     
        $("#regMonitoringAlertWarnEventCount").val('');     
        $("#regMonitoringAlertCriticEventCount").val('');  
        
        getMonitoringAlertPolicyList("alertName");
    } else if ( targetAction == "DEL_SUCCESS" ) {
        console.log("##########AddMonitoringAlertPolicy DEL_SUCCESS")
        $(".dashboard.register_cont").removeClass("active");
        $(".dashboard.server_status").removeClass("view");
        $(".dashboard .status_list tbody tr").addClass("on");

        var offset = $("#CreateBox").offset();
        $("#wrap").animate({scrollTop : offset.top}, 0);

        getMonitoringAlertPolicyList("alertName");
    }
}


function getMonitoringAlertPolicyList(sortType) {
    console.log("#####################getMonitoringAlertPolicyList : ", sortType);

    var url = "/operation/policies/monitoringalertpolicy/list"
    axios.get(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get Threshold List : ", result.data);
        var data = result.data.MonitoringAlertPolicyList;
        console.log("$$$Alert DATA$$$");
        console.log(data);

        var html = ""
        var cnt = 0;

        if (data.length) {
            if (sortType) {
                cnt++;
                console.log("check2 : ", sortType);
                data.filter(list => list.Name !== "").sort((a, b) => (a[sortType] < b[sortType] ? - 1 : a[sortType] > b[sortType] ? 1 : 0)).map((item, index) => (
                    html += addMonitoringAlertRow(item, index)
                ))
            } else {
                console.log("check3 : ", sortType);
                data.filter((list) => list.Name !== "").map((item, index) => (
                    html += addMonitoringAlertRow(item, index)
                ))
            }

            $("#alertList").empty();
            $("#alertList").append(html)

            ModalDetail();
        }
    }).catch(function(error){
        console.log("Threshold list error : ", error);        
    })    
}


// Threshold목록에 Item 추가
function addMonitoringAlertRow(item, index){
    console.log("addMonitoringAlertRow " + index);
    console.log(item)
    var html = ""
    html += '<tr onclick="showMonitoringAlertPolicyInfo(\'' + item.name + '\');">'
        + '<td class="overlay hidden" data-th="">'
        + '<input type="hidden" id="alertpolicy_info_' + index + '" value="' + item.name + '"/>'
        + '<input type="checkbox" name="chk" value="' + item.name + '" id="raw_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>'
        + '<td class="overlay hidden" data-th="Name">' + item.name + '</td>'
        + '<td class="overlay hidden" data-th="Measurement">' + item.measurement + '</td>'
        + '<td class="overlay hidden" data-th="TargetType">' + item.target_type + '</td>'
        + '<td class="overlay hidden" data-th="TargetId">' + item.target_id + '</td>'
        + '<td class="overlay hidden" data-th="AlertEventType">' + item.alert_event_type + '</td>'
        + '<td class="overlay hidden" data-th="AlertEventName">' + item.alert_event_name + '</td>'
        + '<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>'
        + '</tr>'	
    return html;
}


function ModalDetail() {
    $(".dashboard .status_list tbody tr").each(function () {
        var $td_list = $(this),
            $status = $(".server_status"),
            $detail = $(".server_info");
        $td_list.off("click").click(function () {
            $td_list.addClass("on");
            $td_list.siblings().removeClass("on");
            $status.addClass("view");
            $status.siblings().removeClass("on");
            $(".dashboard.register_cont").removeClass("active");
            $td_list.off("click").click(function () {
                if ($(this).hasClass("on")) {
                    console.log("reg ok button click")
                    $td_list.removeClass("on");
                    $status.removeClass("view");
                    $detail.removeClass("active");
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


// Add Monitoring Alert Event-Handler 
function addMonitoringAlertEventHandler() {
    $("#Add_MonitoringAlertEventHandler_Register").modal(); 
}


// Save Monitoring Alert Event-Handler 
function saveNewMonitoringAlertEventHandler() {
    // valid check
    var monitoringAlertEventHandlerModalType = $("#regMonitoringAlertEventHandlerModalType").val();
    var monitoringAlertEventHandlerModalName = $("#regMonitoringAlertEventHandlerModalName").val();

    console.log(monitoringAlertEventHandlerModalType + ", " + monitoringAlertEventHandlerModalName);

    // if(!monitoringAlertEventHandlerModalType || !monitoringAlertEventHandlerModalName){
    //     alert("!!!!");
    //     $("#modalMonitoringAlertEventHandlerRequired").modal()// TODO : requiredCloudConnection 로 바꿔 공통으로 쓸까?
    //     return;
    // }

    // var optionsVal = {};
    // optionsVal['url'] = "https://cloud-barista.slack.com/archives/C022PB8K7NG";
    // optionsVal['channel'] = "#monitoring-alert-event-handler";
    // var monitoringAlertEventHandlerInfo = {         
    //     type: monitoringAlertEventHandlerModalType,
    //     name: monitoringAlertEventHandlerModalName,
    //     options: optionsVal           
    // }

    var monitoringAlertEventHandlerInfo = {         
        type: monitoringAlertEventHandlerModalType,
        name: monitoringAlertEventHandlerModalName,
        url: "https://hooks.slack.com/services/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",
        channel: "#kapacitor-alert"
    }

    console.log(monitoringAlertEventHandlerInfo);

    var url = "/operation/policies/monitoringalerteventhandler/reg/proc";
    console.log("Monitoring Alert Event-Handler Reg URL : ",url)
    
    if(monitoringAlertEventHandlerModalType || monitoringAlertEventHandlerModalName) {
        axios.post(url, monitoringAlertEventHandlerInfo, {
            headers: {
                'Content-type': 'application/json',
                // 'Authorization': apiInfo,
            }
        }).then(result => {
            console.log("result add monitoring alert event handler : ", result);
            var data = result.data;
            console.log(data);

            if (data.status == 200 || data.status == 201) {
                commonAlert("Success Create Monitoring Alert Event-Handler!!")
                
                displayMonitoringAlertEventHandlerInfo("REG_SUCCESS")
            } else {
                commonAlert("Fail Create Monitoring Alert Event-Handler " + data.message)
            }
        }).catch(function(error){
            var data = error.data;
            console.log(data);
            console.log(error);        
            commonAlert("Monitoring Alert Event-Handler create error : ",error)            
        });
    } else {
        commonAlert("Input Monitoring Alert Event-Handler Type or Name");
        $("#regMonitoringAlertEventHandlerModalType").focus()
        return;
    }

    if (monitoringAlertEventHandlerModalType == "smtp") {

    } else if (monitoringAlertEventHandlerModalType == "slack"){

    }

}


function deleteMonitoringAlertEventHandler() {
    console.log("##########deleteMonitoringAlertEventHandler")

    var count = 0;

    var selectedIndex = "";
    var selectedType = "";
    var selectedName = "";

    // var chkIdArr = $(this).attr('id').split("_");// 0번째와 2번째를 합치면 id 추출가능  ex) securityGroup_Raw_0
    //   if( $(this).is(":checked")){
    $( "input[name='chk']:checked" ).each( function () {
        count++;
        selectedIndex = $(this).attr('id').split("_")[1]; // raw_1  [0] = raw, [1] = 1
        selectedType = $("#monitoringAlertEventHandlerType_info_" + selectedIndex).val();
        selectedName = $("#monitoringAlertEventHandlerName_info_" + selectedIndex).val();
    });

    console.log("selectedType : ", selectedType);
    console.log("selectedName : ", selectedName);
    console.log("count : ", count);

    if(selectedIndex == ''){
        commonAlert("삭제할 대상을 선택하세요.");
        return false;
    }

    if(count != 1){
        commonAlert("삭제할 대상을 하나만 선택하세요.");
        return false;
    }
    
    var url = "/operation/policies/monitoringalerteventhandler/del/type/" + selectedType + "/event/" + selectedName
    console.log("del monitoringAlertEventHandler url : ", url);

    axios.delete(url, {
        headers: {
            // 'Authorization': "{{ .apiInfo}}",
            'Content-Type': "application/json"
        }
    }).then(result => {
        var data = result.data;
        console.log(data);
        if (result.status == 200 || result.status == 201) {
           commonAlert("Success Delete Monitoring Alert Event-Handler");
           displayMonitoringAlertEventHandlerInfo("DEL_SUCCESS");
        }else{
            commonAlert(data)
        }
    }).catch(function(error){
        commonAlert(error)
        console.log("Monitoring Alert Event-Handler delete error : ",error);        
    });
}


function displayMonitoringAlertEventHandlerInfo(targetAction) {
    if ( targetAction == "DEL_SUCCESS" ) {
        console.log("##########MonitoringAlertEventHandler DEL_SUCCESS")
        
        getMonitoringAlertEventHandlerList();
    } else if ( targetAction == "REG_SUCCESS" ) {
        console.log("##########MonitoringAlertEventHandler REG_SUCCESS")
        
        $("#regMonitoringAlertEventHandlerModalType").val('');  
        $("#regMonitoringAlertEventHandlerModalName").val('');  

        getMonitoringAlertEventHandlerList();

        $("#Add_MonitoringAlertEventHandler_Register").modal("hide");
    }
  
}

// Monitoring Alert Event-Handler목록 조회
function getMonitoringAlertEventHandlerList() {
    console.log("#####################getMonitoringAlertEventHandlerList : ");

    var url = "/operation/policies/monitoringalerteventhandler/list"
    axios.get(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get Threshold List : ", result.data);
        var data = result.data.MonitoringAlertEventHandlerList;
        console.log("$$$EventHandler DATA$$$");
        console.log(data);

        var html = ""
        var cnt = 0;

        if (data.length) {
            data.filter((list) => list.Name !== "").map((item, index) => (
                html += addMonitoringAlertEventHandlerRow(item, index)
            ))
        }

        $("#monitoringAlertEventHandlerList").empty();
        $("#monitoringAlertEventHandlerList").append(html);

        //ModalDetail();
        
    }).catch(function(error){
        console.log("Event-Handler list error : ", error);        
    })    
}

// Monitoring Alert Event-Handler목록에 Item 추가
function addMonitoringAlertEventHandlerRow(item, index){
    console.log("addMonitoringAlertEventHandlerRow " + index);
    console.log(item)
    var html = ""

    html += '<tr><td class="overlay hidden" data-th="">'
        + '<input type="hidden" id="monitoringAlertEventHandlerType_info_' + index + '" value="' + item.type + '"/>'
        + '<input type="hidden" id="monitoringAlertEventHandlerName_info_' + index + '" value="' + item.name + '"/>'
        + '<input type="checkbox" name="chk" value="' + item.type +'" id="raw_' + index + '" title="" />'
        + '<label for="td_ch1"></label> <span class="ov off"></span></td>'
        + '<td class="btn_mtd ovm" data-th="Type">' + item.type + '<span class="ov"></span>'
        + '<input type="hidden" id="monitoringAlertEventHandler_info_' + index + '" value="' + item.type + '"/>'
        + '<td class="overlay hidden" data-th="Name">' + item.name + '</td></tr>'
	
    return html;
}