<template>
    <div class="grid grid-cols-12 gap-4">
        <template v-for="i in projects_data">
            <div class="col-span-12 xl:col-span-3 lg:col-span-4 md:col-span-6">
                <div class="basecard_custom p-5 text-center h-full grid">
                    <img class="w-20 h-20 object-cover object-center m-auto rounded-full shadow-lg mb-2" 
                        :src="profile_photos[keyOfPhotos[i.logoImg]]['name'].replace(/photos\//g, 'photos/small_')" 
                        :alt="profile_photos[keyOfPhotos[i.logoImg]]['desc']"
                    />
                    <h5 class="text-lg text-primary">{{i.title}}</h5>
                    <p class="text-sm">{{i.desc}}</p>
                    <p class="keyword text-xs font-semibold text-blue-500 pt-1"># {{i.keyword}}</p>
                    <p class="project_content my-3 text-left indent-8" v-html="i.content"></p>
                    <div class="self-end">
                        <!-- 試玩按鈕區塊 -->
                        <BaseBtn v-if="i.links[0]!=''" rounded class="bg-purple-500 text-white hover:animate-pulse">
                            <i class="i-Gamepad-2 mr-2"></i>
                            <a :href="i.links[0]" target="_blank">
                                {{ i.links[1] }}
                            </a>
                        </BaseBtn>
                        <!-- 小圖片區塊 -->
                        <div class="mt-3 text-sm" v-if="i.photoNames.length>0">photos:</div>
                        <div class="mt-0 flex gap-x-1">
                            <div :class="'linkimg_'+i.photoNames.length+'_block'"></div> <!-- 預留空間 -->
                            <!-- 小圖片本體 -->
                            <img v-for="i2 in i.photoNames" class="
                                shadow-lg mb-3 rounded-lg cursor-pointer
                                linkimg h-1/3
                                hover:opacity-50
                            " 
                                @click = "currentImageName = profile_photos[keyOfPhotos[i2]]['name']"
                                :src="profile_photos[keyOfPhotos[i2]]['name'].replace(/photos\//g, 'photos/small_')" 
                                :alt="profile_photos[keyOfPhotos[i2]]['desc']"
                            />
                        </div>
                    </div>
                </div>
            </div>
        </template>
    </div>
</template>


