function importVmInfoFromFile() {
    var input = document.createElement("input");
    input.type = "file";
    // input.accept = "text/plain"; // 확장자가 xxx, yyy 일때, ".xxx, .yyy"
	input.accept = ".json";
    input.onchange = function (event) {
        processFile(event.target.files[0]);
    };
    input.click();
}

// 선택한 파일을 읽어 화면에 보여줌
function processFile(file) {
	try{
		var reader = new FileReader();
		reader.onload = function () {
			console.log(reader.result);
			console.log("---1")
			// $("#fileContent").val(reader.result);
			
			var jsonStr = JSON.stringify(reader.result)
			console.log(JSON.stringify(jsonStr));

			// 요거는 string으로만 나오네... 
			// console.log("---2")
			// var jsonObj = JSON.parse(reader.result);
			// var jsonObj = JSON.parse(jsonStr);
			// console.log(jsonObj);
			// console.log(jsonObj[0]);
			// console.log(jsonObj[10]);
			// console.log(jsonObj.name);
			// console.log("---3")

			// 요거 작동 하네.  param, value 모두 따옴표로 묶여진 json 형태여야 함.
			var newJ= $.parseJSON(reader.result);

			console.log(newJ.name);
			console.log(newJ.imageId);
			console.log(newJ.connectionName);
			console.log(newJ.securityGroupIds);
			setVmInfoToForm(newJ);
			jsonFormatter(newJ)
			//securityGroupIds: [ 	"sg-mz-aws-us-east-01"	],
		};
		//reader.readAsText(file, /* optional */ "euc-kr");
		reader.readAsText(file);
	}catch(error){
		commonAlert("File Load Failed");
		console.log(error);
	}
}

// json 객체를 textarea에 표시할 때 예쁘게
function jsonFormatter(vmInfoObj){
	// var fmt = JSON.stringify(vmInfoObj, null, "\t"); // stringify with tabs inserted at each level
	var fmt = JSON.stringify(vmInfoObj, null, 4);    // stringify with 4 spaces at each level
	$("#vmImportScriptPretty").val(fmt);	
}

// 선택한 파일을 읽어 form에 Set
function setVmInfoToForm(vmInfoObj){
	//export form
	$("#i_name").val(vmInfoObj.name);
	$("#i_description").val(vmInfoObj.description);
	$("#i_connectionName").val(vmInfoObj.connectionName);
	$("#i_imageId").val(vmInfoObj.imageId);	
	$("#i_specId").val(vmInfoObj.specId);
	$("#i_subnetId").val(vmInfoObj.subnetId);
	$("#i_vNetId").val(vmInfoObj.vNetId);
	$("#i_securityGroupIds").val(vmInfoObj.securityGroupIds);
	$("#i_sshKeyId").val(vmInfoObj.sshKeyId);
	$("#i_label").val(vmInfoObj.label);

	$("#i_vmUserAccount").val(vmInfoObj.vmUserAccount);
	$("#i_vmUserPassword").val(vmInfoObj.vmUserPassword);

	var addServerCnt = vmInfoObj.vmGroupSize == "" ? 0: vmInfoObj.vmGroupSize;
	$("#i_vm_add_cnt").val(addServerCnt);

	$("#vmImportScript").val(JSON.stringify(vmInfoObj));
	
}

			
const Import_Server_Config_Arr = new Array();
var import_data_cnt = 0
const importServerCloneObj = obj=>JSON.parse(JSON.stringify(obj))
function importDone_btn(){
	var import_form = $("#import_form").serializeObject()
	var server_name = import_form.name
	var server_cnt = parseInt(import_form.i_vm_add_cnt)
	console.log('server_cnt : ',server_cnt)
	var add_server_html = "";
	
	if(server_cnt > 1){
		for(var i = 1; i <= server_cnt; i++){
			var new_vm_name = server_name+"-"+i;
			var object = importServerCloneObj(import_form)
			object.name = new_vm_name
			
			add_server_html +='<li onclick="view_import(\''+import_data_cnt+'\')">'
					+'<div class="server server_on bgbox_b">'
					+'<div class="icon"></div>'
					+'<div class="txt">'+new_vm_name+'</div>'
					+'</div>'
					+'</li>';
			Import_Server_Config_Arr.push(object)
			console.log(i+"번째 import form data 입니다. : ",object);
		}
	}else{
		Import_Server_Config_Arr.push(import_form)
		add_server_html +='<li onclick="view_import(\''+import_data_cnt+'\')">'
						+'<div class="server server_on bgbox_b">'
						+'<div class="icon"></div>'
						+'<div class="txt">'+server_name+'</div>'
						+'</div>'
						+'</li>';

	}

	// Done 버튼 클릭 시 form은 비활성
	$(".import_servers_config").removeClass("active");

	// server List에 추가
	$("#mcis_server_list").prepend(add_server_html)
	console.log("import btn click and import form data : ",import_form)
	console.log("import data array : ",Import_Server_Config_Arr);
	import_data_cnt++;
	$("#import_form").each(function(){
		this.reset();
	})
}
function view_import(cnt){
	console.log('view import cnt : ',cnt);
	var select_form_data = Import_Server_Config_Arr[cnt]
	console.log('select_form_data : ', select_form_data);
	$(".simple_servers_config").removeClass("active")
	$(".expert_servers_config").removeClass("active")
	$(".import_servers_config").addClass("active")
}

