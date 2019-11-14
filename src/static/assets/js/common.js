funtcion requestAjax(url, method, data){
    console.log("Request URL : ",url)
    var met = method.toLowerCase
    $.ajax({
        url : url,
        type: method,
        data: data

    }).then(function(result){
        console.log(result)
    })
}

function requestAxios()
