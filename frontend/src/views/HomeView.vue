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
    methods:{
    liGen: function(){
      this.cate = "General"
    },
    liHel: function(){
      this.cate = "Help"
    },
    liNew: function(){
      this.cate = "News"
    },
    liDis: function(){
      this.cate = "Discussion"
    },
    liOff: function(){
      this.cate = "Off-Topic"
    },
    liLik: function(){
      // this.cate = "Liked-Posts" QOTD: HOW TO FILTER LIKED POSTS 
    },
    liMyP: function(){
      this.cate = "My-Posts"
      this.auth = this.$user.current
    },
  }
}
</script>

<template>
  <div id="app">
      <div v-if="!$user.isAuthorised">
          <p>Sort posts by:</p>
        <div class="btn-group" style="width:100%">
          <button type="button" @click="liGen()" style="width:20%">General</button>
          <button type="button" @click="liHel()" style="width:20%">Help</button>
          <button type="button" @click="liNew()" style="width:20%">News</button>
          <button type="button" @click="liDis()" style="width:20%">Discussion</button>
          <button type="button" @click="liOff()" style="width:20%">Off-Topic</button>
        </div>
      <br><br>
      </div>

      <div v-if="$user.isAuthorised">
          <p>Sort posts by:</p>
      <div class="btn-group" style="width:100%">
          <button type="button" @click="liGen()" style="width:16.667%">General</button>
          <button type="button" @click="liHel()" style="width:16.667%">Help</button>
          <button type="button" @click="liNew()" style="width:16.667%">News</button>
          <button type="button" @click="liDis()" style="width:16.667%">Discussion</button>
          <button type="button" @click="liOff()" style="width:16.667%">Off-Topic</button>
          <button type="button" @click="liMyP()" style="width:16.667%">My-Posts</button>
        </div>
        <br><br>
        </div>

    <div id="threadList">
    <div v-if="Threads != null">
      <div v-for="Thread in Threads" :key="Thread">
        <li v-if="Thread.category==cate">
      <h1>{{ Thread.title }}</h1>
      <p>{{ Thread }}</p>
      </li>
      <li v-else-if="'My-Posts'==cate && Thread.author==auth">
        <h1>{{ Thread.title }}</h1>
      <p>{{ Thread }}</p>
      </li>
      <li v-else-if="!cate && !auth">
        <h1>{{ Thread.title }}</h1>
      <p>{{ Thread }}</p>
      </li>
    </div>
    </div>
    </div>
</div>
</template>

