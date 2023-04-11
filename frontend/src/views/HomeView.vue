<script>
import Vue from "vue";
import { ref } from 'vue';
export default {
    name: "HomeView.vue",
    components: {},
    setup() {
        const selected = ref(0) // index of the selected el
        const changeSelected = (i) => { selected.value = i; };
        return {
            changeSelected, selected,
        }
    },
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
                <button type="button" :class=' { "selected": selected === 0 } ' @click="liAll(), changeSelected(0)" style="width:16.66%">All</button>
                <button type="button" :class=' { "selected": selected === 1 } ' @click="liGen(), changeSelected(1)" style="width:16.66%">General</button>
                <button type="button" :class=' { "selected": selected === 2 } ' @click="liHel(), changeSelected(2)" style="width:16.66%">Help</button>
                <button type="button" :class=' { "selected": selected === 3 } ' @click="liNew(), changeSelected(3)" style="width:16.66%">News</button>
                <button type="button" :class=' { "selected": selected === 4 } ' @click="liDis(), changeSelected(4)" style="width:16.66%">Discussion</button>
                <button type="button" :class=' { "selected": selected === 5 } ' @click="liOff(), changeSelected(5)" style="width:16.66%">Off-Topic</button>
            </div>
            <br><br>
        </div>
        <div v-if="$user.isAuthorised">
            <p>Sort posts by:</p>
            <div class="btn-group" style="width:100%">
                <button type="button" :class=' { "selected": selected === 0 } ' @click="liAll(), changeSelected(0)" style="width:14.28%">All</button>
                <button type="button" :class=' { "selected": selected === 1 } ' @click="liGen(), changeSelected(1)" style="width:14.28%">General</button>
                <button type="button" :class=' { "selected": selected === 2 } ' @click="liHel(), changeSelected(2)" style="width:14.28%">Help</button>
                <button type="button" :class=' { "selected": selected === 3 } ' @click="liNew(), changeSelected(3)" style="width:14.28%">News</button>
                <button type="button" :class=' { "selected": selected === 4 } ' @click="liDis(), changeSelected(4)" style="width:14.28%">Discussion</button>
                <button type="button" :class=' { "selected": selected === 5 } ' @click="liOff(), changeSelected(5)" style="width:14.28%">Off-Topic</button>
                <button type="button" :class=' { "selected": selected === 6 } ' @click="liMyP(), changeSelected(6)" style="width:14.28%">My-Posts</button>
            </div>
            <br><br>
        </div>
        <div id="threadList">
                <div class="thread" v-for="(Thread, idx,) in sortPOST()" :key="idx">
                    <div data-id=Thread.id @click="clickThread(Thread.id)">
                  <h1>{{ Thread.title }}</h1>
                  <p>Category:{{ Thread.category }} || Author:{{ Thread.author }} || Date:{{ Thread.date }} </p>
                  <p>{{ Thread.content }}</p>
                </div>
        </div>
    </div>
    </div>
</template>
<style scoped>
.selected {
    background-color: #E6AF2E !important;
}
.selected:hover {
    background-color: #a8861d!important;
}
</style>