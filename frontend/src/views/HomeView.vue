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
      <div v-if="!$user.isAuthorised">
          <p>Sort posts by:</p>
        <div class="btn-group" style="width:100%">
          <button href="general" style="width:20%">General</button>
          <button href="help" style="width:20%">Help</button>
          <button href="news" style="width:20%">News</button>
          <button href="discussion" style="width:20%">Discussion</button>
          <button href="offTopic" style="width:20%">Off-Topic</button>
        </div>
      <br><br>
      </div>

      <div v-if="$user.isAuthorised">
          <p>Sort posts by:</p>
      <div class="btn-group" style="width:100%">
          <button href="general" style="width:14.28%">General</button>
          <button href="help" style="width:14.28%">Help</button>
          <button href="news" style="width:14.28%">News</button>
          <button href="discussion" style="width:14.28%">Discussion</button>
          <button href="offTopic" style="width:14.28%">Off-Topic</button>
          <button href="likedPosts" style="width:14.28%">Liked-Posts</button>
          <button href="myPosts" style="width:14.28%">My-Posts</button>
        </div>
        <br><br>
        </div>

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

