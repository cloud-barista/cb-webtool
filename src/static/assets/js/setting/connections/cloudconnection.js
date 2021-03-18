
$(document).ready(function(){
    order_type = "ConfigName"

    
    // // defaultnamespace 확인 
    // defaultNamespace = '{{.LoginInfo.DefaultNameSpaceName}}'

    // if( defaultNamespace == ''){
    //     alert("Namespace를 선택바랍니다.")
    // }else{
    //     alert(defaultNamespace + "Dashboard로 이동.")
    // }

    //getNameSspaceList();// 이미 가져왔음

    // getCloudConnectionList(order_type);

    // getCloudOS("{{ .apiInfo}}",'RegionModalProviderName');
    // getCloudOS("{{ .apiInfo}}",'CredentialModalProviderName');
    // getCloudOS("{{ .apiInfo}}",'DriverModalProviderName');


})


function fnMove(target){
    var offset = $("#" + target).offset();
    console.log("fn move offset : ",offset)
    $('html, body').animate({scrollTop : offset.top}, 400);
}

// 현재생성된 connection config 목록. 
function getCloudConnectionList(sort_type){
    // 원래는 목록을 조회해서 filterling 하는 function
    // 이미 목록을 가져왔으므로 
    // TODO : 가져온 것을 filtering 하는 것으로 변경 필요
    alert(sort_type + "만 filtering 하자.")
    // var url = "{{ .comURL.SpiderURL}}"+"/connectionconfig";
    // axios.get(url,{
    //     headers:{
    //         'Authorization': "{{ .apiInfo}}",
    //         'Content-Type' : "application/json"
    //     }
    // }).then(result=>{
    //     console.log("get CloudConnection Data : ",result.data);
    //     var data = result.data.connectionconfig;
    //     var html = ""
    //     if(data.length){ 
    //         if(sort_type){
                
    //             data.filter(list=> list.ConfigName !=="" ).sort((a,b) => ( a[sort_type] < b[sort_type] ? -1 : a[sort_type] > b[sort_type] ? 1 : 0)).map((item,index)=>(
    //                 html +='<tr onclick="showInfo(\'cc_info_'+index+'\');">'
    //                     +'<td class="overlay hidden" data-th="">'
    //                     +'<input type="hidden" id="cc_info_'+index+'" value="'+item.ConfigName+'|'+item.ProviderName+'|'+item.RegionName+'|'+item.CredentialName+'|'+item.DriverName+'"/>'
    //                     +'<input type="checkbox" name="chk" value="'+item.ConfigName+'" id="raw_'+index+'" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>'
    //                     +'<td class="btn_mtd ovm td_left" data-th="Name">'+item.ConfigName+'<a href="javascript:void(0);"><img src="/assets/img/contents/icon_copy.png" class="td_icon" alt=""/></a> <span class="ov"></span></td>'
    //                     +'<td class="overlay hidden" data-th="Cloud Provider">'+item.ProviderName+'</td>'
    //                     +'<td class="overlay hidden" data-th="Region">'+item.RegionName+'</td>'
    //                     +'<td class="overlay hidden" data-th="Zone">'+item.RegionName+'</td>'
    //                     +'<td class="overlay hidden" data-th="Credential">'+item.CredentialName+'</td>'
    //                     +'<td class="overlay hidden" data-th="Driver">'+item.DriverName+'</td>'
    //                     +'<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
    //                     +'</tr>' 
    //             ))
    //         }else{
    //             data.filter((list)=> list.ConfigName !== "" ).map((item,index)=>(
    //                 html +='<tr onclick="showInfo(\'cc_info_'+index+'\');">'
    //                     +'<td class="overlay hidden" data-th="">'
    //                     +'<input type="hidden" id="cc_info_'+index+'" value="'+item.ConfigName+'|'+item.ProviderName+'|'+item.RegionName+'|'+item.CredentialName+'|'+item.DriverName+'"/>'
    //                     +'<input type="checkbox" name="chk" value="'+item.ConfigName+'" id="raw_'+index+'" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>'
    //                     +'<td class="btn_mtd ovm td_left" data-th="Name">'+item.ConfigName+'<span class="ov"></span></td>'
    //                     +'<td class="overlay hidden" data-th="Cloud Provider">'+item.ProviderName+'</td>'
    //                     +'<td class="overlay hidden" data-th="Region">'+item.RegionName+'</td>'
    //                     +'<td class="overlay hidden" data-th="Zone">'+item.RegionName+'</td>'
    //                     +'<td class="overlay hidden" data-th="Credential">'+item.CredentialName+'</td>'
    //                     +'<td class="overlay hidden" data-th="Driver">'+item.DriverName+'</td>'
    //                     +'<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
    //                     +'</tr>'        
    //             ))

    //         }		
            
    //         $("#cList").empty();
    //         $("#cList").append(html);
    //     //  nsModal()
    //     ModalDetail()
        
    // }
//     })
}

