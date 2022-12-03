<template>
      <div id="app">
      <b-navbar class="navBar" toggleable="lg" type="dark" variant="info">
        <b-navbar-brand href="#">NavBar</b-navbar-brand>
          <b-navbar-nav class="ml-auto">
           <b-nav-item-dropdown right>
              <template #button-content>
              </template>
              <b-dropdown-item href="#"><router-link class="logButton" to="login">Login</router-link></b-dropdown-item>
              <b-dropdown-item href="#"><router-link  class="logButton" to="register">Register</router-link></b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
      
      </b-navbar>
      </div>
    <!-- <div class="menu">
        <router-link class="logButton" to="login">Login</router-link>
        or
        <router-link  class="logButton" to="register">Register</router-link>
    </div> -->
</template>
