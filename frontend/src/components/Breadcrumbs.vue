<script setup>
const props = defineProps({
    parentTitle: String,
    subParentTitle: String
})
</script>

<template>
    <div class="breadcrumb flex items-end border-b border-gray-300 pb-4 mb-6 pl-3">
        <h1 class="text-xl mr-1 font-semibold mr-2 dark:text-white"> {{ props.parentTitle }} </h1>
        <ul>
            <li class="inline pl-2 dark:text-gray-400"><h2 class="text-lg"> {{ props.subParentTitle }}</h2></li>
        </ul>
    </div>
</template>

