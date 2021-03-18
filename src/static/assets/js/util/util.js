
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