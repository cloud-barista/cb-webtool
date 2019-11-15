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

function getNameSpace(){
    var url = CommonURL+"/ns"
    axios.get(url).then(result =>{
        var data = result.data.ns
        var namespace = ""
        for( var i in data){
            if(i == 0 ){
                namespace = data[i].id
            }
        }
        $("#namespace1").val(namespace);

    })
}
