<template>
  <div v-if="visible" class="chatbox">
    <div class="header">
      <h3>{{ title }}</h3>
      <button class="logButton" @click="close()">&times;</button>
    </div>
    <div class="content">
      <!-- Chat content goes here -->
      <div v-for="chatBox in chatBoxes" :key="chatBox.id">
         <!-- Only display messages when the chatbox is open -->
        <div v-if="chatBox.id === targetId">
          <div
            v-for="(message, index) in chatBox.messages" :key="index"
            v-show="message.recip == title"
            :class="{'sent-message': message.author !== title, 'received-message': message.author == title}"
            class="message"
          >
            {{ message.author }}: {{ message.message }}
          </div>
        </div>
      </div>
    </div>
    <div class="card-footer">
      <input
        v-model="this.newMessage"
        class="messageTextArea"
        name=""
        placeholder="Type your message..."
        @keyup.enter.exact="sendMessage(chatBox)"
      >
      <button 
        class="logButton input-group-append" 
        @click="sendMessage(chatBox)"
      >Send</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: "ChatBox",
  data() {
    return {
      roomInput: null,
      chatBoxes: [],
      chatBox: {
        id: "",
        messages: [],
        name: "",
        private: true,
      },
      message: {
        author: "",
        recip: "",
        message: "",
        date: "",
      },
      newMessage: "",
      users: [],
      user: {
        id: "",
        name: "",
      },
      targetId: "",
    };
  },
  props: {
    title: {
      type: String,
      default: "Chat Box",
    },
    visible: {
      type: Boolean,
      default: true,
    },
  },
  mounted: function () {
    this.connectToWebsocket();
  },
  methods: {
    connectToWebsocket() {
      if (this.$socket) {
        console.log("event listener open");
        this.$socket.addEventListener("message", (event) => {
          //console.log("HandlingnewMessage, event:", event)
          this.handleNewMessage(event);
        });
      }
    },
    close() {
      this.$emit("close");
    },
    
    handleNewMessage(event) {
      this.joinPrivateRoom()
      console.log("Handling new event message")
      let data = event.data;
      data = data.split(/\r?\n/);

      for (let i = 0; i < data.length; i++) {
        let msg = JSON.parse(data[i]);
        console.log("msg:", msg);
          // Check if the 'sender' property is present and has a value
    if (msg.sender) {
      console.log("Sender:", msg.sender);
    } else {
      console.log("Sender is null or undefined");
    }
        switch (msg.action) {
          case "send-message":
            console.log("Case: Send-message");
            this.handleChatMessage(msg);
            break;
          case "user-join":
            console.log("Case: user-join");
            this.handleUserJoined(msg);
            break;
          case "user-left":
            console.log("Case: user-left");
            this.handleUserLeft(msg);
            break;
          case "room-joined":
            console.log("Case: room-joined");
            this.handleRoomJoined(msg);
            break;
          case "join-room-private":
            this.joinPrivateRoom(msg);
            break;
          default:
          console.log("No case msg:",msg); 
            break;
        }
      }
    },
    handleChatMessage(msg) {
      const chatBox = this.findRoom(msg.target.id);
      console.log("Msg Target id:", msg.target.id, "Message Target name:", msg.target.name)
      if (typeof chatBox !== "undefined") {
        chatBox.messages.push({
        recip: this.title,
        author: msg.sender.name,
        message: msg.message,
    });
      }
      console.log("ChatBoxes!", this.chatBoxes)
      //console.log("chatBox:", chatBox,"\n ChatBox's name:", chatBox.name, "\n Target Chatbox id:", this.targetId, "\n Newmessage:", chatBox.newMessage)
      console.log("ChatBox messages after push:", chatBox.messages)
    },
    handleUserJoined(msg) {
      console.log("msg.sender:", msg.sender);
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
      console.log("ChatBoxes", this.chatBoxes)
      let chatBox = msg.target;
      chatBox.name = chatBox.private ? msg.sender.name : chatBox.name;
      chatBox.id = msg.target.id;
      //console.log("Room id?", chatBox.id,)
      chatBox["messages"] = [];
      this.chatBoxes.push(chatBox);
      this.targetId = chatBox.id
      console.log("roomJoined chatbox:", chatBox)
    },
    loadMessages(chatBox) {
    // Replace this with your actual API call to fetch messages from the server
    // You'll need to adjust the endpoint and request parameters based on your server implementation
    axios.get(`/api/messages/${chatBox.id}`)
      .then(response => {
        // Assuming the server returns an array of messages in the response
        chatBox.messages = response.data;
      })
      .catch(error => {
        console.error('Failed to load messages:', error);
      });
    },
    sendMessage(chatBox) {
      // console.log("chatBox.messages:", chatBox.message)
       
      //console.log("yuyi", this.targetId)
      if (chatBox.newMessage !== "") {
        console.log("chatBox:", chatBox,"\n ChatBox's name:", chatBox.name, "\n Target Chatbox id:", this.targetId, "\n Newmessage:", chatBox.newMessage)
          this.$socket.send(
            JSON.stringify({ 
              action: "send-message", 
              message: this.newMessage, 
              target: { 
                id: chatBox.id,
                name: chatBox.title
              }
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
  },
};

</script>
<style>
.sent-message {
  background-color: #F5D061;
  border-radius: 0.4rem 0.4rem 0rem 0.4rem;
  padding-right: 0.5rem;
  padding-left: 0.5rem;
  border: #282F44 solid 1px;

  /* Add other styles for sent messages */
}

.received-message {
  background-color: #F5D061;
  border-radius: 0.4rem 0.4rem 0.4rem 0;
  padding-right: 0.5rem;
  padding-left: 0.5rem;
  border: #282F44 solid 1px;
  
  /* Add other styles for received messages */
}
.chatBox{
  border-radius: 5px;
}
.messageTextArea{
  background-color: #E6AF2E;
  border-radius: 1rem;
  border: none;
  padding-left: 0.5rem;
  box-shadow: 0 0 1rem rgba(black, 0.1),
  0rem 1rem 1rem -1rem rgba(black, 0.2);
  margin: 0 0.5rem;
  width: 12rem;
}
.card-footer{
  display: flex;
  padding-bottom: 0.5rem !important;
  padding-top: 0.5rem !important;
  background-color: #F5D061 !important;
}
.content{
  height: 14rem;
  background-color: #E6AF2E !important;
}
.header{
  background-color: #F5D061 !important;
  border-color: #282F44 !important;
  border-width: 2px !important;
}
.message{
  margin-top: 10px;
  background-color: #F5D061;
  padding-left: 0.5rem;
  border: #282F44 solid 1px;
  
}
</style>
