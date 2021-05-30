
$(document).ready(function(){
    order_type = "ConfigName"
    
    var defaultNameSpaceID = $('#topboxDefaultNameSpaceID').val();// Topbox에 기본 namespaceID를 set 함.
    if( defaultNameSpaceID == '' || defaultNameSpaceID == undefined){
        namespaceModalOkbtn();
        $("#popNameSpace").modal()
    }

    // css class 의 .btn_ok 에 대한 event를 따로 정의 함.
    // $('#AddBox .btn_ok.register').click(function(){
    // }
    
    // // defaultnamespace 확인 
    // defaultNamespace = '{{.LoginInfo.DefaultNameSpaceName}}'

    // if( defaultNamespace == ''){
    //     alert("Namespace를 선택바랍니다.")
    // }else{
    //     alert(defaultNamespace + "Dashboard로 이동.")
    // }

    //getNameSspaceList();// 이미 가져왔음

    // getCloudConnectionList(order_type);

    // getCloudOS("{{ .apiInfo}}",'RegionModalProviderName');// 이미 가져왔음
    // getCloudOS("{{ .apiInfo}}",'CredentialModalProviderName');// 이미 가져왔음
    // getCloudOS("{{ .apiInfo}}",'DriverModalProviderName');// 이미 가져왔음

    /* scroll */
    //checkbox all
    $("#th_chall").click(function() {
        if ($("#th_chall").prop("checked")) {
        $("input[name=chk]").prop("checked", true);
        } else {
        $("input[name=chk]").prop("checked", false);
        }
    })
        
    //table 스크롤바 제한
    $(window).on("load resize",function(){
        var vpwidth = $(window).width();
        if (vpwidth > 768 && vpwidth < 1800) {
        $(".dashboard_cont .dataTable").addClass("scrollbar-inner");
            $(".dataTable.scrollbar-inner").scrollbar();
        } else {
        $(".dashboard_cont .dataTable").removeClass("scrollbar-inner");
        }

        // Table 높이 조절, hidden인 상태인 Table은 show 될 때 set 하도록
        setTableHeightForScroll('connectionListTable', 300)
    });

    //Create popup - Region / Credential / Driver
    $(function() {
        return $(".modal").on("show.bs.modal", function() {
          var curModal;
          curModal = this;
          $(".modal").each(function() {
            if (this !== curModal) {
              console.log(".modal on show.bs.modal" + this + " : " + curModal);
              $(this).modal("hide");
            }
          });
        });
      });

})

// common.js에 정의 됨.
// function fnMove(target){
//     var offset = $("#" + target).offset();
//     console.log("fn move offset : ",offset)
//     $('html, body').animate({scrollTop : offset.top}, 400);
// }

