<script setup lang="ts">

    import Breadcrumbs from "@/components/Breadcrumbs.vue"
    import { onMounted, ref, watch, nextTick, computed } from "vue"
    import { piniaStock } from '@/store/piniaStock'

  import { useSeoMeta } from '@unhead/vue'
  useSeoMeta({
    title: '儀錶板|台股恐慌資訊站 TwVixStock',
    description: '台股恐慌資訊站提供加權、恐慌、本益比指數已及AI估值分析',
    ogImage: 'http://twvixstock.qzz.io/images/profile_photos/small_2026_twvix2.jpg',
    ogTitle: '儀錶板|台股恐慌資訊站 TwVixStock',
    ogType: 'website',
    ogLocale: 'zh_TW',
    keywords: '台股, 恐慌指數, VIX, 股票投資, TwVixStock, 台灣加權, 本益比, 儀錶板',
  });

    // 使用 pinia
    const piniaStockMain = piniaStock()

    // 建議值
    const vixSuggest = 25;
    const peSuggest = 23;

    // 系列名
    const seriesNameEX  = "台灣加權指數 TAIEX";
    const seriesNameVIX = "台灣恐慌指數 VIXTWN";
    const seriesNamePE  = "台灣本益比估值 TAI";



// 折線圖

    // 定義圖表用的參數
    let DateArr = ref<string[]>([])
    let TAIEX = ref<number[]>([])
    let VIXTWN = ref<number[]>([])
    let VIXTWN_base = ref<number[]>([])  //vix 計算的基準
    let TAIPE = ref<number[]>([])
    let TAIPE_base = ref<number[]>([])  //pe 計算的基準
    let hidden_EX = ref<boolean>(false)
    let hidden_VIX = ref<boolean>(false)
    let hidden_PE = ref<boolean>(false)

    // 運算用的
    const myChart = ref(null);  //圖表本體
    let vixIsReverse = ref<boolean>(false);  //VIX 是否顛倒
    let vixIsInflation = ref<boolean>(false);  //VIX 是否通膨
    let peIsInflation = ref<boolean>(false);  //PE 是否通膨
    let showDays = ref<number>(0);  //顯示日子數量
    let lastDate = ref<string>("0000-00-00");  //資料日期
    let lastEX = ref<number>(0);  //當前EX
    let lastVIX = ref<number>(0);  //當前VIX
    let lastPE = ref<number>(0);  //當前PE

    // GEMIMI用了
    let geminiDate = ref<string>("0000-00-00");
    let geminiPoint = ref<number>(0); 
    let geminiPrice = ref<number>(0); 
    let geminiReason = ref<string>("Loading...");

    // =============================== 以下計算 ===============================

    // 設定日期範圍
    const setDateRange = (days:number)=>{

        // 記錄當前隱藏系列並設置
        getHiddenSeries();

        // 日子參數
        showDays.value = days;

        // 重新初始化參數
        DateArr.value = [];
        TAIEX.value = [];
        VIXTWN.value = [];
        VIXTWN_base.value = [];
        TAIPE.value = [];
        TAIPE_base.value = [];

        // 插入參數
        let nowstamp = new Date().getTime(); //現在時間戳
        piniaStockMain.DateArr.forEach((v,k) => {
            //紀錄的日期時間戳
            let timestamp = new Date(v.replace(/-/g, "/")).getTime();

            //顯示時間吻合才顯示
            if ( timestamp > (nowstamp - (86400000*days)) ){
                DateArr.value.push(piniaStockMain.DateArr[k])
                TAIEX.value.push(piniaStockMain.TAIEX[k])
                VIXTWN.value.push(piniaStockMain.VIXTWN[k])
                VIXTWN_base.value.push(piniaStockMain.VIXTWN[k])
                TAIPE.value.push(piniaStockMain.TAIPE[k])
                TAIPE_base.value.push(piniaStockMain.TAIPE[k])
            }
        });

        //運算  vix
        vixCompute(); 

        //運算 pe
        peCompute();

        // 刷新 x 軸文字
        nextTick(()=>{
            (myChart.value as any).updateOptions({
                labels: DateArr.value 
            });
        });

        // 隱藏應該隱藏的系列
        //setHiddenSeries(hidedenSeriesName);
    }

    // 監控變化 進行運算
    watch(() => vixIsReverse, (newval)=>{    
        vixCompute();
    }, {deep:true});
    watch(() => vixIsInflation, (newval)=>{    
        vixCompute();
    }, {deep:true});
    watch(() => peIsInflation, (newval)=>{    
        peCompute();
    }, {deep:true});


    // vix 運算(依據顛倒 通膨旗標)
    const vixCompute = ()=>{

        // 記錄當前隱藏系列並設置
        getHiddenSeries();

        //還原基準 vix
        VIXTWN.value = [...VIXTWN_base.value]; 

        //顛倒旗標
        let AvgTmp = 0;
        if(vixIsReverse.value){
            //先算平均值
            VIXTWN.value.forEach((v)=>{
                AvgTmp += v;
            })
            AvgTmp /= VIXTWN.value.length;

            //再顛倒
            VIXTWN.value.forEach((v, k)=>{
                VIXTWN.value[k] = (v*-1) + (AvgTmp*2); //乘 -1 加兩倍平均 就是顛倒
            })
        }

        //通膨旗標
        if(vixIsInflation.value){
            VIXTWN.value = inflationCompute(VIXTWN.value);
        }

        // 隱藏應該隱藏的系列
        //setHiddenSeries(hidedenSeriesName);
    }

    // pe 運算(依據 通膨旗標)
    const peCompute = ()=>{

        // 記錄當前隱藏系列並設置
        getHiddenSeries();

        //還原基準 pe
        TAIPE.value = [...TAIPE_base.value];

        //通膨旗標
        if(peIsInflation.value){
            TAIPE.value = inflationCompute(TAIPE.value);
        }

        // 隱藏應該隱藏的系列
        //setHiddenSeries(hidedenSeriesName);
    }

    //通膨率運算
    const inflationCompute = (NumArr:number[]):number[] =>{

        //平均保存備用
        let AvgTmp = 0;
        NumArr.forEach((v)=>{
            AvgTmp += v;
        })
        AvgTmp /= NumArr.length;

        //先算 TAIEX.value 通膨率
        let exAvgL = 0;
        let exAvgR = 0;
        for(let i=0; i<(TAIEX.value.length/2-1); i++){
            exAvgL += TAIEX.value[i];  //左半部價格
        }
        for(let i=(TAIEX.value.length-1); i>TAIEX.value.length/2; i--){
            exAvgR += TAIEX.value[i];  //右半部價格
        }

        //半個視窗的通膨率
        let rate = (exAvgR / exAvgL) - 1;

        //套用通膨率
        let ni = (NumArr.length/2-1); //長度一半
        for(let i=0; i<ni; i++){
            let Lrate = (ni-i)/ni; //左半部，越左越1 越中越0 
            NumArr[i] = NumArr[i] - (AvgTmp * (rate * Lrate)); //調整左側
        }
        ni = (NumArr.length/2);
        for(let i=(NumArr.length-1); i>NumArr.length/2; i--){
            let Rrate = (i-ni)/ni; //右半部，越右越1 越中越0 
            NumArr[i] = NumArr[i] + (AvgTmp * (rate * Rrate)); //調整右側
        }

        return NumArr;
    }
    

        //顯示隱藏系列相關方法整理
            // myChart.value.chart.w.globals.seriesNames  全系列名
            // myChart.value.chart.w.globals.collapsedSeries  被隱藏的系列 data
            // myChart.value.chart.w.globals.collapsedSeriesIndices  被隱藏的系列 id
            // myChart.value.hideSeries("台灣加權指數 TAIEX")  執行隱藏
            // myChart.value.showSeries("台灣加權指數 TAIEX")  執行顯示

    // 取得隱藏系列名
    const getHiddenSeries = ()=>{
        // 還原顯示參數
        hidden_EX.value = false; 
        hidden_VIX.value = false; 
        hidden_PE.value = false; 
        myChart.value.chart.w.globals.collapsedSeriesIndices.forEach((v,_) => {
            // 系列名
            let nameTmp = myChart.value.chart.w.globals.seriesNames[v]; 

            // 設置預設是否隱藏
            if ( nameTmp == seriesNameEX){ hidden_EX.value = true; }
            if ( nameTmp == seriesNameVIX){ hidden_VIX.value = true; }
            if ( nameTmp == seriesNamePE){ hidden_PE.value = true; }
        });
    };

    // =============================== 以下初始 ===============================

    // pinia GEMINI 通知初始完成時(初次進首頁) --> 初始
    watch(() => piniaStockMain.InitialGeminiReady, (newval, oldval)=>{    
        if(newval){ //初始化剛剛完成
            initialGemini(); //GEMINI參數
        }
    }, {deep:true});

    // pinia STOCK 通知初始完成時(初次進首頁) --> 初始
    watch(() => piniaStockMain.InitialStockReady, (newval, oldval)=>{    
        if(newval){ //初始化剛剛完成
            initialChart(); //繪圖
            initialStock(); //STOCK參數初始
        }
    }, {deep:true});

    // 掛載時(route進來的時候) --> 初始
    onMounted(async (): Promise<void> => {
        if(piniaStockMain.InitialStockReady){ //Stock初始化早就完成了
            initialChart(); //圖表初始
            initialStock(); //STOCK參數初始
        }
        if(piniaStockMain.InitialGeminiReady){ //gemini初始化早就完成了
            initialGemini(); //GEMINI參數初始
        }
    })

    //圖表初始
    const initialChart = ()=>{  //圖表
        setTimeout(()=>{ //等畫布完全載入
            setDateRange(365);    //給初始日期
        }, 200);
    }

    //STOCK參數初始
    const initialStock = ()=>{
        lastDate.value = piniaStockMain.DateArr[ piniaStockMain.DateArr.length -1 ]  //資料日期
        lastEX.value = piniaStockMain.TAIEX[ piniaStockMain.TAIEX.length -1 ]
        lastPE.value = piniaStockMain.TAIPE[ piniaStockMain.TAIPE.length -1 ]
        lastVIX.value = piniaStockMain.VIXTWN[ piniaStockMain.VIXTWN.length -1 ]
    }

    //GEMINI參數初始
    const initialGemini = ()=>{  
        geminiDate.value = piniaStockMain.geminiData.data.Date;
        geminiPoint.value = piniaStockMain.geminiData.data.Point;
        geminiPrice.value = piniaStockMain.geminiData.data.Price;
        geminiReason.value = piniaStockMain.geminiData.data.Reason;
    }


    // =============================== 以下圖表 折線圖 ===============================

    // 1. 資料定義：series 陣列裡放幾個物件，就有幾條線
    const series = ref([
        {
            name: seriesNameEX,
            type: 'area',
            //data: [2300, 1100, 2200, 2700, 1300, 2200, 3700, 2100, 4400, 2200, 3000]
            data: TAIEX,
            hidden: hidden_EX
        },{
            name: seriesNameVIX,
            type: 'line',
            //data: [84, 95, 81, 97, 82, 73, 81, 71, 96, 87, 43]
            data: VIXTWN,
            hidden: hidden_VIX
        },{
            name: seriesNamePE,
            type: 'line',
            //data: [30, 25, 36, 30, 45, 35, 64, 52, 59, 36, 39]
            data: TAIPE,
            hidden: hidden_PE
        }
    ]);

    // 2. 圖表配置
    const chartOptions = ref({
        colors: ['#a855f7', '#FF4560', '#3b82f6'],
        chart: {
            //height: 350,
            //type: 'line',
            stacked: false,
            parentHeightOffset: 0,
            toolbar: { show: true },
            zoom: { enabled: false },
            selection: { enabled: false },
        },
        stroke: {
            width: [1.5, 1.5, 1.5],
            //curve: 'smooth'
        },
        fill: {
            opacity: [0.7, 0.7, 0.7],
            type: ['gradient', null, null],
            gradient: {
                inverseColors: false,
                shade: 'dark',            // 改為 dark 會增加顏色的深邃感
                type: "vertical",
                opacityFrom: 0.85,
                opacityTo: 0.2,         // 底部依然保持清爽
                stops: [20, 100, 100], // 在 50% 處有一個過渡，層次更明顯
                colorStops: []    // 如果你想自訂顏色顏色，可以在這裡精確控制
            }
        },
        //labels: ['01/01/2003', '02/01/2003', '03/01/2003', '04/01/2003', '05/01/2003', '06/01/2003', '07/01/2003','08/01/2003', '09/01/2003', '10/01/2003', '11/01/2003'],
        labels: DateArr,
        markers: {
            size: 0, // 平時隱藏點
            hover: {
                size: 5    // 滑鼠移上去才顯現
            }
        },
        xaxis: {
            type: 'datetime',
            labels: {
                datetimeFormatter: {
                    year: 'yyyy',
                    month: 'yyyy-MM',
                    day: 'MM-dd',
                }
            }
        },
        yaxis: [
            {
                seriesName: seriesNameEX, // 與 series 裡的 name 對應
                //title: { text: "TEAM A (Points)" },
                    labels: {
                    formatter: function (val) {
                        return val.toFixed(0) / 1000 + "k"; 
                    },
                    offsetX: -15,
                }
            },{
                seriesName: [seriesNameVIX, seriesNamePE],
                opposite: true, // 放在右邊
                //title: { text: "TEAM B B (Percentage)" },
                labels: {
                    formatter: function (val) {
                        return val.toFixed(0); 
                    },
                    offsetX: -20,
                }
            }
        ],
        tooltip: {
            shared: true,
            intersect: false,
            y: {
                formatter: function (y) {
                    if (typeof y !== "undefined") {
                        return y.toFixed(2) + " 點";
                    }
                    return y;
                }
            },
            x: {
                format: 'yyyy-MM-dd' // 這裡可以自訂格式
            }
        },
        grid: {
            padding: {
                left: -5,
                right: -15
            }
        }
    });


    // =============================== 以下圖表 圓環圖 ===============================

    const series2  = computed(() => [geminiPoint.value]);
    //const series2 = ref([80]);
    const chartOptions2 = ref({
        colors: ['#d96570'],
        chart: {
            height: 350,
            type: 'radialBar',
            toolbar: { show: false },
            zoom: { enabled: false },
            selection: { enabled: false },
        },
        plotOptions: {
            radialBar: {
            startAngle: -135,
            endAngle: 225,
            hollow: {
                margin: 0,
                size: '65%',
                background: '#fff',
                image: undefined,
                imageOffsetX: 0,
                imageOffsetY: 0,
                position: 'front',
                dropShadow: {
                    enabled: true,
                    top: 3,
                    left: 0,
                    blur: 4,
                    opacity: 0.5
                }
            },
            track: {
                background: '#fff',
                strokeWidth: '67%',
                margin: 0, // margin is in pixels
                dropShadow: {
                    enabled: true,
                    top: -3,
                    left: 0,
                    blur: 4,
                    opacity: 0.7
                }
            },
            dataLabels: {
                show: true,
                name: {
                    offsetY: -10,
                    show: true,
                    color: '#555',
                    fontSize: '17px'
                },
                value: {
                    formatter: function(val) {
                        return parseInt(val);
                    },
                    color: '#333',
                    fontSize: '42px',
                    show: true,
                }
            }
            }
        },
        fill: {
            type: 'gradient',
            gradient: {
                shade: 'light',
                type: 'horizontal',
                gradientToColors: ['#4285f4'], // 最後顏色 (藍)
                stops: [0, 33, 66, 100], // 顏色位置
                colorStops: [
                    { offset: 0, color: '#d96570' },
                    { offset: 33, color: '#9b72cb' },
                    { offset: 66, color: '#4285f4' },
                    { offset: 100, color: '#4285f4' }
                ]
            }
        },
        stroke: {
            lineCap: 'round'
        },
        labels: ["📈台灣加權指數評分"],
    });

    // =============================== 以下圖表 長條圖 ===============================

    const series3 = computed(() => [{
        name: 'Actual',
        data: [{
            x: '恐慌指數VIX',
            y: lastVIX.value,
            fillColor: '#FF4560',
            goals: [{
                name: 'Expected',
                value: vixSuggest,
                strokeWidth: 4,
                strokeHeight: 20,
                strokeColor: '#00f0f9',
                strokeDashArray: 8 
            }]
        },{
            x: '本益比PE',
            y: lastPE.value,
            fillColor: '#3b82f6',
            goals: [{
                name: 'Expected',
                value: peSuggest,
                strokeWidth: 4,
                strokeHeight: 20,
                strokeColor: '#00f0f9',
                strokeDashArray: 8
            }]
        }]
    }]);
    const chartOptions3 = {
        chart: {
            height: 200,
            type: 'bar'
        },
        plotOptions: {
            bar: {
                horizontal: true,
            }
        },
        colors: ['#00E396'],
        fill: {
            type: 'gradient',
            gradient: {
                shade: 'dark',
                type: 'horizontal',
                shadeIntensity: 0.4,
                gradientToColors: ['#a855f7'],
                opacityFrom: 1,
                opacityTo: 0.85,
                stops: [0, 90, 100]
            }
        },
        dataLabels: {
            formatter: function(val, opt) {
                const goals = opt.w.config.series[opt.seriesIndex].data[opt.dataPointIndex].goals;
                if (goals && goals.length) {
                    return `${val} / ${goals[0].value}`;
                }
                return val;
            }
        },
        legend: {
            show: true,
            showForSingleSeries: true,
            customLegendItems: ['當前值', '合理值'],
            markers: {
                fillColors: ['#a855f7', '#00f0f9']
            }
        }
    };

