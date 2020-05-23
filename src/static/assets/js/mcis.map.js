//서버에서 처리 필요 없다.ㅜㅡ
function getIPStackRegion(ip_address){
    var apiUrl = "http://api.ipstack.com/"
    var access_key = "86c895286435070c0369a53d2d0b03d1"
    var url = apiUrl+ip_address+"?access_key="+access_key

    console.log("api get region url:",url);
    axios.get(url).then((result)=>{
        console.log("api get result : ",result);
        var data = result.data
        var lat = data.latitude
        var long = data.longitude
        
    })
}
function getGeoLocationInfo(mcis_id,map){
  $.ajax({
    type:'GET',
    url: '/map/geo/'+mcis_id,
   // async:false,
    }).done(function(result){
        console.log("region Info : ",result)
        var polyArr = new Array();
        for(var i in result){
            console.log("region lat long info : ",result[i])
            // var json_parse = JSON.parse(result[i])
            // console.log("json_parse : ",json_parse.longitude)
            var long = result[i].longitude
            var lat = result[i].latitude
            var fromLonLat = long+" "+lat;
            polyArr.push(fromLonLat)
            drawMap(map,long,lat)
        }
        var polygon = "";
        if(polyArr.length > 1){
          polygon = polyArr.join(", ")
          polygon = "POLYGON(("+polygon+"))";
        }else{
          polygon = "POLYGON(("+polyArr[0]+"))";
        }
        drawPoligon(map,polygon);
    })
}
function map_init(){
  if(!JZMap){
    const osmLayer = new ol.layer.Tile({
      source: new ol.source.OSM(),
    });

    var JZMap = new ol.Map({
      target: 'map',
      layers: [
        osmLayer
      ],
      view: new ol.View({
        center: ol.proj.fromLonLat([37.41, 8.82]),
        zoom: 2
      })
    });
  }
}
function drawMap(JZMap,long,lat){
  var icon = new ol.style.Style({
    image: new ol.style.Icon({
        src:'/assets/img/marker/purple.png', // pin Image
        anchor: [0.5, 1],
    
    })
})
  var map_center = ol.proj.fromLonLat([long, lat]);
  var point_gem = new ol.geom.Point(map_center);
  var point_feature = new ol.Feature(point_gem);
  point_feature.setStyle([icon])
  var stackVectorMap = new ol.source.Vector({
    features : [point_feature]
  })

  var stackLayer = new ol.layer.Vector({
    source: stackVectorMap
  })
  JZMap.addLayer(stackLayer)
  
 
}

function drawPoligon(JZMap,polygon){
  var wkt = polygon;
  console.log(wkt)
  var format = new ol.format.WKT();

  var feature = format.readFeature(wkt, {
    dataProjection: "EPSG:4326",
    featureProjection: "EPSG:3857"
  });
  var stackVectorMap = new ol.source.Vector({
    features : [feature]
  })

  var stackLayer = new ol.layer.Vector({
    source: stackVectorMap
  })
  JZMap.addLayer(stackLayer);
  
}



function escapeXml(string) {
    return string.replace(/[<>]/g, function (c) {
      switch (c) {
        case '<': return '\u003c';
        case '>': return '\u003e';
      }
    });
  }