// 현재생성된 connection config 목록. 
function getCloudConnectionList(sort_type){
    // 원래는 목록을 조회해서 filterling 하는 function
    // 이미 목록을 가져왔으므로 
    // TODO : 가져온 것을 filtering 하는 것으로 변경 필요
    // alert(sort_type + "만 filtering 하자.")
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

// 원래는 confirm창이었으나 입력form에 물어볼 필요 없으므로 config Area만 show
function addCloudConnectionConfirm(){
    // console.log("cloud connection 생성하겠냐는 물음")
    $(".dashboard.register_cont").toggleClass("active");
    $(".dashboard.server_status").removeClass("view");
    $(".dashboard .status_list tbody tr").removeClass("on");
    //ok 위치이동
    // $('#RegistBox').on('hidden.bs.modal', function () {
        console.log(" ok 눌렀음");// cancel 눌러도 들어옴.
        var offset = $("#CreateBox").offset();
        $("#wrap").animate({scrollTop : offset.top}, 300);        
    // })

    setTableHeightForScroll("regionListTable", 300);
    setTableHeightForScroll("credentialListTable", 300);
    setTableHeightForScroll("driverListTable", 300);
}

// connection 클릭시 나타남.
function showConnectionConfigInfo(target){
    console.log("target : ",target);
    var infos = $("#"+target).val()
    infos = infos.split("|")
    $("#infoName").val(infos[0])
    $("#infoProvider").val(infos[1])
    $("#infoRegion").val(infos[2])
    $("#infoCredential").val(infos[3])
    $("#infoDriver").val(infos[4])

    $("#infoNamePrs").text(infos[0])
    $("#infoProviderPrs").text(infos[1])
    $("#infoRegionPrs").text(infos[2])
    $("#infoCredentialPrs").text(infos[3])
    $("#infoDriverPrs").text(infos[4])

    getRegionDetail(infos[2]);

    $("#infoName").focus();
}

// function goFocus(target){

//     console.log(event)
//     event.preventDefault()

//     $("#"+target).focus();
//     fnMove(target)

// }

function getRegionInfo(type, target){

    console.log("getRegionInfo target : ",target);
    var infos = $("#"+target).val()
    infos = infos.split("|")
    console.log("name : ",infos[0]);
    $("#regRegion").val(infos[0]);
}

function getCredentialInfo(type, target){
console.log("getCredentialInfo target : ",target);
var infos = $("#"+target).val()
infos = infos.split("|")
console.log("name : ",infos[0]);
$("#regCredential").val(infos[0]);
}
function getDriverInfo(type, target){
console.log("target : ",target);
var infos = $("#"+target).val()
infos = infos.split("|")
console.log("name : ",infos[0]);
$("#regDriver").val(infos[0]);
}

// connection 정보 저장버튼 클릭
function createCloudConnection(){
    var configname = $("#regConfigName").val()
    var providername = $("#regProvider").val()
    var regionname = $("#regRegion").val()
    var credentialname = $("#regCredential").val()
    var drivername = $("#regDriver").val()

    console.log("info param configname : ",configname);
    console.log("info param providername : ",providername);
    console.log("info param regionname : ",regionname);
    console.log("info param credentialname : ",credentialname);
    console.log("info param drivername : ",drivername);

    if(!configname){
        alert("Input New Cloud Connection Name")
        $("#regConfigName").focus()
        return;
    }
    if(!providername){
        alert("Input Provider Name")
        $("#regProvider").focus()
        return;
    }
    
    var url = "/setting/connections/cloudconnectionconfig" + "/reg/proc";
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
        $("#regConfigName").focus()
        return;
    }
}

// confirm에서 Ok를 눌렀으므로 바로 처리
function deleteCloudConnection(){

    var cnt = 0;
    var mcc_id = "";
    // var apiInfo = ApiInfo;
    console.log("start cloudConnection_delete ")
    console.log("info chk : ", $(".chk"))
    $('input:checkbox[name="cloudconnection_chk"]').each(function(){
    // $(".chk").each(function(){
        if($(this).is(":checked")){
            //alert("chk");
            cnt++;
            mcc_id = $(this).val();        
        }
    })

        if(cnt < 1 ){
            commonAlert("삭제할 대상을 선택해 주세요.");
            return;
        }
        if(cnt >1){
            commonAlert("한개씩만 삭제 가능합니다.")
            return;
        }

        if(cnt == 1){
            console.log("mcc_id ; ",mcc_id)
            //var url = CommonURL+"/ns/"+NAMESPACE+"/mcis/"+mcc_id
            //var url = "{{ .comURL.SpiderURL}}"+"/connectionconfig/"+mcc_id;
            var url = "/setting/connections/cloudconnectionconfig" + "/del/"+mcc_id;
            // if(confirm("삭제하시겠습니까?")){
                axios.delete(url,{
                    headers :{
                        'Content-type': 'application/json',
                        // 'Authorization': apiInfo,
                        }
                }).then(result=>{
                    var data = result.data
                    console.log(data);        
                    if(result.status == 200){
                        commonAlert(data.message)
                        location.reload(true)
                    }
                // }).catch(function(error){
                //     console.log("connection delete error : ",error);        
                // });
                }).catch((error) => {
                    console.warn(error);
                    console.log(error.response)
                    var errorMessage = error.response.data.error;
                    commonErrorAlert(statusCode, errorMessage) 
                });
            // }
        }
}

