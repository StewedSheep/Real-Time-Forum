<template>
  <div id="app">
    <b-navbar
      style="background-color: #f5d061 !important"
      class="navBar"
      toggleable="lg"
      type="dark"
      variant="info"
    >
      <ChatBar />
      <b-navbar-nav class="ml-auto">
        <em class="loggedInAs">Logged in as: {{ $user.current }} </em>
        <router-link class="logButton" to="newThread">Create a New Thread</router-link>
        <router-link class="logButton" @click.native="logOut()" to="/"
          >Log Out</router-link
        >
      </b-navbar-nav>
    </b-navbar>
  </div>
</template>

<script>
import Vue from "vue";
import { toRef } from "vue";
import ChatBar from "./ChatBar.vue";

export default {
  data() {
    return {
      user: {
        name: "",
      },
      Users: [],
    };
  },
  methods: {
    logOut: function () {
      this.$emit("logoutEvent");
      Vue.axios.post("/logout", {
        headers: {
          "Content-Type": "text/plain",
          "Access-Control-Allow-Origin": "*",
        },
      });
      // .then((_) => {
      //   if (!this.isLoginCookieExists) {
      //     console.log(_);
      //     App.$forceUpdate();
      //   }
      // });
    },
    isLoginCookieExists() {
      return this.getLoginCookieValue() != undefined;
    },
    getLoginCookieValue() {
      const cookie = document.cookie;
      return cookie.split("::")[1];
    },
  },
  setup(props) {
    const usersRef = toRef(props, "users");
    return { usersRef };
  },
  components: {
    ChatBar: ChatBar,
  },
};
</script>
