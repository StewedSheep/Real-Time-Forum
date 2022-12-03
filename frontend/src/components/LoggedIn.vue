<template>

     <div id="app">
      <b-navbar class="navBar" toggleable="lg" type="dark" variant="info">
        <b-navbar-brand href="#">Biitch!</b-navbar-brand>
          <b-navbar-nav class="ml-auto">
           <b-nav-item-dropdown right>
                <em>Logged in as: {{$user.current}} </em>     
              <b-dropdown-item href="#"><router-link class="logButton" to="newThread">Create a New Thread</router-link></b-dropdown-item>
              <b-dropdown-item href="#"><router-link class="logButton" @click.native="logOut()" to="/">Log Out</router-link></b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
      </b-navbar>
      <!--
      <div class="menu">
        <h4 class="logged">Logged in as: {{$user.current}} </h4>  
        <router-link class="logButton" to="newThread">Create a New Thread</router-link>
        or
       <router-link class="logButton" @click.native="logOut()" to="/">Log Out</router-link>
     </div>-->

    </div>
</template>

<script>
import Vue from 'vue'
export default {
    props: ['loggedUsername'],
    methods : {
      logOut: function() {
      Vue.axios.post('/logout', {
    headers: {
          'Content-Type': 'text/plain',
          'Access-Control-Allow-Origin': '*'
        }
      })
    },
    // redirecting???   sucReg: function() {{setTimeout(() => this.$router.push({ name: 'login' }), 3000)}},
  }
}
</script>>
