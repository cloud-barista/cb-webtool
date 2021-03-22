$(document).ready(function () {
    order_type = "name"

    // var defaultNameSpace = "{{ .DefaultNameSpaceID }}"
    // alert(defaultNameSpace)
    // var nameSpaceList = "{{ .NameSpaceList }}"
    // alert(nameSpaceList);
    // page load시 이미 가져왔음
    // getVpcList(order_type);
    // getCloudOS(apiInfo,'provider');
})                      

function goFocus(target) {
    console.log(event)
    event.preventDefault();

    $("#" + target).focus();
    fnMove(target)
}

function fnMove(target) {
    var offset = $("#" + target).offset();
    console.log("fn move offset : ", offset);
    $('html, body').animate({
        scrollTop: offset.top
    }, 400);
}

function goDelete() {
    var vNetId = "";
    var count = 0;

    $( "input[name='chk']:checked" ).each (function (){
        count++;
        vNetId = vNetId + $(this).val()+"," ;
    });
    vNetId = vNetId.substring(0,vNetId.lastIndexOf( ","));
    
    console.log("vNetId : ", vNetId);
    console.log("count : ", count);

    if(vNetId == ''){
        alert("삭제할 대상을 선택하세요.");
        return false;
    }

    if(count != 1){
        alert("삭제할 대상을 하나만 선택하세요.");
        return false;
    }

    var url = CommonURL + "/ns/" + NAMESPACE + "/resources/vNet/" + vNetId;
    console.log("del vnet url : ", url);

    axios.delete(url, {
        headers: {
            'Authorization': "{{ .apiInfo}}",
            'Content-Type': "application/json"
        }
    }).then(result => {
        var data = result.data;
        console.log(data);
        if (result.status == 200 || result.status == 201) {
            alert("Success Delete Network.");
            location.reload(true);
        }
    })
}          

