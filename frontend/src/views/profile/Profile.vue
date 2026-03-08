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
    <Breadcrumb parentTitle='Profile' subParentTitle='ProfileV2' />
    
    <BaseCard noPadding class="user-profile overflow-hidden relative">
        <div class="header-cover">
          <div class="text-white text-center text-4xl">Allen Zheng</div>
          <div class="text-white text-center">資深網路工程師</div>
        </div>
        <div class="flex justify-center z-10 -m-10">
            <div class="text-center"><img class="relative z-1 w-20 h-20 m-auto rounded-full border-2 border-white" src="/images/faces/allen_small_head.jpg" />
                <p class="profile_desc text-gray-600">
                  ▪️15年全端開發經歷，6年管理經歷，深耕<br class="rwd">高併發架構與自動化交易系統。<br>
                  ▪️撰寫超過20個上線網站，對接超過30個<br class="rwd">金流API、超過10個遊戲API。<br>
                  ▪️帶領5人團隊，從零開始獨立開發過三種不同類型<br class="rwd">遊戲，總不重複註冊人數超過27萬人。<br>
                  ▪️協作開發的體育系統同時在線人數超過200人，自<br class="rwd">動部屬管理該系統20個以上微服務。<br>
                  ▪️登入爬蟲超過30間合作廠商後台，達成一鍵式對單<br class="rwd">及開代理等系列合作商後台批次操作。
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
    height: 300px;
    padding-top: 195px;
    text-shadow: 0px 0px 20px rgb(34, 2, 37);
  }
  &::after {
    content: "";
    width: 100%;
    height: 300px;
    position: absolute;
    background: rgba(0,0,0,.1)
  }
}

.profile_desc .rwd {
  display: none;
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
}


</style>
