<script>
import Vue from "vue";
import { toRef } from "vue";
import ChatBox from "./ChatBox.vue";

export default {
  name: "ChatBar.vue",
  props: {
    users: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      chatBoxes: [],
      user: {
        name: "",
      },
      allUsers: [],
    };
  },
  setup(props) {
    const usersRef = toRef(props, "users");
    return { usersRef };
  },
  mounted() {
    Vue.axios.get("/totUsers").then((response) => (this.allUsers = response.data));
  },
  methods: {
    openChatBox(user) {
      this.chatBoxes.push({
        title: user,
        visible: true,
      });
    },

    closeChatBox(index) {
      this.chatBoxes.splice(index, 1);
    },
    // based on last message and then in alphabetical order
    sortUsers() {
      return this.allUsers.filter((user) => user !== this.$user.current);
    },

    clickUser() {
      this.$router.push({ path: "/chat" });
    },
  },
  components: {
    ChatBox,
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
        <!-- sorts user order -->
        <div id="chatButton" v-for="user in sortUsers()" :key="user">
          <h5 id="chatBarButton">{{ user }}</h5>
          <div>
            <button @click="openChatBox(user)">Message</button>
          </div>
          <!-- sets online indicator -->
          <div v-for="actUser in users" :key="actUser.name">
            <!-- <p>
              {{ actUser.name }}
            </p> -->
            <span
              v-if="user == actUser.name"
              id="chatBarButton"
              class="statusDotOnline"
            />
            <span v-else id="chatBarButton" class="statusDotOffline" />
          </div>

          <br />
          <!-- last message data -->
          <p id="chatBarButton">Last msg.</p>
          <p id="chatBarButton" style="float: right">19.03 11:11</p>
        </div>
      </div>
    </b-sidebar>
    <!-- message box containers -->
    <div class="chat-box-container">
      <div
        v-for="(chatBox, index) in chatBoxes"
        :key="index"
        class="chat-box"
        :class="{ visible: chatBox.visible }"
      >
        <ChatBox
          :title="chatBox.title"
          :visible="chatBox.visible"
          @close="closeChatBox(index)"
        />
      </div>
    </div>
  </div>
</template>
