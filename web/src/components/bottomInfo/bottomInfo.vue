<template>
  <div class="flex flex-col md:flex-row gap-2 items-center text-sm text-slate-700 dark:text-slate-500 justify-center py-2">
    <div class="text-center">
      <span class="mr-1">Powered by</span>
      <span>
        <a
          class="font-bold text-active"
          href="https://github.com/richardgong1987"
        >GONG</a>
      </span>
    </div>
    <slot />
    <div class="text-center">
      <span class="mr-1">Copyright</span>
      <span>
        <a
          class="font-bold text-active"
          href=""
        >GONG团队</a>
      </span>
    </div>
  </div>
</template>

<script  setup>
defineOptions({
  name: 'BottomInfo'
})
</script>

