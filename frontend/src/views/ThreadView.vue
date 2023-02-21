<script>
import Vue from "vue";
export default {
    name: "HomeView.vue",
    components: {},
    data() {
        return {
            Thread: [],
        };
    },
    mounted() {
        Vue.axios.get("/thread").then((response) => (this.Threads = response.data));
        this.cate = ""
    },
    methods: {}
}
</script>

<template>
    <div id="app">
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

