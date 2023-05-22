<template>
  <div>
    <PageHeader />
    <component v-bind:is="comp" @logoutEvent="removeUser" :users="users" />
    <router-view />
    <!-- message box containers -->
    <!-- <div class="chat-box-container">
      <div
        v-for="(chatBox, index) in chatBoxes"
        :key="index"
        class="chat-box"
        :class="{ visible: chatBox.visible }"
      >
        <ChatBox
          :title="chatBox.title"
          :visible="chatBox.visible"
          @close="closeChatBox(index)"
        />
      </div>
    </div> -->
  </div>
</template>

<script>
import EventBus from "@/stores/event-bus.js";
import Vue from "vue";
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
      messages: [],
      newMessage: "",
      //websocket data here on out
      chatBoxes: [],
      roomInput: null,
      user: {
        name: "",
      },
      users: [],
    };
  },

  //checks for cookie and existing websocket, updates every 100ms
  created: function () {
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

  mounted() {
    EventBus.$on("chatbox-opened", (eventData) => {
      for (let i = 0; i < this.users.length; i++) {
        if (this.users[i].name == eventData) {
          console.log("attempting to open chat with ", this.users[i]);
          this.joinRoom(this.users[i]);
        }
      }
    });
  },

  methods: {
    connectToWebsocket() {
      const socket = new WebSocket(this.serverUrl + "?name=" + this.$user.current);
      Vue.prototype.$socket = socket;
      if (this.$socket == undefined) {
        this.$socket = socket;
      }
      this.$socket.addEventListener("open", (event) => {
        this.onWebsocketOpen(event);
        if (this.$socket) {
          console.log("event listener open");
          this.$socket.addEventListener("message", (event) => {
            this.handleNewMessage(event);
          });
        }
      });
    },
    onWebsocketOpen() {
      console.log("connected to WS!");
    },
    handleNewMessage(event) {
      let data = event.data;
      data = data.split(/\r?\n/);

      for (let i = 0; i < data.length; i++) {
        let msg = JSON.parse(data[i]);
        switch (msg.action) {
          case "send-message":
            this.handleChatMessage(msg);
            break;
          case "user-join":
            this.handleUserJoined(msg);
            break;
          case "user-left":
            this.handleUserLeft(msg);
            break;
          case "room-joined":
            this.handleRoomJoined(msg);
            break;
          case "join-room-private":
            this.joinPrivateRoom(msg);
            break;
          default:
            break;
        }
      }
    },
    handleChatMessage(msg) {
      const chatBox = this.findRoom(msg.target.id);
      if (typeof chatBox !== "undefined") {
        chatBox.messages.push(msg);
      }
    },
    handleUserJoined(msg) {
      this.users.push(msg.sender);
    },
    handleUserLeft(msg) {
      for (let i = 0; i < this.users.length; i++) {
        if (this.users[i].id == msg.sender.id) {
          this.users.splice(i, 1);
        }
      }
    },
    handleRoomJoined(msg) {
      let chatBox = msg.target;
      chatBox.name = chatBox.private ? msg.sender.name : chatBox.name;
      chatBox["messages"] = [];
      this.chatBoxes.push(chatBox);
    },
    sendMessage(chatBox) {
      if (chatBox.newMessage !== "") {
        this.$socket.send(
          JSON.stringify({
            action: "send-message",
            message: this.newMessage,
            target: {
              id: chatBox.id,
              name: chatBox.name,
            },
          })
        );
        chatBox.newMessage = "";
      }
    },
    findRoom(roomId) {
      for (let i = 0; i < this.chatBoxes.length; i++) {
        if (this.chatBoxes[i].id === roomId) {
          return this.chatBoxes[i];
        }
      }
    },
    joinRoom() {
      this.$socket.send(JSON.stringify({ action: "join-room", message: this.roomInput }));
      this.roomInput = "";
    },
    leaveRoom(chatBox) {
      this.$socket.send(JSON.stringify({ action: "leave-room", message: chatBox.id }));

      for (let i = 0; i < this.chatBoxes.length; i++) {
        if (this.chatBoxes[i].id === chatBox.id) {
          this.chatBoxes.splice(i, 1);
          break;
        }
      }
    },
    joinPrivateRoom(chatBox) {
      this.$socket.send(
        JSON.stringify({ action: "join-room-private", message: chatBox.id })
      );
    },
    removeUser() {
      this.users = [];
    },
    closeChatBox(index) {
      this.chatBoxes.splice(index, 1);
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
      console.log(newVal);
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
</style>
