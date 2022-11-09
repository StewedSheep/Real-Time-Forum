<template>
  <div>
    <PageHeader/>
    <component :key="componentKey" v-bind:is="comp" v-bind:loggedUsername="loggedUsername" />
    <component :key="componentKey" v-bind:is="component" />
    <router-view/>
  </div>
</template>

<script>
import LoggedIn from './components/LoggedIn.vue'
import LoggedOut from './components/LoggedOut.vue'
import LoggedNavBar from './components/LoggedNavBar.vue';
import UnloggedNavBar from './components/UnloggedNavBar.vue';
import PageHeader from './components/PageHeader.vue';

let loggedUsername;

console.log($store.auth.user);

export default {
  data() {
    return{loggedUsername, componentKey: 0, comp: LoggedOut, component: UnloggedNavBar}
   },
  //checks for cookie updates every 100ms
created: function() {
  const timer = setInterval(() => {
    this.loggedUsername = document.cookie.split('::')[1];
    $store.auth.set(this.loggedUsername);
  }, 100);
  //destroys last check hook so new can be run
  this.$once("hook:beforeDestroy", () => {
    clearInterval(timer);
  });
} ,
  //watches for changes in logged in username and changes components
  watch: {
    loggedUsername(newVal, oldVal) {
      if (newVal != ""){
      this.comp = LoggedIn;
      this.component = LoggedNavBar;
      }
      if (newVal == null){
      this.comp = LoggedOut;
      this.component = UnloggedNavBar;
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
    "LoggedNavBar": LoggedNavBar,
    "UnloggedNavBar": UnloggedNavBar,
},
}
</script>

<style>
  @import './assets/styles/main.css';
  #app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  width: 100%;
}
</style>