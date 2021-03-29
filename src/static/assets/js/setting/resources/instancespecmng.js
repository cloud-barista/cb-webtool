$(document).ready(function(){
    //action register open / table view close
    // $('#RegistBox .btn_ok.register').click(function(){
    //     $(".dashboard.register_cont").toggleClass("active");
    //     $(".dashboard.server_status").removeClass("view");
    //     $(".dashboard .status_list tbody tr").removeClass("on");
    //     //ok 위치이동
    //     $('#RegistBox').on('hidden.bs.modal', function () {
    //         var offset = $("#CreateBox").offset();
    //         $("#wrap").animate({scrollTop : offset.top}, 300);
    //     })		
    // });
});

/* scroll */
$(document).ready(function(){
    //checkbox all
    // $("#th_chall").click(function() {
    //     if ($("#th_chall").prop("checked")) {
    //         $("input[name=chk]").prop("checked", true);
    //     } else {
    //         $("input[name=chk]").prop("checked", false);
    //     }
    // })
        
    //     //table 스크롤바 제한
    // $(window).on("load resize",function(){
    //         var vpwidth = $(window).width();
    //     if (vpwidth > 768 && vpwidth < 1800) {
    //         $(".dashboard_cont .dataTable").addClass("scrollbar-inner");
    //             $(".dataTable.scrollbar-inner").scrollbar();
    //     } else {
    //         $(".dashboard_cont .dataTable").removeClass("scrollbar-inner");
    //     }
    // });
});

$(document).ready(function () {
    order_type = "name"
    getInstanceSpecList(order_type);

    // var apiInfo = "{{ .apiInfo}}";
    // getCloudOS(apiInfo,'provider');
});

// function goFocus(target) {
//     console.log(event)
//     event.preventDefault();

//     $("#" + target).focus();
//     fnMove(target)
// }

// function fnMove(target) {
//     var offset = $("#" + target).offset();
//     console.log("fn move offset : ", offset);
//     $('html, body').animate({
//         scrollTop: offset.top
//     }, 400);
// }

// 등록/상세 area 보이기 숨기기
function displayInstanceSpecInfo(targetAction){
    if( targetAction == "REG"){
        $('#instanceSpecCreateBox').toggleClass("active");
        $('#instanceSpecInfoBox').removeClass("view");
        $('#instanceSpecListTable').removeClass("on");
        var offset = $("#instanceSpecCreateBox").offset();
        // var offset = $("#" + target+"").offset();
    	$("#TopWrap").animate({scrollTop : offset.top}, 300);

        // form 초기화
        $("#regSpecName").val('')
        $("#regCspSpecName").val('')

    }else if ( targetAction == "REG_SUCCESS"){
        $('#instanceSpecCreateBox').removeClass("active");
        $('#instanceSpecInfoBox').removeClass("view");
        $('#instanceSpecListTable').addClass("on");
        
        // form 초기화
        $("#regSpecName").val('')
        $("#regCspSpecName").val('')
        
        var offset = $("#instanceSpecCreateBox").offset();
        $("#TopWrap").animate({scrollTop : offset.top}, 0);
        
        getInstanceSpecList("name");
    }else if ( targetAction == "DEL"){
        $('#instanceSpecCreateBox').removeClass("active");
        $('#instanceSpecInfoBox').addClass("view");
        $('#instanceSpecListTable').removeClass("on");

        var offset = $("#instanceSpecInfoBox").offset();
    	$("#TopWrap").animate({scrollTop : offset.top}, 300);

    }else if ( targetAction == "DEL_SUCCESS"){
        $('#instanceSpecCreateBox').removeClass("active");
        $('#instanceSpecInfoBox').removeClass("view");
        $('#instanceSpecListTable').addClass("on");

        var offset = $("#instanceSpecInfoBox").offset();
        $("#TopWrap").animate({scrollTop : offset.top}, 0);

        getInstanceSpecList("name");
    }
}

