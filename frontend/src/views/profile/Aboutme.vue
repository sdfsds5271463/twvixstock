<template>
    <div class="p-5">
        <p class="text-xl font-semibold mb-2">Personal Information</p>
        <p class="border-b border-gray-300 pb-10 mb-8">Lorem ipsum dolor sit amet consectetur adipisicing elit. Eveniet, commodi quam! Provident quis voluptate asperiores ullam, quidem odio pariatur. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptatem, nulla eos? Cum non ex voluptate corporis id asperiores doloribus dignissimos blanditiis iusto qui repellendus deleniti aliquam, vel quae eligendi explicabo.</p>
        <div class="grid grid-cols-12 gap-5 border-b border-gray-300 pb-10 mb-10">
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Calendar text-base"></i><span class="ml-1">Birth Date</span></p><span>1, Jan 1994</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-MaleFemale text-base"></i><span class="ml-1">Gender</span></p><span>1, Jan 1994</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Face-Style-4 text-base"></i><span class="ml-1">Profession</span></p><span>Digital Marketer</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Calendar text-base"></i><span class="ml-1">Birth Date</span></p><span>1, Jan 1994</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-MaleFemale text-base"></i><span class="ml-1">Gender</span></p><span>1, Jan 1994</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Face-Style-4 text-base"></i><span class="ml-1">Profession</span></p><span>Digital Marketer</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Globe text-base"></i><span class="ml-1"> Lives In </span></p><span>8 Years</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Edit-Map text-base"></i><span class="ml-1">Email</span></p><span>example@ui-lib.com</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Professor text-base"></i><span class="ml-1"> Experience</span></p><span>8 Years</span>
            </div>
        </div>
        <div class="my-5">
            <p class="text-2xl">Timothy Carlson</p>
            <p class="text-gray-600">Digital Marketer</p>
        </div>
        <div class="flex justify-around flex-wrap">
            <div class="text-center mx-4 my-2"><i class="i-Plane text-4xl text-purple-500"></i>
                <p class="text-base">Travelling </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Camera text-4xl text-purple-500"></i>
                <p class="text-base">Photography </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Car-3 text-4xl text-purple-500"></i>
                <p class="text-base">Driving </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Gamepad-2 text-4xl text-purple-500"></i>
                <p class="text-base">Gaming </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Music-Note-2 text-4xl text-purple-500"></i>
                <p class="text-base">Music </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Shopping-Bag text-4xl text-purple-500"></i>
                <p class="text-base">Shopping </p>
            </div>
        </div>
    </div>


  <div class="p-6 bg-white rounded-xl shadow-lg">
    <h2 class="text-xl font-bold mb-4 text-gray-800">股票走勢比較</h2>
    <apexchart
      type="line"
      height="400"
      :options="chartOptions"
      :series="series"
    ></apexchart>
  </div>






</template>



<script setup lang="ts">
import { ref } from 'vue';

// 1. 資料定義：series 陣列裡放幾個物件，就有幾條線
const series = ref([
  {
    name: "TAIEX",
    data: [23.22, 25.84, 27.46, 26.10, 28.50] // 你的 Close 數據
  },
  {
    name: "TSMC",
    data: [15.10, 18.20, 17.50, 19.80, 21.00]
  },
  {
    name: "AAPL",
    data: [10.50, 12.30, 11.80, 13.20, 14.50]
  }
]);

// 2. 圖表配置
const chartOptions = ref({
  chart: {
    zoom: { enabled: true }, // 開啟縮放功能，看股票必備
    toolbar: { show: true }
  },
  colors: ['#2563eb', '#dc2626', '#16a34a'], // 藍、紅、綠 三色
  stroke: {
    curve: 'smooth', // 曲線平滑化
    width: 3
  },
  xaxis: {
    categories: ['2026-01', '2026-02', '2026-03', '2026-04', '2026-05'], // 日期
    title: { text: '日期' }
  },
  yaxis: {
    title: { text: '收盤價' }
  },
  legend: {
    position: 'top' // 圖例放在上方
  },
  tooltip: {
    shared: true, // 滑鼠移上去時同時顯示三條線的數值
    intersect: false
  }
});

</script>