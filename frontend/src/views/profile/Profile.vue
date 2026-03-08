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
        <div class="header-cover"></div>
        <div class="flex justify-center z-10 -m-10">
            <div class="text-center"><img class="relative z-1 w-20 h-20 m-auto rounded-full border-2 border-white" src="/images/faces/allen_small_head.jpg" />
                <p class="text-2xl">Allen Zheng</p>
                <p class="text-gray-600">資深網路工程師</p>
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
                    class="mb-3 mr-1 border border-primary text-primary"
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
   background-image: url("../../../images/photo-wide-5.jpeg") ;
   position: relative;
   background-size: cover;
   background-repeat: no-repeat;
   height: 300px;
 }
  &::after {
    content: "";
    width: 100%;
    height: 300px;
    position: absolute;
    background: rgba(0,0,0,.1)
  }
}
/*@media (max-width: 767px) {
  .user-profile {
      .header-cover {
          height: 200px;
      }
  }
}*/
@media (max-width: 414px) {
  .tablist_icon {
      display: none;
  }
}


</style>
