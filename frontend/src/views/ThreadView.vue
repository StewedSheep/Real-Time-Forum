<script>
import Vue from "vue";
export default {
    name: "ThreadView.vue",
    components: {},
    data() {
        return {
            Threads: [],
        };
    },
    mounted() {
        console.log(this.$route.query.id)
        var data = {threadId: this.$route.query.id,}
        console.log(data)
 
      Vue.axios.post('/thread', data, {
        headers: {
          'Content-Type': 'text/plain',
          'Access-Control-Allow-Origin': '*'
        }
        }).then((response) => (this.Threads = response.data));
    
    },
     methods: {}
}
</script>

<template>
    <div id="app">
        <div id="threadList">
                <div class="thread" v-for="Thread in Threads" :key="Thread">
                  <h1>{{ Thread.title }}</h1>
                  <p>Category:{{ Thread.category }} || Author:{{ Thread.author }} || Date:{{ Thread.date }} </p>
                  <p>{{ Thread.content }}</p>
                </div>
        </div>
    </div>
</template>

