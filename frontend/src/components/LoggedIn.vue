<template>

     <div id=app>
       <b-navbar style=" background-color: #F5D061 !important;" class="navBar" toggleable="lg" type="dark" variant="info">
        <b-navbar-brand href="#">NavBar</b-navbar-brand>
        
        <ChatBar />
          <b-navbar-nav class="ml-auto">
            <em class="loggedInAs">Logged in as: {{$user.current}} </em> 
            <b-nav-item-dropdown right>
              <template #button-content>
              </template>
              <b-dropdown-item href="#"><router-link class="logButton" to="newThread">Create a New Thread</router-link></b-dropdown-item>
              <b-dropdown-item href="#"><router-link class="logButton" @click.native="logOut()" to="/">Log Out</router-link></b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
      
      </b-navbar>    
              
              
    </div>
</template>

<script>
import Vue from 'vue'

import ChatBar from './ChatBar.vue'

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
  },
  components: {
    "ChatBar": ChatBar,

},
}
</script>>
