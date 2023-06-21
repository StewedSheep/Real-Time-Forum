<template>
  <div>
    <PageHeader />
    <div class="notification" v-if="notification.visible">
      {{ notification.message }}
    </div>
    <component v-bind:is="comp" :users="users" />
    <router-view />
  </div>
</template>

<script>
import Vue from "vue";
import VueNativeSock from "vue-native-websocket";
//import EventBus from "@/stores/event-bus.js";
import { reactive } from "vue";
import ChatBox from "./components/ChatBox.vue";
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
      newMessage: "",
      //websocket data here on out
      recipient: "",
      message: "",
      messages: [],
      chatBoxes: [],
      roomInput: null,
      user: {
        name: "",
      },
      users: [],
      notification: {
        message: "",
        visible: false,
      },
    };
  },

  //checks for cookie and existing websocket, updates every 100ms
  created() {
    let timer = setInterval(() => {
      this.loggedUsername = document.cookie.split("::")[1];
      //sets username to global storage
      this.$user.authorised({
        username: this.loggedUsername,
      });
      //checks if ws connection should be open or closed
      if (
        (this.$socket == null && this.$user.isAuthorised) ||
        (!this.$socket.readyState === WebSocket.OPEN && this.$user.isAuthorised)
      ) {
        this.connectToWebsocket();
        this.addMessegeListener();
      } else if (this.$socket.readyState === WebSocket.OPEN && !this.$user.isAuthorised) {
        this.$socket.close();
        this.$socket = undefined;
      }
    }, 200);
    //destroys last check hook so new can be run
    this.$once("hook:beforeDestroy", () => {
      clearInterval(timer);
    });
  },

  beforeDestroy() {
    if (this.socket) {
      this.socket.close();
    }
  },

  mounted() {
    // EventBus.$on("chatbox-opened", (eventData) => {
    //   for (let i = 0; i < this.users.length; i++) {
    //     if (this.users[i].name == eventData) {
    //       console.log("attempting to open chat with ", this.users[i].name);
    //       this.joinRoom(this.users[i]);
    //     }
    //   }
    // });
  },

  sockets: {
    onmessage(data) {
      console.log("recived, ", data);
    },
  },

  methods: {
    connectToWebsocket() {
      Vue.use(VueNativeSock, this.serverUrl + `?username=${this.$user.current}`, {
        format: "json",
        reconnection: true,
        reconnectionAttempts: 5,
        reconnectionDelay: 3000,
      });
    },
    addMessegeListener() {
      this.$socket.onmessage = (event) => {
        const message = JSON.parse(event.data);
        if (message.Type === "activeUsers") {
          this.users = message.content;
          console.log(this.users);
        }
      };
    },
    showNotification(message) {
      this.notification.message = message;
      this.notification.visible = true;

      // Hide the notification after a certain duration (4 seconds)
      setTimeout(() => {
        this.notification.visible = false;
      }, 4000);
    },
  },
  setup() {
    const state = reactive({
      users: [],
    });
    return { state };
  },

  //watches for changes in logged-in username
  watch: {
    loggedUsername(newVal, oldVal) {
      if (newVal != "") {
        this.comp = LoggedIn;
      }
      if (newVal == null) {
        this.comp = LoggedOut;
      }
      if (oldVal != newVal) {
        this.componentKey++;
      }
    },
    immediate: true,
  },
  name: "app",
  components: {
    PageHeader: PageHeader,
    LoggedIn: LoggedIn,
    LoggedOut: LoggedOut,
    ChatBox,
  },
};
</script>

<style>
@import "./assets/styles/main.css";
.notification {
  position: fixed;
  top: 20px;
  right: 20px;
  padding: 10px;
  background-color: #f5f5f5;
  border: 1px solid #ccc;
  border-radius: 4px;
  animation: fade-in 0.3s ease-in-out;
}

@keyframes fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
</style>
