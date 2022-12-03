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
    liAll: function(){
      this.cate = ""
    },
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
          <button type="button" @click="liAll()" style="width:16.66%">All</button>
          <button type="button" @click="liGen()" style="width:16.66%">General</button>
          <button type="button" @click="liHel()" style="width:16.66%">Help</button>
          <button type="button" @click="liNew()" style="width:16.66%">News</button>
          <button type="button" @click="liDis()" style="width:16.66%">Discussion</button>
          <button type="button" @click="liOff()" style="width:16.66%">Off-Topic</button>
        </div>
      <br><br>
      </div>
      <div v-if="$user.isAuthorised">
          <p>Sort posts by:</p>
      <div class="btn-group" style="width:100%">
          <button type="button" @click="liAll()" style="width:14.28%">All</button>
          <button type="button" @click="liGen()" style="width:14.28%">General</button>
          <button type="button" @click="liHel()" style="width:14.28%">Help</button>
          <button type="button" @click="liNew()" style="width:14.28%">News</button>
          <button type="button" @click="liDis()" style="width:14.28%">Discussion</button>
          <button type="button" @click="liOff()" style="width:14.28%">Off-Topic</button>
          <button type="button" @click="liMyP()" style="width:14.28%">My-Posts</button>
        </div>
        <br><br>
        </div>
    <div id="threadList">
    <div v-if="Threads != null">
      <div class="thread" v-for="Thread in Threads" :key="Thread">
        <li v-if="Thread.category==cate">
          <h1>{{ Thread.title }}</h1>
          <p>Category:{{ Thread.category }} Author:{{ Thread.author }}
          Date:{{ Thread.date }} </p>
          <p>{{ Thread.content }}</p>
        </li>
        <li v-else-if="'My-Posts'==cate && Thread.author==auth">
          <h1>{{ Thread.title }}</h1>
          <p>Category:{{ Thread.category }} Author:{{ Thread.author }}
          Date:{{ Thread.date }} </p>
          <p>{{ Thread.content }}</p>
        </li>
        <li v-else-if="!cate && !auth">
          <h1>{{ Thread.title }}</h1>
          <p>Category:{{ Thread.category }} Author:{{ Thread.author }}
          Date:{{ Thread.date }} </p>
          <p>{{ Thread.content }}</p>      
        </li>
    </div>
    </div>
    </div>
</div>
</template>

