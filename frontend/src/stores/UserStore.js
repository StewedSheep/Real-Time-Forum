import Vue from 'vue';

class UserStore{
    constructor(){
        this.state = Vue.observable({
            user: null
        })
    }

    authorised(user){
        this.state.user = user;
    }

    get isAuthorised(){
        return this.state.user.username != undefined;
    }

    get current(){
        return this.state.user.username;
    }
}

export default new UserStore();