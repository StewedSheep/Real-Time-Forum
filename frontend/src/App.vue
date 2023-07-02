<template>
  <div>
    <PageHeader />
    <div class="notification" v-if="notification.visible">
      {{ notification.message }}
    </div>
    <component :key="componentKey" v-bind:is="this.comp" :users="users" :list="list" />
    <router-view />
  </div>
</template>

<script>
import Vue from "vue";
import VueNativeSock from "vue-native-websocket";
import EventBus from "@/stores/event-bus.js";
import { reactive } from "vue";
import ChatBox from "./components/ChatBox.vue";
import LoggedIn from "./components/LoggedIn.vue";
import LoggedOut from "./components/LoggedOut.vue";
import PageHeader from "./components/PageHeader.vue";

export default {
  data() {
    return {
      componentKey: 0,
      comp: LoggedOut,
      $socket: undefined,
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
      list: {},
      msgList: [],
      notification: {
        message: "",
        visible: false,
      },
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
        (!this.$socket.readyState === WebSocket.OPEN && this.$user.isAuthorised)
      ) {
        this.connectToWebsocket();
        this.addMessegeListener();
      } else if (this.$socket.readyState === WebSocket.OPEN && !this.$user.isAuthorised) {
        this.closeWebSocket();
        this.$socket = undefined;
      }
    }, 200);
    //destroys last check hook so new can be run
    this.$once("hook:beforeDestroy", () => {
      clearInterval(timer);
    });
  },

  beforeDestroy() {
    if (this.$socket) {
      this.$socket.close();
    }
  },

  mounted() {
    EventBus.$on("chatbox-opened", (eventData) => {
      console.log("opened chat, total chats: ", eventData);
      this.chatBoxes = eventData;
    });
    EventBus.$on("chatbox-closed", (eventData) => {
      console.log("closed chat, remaining chats: ", eventData);
    });
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
        // reconnection: true,
        // reconnectionAttempts: 5,
        // reconnectionDelay: 3000,
      });
    },
    closeWebSocket() {
      this.$socket.close();
    },

    addMessegeListener() {
      if (this.$socket) {
        this.$socket.onmessage = (event) => {
          const message = JSON.parse(event.data);
          console.log("from app.vue", message);
          if (message.Type === "activeUsers") {
            this.users = message.content;
          } else if (message.Type === "listMsgs" && message.to == this.$user.current) {
            this.list = message.content;
            console.log(this.list);
            // } else if (message.Type !== "activeUsers" && message.Type !== "listMsgs") {
            // for (const key in this.chatBoxes) {
            //   const chatBox = this.chatBoxes[key];
            //   console.log(chatBox.title);
            //   if (message.from == chatBox.title) {
            //     this.chatBoxes[message.from].messages.push(message);
            //   }
            // }
          }
        };
      }
    },
    showNotification(message) {
      this.notification.message = message;
      this.notification.visible = true;

      // Hide the notification after a certain duration (4 seconds)
      setTimeout(() => {
        this.notification.visible = false;
      }, 4000);
    },
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
  setup() {
    const state = reactive({
      users: [],
      list: {},
      msgList: [],
    });
    return { state };
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
