<template>
    <div id="app">
    <div class="menu">
                <h4>Logged in as: {{$user.current}} </h4>
                <router-link @click.native="logOut()" to="/">Log Out</router-link>
                    or
                <router-link to="newThread">Create a New Thread</router-link>
            </div>
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
