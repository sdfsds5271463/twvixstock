<template>
  <div class="p-4">
    <span @click="currentImageName = 'images/Sass_icon.png'">123</span>
  </div>
</template>

<script setup lang="ts">
    import { inject, Ref } from 'vue'
    import type { ImageItem } from '@/types/ImageItem'
    const currentImageName = inject<Ref<string>>('currentImageName')!
    const profile_photos = inject<Ref<ImageItem[]>>('profile_photos')!
    console.log(profile_photos.value[0]['name']);
    //<span @click="currentImageName = 'images/Sass_icon.png'">123</span>  這樣可以觸發
</script>
