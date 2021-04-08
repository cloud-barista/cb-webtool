
$(document).ready(function(){
    //  기존에 로그인을 바로 시키기 위해 /regUser를 한번 호출 함.
    //  axios.post("/regUser",{})
    //     .then(result =>{
    //          console.log(result);
    //      });
    // var nsUrl = "http://localhost:1234/"
     $("#sign_btn").on("click",function(){         
         try{
            //  var email = $("#email").val();
            var userID = $("#userID").val();
            var password = $("#password").val();
            if(!password || !userID){
                $("#required").modal()
                return;
            }

            var req = {
                userID : userID,
                password : password,
            };
            console.log(req)
            axios.post("/login/proc",{
                headers: { },
                userID:userID,password:password 
                })
                .then(result =>{
                    console.log(result);
                    if(result.status == 200){
                        console.log("get result Data : ",result.data.LoginInfo);
                        tokenSuccess(result.data.LoginInfo)
                        
                        // location.href = "/setting/connections/cloudconnectionconfig/mngform" // --> TODO : Dashboard로 보낼 것, namespace 없을 때만 connection으로
                        location.href = "/main/" // --> TODO : Dashboard로 보낼 것, namespace 없을 때만 connection으로
                 }else{
                     alert("ID or PASSWORKD MISMATCH!!Check yourself!")
                    //  location.reload(true); 
                 }
             }).catch(function(error){
                 console.log("login error : ",error);
                 alert("ID or PASSWORKD MISMATCH!!Check yourself!!")
                //  location.reload(true);
             })
         }catch(e){
                 alert(e);
         }
 
     })

     // namespace 등록영역 보이기/숨기기
     $("#btnToggleNamespace").on("click",function(){         
        try{
            //addNamespaceForm
            showHideByButton("btnToggleNamespace", "addNamespaceForm")
        }catch(e){
            alert(e);
        }
    })
 })
 
 // 로그인 성공 시 Token저장
 function tokenSuccess(loginInfo) {    
    // console.log("username1 = " + loginInfo['Username'])
    // console.log("username2 = " + loginInfo.Username)
    // console.log("AccessToken = " + loginInfo['AccessToken'])
    // localStorage.Username = loginInfo['Username'];
    // localStorage.AccessToken = loginInfo['AccessToken'];    
    // localStorage.LLL = "ABCD";
    console.log(loginInfo)
    // alert("tokenSuccess")
    document.cookie = "UserID="+loginInfo['UserID'] +";AccessToken="+ loginInfo['AccessToken'];
    
}

 // Toggle 기능 : 원래는 namespace와 connection driver에서 사용한 것 같으나 현재는 namespace만 사용. 그럼 굳이 function으로 할 필요있나?
 // 버튼의 display를 ADD / HIDE, 대상 area를 보이고 숨기고
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

// 커넥션 정보 조회 : getConfig() -> getCloudConnectionConfig 로 변경
function getCloudConnectionConfig(){
    var connectionURL = "/connections/cloudconnectionconfig/list"
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
    if(namespaceList == null || namespaceList == undefined){
        console.log("사용자의 namespace 정보가 없음")
        return;
    }

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


// namespace 선택modal
function nsModal(){
    console.log("nsModal called");
    $(".popboxNS .cloudlist .list").each(function () {
        $(this).click(function () {
            var $list = $(this);
            // var $ok = $(".btn_ok");// --class 말고 id로
            var $ok = $("#select_ns_ok_btn");
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
    console.log("configModal called");
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
//function createNS(){
function createNameSpace(){
    var addNamespaceValue = $("#addNamespace").val()
    var addNamespaceDescValue = $("#addNamespaceDesc").val()
    
    //var url = "/ns";
    var url = "/setting/namespaces/namespace/reg/proc";
    var obj = {
        name: addNamespaceValue,
        description : addNamespaceDescValue
    }
    console.log(obj)
    if(addNamespaceValue){
        axios.post(url,obj,{
            headers: { 
                'Content-type': 'application/json',
                // 'Authorization': apiInfo, 
            }
        }).then(result =>{
            console.log(result);
            if(result.status == 200 || result.status == 201){
                var namespaceList = result.data.nsList;
                getUserNamespace(namespaceList)
                
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
        $("#addNamespace").focus()
        return;
    }
}

function selectNS(ns){
    console.log("select namespace : ",ns)
    $("#sel_ns").val(ns);
}


function clickOK(){
    var select_ns =   $("#sel_ns").val();
    console.log("slect ns is : ",select_ns);
    setNS(select_ns);   
}

function setNS(nsid){
    if(nsid){
        reqUrl = "/setting/namespaces/"+nsid;
        console.log(reqUrl);
        axios.get(reqUrl,{
            headers:{
                'Authorization': "{{ .apiInfo}}"
            }
        }).then(result=>{
            var data = result.data
            console.log(data);
            location.href = "/Dashboard/NS"
        }).catch(function (error) {
            if (error.response) {
                console.log(error.response)              
            }
            else if (error.request) {
                console.log(error.request)
            }
            else {
                console.log(error.message)
            }
        })
        
    }else{
        alert("NameSpace가 등록되어 있지 않습니다.\n등록페이지로 이동합니다.")
        
    }
    
} 