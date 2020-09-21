<template>
  <div id="app">
    <div id="nav">
      <top-header></top-header>
    </div>
    <router-view />
  </div>
</template>

<script>
import TopHeader from "@/components/TopHeader.vue";
export default {
  name: 'App',
  components: { "top-header": TopHeader }
}
</script>