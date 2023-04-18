<script>
import Vue from "vue";

export default {
  name: "ChatBar.vue",
  components: {},
  data() {
    return {
      Users: [],
    };
  },

  created() {
    // this.connectToWebsocket();
  },
  mounted() {
    Vue.axios.get("/totUsers").then((response) => (this.Users = response.data));
  },
  methods: {
    connectToWebsocket() {
      //   this.user.name = this.$user.current;
      this.$socket.addEventListener("message", (event) => {
        this.handleNewMessage(event);
      });
      this.getAllConnections();
    },

    handleNewMessage(event) {
      let data = event.data;
      data = data.split(/\r?\n/);

      for (let i = 0; i < data.length; i++) {
        let msg = JSON.parse(data[i]);
        switch (msg.action) {
          case "all-connections":
            this.getAllConnections(msg);
            break;
          default:
            break;
        }
      }
    },
    getAllConnections() {
      console.log("got here");
      this.$socket.send(JSON.stringify({ action: "all-connections" }));
    },
    handleAllConnections(msg) {
      let connections = msg.message;
      console.log(connections);
    },
    // sortUsers(){
    //   let filteredUsers =  this.Users.filter((t) => {
    //     if("My-Posts"==categor){
    //       return t.author === this.$user.current
    //     }else{return t.category.includes(categor)}
    // })
    //   let orderedStories = filteredUsers.sort((a, b) => {
    //     return b.id - a.id;})
    //     return orderedStories
    // },
    clickUser(user) {
      //console.log(user); logs the user
      this.$router.push({ path: "/chat" });
    },
  },
};
</script>

<template>
  <div>
    <b-button id="chatButton" v-b-toggle.sidebar-right style="margin-left: 10px"
      >Chat</b-button
    >
    <b-sidebar class="chatBar" id="sidebar-right" title="Let's chat!" right shadow>
      <br />
      <div class="px-3 py-2">
        <div id="chatButton" v-for="user in Users" :key="user">
          <h5 id="chatBarButton">{{ user }}</h5>
          <button @click="clickUser(user)">Send Msg</button>
          <span id="chatBarButton" class="statusDot"></span>
          <br />
          <p id="chatBarButton">Last msg.</p>
          <p id="chatBarButton" style="float: right">19.03 11:11</p>
        </div>
      </div>
    </b-sidebar>
  </div>
</template>
