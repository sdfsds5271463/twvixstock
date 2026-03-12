<template>

    <ul v-for="i in (Math.ceil(TD.length/2))"   class="timeline clearfix">

        <li class="timeline-line"></li>
        <li v-if="TD[2*i-2].float_right" class="timeline-line"></li> <!-- 兩個元素會給奇偶數判斷 造成下個元素 float-right -->

        <li class="timeline-item">
            
            <!--li class="timeline-group text-center">
                <BaseBtn class="bg-primary text-white">
                    <i class="i-Calendar-4"></i> 2021
                </BaseBtn>
            </li-->

            <div class="timeline-badge">
                <div :style="'background-image: \
                    linear-gradient(to bottom, rgba(139, 92, 246, 0), rgba(139, 92, 246, 0.6)),\
                    url('+ profile_photos[keyOfPhotos[TD[2*i-2].ball_photo]]['name'].replace(/photos\//g, 'photos/small_') +');'"
                class="
                    ballimage
                    rounded-full flex items-center justify-center h-12 w-12 bg-cover
                    text-center text-purple-100 text-sm
                ">
                    <!--i class="i-Cloud-Picture"></i-->
                    <span v-html="TD[2*i-2].ball_text"></span>
                </div>
            </div>
            <BaseCard class="timeline-card">
                <div class="mb-1">
                    <strong class="mr-1" v-html="TD[2*i-2].title"></strong>
                    <span class="text-sm" v-html="TD[2*i-2].sub_title"></span>
                    <p class="text-muted text-blue-500 key_word">
                        <span class="key_word_big" v-html="'# '+TD[2*i-2].key_word_big"></span>
                        <span class="key_word_small" v-html="' '+TD[2*i-2].key_word_small"></span>
                    </p>
                </div>
                <span class="data_desc" v-html="TD[2*i-2].desc1"></span>
                <div class="flex flex-wrap rounded shadow-lg border-2 border-solid mb-2 mt-1">
                    <div class="w-1/2" v-for="i2 in TD[2*i-2].photos" >
                        <div class="absolute img_desc text-white pl-2 pt-1 text-sm" 
                            v-html="profile_photos[keyOfPhotos[i2]]['desc']"></div>
                        <img 
                            class="timeline_img rounded cursor-pointer hover:opacity-50" 
                            :src="profile_photos[keyOfPhotos[i2]]['name'].replace(/photos\//g, 'photos/small_')" 
                            :alt="profile_photos[keyOfPhotos[i2]]['desc']"
                            @click = "currentImageName = profile_photos[keyOfPhotos[i2]]['name']"
                        >
                    </div>
                </div>
                <span class="data_desc" v-html="TD[2*i-2].desc2"></span>
            </BaseCard>
        </li>

        <li v-if="TD[2*i-1]" :class="'timeline-item'">

            <div v-if="i==1" class="timeline-first-space m-20"></div> <!-- 右側首行空行 -->

            <div class="timeline-badge">
                <div :style="'background-image: \
                    linear-gradient(to bottom, rgba(139, 92, 246, 0), rgba(139, 92, 246, 0.6)),\
                    url('+ profile_photos[keyOfPhotos[TD[2*i-1].ball_photo]]['name'].replace(/photos\//g, 'photos/small_') +');'"
                class="
                    ballimage
                    rounded-full flex items-center justify-center h-12 w-12 bg-cover
                    text-center text-purple-100 text-sm
                ">
                    <!--i class="i-Cloud-Picture"></i-->
                    <span v-html="TD[2*i-1].ball_text"></span>
                </div>
            </div>
            <BaseCard class="timeline-card">
                <div class="mb-1">
                    <strong class="mr-1" v-html="TD[2*i-1].title"></strong>
                    <span class="text-sm" v-html="TD[2*i-1].sub_title"></span>
                    <p class="text-muted text-blue-500 key_word">
                        <span class="key_word_big" v-html="'# '+TD[2*i-1].key_word_big"></span>
                        <span class="key_word_small" v-html="' '+TD[2*i-1].key_word_small"></span>
                    </p>
                </div>
                <span class="data_desc" v-html="TD[2*i-1].desc1"></span>
                <div class="flex flex-wrap rounded shadow-lg border-2 border-solid mb-2 mt-1">
                    <div class="w-1/2" v-for="i2 in TD[2*i-1].photos" >
                        <div class="absolute img_desc text-white pl-2 pt-1 text-sm" 
                            v-html="profile_photos[keyOfPhotos[i2]]['desc']"></div>
                        <img 
                            class="timeline_img rounded cursor-pointer hover:opacity-50" 
                            :src="profile_photos[keyOfPhotos[i2]]['name'].replace(/photos\//g, 'photos/small_')" 
                            :alt="profile_photos[keyOfPhotos[i2]]['desc']"
                            @click = "currentImageName = profile_photos[keyOfPhotos[i2]]['name']"
                        >
                    </div>
                </div>
                <span class="data_desc" v-html="TD[2*i-1].desc2"></span>
            </BaseCard>
        </li>
    </ul>


    <ul class="timeline clearfix" style="pointer-events: none">
        <li class="timeline-line"></li>
        <li class="timeline-group text-center">
            <BaseBtn class="bg-yellow text-gray-600">
                <i class="i-Calendar-4"></i> Joined in 2012
            </BaseBtn>
        </li>
    </ul>