function addCloudConnectionConfirm(){
    console.log("cloud connection 생성하겠냐는 물음")
    $(".dashboard.register_cont").toggleClass("active");
    $(".dashboard.server_status").removeClass("view");
    $(".dashboard .status_list tbody tr").removeClass("on");
    //ok 위치이동
    $('#RegistBox').on('hidden.bs.modal', function () {
        console.log(" ok 눌렀음");// cancel 눌러도 들어옴.
        var offset = $("#CreateBox").offset();
        $("#wrap").animate({scrollTop : offset.top}, 300);

        getRegionList()
        getCredentialList()
        getDriverList()
    })
}

function deleteCloudConnectionConfirm(){
    console.log("cloud connection 삭제하겠냐는 물음")
    deleteCloudConnection()
}

// connection 클릭시 나타남.
function showConnectionConfigInfo(target){
    console.log("target : ",target);
    var infos = $("#"+target).val()
    infos = infos.split("|")
    $("#info_name").val(infos[0])
    $("#info_provider").val(infos[1])
    $("#info_region").val(infos[2])
    $("#info_credential").val(infos[3])
    $("#info_driver").val(infos[4])

    $("#info_name_prs").text(infos[0])
    $("#info_provider_prs").text(infos[1])
    $("#info_region_prs").text(infos[2])
    $("#info_credential_prs").text(infos[3])
    $("#info_driver_prs").text(infos[4])

    getRegionDetail(infos[2]);

    $("#info_name").focus();
}

function goFocus(target){

    console.log(event)
    event.preventDefault()

    $("#"+target).focus();
    fnMove(target)

}

function getRegionInfo(type, target){

    console.log("getRegionInfo target : ",target);
    var infos = $("#"+target).val()
    infos = infos.split("|")
    console.log("name : ",infos[0]);
    $("#reg_region").val(infos[0]);
}

function getCredentialInfo(type, target){
console.log("getCredentialInfo target : ",target);
var infos = $("#"+target).val()
infos = infos.split("|")
console.log("name : ",infos[0]);
$("#reg_credential").val(infos[0]);
}
function getDriverInfo(type, target){
console.log("target : ",target);
var infos = $("#"+target).val()
infos = infos.split("|")
console.log("name : ",infos[0]);
$("#reg_driver").val(infos[0]);
}

// connection 정보 저장버튼 클릭
function createCloudConnection(){
    var configname = $("#reg_ConfigName").val()
    var providername = $("#reg_Provider").val()
    var regionname = $("#reg_region").val()
    var credentialname = $("#reg_credential").val()
    var drivername = $("#reg_driver").val()

    console.log("info param configname : ",configname);
    console.log("info param providername : ",providername);
    console.log("info param regionname : ",regionname);
    console.log("info param credentialname : ",credentialname);
    console.log("info param drivername : ",drivername);

    if(!configname){
        alert("Input New Cloud Connection Name")
        $("#reg_ConfigName").focus()
        return;
    }
    if(!providername){
        alert("Input Provider Name")
        $("#reg_Provider").focus()
        return;
    }
    
    var apiInfo = "{{ .apiInfo}}";
    var url = "{{ .comURL.SpiderURL}}"+"/connectionconfig";
    var obj = {
        ConfigName: configname,
        ProviderName: providername,
        RegionName: regionname,
        CredentialName: credentialname,
        DriverName : drivername
    }
    console.log("info connectionconfig obj Data : ",obj);
    if(configname){
        axios.post(url,obj,{
            headers: { 
                'Content-type': 'application/json',
                'Authorization': apiInfo, 
            }
        }).then(result =>{
            console.log(result);
            if(result.status == 200 || result.status == 201){
                alert("Success Create Cloud Connection")
                //등록하고 나서 화면을 그냥 고칠 것인가?
                getCloudConnectionList();
                //아니면 화면을 리로딩 시킬것인가?
                location.reload();
                // $("#btn_add2").click()
                // $("#namespace").val('')
                // $("#nsDesc").val('')
            }else{
                alert("Fail Create Cloud Connection")
            }
        });
    }else{
        alert("Input Cloud Connection Name")
        $("#reg_ConfigName").focus()
        return;
    }
}

