<script>
import Vue from 'vue'

export default {
  name: "NewThreadView.vue",
  components: {},
  methods : {
      createThread: function() {
      var data = {title: this.title,
                  category: this.category,
                  content: this.content,
                }
      Vue.axios.post('/threaddat', data, {
    headers: {
          'Content-Type': 'text/plain',
          'Access-Control-Allow-Origin': '*'
        }
      }).then((response) => (this.RegisterError = response.data))
    },
  }
}

</script>


<template>
  <div v-if="$user.isAuthorised">
    <br>
        <form>
        <tr><br><td>
                <label>Thread Topic:</label><br>
            <input type="text" v-model="title" name="title" required /><br><br>
                <br>
                <label>Choose a category:</label>
                <v-select v-model="category" :options="['General', 'Help' ,'News', 'Discussion', 'Off-Topic']"></v-select>
            </td>
        </tr>
        <br>  
        <textarea v-model="content" required></textarea>
    <br>
    <button type="button" @click="createThread()">Post new Thread</button> 
    </form>
</div>
<div v-else id="app">
    <h3>Login to create a new thread.</h3>
</div>
</template>

<style scoped>
textarea {
  width: 50%;
  height: 150px;
}
</style>