function getIPStackRegion(ip_address){
    var apiUrl = "http://api.ipstack.com/"
    var access_key = "86c895286435070c0369a53d2d0b03d1"
    var url = apiUrl+ip_address+"?access_key="+access_key

    console.log("api get region url:",url);
    axios.get(url).then((result)=>{
        console.log("api get result : ",result);
    })
}

function test_setFocusRegion(code){
    var map = jQuery('#vmap').vectorMap('get','mapObject');
    map.setFocus = code;
}


function escapeXml(string) {
    return string.replace(/[<>]/g, function (c) {
      switch (c) {
        case '<': return '\u003c';
        case '>': return '\u003e';
      }
    });
  }

//   function zoomIn(){
//     jQuery('#vmap').vectorMap('zoomIn');
//   }
//   function zoomOut(){
//     jQuery('#vmap').vectorMap('zoomOut');
//   }