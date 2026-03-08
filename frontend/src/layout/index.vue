<script setup lang="ts">
import { useStore } from 'vuex'
import Header from './Header.vue'
import Sidebar from './Sidebar.vue'
import Footer from './Footer.vue'
import LightboxUI from '@/components/LightboxUI.vue'

let store = useStore();

// 給圖片輪播用的資料
import { ref,provide } from 'vue'
import type { ImageItem } from '@/types/ImageItem'
const profile_photos = ref<ImageItem[]>([
  { name: 'images/profile_photos/2024_jyb3.jpg',    desc: '遊戲平台手機板' , year: '2024' },
  { name: 'images/profile_photos/2024_jyb2.jpg',    desc: '遊戲平台前台2' , year: '2024' },
  { name: 'images/profile_photos/2024_jyb1.jpg',    desc: '遊戲平台前台1' , year: '2024' },
  { name: 'images/profile_photos/2022_sb4.jpg',     desc: '體育系統手機板' , year: '2022' },
  { name: 'images/profile_photos/2022_sb3.jpg',     desc: '體育系統2' , year: '2022' },
  { name: 'images/profile_photos/2022_sb2.jpg',     desc: '體育系統1' , year: '2022' },
  { name: 'images/profile_photos/2022_sb1.jpg',     desc: '體育系統架構圖' , year: '2022' },
  { name: 'images/profile_photos/2020_eg4.jpg',     desc: '電子遊戲前台3' , year: '2020' },
  { name: 'images/profile_photos/2020_eg3.jpg',     desc: '電子遊戲前台2' , year: '2020' },
  { name: 'images/profile_photos/2020_eg2.jpg',     desc: '電子遊戲前台1' , year: '2020' },
  { name: 'images/profile_photos/2020_eg1.jpg',     desc: '電子遊戲後台' , year: '2020' },
  { name: 'images/profile_photos/2019_li4.jpg',     desc: '百家樂手機板' , year: '2019' },
  { name: 'images/profile_photos/2019_li3.jpg',     desc: '百家樂前台2' , year: '2019' },
  { name: 'images/profile_photos/2019_li2.jpg',     desc: '百家樂前台1' , year: '2019' },
  { name: 'images/profile_photos/2019_li1.jpg',     desc: '百家樂後台' , year: '2019' },
  { name: 'images/profile_photos/2018_lo4.jpg',     desc: '彩票遊戲手機板' , year: '2018' },
  { name: 'images/profile_photos/2018_lo3.jpg',     desc: '彩票遊戲前台2' , year: '2018' },
  { name: 'images/profile_photos/2018_lo2.jpg',     desc: '彩票遊戲前台1' , year: '2018' },
  { name: 'images/profile_photos/2018_lo1.jpg',     desc: '彩票遊戲後台' , year: '2018' },
  { name: 'images/profile_photos/2017_bos2.jpg',    desc: '對接金流文件' , year: '2017' },
  { name: 'images/profile_photos/2017_bos.jpg',     desc: '遊戲系統後台' , year: '2017' },
  { name: 'images/profile_photos/2016_yuntech3.jpg',     desc: '校園徵才廠商登錄網' , year: '2016' },
  { name: 'images/profile_photos/2016_yuntech2.jpg',     desc: '雲科研發處官網' , year: '2016' },
  { name: 'images/profile_photos/2016_yuntech1.jpg',     desc: '雲科卓越典範人物' , year: '2016' },
  { name: 'images/profile_photos/2015_yuntech2.jpg',     desc: '雲科25周年校慶官網' , year: '2015' },
  { name: 'images/profile_photos/2015_yuntech1.jpg',     desc: '反毒紫錐花運動' , year: '2015' },
  { name: 'images/profile_photos/2014_yuntech3.jpg',     desc: '中華民國馬術協會' , year: '2014' },
  { name: 'images/profile_photos/2014_yuntech2.jpg',     desc: '雲科數位校史館' , year: '2014' },
  { name: 'images/profile_photos/2014_yuntech1.jpg',     desc: '雲科資工系網' , year: '2014' },
  { name: 'images/profile_photos/2012_yuntech2.jpg',     desc: '雲科人命名活動' , year: '2012' },
  { name: 'images/profile_photos/2012_yuntech1.jpg',     desc: '政府投標審核系統' , year: '2012' },
])
const currentImageName = ref<string>('')

provide('currentImageName', currentImageName); 
provide('profile_photos', profile_photos); 

</script>

<template>
    <LightboxUI />
    <div class="app-admin-wrap-layout-2">
        <Header />
        <Sidebar />
        <div :class="store.state.largeSidebar.sidebarToggleProperties.isSideNavOpen === true ? '': 'full'" class="main-content-wrap">
            <main>
                <div class="main-content-wrap flex flex-col flex-grow print-area pt-10">
                    <div>
                        <router-view />
                    </div>
                    
                    <Footer />
                </div>
            </main>
        </div>
    </div>
</template>

<style lang="scss" scoped>
    .app-admin-wrap-layout-2 {
        width: 100%;
        height: 100%;
        .main-content-wrap {
            width: calc(100% - 100px);
            margin-left: 80px;
            // min-height: 100vh;
            margin-top: 80px;
            transition: all 0.24s ease-in-out;
            .main-content-body{
                min-height: calc(100vh - 80px);
            }
            &.full {
                width: 100%;
                margin-left: 0px;
                transition: all 0.24s ease-in-out;
            }
            @media screen and (max-width: 991px) {
                width: 100%;
                margin-left: 0px;
                padding-right: 0px;
                padding-left: 0px;
            }
        }
    }
</style>
