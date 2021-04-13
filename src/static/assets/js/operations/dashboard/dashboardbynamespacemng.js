
// dashboard 의 MCIS 목록에서 mcis 선택 : 색상반전, 선택한 mcis id set -> status변경에 사용
function selectMcis(id,name,target){
    var mcis_id = id
    var mcis_name = name
    var init_select_areabox = $("#init_select_areabox").val()
    $target = $("#"+target)
    if($target.hasClass("active")){
        location.href = "/Manage/MCIS/list/"+mcis_id+"/"+mcis_name
        return;
    }
    $("[id^='mcis_areabox_']").each(function(){
        var s_id = $(this).attr("id");
        if(s_id == target){
            $(this).addClass("active");
         
        }else{
            $(this).removeClass("active");
           
        }
    })

    $("#mcis_id").val(mcis_id)
    $("#mcis_name").val(mcis_name)    
 }

 //
 function mcisLifeCycle(type){

    // 선택된 mcis 가 있는지 체크.
    /////// TODO : util.mcislifecycle.js 를 호출하도록 변경
    // 이름 바꿀 것.   
 }