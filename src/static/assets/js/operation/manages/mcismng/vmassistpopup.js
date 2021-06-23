
$(document).ready(function(){
	
	// #ID 에 .클래스명_assist
	//	대상 class명.toggleClass
	$('#OS_HW_Spec .btn_spec_assist').click(function(){
		$(".spec_select_box").toggleClass("active");
	});

	$('#OS_HW_Spec .btn_image_assist').click(function(){
		$(".spec_select_box").toggleClass("active");
	});
});
															
function openTextFile() {
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
    var reader = new FileReader();
    reader.onload = function () {
		console.log(reader.result);
        $("#fileContent").val(reader.result);
    };
    //reader.readAsText(file, /* optional */ "euc-kr");
	reader.readAsText(file);
}


// function exportVmScript(vmIndex){
	
// 	var connectionNameVal = $("#p_connectionName_" + vmIndex).val();
// 	var descriptionVal = $("#p_description_" + vmIndex).val();
// 	var imageIdVal = $("#p_imageId_" + vmIndex).val();
// 	var labelVal = $("#p_label_" + vmIndex).val();
// 	var nameVal = $("#p_name_" + vmIndex).val();
// 	var securityGroupIdsVal = $("#p_securityGroupIds_" + vmIndex).val();
// 	var specIdVal = $("#p_specId_" + vmIndex).val();
// 	var sshKeyIdVal = $("#p_sshKeyId_" + vmIndex).val();
// 	var subnetIdVal = $("#p_subnetId_" + vmIndex).val();
// 	var vNetIdVal = $("#p_vNetId_" + vmIndex).val();
// 	var vmGroupSizeVal = $("#p_vmGroupSize_" + vmIndex).val();
// 	var vmUserAccountVal = $("#p_vmUserAccount_" + vmIndex).val();
// 	var vmUserPasswordVal = $("#p_vmUserPassword_" + vmIndex).val();

// 	var paramValueAppend = '"';
// 	var vmCreateScript = "";
// 	vmCreateScript += '{	';
// 	vmCreateScript += paramValueAppend + 'connectionName' + paramValueAppend + ' : ' + paramValueAppend + connectionNameVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'description' + paramValueAppend + ' : ' + paramValueAppend + descriptionVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'imageId' + paramValueAppend + ' : ' + paramValueAppend + imageIdVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'label' + paramValueAppend + ' : ' + paramValueAppend + labelVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'name' + paramValueAppend + ' : ' + paramValueAppend + nameVal + paramValueAppend;
// 	// vmCreateScript += ',securityGroupIds: ';
//     // vmCreateScript += '	' + paramValueAppend + securityGroupIdsVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'specId' + paramValueAppend + ' : ' + paramValueAppend + specIdVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'sshKeyId' + paramValueAppend + ' : ' + paramValueAppend + sshKeyIdVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'subnetId' + paramValueAppend + ' : ' + paramValueAppend + subnetIdVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'vNetId' + paramValueAppend + ' : ' + paramValueAppend + vNetIdVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'vmGroupSize' + paramValueAppend + ' : ' + paramValueAppend + vmGroupSizeVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'vmUserAccount' + paramValueAppend + ' : ' + paramValueAppend + vmUserAccountVal + paramValueAppend;
// 	vmCreateScript += ',' + paramValueAppend + 'vmUserPassword' + paramValueAppend + ' : ' + paramValueAppend + vmUserPasswordVal + paramValueAppend;
// 	vmCreateScript += '}';

	
// 	$("#exportFileName").val(nameVal);
// 	$("#vmExportScript").val(vmCreateScript);
// }

// function saveVmInfoToFile(){
// 	var fileName = $("#exportFileName").val();
// 	var exportScript = $("#vmExportScript").val();
	
// 	var element = document.createElement('a');
// 	// element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(exportScript));
// 	element.setAttribute('href', 'data:text/json;charset=utf-8,' + encodeURIComponent(exportScript));
// 	// element.setAttribute('download', fileName);
// 	element.setAttribute('download', fileName + ".json");

// 	element.style.display = 'none';
// 	document.body.appendChild(element);

// 	element.click();

// 	document.body.removeChild(element);

// }

// assist에서 provider 선택시 retion filter
function getRegionListFilterAtAssist(provider, targetRegionObj){
	// region 목록 filter
	selectBoxFilterByText(targetRegionObj, provider)
	$("#" + targetRegionObj + " option:eq(0)").attr("selected", "selected");
}

