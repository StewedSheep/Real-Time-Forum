<template>
  <div>
    <PageHeader />
    <component v-bind:is="comp" />
    <router-view />
  </div>
</template>

<script>
import Vue from "vue";
import LoggedIn from "./components/LoggedIn.vue";
import LoggedOut from "./components/LoggedOut.vue";
import PageHeader from "./components/PageHeader.vue";

let loggedUsername;

export default {
  data() {
    return {
      loggedUsername,
      componentKey: 0,
      comp: LoggedOut,
      socket: null,
      serverUrl: "ws://localhost:8000/api/v1/ws",
      messages: [],
      newMessage: "",
    };
  },
  // connects to websocket on loading page
  updated: function () {
    setInterval(() => {
      if (!this.$socket) {
        this.connectToWebsocket();
      }
    }, 5000);
  },
  //checks for cookie updates every 100ms
  created: function () {
    const timer = setInterval(() => {
      this.loggedUsername = document.cookie.split("::")[1];
      //sets username to global storage
      this.$user.authorised({
        username: this.loggedUsername,
      });
    }, 100);
    //destroys last check hook so new can be run
    this.$once("hook:beforeDestroy", () => {
      clearInterval(timer);
    });
  },

  methods: {
    connectToWebsocket() {
      const socket = new WebSocket(this.serverUrl + "?name=" + this.$user.current);
      Vue.prototype.$socket = socket;
      this.$socket.addEventListener("open", (event) => {
        this.onWebsocketOpen(event);
      });
    },
    onWebsocketOpen() {
      console.log("connected to WS!");
    },
  },

  //watches for changes in logged in username and changes components
  watch: {
    loggedUsername(newVal, oldVal) {
      if (newVal != "") {
        this.comp = LoggedIn;
      }
      if (newVal == null) {
        this.comp = LoggedOut;
      }
      console.log(newVal);
      if (oldVal != newVal) {
        this.componentKey++;
      }
      console.log("loggedUsername watcher triggered");
    },
    immediate: true,
  },
  name: "app",
  components: {
    PageHeader: PageHeader,
    LoggedIn: LoggedIn,
    LoggedOut: LoggedOut,
  },
};
</script>

<style>
@import "./assets/styles/main.css";
</style>
