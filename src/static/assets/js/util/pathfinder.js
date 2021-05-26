// path 와 매핑되는 controller의 이름 = key가 되어 
// 해당 key입력 시 main.go의 path를 return
// 필요한 param을 path에 적용하여 호출 url return
// leftmenu에서 script import



// map에 담긴 Key를 value로 바꿔 url을 return한다.
// url에는 main.go 에서 사용하는 path를 넣는다.
function setUrlByParam(url, urlParamMap){
    //resultVmCreateMap.set(resultVmKey, resultStatus)
    // var url = "/operation/manages/mcksmng/:clusteruID/:clusterName/del/:nodeID/:nodeName";
    var returnUrl = url;
    for (let key of urlParamMap.keys()) { 
        console.log("urlParamMap " + key + " : " + urlParamMap.get(key) );
        
        var urlParamValue = urlParamMap.get(key)
        returnUrl = returnUrl.replace(key, urlParamValue);        
    }
    return returnUrl;
}