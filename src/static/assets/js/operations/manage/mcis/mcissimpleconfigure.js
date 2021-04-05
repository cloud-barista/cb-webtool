			// $(document).ready(function(){
				
			// })

            // getConnectionListForSelectbox 로 변경
			// function changeProvider(provider, target){
			// }

			// Connection 정보가 바뀌면 등록에 필요한 목록들을 다시 가져온다.
			function changeConnectionInfo(configName){
				console.log("config name : ",configName)
                getImageInfo(configName);
                getSecurityInfo(configName);
                getSSHKeyInfo(configName);
				getVnetInfo(configName);
				getSpecInfo(configName);
			}
			
			function getImageInfo(configName){
				console.log("1 : ",configName);
				 var configName = configName;
				 if(!configName){
					 configName = $("#regConnectionName option:selected").val();
				 }
				 console.log("2 : ",configName);
			 
				 getCommonVirtualMachineImageList("mcissimpleconfigure", "name");

				//  var url = CommonURL+"/ns/"+NAMESPACE+"/resources/image";
				//  var html = "";
				//  var apiInfo = ApiInfo
				//  axios.get(url,{
				// 	 headers:{
				// 		 'Authorization': apiInfo
				// 	 }
				//  }).then(result=>{
				// 	 console.log("Image Info : ",result.data)
				// 	 data = result.data.image
				// 	 if(!data){
				// 		 alert("등록된 이미지 정보가 없습니다.")
				// 		 location.href = "/Image/list"
				// 		 return;
				// 	 }
				// 	 for(var i in data){
				// 		 if(data[i].connectionName == configName){
				// 			 html += '<option value="'+data[i].id+'" >'+data[i].name+'('+data[i].id+')</option>'; 
				// 		 }
				// 	 }
				// 	 $("#s_imageId").empty();
				// 	 $("#s_imageId").append(html);
					 
				//  })
			}

			// 
			function setVirtualMachineImageListAtSimpleConfigure(data, sortType){
				var html = "";
				if(!data){
					alert("등록된 이미지 정보가 없습니다. 이미지 등록 화면으로 이동합니다.")
					location.href = "/Image/list"
					return;
				}
				for(var i in data){
					if(data[i].connectionName == configName){
						html += '<option value="'+data[i].id+'" >'+data[i].name+'('+data[i].id+')</option>'; 
					}
				}
				$("#s_imageId").empty();
				$("#s_imageId").append(html);
			}

 			function getSecurityInfo(configName){
				 var configName = configName;
				 if(!configName){
					 configName = $("#regConnectionName option:selected").val();
				 }
				 var url = CommonURL+"/ns/"+NAMESPACE+"/resources/securityGroup";
				 var html = "";
				 var apiInfo = ApiInfo
				 var default_sg = "";
				 axios.get(url,{
					 headers:{
						 'Authorization': apiInfo
					 }
				 }).then(result=>{
					 data = result.data.securityGroup
					 var cnt = 0
					 for(var i in data){
						 if(data[i].connectionName == configName){
							 cnt ++;
							 html += '<option value="'+data[i].id+'" >'+data[i].cspSecurityGroupName+'('+data[i].id+')</option>'; 
							if(cnt ==1 ){
								default_sg = data[i].id
							}
								
						
						}
					 }
				   
					 $("#sg").empty();
					 $("#sg").append(html);
					 $("#securityGroupIds").val(default_sg)
					 
				 })
			 }
			 function getSpecInfo(configName){
				var configName = configName;
				if(!configName){
					configName = $("#regConnectionName option:selected").val();
				}
				var url = CommonURL+"/ns/"+NAMESPACE+"/resources/spec";
				var html = "";
				var apiInfo = ApiInfo
				axios.get(url,{
					headers:{
						'Authorization': apiInfo
					}
				}).then(result=>{
					var data = result.data.spec
					console.log("spec result : ",data)
					if(data){
						html +="<option value=''>Select SpecName</option>"
						data.filter(csp => csp.connectionName === configName).map(item =>(
							html += '<option value="'+item.id+'">'+item.name+'('+item.cspSpecName+')</option>'	
						))

					}else{
						html +=""
					}
					
				  
					$("#s_spec").empty();
					$("#s_spec").append(html);
					
				})
			}
			function getSSHKeyInfo(configName){
				 var configName = configName;
				 if(!configName){
					 configName = $("#regConnectionName option:selected").val();
				 }
				 var url = CommonURL+"/ns/"+NAMESPACE+"/resources/sshKey";
				 var html = "";
				 var apiInfo = ApiInfo
				 axios.get(url,{
					 headers:{
						 'Authorization': apiInfo
					 }
				 }).then(result=>{
					 console.log("sshKeyInfo result :",result)
					 data = result.data.sshKey
					 for(var i in data){
						 if(data[i].connectionName == configName){
							 html += '<option value="'+data[i].id+'" >'+data[i].cspSshKeyName+'('+data[i].id+')</option>'; 
						 }
					 }
					 $("#sshKey").empty();
					 $("#sshKey").append(html);
					 
				 })
			 }
			function getVnetInfo(configName){
				var configName = configName;
				console.log("get vnet INfo config name : ",configName)
                if(!configName){
                    configName = $("#regConnectionName option:selected").val();
				}
				console.log("get vnet INfo config name : ",configName)
                var url = CommonURL+"/ns/"+NAMESPACE+"/resources/vNet";
                var html = "";
                var html2 = "";
                var apiInfo = ApiInfo
                axios.get(url,{
                    headers:{
                        'Authorization': apiInfo
                    }
                }).then(result=>{
                    data = result.data.vNet
					console.log("vNetwork Info : ",result)
					var init_vnet = "";
					var init_subnet = "";
					var v_net_cnt = 0
					var subnet_cnt = 0;
                    for(var i in data){
                        if(data[i].connectionName == configName){
                            html += '<option value="'+data[i].id+'" selected>'+data[i].cspVNetName+'('+data[i].id+')</option>'; 
							v_net_cnt++;
							var subnetInfoList = data[i].subnetInfoList
							if(v_net_cnt == 1){
								init_vnet = data[i].id
								console.log("init_vnet :",init_vnet)
							}
							
                            for(var k in subnetInfoList){
								
									init_subnet = subnetInfoList[0].IId.NameId
									console.log("init_subnet :",init_subnet)
							
                                html2 += '<option value="'+subnetInfoList[k].IId.NameId+'" >'+subnetInfoList[k].IPv4_CIDR+'</option>'; 
                            }
                        }
                    }
                    $("#vnet").empty();
                    $("#vnet").append(html);
                    $("#subnet").empty();
					$("#subnet").append(html2);
					
					//setting default
					$("#s_subnetId").val(init_subnet);
					$("#s_vNetId").val(init_vnet);				
                    
                })
            }
			
			function getConnectionInfo(provider, target){
                // var url = SpiderURL+"/connectionconfig";
                console.log("provider : ",provider)
				if(!provider){
					alert("Please select Provider!!")
					$("#regProvider").focus();
					return;
				}
                //var provider = $("#provider option:selected").val();
                var html = "";
                var apiInfo = ApiInfo
                axios.get(url,{
                    headers:{
                        'Authorization': apiInfo
                    }
                }).then(result=>{
                    var data = result.data.connectionconfig
                    console.log("connection data : ",data);
					if(data){
						data.filter(item => item.ProviderName === provider).map(csp=>(
							html += '<option value' 
						))
					}
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
                    
                    $("#"+target).empty();
                    $("#"+target).append(html);
               

                    
                })
            }

            // networkmng.js 에 동일한 function있으므로 참고할 것
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
                    var data = result.data.connectionconfig
                    console.log("connection data : ",data);

					if(data){
						data.filter(csp => csp.ProviderName === provider).map(item =>(
							html += '<option value="'+item.ConfigName+'" item="'+item.ProviderName+'">'+item.ConfigName+'</option>'
						))
					}
                   
                    $("#regConnectionName").empty();
					$("#regConnectionName").append(html);
					
                    getImageInfo(configName);
                    getSecurityInfo(configName);
                    getSSHKeyInfo(configName);
                    getVnetInfo(configName);

                    
                })
            }
			
						
			const Simple_Server_Config_Arr = new Array();
			var simple_data_cnt = 0
			const cloneObj = obj=>JSON.parse(JSON.stringify(obj))
			function simple_btn(){
				var simple_form = $("#simple_form").serializeObject()
				var server_name = simple_form.name
				var server_cnt = parseInt(simple_form.vm_add_cnt)
				console.log('server_cnt : ',server_cnt)
				var add_server_html = "";
				
				if(server_cnt > 1){
					for(var i = 1; i <= server_cnt; i++){
						var new_vm_name = server_name+"-"+i;
						var object = cloneObj(simple_form)
						object.name = new_vm_name
						
						add_server_html +='<li onclick="view_simple(\''+simple_data_cnt+'\')">'
								+'<div class="server server_on bgbox_b">'
								+'<div class="icon"></div>'
								+'<div class="txt">'+new_vm_name+'</div>'
								+'</div>'
								+'</li>';
						Simple_Server_Config_Arr.push(object)
						console.log(i+"번째 Simple form data 입니다. : ",object);
					}
				}else{
					Simple_Server_Config_Arr.push(simple_form)
					add_server_html +='<li onclick="view_simple(\''+simple_data_cnt+'\')">'
									+'<div class="server server_on bgbox_b">'
									+'<div class="icon"></div>'
									+'<div class="txt">'+server_name+'</div>'
									+'</div>'
									+'</li>';

				}
				$(".simple_servers_config").removeClass("active");
				$("#mcis_server_list").prepend(add_server_html)
				console.log("simple btn click and simple form data : ",simple_form)
				console.log("simple data array : ",Simple_Server_Config_Arr);
				simple_data_cnt++;
				$("#simple_form").each(function(){
					this.reset();
				})
			}
			function view_simple(cnt){
				console.log('view simple cnt : ',cnt);
				var select_form_data = Simple_Server_Config_Arr[cnt]
				console.log('select_form_data : ', select_form_data);
				$(".simple_servers_config").addClass("active")
				$(".new_servers_config").removeClass("active")

			}
			
