static files : assets



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
----