<script setup lang="ts">
import { ref,Ref,inject } from 'vue'
import Lightbox from 'vue-my-photos'
import type { ImageItem } from '@/types/ImageItem'
import router from '../router'

/*
使用時先引入
script setup lang="ts">
    import { inject, Ref } from 'vue'
    import type { ImageItem } from '@/types/ImageItem'
    const currentImageName = inject<Ref<string>>('currentImageName')!
    const profile_photos = inject<Ref<ImageItem[]>>('profile_photos')!
測試console.log(profile_photos.value[0]['name']);
然後<span @click="currentImageName = 'images/Sass_icon.png'">123</span>  這樣可以觸發
*/

//參數
const profile_photos = inject<Ref<ImageItem[]>>('profile_photos')!
const currentImageName = inject<Ref<string>>('currentImageName')!

// UI本體
const lightboxRef = ref<any>(null)

//上一頁可以關閉 UI
router.beforeEach((to, from, next) => {
  if (currentImageName.value != "") {
    currentImageName.value = "";
    lightboxRef.value.hide()
    return next(false) // 阻止返回
  }
  next()
})

//點擊圖片可以關閉
const SideClose = (event: any)=>{
    let close:boolean = true
    if(event.target.tagName == "circle"){
        close = false;  // 按下箭頭符號不關閉
    }else if(event.target.className.indexOf("lightbox-arrow") != -1){
        close = false;  // 按下箭頭區域不關閉
    }
    if(close){
        lightboxRef.value.hide()
    }
}
</script>

<template>
    <!-- 輪播本體 -->
    <Lightbox
        ref="lightboxRef"
        @click="SideClose($event)"
        :images="profile_photos"
        :current-image-name="currentImageName"
        @on-lightbox-close="currentImageName = ''"
    />
</template>

<style lang="scss">
    .lightbox .lightbox-image {
        animation: lightboxFadeIn 0.1s ease;
    }
    .lightbox {
        animation: lightboxFadeIn 0.3s ease;
        padding-top: 10vh;
        .lightbox-image-container {
            max-height: 80vh;
        }
    }
    @keyframes lightboxFadeIn {
        from {
            opacity: 0;
            transform: scale(0.95);
        }
        to {
            opacity: 1;
            transform: scale(1);
        }
    }
</style>

