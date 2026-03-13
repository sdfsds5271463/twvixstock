<template>
    <div class="px-5">
        <!-- p class="text-xl font-semibold mb-2 text-primary">Allen Zheng</p -->
        <div class="text-sm text-gray-400">自我介紹:</div>
        <div class="border-b border-gray-300 pb-6 mb-6 aboutme_content">
          <p>
            資深網路工程師，在業界任職13年（雲科研發處網管4年、博斯特產品開發部9年）。精通領域有 <b>前端、後端、自動部屬、API對接、爬蟲</b>，
            經歷了 <b>PHP→GO、jQuery→Vue、JavaScript→TypeScript、Nginx→CI/CD</b> 的技術變革，
            處理過 <b>大量金流API、併發多線爬蟲、複雜微服務部屬、多線程與多進程運算</b>，以 <b>主管</b> 身分參與過 <b>5個大型專案</b>，
            精通 <b>5種</b> 主流網頁程式語言。
          </p>
          <p>
            就學階段，從甲工資訊科 → 雲科資工系 → 雲科資工所，一路都在領域內發展，於大三開始產學合作，並以網管技術人員任職於雲科研發處。
            畢業後亦投入網頁領域的職場中，在任職的第二年底成為部門主管，經歷9年的變遷後，於2025年底離開該公司。
          </p>
        </div>

        <div class="text-sm text-gray-400">大型專案:</div>
        <div class="bigProjects grid grid-cols-12 gap-3 border-b border-gray-300 pb-10 mb-4">
          <div class="basecard_custom col-span-12 md:col-span-6 xl:col-span-4">
            <p class="font-medium">▪️2026 台灣恐慌資料站 
              <a class="btn text-blue-500"
                href="https://twvixstock.qzz.io/" target="_blank">
                🌐點此訪問
              </a>
            </p>
            <p class="text-sm text-gray-500">排程爬蟲股市並 AI 分析的<b>自動部屬</b>全端系統</p>
          </div>
          <div class="basecard_custom col-span-12 md:col-span-6 xl:col-span-4">
            <p class="font-medium">▪️2024 大型遊戲平台
              <a class="btn text-blue-500"
                @click = "currentImageName = profile_photos[keyOfPhotos['2024_jyb3.jpg']]['name']" >
                🎞️點此預覽
              </a>
            </p>
            <p class="text-sm text-gray-500">對接超過<b>30個</b>遊戲<b>9個金流</b>的完整遊戲系統</p>
          </div>
          <div class="basecard_custom col-span-12 md:col-span-6 xl:col-span-4">
            <p class="font-medium">▪️2022 體育遊戲系統
              <a class="btn text-blue-500"
                @click = "currentImageName = profile_photos[keyOfPhotos['2022_sb2.jpg']]['name']" >
                🎞️點此預覽
              </a>
            </p>
            <p class="text-sm text-gray-500"><b>自動部屬</b>超過<b>20個</b>微服務<b>200人</b>同時在線的<b>高併發</b>系統</p>
          </div>
          <div class="basecard_custom col-span-12 md:col-span-6 xl:col-span-4">
            <p class="font-medium">▪️2020 電子遊戲
              <a class="btn text-blue-500"
                href="http://lottsample.ddns.net:9527/egames/public/trygame" target="_blank">
                🎮點此試玩
              </a>
            </p>
            <p class="text-sm text-gray-500">每日注單遠超<b>上萬張</b>的動畫遊戲</p>
          </div>
          <div class="basecard_custom col-span-12 md:col-span-6 xl:col-span-4">
            <p class="font-medium">▪️2019 區塊鏈百家樂
              <a class="btn text-blue-500"
                href="http://livesample.ddns.net:9527/game/trygame" target="_blank">
                🎮點此試玩
              </a>
            </p>
            <p class="text-sm text-gray-500">開設超過<b>120子代理</b>的hash加密紙牌遊戲</p>
          </div>
          <div class="basecard_custom col-span-12 md:col-span-6 xl:col-span-4">
            <p class="font-medium">▪️2018 彩票遊戲
              <a class="btn text-blue-500"
                href="http://lottsample.ddns.net:9527/index.php?m=common&c=trygame&a=trygame&type=0" target="_blank">
                🎮點此試玩
              </a>
            </p>
            <p class="text-sm text-gray-500">註冊超過<b>27萬人</b>的概率精算遊戲</p>
          </div>
        </div>

        <div class="grid grid-cols-12 gap-3">
            <div class="col-span-12 md:col-span-6 lg:col-span-3 break-all">
                <p class="text-purple-500 mb-1"><i class="i-Professor text-base"></i><span class="ml-1">年資</span></p><span>13 年</span>
            </div>
            <div class="col-span-12 md:col-span-6 lg:col-span-3 break-all">
                <p class="text-purple-500 mb-1"><i class="i-Edit-Map text-base"></i><span class="ml-1">學歷</span></p><span>雲科資訊工程研究所</span>
            </div>
            <div class="col-span-12 md:col-span-6 lg:col-span-3 break-all">
                <p class="text-purple-500 mb-1"><i class="i-ID-Card text-base"></i><span class="ml-1">證照</span></p><span>電腦軟體應用丙級<br>電腦硬體裝修乙級</span>
            </div>
            <div class="col-span-12 md:col-span-6 lg:col-span-3 break-all">
                <p class="text-purple-500 mb-1"><i class="i-Calendar text-base"></i><span class="ml-1">生日</span></p><span>23, Jan 1991</span>
            </div>
            <div class="col-span-12 md:col-span-6 lg:col-span-3 break-all">
                <p class="text-purple-500 mb-1"><i class="i-Professor text-base"></i><span class="ml-1">個人網站</span></p>
                <a class="text-blue-500" href="https://twvixstock.qzz.io" target="_blank">https://twvixstock.qzz.io</a>
            </div>
            <div class="col-span-12 md:col-span-6 lg:col-span-3 break-all">
                <p class="text-purple-500 mb-1"><i class="i-Edit-Map text-base"></i><span class="ml-1">Github</span></p>
                <a class="text-blue-500" href="https://github.com/sdfsds5271463/twvixstock" target="_blank">https://github.com/sdfsds5271463/twvixstock</a>
            </div>
            <div class="col-span-12 md:col-span-6 lg:col-span-3 break-all">
                <p class="text-purple-500 mb-1"><i class="i-ID-Card text-base"></i><span class="ml-1">Email</span></p><span>sdfsds5271463@gmail.com</span>
            </div>
        </div>
    </div>

    <div class="p-6 mt-2 rounded-xl shadow-lg bg-gray-200 border-b border-gray-300 mb-6">
      <h2 class="text-xl font-bold mb-4 text-gray-600">參與專案 & 語言使用率</h2>
      <apexchart
        type="line"
        height="400"
        :options="chartOptions"
        :series="series"
      ></apexchart>
    </div>

    <div class="p-5">
        <div class="text-sm text-gray-400">精通技能關鍵字:</div>
        <div class="border-b border-gray-300 pb-10 mb-8 text-blue-500 aboutme_content font-semibold">
          <p>前端 # Vue3 Vite Pinia Vuex VueRouter TypeScript TailwindCSS BootStrap jQuery Javascript Twig Sass Css Html PixiJS</p>
          <p>後端 # Golang Gin PHP Laravel Python Scrapy RESTfulAPI WebSocket Redis MySQL SQL Gemini API </p>
          <p>部屬 # DevOps Kubernetes K8s CI/CD Docker GCP GKE CloudBuild Prometheus ELK GitHub Cloudflare Linux Nginx</p>
          <p>其他 # UnitTest WebCrawler Git SEO MVC RWD Ajax Json Selenium Playwright Photoshop WordPress Spine JWT Apexchart</p>
        </div>
    </div>