// 선택한 resion의 상세 정보 조회
function getRegionDetail(target){
    var url = "/setting/connections/region/" + target;
    console.log("info urls : ", url);
    axios.get(url)
        .then(result=>{
            console.log("get Region Data : ",result.data);
            var data = result.data.Region;
            // if(data.RegionName){ 
            //     console.log("info Region Detail, regionName : ",target,", region : ",data.KeyValueInfoList[0].Value, ", zone : ",data.KeyValueInfoList[1].Value)
            //     setRegionDispInfo(data.KeyValueInfoList[0].Value, data.KeyValueInfoList[1].Value)
            // }
            if(data.RegionName){ 
            var keyValueInfoList = data.KeyValueInfoList;
                var regionID = "-"
                var zoneID = "-"
                // console.log("found: ", item)
                // console.log("found id: ", item.KeyValueInfoList)
                // console.log("found id: ", keyValueInfoList.length)
                if( keyValueInfoList.length == 1 ){
                    regionID = keyValueInfoList[0].Value;
                }else if ( keyValueInfoList.length == 2){
                    regionID = keyValueInfoList[0].Value
                    zoneID = keyValueInfoList[1].Value                    
                }
                console.log("info Region Detail, regionName : ",target,", region : ",regionID, ", zone : ",zoneID)
                setRegionDispInfo(regionID, zoneID)
            }
        // }).catch(function(error){
        //     console.log("region detail error : ",error);        
        // });
        }).catch((error) => {
            console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            commonErrorAlert(statusCode, errorMessage) 
        });
}

// Region에서 zone 정보까지 표시
function setRegionDispInfo(var_region, var_zone){
    var region_and_zone = ""
    var region_and_zone_code = ""
    region_and_zone = "" + var_region + " (" + var_zone + ")";
    region_and_zone_code = "" + var_region + " (" + var_zone + ")";
    $("#infoRegion").val(region_and_zone);// 상세내역에서 Region/zone 부분 
    $("#infoRegionPrs").text(region_and_zone_code);// 지도 위에 zone 표시영역 Region/zone 에서 zone

    console.log("info region_and_zone : ", region_and_zone);
}

// region 목록 : 저장 후 갱신용
function getRegionList(){
    console.log("getRegionList")
    var url = "/setting/connections/region"
    axios.get(url,{})
        .then(result=>{
            console.log("get Region Data : ",result.data);
            var data = result.data.Region;

            if(data.length){
                var html = ""
                data.forEach(function(item, index) {
                    var keyValueInfoList = item.KeyValueInfoList;
                    var regionID = "-"
                    var zoneID = "-"
                    // console.log("found: ", item)
                    // console.log("found id: ", item.KeyValueInfoList)
                    // console.log("found id: ", keyValueInfoList.length)
                    if( keyValueInfoList.length == 1 ){
                        regionID = keyValueInfoList[0].Value;
                    }else if ( keyValueInfoList.length == 2){
                        regionID = keyValueInfoList[0].Value
                        zoneID = keyValueInfoList[1].Value                    
                    }
                    
                    html +='<tr onclick="getRegionInfo(\'region\', \'region_info_'+index+'\');">'
                        +'<td class="btn_mtd ovm" data-th="Name">'+item.RegionName+'<span class="ov"></span></td>'
                        +'<input type="hidden" id="region_info_'+index+'" value="'+item.RegionName+'"/>'
                        +'<td class="overlay hidden" data-th="region ID">'+regionID+'</td>'
                        +'<td class="overlay hidden" data-th="Zone ID">'+zoneID+'</td>'
                        +'<td class="overlay hidden" data-th="CP">'+item.ProviderName+'</td>'
                        +'</tr>'    
                });
            
                $("#regionList").empty();
                $("#regionList").append(html);
                $("#regionList tr").each(function(){
                    $selector = $(this)
                    
                    $selector.click(function(){                        
                        if($(this).hasClass("on")){
                            $(this).removeClass("on");
                        }else{
                            $(this).addClass("on")
                        }
                    })
                    
                })
            }// end of data.length
        }
    // ).catch(function(error){
    //     console.log("region display error : ",error);        
    // });
    ).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        commonErrorAlert(statusCode, errorMessage) 
    });
}

