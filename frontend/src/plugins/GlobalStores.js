import UserStore from '@/stores/UserStore'

export default{

    install(Vue){
        Vue.prototype.$user = UserStore;
    }

}