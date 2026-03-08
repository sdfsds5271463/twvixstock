<script setup lang="ts">
  import Breadcrumb from '@/components/Breadcrumbs.vue'
  import Timeline from './Timeline.vue'
  import Aboutme from './Aboutme.vue'
  import Projects from './Projects.vue'
  import Photos from './Photos.vue'
  import { ref } from 'vue'
  import { TabGroup, TabList, Tab, TabPanels, TabPanel } from '@headlessui/vue'
  import About from '../About.vue'
  let categories = ref(['Timeline', 'About', 'Projects', 'Photos'])
  let categories_icon = ref(['i-Calendar-4', 'i-Business-Man', 'i-Professor', 'i-Camera'])
</script>

<template>
 
  <div class="container mx-auto">
    <Breadcrumb parentTitle='Profile' subParentTitle='個人資料介紹' />
    
    <BaseCard noPadding class="user-profile overflow-hidden relative">
        <div class="header-cover">
          <div class="text-white text-center text-4xl">Allen Zheng</div>
          <div class="text-white text-center">資深網路工程師</div>
        </div>
        <div class="flex justify-center z-10 -m-10">
            <div class="text-center"><img class="relative z-1 w-20 h-20 m-auto rounded-full border-2 border-white" src="/images/faces/allen_small_head.jpg" />
                <p class="profile_desc text-gray-600">
                  <p>▪️<b>15年</b>全端開發經歷，<b>6年</b>管理經歷，深耕<br class="rwd">高併發架構與自動化交易系統。</p>
                  <p>▪️撰寫超過<b>20個</b>上線網站，對接超過<b>10個</b>遊戲API<br class="rwd">，亦主力開發過<b>4個</b>大型專案的API。</p>
                  <p>▪️對接超過<b>30個</b>金流API，由我完全負責的金流<br class="rwd">接口每月處理超過<b>1萬張</b>有效單。</p>
                  <p>▪️帶領<b>5人</b>團隊，從零開始獨立開發過<b>3種</b>不同類型<br class="rwd">遊戲，不重複註冊總人數超過<b>27萬</b>人。</p>
                  <p>▪️協作開發的體育系統同時在線人數超過<b>200人</b>，自<br class="rwd">動部屬管理該系統<b>20個</b>以上微服務。</p>
                  <p>▪️登入爬蟲超過<b>30間</b>合作廠商後台，達成一鍵式對單<br class="rwd">及開代理等系列合作商後台批次操作。</p>
                </p>
            </div>
        </div>
        <div class="mt-10 p-5">
         
          <TabGroup>
            <TabList class="flex bg-white justify-center">
                <Tab
                    v-for="(category, index) in categories"
                    as="template"
                    :key="category"
                    v-slot="{ selected }"
                >
                <BaseBtn 
                    class="mx-1 mb-3 border border-primary text-primary"
                    :class="[
                    'hover:bg-primary hover:text-white',
                    selected
                        ? 'bg-primary text-white'
                        : '',
                    ]"
                ><i class="tablist_icon pr-3" :class="categories_icon[index]"></i>{{ category }} 
                </BaseBtn>
              </Tab>
            </TabList>

            <TabPanels class="mt-10">
              <TabPanel>
                <Timeline></Timeline>
              </TabPanel>
              <TabPanel>
                <Aboutme></Aboutme>
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
</template>

<style lang="scss" scoped>
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
  }
  &::after {
    content: "";
    width: 100%;
    height: 202px;
    position: absolute;
    background: rgba(0,0,0,.1)
  }
}

.profile_desc .rwd {
  display: none;
}
.profile_desc p {
  line-height: 18px;
  padding: 4px 0px;
}
@media (max-width: 767px) {
  .profile_desc .rwd {
    display: inline;
  }
}
@media (max-width: 414px) {
  .tablist_icon {
      display: none;
  }
  .profile_desc {
    font-size: 14px;
  }
  .profile_desc p {
    line-height: 16px;
  }
}


</style>
