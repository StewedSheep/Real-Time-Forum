<script>
import EventBus from "@/stores/event-bus.js";
import Vue from "vue";
import { toRef } from "vue";
import ChatBox from "./ChatBox.vue";

export default {
  name: "ChatBar.vue",
  props: {
    msgList: {
      type: Array,
      required: true,
    },
    users: {
      type: Array,
      required: true,
    },
    list: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      user: {
        name: "",
      },
      allUsers: [],
      chatBoxes: [],
    };
  },
  setup(props) {
    const usersRef = toRef(props, "users");
    return { usersRef };
  },
  mounted() {
    Vue.axios.get("/totUsers").then((response) => (this.allUsers = response.data));
    this.sortUsers();
  },
  methods: {
    openChatBox(user) {
      this.chatBoxes.forEach((chatBox) => {
        if (chatBox.title === user) {
          this.chatBoxes = this.chatBoxes.filter((item) => item.title !== user);
        }
      });
      this.chatBoxes.push({
        title: user,
        visible: true,
      });
      EventBus.$emit("chatbox-opened", this.chatBoxes);
    },
    closeChatBox(index) {
      this.chatBoxes.splice(index, 1);
      EventBus.$emit("chatbox-closed", this.chatBoxes);
    },

    // based on last message and then in alphabetical order
    sortUsers() {
      const customSort = (a, b) => {
        const dateA = this.list[a];
        const dateB = this.list[b];
        if (
          (dateA === "no date" || dateA === "no msgs") &&
          (dateB === "no date" || dateB === "no msgs")
        ) {
          return a.localeCompare(b);
        } else if (dateA === "no date" || dateA === "no msgs") {
          return 1;
        } else if (dateB === "no date" || dateB === "no msgs") {
          return -1;
        }
        return dateB.localeCompare(dateA);
      };
      this.allUsers = this.allUsers.sort(customSort);

      // Remove the users name from the array
      const index = this.allUsers.findIndex((name) => name === this.$user.current);
      if (index !== -1) {
        this.allUsers.splice(index, 1);
      }
    },
  },
  watch: {
    list: {
      deep: true,
      handler() {
        this.sortUsers();
      },
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
        <div
          id="chatButton"
          v-for="user in this.allUsers"
          :key="user"
          @click="openChatBox(user)"
        >
          <h5 id="chatBarButton">{{ user }}</h5>
          <!-- sets online indicator -->
          <div v-for="actUser in users" :key="actUser">
            <span v-if="user == actUser" id="chatBarButton" class="statusDotOnline" />
            <span v-else id="chatBarButton" class="statusDotOffline" />
          </div>
          <br />
        </div>

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
    </b-sidebar>
  </div>
</template>
