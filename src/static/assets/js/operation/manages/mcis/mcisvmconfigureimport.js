function importVmInfoFromFile() {
    var input = document.createElement("input");
    input.type = "file";
    input.accept = "text/plain"; // 확장자가 xxx, yyy 일때, ".xxx, .yyy"
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
			console.log("---2")
			// var jsonObj = JSON.parse(reader.result);
			var jsonObj = JSON.parse(jsonStr);
			console.log(jsonObj);
		};
		//reader.readAsText(file, /* optional */ "euc-kr");
		reader.readAsText(file);
	}catch(error){
		commonAlert("File Load Failed");
		console.log(error);
	}
}

// // 파일 선택
// function importVmInfoFromFile(){
// 	var input = document.createElement("input");
//     input.type = "file";
//     input.accept = "text/plain"; // 확장자가 xxx, yyy 일때, ".xxx, .yyy"
//     input.onchange = function (event) {
//         //processFile(event.target.files[0]);
// 		setVmInfoToForm(event.target.files[0]);
//     };
//     input.click();
// }
// // 선택한 파일을 읽어 화면에 보여줌
// function setVmInfoToForm(file) {
//     var reader = new FileReader();
//     reader.onload = function () {
// 		console.log(reader.result);
// 		var jsonObj = JSON.parse(reader.result);
//         // $("#fileContent").val(reader.result);
// 		console.log(jsonObj);
//     };
//     //reader.readAsText(file, /* optional */ "euc-kr");
// 	// reader.readAsText(file);
// }



// 선택한 파일을 읽어 form에 Set
// function setVmInfoToForm(vmInfoStr){
// 	console.log("setVmInfo");
// 	console.log(vmInfoStr);
// 	//Split
// 	// 1. 콤마
// 	// 2. : 콜론
// 	//var params = vmInfoStr.split(",")

// 	var jsonObj = JSON.parse(vmInfoStr);
// 	console.log(jsonObj);
// 	// $.getJSON(vmInfoStr, function(data) {
// 	// 	console.log("getJson");
// 	// 	console.log(data);
// 	// 	var html = '';
// 	// 	$.each(data, function(entryIndex, entry) {
// 	// 		console.log(entryIndex + " : " + entry)
// 	// 		// html += '<div class="entry">';
// 	// 		// html += '<h3 class="term">' + entry.term + '</h3>';
// 	// 		// html += '<div class="part">' + entry.part + '</div>';
// 	// 		// html += '<div class="definition">';
// 	// 		// html += entry.definition;
// 	// 		// html += '</div>';
// 	// 		// html += '</div>';
// 	// 	});
// 	// 	// console.log(html);
// 	// 	// $('#dictionary').html(html);
// 	// });
// }
			
			
const Import_Server_Config_Arr = new Array();
var import_data_cnt = 0
const importServerCloneObj = obj=>JSON.parse(JSON.stringify(obj))
function import_btn(){
	var import_form = $("#import_form").serializeObject()
	var server_name = import_form.name
	var server_cnt = parseInt(import_form.s_vm_add_cnt)
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
	$(".import_servers_config").removeClass("active");
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
	$(".import_servers_config").addClass("active")
	$(".new_servers_config").removeClass("active")

}

