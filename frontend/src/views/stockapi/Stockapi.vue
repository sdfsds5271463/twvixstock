<script setup lang="ts">
  import Breadcrumb from '@/components/Breadcrumbs.vue'
</script>

<template>
  <div class="container mx-auto">
    <Breadcrumb parentTitle='Stock Api' subParentTitle='ProfileV2' />
  </div>
    <div class="p-5">
        <p class="text-xl font-semibold mb-2">Personal Information</p>
        <p class="border-b border-gray-300 pb-10 mb-8">Lorem ipsum dolor sit amet consectetur adipisicing elit. Eveniet, commodi quam! Provident quis voluptate asperiores ullam, quidem odio pariatur. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptatem, nulla eos? Cum non ex voluptate corporis id asperiores doloribus dignissimos blanditiis iusto qui repellendus deleniti aliquam, vel quae eligendi explicabo.</p>
        <div class="grid grid-cols-12 gap-5 border-b border-gray-300 pb-10 mb-10">
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Calendar text-base"></i><span class="ml-1">Birth Date</span></p><span>1, Jan 1994</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-MaleFemale text-base"></i><span class="ml-1">Gender</span></p><span>1, Jan 1994</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Face-Style-4 text-base"></i><span class="ml-1">Profession</span></p><span>Digital Marketer</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Calendar text-base"></i><span class="ml-1">Birth Date</span></p><span>1, Jan 1994</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-MaleFemale text-base"></i><span class="ml-1">Gender</span></p><span>1, Jan 1994</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Face-Style-4 text-base"></i><span class="ml-1">Profession</span></p><span>Digital Marketer</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Globe text-base"></i><span class="ml-1"> Lives In </span></p><span>8 Years</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Edit-Map text-base"></i><span class="ml-1">Email</span></p><span>example@ui-lib.com</span>
            </div>
            <div class="col-span-12 sm:col-span-6 md:col-span-4">
                <p class="text-purple-500 mb-1"><i class="i-Professor text-base"></i><span class="ml-1"> Experience</span></p><span>8 Years</span>
            </div>
        </div>
        <div class="my-5">
            <p class="text-2xl">Timothy Carlson</p>
            <p class="text-gray-600">Digital Marketer</p>
        </div>
        <div class="flex justify-around flex-wrap">
            <div class="text-center mx-4 my-2"><i class="i-Plane text-4xl text-purple-500"></i>
                <p class="text-base">Travelling </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Camera text-4xl text-purple-500"></i>
                <p class="text-base">Photography </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Car-3 text-4xl text-purple-500"></i>
                <p class="text-base">Driving </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Gamepad-2 text-4xl text-purple-500"></i>
                <p class="text-base">Gaming </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Music-Note-2 text-4xl text-purple-500"></i>
                <p class="text-base">Music </p>
            </div>
            <div class="text-center mx-4 my-2"><i class="i-Shopping-Bag text-4xl text-purple-500"></i>
                <p class="text-base">Shopping </p>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>

/*@media (max-width: 767px) {
  .user-profile {
      .header-cover {
          height: 200px;
      }
  }
}
@media (max-width: 414px) {
  .tablist_icon {
      display: none;
  }
}*/


</style>