</template>



<script setup lang="ts">
    import { inject, Ref, ref, computed } from 'vue'
    import type { ImageItem } from '@/types/ImageItem'
    const currentImageName = inject<Ref<string>>('currentImageName')!
    const profile_photos = inject<Ref<ImageItem[]>>('profile_photos')!

    //對應 currentImageName 的 ID
    let keyOfPhotos:{[key: string]: number} = {}
    profile_photos.value.forEach((photo, index) => {
        const photoName = photo.name.split("/")[2]
        keyOfPhotos[photoName] = index
    })

    const ttt =  ref("small_2022_sb2.jpg");

    type TimelineStruct = {
        ball_photo: string;
        ball_text: string;
        title: string;
        sub_title: string;
        key_word_big: string;
        key_word_small: string;
        desc1: string;
        desc2: string;
        photos: string[];
        float_right?: boolean;
    }

    let TD:TimelineStruct[] = [
        {
            ball_photo: "2018_lo3.jpg",
            ball_text: "2026",
            title: "台股恐慌資訊站",
            sub_title: "自動爬蟲股市並分析的自動部屬系統",
            key_word_big: "Golang Vue3Vite TypeScript TailwindCSS DevOps GCP Kubernetes K8s Docker CI/CD",
            key_word_small: "RESTful API Gin Vuex VueRouter Sass CloudBuild GKE Prometheus GitHub WebCrawler RWD Linux UnitTest SEO Cloudflare",
            desc1: `
                <a href='https://github.com/sdfsds5271463/twvixstock' target='_blank'>開啟本站 GitHub</a>
                獨立開發的個人網站，<b>自動部屬</b>在 GCP 上並開源於 GitHub。系統排程至證交所進行恐慌指數爬蟲，資料送至前端可查詢分析，並開放 API 接口服務。
                本系統有三個 GCP 服務，Mysql、Server、CronJob。其中 Vue 編譯後靜態檔存放在 Server 的 Container；CronJob 爬蟲到的資料交給 Mysql，最後由 Server 實現前端查詢分析與 API 等服務。
            `,
            desc2: `
                本站是一個完整專案架構，具有前端、後端、排程、部屬，其 GCP 自動部屬流程為 Git → GitHub → CloudBuild → Registry → GKE → gke-deploy → (運行服務開放端口) → Cloudflare。
                技術實現上，使用 Go + Gin 框架實現後端響應及排程，並使用 Vue3Vite + Vuex + VueRouter + TailwindCSS 撰寫前端分析頁面。
            `,
            photos: ["2018_lo3.jpg","2018_lo3.jpg","2018_lo3.jpg","2018_lo3.jpg"],
            float_right: false, //奇數物件根據祖先長度，決定是否飄右邊
        },
        {
            ball_photo: "2024_jyb1.jpg",
            ball_text: "2025<br>~<br>2024",
            title: "大型線上遊戲平台",
            sub_title: "對接超過<b>30個</b>遊戲商的完整遊戲系統",
            key_word_big: "RESTful API PHP Laravel Python Vue3 TypeScript MySQL Redis Bootstrap",
            key_word_small: "MVC Vuex Sass WebCrawler Scrapy Selenium Playwright UnitTest SEO",
            desc1: `
                我領導的<b>5人</b>後端技術團隊，負責大量的接口對接及爬蟲支援。
                本平台對接了<b>9個</b>金流商由我一人完全負責，包含提交及回調查詢等系列動作，每月有效金流單超過<b>1萬張</b>，金額達<b>八位數</b>。
                另外我與我的團隊總共對接了超過<b>30個</b>遊戲商，並撰寫了大量的爬蟲，實現併發登入遊戲商後台對單與開設代理等功能。
            `,
            desc2: `
                本站使用 PHP + Laravel 框架為主體，前端使用 PHP View + Vue3 CDN + Bootstrap 進行渲染，所有 API 對接都以 PHP 進行撰寫。查帳、回調、索取遊戲狀態等定期操作，我們使用 Linux 的 Crontab 觸發 Laravel Console 進行排程。
                管理上，需要併發同時登入30間以上遊戲商後台實現一鍵查帳、一鍵開設代理、一鍵關閉遊戲等功能，因此我們使用了大量的 Python + Scrapy 框架進行多線程併發爬蟲的撰寫，並使用 Playwright 做為登入與紀錄登入狀態的主要工具。
            `,
            photos: ["2024_jyb5.jpg","2024_jyb4.jpg","2024_jyb3.jpg","2024_jyb1.jpg"],
        },
        {
            ball_photo: "2022_sb2.jpg",
            ball_text: "2024<br>~<br>2022",
            title: "體育遊戲系統",
            sub_title: "自動擴展pod讓超過<b>200人</b>同時在線的微服務架構",
            key_word_big: "Golang gRPC Vue3 TypeScript DevOps GCP Kubernetes K8s Docker CI/CD WebSocket API Redis",
            key_word_small: "Gin Vuex Sass CloudBuild GKE Prometheus ELK GitHub Linux UnitTest JWT",
            desc1: `
                我領導的<b>8人</b>全端技術團隊，負責將舊版 PHP + jQuery 體育系統源碼重新撰寫成 Golang + Vue3 並<b>自動部屬</b>至 GCP，亦包含後續開發維護等全端工作。
                系統架構包含平台、前端、後端、控端、賽程、訊源、視訊等，由<b>20個</b>以上的 GCP 容器微服務部屬整合運行，具備併發及自動拓展等特性，讓系統可以在超過<b>200人</b>同時在線且每秒更新超過<b>百次</b>賠率的狀況下，穩定運行。
            `,
            desc2: `
                本案在我們團隊接手之後，共分為三個時期：
                初期我們首要任務是要將一個未完成的大量微服務系統依序部屬到全新的 GCP 上，設定其關聯，並將 GitHub 資料與 CloudBuild 觸發 gke-deploy 實現自動部屬功能。
                中期我們開始將 PHP + Laravel 的專案改寫成 Go + Gin，需要大量運算的服務我們使用 Go Channel 集中狀態，併發 gRPC 給自拓展的 pod 進行運算服務。
                後期我們前端版面重構，撰寫新版本的前端連線 WebSocket，並使用，並使用 Vue3 + TypeScript + Sass 重新刻劃 UI 介面。
                營運方面，我們提供標準 API (JWT) 的平台對接，提供 WebSocket API 的賠率對接，也提供訊源 API 等服務。並使用 Prometheus + ELK 監控負載狀態，必要時手動調整 GCP 負載上限。
            `,
            photos: ["2022_sb5.jpg","2022_sb4.jpg","2022_sb2.jpg","2022_sb1.jpg"],
            float_right: false, //奇數物件根據祖先長度，決定是否飄右邊
        },
        {
            ball_photo: "2020_eg4.jpg",
            ball_text: "2022<br>~<br>2018",
            title: "獨立開發多款遊戲",
            sub_title: "電子、百家樂、彩票系統總註冊人數超過<b>27萬人</b>",
            key_word_big: "RESTful API PHP Laravel MVC jQuery WebCrawler API MySQL Bootstrap UnitTest Nginx Git",
            key_word_small: "Twig Javascript Ajax RWD Redis Photoshop Spine PixiJS",
            desc1: `
                試玩: <a href='http://lottsample.ddns.net:9527/index.php?m=common&c=trygame&a=trygame&type=0' target='_blank'>彩票遊戲</a><a href='http://livesample.ddns.net:9527/game/trygame' target='_blank'>區塊鏈百家樂</a><a href='http://lottsample.ddns.net:9527/egames/public/trygame' target='_blank'>電子遊戲</a>
                我領導的<b>5~6人</b>全端技術團隊，從零開始原創三款完全不同類型的遊戲，每款遊戲都有完整的前台、後台、排程、API 等架構，三款遊戲總計超過<b>400個</b>對接商戶。
                2018 彩票遊戲，主體為基於彩票規則的爬蟲獎號概率精算遊戲，累計註冊且遊玩的會員超過<b>27萬人</b>。
                2019 區塊鏈百家樂，主體為 SHA512 預加密可驗證的公平紙牌機率遊戲，與著名遊戲平台合作，內崁為旗下遊戲之一，該平台就開設超過<b>120個</b>子代理。
                2020 電子遊戲，主體為精美的動畫與機率計算，內部崁入合作商捕魚遊戲，整個平台單日注單可遠超<b>上萬張</b>。
            `,
            desc2: `
                三個專案都是基於 PHP + Laravel + jQuery + Bootstrap 所開發的。
                彩票遊戲使用了 Twig 的 View，並寫了大量的爬蟲訪問了彩票官網、獎源、以及我們合作遊戲商的前後台，使用排程進行精算與開獎。
                區塊鏈百家樂使用 Canvas 畫布繪製路圖，排程預先模擬所有場次結果並 SHA512 加密之提供後續驗證，與合作平台使用無縫分數模式對接。
                電子遊戲使用 Spine 編輯成動畫，至前端使用 PixiJS 動畫引擎呈現，並於遊戲內呼叫捕魚遊戲商 API 提供遊玩。
            `,
            photos: ["2020_eg4.jpg","2020_eg3.jpg","2019_li4.jpg","2019_li3.jpg","2018_lo4.jpg","2018_lo3.jpg"],
        },
        {
            ball_photo: "2017_bos1.jpg",
            ball_text: "2018<br>~<br>2017",
            title: "遊戲後台API",
            sub_title: "對接超過<b>30個</b>金流API及<b>10個</b>遊戲API",
            key_word_big: "RESTful API Json PHP Laravel MVC MySQL Redis Git",
            key_word_small: "",
            desc1: `
                項目成立初期加入該團隊，專精負責 API 對接的業務，本專案我總計對接了<b>30個</b>以上金流 API、<b>10個</b>以上遊戲 API。
            `,
            desc2: `
                核心技術 PHP、Laravel、API、RESTful API、Json、Git。
                主要工作流程為接洽合作商、開設代理加白並索取對接文件，經過測試環境正式環境到對接完成上線。
                遊戲 API 有創建會員、無縫或轉帳、啟動遊戲、報表等固定接口；金流 API 有代付代收、請求、回調、查詢等固定接口。我的任務就是完善這個流程。
            `,
            photos: ["2017_bos4.jpg","2017_bos3.jpg","2017_bos2.jpg","2017_bos1.jpg"],
            float_right: false, //奇數物件根據祖先長度，決定是否飄右邊
        },
        {
            ball_photo: "2012_yuntech2.jpg",
            ball_text: "2016<br>~<br>2012",
            title: "雲科研發處網管",
            sub_title: "CMS內容管理系統及E化政府作業流程",
            key_word_big: "WordPress PHP MySQL Javascript jQuery Photoshop SEO",
            key_word_small: "Html Css Apache",
            desc1: `
                自2012年起職位為雲科研發處網管，任職期間總共撰寫了超過<b>20個</b>上線網站。
                網站分為靜態網站、CMS 內容管理系統、E化政府單位流程的介面化資料庫網頁系統。
            `,
            desc2: `
                核心技術為 WordPress、PHP、MySQL、jQuery、Html、Css、SEO、Apache。
                撰寫了大量校內活動靜態網頁，學會了極強的 Photoshop 素材編輯能力以及基礎的 SEO 推廣，自我 2012 任職起連續四年的雲科官方校慶系列網站全部都是我做的。
                協助將各處室的靜態網站以 Apache + WordPress 建成教職員自行編輯內容、上傳圖片的 CMS 內容管理系統。
                協助政府單位建置了數個 PHP + MySQL 的E化投標或公文等系統，項目內容包含介面化填寫、比對資料，寄發郵件等。
            `,
            photos: ["2016_yuntech3.jpg","2016_yuntech2.jpg","2012_yuntech2.jpg","2012_yuntech1.jpg"],
        },
    ];
    
    //const evenTimeline = computed(()=> TD.filter((v,i) => i % 2 === 0));

    //desc1 desc2 讓 \n 變成 <p><\p>
    TD.forEach((v,k)=>{
        //整理內容 (變成 html顯示好看的)
        for(let i=0; i<=1; i++){
            //desc1 desc2 都歸整
            let desc = v.desc1;
            if(i==1){ desc = v.desc2; }

            desc = desc.trim();
            let exp1 = desc.split("\n");
            for(let k2 in exp1){
                exp1[k2] = exp1[k2].trim();  //每行去掉空白
            }
            desc = "<p>" + exp1.join("<\p><p>") + "<\p>"; //段落包好

            //desc1 desc2 賦予值
            if(i==1){ 
                TD[k].desc2 = desc
            }else{
                TD[k].desc1 = desc
            }
        }
    });

