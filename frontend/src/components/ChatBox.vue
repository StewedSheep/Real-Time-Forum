<template>
  <div v-if="visible" class="chatbox">
    <div class="header">
      <h3>{{ title }}</h3>
      <button @click="close">&times;</button>
    </div>
    <div class="content">
      <!-- Chat content goes here -->
    </div>
  </div>
</template>

<script>
export default {
  name: "ChatBox",
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
  methods: {
    close() {
      this.$emit("close");
    },
  },
};
</script>
