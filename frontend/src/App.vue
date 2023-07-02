<template>
  <div>
    <PageHeader />
    <div
      v-for="notification in notifications"
      :key="notification.id"
      class="notification"
    >
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
      notifications: [],
      notificationId: 1,
      //websocket data here on out
      recipient: "",
      message: "",
      messages: [],
      chatBoxes: [],
      roomInput: null,
      chatopen: false,
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

  // sockets: {
  //   onmessage(data) {
  //     console.log("recived, ", data);
  //   },
  // },

  methods: {
    connectToWebsocket() {
      Vue.use(VueNativeSock, this.serverUrl + `?username=${this.$user.current}`, {
        format: "json",
        // reconnection: true,
        // reconnectionAttempts: 5,
        // reconnectionDelay: 3000,
      });
      console.log("connected to ws!");
    },
    closeWebSocket() {
      this.$socket.close();
    },

    addMessegeListener() {
      if (this.$socket !== undefined) {
        this.$socket.onmessage = (event) => {
          const message = JSON.parse(event.data);
          console.log("from app.vue", message);
          if (message.Type === "activeUsers") {
            this.users = message.content;
          } else if (message.Type === "listMsgs" && message.to == this.$user.current) {
            this.list = message.content;
            console.log(this.list);
          } else if (message.Type == "chat") {
            this.chatopen = false;
            for (const key in this.chatBoxes) {
              const chatBox = this.chatBoxes[key];
              console.log(chatBox.title);
              if (message.from == chatBox.title && chatBox.visible == true) {
                EventBus.$emit("chatbox-data", message);
                this.chatopen = true;
              }
            }
            if (this.chatopen == false) {
              this.showNotification(message.from);
              console.log("notification fired");
            }
          }
        };
      }
    },
    showNotification(message) {
      const notification = {
        id: this.notificationId++,
        message: "New message from: " + message,
      };

      this.notifications.push(notification);

      setTimeout(() => {
        this.removeNotification(notification.id);
      }, 5000);
    },
    removeNotification(id) {
      this.notifications = this.notifications.filter(
        (notification) => notification.id !== id
      );
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
  left: 20px;
  padding: 10px;
  background-color: #282f44;
  border: 1px solid #e6af2e;
  border-radius: 4px;
  animation: fade-in 0.3s ease-in-out;
  padding: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  z-index: 9999;
  margin-bottom: 10px;
}

.notification:first-child {
  margin-top: 0;
}

.notification + .notification {
  margin-top: 50px;
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