function getCredentialList(){
    var url = "/setting/connections/credential";
    axios.get(url,{
    //     headers:{
    //         'Authorization': "", // TODO : token 넣을 것
    //         'Content-Type' : "application/json"
    //     }
    }).then(result=>{
        console.log("get Credential Data : ",result.data);
        var data = result.data.Credential;
        if(data.length){
            var html = ""
            data.forEach(function(item, index) {
                var keyValueInfoList = item.KeyValueInfoList;
                var key1 = "-"
                var value1 = "-"
                var key2 = "-"
                var value2 = "-"
                // console.log("found: ", item)
                // console.log("found id: ", item.KeyValueInfoList)
                // console.log("found id: ", keyValueInfoList.length)
                if( keyValueInfoList.length == 1 ){
                    key1 = keyValueInfoList[0].Key;
                    value1 = keyValueInfoList[0].Value;
                }else if ( keyValueInfoList.length == 2){
                    key1 = keyValueInfoList[0].Key;
                    value1 = keyValueInfoList[0].Value;
                    key2 = keyValueInfoList[1].Key;
                    value2 = keyValueInfoList[1].Value;          
                }
            
                html +='<tr onclick="getCredentialInfo(\'credential\', \'credential_info_'+index+'\');">'
                    +'<td class="btn_mtd ovm" data-th="Name">'+item.CredentialName+'<span class="ov"></span></td>'
                    +'<input type="hidden" id="credential_info_'+index+'" value="'+item.CredentialName+'"/>'
                    // +'<td class="overlay hidden" data-th="accesskey">'+value1+'</td>'
                    +'<td class="overlay hidden" data-th="accesskey">...</td>'
                    +'<td class="overlay hidden" data-th="CP">'+item.ProviderName+'</td>'
                    +'</tr>'
            });
        // if(data.length){ 
        //         data.filter((list)=> list.CredentialName !== "" ).map((item,index)=>(
        //             html +='<tr onclick="getCredentialInfo(\'credential\', \'credential_info_'+index+'\');">'
        //                 +'<td class="btn_mtd ovm" data-th="Name">'+item.CredentialName+'<span class="ov"></span>'
        //                 +'<input type="hidden" id="credential_info_'+index+'" value="'+item.CredentialName+'"/>'
        //                 //  +'<td class="overlay hidden" data-th="accesskey">'+item.KeyValueInfoList[1].Value+'</td>'
        //                 +'<td class="overlay hidden" data-th="accesskey">...</td>'
        //                 +'<td class="overlay hidden" data-th="CP">'+item.ProviderName+'</td>'
        //                 +'</tr>'        
        //         ))
            
            $("#credentialList").empty();
            $("#credentialList").append(html);
        // ModalDetail()
            $("#credentialList tr").each(function(){
                    $selector = $(this)
                    
                    $selector.click(function(){
                        
                        if($(this).hasClass("on")){
                            $(this).removeClass("on");
                        }else{
                            $(this).addClass("on")
                        }
                    })
            })
        }// end of data.length
    // }).catch(function(error){
    //     console.log("region display error : ",error);        
    // });
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        commonErrorAlert(statusCode, errorMessage) 
    });
}

function getDriverList(){
    var url = "/setting/connections"+"/driver";
    axios.get(url,{
        // headers:{
        //     'Authorization': "{{ .apiInfo}}",
        //     'Content-Type' : "application/json"
        // }
    }).then(result=>{
        console.log("get Driver Data : ",result.data);
        var data = result.data.Driver;
        var html = ""
        if(data.length){ 
                data.filter((list)=> list.DriverName !== "" ).map((item,index)=>(
                    html +='<tr onclick="getDriverInfo(\'driver\', \'driver_info_'+index+'\');">'
                        +'<td class="btn_mtd ovm" data-th="Name">'+item.DriverName+'<span class="ov"></span></td>'
                        +'<input type="hidden" id="driver_info_'+index+'" value="'+item.DriverName+'|'+item.ProviderName+'|'+item.DriverLibFileName+'"/>'
                        +'<td class="overlay hidden" data-th="Driver SDK">'+item.DriverLibFileName+'</td>'
                        +'<td class="overlay hidden" data-th="CP">'+item.ProviderName+'</td>'
                        +'</tr>'        
                ))
            
            $("#driverList").empty();
            $("#driverList").append(html);
        
            // ModalDetail()
            $("#driverList tr").each(function(){
                    $selector = $(this)
                    
                    $selector.click(function(){
                        
                        if($(this).hasClass("on")){
                            $(this).removeClass("on");
                        }else{
                            $(this).addClass("on")
                        }
                    })
            })
        }
    // }).catch(function(error){
    //     console.log("region display error : ",error);        
    // });
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage) 
    });
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
    var regionName = $("#RegionModalRegionName").val();
    var providerName = $("#RegionModalProviderName").val();
    var regionID = $("#RegionModalRegionID").val();
    var zoneID = $("#RegionModalZoneID").val();	
	
    if(!regionName || !providerName || !regionID){
        $("#modalRegionRequired").modal()// TODO : requiredCloudConnection 로 바꿔 공통으로 쓸까?
        return;
    }
    //
    console.log("saveNewRegion popup");
    var regionInfo = {            
        RegionName:regionName,
        ProviderName: providerName,
        KeyValueInfoList:[ {"Key":"Region","Value":regionID},{"Key":"Zone","Value":zoneID}]

        // RegionKey: "Region",
        // RegionValue: regionID,
        // ZoneKey: "Zone",
        // ZoneValue: zoneID        
    }
    console.log(regionInfo)
    axios.post("/setting/connections/region/reg/proc",regionInfo,{
        
    }).then(result =>{
        console.log(result);
        if(result.status == 200 || result.status == 201){
            alert("Success Save Cloud Region");
            // 성공하면 내용 초기화
            $("#RegionModalRegionName").val('');
            $("#RegionModalProviderName option:eq(0)").attr("selecte", "selected");
            $("#RegionModalRegionID").val('');
            $("#RegionModalZoneID").val('');
            // Region table 갱신
            getRegionList();
        }else{
            alert("Fail Create Cloud Region")
        }
  
    // }).catch(function(error){
    //     console.log("save error : ",error);    
    // });
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage) 
    });
}

