var subnetJsonList = "";//저장시 subnet목록을 담을 array 
$(document).ready(function(){
    order_type = "name"
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
    });
});

// $(document).ready(function () {
    
//     // var defaultNameSpace = "{{ .DefaultNameSpaceID }}"
//     // alert(defaultNameSpace)
//     // var nameSpaceList = "{{ .NameSpaceList }}"
//     // alert(nameSpaceList);
//     // page load시 이미 가져왔음
//     // getVpcList(order_type);
//     // getCloudOS(apiInfo,'provider');
// })                      

// function goFocus(target) {
//     console.log(event)
//     event.preventDefault();

//     $("#" + target).focus();
//     fnMove(target)
// }

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
    }).catch(function(error){
        console.log("Network delete error : ",error);        
    });
}          

function getVpcList(sort_type) {
    console.log(sort_type);
    // var url = CommonURL + "/ns/" + NAMESPACE + "/resources/vNet";
    //var currentNameSpace = $('$topboxDefaultNameSpaceID').val()
    // defaultNameSpace 기준으로 가져온다. (server단 session에서 가져오므로 변경하려면 현재 namesapce를 바꾸고 호출하면 됨)
    var url = "/setting/resources/network/list";
    axios.get(url, {
        headers: {
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
                data.filter(list => list.Name !== "").sort((a, b) => (a[sort_type] < b[sort_type] ? - 1 : a[sort_type] > b[sort_type] ? 1 : 0)).map((item, index) => (
                    html += '<tr onclick="showVNetInfo(\'' + item.Name + '\');">' 
                        + '<td class="overlay hidden" data-th="">' 
                        + '<input type="hidden" id="sg_info_' + index + '" value="' + item.Name + '|' + item.CidrBlock + '"/>' 
                        + '<input type="checkbox" name="chk" value="' + item.Name + '" id="raw_'  + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' 
                        + '<td class="btn_mtd ovm" data-th="name">' + item.Name + '</td>'
                        + '<td class="overlay hidden" data-th="cidrBlock">' + item.CidrBlock + '</td>' 
                        + '<td class="overlay hidden" data-th="description">' + item.Description + '</td>'  
                        + '<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
                        + '</tr>'
                ))
            } else {
                data.filter((list) => list.Name !== "").map((item, index) => (
                    html += '<tr onclick="showVNetInfo(\'' + item.Name + '\');">' 
                        + '<td class="overlay hidden" data-th="">' 
                        + '<input type="hidden" id="sg_info_' + index + '" value="' + item.Name  + '"/>'
                        + '<input type="checkbox" name="chk" value="' + item.Name + '" id="raw_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span></td>' 
                        + '<td class="btn_mtd ovm" data-th="name">' + item.Name + '<span class="ov"></span></td>' 
                        + '<td class="overlay hidden" data-th="cidrBlock">' + item.CidrBlock + '</td>' 
                        + '<td class="overlay hidden" data-th="description">' + item.Description + '</td>' 
                        + '<td class="overlay hidden" data-th=""><a href="javascript:void(0);"><img src="/assets/img/contents/icon_link.png" class="icon" alt=""/></a></td>' 
                        + '</tr>'
                ))
            }

            $("#vpcList").empty()
            $("#vpcList").append(html)
            
            ModalDetail()
        }
    }).catch(function(error){
        console.log("Network list error : ",error);        
    });
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

function displayVNetInfo(targetAction){
    if( targetAction == "REG"){
        $('#vnetCreateBox').toggleClass("active");
        $('#vNetInfoBox').removeClass("view");
        $('#vNetListTable').removeClass("on");
        var offset = $("#vnetCreateBox").offset();
        // var offset = $("#" + target+"").offset();
    	$("#TopWrap").animate({scrollTop : offset.top}, 300);
    }


    //CreateBox
    // $('#RegistBox .btn_ok.register').click(function(){
    // 		$(".dashboard.register_cont").toggleClass("active");
    // 		$(".dashboard.server_status").removeClass("view");
    // 		$(".dashboard .status_list tbody tr").removeClass("on");
    // 		//ok 위치이동
    // 		$('#RegistBox').on('hidden.bs.modal', function () {
    // 			var offset = $("#CreateBox").offset();
    // 			$("#wrap").animate({scrollTop : offset.top}, 300);
    // 		})		
}

