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
// import Vue from "vue";
// import VueNativeSock from "vue-native-websocket";
// import EventBus from "@/stores/event-bus.js";
// import { reactive } from "vue";
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
      // $socket: undefined,
      // serverUrl: "ws://localhost:8000/api/v1/ws",
      // notifications: [],
      // notificationId: 1,

      //websocket data here on out

      showWS: false,
      // message: "",
      // messages: [],
      // chatBoxes: [],
      // chatopen: false,
      // user: {
      //   name: "",
      // },
      // users: [],
      // list: {},
      // msgList: [],
      // notification: {
      //   message: "",
      //   visible: false,
      // },
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
        console.log("open");
      } else if (this.$socket.readyState === WebSocket.OPEN && !this.$user.isAuthorised) {
        this.showWS = false;
        console.log("closed");
      }
      console.log(this.$socket);
    }, 333);
    //destroys last check hook so new can be run
    this.$once("hook:beforeDestroy", () => {
      clearInterval(timer);
    });
  },
  // beforeDestroy() {
  //   if (this.$socket) {
  //     this.closeWebSocket();
  //   }
  // },
  // mounted() {
  //   EventBus.$on("chatbox-opened", (eventData) => {
  //     console.log("opened chat, total chats: ", eventData);
  //     this.chatBoxes = eventData;
  //   });
  //   EventBus.$on("chatbox-closed", (eventData) => {
  //     console.log("closed chat, remaining chats: ", eventData);
  //   });
  // },
  methods: {
    // openWS() {
    //   this.showWS = true;
    // },
    // closeWS() {
    //   this.showWS = false;
    // },
    // connectToWebsocket() {
    //   if (this.$socket) {
    //     this.$socket.close();
    //     this.$socket = undefined;
    //   }
    //   Vue.use(VueNativeSock, this.serverUrl + `?username=${this.$user.current}`, {
    //     format: "json",
    //   });
    //   this.addMessegeListener();
    // },
    // closeWebSocketComp() {
    //   if (this.$socket) {
    //     this.$socket.close();
    //     this.$socket = undefined;
    //   }
    // },
    // //websocket message listener
    // addMessegeListener() {
    //   if (this.$socket !== undefined) {
    //     this.$socket.onmessage = (event) => {
    //       const message = JSON.parse(event.data);
    //       console.log("from app.vue", message);
    //       //lists active users connected to ws
    //       if (message.Type === "activeUsers") {
    //         this.users = message.content;
    //         //sends latest messages
    //       } else if (message.Type === "listMsgs" && message.to == this.$user.current) {
    //         this.list = message.content;
    //         console.log(this.list);
    //         //listens for new chats
    //       } else if (message.Type == "chat") {
    //         this.chatopen = false;
    //         for (const key in this.chatBoxes) {
    //           const chatBox = this.chatBoxes[key];
    //           console.log(chatBox.title);
    //           if (message.from == chatBox.title && chatBox.visible == true) {
    //             EventBus.$emit("chatbox-data", message);
    //             this.chatopen = true;
    //           }
    //         }
    //         if (this.chatopen == false) {
    //           this.showNotification(message.from);
    //           console.log("notification fired");
    //         }
    //       }
    //     };
    //   }
    // },
    // //notification for unopened chats
    // showNotification(message) {
    //   const notification = {
    //     id: this.notificationId++,
    //     message: "New message from: " + message,
    //   };
    //   this.notifications.push(notification);
    //   setTimeout(() => {
    //     this.removeNotification(notification.id);
    //   }, 5000);
    // },
    // removeNotification(id) {
    //   this.notifications = this.notifications.filter(
    //     (notification) => notification.id !== id
    //   );
    // },
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
  // setup() {
  //   const state = reactive({
  //     users: [],
  //     list: {},
  //     msgList: [],
  //   });
  //   return { state };
  // },
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
