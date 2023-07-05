<template>
  <div>
    <PageHeader />
    <div v-if="this.showWS">
      <WebSocketComp />
    </div>
    <component :key="componentKey" v-bind:is="this.comp" />
    <router-view />
  </div>
</template>
<script>
import ChatBox from "./components/ChatBox.vue";
import LoggedIn from "./components/LoggedIn.vue";
import LoggedOut from "./components/LoggedOut.vue";
import PageHeader from "./components/PageHeader.vue";
import WebSocketComp from "./components/WebSocket.vue";
export default {
  data() {
    return {
      componentKey: 0,
      comp: LoggedOut,
      showWS: false,
    };
  },
  //checks for cookie and existing websocket, updates every 200ms
  created() {
    let timer = setInterval(() => {
      if (this.isLoginCookieExists()) {
        this.$user.authorised({
          username: this.getLoginCookieValue(),
        });
        this.comp = LoggedIn;
      } else {
        this.$user.authorised({
          username: undefined,
        });
        this.comp = LoggedOut;
      }
      //checks if ws connection should be open or closed
      if (
        (this.$socket == undefined && this.$user.isAuthorised) ||
        (this.$socket.readyState === WebSocket.CLOSED && this.$user.isAuthorised)
      ) {
        this.showWS = true;
        // console.log("open");
      } else if (this.$socket.readyState === WebSocket.OPEN && !this.$user.isAuthorised) {
        this.showWS = false;
        // console.log("closed");
      }
      // console.log(this.$socket);
    }, 333);
    //destroys last check hook so new can be run
    this.$once("hook:beforeDestroy", () => {
      clearInterval(timer);
    });
  },
  methods: {
    //checks cookie from browser storage
    isLoginCookieExists() {
      return this.getLoginCookieValue() != undefined;
    },
    //gets cookie name from clients browser
    getLoginCookieValue() {
      const cookie = document.cookie;
      return cookie.split("::")[1];
    },
  },

  name: "app",
  components: {
    PageHeader: PageHeader,
    LoggedIn: LoggedIn,
    LoggedOut: LoggedOut,
    ChatBox,
    WebSocketComp,
  },
};
</script>
<style>
@import "./assets/styles/main.css";
@keyframes fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
</style>
