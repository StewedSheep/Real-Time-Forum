<script>
import Vue from 'vue'

export default {
  name: "RegisterView",
  components: {},
  data: function (){
    return {
        RegisterError: ""
    }
  },
  methods : {
      postReg: function() {
      var data = {email: this.email,
                  username: this.username,
                  password: this.password,
                  firstName: this.firstName,
                  lastName: this.lastName,
                  gender: this.gender,
                  birthday: this.birthday,
                }
      Vue.axios.post('/register', data, {
    headers: {
          'Content-Type': 'text/plain',
          'Access-Control-Allow-Origin': '*'
        }
      }).then((response) => (this.RegisterError = response.data))
    },
    // redirecting???   sucReg: function() {{setTimeout(() => this.$router.push({ name: 'login' }), 3000)}},
  }
}
</script>


<template>
        <div id="app">
        <center>
            <h3>Register a New Account</h3>
        <br>
  <div v-if="RegisterError != null">
    <p>{{ RegisterError }}</p>
    </div>
        <br>
        <form>
            <label>Email:</label><br>
            <input type="text" v-model="email" name="email" /><br><br>

            <label for="username">Username:</label><br>
            <input type="text" v-model="username" name="username" /><br><br>

            <label for="password">Password:</label><br>
            <input type="password" v-model="password" name="password" /><br><br>

            <label for="firstName">First Name:</label><br>
            <input type="text" v-model="firstName" name="firstName" /><br><br>

            <label for="lastName">Last Name:</label><br>
            <input type="text" v-model="lastName" name="lastName" /><br><br>

            <label for="gender">Gender:</label><br>
           <select v-model="gender">
            <option>Male</option>
            <option>Female</option>
            </select><br><br>

            <label for="birthday">Birthday:</label><br>
            <input type="date" v-model="birthday" name="birthday">
            <br><br><br>
          </form> 

            <div>
      <button type="button" @click="postReg()">Submit</button>  
    </div>
  <br><br><br>

<li>Username has to be between 5 and 16 characters long and contain no space</li>
<li>usernames must contain only letters or numbers</li>
<li>passwords must contain a uppercase letter, lowercase letter and number</li>
<li>passwords must be 6 - 16</li>
<li>First and last name must be in latin letters</li>
<li>Only users above the age of 13 can register</li>
</center>
</div>
</template>

<style>
.button {
  background-color: #f4511e;
  border: none;
  color: white;
  padding: 16px 32px;
  text-align: center;
  font-size: 16px;
  margin: 4px 2px;
  opacity: 0.6;
  transition: 0.3s;
  display: inline-block;
  text-decoration: none;
  cursor: pointer;
}

.button:hover {opacity: 1}
</style>