function getVpcList(sort_type) {
    console.log(sort_type);
    // var url = CommonURL + "/ns/" + NAMESPACE + "/resources/vNet";
    var url = "";
    axios.get(url, {
        headers: {
            'Authorization': "{{ .apiInfo}}",
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get VPC List : ", result.data);
        var data = result.data.vNet;
        var html = ""
        var cnt = 0;
        
        if (data.length) {
            if (sort_type) {
                cnt++;
                console.log("check : ", sort_type);
                data.filter(list => list.name !== "").sort((a, b) => (a[sort_type] < b[sort_type] ? - 1 : a[sort_type] > b[sort_type] ? 1 : 0)).map((item, index) => (
                    html += '<tr onclick="showInfo(\'' + item.name + '\');">' 
                        + '<td class="overlay hidden" data-th="">' 
                        + '<input type="hidden" id="sg_info_' + index + '" value="' + item.name + '|' + item.cidrBlock + '"/>' 
                        + '<input type="checkbox" name="chk" value="' + item.name + '" id="raw_'  + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' 
                        + '<td class="btn_mtd ovm" data-th="name">' + item.name + '</td>'
                        + '<td class="overlay hidden" data-th="cidrBlock">' + item.cidrBlock + '</td>' 
                        + '<td class="overlay hidden" data-th="description">' + item.description + '</td>'  
                        + '<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
                        + '</tr>'
                ))
            } else {
                data.filter((list) => list.name !== "").map((item, index) => (
                    html += '<tr onclick="showInfo(\'' + item.name + '\');">' 
                        + '<td class="overlay hidden" data-th="">' 
                        + '<input type="hidden" id="sg_info_' + index + '" value="' + item.name  + '"/>'
                        + '<input type="checkbox" name="chk" value="' + item.name + '" id="raw_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' 
                        + '<td class="btn_mtd ovm" data-th="name">' + item.name + '<span class="ov"></span></td>' 
                        + '<td class="overlay hidden" data-th="cidrBlock">' + item.cidrBlock + '</td>' 
                        + '<td class="overlay hidden" data-th="description">' + item.description + '</td>' 
                        + '<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
                        + '</tr>'
                ))
            }

            $("#vpcList").empty()
            $("#vpcList").append(html)
            
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

function getConnectionInfo(provider){
    var url = SpiderURL+"/connectionconfig";
    console.log("provider : ",provider)
    //var provider = $("#provider option:selected").val();
    var html = "";
    var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            'Authorization': apiInfo
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

function applySubnet() {
    var subnetNameValue = $("input[name='reg_subnetName']").length;
    var subnetCIDRBlockValue = $("input[name='reg_subnetCidrBlock']").length;
    
    var subnetNameData = new Array(subnetNameValue);
    var subnetCIDRBlockData = new Array(subnetCIDRBlockValue);
    
    for(var i=0; i<subnetNameValue; i++){                          
        subnetNameData[i] = $("input[name='reg_subnetName']")[i].value;
        console.log("subnetNameData" + [i] + " : ", subnetNameData[i]);
    }
    for(var i=0; i<subnetCIDRBlockValue; i++){                          
        subnetCIDRBlockData[i] = $("input[name='reg_subnetCidrBlock']")[i].value;
        console.log("subnetCIDRBlockData" + [i] + " : ", subnetCIDRBlockData[i]);
    }
    
    subnetJsonList = new Array();
    
    for(var i=0; i<subnetNameValue; i++){
        var SNData = "SNData" + i;
        var SNData = new Object();
        SNData.name = subnetNameData[i];
        SNData.ipv4_CIDR = subnetCIDRBlockData[i];
        subnetJsonList.push(SNData);
    }

    var infoshow = "";
    for (var i in subnetJsonList) {
        infoshow += subnetJsonList[i].name + " (" + subnetJsonList[i].ipv4_CIDR + ") ";
    }
    
    $("#reg_subnet").empty();
    $("#reg_subnet").val(infoshow);
    
    $("#register_box").modal("hide");
}

function createVNet() {
    var vpcName = $("#reg_vpcName").val();
    var description = $("#reg_description").val();
    var connectionName = $("#reg_connectionName").val();
    var cidrBlock = $("#reg_cidrBlock").val();
    
    if (!vpcName) {
        alert("Input New VPC Name")
        $("#reg_vpcName").focus()
        return;
    }

    var apiInfo = "{{ .apiInfo}}";
    var url = CommonURL+"/ns/"+NAMESPACE+"/resources/vNet"
    console.log("vNet Reg URL : ",url)
    var obj = {
        cidrBlock: cidrBlock,
        connectionName: connectionName,
        description: description,
        name: vpcName,
        subnetInfoList: subnetJsonList
    }
    console.log("info vNet obj Data : ", obj);
    
    if (vpcName) {
        axios.post(url, obj, {
            headers: {
                'Content-type': 'application/json',
                'Authorization': apiInfo,
            }
        }).then(result => {
            console.log("result vNet : ", result);
            if (result.status == 200 || result.status == 201) {
                alert("Success Create Network(vNet)!!")
                //등록하고 나서 화면을 그냥 고칠 것인가?
                getVpcList();
                //아니면 화면을 리로딩 시킬것인가?
                location.reload();
                // $("#btn_add2").click()
                // $("#namespace").val('')
                // $("#nsDesc").val('')
            } else {
                alert("Fail Create Network(vNet)")
            }
        });
    } else {
        alert("Input VPC Name")
        $("#reg_vpcName").focus()
        return;
    }
}

function showInfo(target) {
    console.log("target showInfo : ", target);
    var apiInfo = "{{ .apiInfo}}";
    var vNetId = encodeURIComponent(target);
    $('.stxt').html(target);
    
    var url = CommonURL+"/ns/"+NAMESPACE+"/resources/vNet/"+ vNetId;
    console.log("vnet detail URL : ",url)

    return axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    
    }).then(result=>{
        var data = result.data
        console.log("Show Data : ",data);
        
        var dtlVpcName = data.name;
        var dtlDescription = data.description;
        var dtlConnectionName = data.connectionName;
        var dtlCidrBlock = data.cidrBlock;
        var dtlSubnet = "";

        var subList = data.subnetInfoList;
        for (var i in subList) {
            dtlSubnet = subList[i].IId.NameId + " (" + subList[i].IPv4_CIDR + ")";
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

        getProvider(dtlConnectionName);
    }) 
}

function getProvider(target) {
    console.log("getProvidergetProvider : ",target);
    var url = SpiderURL+"/connectionconfig/" + target;
        
    return axios.get(url,{
        headers:{
            'Authorization': apiInfo
        }
    
    }).then(result=>{
        var data = result.data;
        
        var Provider = data.ProviderName;

        $("#dtlProvider").val(Provider);
    })        
}

$(document).ready(function() {
    var subnetJsonList = "";
    //Subnet pop table scrollbar
      $('.btn_register').on('click', function() {
        $("#register_box").modal();
        $('.dtbox.scrollbar-inner').scrollbar();
    });	
    
    /*
    $('.register_cont .btn_cancel').click(function(){
        $(".dashboard.register_cont").toggleClass("active");
    });
    */
});