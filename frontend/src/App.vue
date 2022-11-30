<template>
  <div>
    <PageHeader/>
    <component v-bind:is="comp" />
    <router-view/>
  </div>
</template>

<script>
import LoggedIn from './components/LoggedIn.vue'
import LoggedOut from './components/LoggedOut.vue'
import PageHeader from './components/PageHeader.vue';

let loggedUsername;

export default {
  data() {
    return{loggedUsername,
      componentKey: 0,
      comp: LoggedOut,}
   },
  //checks for cookie updates every 100ms
created: function() {
  const timer = setInterval(() => {
    this.loggedUsername = document.cookie.split('::')[1];
    //sets username to global storage
    this.$user.authorised({
      username : this.loggedUsername,
    })
  }, 100);
  //destroys last check hook so new can be run
  this.$once("hook:beforeDestroy", () => {
    clearInterval(timer);
  });
},
  //watches for changes in logged in username and changes components
  watch: {
    loggedUsername(newVal, oldVal) {
      if (newVal != ""){
      this.comp = LoggedIn;
      }
      if (newVal == null){
      this.comp = LoggedOut;
      }
      console.log(newVal);
      if(oldVal != newVal){
        this.componentKey++;
      }
      console.log("loggedUsername watcher triggered")
    },
    immediate: true
  },
  name: "app",
  components: {
    "PageHeader": PageHeader,
    "LoggedIn": LoggedIn,
    "LoggedOut": LoggedOut,
},
}
</script>

<style>
  @import './assets/styles/main.css';

</style>