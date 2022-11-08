$(document).ready(function () {
	$('.btn_recommend').on('click', function () {
		showRecommendAssistPopup();
	});

})

// assist를 통해 선택한 spec 정보를 받음
function setAssistSpecToExpress(specInfo){
	var specName = specInfo.SpecID;
	var provider = specInfo.Provider;
	var connectionName = specInfo.ConnectionName;
	
	
	$("#ep_provider").val(provider)
	$("#ep_connectionName").val(connectionName)
	$("#ep_spec").val(specName)
	$("#ep_imageId").val("ubuntu18.04")
	console.log("setAssist", specInfo)

	// rootDisk의 Type 조회
	getCommonLookupDiskInfo('vmsimple', provider, connectionName)
}

function getSshKeyListCallbackSuccessForExpress(caller,data){
	console.log(data);
	var html = ""
	//data = data.SshKeyList
	html += '<option value="">Select SSH Key</option>'
	for (var i in data) {
		html += '<option value="' + data[i].id + '" >' + data[i].cspSshKeyName + '(' + data[i].id + ')</option>';
	}
	$("#ep_sshKey").empty();
	$("#ep_sshKey").append(html);
}
	

// // 공통으로 빼야할텐데
// var DISK_SIZE = [];
// function getCommonLookupDiskInfoSuccess(caller, provider, data){
	
// 	console.log("getCommonLookupDiskInfoSuccess",data[0]);
// 	var root_disk_type = [];
// 	var res_item = data
// 	res_item.forEach(item=>{
// 		var temp_provider = item.provider
// 		if(temp_provider == provider){
// 			root_disk_type = item.rootdisktype
// 			DISK_SIZE = item.disksize
// 		}
// 	})

// 	var html = '<option value="">Select Disk Type</option>'
// 	console.log("disk_type : ",root_disk_type);
// 	root_disk_type.forEach(item=>{
// 		html += '<option value="'+item+'">'+item+'</option>'
// 	})
// if(caller == "vmsimple_root"){
// 	$("#ss_root_disk_type").empty();
// 	$("#ss_root_disk_type").append(html);
// }else if(caller == "vmsimple_datadisk"){
// 	$("#ss_data_disk_type").empty();
// 	$("#ss_data_disk_type").append(html);
// }else if(caller == "vmexpress_root"){
// 	$("#p_root_disk_type").empty();
// 	$("#p_root_disk_type").append(html);
// }else if(caller == "vmexpress_data"){
// 	$("#p_data_disk_type").empty();
// 	$("#p_data_disk_type").append(html);
// }else if(caller == "vmexpert_data"){
// 	$("#es_data_disk_type").empty();
// 	$("#es_data_disk_type").append(html);
// }else if(caller == "vmexpert_data"){
// 	$("#es_data_disk_type").empty();
// 	$("#es_data_disk_type").append(html);
// }else{
// 	$("#e_root_disk_type").empty()
// 	$("#e_root_disk_type").append(html)
// }
// console.log("const valie DISK_SIZE : ",DISK_SIZE);


// }

// // 공통으로 빼야할텐데
// var DISK_MAX_VALUE = 0;
// var DISK_MIN_VALUE = 0;
// // 공통으로 빼야할텐데
// function changeDiskSize(type){
// 	var disk_size = DISK_SIZE;

// 	if(disk_size){
// 		disk_size.forEach(item=>{
// 			var temp_size = item.split("|")
// 			var temp_type = temp_size[0];
// 			if(temp_type == type){
// 				DISK_MAX_VALUE = temp_size[1];
// 				DISK_MIN_VALUE = temp_size[2]
// 			}
// 		})
// 	}
// 	console.log("ROOT_DISK_MAX_VALUE : ",DISK_MAX_VALUE)
// 	console.log("ROOT_DISK_MIN_VALUE : ",DISK_MIN_VALUE)
// 	$("#s_rootDiskType").val(type);
// 	$("#e_rootDiskType").val(type);
// 	$("#p_rootDiskType").val(type);

// }

const Express_Server_Config_Arr = new Array();
var express_data_cnt = 0
function expressDone_btn() {
	// express 는 common resource를 하므로 별도로 처리(connection, spec만)
	$("#p_provider").val($("#ep_provider").val())
	$("#p_connectionName").val($("#ep_connectionName").val())
	$("#p_name").val($("#ep_name").val())
	$("#p_description").val($("#ep_description").val())
	$("#p_spec").val($("#ep_spec").val())
	$("#p_subGroupSize").val($("#ep_vm_add_cnt").val() + "")
	$("#p_vm_cnt").val($("#ep_vm_add_cnt").val() + "")


	//var express_form = $("#express_form").serializeObject()
	// commonSpec 으로 set 해야하므로 재설정
	var express_form = {}
	express_form["name"] = $("#p_name").val();
	express_form["connectionName"] = $("#p_connectionName").val();
	express_form["description"] = $("#p_description").val();
	express_form["subGroupSize"] = $("#p_subGroupSize").val();
	express_form["commonImage"] = "ubuntu18.04";
	express_form["commonSpec"] = $("#p_spec").val();

	console.log("express_form form : ", express_form);

	var server_name = express_form.name
	
	var server_cnt = parseInt(express_form.subGroupSize)
	
	var add_server_html = "";

	Express_Server_Config_Arr.push(express_form)


	var displayServerCnt = '(' + server_cnt + ')'

	add_server_html += '<li onclick="view_express(\'' + express_data_cnt + '\')">'
		+ '<div class="server server_on bgbox_b">'
		+ '<div class="icon"></div>'
		+ '<div class="txt">' + server_name + displayServerCnt + '</div>'
		+ '</div>'
		+ '</li>';

	// }
	$(".express_servers_config").removeClass("active");

	console.log("add server html");
	$("#mcis_server_list").prepend(add_server_html)
	
	$("#plusVmIcon").remove();
	$("#mcis_server_list").prepend(getPlusVm());

	console.log("express btn click and express form data : ", express_form)
	console.log("express data array : ", Express_Server_Config_Arr);
	express_data_cnt++;
	$("#express_form").each(function () {
		this.reset();
	})
	$("#ep_data_disk").val("");

}

function view_express(cnt) {
	console.log('view simple cnt : ', cnt);
	var select_form_data = Simple_Server_Config_Arr[cnt]
	console.log('select_form_data : ', select_form_data);
	$(".express_servers_config").addClass("active")
	$(".simple_servers_config").removeClass("active")
	$(".expert_servers_config").removeClass("active")
	$(".import_servers_config").removeClass("active")

}
