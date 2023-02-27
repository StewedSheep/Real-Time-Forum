<template>
<html>
  <head>
    <title>Chat</title>
    <!-- Load required Bootstrap and BootstrapVue CSS -->
    <link
      type="text/css"
      rel="stylesheet"
      href="//unpkg.com/bootstrap/dist/css/bootstrap.min.css"
    />
    <link
      type="text/css"
      rel="stylesheet"
      href="//unpkg.com/bootstrap-vue@latest/dist/bootstrap-vue.min.css"
    />
    <link type="text/css" rel="stylesheet" href="assets/style.css" />
    <!-- Load polyfills to support older browsers -->
    <!-- <script
      src="//polyfill.io/v3/polyfill.min.js?features=es2015%2CIntersectionObserver"
      crossorigin="anonymous"
    ></script> -->

    <!-- Load Vue followed by BootstrapVue -->
    <!-- <script src="https://unpkg.com/vue"></script>
    <script src="//unpkg.com/bootstrap-vue@latest/dist/bootstrap-vue.min.js"></script> -->
  </head>

  <body>
    <div id="app">
      <div class="container-fluid h-100">
        <div class="row justify-content-center h-100">
          <div class="col-md-8 col-xl-6 chat">
            <div class="card">
              <div class="card-header msg_head">
                <div class="d-flex bd-highlight justify-content-center">
                  Chat
                </div>
              </div>
              <div class="card-body msg_card_body">
                <div
                  v-for="(message, key) in messages"
                  :key="key"
                  class="d-flex justify-content-start mb-4"
                >
                  <div class="msg_cotainer">
                    {{message.message}}
                    <span class="msg_time"></span>
                  </div>
                </div>
              </div>
              <div class="card-footer">
                <div class="input-group">
                  <textarea
                    v-model="newMessage"
                    name=""
                    class="form-control type_msg"
                    placeholder="Type your message..."
                    @keyup.enter.exact="sendMessage"
                  ></textarea>
                  <div class="input-group-append">
                    <span class="input-group-text send_btn" @click="sendMessage"
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
data(){
  return{
      ws: null,
      serverUrl: "ws://localhost:8000/api/v1/ws",
      messages: [],
      newMessage: ""
}},
    mounted: function() {
      this.connectToWebsocket();
    },
    methods: {
      connectToWebsocket() {
        this.ws = new WebSocket( this.serverUrl );
        this.ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
        this.ws.addEventListener('message', (event) => { this.handleNewMessage(event) });
      },
      onWebsocketOpen() {
        console.log("connected to WS!");        
      },
      handleNewMessage(event) {
        let data = event.data;
        data = data.split(/\r?\n/);

        for (let i = 0; i < data.length; i++) {
            let msg = JSON.parse(data[i]);
            this.messages.push(msg);

        }   
      },
      sendMessage() {
        if(this.newMessage !== "") {
          this.ws.send(JSON.stringify({message: this.newMessage}));
          console.log(this.newMessage)
          this.newMessage = "";
        }
      }

    }
}
</script>
