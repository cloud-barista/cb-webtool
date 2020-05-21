var config = {
    type: 'line',
    data: {
        labels:[] ,// 시간을 배열로 받아서 처리
        datasets: [{
            label : "",//cpu 관련 내용들 
            //backgroundColor:window.chartColors.red,
           // borderColor:window.chartColors.red,
            data:[],//
        }]

    }
}



function time_arr(obj, title){
    //data sets
   var labels = obj.columns;
   var datasets = obj.values;
    // 각 값의 배열 데이터들
   
   var series_label = new Array();
   var data_set = new Array();
   // 최종 객체 data
   var new_obj = {}
   var color_arr = ['rgb(255, 99, 132)','rgb(255, 159, 64)', 'rgb(255, 205, 86)','rgb(75, 192, 192)','rgb(54, 162, 235)','rgb(153, 102, 255)','rgb(201, 203, 207)','rgb(99, 255, 243)']   

   for(var i in labels){
    var dt = {}  
    var series_data = new Array();  
    for(var k in datasets){
        if(i == 0){
            series_label.push(datasets[k][i]) //이건 시간만 담는다.
        }else{
            dt.label = labels[i];
            series_data.push(datasets[k][i]) //그외 나머지 데이터만 담는다.
            dt.borderColor = color_arr[i];
            dt.backgroundColor = color_arr[i];
            dt.fill= false;
           // dt.data
        }  
    }
    if(i > 0){
       dt.data = series_data
       data_set.push(dt)
    }
   
    
   
   }
   console.log("data set : ",data_set);
   console.log("time series : ",series_label);
   new_obj.labels = series_label //시간만 담김 배열
   new_obj.datasets =  data_set//각 데이터 셋의 배열
   console.log("Chart Object : ",new_obj);
   config.type = 'line',
   config.data = new_obj
   config.options = {
    responsive: true,
    title: {
        display: true,
        text: title
    },
    tooltips: {
        mode: 'index',
        intersect: false,
    },
    hover: {
        mode: 'nearest',
        intersect: true
    },
    scales: {
        x: {
            display: true,
            scaleLabel: {
                display: true,
                labelString: 'Time'
            }
        },
        y: {
            display: true,
            scaleLabel: {
                display: true,
                labelString: 'Value'
            }
        }
    }
}
   return config;
}

window.chartColors = {
	red: 'rgb(255, 99, 132)',
	orange: 'rgb(255, 159, 64)',
	yellow: 'rgb(255, 205, 86)',
	green: 'rgb(75, 192, 192)',
	blue: 'rgb(54, 162, 235)',
	purple: 'rgb(153, 102, 255)',
    grey: 'rgb(201, 203, 207)',
    mint: 'rgb(99, 255, 243)'
};