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
            <P>By: {{ Thread.author }}</P>
            <p>Category: {{ Thread.category }}</p>
            <p class="threadContent">{{ Thread.content }}</p>
            <p class="date">Date: {{ Thread.date }}</p>
          
            <div class="comments">
            <h1>Comment</h1>
            <p>By: urmom</p>
            <P class="threadContent">Haha I'm a stupid comment hahah</P>
            <p class="date">Date: Wanna Go out?</p>
          </div>
          </div>
        </div>
    </div>
</template>

