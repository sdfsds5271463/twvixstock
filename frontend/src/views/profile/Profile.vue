<script setup lang="ts">
  import Breadcrumb from '@/components/Breadcrumbs.vue'
  import Timeline from './Timeline.vue'
  import Aboutme from './Aboutme.vue'
  import Projects from './Projects.vue'
  import Photos from './Photos.vue'
  import { ref,onMounted,onUnmounted } from 'vue'
  import { TabGroup, TabList, Tab, TabPanels, TabPanel } from '@headlessui/vue'
  //import About from '../About.vue'
  import { useSeoMeta } from '@unhead/vue'
  useSeoMeta({
    title: '人物簡介|台股恐慌資訊站 TwVixStock',
    description: '台股恐慌資訊站作者經歷以及個人相關作品',
    ogImage: 'http://twvixstock.qzz.io/images/profile_photos/small_2026_twvix2.jpg',
    ogTitle: '人物簡介|台股恐慌資訊站 TwVixStock',
    ogType: 'website',
    ogLocale: 'zh_TW',
    keywords: '台股, 恐慌指數, VIX, 股票投資, TwVixStock, 台灣加權, 本益比, 人物簡介',
  });

  let categories = ref(['關於我', '時間線', '專案清單', '專案圖'])
  let categories_icon = ref(['i-Business-Man', 'i-Calendar-4', 'i-Professor', 'i-Camera'])

  //以下 試玩球功能
  const tryGameBall = ref(null);
  const ballIndex = ref(0);
  let tryBallLink:[string,string,string][] = [  //連結清單
    ["2018","彩票","http://lottsample.ddns.net:9527/index.php?m=common&c=trygame&a=trygame&type=0"],
    ["2019","百家樂","http://livesample.ddns.net:9527/game/trygame"],
    ["2020","電子","http://lottsample.ddns.net:9527/egames/public/trygame"],
    ["2026","GitHub","https://github.com/sdfsds5271463/twvixstock"],
  ];
  let timer = null

  onMounted(() => {
    timer = setInterval(()=>{
      let ballWord = tryGameBall.value.childNodes[1];  //文字元素
      ballWord.classList.add('blinking');  //開始閃爍
      setTimeout(()=>{
        ballIndex.value += 1;
        ballIndex.value %= 4;  //換文字
      },1000);
      setTimeout(()=>{
        ballWord.classList.remove("blinking");  //重置閃爍
      },2000);
    },6000); 

  })

  onUnmounted(() => {
    clearInterval(timer)  //很重要!!
  })

  function openLink() {
    window.open(tryBallLink[ballIndex.value][2], '_blank')
  }
  //以上 試玩球功能


  //上跳錨點
  const jumpTop = () => {
    const el = document.getElementById('target-jumpTop');
    const rect = el.getBoundingClientRect(); //取得位置

    if(rect.y < 0){  //往下滾之後才需要上彈
      el?.scrollIntoView({ 
        behavior: 'smooth', 
        block: 'start' // 讓元素對齊視窗頂部
      });
    }
  };


</script>

