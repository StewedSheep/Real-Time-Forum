<script>
import Vue from "vue";
export default {
    name: "HomeView.vue",
    components: {},
    data() {
        return {
            Threads: [],
        };
    },
    mounted() {
        Vue.axios.get("/totThread").then((response) => (this.Threads = response.data));
        this.cate = ""
    },
    methods: {
        sortPOST(){
          let filteredStories =  this.Threads.filter((t) => {
            let categor = this.cate
            if("My-Posts"==categor){
              return t.author === this.$user.current
            }else{return t.category.includes(categor)}
        })
          let orderedStories = filteredStories.sort((a, b) => {
            return b.id - a.id;})
            return orderedStories
        },
        clickThread(id) {
        console.log(id)
        this.$router.push({ path: '/thread', query: { id } })
        },


        liAll: function() {
            this.cate = ""
        },
        liGen: function() {
            this.cate = "General"
        },
        liHel: function() {
            this.cate = "Help"
        },
        liNew: function() {
            this.cate = "News"
        },
        liDis: function() {
            this.cate = "Discussion"
        },
        liOff: function() {
            this.cate = "Off-Topic"
        },
        liMyP: function() {
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
                <div class="thread" v-for="Thread in sortPOST()" :key="Thread">
                    <div data-id=Thread.id @click="clickThread(Thread.id)">
                  <h1>{{ Thread.title }}</h1>
                  <p>Category:{{ Thread.category }} || Author:{{ Thread.author }} || Date:{{ Thread.date }} </p>
                  <p>{{ Thread.content }}</p>
                </div>
        </div>
    </div>
</template>