</script>





<style lang="scss" scoped>
.timeline {
  position: relative;
  list-style: none;
  padding: 0;
  margin: 0;
  
  li.timeline-item {
      position: relative;
      width: 50%;
      display: inline-block;
      
      &:nth-child(even) {
          float: left;
          //margin-top: 6rem;
          padding: 0 3rem 3rem 0;
          .timeline-badge {
              left: calc(100% - 24px);
          }
      }
      &:nth-child(odd) {
          float: right;
          //margin-top: 6rem;
          padding: 0 0 3rem 3rem;
          .timeline-badge {
              right: calc(100% - 24px);
          }
      }
      .timeline-badge {
          position: absolute;
          width: 48px;
          height: 48px;
          z-index: 1;
          .ballimage{
            text-shadow: 1px 1px 2px black;
            line-height: 10px;
          }
      }
      .badge-icon {
          display: inline-block;
          text-align: center;
          font-size: 22px;
          border-radius: 50%;
          height: 100%;
          width: 100%;
          line-height: 48px;
      }
      .badge-img {
          display: inline-block;
          border-radius: 50%;
          height: 100%;
          width: 100%;
      }
  }
  li.timeline-group {
      position: relative;
      z-index: 99;
      padding: 0 0 2rem 0;
      pointer-events: none;
  }
  .timeline-line {
      position: absolute;
      content: "";
      width: 1px;
      height: 100%;
      background: #D1D5DB;
      left: 0;
      right: 0;
      margin: auto;
      z-index: 0;
  }
  .img_desc{ //圖片敘述文字
    pointer-events: none;
    text-shadow: 1px 1px 5px black;
    z-index: 50;
  }
  .timeline_img{ //圖片滑動效果
    transition: opacity 0.2s ease;
  }
  .key_word{ //關鍵字
    line-height: 14px;
  }
  .key_word_big{ //關鍵字 大
    font-size: 14px;
    font-weight: 600;
  }
  .key_word_small{ //關鍵字 小
    font-size: 12px;
    font-weight: 400;
  }
}

@media (max-width: 767px) {
  .timeline {
      .timeline-line, .timeline-first-space { //.timeline-first-space 是首行空行
        display: none !important;
      }
      li.timeline-item {
          width: 100%;
          padding: 3rem 0 3rem !important;
          &:nth-child(odd) {
              margin-top: 1rem;
          }
          .timeline-badge {
              left: 0 !important;
              right: 0 !important;
              top: -16px;
              margin: auto;
          }
      }
      
      li.timeline-group {
          padding: 0 0 3rem;
      }
  }
}
</style>

<style lang="scss">
    //這裡不能 scoped
    .data_desc p{ //時間線desc格式化
        line-height: 22px;
        text-indent: 2em;
        padding: 2px;
        a{
            color: white;
            background-color: rgba(139, 92, 246,0.9);
            border-radius: 5px;
            padding: 3px 12px;
            margin: 3px;
            box-shadow: 1px 1px 3px rgb(112, 112, 112);
            transition: opacity 0.3s ease;
            font-weight: 300;
            font-size: 14px;
            line-height: 28px;
            white-space: nowrap;
        }
        a:hover{
            opacity: 60%;
        }
    }
</style>

