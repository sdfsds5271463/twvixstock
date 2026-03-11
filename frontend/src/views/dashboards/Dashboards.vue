<script setup lang="ts">

import Breadcrumbs from "@/components/Breadcrumbs.vue"
import { onMounted,ref } from "vue"

import { piniaStock } from '@/store/piniaStock'

const piniaStockMain = piniaStock()

onMounted(async (): Promise<void> => {
  piniaStockMain.submitStockDB({ start: "2026-02-01", end: "2026-02-28", type: "TAIEX" });
})

  // 1. 資料定義：series 陣列裡放幾個物件，就有幾條線
  const series = ref([{
          name: 'TEAM A',
          type: 'column',
          data: [2300, 1100, 2200, 2700, 1300, 2200, 3700, 2100, 4400, 2200, 3000]
        }, {
          name: 'TEAM B',
          type: 'area',
          data: [44, 55, 41, 67, 22, 43, 21, 41, 56, 27, 43]
        }, {
          name: 'TEAM C',
          type: 'line',
          data: [30, 25, 36, 30, 45, 35, 64, 52, 59, 36, 39]
        }]);

  // 2. 圖表配置
  const chartOptions = ref({chart: {
          height: 350,
          type: 'line',
          stacked: false,
        },
        stroke: {
          width: [0, 2, 5],
          curve: 'smooth'
        },
        plotOptions: {
          bar: {
            columnWidth: '50%'
          }
        },
        
        fill: {
          opacity: [0.85, 0.25, 1],
          gradient: {
            inverseColors: false,
            shade: 'light',
            type: "vertical",
            opacityFrom: 0.85,
            opacityTo: 0.55,
            stops: [0, 100, 100, 100]
          }
        },
        labels: ['01/01/2003', '02/01/2003', '03/01/2003', '04/01/2003', '05/01/2003', '06/01/2003', '07/01/2003',
          '08/01/2003', '09/01/2003', '10/01/2003', '11/01/2003'
        ],
        markers: {
          size: 0
        },
        xaxis: {
          type: 'datetime'
        },
        /*yaxis: {
          title: {
            text: 'Points',
          }
        },*/
yaxis: [
  {
    seriesName: 'TEAM A', // 與 series 裡的 name 對應
    title: { text: "TEAM A (Points)" },
  },
  {
    seriesName: 'TEAM B',
    opposite: true, // 放在右邊
    title: { text: "TEAM B (Percentage)" },
  },
  {
    seriesName: 'TEAM C',
    opposite: true, // 也放在右邊（或不加此行與 B 並列）
    title: { text: "TEAM C (Ratio)" },
    // 如果不希望座標軸重疊，可以設定 offset
  }
],
        tooltip: {
          shared: true,
          intersect: false,
          y: {
            formatter: function (y) {
              if (typeof y !== "undefined") {
                return y.toFixed(0) + " points";
              }
              return y;
        
            }
          }
        }
  });



</script>

<template>
    <div class="container mx-auto">





        <div class="p-6 mt-2 rounded-xl shadow-lg bg-gray-200 border-b border-gray-300 mb-6">
        <h2 class="text-xl font-bold mb-4 text-gray-600">參與專案 & 語言使用率</h2>
        <apexchart
            type="line"
            height="400"
            :options="chartOptions"
            :series="series"
        ></apexchart>
        </div>


        <Breadcrumbs parentTitle="Dashboard" subParentTitle="台股恐慌儀表板" />

    </div>
</template>
