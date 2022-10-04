<template>
  <div>
    <PageHeader/>
    <LoginRegister v-bind:loggedUsername="loggedUsername"/>
    <NavigationBar v-bind:loggedUsername="loggedUsername"/>
    <router-view/>
  </div>
</template>

<script>
import LoginRegister from './components/LoginRegister.vue';
import NavigationBar from './components/NavigationBar.vue';
import PageHeader from './components/PageHeader.vue';

const cookie = document.cookie.split('::');
const loggedUsername = cookie[1];

browser.cookies.onChanged.addListener(removed)=> {
    if (removed) {
        console.log("cookie gone")}};

export default {
   data() {
    return{loggedUsername}
   },
  name: "app",
  components: {
    "LoginRegister": LoginRegister,
    "NavigationBar": NavigationBar,
    "PageHeader": PageHeader,
    
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