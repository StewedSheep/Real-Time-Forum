<script>
import Vue from "vue";

export default {
  name: "HomeView.vue",
  components: {},
  data (){
    return {
      Threads: [],
    };
  },
  mounted() {
          Vue.axios.get("/totThread").then((response) => (this.Threads = response.data));
    },
}
</script>

<template>
  <div id="app">
    <div id="threadList">
    <div v-if="Threads != null">
      <div v-for="Thread in Threads" :key="Thread">
      <h1>{{ Thread.title }}</h1>
      <p>{{ Thread }}</p>
    </div>
    </div>
    </div>
</div>
</template>

