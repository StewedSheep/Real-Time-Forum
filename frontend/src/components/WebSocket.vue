<template>
  <div>
    <div
      v-for="notification in notifications"
      :key="notification.id"
      class="notification"
    >
      {{ notification.message }}
    </div>
  </div>
</template>
<script>
import Vue from "vue";
import VueNativeSock from "vue-native-websocket";
import EventBus from "@/stores/event-bus.js";
import { reactive } from "vue";

export default {
  data() {
    return {
      message: "",
      messages: [],
      chatBoxes: [],
      chatopen: false,
      user: {
        name: "",
      },
      users: [],
      list: {},
      msgList: [],
      notifications: [],
      notificationId: 1,
      notification: {
        message: "",
        visible: false,
      },
      vm: "",
      $socket: undefined,
      serverUrl: "ws://localhost:8000/api/v1/ws",
    };
  },
  setup() {
    const state = reactive({
      users: [],
      list: {},
      msgList: [],
    });
    return { state };
  },
  created() {
    this.connectToWebsocket();
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
  beforeDestroy() {
    this.vm.$disconnect();
  },
  methods: {
    connectToWebsocket() {
      Vue.use(VueNativeSock, this.serverUrl + `?username=${this.$user.current}`, {
        format: "json",
        connectManually: true,
      });
      this.vm = new Vue();
      // Connect to the websocket target specified in the configuration
      this.vm.$connect();
      this.addMessegeListener();
    },
    //websocket message listener
    addMessegeListener() {
      if (this.$socket !== undefined) {
        this.$socket.onmessage = (event) => {
          const message = JSON.parse(event.data);
          console.log("from app.vue", message);
          //lists active users connected to ws
          if (message.Type === "activeUsers") {
            EventBus.$emit("users", message.content);
            //sends latest messages
          } else if (message.Type === "listMsgs" && message.to == this.$user.current) {
            EventBus.$emit("list", message.content);
            console.log(this.list);
            //listens for new chats
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
    //notification for unopened chats
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
  },
};
</script>

<style>
@import "@/assets/styles/main.css";
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
