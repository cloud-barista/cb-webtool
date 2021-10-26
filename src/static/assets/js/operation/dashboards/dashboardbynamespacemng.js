$(document).ready(function () {
    jQuery('.sc_box.scrollbar-inner').scrollbar();

    getCommonCloudConnectionList("dashboard", true)

    getCommonMcisList("dashboard", true, "", "status")
})

// CloudConnectionList가져온 결과를 set
function getCloudConnectionListCallbackSuccess(caller, connectionConfigList, sortType) {
    var totalProviderCount = 0;
    var totalConnectionConfigCount = 0;
    var providerConnectionMap = new Map();
    if (!isEmpty(connectionConfigList) && connectionConfigList.length) {
        totalConnectionConfigCount = connectionConfigList.length;

        var providerArr = new Array();
        for (var itemIndex in connectionConfigList) {
            var aConnectionConfig = connectionConfigList[itemIndex]
            // console.log(aConnectionConfig)
            if (providerConnectionMap.has(aConnectionConfig.ProviderName)) {
                providerConnectionMap.set(aConnectionConfig.ProviderName, providerConnectionMap.get(aConnectionConfig.ProviderName) + 1)
            } else {
                providerConnectionMap.set(aConnectionConfig.ProviderName, 1)
            }

        }
        totalProviderCount = providerArr.length
    }
    // console.log(totalConnectionConfigCount + " : " + providerConnectionMap.size)
    $("#connectionCount").text(totalConnectionConfigCount)
    $("#providerCount").text(providerConnectionMap.size)


    for (let item of providerConnectionMap) {
        console.log(item[0] + ' , ' + item[1]);
        $("#cpConnectionDetail").append('<li class="bg_etc bg_' + item[0].toLowerCase() + '"><a href="javascript:void(0);"><span class="conn_cnt">' + item[1] + '<div class="conn_tit">' + item[0] + '</div></span></a></li>');
    }
}

