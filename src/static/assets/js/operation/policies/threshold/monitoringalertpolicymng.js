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
}


function deleteMonitoringAlertPolicy(){
    console.log("##########DeleteMonitoringAlertPolicy")
}


function createMonitoringAlertPolicy(){
    console.log("##########CreateMonitoringAlertPolicy")
}


function showMonitoringAlertPolicyInfo(monitoringAlertName) {
    console.log("showMonitoringAlertPolicyInfo : ", monitoringAlertName);
}