</template>



<script setup lang="ts">

    import { ref, inject, Ref } from 'vue'
    import type { ImageItem } from '@/types/ImageItem'
    const currentImageName = inject<Ref<string>>('currentImageName')!
    const profile_photos = inject<Ref<ImageItem[]>>('profile_photos')!

    //對應 currentImageName 的 ID
    let keyOfPhotos:{[key: string]: number} = {}
    profile_photos.value.forEach((photo, index) => {
        const photoName = photo.name.split("/")[2]
        keyOfPhotos[photoName] = index
    })


  // 1. 資料定義：series 陣列裡放幾個物件，就有幾條線
  const series = ref([
    {
      name: ["PHP","Laravel"],
      data: [55, 60, 95, 50, 30, 25, 40,  10,  0, 50, 50,  0],
    },
    {
      name: ["GOLANG","Gin"],
      data: [ 0,  0,  0,  0,  0,  0,  0,  60, 65, 15, 5, 35],
    },
    {
      name: ["PYTHON","Scrapy"],
      data: [ 0,  0,  0, 20, 10,  0, 10,  0,  0,  35, 30, 10],
    },
    {
      name: ["JAVASCRIPT","jQuery"],
      data: [45, 40,  5, 30, 60, 75, 50,  20,  5,  0,  0, 5],
    },
    {
      name: ["TYPESCRIPT","Vue"],
      data: [ 0,  0,  0,  0,  0,  0,  0,  10, 30,  0, 15, 50],
    },
  ]);

  // 2. 圖表配置
  const chartOptions = ref({
    chart: {
      zoom: { enabled: false }, // 是否開啟縮放功能
      toolbar: { show: false }
    },
    //colors: ['#2563eb', '#dc2626', '#16a34a'], // 藍、紅、綠 三色
    stroke: {
      curve: 'smooth', // 曲線平滑化
      width: 3
    },
    xaxis: {
      categories: [['2015-', '任職學界'], "2016", ['2017', 'API專案'], ['2018', '彩票平台'], ['2019', '百家樂'], ['2020', '電子遊戲'], 
        '2021', ['2022', '體育系統'], '2023', ['2024', '遊戲平台'], '2025', ['2026+', '股票系統']
      ],
      title: { text: '年代 專案' },
      labels: {
        style: {
          fontSize: '11px',
        }
      }
    },
    yaxis: {
      title: { text: '使用率' }
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


<style lang="scss">
  .aboutme_content p{ //關於我內容格式化
    line-height: 22px;
    text-indent: 2em;
    padding: 2px;
    b{
      font-size: 105%;
    }
    a{
      font-size: 105%;
      font-weight: 600;
      color: rgb(59, 130, 246);
    }
  }

</style>

<style lang="scss" scoped>


  //大型遊戲專案
  .bigProjects{
    
    //瀏覽試玩等按鈕
    .btn{
        cursor: pointer;
        user-select: none;
        transition: color 0.3s, background-color 0.3s; 
        background-color: #EEE;
        padding: 2px 4px;
        border-radius: 5px;
    }
    .btn:hover{
        color: #EEE;
        background-color: #a457e6;
    }

    .basecard_custom {  
        border-radius: 4px;
        box-shadow: 0 4px 20px 1px rgba(0, 0, 0, 0.06), 0 1px 4px rgba(0, 0, 0, 0.08);
        padding: 8px 12px;
    }

  }

</style>