<script setup>
const props = defineProps({
    parentTitle: String,
    subParentTitle: String
})
</script>

<template>
    <div class="breadcrumb flex items-end border-b border-gray-300 pb-4 mb-6 pl-3">
        <p class="text-xl mr-1 font-semibold mr-2 dark:text-white"> {{ props.parentTitle }} </p>
        <ul>
            <li class="inline pl-2 dark:text-gray-400"> {{ props.subParentTitle }}</li>
        </ul>
    </div>
</template>

