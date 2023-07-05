<script>
import Vue from "vue";

var today = new Date();
var date = today.getFullYear() + "-" + (today.getMonth() + 1) + "-" + today.getDate();
var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
var dateTime = date + " " + time;

export default {
  name: "NewThreadView.vue",
  components: {},
  data() {
    return {
      form: {
        title: null,
        category: null,
        content: null,
      },
    };
  },
  methods: {
    createThread: function () {
      const formIsValid =
        !!this.form.category && !!this.form.title && !!this.form.content;
      if (formIsValid) {
        var data = {
          title: this.form.title,
          author: this.$user.current,
          date: dateTime,
          category: this.form.category,
          content: this.form.content,
        };
        Vue.axios
          .put("/threaddat", data, {
            headers: {
              "Content-Type": "text/plain",
              "Access-Control-Allow-Origin": "*",
            },
          })
          .then(this.$router.push("/"));
      } else {
        console.log("form is not valid");
      }
    },
  },
};
</script>

<template>
  <div v-if="$user.isAuthorised">
    <br />
    <form class="thread">
      <tr>
        <br />
        <td>
          <label>Thread Topic:</label><br />
          <input
            class="logButton"
            type="text"
            v-model="form.title"
            name="title"
            required
          /><br /><br />
          <br />
          <label>Choose a category:</label>
          <!-- <v-select class="logButton" v-model="form.category" :options="['General', 'Help' ,'News', 'Discussion', 'Off-Topic']" /> -->
          <select class="logButton" v-model="form.category">
            <option selected="selected">General</option>
            <option>Help</option>
            <option>News</option>
            <option>Discussion</option>
            <option>Off-Topic</option>
          </select>
        </td>
      </tr>
      <br />
      <textarea class="logButton" v-model="form.content" required></textarea>
      <br />
      <p v-if="!formIsValid" class="error-message">*All the fields are required.</p>
      <button class="logButton" type="button" @click="createThread">
        Post new Thread
      </button>
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
