<script>
import Vue from "vue";
var today = new Date();
var date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();
var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
var dateTime = date+' '+time;

export default {
    name: "ThreadView.vue",
    components: {},
    data() {
        return {
            Thread: [],
            form: {
              relId: null,
              author: null,
              content: null,
            },
        };
    },
    mounted() {
        //console.log("ThreadId on ThreadView:", this.$route.query.id)
        var data = {threadId: this.$route.query.id,}
        console.log("ThreadView thread id data:", data)
      Vue.axios.post('/thread', data, {
        headers: {
          'Content-Type': 'text/plain',
          'Access-Control-Allow-Origin': '*'
        }
        }).then((response) => (
          //console.log("Comments:", response.data.comments),
          this.Thread = response.data));
    
    },
     methods: {  createComment: function() {
        const formIsValid = !!this.form.content
        if (formIsValid){
          var data = {author: this.$user.current,
            date: dateTime,
            content: this.form.content,
            relId: this.$route.query.id,
          }
          console.log("Comment data:", data)
          Vue.axios.put('/commentdat', data, {
            headers: {
              'Content-Type': 'text/plain',
              'Access-Control-Allow-Origin': '*'
            }
          })} else {console.log("form is not valid")}
      },
    }
}
</script>

<template>
    <div id="app">
        <div id="threadList">
          <div class="thread" v-for="(Thread, idx) in Thread" :key="idx">
            <h1>{{ Thread.title }}</h1>
            <P>By: {{ Thread.author }}</P>
            <p>Category: {{ Thread.category }}</p>
            <p class="threadContent">{{ Thread.content }}</p>
            <p class="date">Date: {{ Thread.date }}</p>
            <form class="comments">
              <label for="comment">Add a comment:</label>
              <input type="text" v-model="form.content" id="comment" name="comment"><br><br>
              <input class="SubmitButton" type="Button" @click="createComment()" value="Submit">
            </form>
            <br>
            <div class="commentList">
              <div class="comments" v-for="(Comment, idx) in Thread.comments" :key="idx">
                <h1>{{Comment.author}}</h1>
                <p class="threadContent">{{ Comment.content }}</p>
                <p class="date">{{ Comment.date }}</p>
              </div>
            </div>
          </div>
        </div>
    </div>
</template>

