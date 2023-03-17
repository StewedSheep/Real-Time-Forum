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
  mounted() {
    Vue.axios.get("/totUsers").then((response) => (this.Users = response.data));
    console.log(this.Users);
  },

  methods: {
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
      console.log(user);
      this.$router.push({ path: "/chat", query: { user } });
    },
  },
};
</script>

<template>
  <div>
    <b-button id="chatButton" v-b-toggle.sidebar-right style="margin-left: 10px;">Chat</b-button>
    <b-sidebar class="chatBar" id="sidebar-right" title="Let's chat!" right shadow>
      <br>
      <div class="px-3 py-2">
        <div data-id="User.id" @click="clickUser(user)">
          <div  id="chatButton" v-for="user in Users" :key="user">
            <h5 id="chatBarButton">{{ user }}</h5> 
            <span id="chatBarButton" class="statusDot"></span>
            <br>
            <p id="chatBarButton">Last msg.</p>
            <p id="chatBarButton" style="float: right;">19.03 11:11</p>
          </div>
        </div>
      </div>
    </b-sidebar>
  </div>
</template>
