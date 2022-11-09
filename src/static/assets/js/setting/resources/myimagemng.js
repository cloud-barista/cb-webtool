$(document).ready(function () {
    order_type = "name"
    //checkbox all
    $("#th_chall").click(function () {
        if ($("#th_chall").prop("checked")) {
            $("input[name=chk]").prop("checked", true);
        } else {
            $("input[name=chk]").prop("checked", false);
        }
    })

    //table 스크롤바 제한
    $(window).on("load resize", function () {
        var vpwidth = $(window).width();
        if (vpwidth > 768 && vpwidth < 1800) {
            $(".dashboard_cont .dataTable").addClass("scrollbar-inner");
            $(".dataTable.scrollbar-inner").scrollbar();
        } else {
            $(".dashboard_cont .dataTable").removeClass("scrollbar-inner");
        }

        setTableHeightForScroll('myImageListTable', 300);
    });

    getMyImageList(order_type);
});

// function goDelete() {
function deleteDataDisk() {
    var dataDiskId = "";
    var count = 0;
    var diskStatus = "";
    $("input[name='chk']:checked").each(function () {
        count++;
        dataDiskId = dataDiskId + $(this).val() + ",";
        diskStatus = $(this).attr("diskStatus")
        if (diskStatus == "attached") {
            commonAlert("Disk에 할당된 VM이 있어 삭제 할 수 없습니다.");
            return false;
        }
    });

    if (diskStatus == "attached") {
        var checkBox = document.getElementsByName("chk");
        checkBox.forEach(item => {
            item.checked = false;
        })
        return false;
    }

    dataDiskId = dataDiskId.substring(0, dataDiskId.lastIndexOf(","));

    console.log("dataDiskId : ", dataDiskId);
    console.log("count : ", count);

    if (dataDiskId == '') {
        commonAlert("삭제할 대상을 선택하세요.");
        return false;
    }

    if (count != 1) {
        commonAlert("삭제할 대상을 하나만 선택하세요.");
        return false;
    }

    var url = "/setting/resources" + "/datadisk/del/" + dataDiskId
    console.log("del dataDisk url : ", url);

    axios.delete(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        var data = result.data;
        console.log(result);
        console.log(data);
        if (result.status == 200 || result.status == 201) {
            commonAlert(data.message)
            displayDataDiskInfo("DEL_SUCCESS")

        } else {
            commonAlert(result.data.error)
        }
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

function deleteMyImageDisk() {
    var imageId = "";
    var count = 0;
    var chk_length = $("input[name='chk']:checked").length
    if (chk_length > 1) {
        alert("한개만 삭제할 수 있습니다.")
        return;
    } else if (chk_length == 1) {
        $("input[name='chk']:checked").each(function () {
            count++;
            imageId = $(this).val();
            console.log("image ID : ", imageId);
            var url = "/setting/resources/myimage/del/" + imageId;
            console.log("del dataDisk url : ", url);
            axios.delete(url, {
                headers: {
                    'Content-Type': "application/json"
                }
            }).then(result => {
                var data = result.data;
                console.log(result);
                console.log(data);
                if (result.status == 200 || result.status == 201) {
                    commonAlert(data.message)
                    getMyImageList("name")
                    $("#myImageInfoBox").hide();

                } else {
                    commonAlert(result.data.error)
                }
            }).catch((error) => {
                console.warn(error);
                console.log(error.response)
                var errorMessage = error.response.data.error;
                var statusCode = error.response.status;
                commonErrorAlert(statusCode, errorMessage);
            });

        });
    }
}

function deleteMyImage() {
    var myImageId = $("#dtlMyImageName").val()
    var url = "/setting/resources/myimage/del/" + myImageId
    console.log("del dataDisk url : ", url);

    axios.delete(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        var data = result.data;
        console.log(result);
        console.log(data);
        if (result.status == 200 || result.status == 201) {
            commonAlert(data.message)
            getMyImageList("name")
            $("#myImageInfoBox").hide();
        } else {
            commonAlert(result.data.error)
        }
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

function chkDiskStatus(attr) {

}
function attachDataDisk() {
    var dataDiskId = "";
    var count = 0;
    var connectionName = [];
    var connectionDup = false;
    var selectDiskName = [];
    var chkDiskId = [];
    var diskStatus = "";
    $("input[name='chk']:checked").each(function () {
        count++;
        dataDiskId = dataDiskId + $(this).val() + ",";
        var tempConnectionName = $(this).attr("item");
        var tempDiskName = $(this).attr("diskName");
        var tempDiskId = $(this).val();
        diskStatus = $(this).attr("diskStatus")
        if (diskStatus == "attached") {
            commonAlert("이미 Attach된 Disk 를 선택하셨습니다.");
            return false;
        }
        connectionName.push(tempConnectionName);
        selectDiskName.push(tempDiskName);
        chkDiskId.push(tempDiskId);

    });
    if (diskStatus == "attached") {
        var checkBox = document.getElementsByName("chk");
        checkBox.forEach(item => {
            item.checked = false;
        })
        return false;
    }
    console.log("connectionName arr :", connectionName)
    console.log("selectDiskName arr :", selectDiskName)
    // 선택된 디스크 이름 가져와서 뿌려주기
    console.log("datadiskID : ", dataDiskId);
    var selectDisk = selectDiskName.join(",");
    var selectDiskId = chkDiskId.join(",")
    $("#selected_disk").val(selectDisk);
    $("#selected_disk_id").val(selectDiskId);
    //connection  다를경우 걸르기
    connectionName.forEach((item) => {
        if (connectionName[0] == item) {

        } else {
            connectionDup = true
        }
    })

    if (connectionDup) {
        commonAlert("같은 Connection Name을 선택하세요.")
        return false;
    } else {
        connectionName = connectionName[0];
    }

    dataDiskId = dataDiskId.substring(0, dataDiskId.lastIndexOf(","));

    console.log("dataDiskId : ", dataDiskId);
    console.log("count : ", count);

    if (dataDiskId == '') {
        commonAlert("Attach할 대상을 선택하세요.");
        return false;
    }

    getCommonMcisList("dataDiskMng", true, "", "", "connection=" + connectionName)

}

function runDetachDataDisk(command) {
    var count = 0;
    var diskStatus = "";
    $("input[name='chk']:checked").each(function (index) {
        console.log("detach index :", index);
        count++;
        diskStatus = $(this).attr("diskstatus")
        if (diskStatus != "attached") {
            commonAlert("Disk에 할당된 VM이 없어 Detach 할 수 없습니다.");
            return false;
        } else {
            var mcis_id = $(this).attr("mcis_id");
            var vm_id = $(this).attr("vm_id");
            var dataDiskId = $(this).val();
            var url = "/operation/manages/mcismng/" + mcis_id + "/vm/" + vm_id + "/datadisk?option=" + command;
            console.log("datach url : ", url)
            var obj = {
                dataDiskId
            }
            axios.put(url, obj).then(result => {
                var data = result.data;
                console.log(data);
                if (data.status == 200 || data.status == 201) {
                    if (index == count - 1) {
                        commonAlert("Success Detach DataDisk!")
                        $("#dataDiskInfoBox").hide();
                        // displayDataDiskInfo("MODIFY_SUCCESS");
                        location.reload()
                    }

                } else {
                    commonAlert("Fail Detach DataDisk at " + item + data.message)
                    showMyImageInfo(diskId, diskName);
                }
            }).catch(error => {
                console.log(error.response);
            })
        }

    });
    console.log("disk status at run detach : ", diskStatus)
    if (diskStatus != "attached") {
        var checkBox = document.getElementsByName("chk");
        checkBox.forEach(item => {
            item.checked = false;
        })
        return false;
    }
}
function runAttachDataDisk(command) {
    var mcis_id = $("#attach_mcis_id").val();
    var vm_id = $("#attach_vm_id").val();
    var dataDiskId = $("#selected_disk").val()
    var diskId = dataDiskId.split(",");
    console.log("command : ", command);
    var count = diskId.length;
    console.log("count : ", count);
    diskId.forEach((item, index) => {
        var url = "/operation/manages/mcismng/" + mcis_id + "/vm/" + vm_id + "/datadisk?option=" + command;
        console.log("attach url : ", url)
        console.log("attach diskid : ", item);
        var obj = {
            dataDiskId: item
        }
        axios.put(url, obj).then(result => {
            var data = result.data;
            console.log(data);
            if (data.status == 200 || data.status == 201) {
                if (index == count - 1) {
                    commonAlert("Success Attach DataDisk!")
                    displayDataDiskInfo("MODIFY_SUCCESS");
                    location.reload();
                } else {

                }

            } else {
                commonAlert("Fail attach DataDisk at " + item + data.message)
                showMyImageInfo(diskId, diskName);
                location.reload();
            }
        }).catch(error => {
            console.log("error response : ", error.response);
        })
    })

}



function getMyImageList(sort_type) {
    console.log(sort_type);
    // defaultNameSpace 기준으로 가져온다. (server단 session에서 가져오므로 변경하려면 현재 namesapce를 바꾸고 호출하면 됨)
    var url = "/setting/resources/myimage/list";
    axios.get(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get MyImage List : ", result.data);
        // var data = result.data.dataDisk;
        var data = result.data.myImageInfoList;
        if (data == null) {
            data = []
        }
        var html = ""
        var cnt = 0;

        if (data.length == 0) {
            html += '<tr><td class="overlay hidden" data-th="" colspan="5">No Data</td></tr>'

            $("#myImageList").empty()
            $("#myImageList").append(html)
            ModalDetail()
        } else {
            if (data.length) {
                if (sort_type) {
                    cnt++;
                    console.log("check : ", sort_type);
                    data.filter(list => list.Name !== "").sort((a, b) => (a[sort_type] < b[sort_type] ? - 1 : a[sort_type] > b[sort_type] ? 1 : 0)).map((item, index) => (
                        html += addMyImageRow(item, index)
                    ))
                } else {
                    data.filter((list) => list.Name !== "").map((item, index) => (
                        html += addMyImageRow(item, index)
                    ))
                }

                $("#myImageList").empty()
                $("#myImageList").append(html)

                ModalDetail()
            }
        }
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

// dataDisk목록에 Item 추가
function addMyImageRow(item, index) {
    console.log("addMyImageRow " + index);
    console.log(item)
    //Disk Attach 여부 확인
    var assoObjList = item.associatedObjectList;
    var diskStatus = ""; //disk Attach 여부 상태. 1) attach되어 있을경우 attached 그렇지 않으면 빈값; 
    var ns = "";
    var mcis_id = "";
    var vm_id = "";
    if (assoObjList != null) {
        var tempAssoObjList = assoObjList[0];
        var parse = tempAssoObjList.split("/")
        ns = parse[2];
        mcis_id = parse[4];
        vm_id = parse[6];

    }

    var html = ""
    html += '<tr onclick="showMyImageInfo(\'' + item.id + '\',\'' + item.name + '\');">'
        + '<td class="overlay hidden column-50px" data-th="">'
        + '<input type="hidden" id="dataDisk_info_' + index + '" value="' + item.name + '"/>'
        + '<input type="checkbox" name="chk" value="' + item.id + '" id="raw_' + index + '" title="" item="' + item.connectionName + '" diskname="' + item.name + '" imageStatus="' + item.status + '" mcis_id="' + mcis_id + '" vm_id="' + vm_id + '" /><label for="td_ch1"></label> <span class="ov off"></span>'

        + '</td>'
        + '<td class="btn_mtd ovm" data-th="name">' + item.name + '<span class="ov"></span></td>'
        + '<td class="overlay hidden" data-th="diskType">' + item.sourceVmId + '</td>'
        + '<td class="overlay hidden" data-th="attachedVm">' + item.connectionName + '</td>'
        + '<td class="overlay hidden" data-th="connectionName">' + item.status + '</td>'
        + '<td class="overlay hidden" data-th="description">' + item.creationDate + '</td>'
        + '</tr>'
    return html;
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

function myImageInfo(targetAction) {
    if (targetAction == "REG") {
        $('#dataDiskCreateBox').toggleClass("active");
        $('#dataDiskInfoBox').removeClass("view");
        $('#dataDiskListTable').removeClass("on");
        var offset = $("#dataDiskCreateBox").offset();
        // var offset = $("#" + target+"").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 300);

        // form 초기화
        $("#regDataDiskName").val('')
        $("#regDataDiskSize").val('')
        $("#regDataDiskType").val('')
        $("#regDescription").val('')
        goFocus('dataDiskCreateBox');
    } else if (targetAction == "REG_SUCCESS") {
        $('#dataDiskCreateBox').removeClass("active");
        $('#dataDiskInfoBox').removeClass("view");
        $('#dataDiskListTable').addClass("on");

        var offset = $("#dataDiskCreateBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);

        // form 초기화
        $("#regDataDiskName").val('')
        $("#regDataDiskSize").val('')
        $("#regDataDiskType").val('')
        $("#regDescription").val('')
        getDataDiskList("name");
    } else if (targetAction == "MODIFY_SUCCESS") {

        getDataDiskList("name");
    } else if (targetAction == "DEL") {
        $('#dataDiskCreateBox').removeClass("active");
        $('#dataDiskInfoBox').addClass("view");
        $('#dataDiskListTable').removeClass("on");

        var offset = $("#dataDiskInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 300);

    } else if (targetAction == "DEL_SUCCESS") {
        $('#dataDiskCreateBox').removeClass("active");
        $('#dataDiskInfoBox').removeClass("view");
        $('#dataDiskListTable').addClass("on");

        var offset = $("#dataDiskInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);

        getDataDiskList("name");
    } else if (targetAction == "CLOSE") {
        $('#dataDiskCreateBox').removeClass("active");
        $('#dataDiskInfoBox').removeClass("view");
        $('#dataDiskListTable').addClass("on");

        var offset = $("#dataDiskInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);
    } else if (targetAction == "Attach") {
        attachDataDisk();
    }
}

// provider에 등록 된 connection 목록 표시
function getConnectionInfo(provider) {
    var url = "/setting/connections/cloudconnectionconfig/" + "list"
    var html = "";
    axios.get(url, {
        headers: {
        }
    }).then(result => {
        console.log('getConnectionConfig result: ', result)
        var data = result.data.ConnectionConfig
        console.log("connection data : ", data);
        var count = 0;
        var configName = "";
        var confArr = new Array();
        for (var i in data) {
            if (provider == data[i].ProviderName) {
                count++;
                html += '<option value="' + data[i].ConfigName + '" item="' + data[i].ProviderName + '">' + data[i].ConfigName + '</option>';
                configName = data[i].ConfigName
                confArr.push(data[i].ConfigName)
            }
        }
        if (count == 0) {
            commonAlert("해당 Provider에 등록된 Connection 정보가 없습니다.")
            html += '<option selected>Select Configname</option>';
        }
        if (confArr.length > 1) {
            configName = confArr[0];
        }
        $("#regConnectionName").empty();
        $("#regConnectionName").append(html);

    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
    });
}



function createDataDisk() {
    var diskName = $("#regDataDiskName").val();
    var diskSize = $("#regDataDiskSize").val();
    var diskType = $("#regDataDiskType").val();
    var description = $("#regDescription").val();
    var provider = $("#provider").val();
    var connectionName = $("#regConnectionName").val();


    var url = "/setting/resources" + "/datadisk/reg"
    console.log("dataDisk Reg URL : ", url)
    var obj = {
        connectionName,
        description,
        name: diskName,
        diskSize,
        diskType,
    }
    console.log("info dataDisk obj Data : ", obj);

    if (diskName) {
        axios.post(url, obj, {
            // headers: {
            //     'Content-type': 'application/json',
            // }
        }).then(result => {
            console.log("result dataDisk : ", result);
            var data = result.data;
            console.log(data);
            if (data.status == 200 || data.status == 201) {
                commonAlert("Success Create DataDisk!")

                displayDataDiskInfo("REG_SUCCESS")

            } else {
                commonAlert("Fail Create DataDisk " + data.message)
            }
        }).catch((error) => {
            console.log(error.response)
            // var errorMessage = error.response.data.error;
            // var statusCode = error.response.status;
            // commonErrorAlert(statusCode, errorMessage)
        });
    } else {
        commonAlert("Input Disk Name")
        $("#regDataDiskName").focus()
        return;
    }
}

// 선택한 dataDisk의 상세정보 : 이미 가져왔는데 다시 가져올 필요있나?? dataDiskID
function showMyImageInfo(myImageId, myImageName) {
    console.log("showMyImageInfo : ", myImageName);

    $('#myImageName').text(myImageName)

    var url = "/setting/resources" + "/myimage/" + encodeURIComponent(myImageId);
    console.log("dataDisk detail URL : ", url)

    return axios.get(url, {
    }).then(result => {
        console.log(result);
        console.log(result.data);
        var data = result.data.myImageInfo
        console.log("Show Data : ", data);

        var dtlMyImageName = data.name;
        var dtlSouceVmId = data.sourceVmId;
        var dtlAssoObj = data.associatedObjectList;
        var dtlStatus = data.status;
        var dtlConnectionName = data.connectionName;
        var dtlCreationDate = data.creationDate;

        if (dtlAssoObj != null) {
            var parse = dtlAssoObj[0];
            var temp_parse = parse.split("/");
            vm_id = temp_parse[6];
        }


        // var subList = data.subnetInfoList;
        // for (var i in subList) {
        //     dtlSubnet += subList[i].id + " (" + subList[i].ipv4_CIDR + ")";
        // }
        // console.log("dtlSubnet : ", dtlSubnet);

        $("#dtlMyImageName").empty();
        $("#dtlSouceVmId").empty();
        $("#dtlAssoObj").empty();
        $("#dtlStatus").empty();
        $("#dtlConnectionName").empty();
        $("#dtlCreationDate").empty();
        console.log("dtlMyImageName : ", dtlMyImageName);
        $("#dtlMyImageName").val(dtlMyImageName);
        $("#dtlSouceVmId").val(dtlSouceVmId);
        $("#dtlAssoObj").val(dtlAssoObj);
        $("#dtlStatus").val(dtlStatus);
        $("#dtlConnectionName").val(dtlConnectionName);
        $("#dtlCreationDate").val(dtlCreationDate);


        if (dtlConnectionName == '' || dtlConnectionName == undefined) {
            commonAlert("ConnectionName is empty")
        } else {
            getProviderNameByConnection(dtlConnectionName, 'dtlProvider')// provider는 connection 정보에서 가져옴            
        }

    }).catch(function (error) {
        console.log("Network detail error : ", error);
    });
}

function putDataDisk() {
    var diskId = $("#dtlDiskId").val();
    var preDiskSize = $("#dtlPreDiskSize").val();
    var diskSize = $("#dtlDiskSize").val();
    var diskName = $("#dtlDiskName").val();
    var description = $("#dtlDescription").val();

    console.log("preDiskSize : ", preDiskSize)
    console.log("diskSize : ", diskSize)
    if (preDiskSize >= diskSize) {
        commonAlert("변경할 디스크 사이즈는 기존 디스크 사이즈 보다 커야 합니다.")
        $("#dtlDiskSize").focus();
        return;
    }
    var url = "/setting/resources" + "/datadisk/" + encodeURIComponent(diskId);

    if (diskId) {
        console.log("disk id : ", diskId)
        console.log("disk modify by put url :", url);
        //put
        var obj = {
            description,
            diskSize
        }
        console.log("modify disk info : ", obj);
        axios.put(url, obj, {
            // headers: {
            //     'Content-type': 'application/json',
            //     // 'Authorization': apiInfo,
            // }
        }).then(result => {
            var data = result.data;
            console.log(data);
            if (data.status == 200 || data.status == 201) {
                commonAlert("Success Modify DataDisk!")
                displayDataDiskInfo("MODIFY_SUCCESS");
                showMyImageInfo(diskId, diskName);

            } else {
                commonAlert("Fail Create DataDisk " + data.message)
                showMyImageInfo(diskId, diskName);
            }
        }).catch(error => {
            console.log(error.response);
        })
    } else {
        commonAlert("Disk ID is empty")
        return;
    }
}


function displayDiskAttachModal(isShow) {
    if (isShow) {
        $("#vmSelectBox").modal();
        $('.dtbox.scrollbar-inner').scrollbar();
    } else {
        $("#vnetCreateBox").toggleClass("active");
    }
}

function getMcisListCallbackSuccess(caller, data) {
    console.log("getMcis List data : ", data);
    var html = "";
    data.forEach((item) => {

        console.log("vm: ", item.vm);
        vm_list = item.vm;
        var mcis_name = item.name;
        var mcis_id = item.id

        vm_list.forEach((item, i) => {
            console.log()
            html += '<tr>'
                + '<td class="overlay hidden column-50px" data-th="">'
                + '<input type="hidden" id="vm_info_' + i + '" value="' + item.name + '"/>'
                + '<input type="checkbox" name="vmChk" value="' + item.id + '" title="" diskId="' + item.id + '" mcisId="' + mcis_id + '" onclick="checkOnlyOne(this,\'' + mcis_id + '\',\'' + item.id + '\')"/><label for="td_ch1"></label> <span class="ov off"></span>'

                + '</td>'
                + '<td class="btn_mtd ovm" data-th="name">' + mcis_name + '<span class="ov"></span></td>'
                + '<td class="btn_mtd ovm" data-th="name">' + item.name + '<span class="ov"></span></td>'
                + '<td class="overlay hidden" data-th="diskType">' + item.connectionName + '</td>'
                + '</tr>'
        })


    })
    console.log(html)
    $("#vmList").empty()
    $("#vmList").append(html)
    displayDiskAttachModal(true)
}

function checkOnlyOne(element, mcis_id, vm_id) {
    var checkBox = document.getElementsByName("vmChk");
    checkBox.forEach(item => {
        item.checked = false;
    })
    element.checked = true;
    $("#attach_mcis_id").val(mcis_id);
    $("#attach_vm_id").val(vm_id);

}