function deleteCloudConnection(){

    var cnt = 0;
    var mcc_id = "";
    var apiInfo = ApiInfo;
    console.log("start cloudConnection_delete ")
    console.log("info chk : ", $(".chk"))
    $(".chk").each(function(){
        if($(this).is(":checked")){
            //alert("chk");
            cnt++;
            mcc_id = $(this).val();        
        }
        if(cnt < 1 ){
            alert("삭제할 대상을 선택해 주세요.");
            return;
        }

        if(cnt == 1){
        console.log("mcc_id ; ",mcc_id)
            //var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcc_id
            var url = "{{ .comURL.SpiderURL}}"+"/connectionconfig/"+mcc_id;
            
            if(confirm("삭제하시겠습니까?")){
            axios.delete(url,{
                headers :{
                    'Content-type': 'application/json',
                    'Authorization': apiInfo,
                    }
            }).then(result=>{
                var data = result.data
                if(result.status == 200){
                    alert(data.message)
                    location.reload(true)
                }
            })
            }
        }

        if(cnt >1){
            alert("한개씩만 삭제 가능합니다.")
            return;
        }

    })
}

// 선택한 resion의 상세 정보 조회
function getRegionDetail(target){
    var url = "/setting/connections/region/" + target;
    console.log("info urls : ", url);
    axios.get(url)
        .then(result=>{
            console.log("get Region Data : ",result.data);
            var data = result.data.Region;
            if(data.RegionName){ 
                console.log("info Region Detail, regionName : ",target,", region : ",data.KeyValueInfoList[0].Value, ", zone : ",data.KeyValueInfoList[1].Value)
                setRegionDispInfo(data.KeyValueInfoList[0].Value, data.KeyValueInfoList[1].Value)
            }
        }
    )
}

// Region에서 zone 정보까지 표시
function setRegionDispInfo(var_region, var_zone){
    var region_and_zone = ""
    var region_and_zone_code = ""
    region_and_zone = "" + var_region + " (" + var_zone + ")";
    region_and_zone_code = "" + var_region + " (" + var_zone + ")";
    $("#info_region").val(region_and_zone);// 상세내역에서 Region/zone 부분 
    $("#info_region_prs").text(region_and_zone_code);// 지도 위에 zone 표시영역 Region/zone 에서 zone

    console.log("info region_and_zone : ", region_and_zone);
}

