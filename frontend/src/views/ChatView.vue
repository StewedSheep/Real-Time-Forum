<template>
  <div id="app">
    <input type="text" v-model="recipient" placeholder="Recipient username">
    <input type="text" v-model="message" @keyup.enter="sendMessage">
    <div v-for="msg in messages" :key="msg.from + msg.content">
      <b>{{ msg.from }}</b> to <i>{{ msg.to }}</i>: {{ msg.content }}
    </div>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      recipient: '',
      message: '',
      messages: [],
    };
  },
  methods: {
    sendMessage() {
      this.$socket.send(JSON.stringify({to: this.recipient, from: this.$user.current, content: this.message}));
      this.message = '';
    },
  },
  sockets: {
    onmessage(data) {
      this.messages.push(data.data);
    },
  },
};
</script>