
$(document).ready(function(){
    getConfig();// 세션과 상관없으므로 바로 호출.
    
    //  기존에 로그인을 바로 시키기 위해 /regUser를 한번 호출 함.
    //  axios.post("/regUser",{})
    //     .then(result =>{
    //          console.log(result);
    //      });
    // var nsUrl = "http://localhost:1234/"
     $("#sign_btn").on("click",function(){         
         try{
             var username = $("#email").val();
             var password = $("#password").val();
             if(!password || !username){
                 $("#required").modal()
                 return;
             }

             var req = {
                 username : username,
                 password : password,
             };
             console.log(req)
             axios.post("/login/proc",{headers: { },username:username,password:password })
                .then(result =>{
                    console.log(result);
                    if(result.status == 200){
                                         
                        alert("Login Success");
                        $("#popLogin").modal();
                        

                        var namespaceList = result.data.ns;
                        if(namespaceList == null || namespaceList == undefined){
                            console.log("사용자의 namespace 정보가 없음")
                        }else{
                            showUserNamespace(namespaceList);
                        }
                        
                        // getNameSpace();
                        nsModal();
                 }else{
                     alert("ID or PASSWORKD MISMATCH!!Check yourself!")
                     location.reload(true);
 
                 }
             }).catch(function(error){
                 console.log("login error : ",error);
                 alert("ID or PASSWORKD MISMATCH!!Check yourself!!")
                 location.reload(true);
             })
         }catch(e){
                 alert(e);
         }
 
     })

     $("#btnToggleNamespace").on("click",function(){         
        try{
            //addNamespaceForm
            showHideByButton("btnToggleNamespace", "addNamespaceForm")
        }catch(e){
            alert(e);
        }
    })
     
    
 })
 

 // Toggle 기능
function showHideByButton(origin, target){
    var originObj = $("#"+origin);
    var targetObj = $("#"+target)
    if( originObj.html() == 'ADD +' ){
        originObj.html('HIDE -');
        targetObj.fadeIn();
    }else{
        originObj.html('ADD +');
        targetObj.fadeOut();
    }
}

 // 엔터키가 눌렸을 때 실행할 내용
function enterKeyForLogin() {
    if (window.event.keyCode == 13) {         
         $("#sign_btn").click();
    }
}

// 커넥션 정보 조회
function getConfig(){
    var connectionURL = "/connectionconfig"
    axios.get(connectionURL,{
    }).then(result=>{
        console.log("get Connection config Data : ",result.data);
        // console.log("get Connection config Data : ",result);
        var data = result.data.connectionconfig;
        var html = ""
        
        if(data){
            data.map(item=>(
                html += '<div class="list">'
                        +'<div class="stit">'+item.ConfigName+'</div>'
                        +'<div class="stxt">'+item.ProviderName+' / '+item.RegionName+' </div>'
                        +'</div>'
            ))
            $("#cloudList").empty()
            $("#cloudList").append(html)
           configModal()
        }
    })
}

function getUserNamespace(namespaceList){    
    var html = ""
    console.log(namespaceList);
    // console.log(result);
    //최초 로그인시에는 호출되지 않도록 버그 수정
    if(!isEmpty(namespaceList) && namespaceList.length){
        namespaceList.filter((list)=> list.name !== "" ).map((item,index)=>(
            html += '<div class="list" onclick="selectNS(\''+item.id+'\');">'
                +'<div class="stit">'+item.name+'</div>'
                +'<div class="stxt">'+item.description+' </div>'
                +'</div>'                
        ))
        $("#nsList").empty();
        $("#nsList").append(html);
    }
}
// 유저의 namespace 목록 조회
function getNameSpace(){
    var url = "/ns";
    // token
    axios.get(url,{
        headers:{
            'Authorization': "{{ .apiInfo}}",
            'Content-Type' : "application/json"
        }
    }).then(result=>{
        console.log("get NameSpace Data : ",result.data);        
        var data = result.data.ns;
        var html = ""

        //최초 로그인시에는 호출되지 않도록 버그 수정
        if(!isEmpty(data) && data.length){
            data.filter((list)=> list.name !== "" ).map((item,index)=>(
                html += '<div class="list" onclick="selectNS(\''+item.id+'\');">'
                    +'<div class="stit">'+item.name+'</div>'
                    +'<div class="stxt">'+item.description+' </div>'
                    +'</div>'                
            ))
            $("#nsList").empty();
            $("#nsList").append(html);
        }
    }) .catch(function (error) {
        if (error.response) {
            // 서버가 2xx 외의 상태 코드를 리턴한 경우
        //error.response.headers / error.response.status / error.response.data
            alert("There is a problem communicating with cb-tumblebug server\nCheck the cb-tumblebug server\nCall Url : " + url + "\nStatus Code : " + error.response.status);
        }
        else if (error.request) {
            // 응답을 못 받음
            alert("No response was received from the cb-tumblebug server.\nCheck the cb-tumblebug server\nCall Url : " + url);
        }
        else {
            alert("Error communicating with cb-tumblebug server.\n" + error.message);
        }
        //console.log(error.config);
    })
}


