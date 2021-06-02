static files : assets

js 파일 구조는 static/assets/js 아래에 메뉴구조와 동일한 경로를 가진다.
( view 파일과 구조도 동일)

main.go 에 호출 경로가 있으며
화면을 return하는 경우 : mngform
화면이 없는 경우

- 리스트 조회 : Get + 업무명 + List
- 단건 조회 : Get + 업무명 + Data
- 등록 처리 : 업무명 + RegProc
- 삭제 처리 : 업무명 + DelProc

axio 호출 url 은 메뉴구조를 맞춰야 한다.
templates/MenuLeft.html

menu level1 = operation or settings
menu level2 = 1단계 밑의 folder명 ex) menu_level2_dashboards
menu level3 = 2단계 밑의 folder명 또는 파일명 ex) menu_level3_dashboardnamespace

operation > dashboards > dashboardnamespace

axios 호출 url = "/operations/dashboards/dashboardnamespace" + ...

---- call axios
ex)

    var url = "/ab/cd/";
    console.log("URL : ",url)
    // get, post, put, delete
    axios.get(url, {
        headers: {
        }
    }).then(result => {
        var statusCode = result.data.status;
        if( statusCode == 200 || statusCode == 201) {
            commonAlert("Success ");
        } else {
            var message = result.data.message;
            commonErrorAlert(statusCode, message)
        }
    }).catch(function(error){

        var statusCode = error.response.data.status;
        var message = error.response.data.message;
        commonErrorAlert(statusCode, message)

    });

---

catch 를 function 대신 object(error)로 변경
}).catch((error) => {
console.warn(error);
console.log(error.response)
});