function getRegionList(){
    // page load시 가져옴
    // var url = "{{ .comURL.SpiderURL}}"+"/region";
    // axios.get(url,{
    //     headers:{
    //         'Authorization': "{{ .apiInfo}}",
    //         'Content-Type' : "application/json"
    //     }
    // }).then(result=>{
    //         console.log("get Region Data : ",result.data);
    //         var data = result.data.region;
    //         var html = ""
    //         if(data.length){ 
    //                 data.filter((list)=> list.RegionName !== "" ).map((item,index)=>(
    //                     html +='<tr onclick="getRegionInfo(\'region\', \'region_info_'+index+'\');">'
    //                         +'<td class="btn_mtd ovm" data-th="Name">'+item.RegionName+'<span class="ov"></span></td>'
    //                         +'<input type="hidden" id="region_info_'+index+'" value="'+item.RegionName+'"/>'
    //                         +'<td class="overlay hidden" data-th="region ID">'+item.KeyValueInfoList[0].Value+'</td>'
    //                         +'<td class="overlay hidden" data-th="Zone">'+item.KeyValueInfoList[1].Value+'</td>'
    //                         +'<td class="overlay hidden" data-th="CP">'+item.ProviderName+'</td>'
    //                         +'</tr>'        
    //                 ))
                
    //             $("#regionList").empty();
    //             $("#regionList").append(html);
    //             $("#regionList tr").each(function(){
    //                 $selector = $(this)
                    
    //                 $selector.click(function(){
                        
    //                     if($(this).hasClass("on")){
    //                         $(this).removeClass("on");
    //                     }else{
    //                         $(this).addClass("on")
    //                     }
    //                 })
                    
    //             })
                    
    //             //  nsModal()
    //             ModalDetail()
    //         }
    //     }
    // )
}
function getCredentialList(){
// var url = "{{ .comURL.SpiderURL}}"+"/credential";
// axios.get(url,{
//     headers:{
//         'Authorization': "{{ .apiInfo}}",
//         'Content-Type' : "application/json"
//     }
// }).then(result=>{
//     console.log("get Credential Data : ",result.data);
//     var data = result.data.credential;
//     var html = ""
//     if(data.length){ 
//             data.filter((list)=> list.CredentialName !== "" ).map((item,index)=>(
//                 html +='<tr onclick="getCredentialInfo(\'credential\', \'credential_info_'+index+'\');">'
//                      +'<td class="btn_mtd ovm" data-th="Name">'+item.CredentialName+'<span class="ov"></span>'
//                      +'<input type="hidden" id="credential_info_'+index+'" value="'+item.CredentialName+'"/>'
//                      //  +'<td class="overlay hidden" data-th="accesskey">'+item.KeyValueInfoList[1].Value+'</td>'
//                      +'<td class="overlay hidden" data-th="accesskey">...</td>'
//                      +'<td class="overlay hidden" data-th="CP">'+item.ProviderName+'</td>'
//                      +'</tr>'        
//             ))
        
//         $("#credentialList").empty();
//         $("#credentialList").append(html);
//       //  nsModal()
//       ModalDetail()
//       $("#credentialList tr").each(function(){
//             $selector = $(this)
            
//             $selector.click(function(){
                
//                 if($(this).hasClass("on")){
//                     $(this).removeClass("on");
//                 }else{
//                     $(this).addClass("on")
//                 }
//             })
//     })
// }
// })
}
function getDriverList(){
// var url = "{{ .comURL.SpiderURL}}"+"/driver";
// axios.get(url,{
//     headers:{
//         'Authorization': "{{ .apiInfo}}",
//         'Content-Type' : "application/json"
//     }
// }).then(result=>{
//     console.log("get Driver Data : ",result.data);
//     var data = result.data.driver;
//     var html = ""
//     if(data.length){ 
//             data.filter((list)=> list.DriverName !== "" ).map((item,index)=>(
//                 html +='<tr onclick="getDriverInfo(\'driver\', \'driver_info_'+index+'\');">'
//                      +'<td class="btn_mtd ovm" data-th="Name">'+item.DriverName+'<span class="ov"></span></td>'
//                      +'<input type="hidden" id="driver_info_'+index+'" value="'+item.DriverName+'|'+item.ProviderName+'|'+item.DriverLibFileName+'"/>'
//                      +'<td class="overlay hidden" data-th="Driver SDK">'+item.DriverLibFileName+'</td>'
//                      +'<td class="overlay hidden" data-th="CP">'+item.ProviderName+'</td>'
//                      +'</tr>'        
//             ))
        
//         $("#driverList").empty();
//         $("#driverList").append(html);
//       //  nsModal()
//       ModalDetail()
//       $("#driverList tr").each(function(){
//             $selector = $(this)
            
//             $selector.click(function(){
                
//                 if($(this).hasClass("on")){
//                     $(this).removeClass("on");
//                 }else{
//                     $(this).addClass("on")
//                 }
//             })
//     })
// }}
// )
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


function ModalDetail2(){
    $(".dashboard .dataTable tbody tr").each(function(){
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

// connection 화면에서 팝업으로 region등록.
function saveNewRegion(){
    // valid check
    var regionName = $("#RegionMoalRegionName").val();
    var providerName = $("#RegionModalProviderName").val();
    var regionID = $("#RegionModalRegionID").val();
    var zoneID = $("#RegionModalZoneID").val();	
	
    if(!regionName || !providerName || !regionID){
        $("#required").modal()
        return;
    }
    //
    console.log("saveNewRegion popup");
    var regionInfo = {            
        RegionName:regionName,
        ProviderName: providerName,
        RegionKey: "Region",
        RegionValue: regionID,
        ZoneKey: "Zone",
        ZoneValue: zoneID        
    }
    console.log(req)
    axios.post(url,req,{


    axios.post(url,regionInfo,{
        
    }).then(result =>{
        console.log(result);
        if(result.status == 200 || result.status == 201){
            alert("Success Save Cloud Region");
            // 성공하면 내용 초기화
            $("#RegionMoalRegionName").val() = "";
            $("#RegionModalProviderName option:eq(0)").attr("selecte", "selected");
            $("#RegionModalRegionID").val() = "";
            $("#RegionModalZoneID").val() = "";	
            // TODO : region 목록 조회하여 Region table 갱신    
        }else{
            alert("Fail Create Cloud Region")
        }
    });
    
}