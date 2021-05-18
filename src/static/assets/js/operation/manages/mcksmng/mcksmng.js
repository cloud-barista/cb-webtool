function clickListOfMcks(uid, mcksIndex){
    console.log("click view mcks id :",uid)
    $(".server_status").addClass("view");

    // List Of MCKS에서 선택한 row 외에는 안보이게
    $("[id^='server_info_tr_']").each(function(){
        var item = $(this).attr("item").split("|")
        console.log(item)
        if(id == item[0]){           
            $(this).addClass("on")
        }else{
            $(this).removeClass("on")
        }
    })

    $("#mcks_uid").val(uid);

    // MCIS Info area set
    showServerListAndStatusArea(uid,mcksIndex);
}


// MCKS Info area 안의 Node List 내용 표시
// 해당 MCKS의 모든 Node 표시
// TODO : 클릭했을 때 서버에서 조회하는것으로 변경할 것.
function showServerListAndStatusArea(uid, mcksIndex){
    
    var mcksUID =  $("#mcksUID" + mcksIndex).val();
    var mcksName =  $("#mcksName" + mcksIndex).val();
    var mcksStatus =  $("#mcksStatus" + mcksIndex).val();
    var mcksConfig = $("#mcksConfig" + mcksIndex).val();
    var nodeTotalCountOfMcks = $("#mcksNodeTotalCount" + mcksIndex).val();

    $(".server_status").addClass("view")
    $("#mcks_info_txt").text("[ "+ mcksName +" ]");
    $("#mcks_server_info_status").empty();
    $("#mcks_server_info_status").append('<strong>Server List </strong>  <span class="stxt">[ '+mcksName+' ]</span>  Server('+nodeTotalCountOfMcks+')')

    //
    $("#mcks_info_name").val(mcksName+" / "+mcksUID)
    $("#mcks_info_cloud_connection").val(mcksConfig) 
    
    $("#mcks_name").val(mcksName)

    var mcks_badge = "";
    var mcksStatusIcon = "";
    if(mcksStatus == "running"){ mcksStatusIcon = "icon_running_db.png"        
    }else if(mcksStatus == "include" ){ mcksStatusIcon = "icon_stop_db.png"
    }else if(mcksStatus == "suspended"){mcksStatusIcon = "icon_stop_db.png"
    }else if(mcksStatus == "terminate"){mcksStatusIcon = "icon_terminate_db.png"
    }else{
        mcksStatusIcon = "icon_stop_db.png"
    }
    mcks_badge = '<img src="/assets/img/contents/' + mcksStatusIcon +'" alt=""/> '
    $("#service_status_icon").empty();
    $("#service_status_icon").append(mcks_badge)

        

    //Manage MCKS Server List on/off : table을 클릭하면 해당 Row 에 active style로 보여주기
    $(".dashboard .ds_cont .area_cont .listbox li.sel_cr").each(function(){
        var $sel_list = $(this);
        var $detail = $(".server_info");
        console.log($sel_list);
        console.log($detail);
        console.log(">>>>>");
        $sel_list.off("click").click(function(){
            $sel_list.addClass("active");
            $sel_list.siblings().removeClass("active");
            $detail.addClass("active");
            $detail.siblings().removeClass("active");
            $sel_list.off("click").click(function(){
                if( $(this).hasClass("active") ) {
                    $sel_list.removeClass("active");
                    $detail.removeClass("active");
                } else {
                    $sel_list.addClass("active");
                    $sel_list.siblings().removeClass("active");
                    $detail.addClass("active");
                    $detail.siblings().removeClass("active");
                }
            });
        });
    }); 
}
