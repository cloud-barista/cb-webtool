$(document).ready(function(){	
			
    //tab menu Server / OS_HW /	Network / Security / Other 위치 표시
    $(".create_tab .nav a").click(function() {
      var idx = $(".create_tab .nav a").index(this);
      for (i = 0; i < $(".create_tab .nav a").length; i++) {
        if (i == idx) {
                $('.config_location > div').removeClass('on');
                $('.config_location > div > span').eq(idx).parent().addClass('on');
        } 
      }
    });
    //tab 내용 다음
  $(".create_tab .btn_next").click(function(e) {
    var $active = $('.create_tab .nav li > .active');
    $active.parent().next().find('.nav-link').removeClass('disabled');
    nextTab($active);
  });

    //tab 내용 이전
  $(".create_tab .btn_prev").click(function(e) {
    var $active = $('.create_tab .nav li > a.active');
    prevTab($active);
  });
  
});
//next
function nextTab(elem) {
  $(elem).parent().next().find('a[data-toggle="tab"]').click();
}
//prev
function prevTab(elem) {
  $(elem).parent().prev().find('a[data-toggle="tab"]').click();
}

$(document).ready(function(){	
    //Deployment Target table scrollbar
    $('.btn_assist').on('click', function() {
        $("#Deployment_box").modal();
        $('.dtbox.scrollbar-inner').scrollbar();
    });
    //Server Configuration clear
    $(".btn_clear").click(function() {
        $('.svc_ipbox').find('input, textarea').val('');
    });


    //OS_HW - Clear
    $("#OS_HW .btn_clear").click(function() {
        $('#OS_HW .tab_ipbox').find('input, textarea').val('');
    });
    //Network - Clear
    $("#Network .btn_clear").click(function() {
        $('#Network .tab_ipbox').find('input, textarea').val('');
    });
    //Security - Clear
    $("#Security .btn_clear").click(function() {
        $('#Security .tab_ipbox').find('input, textarea').val('');
    });
    //Other - Clear
    $("#Other .btn_clear").click(function() {
        $('#Other .tab_ipbox').find('input, textarea').val('');
    });
});

// Expert Mode=on 상태에서 Cloud Provider 를 변경했을 때, 해당 Provider의 region목록 조회 => 실제로는 조회되어 있으므로 filter
// 추가로 connection 정보도 조회하라고 호출
function getRegionListForSelectbox(provider, targetRegionObj, targetConnectionObj){


    //getConnectionListForSelectbox(this.value, 's_regConnectionName');
}

// Expert Mode=on 상태에서 Popup의 Cloud Provider 를 변경했을 때, 해당 Provider의 region목록 조회. 
// getRegionListForSelectbox() 와 동작방식은 동일
function getRegionListForPopSelectbox(provider, targetRegionObj, targetConnectionObj){

}

// region 변경시, 해당 provider, region으로 connection 목록 조회
function getConnectionListByRegionForSelectbox(region, targetProviderObj, targetConnectionObj){

}