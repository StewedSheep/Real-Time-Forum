<script>
import Vue from 'vue'

export default {
  name: "LoginView.vue",
  components: {},
    data (){
        return {
            username : '',
            password : '',
            LoginError : '',
        }
    },
    methods : {
        postLog: function() {
      var data = {username: this.username,
                  password: this.password,
                }
  console.log(data)
      Vue.axios.post('/login', data, {
    headers: {
          'Content-Type': 'text/plain',
          'Access-Control-Allow-Origin': '*'
            }
        }).then((response) => (this.LoginError = response.data));
      },
    }
}
</script>


<template>
<div id="app">
    <center>
        <h3>Login here</h3>

        <div v-if="LoginError != null">
    <p>{{ LoginError }}</p>

    <br/>
    </div> 

<form>
            <label>Username or Email:</label><br>
            <input type="text" v-model="username" name="username" /><br><br>

            <label for="password">Password:</label><br>
            <input type="password" v-model="password" name="password" /><br><br>

            <br>
          </form> 

            <div>
      <button type="button" @click="postLog()">Login</button>  <br><br>
<td>
    No account? <router-link to="register">Register here!</router-link>
</td>
    </div>
<router-view />
</center>
</div>
</template>