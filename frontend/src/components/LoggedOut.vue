<template>
      <div id="app">
      <b-navbar style=" background-color: #F5D061 !important;" class="navBar" toggleable="lg" type="dark" variant="info">
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
</template>
