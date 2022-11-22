<script>
import Vue from 'vue'

export default {
  name: "NewThreadView.vue",
  components: {},
  data() {
    return{
    form: {
      title: null,
      category: null,
      content: null,
      }
    }
  },
  methods : {
      createThread: function() {
        const formIsValid = !!this.form.category && !!this.form.title && !!this.form.content
        if (formIsValid){
      var data = {title: this.form.title,
                  category: this.form.category,
                  content: this.form.content,
                }
      Vue.axios.post('/threaddat', data, {
    headers: {
          'Content-Type': 'text/plain',
          'Access-Control-Allow-Origin': '*'
        }
      }).then((response) => (this.RegisterError = response.data))} else {
        console.log("form is not valid")
      }
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
            <input type="text" v-model="form.title" name="title" required /><br><br>
                <br>
                <label>Choose a category:</label>
                <v-select v-model="form.category" :options="['General', 'Help' ,'News', 'Discussion', 'Off-Topic']" />
            </td>
        </tr>
        <br>  
        <textarea v-model="form.content" required></textarea>
    <br>
    <p v-if="!formIsValid" class="error-message">*All the fields are required.</p>
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