/*********************************************************
 * GeoNet GPS time series chart client application	******
 * modal window
 * ajax queries for static images and source data	
 * Dygraphs chart
 * 
 * baishan 5/7/2012
 **********************************************************/

var $=jQuery.noConflict();

/****** all the chart functionalities defined here ******/
var ldrChartClient = {
    //### 1. constants and vars
    ldrType:0, //0 --gnss, default, 1--utiku slides
    allColors :{
        "e":"#ff0000",
        "n":"#0088ff",
        "u": "#00ff88",
        "avgmp1":"#000000",
        "avgmp2":"#660099",
        "chisq":"#D2691E",
        "gps":"#ff0000",
        "piezo":"#0088ff",
        "rainfall": "#00ff88"
    },
    allParamDesc: {
        "e":"East (mm)",
        "n":"North (mm)",
        "u":"Up (mm)",
        "avgmp1":"Avgmp1",
        "avgmp2":"Avgmp2",
        "chisq":"Chisq",       
        "gps":"Cumulative Displacement (mm)",
        "piezo":"Groundwater Level (m)",
        "rainfall": "Daily Rainfall (mm)"
    },
    //these code-param combinations are unavailable!!
    slidesChartsUnavailable:['utks_gps','utks_piezo','utk1_rainfall','utk2_rainfall','utk3_rainfall'],
    allDygraph: {}, //by code_param
    allDygraphData:{}, //by code_param
    //data/bluf_e.json
    dataUrl : "/ldr-gnss/",
    csvDataPath :  '_design/utils/_show/csv/',
    //http://hutl13447.gns.cri.nz/ldr-charts/services/chart/akto/n/680x380.png
    chartImgPath :'services/chart/',
    chartDataPath :'services/data',
    gpsDataPath:  'services/data/gps/',
    chartWidth : 896,
    chartHeight : 512,
    ieNotice : "Interactive chart not available to Internet Explorer 8 or lower, please use a newer browser (e.g., Chrome, Firefox, IE9 or above.)",
    ieNoticeStyle: {
        'text-align' : 'center',
        'color':'#CC3300',
        'font-style': 'italic',
        'font-size':'11px'
    },
    //modalStyle:{'height': '80%','top':'40%'},
    iev: -1,
	
    //### 2. functions
    
    	
    /***
	 * init parameters, called from page
	 * ***/
    initChartParams: function (dtUrl, csvPath, imgUrl, dataPath, w, h,type){
        if(dtUrl){
            this.dataUrl = dtUrl;
        }
        if(csvPath){
            this.csvDataPath = csvPath;
        }
        if(imgUrl){
            this.chartImgPath = imgUrl;
        }
        if(dataPath){
            this.chartDataPath = dataPath;
        }
        if(w){
            this.chartWidth = w;
        }
        if(h){
            this.chartHeight = h;
        }
        if(type){
            this.ldrType = type;
        }
		
        this.iev =  this.getIEVersion();   
        //
        //init functions
        //this.initModalFunctions();
        //this.initPopoverFunctions();
        this.initChartsFunctions();
        this.initDataFunctions();
    },
    
    	
    /***
	 * re define popover for sites
	 * ***/
    initPopoverFunctions: function (){
        //popover for sites
        $('.site').popover({
            template: '<div class="popover"><div class="arrow"></div><div class="popover-inner popover-site"><div class="popover-content"><p></p></div></div></div>'
        });
		
        //popover for param
        var paramTemplate = '<div class="popover"><div class="arrow bottom"></div><div class="popover-inner popover-param"><h3 class="popover-title"></h3><div class="popover-content"><p></p></div></div></div>';
        if(ldrChartClient.ldrType === 1){//slides, use smaller size and no title
            paramTemplate = '<div class="popover"><div class="arrow bottom"></div><div class="popover-inner popover-site"><div class="popover-content"><p></p></div></div></div>';
        }
        $('.param').popover({
            placement:'bottom',
            template: paramTemplate
        });
        
        //handle img error for site map for chrome and ie
        $('.site').mouseover(function(){
            $('.site-map').error(function(){
                //console.log("error " + this.alt);
                $(this).replaceWith(this.alt);
            //$(this).hide();
            }) ;
        });
    },
    
    	
    /***
	 * init functions for modal window
	 * ***/
    initModalFunctions: function (){
        // wire up the buttons to dismiss the modal when shown
        $("#chart-modal").bind("show", function() {           
            var isMobile = navigator.userAgent.match(/(iPhone|iPod|iPad|Android|BlackBerry)/);
            if (! isMobile) {
                //scroll to top otherwise it offsets the zoom position for the map
                //don't scroll for tablets since it causes some rendering issues
                scroll(0, 0);
            }
        });
	
        // remove the event listeners when the dialog is hidden
        $("#chart-modal").bind("hide", function() {        	
            ldrChartClient.clearDgCharts();
            // remove event listeners on the buttons
            $("#chart-modal .close").unbind();
        });
	
        // finally, wire up the actual modal functionality and show the dialog
        $("#chart-modal").modal({
            "backdrop"  : "static",
            "keyboard"  : true,
            "autoResize": true,
            "show"      : false    // this parameter ensures the modal is shown immediately
        });
    },
	
		
    /***
	 * init chart function, bring up chart in modal window
	 * ***/
    initChartsFunctions: function (){
        $(window).resize(function() {
            //recheck window size
            ldrChartClient.recheckWindowSize();
        });

        $('.modal-link-a').click(function() {
			
            //recheck window size
            ldrChartClient.recheckWindowSize();
		
            //if(console) console.log("img src click" );
            var imgSrc = $(this).children("img").attr('src');
            var strs = imgSrc.split('/');
            var code = strs[strs.length -3];
            var param = strs[strs.length -2];
            if(code && param) {
                if(ldrChartClient.ldrType === 1){//slides
                    if($.inArray(code + "_" + param, ldrChartClient.slidesChartsUnavailable) >= 0){//not available
                        return false;
                    }
                }
                if(ldrChartClient.iev > 0){
                    $('#chart-modal').css('width', '80%');
                }
                //$('.modal-header').children("h3").html("Time Series Chart -- " + code.toUpperCase());
                if(ldrChartClient.iev > 0 && ldrChartClient.iev < 9){//ie show statice imageHUTL
                    //modal-image
                    var sparkelineFile = /s.png$/;
                    //console.log("chartUrl " + chartUrl);
                    if(ldrChartClient.ldrType === 1){//slides
                        sparkelineFile = /l.png$/;
                    }
                    var chartUrl = imgSrc.replace(/sparkline/, "chart").replace(sparkelineFile, ldrChartClient.chartWidth + "x" + ldrChartClient.chartHeight + ".png");
                    if($('#graphdiv').children("img") && $('#graphdiv').children("img").attr('src')){
                        //console.log("existing img" + $('#graphdiv').children("img").attr('src'));
                        $('#graphdiv').children("img").attr('src',chartUrl);
                    }else{
                        //console.log("no img yet");
                        var chartImg = $('<img>').attr('src',chartUrl);
                        // var chartNotice = $('<p>').html(ldrChartClient.ieNotice);
                        $('#graphdiv').append(chartImg).append(chartNotice);
                    }
			    
                    if($('.modal-header').children("p")&& $('.modal-header').children("p").text()){
                        $('.modal-header').children("p").html(ldrChartClient.ieNotice).css(ldrChartClient.ieNoticeStyle);
                    }else{
                        var chartNotice = $('<p>').html(ldrChartClient.ieNotice).css(ldrChartClient.ieNoticeStyle);
                        $('.modal-header').append(chartNotice);
                    }
			    
                    $('#chart-modal').modal('show');
                }else{//
                    ldrChartClient.makeDgChart(code,param);
                }
                //update data src
                var csvDataUrl = ldrChartClient.dataUrl + ldrChartClient.csvDataPath + code + '_' + param;
                if(ldrChartClient.ldrType === 1){//slides
                    csvDataUrl = ldrChartClient.dataUrl + "ldr-slides/" + ldrChartClient.csvDataPath + code + '_' + param;
                    if(param === 'gps'){//need to get from rest service
                        csvDataUrl = ldrChartClient.gpsDataPath + code.toLowerCase() + ".csv";
                    }
                }
                $('#chart-data').attr({
                    'href':csvDataUrl,
                    target: '_blank',
                    title:'Download source data in csv format.'
                });
            }
		
            return false;
        });
    },
	
    /***
	 * calculate the chart and modal size according to window size
	 * ***/
    recheckWindowSize:function(){
        var widthdiff = 0, heightdiff = 0;
        var minW = 	$(window).width() < $(document).width()? $(window).width():$(document).width();
        var minH = 	$(window).height() < $(document).height()? $(window).height():$(document).height();	
        this.chartWidth = Math.round(minW * 0.8) - 100;     
        this.chartHeight = Math.round(minH *0.8) - 100;          
        //resize dygraphs charts
        for(var key in this.allDygraph){
            if(this.allDygraph[key]){           	
                this.allDygraph[key].resize(this.chartWidth, this.chartHeight);
            }
        }
    },
	
    /*** init functions for downloading source data ***/
    initDataFunctions : function(){
        //1. form submit
        $('#data-form').submit(function() {
            var _url =  ldrChartClient.chartDataPath;
            //console.log("_url " + _url)
            $('#data-form').attr('target', '');
            $('#data-form').attr('action', _url);
            $('#data-form').attr('method', 'POST');         
            var selectedItems = '';
            $('.sparkline-check').each(function() {
                if(this.checked ){
                    selectedItems += $(this).attr('value') + ",";
                }
            });
		    
            if(selectedItems.length > 0){
                $('#data-items').attr('value',selectedItems);
                $('#ldrType').attr('value',ldrChartClient.ldrType);
                //console.log("selectedItems " + selectedItems);
                return true;
            }else{
                alert("Select items first!");
                return false;
            }
        });
		
        //2. button select all
        $('#btn-download-all').click(function() {
            //console.log("href " + $(this).attr("href"));
            window.location.href=$(this).attr("href");
        }
        );
		
        //3. buttob clear
        $('#btn-clear').click(function() {
            $('.sparkline-check').each(function() {
                this.checked = false;
            });
        }
        );
        
        //3. buttob clear
        $('#btn-select-all').click(function() {
            $('.sparkline-check').each(function() {
                if($.inArray($(this).attr("value").toLowerCase(), ldrChartClient.slidesChartsUnavailable) < 0){
                    this.checked = true;
                }
            });
        }
        );
        
        //4. checkbox sparkline-check
        $('.sparkline-check').change(function() {  
            if(ldrChartClient.ldrType === 1){//check availability
                //console.log("this.checked  " + this.checked  + " value " + $(this).attr("value") + " available " + $.inArray($(this).attr("value"), ldrChartClient.slidesChartsUnavailable)) ;
                if($.inArray($(this).attr("value").toLowerCase(), ldrChartClient.slidesChartsUnavailable) >= 0){
                    if( this.checked ){
                        this.checked = false;
                    }
                }
            }
        }
        );
    },	
	
    /*
	 * make dygraphs chart
	 */
    makeDgChart: function (code,param){
        var key = code + "_" + param;
        //console.log("makeDgChart key " + key);
        if(this.allDygraphData[key]){//data exist
            //console.log("01 makeDgChart  " );
            this.makeDygraphPlot(this.allDygraphData[key], code, param);
        }else{//fetch data from http
            this.ajaxGetPlotData(this.dataUrl,code, param);
        }
    },
       
    /*
	 * @type AJAX Event
	 * @param Void
	 */
    ajaxGetPlotData: function (dataUrl,code, param) {
        var _url = dataUrl + code.toLowerCase() + "_" + param;
        if(this.ldrType == 1){//slides
            if(param === 'gps'){//get gps displacement from rest service
                _url = this.gpsDataPath + code.toLowerCase() + ".json";
            }else{
                _url = dataUrl + "ldr-slides/" + code.toLowerCase() + "_" + param;
            }
        }
        
        //if(console) console.log("## _url " + _url);
        $.ajax({
            url: _url,
            type: 'GET',
            async: false,
            data: false,
            dataType: 'json',
            processData: false,
            mimeType: 'application/json',
            contentType: 'application/json',
            success: function(response) {
                ldrChartClient.processPlotData(response.data, code, param);
            },
            error: this.showError
        })
    },
  
    //parse, store data and make plot
    processPlotData:function  (_data, code, param){
        var chartData = this.parsePlotData(_data);
        var datakey = code + "_" + param;
        this.allDygraphData[datakey] = chartData;
        //if(console) console.log("2 datakey " + datakey) ;
        this.makeDygraphPlot(chartData, code, param);
    },

    //parse data for a single series
    parsePlotData:function (data){
        var chartData = [];
        var errorbar = true;     
        if(data && data.length){
            for (i = 0; i < data.length; i++){
                if(data[i][1] !== undefined && data[i][1] !== null && !isNaN(data[i][1])){//no null data
                    if(data[i].length == 3){//with std error
                        errorbar = true;
                        chartData.push([
                            new Date(data[i][0]) , //x
                            [data[i][1], data[i][2] ] //[y,y-err]
                            ]
                            );
                    }else{
                        chartData.push([
                            new Date(data[i][0]) ,
                            data[i][1]
                            ]);
                        errorbar = false;
                    }
                }
            }
        }
        //sort data by date
        chartData.sort(function(a, b){
            return a[0]-b[0];
        });	   
        return [chartData,errorbar];
    },


    getDygraphChartOpts:function  (index, codes, param, errorBar){
        var chtlabels = ["Date"];
        for (i = 0; i < codes.length; i++){
            chtlabels.push(codes[i] + "_" + param );
        }
        return {
            sigma: 1.0, //set the base  sigma to 1
            width: this.chartWidth,
            height: this.chartHeight,
            drawPoints : true,
            pointSize : 2,
            highlightCircleSize: 4,
            errorBars: errorBar ,
            fillAlpha:0.1,
            strokeWidth: 2,
            //legend: 'always',
            colors: [this.allColors[param]],
            xAxisLabelWidth: 100,
            axes: {
                x: {
                    valueFormatter: function(ms) {
                        return Date.toUTCTimeString (new Date(ms));
                    },
                    axisLabelFormatter: function(d) {
                        return Date.toUTCDateString (d);
                    },
                    pixelsPerLabel: 100
                },
                y: {
                    valueFormatter: function(val, opts, series_name, g) {
                        if(g && g.getSelection() > -1 && g.getOption("errorBars",series_name)){
                            var yval = g.getValue(g.getSelection(),1);
                            if(yval[1]){
                                return yval[0] + " stdE " + yval[1];
                            }
                        }
                        return val;
                    }
                }
            },
            ylabel: this.allParamDesc[param],
            labelsDivWidth: 150,
            labelsDivStyles: {
                "backgroundColor":"#FFFFFF",
                "border":"1px solid #006ACB",
                "borderRadius":"5px",
                "boxShadow":"1px 1px 4px #CCCCCC",
                "fontFamily":"Lucida Grande , Lucida Sans Unicode, Verdana, Arial, Helvetica, sans-serif",
                "fontSize":"12px",
                "fontWeight":"normal",
                "opacity":"0.85",
                "padding":"3px",
                "width":"150px"
            },
            labelFollow: true,
            labelsSeparateLines: true,
            //title: 'GPS Time series',
            verticalCrosshair: true,
            legend: 'always',
            labels: chtlabels
        };
    },
    /* remove the chart when closed */
    clearDgCharts:function(){
        //console.log("## clearDgCharts  ");
        for(var key in this.allDygraph){
            if(this.allDygraph[key]){ 
                // console.log("## clearDgCharts key " +key);
                this.allDygraph[key].destroy();
                this.allDygraph[key] = null;
            }
        }
    },
    
    /* make new dygraphs chart */
    makeDygraphPlot:function (_data, code, param){
        //show the modal
        $('#chart-modal').modal('show');
		
        //if(console) console.log("## makeDygraphPlot param "  + param  + " code " + code) ;       
        var key = code + "_" + param;
        var chartData = _data[0];
        var errorbar = _data[1];
        //check chart exist
        var opts = this.getDygraphChartOpts(0, [code], param, errorbar);
        if(!this.allDygraph[key]){
            var divId = "graphdiv" ;
            //if(console) console.log("02 key " + key);
            this.allDygraph[key] =  new Dygraph(document.getElementById(divId),
                chartData, opts);
        }else{
            //var opts = getDygraphChartOpts(0, selectedSites, param, errorbar);
            // if(console) console.log("03 key " + key);
            opts.file = chartData;
            this.allDygraph[key].updateOptions( opts );
        }
    },
    /* get IE version */
    getIEVersion : function () {
        var rv = -1; // Return value assumes failure.
        if (navigator.appName == 'Microsoft Internet Explorer') {
            var ua = navigator.userAgent;
            var re = new RegExp("MSIE ([0-9]{1,}[\.0-9]{0,})");
            if (re.exec(ua) != null)
                rv = parseFloat(RegExp.$1);
        }
        return rv;
    },
 
    showError:function (response){
    // $("#chart-modal").modal('hide');
    //if(console) console.log("showError response\n" + response);
    }
  
}


/****** misc functions ******/ 
var padNum = function (number, length) {   
    var str = '' + number;
    while (str.length < length) {
        str = '0' + str;
    }
    return str;
};

var stripNum = function (number, length) {   
    var str = '' + number;
    if (str.length > length) {
        str = str.substring(str.length - length);
    }
    return str;
};
 
Date.toUTCDateString = function (date) {
    return date.getUTCDate()  + "-"	+ (date.getUTCMonth()+1) + "-" + stripNum(date.getUTCFullYear(),4);
};
Date.toUTCTimeString = function (date) {
    return date.getUTCFullYear()  + "-"	+ (padNum((date.getUTCMonth()+1),2)) + "-" + (padNum(date.getUTCDate(),2))
    + " " + (padNum(date.getUTCHours(),2)) + ":" + (padNum(date.getUTCMinutes(),2)) + ":" + (padNum(date.getUTCSeconds(),2));
};