<template>

  <div class="container mx-auto">
    <Breadcrumb parentTitle='本站作者人物簡介' subParentTitle='Profile' />

    <BaseCard noPadding class="user-profile relative">
        <div class="header-cover">
          <div class="text-white text-center text-4xl">Allen Zheng</div>
          <div class="text-white text-center">資深全端工程師 / 技術主管</div>
        </div>
        <div class="flex justify-center z-10 -mt-10">
            <div class="text-center">
                <img class="relative z-1 w-20 h-20 m-auto rounded-full border-2 border-white" src="/images/faces/allen_small_head2.jpg" />
                <div class="profile_desc text-gray-600">
                  <p class="text-blue-500 font-extrabold"># PHP Laravel Golang TypeScript Vue<br class="rwd"> DevOps GCP Docker Kubenetes</p>
                  <p>▪️13年全端經驗，<b>PHP</b>(13年) <b>Javascript</b>(13年) <b>Vue</b>(4年) <b>Typescript</b><br class="rwd">(2年) <b>Golang</b>(2年) <b>Docker</b>(4年) <b>Kubenetes</b>(4年)</p> 
                  <p>▪️從傳統 VM 到自動部屬、GitOps 再到將 AI 引入 PR <br class="rwd">流程，曾管理20+個微服務的分散式系統</p>
                  <p>▪️擅長前後端系統重構與漸進式系統升級，包含容<br class="rwd">器化、自動化、前後端分化、分散式部屬</p>
                  <p>▪️擔任7年主管，帶領5~8人團隊從零開始建置過三個<br class="rwd">完整功能平台，對接超過30+個金流API</p>
                  <p>▪️執掌開發營運的系統，有200+的同時在線人數、27萬+<br class="rwd">的總註冊人數，每月處理超過萬張金流有效單</p>
                </div>
            </div>
        </div>
        <div id="target-jumpTop"></div><!-- 彈跳點 -->
        <div class="mt-0 p-5">
          <TabGroup>
            <TabList class="flex justify-center newBtnDiv">
                <Tab
                    v-for="(category, index) in categories"
                    as="template"
                    :key="category"
                    v-slot="{ selected }"
                >
                <div
                    @click="jumpTop"
                    class="text-sm"
                    :class="[
                    'hover:bg-primary hover:text-white',
                    selected
                        ? 'btnActivy bg-primary text-white'
                        : '',
                    ]"
                ><i class="tablist_icon pr-2" :class="categories_icon[index]"></i>{{ category }} 
                </div>
              </Tab>
            </TabList>

            <TabPanels class="mt-6">
              <TabPanel>
                <Aboutme></Aboutme>
              </TabPanel>
              <TabPanel>
                <Timeline></Timeline>
              </TabPanel>
              <TabPanel>
                <Projects></Projects>
              </TabPanel>
              <TabPanel>
                <Photos></Photos>
              </TabPanel>
            </TabPanels>
          </TabGroup>
        </div>
        
    </BaseCard>
  </div>



  <!-- 試玩球 -->
  <div ref="tryGameBall" class="
    tryGameBall rainbow-flow
    fixed z-50 w-16 h-16 rounded-full text-center text-sm pt-1 
    text-white cursor-pointer
    hover:opacity-90 hover:scale-110"
    @click="openLink"
  >
    try
    <div>
      <b v-text="tryBallLink[ballIndex][1]"></b><br><span v-text="tryBallLink[ballIndex][0]"></span>
    </div>
  </div>

</template>

<style lang="scss" scoped>

//以下試玩球
.tryGameBall{
  box-shadow: 1px 1px 6px black;
  bottom:5vh; 
  right:8px;
  line-height: 18px;
  text-shadow: 1px 1px 8px black;
  transition: transform 0.3s ease;
  b{
    font-size: 16px;
  }
}
.rainbow-flow {
  background: linear-gradient(
    270deg,
    #6f00ff,
    #d400ff,
    #ff0000,
    #ff0077,
    #6f00ff,
    #0004fd,
    #0084ff,
    #0066ff,
    #6f00ff,
  );
  background-size: 800% 800%;
  animation: rainbow-flow 18s ease infinite;
}

@keyframes rainbow-flow {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

.blinking{
  animation: blinking 2s ease;
}
@keyframes blinking {
  0% {
    opacity: 100%;
  }
  50% {
    opacity: 0%;
  }
  100% {
    opacity: 100%;
  }
}
//以上試玩球


.user-profile {
  .header-cover {
    background-image: url("../../../images/photo-wide.jpg") ;
    background-color: purple;
    position: relative;
    background-size: cover;
    background-repeat: no-repeat;
    height: 200px;
    padding-top: 95px;
    text-shadow: 0px 0px 20px rgb(34, 2, 37);
    background-clip: border-box;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
  }
}
.profile_desc .rwd {
  display: none;
}
.profile_desc p {
  line-height: 18px;
  padding: 4px 0px;
}

// 新按鈕
.newBtnDiv{
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    color: #AAA;
    padding: 4px 4px;
    margin: 0px 0px;
    background-color: rgba(255, 255, 255, 0.8);

    position: sticky; //吸頂
    top: 45px;
    z-index: 10;

    >div{
        background-color: #EEE;
        border-radius: 6px;
        padding: 8px 12px;
        margin: 2px 4px;
        cursor: pointer;
        user-select: none;
        box-shadow: 1px 1px 4px #545454;
        transition: background-color 0.3s,color 0.3s, scale 0.3s; 
    }
    >div:hover{
        color: #eee;
        background-color: #d1b8e6;
    }
    >div.btnActivy{
        color: #fff;
        font-weight: 600;
        background-color: #a457e6;
        scale: 110%;
    }
}

// RWD
@media (max-width: 767px) {
  .profile_desc .rwd {
    display: inline;
  }
}
@media (max-width: 500px) {
  .tablist_icon {
      display: none;
  }
  .profile_desc {
    font-size: 14px;
  }
  .profile_desc p {
    line-height: 16px;
  }
  .newBtnDiv div{
    font-size: 12px;
    padding: 8px 10px;
  }
}


</style>
