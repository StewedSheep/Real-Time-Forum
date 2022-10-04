<template>
<div id="app">
    <div v-if="loggedUsername" class="menu">
                <h4>Logged in as: {{ loggedUsername }} </h4>
                <span @click="logOut()">Log Out</span>
                    or
                <router-link to="newThread">Create a New Thread</router-link>
            </div>
            <div v-else class="menu">
                <router-link to="login">Login</router-link>
                or
                <router-link to="register">Register</router-link>
                </div>
            <br>
        </div>
</template>

<script>
import Vue from 'vue'

export default {
    props: ['loggedUsername'],
  data() {},
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
</script>