<script setup lang="ts">
    import { inject, Ref } from 'vue'
    import type { ImageItem } from '@/types/ImageItem'
    const currentImageName = inject<Ref<string>>('currentImageName')!
    const profile_photos = inject<Ref<ImageItem[]>>('profile_photos')!

    //對應 currentImageName 的 ID
    let keyOfPhotos:{[key: string]: number} = {}
    profile_photos.value.forEach((photo, index) => {
        const photoName = photo.name.split("/")[2]
        keyOfPhotos[photoName] = index
    })
    
    type data_type = {
        logoImg: string;
        title: string;
        desc: string;
        keyword: string;
        content: string;
        links: [string, string];
        photoNames: string[]; //放1~3張
    }

    let projects_data:data_type[] = [
        {   
            logoImg:"2026_twvix2.jpg",
            title:"台股恐慌資料站(2026)",
            desc:"排程爬蟲股市並AI分析的自動部屬全端系統",
            keyword:"RESTful API Gemini Golang Gin Vue3Vite Pinia Vuex VueRouter TypeScript Sass TailwindCSS DevOps GCP CloudBuild GKE Kubernetes K8s Prometheus GitHub Docker CI/CD WebCrawler Apexchart RWD Linux UnitTest Cloudflare",
            content:`<div class='text-sm'>
                獨立開發的個人網站，<b>自動部屬</b>在 GCP 上並開源於 GitHub。系統排程至證交所進行指數爬蟲，除了提供 API 接口外，亦將指數正規化、按比例權重推算PE估值，並串接 Gemini API 交由 AI 自動分析給出市場評分及估值。
                核心技術 API、Golang、Gin、VueVite、Pinia、VueRouter、TypeScript、Sass、TailwindCSS、DevOps、GCP、Kubernetes、Docker、CI/CD。
                本系統有三個 GCP 服務，Mysql、Server、CronJob。其中 Vue 編譯後靜態檔存放在 Server 的 Container；CronJob 爬蟲到的資料交給 Mysql，最後由 Server 實現前端查詢分析與 API 等服務。
                本站自動部屬流程為 Git → GitHub → CloudBuild → Registry → GKE → gke-deploy → (運行服務開放端口) → Cloudflare。
            </div>`,
            links:["https://github.com/sdfsds5271463/twvixstock", "GitHub專案源碼"],
            photoNames: ["2026_twvix5.jpg","2026_twvix2.jpg","2026_twvix1.jpg"],
        },
        {   
            logoImg:"2024_jyb1.jpg",
            title:"大型線上遊戲平台(2024)",
            desc:"傳統VM重構專案並部屬 ArgoCD GitOps",
            keyword:"Docker Kubenetes ArgoCD GitOps API PHP Laravel Vue Inertia TypeScript MySQL Redis Tailwind Sass Python Scrapy Playwright UnitTest",
            content:`
                協助將 2017 架構較老舊之系統，實現前端升級 Laravel Blade → Inertia Vue SSR，以及傳統部屬到 ArgoCD GitOps，並引入 Prometheus + Grafana + ELK 實現監控日誌與通知，設計完整 CI/CD 流程 GitHub Action → Unit Test → PR → Deploy。
                梳理舊系統商業邏輯脈絡、撰寫缺失內部文件，在不影響營運的狀況下，補足單元測試並漸進式重構系統。
                對接金流 API 及遊戲 API，並撰寫爬蟲批次控制合作遊戲商後台，重構系統的同時仍能保證營運的例行開發案。
                本站使用 PHP+Laravel 框架為主體，前端使用 Inertia Vue SSR + Tailwind CSS 進行渲染，所有 API 對接都以 PHP 進行撰寫。
                使用 Python + Scrapy + Playwright 併發同時登入30間以上遊戲商後台一鍵批次查帳、開代理。
            `,
            links:["", ""],
            photoNames: ["2024_jyb4.jpg","2024_jyb2.jpg","2024_jyb1.jpg",],
        },
        {   
            logoImg:"2022_sb2.jpg",
            title:"體育遊戲系統(2022)",
            desc:"自動部屬超過20個微服務200人同時在線的高併發系統",
            keyword:"Golang Gin gRPC Vue3 Vuex TypeScript Sass DevOps GCP CloudBuild GKE Kubernetes K8s Prometheus Grafana ELK GitHub Docker CI/CD WebSocket API Redis Linux UnitTest JWT",
            content:`
                我領導的<b>8人</b>全端技術團隊，負責將舊版 PHP 體育系統源碼重新撰寫成 Golang <b>自動部屬</b>至 GCP，包含後續開發維護等全端工作。
                核心技術 Golang、Gin、TypeScript、Vue、DevOps、GCP、Kubernetes、Docker、CI/CD、WebSocket。
                我們負責改寫專案 PHP→Go、jQuery→Vue，維運全端、API等所有服務，並全部整合進行 CI/CD。專案由<b>20個</b>以上的 GCP 容器微服務整合運行，通過 GitHub 觸發進行自動部屬，並使用 Prometheus + Grafana + ELK 監控。
                架構包含平台、前端、後端、控端、賽程、訊源、視訊等，使用 Go Channel 集中狀態併發 gRPC 給自拓展 pod 運算服務，以應付同時<b>200人</b>在線及每秒達<b>百次</b>以上的賠率運算。
            `,
            links:["", ""],
            photoNames: ["2022_sb4.jpg","2022_sb2.jpg","2022_sb1.jpg"],
        },
        {   
            logoImg:"2018_lo3.jpg",
            title:"彩票遊戲平台(2018)",
            desc:"註冊超過27萬人的概率精算遊戲",
            keyword:"RESTful API PHP Laravel MVC Twig jQuery Javascript WebCrawler API Ajax RWD MySQL Redis Bootstrap UnitTest Nginx Git",
            content:`
                我領導的<b>5人</b>全端技術團隊，從零開始獨立開發整套遊戲系統，遊戲主體為基於彩票規則的爬蟲獎號概率精算遊戲，專案主架構為前端、後端、API、排程。
                核心技術 PHP、Laravel、WebCrawler、jQuery、Bootstrap、Nginx、Git。
                本專案累計註冊會員超過<b>27萬人</b>，專案核心為排程開獎以及計算賠率，獎號的來源有官網爬蟲、獎源、我們合作遊戲商的前後台等。
                完整專案包含前台、後台建置、排程設計、撰寫 API 並發送文件給合作商對接，以及後續維運。
            `,
            links:["http://lottsample.ddns.net:9527/index.php?m=common&c=trygame&a=trygame&type=0", "彩票遊戲試玩"],
            photoNames: ["2018_lo4.jpg","2018_lo2.jpg","2018_lo1.jpg"],
        },
        {   
            logoImg:"2020_eg3.jpg",
            title:"電子遊戲(2020)",
            desc:"每日注單遠超上萬張的動畫遊戲",
            keyword:"PHP Laravel MVC Spine PixiJS jQuery Javascript API Ajax MySQL Redis Bootstrap UnitTest Nginx Photoshop",
            content:`
                我領導的<b>6人</b>全端技術團隊，從零開始獨立開發整套遊戲系統，遊戲主體為精美的動畫與機率計算，專案主架構為前端、後端、API。
                核心技術 PHP、Laravel、jQuery、Spine、PixiJS、Nginx、Git。
                使用 Spine 編輯成動畫，至前端使用 PixiJS 動畫引擎呈現，內崁入合作商捕魚遊戲，整個平台日注單可遠超<b>上萬張</b>。
                完整專案包含前台、後台建置、撰寫 API 並發送文件給合作商對接，以及後續維運。
            `,
            links:["http://lottsample.ddns.net:9527/egames/public/trygame", "電子遊戲試玩"],
            photoNames: ["2020_eg4.jpg","2020_eg3.jpg","2020_eg2.jpg"],
        },
        {   
            logoImg:"2019_li3.jpg",
            title:"區塊鏈百家樂(2019)",
            desc:"超過120子代理的hash加密紙牌遊戲",
            keyword:"PHP Laravel MVC jQuery Javascript Canvas API Ajax MySQL Redis Bootstrap UnitTest Nginx Git Photoshop",
            content:`
                我領導的<b>5人</b>全端技術團隊，從零開始獨立開發整套遊戲系統，遊戲主體為 SHA512 預加密可驗證的公平紙牌機率遊戲，專案主架構為前端、後端、API、排程。
                核心技術 PHP、Laravel、jQuery、Canvas、Bootstrap、Nginx、Git。
                使用 jQuery 與畫布 Canvas 呈現生動的牌桌畫面，並通過排程預 hash 以及開獎；本遊戲崁入合作商的遊戲平台，創建了超過<b>120個</b>子代理。
                完整專案包含前台、後台建置、排程設計、撰寫 API 並發送文件給合作商對接，以及後續維運。
            `,
            links:["http://livesample.ddns.net:9527/game/trygame", "區塊鏈百家樂試玩"],
            photoNames: ["2019_li4.jpg","2019_li3.jpg","2019_li2.jpg"],
        },
        {   
            logoImg:"2017_bos1.jpg",
            title:"遊戲後台API對接(2017)",
            desc:"技術對接40個以上合作商接口",
            keyword:"RESTful API Json PHP Laravel MVC MySQL Redis Git",
            content:`
                項目成立初期加入該團隊，專精負責 API 對接的業務，本專案我總計對接了<b>30個</b>以上金流 API、<b>10個</b>以上遊戲 API。
                核心技術 PHP、Laravel、API、RESTful API、Json、Git。
                主要工作流程為接洽合作商、開設代理加白並索取對接文件，經過測試環境正式環境到對接完成上線。
                遊戲 API 有創建會員、無縫或轉帳、啟動遊戲、報表等固定接口；金流 API 有代付代收、請求、回調、查詢等固定接口。我的任務就是完善這個流程。
            `,
            links:["", ""],
            photoNames: ["2017_bos3.jpg","2017_bos2.jpg","2017_bos1.jpg"],
        },
        {   
            logoImg:"2016_yuntech2.jpg",
            title:"內容管理E化流程系統(2016)",
            desc:"後台輕鬆更新前端內容、圖片的管理系統",
            keyword:"WordPress PHP MySQL Javascript jQuery Photoshop SEO Html Css Apache",
            content:`
                自2012年起職位為雲科研發處網管，總共撰寫了超過<b>20個</b>上線網站，協助校內建置靜態網站、CMS 內容管理系統及E化政府單位流程的介面化資料庫網頁系統。。
                核心技術 WordPress、PHP、MySQL、jQuery、Html、Css、SEO、Apache。
                撰寫了大量校內活動靜態網頁，學會了極強的 Photoshop 素材編輯能力以及基礎的 SEO 推廣。
                協助將各處室的靜態網站以 Apache + WordPress 建成教職員自行編輯內容、上傳圖片的 CMS 內容管理系統。
                協助政府單位建置了數個 PHP + MySQL 的E化投標或公文等系統，項目內容包含介面化填寫、比對資料，寄發郵件等。
            `,
            links:["", ""],
            photoNames: ["2016_yuntech3.jpg","2016_yuntech2.jpg","2012_yuntech1.jpg"],
        },
    ];

    //content 讓 \n 變成 <p><\p>
    projects_data.forEach((v,k)=>{
        //整理內容 (變成 html顯示好看的)
        v.content = v.content.trim();
        let exp1 = v.content.split("\n");
        for(let k2 in exp1){
            exp1[k2] = exp1[k2].trim();  //每行去掉空白
        }
        v.content = "<p>" + exp1.join("<\p><p>") + "<\p>"; //段落包好
    });
</script>


<style lang="scss" scoped>
    .linkimg{ //小圖片寬度(1~3張)
        width: calc(33% - 2px);
        transition: opacity 0.3s ease;
    }
    .linkimg_1_block{ //1張小圖片置中預留寬
        width: 32%;
    }
    .linkimg_2_block{ //2張小圖片置中預留寬
        width: 15%;
    }
    .linkimg_3_block{ //3張小圖片就不用留寬度
        display: none;
    }

    .keyword{ //關鍵字
        line-height: 12px;
    }

    .project_content p{ //計畫內容
        line-height: 18px;
        padding: 2px 0px;
    }

    .basecard_custom {  //自訂卡片
        border-radius: 10px;
        box-shadow: 0 4px 20px 1px rgba(0, 0, 0, 0.06), 0 1px 4px rgba(0, 0, 0, 0.08);
    }

</style>