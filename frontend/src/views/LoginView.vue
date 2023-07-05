<script>
import Vue from "vue";

export default {
  name: "LoginView.vue",
  components: {},
  data() {
    return {
      username: "",
      password: "",
      LoginError: "",
    };
  },
  methods: {
    postLog: function () {
      var data = {
        username: this.username,
        password: this.password,
        LoginError: this.LoginError,
      };
      console.log(data);
      Vue.axios
        .post("/login", data, {
          headers: {
            "Content-Type": "text/plain",
            "Access-Control-Allow-Origin": "*",
          },
        })
        .then((response) => (this.LoginError = response.data));
      if (this.LoginError !== "Check login info.") {
        this.$router.push("/");
      }
    },
  },
};
</script>

<template>
  <div v-if="$user.isAuthorised">
    <h3 class="thread" style="text-align: center">
      Already Logged in. Please return to the home page.
    </h3>
  </div>
  <div v-else id="app">
    <center>
      <h3>Login here</h3>
      <div v-if="LoginError != null">
        <p>{{ LoginError }}</p>
        <br />
      </div>
      <form>
        <label>Username or Email:</label><br />
        <input
          class="logButton"
          type="text"
          v-model="username"
          name="username"
        /><br /><br />
        <label for="password">Password:</label><br />
        <input
          class="logButton"
          type="password"
          v-model="password"
          name="password"
        /><br /><br />
        <br />
      </form>
      <div>
        <button class="logButton" type="button" @click="postLog()">Login</button>
        <br /><br />
        <router-link to="register"
          ><img
            class="noAccount"
            src="@/assets/no_account.jpg"
            alt="Click here to register!"
        /></router-link>
      </div>
      <router-view />
    </center>
  </div>
</template>
