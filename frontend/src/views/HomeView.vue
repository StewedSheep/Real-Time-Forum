<script>
import Vue from "vue";

export default {
  name: "HomeView.vue",
  components: {},
  data (){
    return {
      law: null,
    };
  },
  methods : {
    getLaw() {
      Vue.axios.get("/law").then((response) => (this.law = response.data));
    },
  },
}
</script>

<template>
  <div id="app">
    <div id="Law">
    <button type="button" @click="getLaw()">Get a new hacker law</button>
    <div v-if="law != null">
      <h1>{{ law.name }}</h1>
      <p>{{ law.definition }}</p>
    </div>
    </div>
    <router-view />
</div>
</template>

