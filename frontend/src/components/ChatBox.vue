<template>
  <div v-if="visible" ref="chatbox" class="chatbox">
    <div class="header">
      <h3>{{ title }}</h3>
      <button class="logButton" @click="close()">&times;</button>
    </div>
    <div v-chat-scroll ref="content" class="content" @scroll="handleScroll">
      <!-- Chat content goes here -->
      <!-- <div v-for="chatBox in chatBoxes" :key="chatBox.id"> -->
      <!-- Only display messages when the chatbox is open -->
      <div
        v-for="msg in messages"
        :key="msg.id"
        :class="{
          'sent-message': msg.from !== title,
          'received-message': msg.from == title,
        }"
        class="message"
      >
        {{ msg.from }}: {{ msg.content }} {{ msg.date }}
      </div>
      <div ref="bottomEl"></div>
    </div>
    <div class="card-footer">
      <input
        v-model="newMessage"
        class="messageTextArea"
        name=""
        placeholder="Type your message..."
        @keyup.enter.exact="sendMessage()"
      />
      <button class="logButton input-group-append" @click="sendMessage()">Send</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
//import { debounce } from "lodash";
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
      },
      message: {
        to: "",
        from: "",
        content: "",
        date: "",
      },
      newMessage: "",
      users: [],
      user: {
        id: "",
        name: "",
      },
      messages: [],
      page: 1,
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

  updated() {},
  created() {
    // this.$socket.onmessage = (event) => {
    //   try {
    //     const data = JSON.parse(event.data);
    //     this.messages.push(data);
    //   } catch (error) {
    //     console.error("Error parsing message data:", error);
    //   }
    // };
  },

  mounted() {
    this.loadMoreMessages();
  },
  methods: {
    handleScroll() {
      const container = this.$refs.content;
      if (container.scrollTop === 0) {
        this.loadMoreMessages();
      }
    },

    close() {
      this.$emit("close");
    },

    loadMoreMessages() {
      try {
        axios
          .get(
            `http://localhost:8000/api/v1/chatDb?from=${this.$user.current}&to=${this.title}&page=${this.page}`
          )
          .then((response) => {
            console.log(JSON.stringify(response.data));
            if (response.data != null && this.page == 1) {
              this.messages = response.data;
              this.page++;
            } else {
              [...this.messages] = [...response.data, ...this.messages];
              this.page++;
            }
          });
      } catch (error) {
        console.error("Error parsing message data:", error);
      }
    },

    sendMessage() {
      if (this.$socket && this.$socket.readyState === WebSocket.OPEN) {
        if (this.newMessage !== "") {
          this.$socket.send(
            JSON.stringify({
              to: this.title,
              content: this.newMessage,
              date: new Date().toISOString(),
            })
          );
          this.messages.push({
            to: this.title,
            from: this.$user.current,
            content: this.newMessage,
            date: new Date().toISOString(),
          });
          this.newMessage = "";
        } else {
          console.log(
            "Unable to send message. The WebSocket connection is not open or no input."
          );
        }
      }
    },
  },
};
</script>
<style>
.sent-message {
  background-color: #f5d061;
  border-radius: 0.4rem 0.4rem 0rem 0.4rem;
  padding-right: 0.5rem;
  padding-left: 0.5rem;
  border: #282f44 solid 1px;

  /* Add other styles for sent messages */
}

.received-message {
  background-color: #f5c361;
  border-radius: 0.4rem 0.4rem 0.4rem 0;
  padding-right: 0.5rem;
  padding-left: 0.5rem;
  border: #282f44 solid 1px;

  /* Add other styles for received messages */
}
.chatbox {
  border-radius: 5px;
}
.messageTextArea {
  background-color: #e6af2e;
  border-radius: 1rem;
  border: none;
  padding-left: 0.5rem;
  box-shadow: 0 0 1rem rgba(black, 0.1), 0rem 1rem 1rem -1rem rgba(black, 0.2);
  margin: 0 0.5rem;
  width: 12rem;
}
.card-footer {
  display: flex;
  padding-bottom: 0.5rem !important;
  padding-top: 0.5rem !important;
  background-color: #f5d061 !important;
}
.content {
  height: 14rem;
  background-color: #e6af2e !important;
  height: 300px;
  overflow-y: auto;
}
.header {
  background-color: #f5d061 !important;
  border-color: #282f44 !important;
  border-width: 2px !important;
}
.message {
  margin-top: 10px;
  background-color: #f5d061;
  padding-left: 0.5rem;
  border: #282f44 solid 1px;
}
</style>
