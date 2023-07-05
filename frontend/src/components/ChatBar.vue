<script>
import EventBus from "@/stores/event-bus.js";
import Vue from "vue";
import { toRef } from "vue";
import ChatBox from "./ChatBox.vue";

export default {
  name: "ChatBar.vue",
  data() {
    return {
      user: {
        name: "",
      },
      totUsers: [],
      allUsers: [{ name: "", online: false, date: "" }],
      chatBoxes: [],
    };
  },
  setup(props) {
    const usersRef = toRef(props, "users");
    return { usersRef };
  },
  mounted() {
    EventBus.$on("users", (eventData) => {
      this.users = eventData;
      console.log("chatbar.vue", eventData);
      this.sortUsers();
    });
    EventBus.$on("list", (eventData) => {
      this.list = eventData;
      console.log("chatbar.vue", eventData);
      this.sortUsers();
    });
    Vue.axios.get("/totUsers").then((response) => (this.totUsers = response.data));
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
      const options = {
        dateStyle: "short",
        timeStyle: "short",
      };

      this.allUsers = this.totUsers.map((name) => ({
        name: name,
        online: this.users.includes(name),
        date: new Date(this.list[name]).toLocaleString("it-IT", options),
      }));
      // Remove the users name from the array
      this.allUsers = this.allUsers.filter((user) => user.name !== this.$user.current);

      const compareUsernames = (a, b) => {
        // Handle "Invalid Date" or "no date" cases
        if (
          a.date === "Invalid Date" ||
          b.date === "Invalid Date" ||
          a.date === "no date" ||
          b.date === "no date"
        ) {
          if (a.date === "Invalid Date" || a.date === "no date") {
            return 1;
          }
          if (b.date === "Invalid Date" || b.date === "no date") {
            return -1;
          }
        }

        // Compare dates
        const dateA = new Date(this.list[a.name]);
        const dateB = new Date(this.list[b.name]);
        if (dateA > dateB) return -1;
        if (dateA < dateB) return 1;

        // Compare names alphabetically
        return a.name.localeCompare(b.name);
      };

      // Sorting the usernames array
      this.allUsers.sort(compareUsernames);
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
          @click="openChatBox(user.name)"
        >
          <h5 id="chatBarButton">{{ user.name }}</h5>
          <!-- sets online indicator -->
          <span v-if="user.online == true" id="chatBarButton" class="statusDotOnline" />
          <br />
          <div v-if="user.date != 'Invalid Date'">
            <p id="chatBarButton">Last message:</p>
            <p id="chatBarButton" style="float: right">{{ user.date }}</p>
          </div>
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