function getInstanceSpecList(sort_type) {
    console.log(sort_type);
    var url = CommonURL + "/ns/" + NAMESPACE + "/resources/spec";
    axios.get(url, {
        headers: {
            // 'Authorization': "{{ .apiInfo}}",
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get Spec List : ", result.data);
        
        var data = result.data.spec;
        var html = ""
        
        if (data.length) {
            if (sort_type) {
                console.log("check : ", sort_type);
                data.filter(list => list.name !== "").sort((a, b) => (a[sort_type] < b[sort_type] ? - 1 : a[sort_type] > b[sort_type] ? 1 : 0)).map((item, index) => (
                    html += '<tr onclick="showInfo(\'' + item.name + '\');">' 
                        + '<td class="overlay hidden" data-th="">' 
                        + '<input type="hidden" id="spec_info_' + index + '" value="' + item.name + '|' + item.connectionName + '|' + item.cspSpecName + '"/>' 
                        + '<input type="checkbox" name="chk" value="' + item.name + '" id="raw_'  + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' 
                        + '<td class="btn_mtd ovm" data-th="name ">' + item.name  + '<span class="ov"></span></td>'
                        + '<td class="overlay hidden" data-th="connectionName">' + item.connectionName + '</td>' 
                        + '<td class="overlay hidden" data-th="cspSpecName">' + item.cspSpecName + '</td>'  
                        + '<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
                        + '</tr>'
                ))
            } else {
                data.filter((list) => list.name !== "").map((item, index) => (
                    html += '<tr onclick="showInfo(\'' + item.name + '\');">' 
                        + '<td class="overlay hidden" data-th="">' 
                        + '<input type="hidden" id="spec_info_' + index + '" value="' + item.name + '"/>' 
                        + '<input type="checkbox" name="chk" value="' + item.name + '" id="raw_'  + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' 
                        + '<td class="btn_mtd ovm" data-th="name ">' + item.name  + '<span class="ov"></span></td>'
                        + '<td class="overlay hidden" data-th="connectionName">' + item.connectionName + '</td>' 
                        + '<td class="overlay hidden" data-th="cspSpecName">' + item.cspSpecName + '</td>'  
                        + '<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
                        + '</tr>'
                ))
            }
        
            $("#specList").empty()
            $("#specList").append(html)
            
            ModalDetail()
        }
    })
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

function showInfo(target) {
    console.log("target showInfo : ", target);
    // var apiInfo = "{{ .apiInfo}}";
    var specId = encodeURIComponent(target);
    
    var url = CommonURL+"/ns/"+NAMESPACE+"/resources/spec/"+ specId;
    console.log("spec detail URL : ",url)

    return axios.get(url,{
        headers:{
            // 'Authorization': apiInfo
        }
    
    }).then(result=>{
        var data = result.data
        console.log("Show Data : ",data);

        var dtlSpecName = data.name;
        var dtlConnectionName = data.connectionName;
        var dtlCspSpecName = data.cspSpecName;

        $("#dtlSpecName").empty();
        $("#dtlProvider").empty();
        $("#dtlConnectionName").empty();
        $("#dtlCspSpecName").empty();
        

        $("#dtlSpecName").val(dtlSpecName);
        $("#dtlConnectionName").val(dtlConnectionName);
        $("#dtlCspSpecName").val(dtlCspSpecName);

        getProvider(dtlConnectionName);
    }) 
}

function getProvider(target) {
    console.log("getProvidergetProvider : ",target);
    var url = SpiderURL+"/connectionconfig/" + target;
        
    return axios.get(url,{
        headers:{
            // 'Authorization': apiInfo
        }
    
    }).then(result=>{
        var data = result.data;
        
        var Provider = data.ProviderName;
        console.log("Provider : ", Provider);
        $("#dtlProvider").val(Provider);
    })        
}							

function getConnectionInfo(provider){
    var url = SpiderURL+"/connectionconfig";
    console.log("provider : ",provider)
    //var provider = $("#provider option:selected").val();
    var html = "";
    // var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            // 'Authorization': apiInfo
        }
    }).then(result=>{
        console.log('getConnectionConfig result: ',result)
        var data = result.data.connectionconfig
        console.log("connection data : ",data);
        var count = 0; 
        var configName = "";
        var confArr = new Array();
        for(var i in data){
            if(provider == data[i].ProviderName){ 
                count++;
                html += '<option value="'+data[i].ConfigName+'" item="'+data[i].ProviderName+'">'+data[i].ConfigName+'</option>';
                configName = data[i].ConfigName
                confArr.push(data[i].ConfigName)
                
            }
        }
        if(count == 0){
            alert("해당 Provider에 등록된 Connection 정보가 없습니다.")
                html +='<option selected>Select Configname</option>';
        }
        if(confArr.length > 1){
            configName = confArr[0];
        }
        $("#reg_connectionName").empty();
        $("#reg_connectionName").append(html);

    })
}

function createSpec() {
    var specId = $("#reg_specName").val();
    var specName = $("#reg_specName").val();
    var connectionName = $("#reg_connectionName").val();
    var cspSpecName = $("#reg_cspSpecName").val();
    
    if (!specName) {
        alert("Input New Spec Name")
        $("#reg_specName").focus()
        return;
    }

    // var apiInfo = "{{ .apiInfo}}";
    var url = CommonURL+"/ns/"+NAMESPACE+"/resources/spec"
    console.log("Spec Reg URL : ",url)
    var obj = {
        id: specId,
        name: specName,
        connectionName: connectionName,
        cspSpecName: cspSpecName
    }
    console.log("info image obj Data : ", obj);
    
    if (specName) {
        axios.post(url, obj, {
            headers: {
                'Content-type': 'application/json',
                // 'Authorization': apiInfo,
            }
        }).then(result => {
            console.log("result spec : ", result);
            if (result.status == 200 || result.status == 201) {
                alert("Success Create Image!!")
                //등록하고 나서 화면을 그냥 고칠 것인가?
                getInstanceSpecList();
                //아니면 화면을 리로딩 시킬것인가?
                // location.reload();
                // $("#btn_add2").click()
                // $("#namespace").val('')
                // $("#nsDesc").val('')
            } else {
                alert("Fail Create Spec")
            }
        });
    } else {
        alert("Input Spec Name")
        $("#reg_specName").focus()
        return;
    }
}

function goDelete() {
    var selSpecId = "";
    var count = 0;

    $( "input[name='chk']:checked" ).each (function (){
        count++;
        selSpecId = selSpecId + $(this).val()+"," ;
    });
    selSpecId = selSpecId.substring(0,selSpecId.lastIndexOf( ","));
    
    console.log("specId : ", selSpecId);
    console.log("count : ", count);

    if(selSpecId == ''){
        alert("삭제할 대상을 선택하세요.");
        return false;
    }

    if(count != 1){
        alert("삭제할 대상을 하나만 선택하세요.");
        return false;
    }
    
    var url = CommonURL + "/ns/" + NAMESPACE + "/resources/spec/" + selSpecId;
    
    axios.delete(url, {
        headers: {
            // 'Authorization': "{{ .apiInfo}}",
            'Content-Type': "application/json"
        }
    }).then(result => {
        var data = result.data;
        if (result.status == 200 || result.status == 201) {
            alert("Success Delete Spec.");
            // location.reload(true);
        }
    })
}                                                  