</script>

<template>
    <div class="container mx-auto">
        <Breadcrumbs parentTitle="台股恐慌儀表板" subParentTitle="Dashboard" />

        <div class="Dashboard_desc px-4 mb-4">
            <p>▪️以下三個大盤重要指數在特定條件下有一定程度的<b>正相關</b>： 【台灣<b>加權指數TAIEX】</b>、
                 【台灣<b>恐慌指數VIXTWN+<span class="text-primary">垂直翻轉</span>+<span class="text-primary">通膨率</span>】</b>、
                 【台灣<b>加權本益比TAIPE+<span class="text-primary">通膨率</span>】</b>。
                正相關如果脫鉤太嚴重，就是整個大盤的異常警訊。 
            </p>
            <p>▪️合理建倉時機建議是
                <b>恐慌指數<span class="text-primary">VIX小於{{vixSuggest}}</span></b><span class="text-xs">(當前<b>{{ lastVIX }}</b>)</span>、
                <b>本益比<span class="text-primary">PE小於{{peSuggest}}</span></b><span class="text-xs">(當前<b>{{ lastPE }}</b>)</span>。
            </p>
        </div>

        <div class="text-center text-sm flex gap-4 justify-center items-end text-gray-500">
            <span class="text-xs">{{ lastDate }}</span>
            <span class=""><span class="text-xs">EX:</span><b>{{ lastEX }}</b></span>
            <span class=""><span class="text-xs">VIX:</span><b>{{ lastVIX }}</b></span>
            <span class=""><span class="text-xs">PE:</span><b>{{ lastPE }}</b></span>
        </div>

        <!-- 頂端開關 -->
        <div class="basecard_custom py-4 sm:p-4  lg:px-8">

            <div class="topSwitchDiv text-lg font-medium text-gray-600 mb-0">
                <label class="switchWord" :class="{'switchWordAct': vixIsReverse }">
                    <span class="switchText">VIX<br class="inline sm:hidden">垂直翻轉</span>
                    <span class="switch">
                    <input type="checkbox" v-model="vixIsReverse">
                    <span class="slider"></span>
                    </span>
                </label>
                <label class="switchWord" :class="{'switchWordAct': vixIsInflation }">
                    <span class="switchText">VIX<br class="inline sm:hidden">加通膨率</span>
                    <span class="switch">
                    <input type="checkbox" v-model="vixIsInflation">
                    <span class="slider"></span>
                    </span>
                </label>
                <label class="switchWord" :class="{'switchWordAct': peIsInflation }">
                    <span class="switchText">PE<br class="inline sm:hidden">加通膨率</span>
                    <span class="switch">
                    <input type="checkbox" v-model="peIsInflation">
                    <span class="slider"></span>
                    </span>
                </label>
            </div>

            <!-- 畫布區 -->
            <div class="loading-container" v-if="piniaStockMain.stockLoading">正在分析台股數據...
                <div class="spinner"></div>
            </div>
            <apexchart
            v-if="! piniaStockMain.stockLoading"
                    ref="myChart"
                    type="line"
                    height="350"
                    :options="chartOptions"
                    :series="series"
            ></apexchart>

            <!-- 年份切換區 -->
            <div class="setDateRangeDiv">
                <span class="hidden sm:block m-2 text-sm">time range: </span>
                <div @click="setDateRange(1825)" :class="{'daysActivy':showDays==1825 }">5年</div>
                <div @click="setDateRange(1460)" :class="{'daysActivy':showDays==1460 }">4年</div>
                <div @click="setDateRange(1095)" :class="{'daysActivy':showDays==1095 }">3年</div>
                <div @click="setDateRange(730)" :class="{'daysActivy':showDays==730 }">2年</div>
                <div @click="setDateRange(365)" :class="{'daysActivy':showDays==365 }">1年</div>
                <div @click="setDateRange(182)" :class="{'daysActivy':showDays==182 }">6月</div>
                <div @click="setDateRange(91)" :class="{'daysActivy':showDays==91 }">3月</div>
                <!--div @click="setDateRange(60)" :class="{'daysActivy':showDays==60 }">2月</div>
                <div @click="setDateRange(30)" :class="{'daysActivy':showDays==30 }">1月</div-->
            </div>
        </div>

        <div class="grid mt-4 gap-4 grid-cols-1 md:grid-cols-2 xl:grid-cols-4 ">
            <div class="">
                <div class="basecard_custom p-5 h-full">
                    <div>
                        <span class="inline-block w-[20px] h-[20px] align-middle mb-1
                                    bg-gradient-to-tr from-[#4285f4] via-[#9b72cb] to-[#d96570]
                                    [clip-path:polygon(50%_0%,_61%_39%,_100%_50%,_61%_61%,_50%_100%,_39%_61%,_0%_50%,_39%_39%)]">
                        </span>
                        Gemini AI 評分:
                    </div>
                    <apexchart type="radialBar" height="350" :options="chartOptions2" :series="series2"></apexchart>
                    <div class="text-center text-sm">0分為空頭市場，100分為多頭市場</div>
                </div>
            </div>
            <div class="">
                <div class="basecard_custom p-5 h-full flex flex-col justify-between">
                    <div>
                        <span class="inline-block w-[20px] h-[20px] align-middle mb-1
                                    bg-gradient-to-tr from-[#4285f4] via-[#9b72cb] to-[#d96570]
                                    [clip-path:polygon(50%_0%,_61%_39%,_100%_50%,_61%_61%,_50%_100%,_39%_61%,_0%_50%,_39%_39%)]">
                        </span>
                        Gemini AI 估值:
                    </div>
                    <div class="text-sm text-center my-3">台灣加權指數合理價:</div>
                    <div class="
                        text-7xl text-center font-black my-3
                        breathing-text font-bold text-transparent bg-clip-text bg-gradient-to-r from-[#4285f4] via-[#9b72cb] to-[#d96570] 
                        bg-[length:200%_auto] animate-gradient-move
                    ">
                        {{ geminiPrice }}
                    </div>
                    <div class="text-sm text-center mt-8 mb-2">當前收盤價 {{ lastDate }}: </div>
                    <div class="text-3xl text-center font-semibold my-3">{{lastEX}}</div>
                    <div class="text-sm text-center mt-3">AI 估值僅供參考，無任何引導意圖</div>
                </div>
            </div>
            <div class="col-span-1 md:col-span-2">
                <div class="basecard_custom p-5 h-full">
                    <div>
                        <span class="inline-block w-[20px] h-[20px] align-middle mb-1
                                    bg-gradient-to-tr from-[#4285f4] via-[#9b72cb] to-[#d96570]
                                    [clip-path:polygon(50%_0%,_61%_39%,_100%_50%,_61%_61%,_50%_100%,_39%_61%,_0%_50%,_39%_39%)]">
                        </span>
                        Gemini AI 建議:
                        <p class="p-6" v-html="geminiReason.replace(/\n/g, '<br>')"></p>
                        <div class="text-right mr-6 text-sm">資料日期: {{ geminiDate }}</div>
                    </div>
                </div>
            </div>
            <div class="col-span-1 md:col-span-2 xl:col-span-4">
                <div class="basecard_custom p-5 h-full">
                    📈恐慌本益比:
                    <div class="pt-3 pl-3 text-sm">
                        建倉合理值為
                        <b>恐慌指數<span class="text-primary">VIX小於{{vixSuggest}}</span></b>(當前<b>{{ lastVIX }}</b>)、
                        <b>本益比<span class="text-primary">PE小於{{peSuggest}}</span></b>(當前<b>{{ lastPE }}</b>)
                    </div>
                    <div class="flex flex-row">
                        <apexchart class="BarVixPe" type="bar" height="200" :options="chartOptions3" :series="series3"></apexchart>
                        <div class="BarVixPeDesc text-xs">
                            <div class="mt-12">
                                <span class="text-red-400"  v-if="(lastVIX/vixSuggest) > 1">{{ (lastVIX/vixSuggest * 100).toFixed(0) }}% 超標</span>
                                <span class="text-blue-400" v-if="(lastVIX/vixSuggest) <= 1">{{ (lastVIX/vixSuggest * 100).toFixed(0) }}% 達標</span>
                            </div>
                            <div class="mt-10">
                                <span class="text-red-400"  v-if="(lastPE/peSuggest) > 1">{{ (lastPE/peSuggest * 100).toFixed(0) }}% 超標</span>
                                <span class="text-blue-400" v-if="(lastPE/peSuggest) <= 1">{{ (lastPE/peSuggest * 100).toFixed(0) }}% 達標</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="text-center text-gray-500 mt-4 mx-6 text-sm">本站同步證交所資料，加權指數於下午3~4點更新，恐慌指數於下午4~5點更新，所有資料更新後Gemini才給出當日建議。</div>
</template>


<style lang="scss">
    .apexcharts-tooltip {
        background: rgba(255, 255, 255, 0.6) !important;
    }
    .apexcharts-tooltip-title{
        background: rgba(255, 255, 255, 0.6) !important;
    }

</style>

<style lang="scss" scoped>

    // 讀取框
    .loading-container {
    width: 100%;
    height: 350px;
    display: flex;
    align-items: center;      /* 垂直置中 */
    justify-content: center;  /* 水平置中 */
    }

    .spinner {
    width: 40px;
    height: 40px;
    border: 4px solid #e5e5e5;
    border-top: 4px solid #a457e6;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }

    // 開關
    .switchWord {
    display: inline-flex;
    align-items: center;
    gap: 2px;
    cursor: pointer;
    padding: 3px 6px;
    border-radius: 8px;
    transition: background 0.2s;
    }

    .switchWord:hover {
    background: rgba(0,0,0,0.05);
    }

    .switchText {
    font-size: 12px;
    line-height: 14px;
    text-align: center;
    }

    /* switch本體 */
    .switch {
    position: relative;
    width: 44px;
    height: 24px;
    }

    .switch input {
    opacity: 0;
    width: 0;
    height: 0;
    }

    .slider {
    position: absolute;
    inset: 0;
    background: #ccc;
    border-radius: 24px;
    transition: 0.25s;
    }

    .slider:before {
    content: "";
    position: absolute;
    width: 18px;
    height: 18px;
    left: 3px;
    top: 3px;
    background: white;
    border-radius: 50%;
    transition: 0.25s;
    }

    input:checked + .slider {
    background: #a855f7;
    }

    input:checked + .slider:before {
    transform: translateX(20px);
    }

    //我自訂的開關樣式
    .topSwitchDiv{ //大外層
        color: #BBB;
        user-select: none;
        transition: color 0.3s; 
        display: flex;
        justify-content: center;
        flex-wrap: wrap;
    }
    .switchWordAct{
        color: #555;
        font-weight: 600;
    }

    //頂端說明
    .Dashboard_desc{
        p{
            line-height: 22px;
            padding: 2px;
        }
    }

    // 下方 days 選擇用
    .setDateRangeDiv{
        display: flex;
        justify-content: center;
        flex-wrap: wrap;
        color: #AAA;
        padding: 10px 4px;
        
        div{
            background-color: #EEE;
            border-radius: 10px;
            padding: 4px 15px;
            margin: 2px 3px;
            cursor: pointer;
            user-select: none;
            transition: background-color 0.3s,color 0.3s, scale 0.3s; 
        }
        div:hover{
            color: #eee;
            background-color: #d1b8e6;
        }
        div.daysActivy{
            color: #fff;
            font-weight: 600;
            background-color: #a457e6;
            scale: 110%;
        }
    }
    @media (max-width: 500px) {
        .setDateRangeDiv div{
            padding: 6px 10px;
            margin: 1px 1px;
            font-size: 14px;
        }
    }

    // gemini icon
    .gemini-icon {
    display: inline-block;
    width: 20px;
    height: 20px;
    vertical-align: middle;
    margin-bottom: 0.2rem; // 微調對齊文字的高度
    
    // 使用稍微圓潤一點的四角星路徑
    clip-path: polygon(
        50% 0%, 62% 38%, 
        100% 50%, 62% 62%, 
        50% 100%, 38% 62%, 
        0% 50%, 38% 38%
    );

    // 在小尺寸下，對角線漸層最能展現色彩感
    background: linear-gradient(135deg, #4285f4, #9b72cb, #d96570);
    
    // 讓顏色更亮一點，在小圖標上才顯眼
    filter: saturate(1.4) brightness(1.1);
    }



    //自訂卡片
    .basecard_custom {  
        border-radius: 10px;
        box-shadow: 0 4px 20px 1px rgba(0, 0, 0, 0.06), 0 1px 4px rgba(0, 0, 0, 0.08);
    }



    //估值股價動畫文字
    @keyframes flow-and-breathe {
    0% {
        background-position: 0% 50%;
        opacity: 0.8;
    }
    50% {
        background-position: 100% 50%;
        opacity: 1; // 呼吸到最亮
    }
    100% {
        background-position: 0% 50%;
        opacity: 0.8;
    }
    }

    .breathing-text {
    // 1. 漸層文字核心
    font-weight: bold;
    background: linear-gradient(90deg, #4285f4, #9b72cb, #d96570, #4285f4);
    background-size: 300% auto; // 放大背景，讓移動有空間
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;

    // 2. 動畫掛載
    animation: flow-and-breathe 9s ease-in-out infinite;
    }

    @keyframes glow {
    0%, 100% {
        filter: drop-shadow(0 0 2px rgba(155, 114, 203, 0.3));
    }
    50% {
        filter: drop-shadow(0 0 8px rgba(155, 114, 203, 0.8));
    }
    }

    .breathing-text {
    // 疊加剛才的動畫
    animation: flow-and-breathe 9s infinite, glow 9s infinite;
    }


    // 以下長條圖
    .BarVixPe{
        width: calc(100% - 70px);
    }
    .BarVixPeDesc{
        width: 70px;
    }

</style>