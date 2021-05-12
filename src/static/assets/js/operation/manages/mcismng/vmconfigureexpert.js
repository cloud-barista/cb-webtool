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

$(document).ready(function(){
  //table 스크롤바 제한
  $(window).on("load resize",function(){
    var vpwidth = $(window).width();
    if (vpwidth > 768 && vpwidth < 1800) {
      $(".dashboard_cont .dataTable").addClass("scrollbar-inner");
      $(".dataTable.scrollbar-inner").scrollbar();
    } else {
      $(".dashboard_cont .dataTable").removeClass("scrollbar-inner");
    }
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

// TODO : util.js로 옮길 것
// select box의 option text에 compareText가 있으면 show 없으면 hide
function selectBoxFilterByText(targetObject, compareText){
  $('#' + targetObject +' option').filter(function() {
    if( this.value == "") return;
    console.log(this.text + " : " + compareText)
    console.log(this.text.indexOf(compareText) > -1)
    this.text.indexOf(compareText) > -1 ? $(this).show() : $(this).hide();    
  });
}

// TODO : util.js로 옮길 것
// select box의 option text에 compareText1 && compareText2가 모두 있으면 show 없으면 hide
function selectBoxFilterBy2Texts(targetObject, compareText1, compareText2){
  $('#' + targetObject +' option').filter(function() {
    if( this.value == "") return;
    console.log(this.text + " : " + compareText)
    console.log(this.text.indexOf(compareText) > -1)
    if ( this.text.indexOf(compareText1) > -1 && this.text.indexOf(compareText2) > -1 ){
      $(this).show()
    }else{
      $(this).hide(); 
    }
  });
}

// Expert Mode=on 상태에서 Cloud Provider 를 변경했을 때, 해당 Provider의 region목록 조회 => 실제로는 조회되어 있으므로 filter
// 추가로 connection 정보도 조회하라고 호출
function getRegionListFilterForSelectbox(provider, targetRegionObj, targetConnectionObj){

  // region 목록 filter
  selectBoxFilterByText(targetRegionObj, provider)
  $("#" + targetRegionObj + " option:eq(0)").attr("selected", "selected");

  // connection 목록 filter
  selectBoxFilterByText(targetConnectionObj, provider)
  $("#" + targetConnectionObj + " option:eq(0)").attr("selected", "selected");
}

// region변경시 connection 정보 filter
function getConnectionListFilterForSelectbox(region, referenceObj, targetConnectionObj){
  var referenceVal = $('#' + referenceObj).val();
  var regionValue = region.substring(region.indexOf("]") ).trim();  
  console.log(region + ", regionValue = " + regionValue);
  if( referenceVal == ""){
    selectBoxFilterByText(targetConnectionObj, regionValue)
  }else{
    selectBoxFilterBy2Texts(targetConnectionObj, referenceVal, regionValue)
  }
  $("#" + targetConnectionObj + " option:eq(0)").attr("selected", "selected");
}

// TODO : filter 기능 check
// provider, region, connection은 먼저 선택이 필수가 아닐 수 있음.
// 그래도 하위에서 일단 선택되면 변경시 알려줘야할 듯.
// 1. provider 선택시 -> 
// 2. region 선택시
// 3. OS Platform(Image) 선택 시
// 4. HW Spec 선택시
// 5. Vnet 선택시
// 6. SecurityGroup 선택시
// 7. sshKey 선택시
// 8. subnet 선택시??

//e_imageID

// Asist를 클릭했을 때 나타나는 popup에서 provider 변경 시 region selectbox와 connection table을 filter
function popProviderChange(providerObj, regionObj, targetTableObj ){
  var providerVal = $("#" + providerObj).val();
  console.log("popProviderChange " + providerVal);
  selectBoxFilterByText(regionObj, providerVal)

  $("#" + regionObj + " option:eq(0)").attr("selected", "selected");

  // table filter
  getConnectionListFilterForTable(providerObj, regionObj, targetTableObj)
}

function getConnectionListFilterForTable(providerObj, regionObj, targetTableObj){
  var providerVal = $("#" + providerObj).val();
  var regionVal = $("#" + regionObj).val();

  $("#" + targetTableObj + " > tbody >  tr").filter(function() {
    console.log("filter table " + $(this).text());
    var compareText = $(this).text().toLowerCase()
    var toggleStatus = true;
    if( providerVal == "" && regionVal == "" ){
      return;
    }else if( providerVal == "" && compareText.indexOf(regionVal.toLowerCase()) > -1 ){
      toggleStatus = true
    }else if( regionVal == "" && compareText.indexOf(providerVal.toLowerCase()) > -1 ){
      toggleStatus = true
    }else if( compareText.indexOf(providerVal.toLowerCase()) > -1 && compareText.indexOf(regionVal.toLowerCase()) > -1 ){
      toggleStatus = true
    }else {
      toggleStatus = false
    }
    //$(this).toggle(toggleStatus)
    if( toggleStatus){
      $(this).show();
    }else{
      $(this).hide();
    }
  });

}
// Expert Mode=on 상태에서 Popup의 Cloud Provider 를 변경했을 때, 해당 Provider의 region목록 조회. 
// getRegionListForSelectbox() 와 동작방식은 동일
function getRegionListForPopSelectbox(provider, targetRegionObj, targetConnectionObj){
  $('#' + targetRegionObj +' option').filter(function() {
    if( this.value == "") return;

    return this.text.indexOf(provider) > -1 ? $(this).show() : $(this).hide();    
  });

  // connection filter
}

// region 변경시, 해당 provider, region으로 connection 목록 조회
function getConnectionListByRegionForSelectbox(region, targetProviderObj, targetConnectionObj){

}


const Expert_Server_Config_Arr = new Array();
var expert_data_cnt = 0
const expertServerCloneObj = obj=>JSON.parse(JSON.stringify(obj))
function expertDone_btn(){
}



//////////////////// filterling 기능 ///////////////

function hardwareSpecFilterByEnter(targetObjId, keyword){
  console.log(event.KeyCode + " : " + keyword);
  if( event.KeyCode == 13){
    selectBoxFilterByText(targetObjId, keyword)
  }
}

function hardwareSpecFilter(targetObjId, keywordObjId){
  var keyword = $("#" + keywordObjId).val();
  selectBoxFilterByText(targetObjId, keyword)
}


// hidden에 들어있는 값을 기준으로 filter
function filterEnterToHidden(keyworkObjId, filterColumnName, targetTableId){
  if (window.event.keyCode == 13) {
    filterToHidden(keyworkObjId, filterColumnName, targetTableId)
      }
}
// vm image 를 filter. 입력한 단어를 전체에서 찾기
// hidden에 들어있는 값을 기준으로 filter
function filterToHidden(keyworkObjId, filterColumnName, tableId){
  var keyword = $("#" + keyworkObjId).val();
  if( keyword == ''){
    commonAlert("검색할 단어를 입력하세요")
  }

  // hidden fild 의 Data 를 기준으로 filter.
  // 해당 table의 column을 지정하여 filter하는데... hidden을 filterling하는 법 찾자
  filterTableByHiddenColumn(tableId, filterColumnName, keyword)

  // target Table이 display:none이면 보이도록
  document.getElementById(tableId).style.display = "";
}

// 대상 table, 선택한 tr의 index, set할 값, 대상 form의 obj 지정하여
// tr의 check시 해당 값이 obj에 저장
// TODO : 초기화는??	
function setValueToFormObj(tableId, targetTabObjId, prefixName, selectedIndex, targetObjId){
  var selectedId = $("#" + prefixName + "_id_" + selectedIndex).val();
  var selectedInfo = $("#" + prefixName + "_info_" + selectedIndex).val();
  $("#" + targetTabObjId).val(selectedInfo);
  $("#" + targetObjId).val(selectedId);

  // 값이 설정된 후에는 table 안보이게
  document.getElementById(tableId).style.display = "none";
}	

// Textbox 값이 변경 된 경우 해당 값을 form obj에 set
function setTextValueToFormObj(setValue, targetObjId){
  $("#" + targetObjId).val(setValue);
}
//////////////////// filterling 기능 //////////////