// MCIS 목록 조회 후 화면에 Set
function getMcisListCallbackSuccess(caller, mcisList) {

    // MCIS Status
    var totalMcisCnt = 0;
    var mcisStatusCountMap = new Map();
    mcisStatusCountMap.set("running", 0);
    mcisStatusCountMap.set("stopped", 0);  // partial 도 stop으로 보고있음.
    mcisStatusCountMap.set("terminated", 0);
    mcisStatusCountMap.set("total", 0);

    var totalServerCnt = 0;
    var totalVmStatusCountMap = new Map();
    totalVmStatusCountMap.set("running", 0);
    totalVmStatusCountMap.set("stopped", 0);  // partial 도 stop으로 보고있음.
    totalVmStatusCountMap.set("terminated", 0);
    totalVmStatusCountMap.set("total", 0);

    if (!isEmpty(mcisList) && mcisList.length > 0) {
        //totalMcisCnt = mcisList.length;
        var addMcis = "";
        var addVm = "";
        for (var mcisIndex in mcisList) {
            var aMcis = mcisList[mcisIndex]
            var mcisStatus = aMcis.status

            var mcisProviderNames = "";//MCIS에 사용 된 provider
            var totalVmCountOfMcis = 0;//MCIS의 VM 갯 수
            var mcisDispStatus = getMcisStatusDisp(mcisStatus);// 화면 표시용 status
            // mcis status
            try {
                // console.log(aMcis)
                if (mcisStatus != "") {// mcis status 가 없는 경우는 skip
                    if (mcisStatusCountMap.has(mcisDispStatus)) {
                        mcisStatusCountMap.set(mcisDispStatus, mcisStatusCountMap.get(mcisDispStatus) + 1)
                    }
                    totalMcisCnt++;
                } else {
                    continue;// status가 없으면 mcks 일 수 있으므로 mcis에서는 count 제외
                }
            } catch (e) {
                console.log("mcis status error")
            }

            // vm status
            try {
                var vmListOfMcis = aMcis.vm;// array


                var vmStatusCountMap = new Map();
                vmStatusCountMap.set("running", 0);
                vmStatusCountMap.set("stopped", 0);  // partial 도 stop으로 보고있음.
                vmStatusCountMap.set("terminated", 0);
                vmStatusCountMap.set("total", 0);

                var vmCloudConnectionMap = new Map();
                console.log(vmListOfMcis)
                if (typeof vmListOfMcis !== 'undefined' && vmListOfMcis.length > 0) {
                    for (var vmIndex in vmListOfMcis) {
                        var aVm = vmListOfMcis[vmIndex];
                        var vmDispStatus = getVmStatusDisp(aVm.status);
                        console.log("vmDispStatus:", vmDispStatus);
                        totalVmCountOfMcis++;
                        console.log("vmStatus " + aVm.status + " , dispVmStatus " + vmDispStatus)
                        if (vmStatusCountMap.has(vmDispStatus)) {
                            vmStatusCountMap.set(vmDispStatus, vmStatusCountMap.get(vmDispStatus) + 1)// mcis내 count
                            totalVmStatusCountMap.set(vmDispStatus, totalVmStatusCountMap.get(vmDispStatus) + 1)// 전체 count
                        }
                        vmStatusCountMap.set("total", vmStatusCountMap.get("total") + 1)// mcis내 count
                        totalVmStatusCountMap.set("total", totalVmStatusCountMap.get("total") + 1)
                        totalServerCnt++;

                        // connections
                        var location = aVm.location;
                        if (!isEmpty(location)) {
                            var cloudType = location.cloudType;
                            if (vmCloudConnectionMap.has(cloudType)) {
                                vmCloudConnectionMap.set(cloudType, vmCloudConnectionMap.get(cloudType) + 1)
                            } else {
                                vmCloudConnectionMap.set(cloudType, 0)
                            }
                        }

                        // vm항목 미리 생성 후 mcis 생성할 때 붙임
                        addVm += '<div class="shot bgbox_' + vmDispStatus + '">'
                        addVm += '    <a href="javascript:void(0);"><span>' + (Number(vmIndex) + 1).toString() + '</span></a>'
                        // for map
                        addVm += '        <input type="hidden" name="vmID" id="vmID_' + vmIndex + '" value="' + aVm.vmID + '"/>'
                        addVm += '        <input type="hidden" name="vmName" id="vmName_' + vmIndex + '" value="' + aVm.vmName + '"/>'
                        addVm += '        <input type="hidden" name="vmStatus" id="vmStatus_' + vmIndex + '" value="' + vmDispStatus + '"/>'
                        addVm += '        <input type="hidden" name="longitude" id="longitude_' + vmIndex + '" value="' + location.longitude + '"/>'
                        addVm += '        <input type="hidden" name="latitude" id="latitude_' + vmIndex + '" value="' + location.latitude + '"/>'
                        addVm += '</div>'

                    }
                }// end of vm list

                // console.log(vmCloudConnectionMap);
                vmCloudConnectionMap.forEach((value, key) => {
                    mcisProviderNames += key + " ";
                });
                console.log("mcisProviderNames=" + mcisProviderNames);
            } catch (e) {
                console.log("vm status error")
            }



            // List of Mcis table
            try {

                addMcis += '    <div class="areabox dbinfo cursor" id="mcis_areabox_' + mcisIndex + '" onclick="selectMcis(\'' + aMcis.id + '\',\'' + aMcis.name + '\',\'mcis_areabox_' + mcisIndex + '\', this)">'
                addMcis += '        <div class="box">'
                addMcis += '            <div class="top">'
                addMcis += '                <div class="txtbox">'
                addMcis += '                    <div class="tit">' + aMcis.name + '</div>'
                addMcis += '                    <div class="txt"><span class="bgbox_b"></span>Available 01</div>'
                addMcis += '                </div>'

                addMcis += '                <div class="state color_' + mcisDispStatus + '"></div>'
                addMcis += '            </div>'

                addMcis += '            <div class="numbox">'
                addMcis += '                infra <strong class="color_b">' + vmStatusCountMap.get("total") + '</strong>'
                addMcis += '                <span class="line">(</span> <span class="num color_b">' + vmStatusCountMap.get("running") + '</span>'
                addMcis += '                <span class="line">/</span> <span class="num color_y">' + vmStatusCountMap.get("stopped") + '</span>'
                addMcis += '                <span class="line">/</span> <span class="num color_r">' + vmStatusCountMap.get("terminated") + '</span>'
                addMcis += '                <span class="line">)</span>'
                addMcis += '            </div>'

                // 이 항목은 크게 의미가 없는데??
                addMcis += '            <div class="numinfo">'
                addMcis += '                <div class="num">server ' + vmStatusCountMap.get("total") + '</div>'
                addMcis += '            </div>'

                addMcis += '            <div class="shotbox">'
                // 각 vm 의 항목들
                addMcis += addVm
                addMcis += '            </div>'

                addMcis += '        </div>'
                addMcis += '    </div>'

            } catch (e) {
                console.log("list of mcis error")
                console.log(e)
            }
            addVm = "";//
        }// end of mcis loop

        $("#mcisList").empty();
        $("#mcisList").append(addMcis);

        $("#total_mcis").text(totalMcisCnt);
        $("#mcis_status_running").text(mcisStatusCountMap.get("running"));
        $("#mcis_status_stopped").text(mcisStatusCountMap.get("stop"));
        $("#mcis_status_terminated").text(mcisStatusCountMap.get("terminate"));

        $("#total_vm").text(totalServerCnt);
        $("#vm_status_running").text(totalVmStatusCountMap.get("running"));
        $("#vm_status_stopped").text(totalVmStatusCountMap.get("stop"));
        $("#vm_status_terminated").text(totalVmStatusCountMap.get("terminate"));
    } else {
        var addMcis = "";
        addMcis += '<tr>'
        addMcis += '<td class="overlay hidden" data-th="" colspan="8">No Data</td>'
        addMcis += '</tr>'
        $("#mcisList").empty();
        $("#mcisList").append(addMcis);
    }
}

