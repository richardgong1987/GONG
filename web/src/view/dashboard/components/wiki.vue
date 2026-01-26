<template>
  <div class="grid grid-cols-2 gap-2">
    <a
      v-for="item in wikis"
      :key="item.url"
      :href="item.url"
      class="text-sm text-black/70 dark:text-white/70 no-underline hover:text-[var(--el-color-primary)] dark:hover:text-white"
      target="_blank"
    >
      {{ item.title }}
    </a>
  </div>
</template>

<script setup>
const wikis = [


]
</script>

<style scoped lang="scss">

</style>