// assist popup에서 조회조건에 맞는 spec을 검색
function assistFilterSpec(){
	var conditionArr = new Array();
	conditionArr.push("cost_per_hour");
	conditionArr.push("ebs_bw_Mbps");
	// conditionArr.push("evaluationScore_01");
	// conditionArr.push("evaluationScore_01");
	// conditionArr.push("evaluationScore_01");
	// conditionArr.push("evaluationScore_01");
	// conditionArr.push("evaluationScore_01");
	// conditionArr.push("evaluationScore_01");
	// conditionArr.push("evaluationScore_01");
	// conditionArr.push("evaluationScore_01");
	// conditionArr.push("evaluationScore_01");
	// conditionArr.push("evaluationScore_01");

	// conditionArr.push("gpumem_GiB");
	conditionArr.push("max_num_storage");
	// conditionArr.push("max_total_storage_TiB");
	// conditionArr.push("mem_GiB");
	// conditionArr.push("net_bw_Gbps");
	// conditionArr.push("num_core");
	// conditionArr.push("num_gpu");
	// conditionArr.push("num_storage");
	conditionArr.push("num_vCPU");
	// conditionArr.push("storage_GiB");
	
	// 
	var searchObj = {}
	searchObj['connectionName'] = "";
	// var condition_CostPerHour = {}
	// condition_CostPerHour['max'] = Number(costPerHourMax)
	// condition_CostPerHour['min'] = Number(costPerHourMin)
	// searchObj['cost_per_hour'] = condition_CostPerHour;

	// var condition_ebsBwMbps = {}
	// condition_ebsBwMbps['max'] = Number(ebsBwMbpsMax)
	// condition_ebsBwMbps['min'] = Number(ebsBwMbpsMax)
	// searchObj['ebs_bw_Mbps'] = condition_ebsBwMbps;
	// assist_num_vCPU_min
	for( var i = 0 ; i < conditionArr.length; i++){
		var conditionMaxValue = $("#assist_" + conditionArr[i] + "_max").val();
		var conditionMinValue = $("#assist_" + conditionArr[i] + "_min").val();
		console.log("conditionMinValue=" + conditionMinValue);
		console.log("conditionMaxValue=" + conditionMaxValue);
		if( conditionMaxValue && conditionMinValue){
			var conditionParam = {};
			conditionParam['max'] = conditionMaxValue;
			conditionParam['min'] = conditionMinValue;
			searchObj[conditionArr[i]] = conditionParam;
		}
	}
	console.log(searchObj);
	// axios 전송
	getCommonFilterSpecsByRange("vmassistpopup", searchObj);
	// assist_specList 에 append
}

// Spec Range 조회 성공
function filterSpecsByRangeCallbackSuccess(caller, data){
	console.log(data)

    var html = ""
    var vmSpecList = data
    
    $("#register_box").modal()    
    vmSpecList.map(item=>(        html +='<tr>'
                +'<td class="btn_mtd" data-th="spec ID">'+item.fromPort+' <span class="ov off"></span></td>'
                +'<td class="overlay hidden" data-th="toPort">'+item.toPort+'</td>'
                +'<td class="overlay hidden" data-th="toProtocol">'+item.ipProtocol+'</td>'
                +'<td class="overlay hidden " data-th="direction">'+item.direction+'</td>'
                +'</tr>'
    ))
    $("#manage_mcis_popup_sg").empty()
    $("#manage_mcis_popup_sg").append(html)

	// <thead>
	// 	<tr>
	// 		<th>spec ID</th>
	// 		<th>spec Name</th>
	// 		<th>CP</th>
	// 		<th>region</th>
	// 		<th>os type</th>
	// 		<th>Cpu / core / mem / disk</th>
	// 		<th>description</th>
	// 	</tr>
	// </thead>
	// <tbody id="assist_specList">
	// 	<!-- <tr>
	// 		<td class="btn_mtd" data-th="spec ID">aws-spec01 <span class="ov off"></span></td>
	// 		<td class="overlay hidden" data-th="spec Name">aws-spec01</td>
	// 		<td class="overlay hidden" data-th="CP">AWS</td>
	// 		<td class="overlay hidden" data-th="region">ap-northeast-1</td>
	// 		<td class="overlay hidden" data-th="os type">Amazon Linux</td>
	// 		<td class="overlay hidden" data-th="Cpu / core / mem / disk"></td>
	// 		<td class="overlay hidden" data-th="description"></td>
	// 	<tr>
}
// Spec Range 조회 실패
function filterSpecsByRangeCallbackFail(){
	
}