// region 삭제
function deleteRegion(){
    var regionName = $("#regRegion").val()

    if(!regionName){
        $("#requireMessage").text("선택된 Resion이 없습니다.")
        $("#requiredCloudConnection").modal()
        return;
    }
    $("#requireMessage").text("")
    
    alert(regionName + " 을 삭제하겠습니까");//TODO : confirm으로 바꿔야 함.
    
    var url = "/setting/connections/region/del/"+regionName;
    //axios.post("/setting/connections/region/reg/proc",regionInfo,{
    axios.delete(url, {},{
        }).then(result =>{
            console.log(result);
            if(result.status == 200 || result.status == 201){
                alert("Deleted Cloud Region");                
                // Region 갱신 
                getRegionList();   
            }else{
                alert("Fail to delete the Cloud Region")
            }
        // }).catch(function(error){
        //     console.log("delete error : ",error);        
        // });
        }).catch((error) => {
            console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage) 
        });
}

//////
// connection 화면에서 팝업으로 Credential등록.
function saveNewCredential(){
    // valid check
    var credentialName = $("#CredentialModalCredentialName").val();
    var providerName = $("#CredentialModalProviderName").val();
    var key0 = $("#CredentialModalKey0").val();
    var value0 = $("#CredentialModalValue0").val();
    var key1 = $("#CredentialModalKey1").val();
    var value1 = $("#CredentialModalValue1").val();	
    var key2 = $("#CredentialModalKey2").val();
    var value2 = $("#CredentialModalValue2").val();	
	// CredentialName string             `json:"CredentialName"`
	// ProviderName   string             `json:"ProviderName"`
	// KeyValueInfoList   []KeyValueInfoList `json:"KeyValueInfoList"`
    if(!credentialName || !providerName || !key0 || !value0 || !key1 || !value1 ){
        $("#modalCredentialRequired").modal()// TODO : requiredCloudConnection 로 바꿔 공통으로 쓸까?
        return;
    }

    var credentialInfo = "";
    // provider에 따라 사용하는 key가 불규칙적임.
    
    if( providerName == "GCP"){// gcp는 Key가 3개
        credentialInfo = {            
            CredentialName:credentialName,
            ProviderName: providerName,
            KeyValueInfoList:[ {"Key":key0,"Value":value0},{"Key":key1,"Value":value1},{"Key":key2,"Value":value2}]
        }
    }else{
        credentialInfo = {            
            CredentialName:credentialName,
            ProviderName: providerName,
            KeyValueInfoList:[ {"Key":key0,"Value":value0},{"Key":key1,"Value":value1}]
        }
    }
    //
    console.log("saveNewCredential popup");
     
    console.log(credentialInfo)
    axios.post("/setting/connections/credential/reg/proc",credentialInfo,{
        // headers: { 'content-type': 'application/x-www-form-urlencoded' },
    }).then(result =>{
        console.log(result);
        if(result.status == 200 || result.status == 201){
            alert("Success Save Cloud Credential");
            // 성공하면 내용 초기화 : provider가 같으면 key0, key1 은 그대로 사용
            $("#CredentialModalCredentialName").val('');
            // $("#CredentialModalProviderName option:eq(0)").attr("selecte", "selected");
            $("#CredentialModalKey0").val('');
            $("#CredentialModalValue0").val('');
            $("#CredentialModalKey1").val('');
            $("#CredentialModalValue1").val('');
            $("#CredentialModalKey2").val('');
            $("#CredentialModalValue2").val('');
            
            // Credential table 갱신
            getCredentialList();
        }else{
            alert("Fail Create Cloud Credential")
        }
  
    // }).catch(function(error){
    //     console.log("save error : ",error);    
    // });
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage) 
    });
}