// 조회 실패시.
function getMcisListCallbackFail(caller, error) {
    // List table에 no data 표시? 또는 조회 오류를 표시?
    var addMcis = "";

    $("#mcisList").empty();
    $("#mcisList").append(addMcis);
}

// dashboard 의 MCIS 목록에서 mcis 선택 : 색상반전, 선택한 mcis id set -> status변경에 사용
// 1번클릭시 선택
// 2번 클릭 시 해당 MCIS로 이동
function selectMcis(id, name, target, obj) {
    console.log("selectMcis")
    var mcis_id = id
    var mcis_name = name
    var init_select_areabox = $("#init_select_areabox").val()
    $target = $("#" + target)

    if ($target.hasClass("active")) {
        // location.href = "/Manage/MCIS/list/" + mcis_id + "/" + mcis_name
        location.href = "/operation/manages/mcismng/mngform"
        return;
    } else {
        $target.addClass("active")
    }

    // $("[id^='mcis_areabox_']").each(function(){
    //     var s_id = $(this).attr("id");
    //     console.log(s_id + ":" + target)
    //     if(s_id == target){
    //         try{
    //             var s_id = $(this).attr("id");
    //             $(this).addClass("active"); 
    //             console.log(s_id + " addClass active")
    //         }catch(e){
    //             console.log(e)
    //         }

    //     }else{
    //         $(this).removeClass("active");
    //         // console.log(s_id + "removeClass active")
    //     }
    // })
    // console.log(" active / deactive ")
    $("#mcis_id").val(mcis_id)
    $("#mcis_name").val(mcis_name)
    console.log(" mcis_id =" + mcis_id + ", mcis_name = " + mcis_name);
}