function getConnectionInfo(provider){
    // var url = SpiderURL+"/connectionconfig";
    var url = "/setting/connections/cloudconnectionconfig/" + "list"
    // console.log("provider : ",provider)
    // var provider = $("#provider option:selected").val();
    var html = "";
    // var apiInfo = ApiInfo
    axios.get(url,{
        headers:{
            // 'Authorization': apiInfo
        }
    }).then(result=>{
        console.log('getConnectionConfig result: ',result)
        // var data = result.data.connectionconfig
        var data = result.data.ConnectionConfig
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

    }).catch(function(error){
        console.log("Network data error : ",error);        
    });
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
    
    var subnetJsonList = new Array();
    
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
    $("#subnetRegisterBox").modal("hide");
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

    // var apiInfo = "{{ .apiInfo}}";
    // var url = CommonURL+"/ns/"+NAMESPACE+"/resources/vNet"
    var url = "/setting/resources" + "/resources/vNet"
    console.log("vNet Reg URL : ",url)
    var obj = {
        CidrBlock: cidrBlock,
        ConnectionName: connectionName,
        Description: description,
        Name: vpcName,
        SubnetInfoList: subnetJsonList
    }
    console.log("info vNet obj Data : ", obj);
    
    if (vpcName) {
        axios.post(url, obj, {
            headers: {
                'Content-type': 'application/json',
                // 'Authorization': apiInfo,
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
        }).catch(function(error){
            console.log("Network create error : ",error);        
        });
    } else {
        alert("Input VPC Name")
        $("#reg_vpcName").focus()
        return;
    }
}

// 선택한 vNet의 상세정보 : 이미 가져왔는데 다시 가져올 필요있나?? vNetID
function showVNetInfo(vpcName) {
    console.log("showVNetInfo : ", vpcName);
    // var apiInfo = "{{ .apiInfo}}";
    // var vNetId = encodeURIComponent(vNetName);
    // $('.stxt').html(vpcName);
    $('#networkVpcName').text(vpcName)
    
    // var url = CommonURL+"/ns/"+NAMESPACE+"/resources/vNet/"+ vNetId;
    var url = "/setting/resources" + "/network/" + encodeURIComponent(vpcName);
    console.log("vnet detail URL : ",url)

    return axios.get(url,{
        // headers:{
        //     'Authorization': apiInfo
        // }
    }).then(result=>{
        console.log(result);
        console.log(result.data);
        var data = result.data.VNetInfo
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

        if(dtlConnectionName == '' || dtlConnectionName == undefined ){
            alert("dtlConnectionName is empty")
        }else{
            getProvider(dtlConnectionName);
        }
        
    }) .catch(function(error){
        console.log("Network detail error : ",error);        
    });
}

// 특정 connection 정보에서 Privider set
function getProvider(connectionName) {
    console.log("getProvider  : ",connectionName);
    // var url = SpiderURL+"/connectionconfig/" + target;
    var url = "/setting/connections"+"/cloudconnectionconfig/" + connectionName;
    return axios.get(url,{
        // headers:{
        //     'Authorization': apiInfo
        // }    
    }).then(result=>{
        var data = result.data;
        console.log(data)
        console.log(data.ConnectionConfig)
        var provider = data.ConnectionConfig.ProviderName;
        //var Provider = data.ConnectionConfig.providerName;
        console.log(provider)
        $("#dtlProvider").val(provider);
    }).catch(function(error){
        console.log("Network getProvider error : ",error);        
    });
}

function displaySubnetRegModal(isShow){
    if(isShow){
        $("#subnetRegisterBox").modal();
        $('.dtbox.scrollbar-inner').scrollbar();
    }else{
        $("#vnetCreateBox").toggleClass("active");
    }
}
// $(document).ready(function() {
//     var subnetJsonList = "";
//     //Subnet pop table scrollbar
//       $('.btn_register').on('click', function() {
//         $("#register_box").modal();
//         $('.dtbox.scrollbar-inner').scrollbar();
//     });	
    
//     /*
//     $('.register_cont .btn_cancel').click(function(){
//         $(".dashboard.register_cont").toggleClass("active");
//     });
//     */
// });