// Credential 삭제
function deleteCredential(){
    var credentialName = $("#regCredential").val()

    if(!credentialName){
        $("#requireMessage").text("선택된 Credential key가 없습니다.")
        $("#requiredCloudConnection").modal()
        return;
    }
    $("#requireMessage").text("")
    
    alert(credentialName + " 을 삭제하겠습니까");//TODO : confirm으로 바꿔야 함.
    
    var url = "/setting/connections/credential/del/"+credentialName;
    axios.delete(url, {},{
        }).then(result =>{
            console.log(result);
            if(result.status == 200 || result.status == 201){
                alert("Deleted Cloud Credential");                
                // Credential 갱신 
                getCredentialList();   
            }else{
                alert("Fail to delete the Cloud Credential")
            }
        // }).catch(function(error){
        //     console.log("delete error : ",error);        
        // });
        }).catch((error) => {
            console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage) 
        });
}

//////
// Driver 등록. connection 화면에서 팝업으로 
function saveNewDriver(){
    // valid check
    var driverlName = $("#DriverModalDriverName").val();
    var providerName = $("#DriverModalProviderName").val();
    var driverLibFilename = $("#DriverModalDriverLibFileName").val();
	
    if(!driverlName || !providerName || !driverLibFilename ){
        $("#modalDriverRequired").modal()// TODO : requiredCloudConnection 로 바꿔 공통으로 쓸까?
        return;
    }
    //
    console.log("saveNewCredential popup");
    // provider에 따라 사용하는 key가 불규칙적임.
    var driverlNameInfo = {            
        DriverName:driverlName,
        ProviderName: providerName,
        DriverLibFileName:driverLibFilename
    }
    console.log(driverlNameInfo)
    axios.post("/setting/connections/driver/reg/proc",driverlNameInfo,{
        
    }).then(result =>{
        console.log(result);
        if(result.status == 200 || result.status == 201){
            alert("Success Save Cloud Driver");
            // 성공하면 내용 초기화
            $("#DriverModalDriverName").val('');
            // $("#DriverModalProviderName option:eq(0)").attr("selecte", "selected");
            $("#DriverModalDriverLibFileName").val('');
            
            // Driver table 갱신
            getDriverList();
        }else{
            alert("Fail Create Cloud Driver")
        }
  
    // }).catch(function(error){
    //     console.log("save error : ",error);    
    // });
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        commonErrorAlert(statusCode, errorMessage) 
    });
}

// Driver 삭제
function deleteDriver(){
    var driverName = $("#regDriver").val()

    if(!driverName){
        $("#requireMessage").text("선택된 Driver가 없습니다.")
        $("#requiredCloudConnection").modal()
        return;
    }
    $("#requireMessage").text("")
    
    alert(driverName + " 을 삭제하겠습니까");//TODO : confirm으로 바꿔야 함.
    
    var url = "/setting/connections/driver/del/"+driverName;
    axios.delete(url, {},{
        }).then(result =>{
            console.log(result);
            if(result.status == 200 || result.status == 201){
                alert("Deleted Cloud Driver");                
                // Driver 갱신 
                getDriverList();   
            }else{
                alert("Fail to delete the Cloud Driver")
            }
        // }).catch(function(error){
        //     console.log("delete error : ",error);        
        // });
        }).catch((error) => {
            console.warn(error);
            console.log(error.response)
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage) 
        });
}

