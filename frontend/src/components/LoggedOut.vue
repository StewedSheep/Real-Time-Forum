<template>
      <div id="app">
      <b-navbar style=" background-color: #F5D061 !important;" class="navBar" toggleable="lg" type="dark" variant="info">
        <b-navbar-brand href="#">NavBar</b-navbar-brand>
          <b-navbar-nav class="ml-auto">
              <template #button-content>
              </template>
              <router-link class="logButton" to="login">Login</router-link>
              <router-link  class="logButton" to="register">Register</router-link>
          </b-navbar-nav>
      
      </b-navbar>
      </div>
</template>
