<template>
  <html>
    <head>
      <title>Chat</title>
      <link type="text/css" rel="stylesheet" href="assets/style.css" />
    </head>

    <body>
      <div id="app">
        <div class="container h-100">
          <div class="row justify-content-center h-100">
            <div class="col-12">
              <div class="row">
                <div class="col-2 card profile" v-for="user in users" :key="user">
                  <div class="card-header">{{ user.name }}</div>
                  <div class="card-body">
                    <button class="btn btn-primary" @click="joinPrivateRoom(user)">
                      Send Message
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div class="chat" v-for="(room, key) in rooms" :key="key">
              <div class="card">
                <div class="card-header msg_head">
                  <div class="d-flex bd-highlight justify-content-center">
                    {{ room.name }}
                    <span class="card-close" @click="leaveRoom(room)">leave</span>
                  </div>
                </div>
                <div class="card-body msg_card_body">
                  <div
                    v-for="(message, key) in room.messages"
                    :key="key"
                    class="d-flex justify-content-start mb-4"
                  >
                    <div class="msg_cotainer">
                      {{ message.message }}
                      <span class="msg_name" v-if="message.sender">{{
                        message.sender.name
                      }}</span>
                    </div>
                  </div>
                </div>
                <div class="card-footer">
                  <div class="input-group">
                    <textarea
                      v-model="room.newMessage"
                      name=""
                      class="form-control type_msg"
                      placeholder="Type your message..."
                      @keyup.enter.exact="sendMessage(room)"
                    ></textarea>
                    <div class="input-group-append">
                      <span class="input-group-text send_btn" @click="sendMessage(room)"
                        >></span
                      >
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </body>
  </html>
</template>
<script>
export default {
  props: ["username"],
  data() {
    return {
      roomInput: null,
      rooms: [],
      user: {
        name: "",
      },
      users: [],
    };
  },
  mounted: function () {
    this.connectToWebsocket();
  },
  methods: {
    connectToWebsocket() {
      if (this.$socket) {
        console.log("event listener open");
        this.$socket.addEventListener("message", (event) => {
          this.handleNewMessage(event);
        });
      }
    },
  },
};

  //   handleNewMessage(event) {
  //     let data = event.data;
  //     data = data.split(/\r?\n/);

  //     for (let i = 0; i < data.length; i++) {
  //       let msg = JSON.parse(data[i]);
  //       switch (msg.action) {
  //         case "send-message":
  //           this.handleChatMessage(msg);
  //           break;
  //         case "user-join":
  //           this.handleUserJoined(msg);
  //           break;
  //         case "user-left":
  //           this.handleUserLeft(msg);
  //           break;
  //         case "room-joined":
  //           this.handleRoomJoined(msg);
  //           break;
  //         default:
  //           break;
  //       }
  //     }
  //   },
  //   handleChatMessage(msg) {
  //     const room = this.findRoom(msg.target.id);
  //     if (typeof room !== "undefined") {
  //       room.messages.push(msg);
  //     }
  //   },
  //   handleUserJoined(msg) {
  //     console.log("msg.sender");
  //     this.users.push(msg.sender);
  //   },
  //   handleUserLeft(msg) {
  //     for (let i = 0; i < this.users.length; i++) {
  //       if (this.users[i].id == msg.sender.id) {
  //         this.users.splice(i, 1);
  //       }
  //     }
  //   },
  //   handleRoomJoined(msg) {
  //     let room = msg.target;
  //     room.name = room.private ? msg.sender.name : room.name;
  //     room["messages"] = [];
  //     this.rooms.push(room);
  //   },
  //   sendMessage(room) {
  //     if (room.newMessage !== "") {
  //       this.$socket.send(
  //         JSON.stringify({
  //           action: "send-message",
  //           message: room.newMessage,
  //           target: {
  //             id: room.id,
  //             name: room.name,
  //           },
  //         })
  //       );
  //       room.newMessage = "";
  //     }
  //   },
  //   findRoom(roomId) {
  //     for (let i = 0; i < this.rooms.length; i++) {
  //       if (this.rooms[i].id === roomId) {
  //         return this.rooms[i];
  //       }
  //     }
  //   },
  //   joinRoom() {
  //     this.$socket.send(JSON.stringify({ action: "join-room", message: this.roomInput }));
  //     this.roomInput = "";
  //   },
  //   leaveRoom(room) {
  //     this.$socket.send(JSON.stringify({ action: "leave-room", message: room.id }));

  //     for (let i = 0; i < this.rooms.length; i++) {
  //       if (this.rooms[i].id === room.id) {
  //         this.rooms.splice(i, 1);
  //         break;
  //       }
  //     }
  //   },
  //   joinPrivateRoom(room) {
  //     this.$socket.send(
  //       JSON.stringify({ action: "join-room-private", message: room.id })
  //     );
  //   },
  // },
// };
</script>
