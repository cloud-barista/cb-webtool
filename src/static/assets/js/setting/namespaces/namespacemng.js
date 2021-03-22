$(document).ready(function(){
    order_type = "name"
    // getNSList(order_type);-> getNameSpaceList 으로 이름변경. 이미 가져왔음.
})

// commons.js에 정의 됨
// function fnMove(target){
// var offset = $("#" + target).offset();
// console.log("fn move offset : ",offset)
// $('html, body').animate({scrollTop : offset.top}, 400);
// }


// function getNSList(sort_type){
function getNameSpaceList(sort_type){
    clearNamespaceInfo();
    var url = "/setting/namespaces"+"/namespace/list";
    axios.get(url,{
        headers:{
            'Content-Type' : "application/json"
        }
    }).then(result=>{
        console.log("get NameSpace Data : ",result.data);
        // var data = result.data.ns;
        var data = result.data;
        var html = ""
        if(data.length){ 
            if(sort_type){            
                data.filter(list=> list.name !=="" ).sort((a,b) => ( a[sort_type] < b[sort_type] ? -1 : a[sort_type] > b[sort_type] ? 1 : 0)).map((item,index)=>(
                    html +='<tr onclick="showNameSpaceInfo(\'ns_info_'+index+'\');">'
                        +'<td class="overlay hidden" data-th="">'
                        +'<input type="hidden" id="ns_info_'+index+'" value="'+item.id+'|'+item.name+'|'+item.description+'"/>'
                        +'<input type="checkbox" name="chk" value="'+item.name+'" id="raw_'+index+'" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>'
                        +'<td class="btn_mtd ovm" data-th="Name">'+item.name+'<span class="ov"></span></td>'
                        +'<td class="overlay hidden" data-th="ID">'+item.id+'</td>'
                        +'<td class="overlay hidden td_left" data-th="description">'+item.description+'</td>'
                        +'<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
                        +'</tr>' 
                ))
            }else{
                data.filter((list)=> list.name !== "" ).map((item,index)=>(
                    html +='<tr onclick="showNameSpaceInfo(\'ns_info_'+index+'\');">'
                        +'<td class="overlay hidden" data-th="">'
                        +'<input type="hidden" id="ns_info_'+index+'" value="'+item.id+'|'+item.name+'|'+item.description+'"/>'
                        +'<input type="checkbox" name="chk" value="'+item.name+'" id="raw_'+index+'" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>'
                        +'<td class="btn_mtd ovm" data-th="Name">'+item.name+'<span class="ov"></span></td>'
                        +'<td class="overlay hidden" data-th="ID">'+item.id+'</td>'
                        +'<td class="overlay hidden td_left" data-th="description">'+item.description+'</td>'
                        +'<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
                        +'</tr>'        
                ))

            }		
            $("#nsList").empty();
            $("#nsList").append(html);
                       
            ModalDetail()        
        }//end of data length
    })
}

// common.js 로 이동
// function goFocus(target){

// console.log(event)
// event.preventDefault()

// $("#"+target).focus();
// fnMove(target)

// }

// function showInfo(target){
function showNameSpaceInfo(target){
    console.log("target : ",target);
    var infos = $("#"+target).val()
    infos = infos.split("|")
    $("#info_id").val(infos[0])
    $("#info_name").val(infos[1])
    $("#info_desc").val(infos[2])
    
    $("#info_name").focus();
}

// 삭제 처리 후 namespace 상세정보 초기화
function clearNamespaceInfo(){
    $("#info_id").val('')
    $("#info_name").val('')
    $("#info_desc").val('')
}

//function createNS(){
function createNameSpace(){
    var namespace = $("#reg_name").val()
    var desc = $("#reg_desc").val()
    if(!namespace){
        alert("Input New NameSpace")
        $("#reg_name").focus()
        return;
    }
    
    // var apiInfo = "{ { .apiInfo} }";
    var url = "/setting/namespaces"+"/namespace/reg/proc";
    var obj = {
        name: namespace,
        description : desc
    }
    if(namespace){
        axios.post(url,obj,{
            headers: { 
                'Content-type': 'application/json',
                // 'Authorization': apiInfo, 
            }
        }).then(result =>{
            console.log(result);
            if(result.status == 200 || result.status == 201){
                alert("Success Create NameSpace")
                //등록하고 나서 화면을 그냥 고칠 것인가?
                getNameSpaceList();
                //아니면 화면을 리로딩 시킬것인가?
                // location.reload();
                // // $("#btn_add2").click()
                // $("#namespace").val('')
                // $("#nsDesc").val('')
            }else{
                alert("Fail Create NameSpace")
            }
        });
    }else{
        alert("Input NameSpace")
        $("#reg_desc").focus()
        return;
    }
}

// 삭제 : TODO : spider api 확인하여 실제 삭제, 수정작업 되는지 Test 할 것.
function deleteNameSpace () {
    var nameSpaceID = $("#info_id").val()
    var nameSpaceName = $("#info_name").val()
    var nameSpaceDescription = $("#info_desc").val()

    if(!nameSpaceID){
        alert("select NameSpace")        
        return;
    }
    
    var url = "/setting/namespaces"+"/namespace/del/" + nameSpaceID;
    // var obj = {
    //     // id : nameSpaceID,
    //     // name: nameSpaceName,
    //     // description : nameSpaceDescription
    //     nameSpaceID : nameSpaceID
    // }
    if(nameSpaceID){
        // axios.post(url,obj,{
        //     headers: { 
        //         'Content-type': 'application/json',
        //         // 'Authorization': apiInfo, 
        //     }
        axios.delete(url, {},{
        }).then(result =>{
            console.log(result);
            if(result.status == 200 || result.status == 201){
                alert("Success delete NameSpace")
                //등록하고 나서 화면을 그냥 고칠 것인가?
                getNameSpaceList();
                
            }else{
                alert("Fail delete NameSpace")
            }
        }).catch(function(error){
            console.log("namespace delete error : ",error);        
        });
    }else{
        alert("Input NameSpace")
        $("#reg_desc").focus()
        return;
    }
}

function getNS(){

}
function ModalDetail(){
    $(".dashboard .status_list tbody tr").each(function(){
        var $td_list = $(this),
                $status = $(".server_status"),
                $detail = $(".server_info");
        $td_list.off("click").click(function(){
                $td_list.addClass("on");
                $td_list.siblings().removeClass("on");
                $status.addClass("view");
                $status.siblings().removeClass("on");
                $(".dashboard.register_cont").removeClass("active");
            $td_list.off("click").click(function(){
                    if( $(this).hasClass("on") ) {
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
