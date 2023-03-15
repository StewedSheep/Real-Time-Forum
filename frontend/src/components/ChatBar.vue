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
    <b-button v-b-toggle.sidebar-right>Chat</b-button>
    <b-sidebar id="sidebar-right" title="Sidebar" right shadow>
      <div class="px-3 py-2">
        <div data-id="User.id" @click="clickUser(user)">
          <p v-for="user in Users" :key="user">
            {{ user }}
          </p>
        </div>
      </div>
    </b-sidebar>
  </div>
</template>
