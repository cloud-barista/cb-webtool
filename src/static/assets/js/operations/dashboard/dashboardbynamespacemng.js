
// dashboard 의 MCIS 목록에서 mcis 선택 : 색상반전, 선택한 mcis id set -> status변경에 사용
function selectMcis(id,name,target){
    console.log("selectMcis")
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
        console.log(s_id + ":" + target)
        if(s_id == target){
            try{
                var s_id = $(this).attr("id");
                $(this).addClass("active"); 
                console.log(s_id + " addClass active")
            }catch(e){
                console.log(e)
            }

        }else{
            $(this).removeClass("active");
            // console.log(s_id + "removeClass active")
        }
    })
    // console.log(" active / deactive ")
    $("#mcis_id").val(mcis_id)
    $("#mcis_name").val(mcis_name)    
    console.log(" mcis_id =" + mcis_id + ", mcis_name = " + mcis_name);
 }

// callMcisLifeCycle -> McisLifeCycle -> callbackMcisLifeCycle
// confirm창을 띄울 때 mcismng와 동일한 key로 호출하므로 callback함수 이름도 같아야 한다.(util.js 참조)
function callMcisLifeCycle(type){
    var selectedCount = 0;
    // 선택된 mcis 가 있는지 체크.
    $("[id^='mcis_areabox_']").each(function(){        
        if($(this).hasClass("active")){            
            selectedCount++
            mcisLifeCycle($("#mcis_id").val(), type);//mcislifecycle.js 호출
        }
    })

    if( selectedCount == 0){
        commonAlert("Please Select MCIS!!")
    }

    /////// TODO : util.mcislifecycle.js 를 호출하도록 변경
    
}

// McisLifeCycle을 호출 한 뒤 return값 처리
function callbackMcisLifeCycle(resultStatus, resultData, type){
    var message = "MCIS "+type+ " complete!."
    if(status == 200 || status == 201){            
        commonAlert(message);
        location.reload();//완료 후 페이지를 reload -> 해당 mcis만 reload
        // 해당 mcis 조회
        // 상태 count 재설정
    }
}