// 뭐하는 모달이냐?
function nsModal(){
    
    $(".popboxNS .cloudlist .list").each(function () {
        $(this).click(function () {
            var $list = $(this);
            var $ok = $(".btn_ok");
                $ok.fadeIn();
            $list.addClass('selected');
            $list.siblings().removeClass("selected");
            $list.off("click").click(function(){
                if( $(this).hasClass("selected") ) {
                    $ok.stop().fadeOut();
                    $list.removeClass("selected");
                } else {
                    $ok.stop().fadeIn();
                    $list.addClass("selected");
                    $list.siblings().removeClass("selected");
                }
            });
        });
    });
}


function configModal(){
    $(".popboxCO .cloudlist .list").each(function () {
        $(this).click(function () {
            var $list = $(this);
            var $popboxNS = $(".popboxNS");
            var $arr = $('#popLogin .arr');
            var $ok = $(".btn_ok");
                $popboxNS.fadeIn();
                $arr.fadeIn();
            $list.addClass('selected');
            $list.siblings().removeClass("selected");
            $list.off("click").click(function(){
                if( $(this).hasClass("selected") ) {
                    $popboxNS.stop().fadeOut();
                    $arr.stop().fadeOut();
                    $ok.stop().fadeOut();
                    $list.removeClass("selected");
                } else {
                    $popboxNS.stop().fadeIn();
                    $arr.stop().fadeIn();
                    $list.addClass("selected");
                    $list.siblings().removeClass("selected");
                }
            });
        });
    });   
}

// namepace 생성
function createNS(){
    var namespace = $("#namespace").val()
    var desc = $("#nsDesc").val()
    
    var url = "/ns";
    var obj = {
        name: namespace,
        description : desc
    }
    if(namespace){
        axios.post(url,obj,{
            headers: { 
                'Content-type': 'application/json',
                // 'Authorization': apiInfo, 
            }
        }).then(result =>{
            console.log(result);
            if(result.status == 200 || result.status == 201){
                
                var namespaceList = result.data.ns;
                if(namespaceList == null || namespaceList == undefined){
                    console.log("사용자의 namespace 정보가 없음")
                }else{
                    showUserNamespace(namespaceList);
                }
                alert("Success Create NameSpace")

                // getNameSpace();
                $("#btnToggleNamespace").click()
                $("#namespace").val('')
                $("#nsDesc").val('')
            }else{
                alert("Fail Create NameSpace")
            }
        }).catch(function (error) {
            if (error.response) {
                // 서버가 2xx 외의 상태 코드를 리턴한 경우
            //error.response.headers / error.response.status / error.response.data
                alert("There is a problem communicating with cb-tumblebug server\nCheck the cb-tumblebug server\nCall Url : " + url + "\nStatus Code : " + error.response.status);
            }
            else if (error.request) {
                // 응답을 못 받음
                alert("No response was received from the cb-tumblebug server.\nCheck the cb-tumblebug server\nCall Url : " + url);
            }
            else {
                alert("Error communicating with cb-tumblebug server.\n" + error.message);
            }
            //console.log(error.config);
        });
    }else{
        alert("Input NameSpace")
        $("#namespace").focus()
        return;
    }
}