<template>
<div class="container-fluid h-100">
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
                       class="form-control type_msg"
                       placeholder="Type your message..."
                       v-on:keyup.enter.exact="sendMessage"
                       ></textarea>
             <div class="input-group-append">
               <span class="input-group-text send_btn" v-on:click="sendMessage"
                     >></span>
             </div>
           </div>
         </div>
       </div>
     </div>
   </div>
</div>
</template>
<script>


export default {
    data() {
        return{
    ws: null,
    serverUrl: "ws://localhost:8000/api/v1/ws",
    messages: [],
    newMessage: ""}
    },
    mounted: function() {
      this.connectToWebsocket()
    },
    methods: {
            sendMessage: function() {
    if(this.newMessage !== "") {
        this.ws.send(JSON.stringify({message: this.newMessage}));
        console.log(this.newMessage)
        this.newMessage = ""
    }
    },
      connectToWebsocket: function () {
        this.ws = new WebSocket( this.serverUrl );
        this.ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
        this.ws.addEventListener('message', (event) => { this.handleNewMessage(event) });
      },
      onWebsocketOpen: function () {
        console.log("connected to WS!");
      }
    },
    handleNewMessage: function (event) {
  let data = event.data;
  data = data.split(/\r?\n/);
  for (let i = 0; i < data.length; i++) {
    let msg = JSON.parse(data[i]);
    this.messages.push(msg);
  }   
},

}
</script>
<style>
    body,
html {
  height: 100%;
  margin: 0;
  background: #0310EA;
  background: -webkit-linear-gradient(to right, #560A86, #7122FA, #560A86);
  background: linear-gradient(to right, #560A86, #7122FA, #560A86);
}

#app{
    height: 100%;
}

.chat {
  margin-top: auto;
  margin-bottom: auto;
}
.card {
  height: 500px;
  border-radius: 15px;
  background-color: rgba(0, 0, 0, 0.4);
  border: 1px solid rgba(255,255,255, .2);
}

.msg_head{
    color: #FFF;
    text-align: center;
    font-size: 26px;
}

.msg_card_body {
  overflow-y: auto;
}
.card-header {
  border-radius: 15px 15px 0 0;
  border-bottom: 0;
}
.card-footer {
  border-radius: 0 0 15px 15px;
  border-top: 0;
}
.container {
  align-content: center;
}

.type_msg {
  background-color: rgba(86, 10, 134, 0.6);
  border: 0;
  color: white;
  height: 60px;
  overflow-y: auto;
}
.type_msg:focus {
  box-shadow: none;
  outline: 0px;
  background-color: rgba(255, 255, 255, 0.6);
}

.send_btn {
  border-radius: 0 15px 15px 0;
  background-color: rgba(0, 0, 0, 0.3);
  border: 0;
  color: white;
  cursor: pointer;
}

.msg_cotainer {
  margin-top: auto;
  margin-bottom: auto;
  margin-left: 10px;
  border-radius: 25px;
  background-color: #82ccdd;
  padding: 10px;
  position: relative;
}
.msg_cotainer_send {
  margin-top: auto;
  margin-bottom: auto;
  margin-right: 10px;
  border-radius: 25px;
  background-color: #75D5FD;
  padding: 10px;
  position: relative;
}

.msg_head {
  position: relative;
}
</style>