// callMcisLifeCycle -> McisLifeCycle -> callbackMcisLifeCycle
// confirm창을 띄울 때 mcismng와 동일한 key로 호출하므로 callback함수 이름도 같아야 한다.(util.js 참조)
function callMcisLifeCycle(type) {
    var selectedCount = 0;
    // 선택된 mcis 가 있는지 체크.
    $("[id^='mcis_areabox_']").each(function () {
        if ($(this).hasClass("active")) {
            selectedCount++
            mcisLifeCycle($("#mcis_id").val(), type);//mcislifecycle.js 호출
        }
    })

    if (selectedCount == 0) {
        commonAlert("Please Select MCIS!!")
    }

    /////// TODO : util.mcislifecycle.js 를 호출하도록 변경

}

// McisLifeCycle을 호출 한 뒤 return값 처리
function callbackMcisLifeCycle(resultStatus, resultData, type) {
    var message = "MCIS " + type + " complete!."
    if (resultStatus == 200 || resultStatus == 201) {
        commonAlert(message);
        location.reload();//완료 후 페이지를 reload -> 해당 mcis만 reload
        // 해당 mcis 조회
        // 상태 count 재설정
    }
}

function setMap() {
    //show_mcis2(url,JZMap);
    //function show_mcis2(url, map){
    // var JZMap = map;
    var JZMap = map_init()// TODO : map click할 때 feature 에 id가 없어 tooltip 에러나고 있음. 해결필요 

    //지도 그리기 관련
    var polyArr = new Array();

    // $("[id^='vmID_']").each(function(){
    $("input[name=vmID]").each(function (vmIndex, item) {
        // var vmID = $(this).attr("id");
        // var vmIndex = vmID.split ("_")[1];
        var vmIDValue = $("#vmID_" + vmIndex).val();
        var vmNameValue = $("#vmName_" + vmIndex).val();
        var vmStatusValue = $("#vmStatus_" + vmIndex).val();
        var longitudeValue = $("#longitude_" + vmIndex).val();
        var latitudeValue = $("#latitude_" + vmIndex).val();

        var vms = new Object();
        vms.id = vmIDValue;
        vms.name = vmNameValue;
        vms.longitudeValue = longitudeValue;
        vms.latitudeValue = latitudeValue;
        // vms.status = vmStatusValue;
        // vms.status = vmStatusValue;

        var fromLonLat = longitudeValue + " " + latitudeValue;
        console.log(longitudeValue + " : " + latitudeValue);
        if (longitudeValue && latitudeValue) {
            // polyArr.push(fromLonLat)
            drawMap(JZMap, longitudeValue, latitudeValue, vms)

            var polygon = "POLYGON((" + fromLonLat + "))";
            // drawPoligon(JZMap,fromLonLat);
            drawPoligon(JZMap, polygon);
        }
        // for(var i in mcis){
        //     for(var o in vms){
        //         vm_cnt++;
        //         var vm_status = vms[o].status
        //         var lat = vms[0].location.latitude
        //         var long = vms[0].location.longitude
        //         var provider = vms[0].location.cloudType

        //         // console.log("info : ",info)
        //         // point_feature.set('title',info.name)
        //         // point_feature.set('vm_status',info.status)
        //         // point_feature.set('vm_id',info.id)
        //         // point_feature.set('id',info.id)

        //         var fromLonLat = long+" "+lat;
        //         if(long && lat){
        //             polyArr.push(fromLonLat)
        //             drawMap(JZMap,long,lat,vms[o])
        //         }

        //         var polygon = "";
        //          console.log("poly arr : ",polyArr);
        //          if(polyArr.length > 1){
        //            polygon = polyArr.join(", ")
        //            polygon = "POLYGON(("+polygon+"))";
        //          }else{
        //            polygon = "POLYGON(("+polyArr[0]+"))";
        //          }
        //          if(polyArr.length >1){
        //             drawPoligon(JZMap,polygon);
        //           }
        //     }
    })
}