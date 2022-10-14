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

        setTableHeightForScroll('vpcListTable', 300);
    });
});

// function goDelete() {
function deleteDataDisk() {
    var dataDiskId = "";
    var count = 0;

    $("input[name='chk']:checked").each(function () {
        count++;
        dataDiskId = dataDiskId + $(this).val() + ",";
    });
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

function getDataDiskList(sort_type) {
    console.log(sort_type);
    // defaultNameSpace 기준으로 가져온다. (server단 session에서 가져오므로 변경하려면 현재 namesapce를 바꾸고 호출하면 됨)
    var url = "/setting/resources/datadisk/list";
    axios.get(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get VPC List : ", result.data);
        // var data = result.data.dataDisk;
        var data = result.data.dataDiskList;

        var html = ""
        var cnt = 0;

        if (data == null) {
            html += '<tr><td class="overlay hidden" data-th="" colspan="5">No Data</td></tr>'

            $("#dataDiskList").empty()
            $("#dataDiskList").append(html)

            ModalDetail()
        } else {
            if (data.length) {
                if (sort_type) {
                    cnt++;
                    console.log("check : ", sort_type);
                    data.filter(list => list.Name !== "").sort((a, b) => (a[sort_type] < b[sort_type] ? - 1 : a[sort_type] > b[sort_type] ? 1 : 0)).map((item, index) => (
                        html += addDataDiskRow(item, index)
                    ))
                } else {
                    data.filter((list) => list.Name !== "").map((item, index) => (
                        html += addDataDiskRow(item, index)
                    ))
                }

                $("#dataDiskList").empty()
                $("#dataDiskList").append(html)

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
function addDataDiskRow(item, index) {
    console.log("addDataDiskRow " + index);
    console.log(item)
    var html = ""
    html += '<tr onclick="showDataDiskInfo(\'' + item.name + '\');">'
        + '<td class="overlay hidden column-50px" data-th="">'
        + '<input type="hidden" id="sg_info_' + index + '" value="' + item.name + '"/>'
        + '<input type="checkbox" name="chk" value="' + item.name + '" id="raw_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span>'
        + '<input type="hidden" value="' + item.systemLabel + '"/>'
        + '</td>'
        + '<td class="btn_mtd ovm" data-th="name">' + item.name + '<span class="ov"></span></td>'
        + '<td class="overlay hidden" data-th="cidrBlock">' + item.cidrBlock + '</td>'
        + '<td class="overlay hidden" data-th="description">' + item.description + '</td>'
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

function displayDataDiskInfo(targetAction) {
    if (targetAction == "REG") {
        $('#dataDiskkCreateBox').toggleClass("active");
        $('#dataDiskkInfoBox').removeClass("view");
        $('#dataDiskkListTable').removeClass("on");
        var offset = $("#dataDiskkCreateBox").offset();
        // var offset = $("#" + target+"").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 300);

        // form 초기화
        $("#regVpcName").val('')
        $("#regDescription").val('')
        $("#regCidrBlock").val('')
        $("#regSubnet").val('')
        goFocus('dataDiskkCreateBox');
    } else if (targetAction == "REG_SUCCESS") {
        $('#dataDiskkCreateBox').removeClass("active");
        $('#dataDiskkInfoBox').removeClass("view");
        $('#dataDiskkListTable').addClass("on");

        var offset = $("#dataDiskkCreateBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);

        // form 초기화
        $("#regVpcName").val('')
        $("#regDescription").val('')
        $("#regCidrBlock").val('')
        $("#regSubnet").val('')
        getVpcList("name");
    } else if (targetAction == "DEL") {
        $('#dataDiskkCreateBox').removeClass("active");
        $('#dataDiskkInfoBox').addClass("view");
        $('#dataDiskkListTable').removeClass("on");

        var offset = $("#dataDiskkInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 300);

    } else if (targetAction == "DEL_SUCCESS") {
        $('#dataDiskkCreateBox').removeClass("active");
        $('#dataDiskkInfoBox').removeClass("view");
        $('#dataDiskkListTable').addClass("on");

        var offset = $("#dataDiskkInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);

        getVpcList("name");
    } else if (targetAction == "CLOSE") {
        $('#dataDiskkCreateBox').removeClass("active");
        $('#dataDiskkInfoBox').removeClass("view");
        $('#dataDiskkListTable').addClass("on");

        var offset = $("#dataDiskkInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);
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



function createDatadisk() {
    var vpcName = $("#regVpcName").val();
    var description = $("#regDescription").val();
    var connectionName = $("#regConnectionName").val();
    var cidrBlock = $("#regCidrBlock").val();
    if (!vpcName) {
        commonAlert("Input New VPC Name")
        $("#regVpcName").focus()
        return;
    }
    
    var url = "/setting/resources" + "/datadisk/reg"
    console.log("dataDisk Reg URL : ", url)
    var obj = {
        CidrBlock: cidrBlock,
        ConnectionName: connectionName,
        Description: description,
        Name: vpcName,
        SubnetInfoList: subnetJsonList
    }
    console.log("info dataDisk obj Data : ", obj);

    if (vpcName) {
        axios.post(url, obj, {
            headers: {
                'Content-type': 'application/json',
            }
        }).then(result => {
            console.log("result dataDisk : ", result);
            var data = result.data;
            console.log(data);
            if (data.status == 200 || data.status == 201) {
                commonAlert("Success Create Network(VPC)!!")

                displayDatadiskInfo("REG_SUCCESS")

            } else {
                commonAlert("Fail Create Network(VPC) " + data.message)
            }
        }).catch((error) => {
            console.log(error.response)
            var errorMessage = error.response.data.error;
            var statusCode = error.response.status;
            commonErrorAlert(statusCode, errorMessage)
        });
    } else {
        commonAlert("Input VPC Name")
        $("#regVpcName").focus()
        return;
    }
}

// 선택한 dataDisk의 상세정보 : 이미 가져왔는데 다시 가져올 필요있나?? dataDiskID
function showDataDiskInfo(dataDiskName) {
    console.log("showDataDiskInfo : ", vpcName);
    
    $('#networkVpcName').text(vpcName)

    var url = "/setting/resources" + "/datadisk/" + encodeURIComponent(dataDiskName);
    console.log("dataDisk detail URL : ", url)

    return axios.get(url, {
    }).then(result => {
        console.log(result);
        console.log(result.data);
        var data = result.data.dataDiskInfo
        console.log("Show Data : ", data);

        var dtlVpcName = data.name;
        var dtlDescription = data.description;
        var dtlConnectionName = data.connectionName;
        var dtlCidrBlock = data.cidrBlock;
        var dtlSubnet = "";

        var subList = data.subnetInfoList;
        for (var i in subList) {
            dtlSubnet += subList[i].id + " (" + subList[i].ipv4_CIDR + ")";
        }
        console.log("dtlSubnet : ", dtlSubnet);

        $("#dtlVpcName").empty();
        $("#dtlDescription").empty();
        $("#dtlProvider").empty();
        $("#dtlConnectionName").empty();
        $("#dtlCidrBlock").empty();
        $("#dtlSubnet").empty();

        $("#dtlVpcName").val(dtlVpcName);
        $("#dtlDescription").val(dtlDescription);
        $("#dtlConnectionName").val(dtlConnectionName);
        $("#dtlCidrBlock").val(dtlCidrBlock);
        $("#dtlSubnet").val(dtlSubnet);

        if (dtlConnectionName == '' || dtlConnectionName == undefined) {
            commonAlert("ConnectionName is empty")
        } else {
            getProviderNameByConnection(dtlConnectionName, 'dtlProvider')// provider는 connection 정보에서 가져옴            
        }

    }).catch(function (error) {
        console.log("Network detail error : ", error);
    });
}

