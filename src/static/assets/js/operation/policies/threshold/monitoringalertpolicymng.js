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