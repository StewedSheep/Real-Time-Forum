<template>
  <div>
    <PageHeader />
    <component v-bind:is="comp" :users="users" />
    <router-view />
  </div>
</template>

<script>
import Vue from "vue";
import { reactive } from "vue";
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
      roomInput: null,
      rooms: [],
      user: {
        name: "",
      },
      users: [],
    };
  },

  //checks for cookie and existing websocket, updates every 100ms
  created: function () {
    const timer = setInterval(() => {
      this.loggedUsername = document.cookie.split("::")[1];
      //sets username to global storage
      this.$user.authorised({
        username: this.loggedUsername,
      });
      if (!this.$socket && this.$user.isAuthorised) {
        this.connectToWebsocket();
      }
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
          default:
            break;
        }
      }
    },
    handleChatMessage(msg) {
      const room = this.findRoom(msg.target.id);
      if (typeof room !== "undefined") {
        room.messages.push(msg);
      }
    },
    handleUserJoined(msg) {
      this.users.push(msg.sender);
      console.log(this.users);
    },
    handleUserLeft(msg) {
      for (let i = 0; i < this.users.length; i++) {
        if (this.users[i].id == msg.sender.id) {
          this.users.splice(i, 1);
        }
      }
    },
    handleRoomJoined(msg) {
      let room = msg.target;
      room.name = room.private ? msg.sender.name : room.name;
      room["messages"] = [];
      this.rooms.push(room);
    },
    sendMessage(room) {
      if (room.newMessage !== "") {
        this.$socket.send(
          JSON.stringify({
            action: "send-message",
            message: room.newMessage,
            target: {
              id: room.id,
              name: room.name,
            },
          })
        );
        room.newMessage = "";
      }
    },
    findRoom(roomId) {
      for (let i = 0; i < this.rooms.length; i++) {
        if (this.rooms[i].id === roomId) {
          return this.rooms[i];
        }
      }
    },
    joinRoom() {
      this.$socket.send(JSON.stringify({ action: "join-room", message: this.roomInput }));
      this.roomInput = "";
    },
    leaveRoom(room) {
      this.$socket.send(JSON.stringify({ action: "leave-room", message: room.id }));

      for (let i = 0; i < this.rooms.length; i++) {
        if (this.rooms[i].id === room.id) {
          this.rooms.splice(i, 1);
          break;
        }
      }
    },
    joinPrivateRoom(room) {
      this.$socket.send(
        JSON.stringify({ action: "join-room-private", message: room.id })
      );
    },
  },
  setup() {
    const state = reactive({
      users: [],
    });
    return { state };
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
