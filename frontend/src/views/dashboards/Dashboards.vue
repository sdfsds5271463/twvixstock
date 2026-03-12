<script setup lang="ts">

    import Breadcrumbs from "@/components/Breadcrumbs.vue"
    import { onMounted, ref, watch, nextTick } from "vue"
    import { piniaStock } from '@/store/piniaStock'

    // 使用 pinia
    const piniaStockMain = piniaStock()

    // 定義圖表用的參數
    let DateArr = ref<string[]>([])
    let TAIEX = ref<number[]>([])
    let VIXTWN = ref<number[]>([])
    let VIXTWN_base = ref<number[]>([])  //vix 計算的基準
    let TAIPE = ref<number[]>([])
    let TAIPE_base = ref<number[]>([])  //pe 計算的基準

    // 運算用的
    const myChart = ref(null);  //圖表本體
    let vixIsReverse = ref<boolean>(false);  //VIX 是否顛倒
    let vixIsInflation = ref<boolean>(false);  //VIX 是否通膨
    let peIsInflation = ref<boolean>(false);  //PE 是否通膨


    // =============================== 以下計算 ===============================

    // 設定日期範圍
    const setDateRange = (days:number)=>{
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
    }

    // 顛倒 vix 按鈕
    const vixReverseBtn = ()=>{
        vixIsReverse.value = !vixIsReverse.value; //顛倒
        //vixCompute(); //運算 vix
    }

    // 通膨 vix 按鈕
    const vixInflationBtn = ()=>{
        vixIsInflation.value = !vixIsInflation.value; //通膨
        //vixCompute(); //運算 vix
    }

    // 通膨 pe 按鈕
    const peInflationBtn = () => {
        peIsInflation.value = !peIsInflation.value;  //通膨
        //peCompute(); //運算 pe
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
        VIXTWN.value = [...VIXTWN_base.value]; //還原基準 vix

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
    }

    // pe 運算(依據 通膨旗標)
    const peCompute = ()=>{
        TAIPE.value = [...TAIPE_base.value]; //還原基準 pe

        //通膨旗標
        if(peIsInflation.value){
            TAIPE.value = inflationCompute(TAIPE.value);
        }
    }

    //通膨率運算
    const inflationCompute = (NumArr:number[]):number[] =>{

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
            NumArr[i] = NumArr[i] / (1 + (rate * Lrate)); //調整左側
        }
        ni = (NumArr.length/2);
        for(let i=(NumArr.length-1); i>NumArr.length/2; i--){
            let Rrate = (i-ni)/ni; //右半部，越右越1 越中越0 
            NumArr[i] = NumArr[i] * (1 + (rate * Rrate)); //調整右側
        }

        return NumArr;
    }
    

    // =============================== 以下初始 ===============================

    // pinia通知初始完成時(初次進首頁) --> 初始
    watch(() => piniaStockMain.InitialReady, (newval, oldval)=>{    
        if(newval){ //初始化剛剛完成
            initialChart();
        }
    }, {deep:true});

    // 掛載時(route進來的時候) --> 初始
    onMounted(async (): Promise<void> => {
        if(piniaStockMain.InitialReady){ //初始化早就完成了
            initialChart();
        }
    })

    //初始動作
    const initialChart = ()=>{
        setTimeout(()=>{ //等畫布完全載入
            setDateRange(365);    //給初始日期
        }, 200);
    }


    // =============================== 以下圖表 ===============================

    // 1. 資料定義：series 陣列裡放幾個物件，就有幾條線
    const series = ref([
        {
            name: '台灣加權指數 TAIEX',
            type: 'area',
            //data: [2300, 1100, 2200, 2700, 1300, 2200, 3700, 2100, 4400, 2200, 3000]
            data: TAIEX
        },{
            name: '台灣恐慌指數 VIXTWN',
            type: 'line',
            //data: [84, 95, 81, 97, 82, 73, 81, 71, 96, 87, 43]
            data: VIXTWN
        },{
            name: '台灣加權本益比 TAIEX',
            type: 'line',
            //data: [30, 25, 36, 30, 45, 35, 64, 52, 59, 36, 39]
            data: TAIPE
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
                seriesName: '台灣加權指數 TAIEX', // 與 series 裡的 name 對應
                //title: { text: "TEAM A (Points)" },
                    labels: {
                    formatter: function (val) {
                        return val.toFixed(0) / 1000 + "k"; 
                    },
                    offsetX: -15,
                }
            },{
                seriesName: ['台灣恐慌指數 VIXTWN', '台灣加權本益比 TAIEX'],
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

</script>

<template>
    <div class="container mx-auto">


        <Breadcrumbs parentTitle="Dashboard" subParentTitle="台股恐慌儀表板" />



        <BaseCard noPadding class="md:p-8 sm:p-4">
            <div class="text-lg font-medium text-gray-600 m-4">
                <!--span class="text-lg font-medium">VIX垂直翻轉</span-->



        <label class="switchWord">
            <span class="switchText">VIX垂直翻轉</span>
            <span class="switch">
            <input type="checkbox" v-model="vixIsReverse">
            <span class="slider"></span>
            </span>
        </label>
        <label class="switchWord">
            <span class="switchText">VIX加通膨率</span>
            <span class="switch">
            <input type="checkbox" v-model="vixIsInflation">
            <span class="slider"></span>
            </span>
        </label>
        <label class="switchWord">
            <span class="switchText">PE加通膨率</span>
            <span class="switch">
            <input type="checkbox" v-model="peIsInflation">
            <span class="slider"></span>
            </span>
        </label>
        <br>
        <span class="text-sm">VIX翻轉通膨後與加權正相關</span> <span class="text-sm">PE通膨後與加權正相關</span>




            </div>



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


            <span @click="setDateRange(1825)">
                5年
            </span>

            <span @click="setDateRange(1095)">
                3年
            </span>

            <span @click="setDateRange(730)">
                2年
            </span>

            <span @click="setDateRange(365)">
                1年
            </span>

            <span @click="setDateRange(182)">
                6月
            </span>

            <span @click="setDateRange(91)">
                3月
            </span>

            <span @click="setDateRange(60)">
                2月
            </span>

            <span @click="setDateRange(30)">
                1月
            </span>

        </BaseCard>
    </div>
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
    gap: 10px;
    cursor: pointer;
    padding: 6px 10px;
    border-radius: 8px;
    transition: background 0.2s;
    }

    .switchWord:hover {
    background: rgba(0,0,0,0.05);
    }

    .switchText {
    font-size: 14px;
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
    background: #4f46e5;
    }

    input:checked + .slider:before {
    transform: translateX(20px);
    }

        

</style>