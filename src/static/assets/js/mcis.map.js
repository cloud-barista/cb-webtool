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
function viewMap(){
    var mcis_id = $("#mcis_id").val();
    $("#map_detail").show();
    $("#map2").empty();
    var map = map_init_target('map2')
    getGeoLocationInfo(mcis_id,map);
}
function getGeoLocationInfo(mcis_id,map){
  var JZMap = map;
  $.ajax({
    type:'GET',
    url: '/map/geo/'+mcis_id,
   // async:false,
    }).done(function(result){
        console.log("region Info : ",result)
        var polyArr = new Array();
      result = [{
        longitude: 126.990407,
        latitude:37.550246,
        Status: "Running",
        VMID: "VM-aws-developer-01",
        VMName: "VM-aws-developer-01"
      },
     
      {
        Status: "Running",
        VMID: "VM-aws-developer-02",
        VMName: "VM-aws-developer-02",
        
        longitude: 10.403993,
        latitude:51.241497,
       },
       {
        Status: "partial",
        VMID: "VM-aws-developer-03",
        VMName: "VM-aws-developer-03",
        latitude: 39.043701171875,
        longitude: -77.47419738769531
      },
       {
        Status: "Warning",
        VMID: "VM-aws-developer-04",
        VMName: "VM-aws-developer-04",
        longitude: 129.315757,
        latitude: -27.635010
       }
    ]
        for(var i in result){
            console.log("region lat long info : ",result[i])
            // var json_parse = JSON.parse(result[i])
            // console.log("json_parse : ",json_parse.longitude)
            var long = result[i].longitude
            var lat = result[i].latitude
            var fromLonLat = long+" "+lat;
            polyArr.push(fromLonLat)
            drawMap(JZMap,long,lat,result[i])
        }
        var polygon = "";
        if(polyArr.length > 1){
          polygon = polyArr.join(", ")
          polygon = "POLYGON(("+polygon+"))";
        }else{
          polygon = "POLYGON(("+polyArr[0]+"))";
        }
       
        if(polyArr.length >1){
          drawPoligon(JZMap,polygon);
        }
        //drawPoligon(map,wkt);
       
    })
}
function map_init_target(target){
 
  const osmLayer = new ol.layer.Tile({
    source: new ol.source.OSM(),
  });
  

var m = new ol.Map({
    target: target,
    layers: [
      osmLayer
    ],
    view: new ol.View({
      center: [0,0],
      zoom: 1
    })
  });
 
return m;
}
function map_init(){
 
    const osmLayer = new ol.layer.Tile({
      source: new ol.source.OSM(),
    });

  var m = new ol.Map({
      target: 'map',
      layers: [
        osmLayer
      ],
      view: new ol.View({
        center: ol.proj.fromLonLat([37.41, 8.82]),
        zoom: 0
      })
    });
  return m;
}
function drawMap(map,long,lat,info){
  var JZMap = map;
  var element = document.getElementById('popup');

  var popup = new ol.Overlay({
    element: element,
    positioning: 'bottom-center',
    stopEvent: false,
    offset: [0, -50]
  });
  
  var icon = new ol.style.Style({
    image: new ol.style.Icon({
        src:'/assets/img/marker/purple.png', // pin Image
        anchor: [0.5, 1],
        scale: 0.5
    
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
  // JZMap.addOverlay(popup);
  // JZMap.on('click',function(evt){
  //   var feature = map.forEachFeatureAtPixel(evt.pixel,function(feature){
  //     return feature;
  //   })
   
  //   if(feature){
  //     var coordinates = feature.getGeometry().getCoordinates();
  //     popup.setPosition(coordinates);
  //     $(element).empty()
  //     $(element).show()      
  //     // $(element).html('<div class="popover" role="tooltip"><div class="arrow"></div><h3 class="popover-header"></h3><div class="popover-body"></div></div>');
     
  //     $(element).popover({
  //       placement: 'top',
  //       html: true,
  //       content: "ID : "+info.VMID+"\n"+"Status :"+info.Status,
  //       title: info.VMName,
  //     });
  //     $(element).popover('show');
  //   }else{
  //     $(element).popover('disable');
  //   }
  // })
  
 
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

