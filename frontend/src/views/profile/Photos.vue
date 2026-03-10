<template>
    <div class="grid grid-cols-12 gap-5">
        <template v-for="i in profile_photos">
            <div class="col-span-12 lg:col-span-4 md:col-span-6">
                <BaseCard noPadding class="relative overflow-hidden cursor-pointer photos_img"
                    @click="currentImageName=i.name"
                >
                    <img class="w-full" :src="i.name.replace(/photos\//g, 'photos/small_')" alt="" />
                    <div class="card-img-overlay-profile absolute bottom-0 p-4">
                        <div class="p-1 text-left font-light flex">
                            <span class="
                                photos_desc 
                                font-black flex items-center text-white 
                                rounded-lg px-2"
                            >
                                <i class="i-Calendar-4 mr-2"></i>{{i.year}} {{i.desc}}
                            </span>
                        </div>
                    </div>
                </BaseCard>
            </div>
        </template>
    </div>
</template>

<script setup lang="ts">
    import { inject, Ref } from 'vue'
    import type { ImageItem } from '@/types/ImageItem'
    const currentImageName = inject<Ref<string>>('currentImageName')!
    const profile_photos = inject<Ref<ImageItem[]>>('profile_photos')!
</script>

<style lang="scss" scoped>
    .photos_desc{
        text-shadow: 0px 0px 8px black;
    }
    .photos_img{
        transition: transform 0.3s ease;
    }
    .photos_img:hover{
        transform: scale(1.05);
    }
</style>