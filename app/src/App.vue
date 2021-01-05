<template>
  <div id="app">
    <div id="nav">
      <top-header></top-header>
    </div>
    <b-container class="mt-1">
      <router-view />
    </b-container>
  </div>
</template>

<script>
import TopHeader from "@/components/TopHeader.vue";
export default {
  name: "App",
  components: { "top-header": TopHeader },
};
</script>
