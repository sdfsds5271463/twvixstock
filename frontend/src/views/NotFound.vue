<template>
  <div class="notfound text-primary text-center h-screen">
    <h1 class="animate-pulse">Page Not Found</h1>
    <h3 class="opacity-50 cursor-pointer hover:opacity-100">
      <router-link to="/">Let's go home</router-link>
    </h3>
  </div>
</template>

<style lang="scss">
  .notfound{
    padding-top: 40vh
  }
</style>