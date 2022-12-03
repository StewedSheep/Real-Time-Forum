<script>
import Vue from "vue";
export default {
  name: "HomeView.vue",
  components: {},
  data (){
    return {
      Threads: [],
      cate: '',
    };
  },
  computed: {
    sortPOST(){
      return this.Threads.filter((t) => {
        let categor = this.cate
        if("My-Posts"==categor){
          return t.author === this.$user.current
        }else{return t.category.includes(categor)}
     })
    },
    },
    methods:{
    listAll(){
      this.cate = ''
    },
    listGen(){
      this.cate = "General";
    },
    listHel(){
      this.cate = "Help";
    },
    listNew(){
      this.cate = "News";
    },
    listDis(){
      this.cate = "Discussion";
    },
    listOff(){
      this.cate = "Off-Topic"
    },
    listMyP(){
      this.cate = "My-Posts"
    },
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
          <button type="button" @click="listAll()" style="width:16.66%">All</button>
          <button type="button" @click="listGen()" style="width:16.66%">General</button>
          <button type="button" @click="listHel()" style="width:16.66%">Help</button>
          <button type="button" @click="listNew()" style="width:16.66%">News</button>
          <button type="button" @click="listDis()" style="width:16.66%">Discussion</button>
          <button type="button" @click="listOff()" style="width:16.66%">Off-Topic</button>
        </div>
      <br><br>
      </div>
      <div v-if="$user.isAuthorised">
          <p>Sort posts by:</p>
      <div class="btn-group" style="width:100%">
          <button type="button" @click="listAll()" style="width:14.28%">All</button>
          <button type="button" @click="listGen()" style="width:14.28%">General</button>
          <button type="button" @click="listHel()" style="width:14.28%">Help</button>
          <button type="button" @click="listNew()" style="width:14.28%">News</button>
          <button type="button" @click="listDis()" style="width:14.28%">Discussion</button>
          <button type="button" @click="listOff()" style="width:14.28%">Off-Topic</button>
          <button type="button" @click="listMyP()" style="width:14.28%">My-Posts</button>
        </div>
        <br><br>
        </div>
    <div id="threadList">
      <div class="thread" v-for="thread in sortPOST" v-bind:key=thread>
        <p> {{ thread.category }} </p>
          <h1>{{ thread.title }}</h1>
          <p>Author:{{ thread.author }} </p>
          <p class="threadContent">{{ thread.content }}</p>
          <p class="date">Date:{{ thread.date }} </p>
          </div>
    </div>
</div>
</template>

