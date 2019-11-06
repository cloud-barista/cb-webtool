var ChartsFlotcharts = function() {

    return {
        //main function to initiate the module


        initCharts: function() {

            if (!jQuery.plot) {
                return;
            }

            var data = [];
            var totalPoints =60;//x축 range 60초

            // 데이터를 랜덤으로 불러오기 : random data generator////
            function getRandomData() {
                if (data.length > 0) data = data.slice(1);
                // do a random walk
                while (data.length < totalPoints) {
                    var prev = data.length > 0 ? data[data.length - 1] : 50;
                    var y = prev + Math.random() * 10 - 5;
                    if (y < 0) y = 0;
                    if (y > 20) y = 20;
                    data.push(y);
                }
                // zip the generated y values with the x values
                var res = [];
                for (var i = 0; i < data.length; ++i) {
                    res.push([i, data[i]]);
                }

                return res;
            }
            ////End of random data generator/////////

            
            //////Dynamic Chart///////////////////////
            function chart4() {
                if ($('#chart_4').size() != 1) {
                    return;
                }
                //server load
                var options = {
                    series: {
                        shadowSize: 1
                    },
                    lines: {
                        show: true,
                        lineWidth: 0.5,
                        fill: true,
                        fillColor: {
                            colors: [{
                                opacity: 0.1
                            }, {
                                opacity: 1
                            }]
                        }
                    },
                    yaxis: {
                        min: 0,
                        max: 20,
                        tickColor: "#eee",
                        tickFormatter: function(v) {
                            return v + "blocks";
                        }
                    },
                    xaxis: {
                        show: true,
                        tickFormatter: function(s) {
                            return s + "s";
                        }
                    },
                    colors: ["#36c6d3"],
                    grid: {
                        tickColor: "#eee",
                        borderWidth: 0,
                    }
                };

                var updateInterval = 1000;//1초마다 업데이트/이동
                var plot = $.plot($("#chart_4"), [getRandomData()], options);

                function update() {
                    var ran = getRandomData();
                    console.log(ran);
                    plot.setData([ran]); //랜덤 데이터 
                    plot.draw(); //그리기
                    setTimeout(update, updateInterval); //업데이트
                }
                update();
            }
            //////End of Dynamic Chart////////////////

            chart4();//graph 불러오기
        },

    };
}();

jQuery(document).ready(function() {    
    ChartsFlotcharts.initCharts();
});