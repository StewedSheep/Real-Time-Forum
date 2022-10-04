<template>
    <div id="app">
                <p>Sort posts by:</p>
    
            <div v-if="loggedUsername" class="btn-group" style="width:100%">
                <router-link to="general"><button href="general" style="width:14.28%">General</button></router-link>
                <router-link to="help"><button href="help" style="width:14.28%">Help</button></router-link>
                <router-link to="news"><button href="news" style="width:14.28%">News</button></router-link>
                <router-link to="discussion"><button href="discussion" style="width:14.28%">Discussion</button></router-link>
                <router-link to="offTopic"><button href="offTopic" style="width:14.28%">Off-Topic</button></router-link>
                <router-link to="likedPosts"><button href="likedPosts" style="width:14.28%">Liked-Posts</button></router-link>
                <router-link to="myPosts"><button href="myPosts" style="width:14.28%">My-Posts</button></router-link>
              </div>
    
    
              <div v-else class="btn-group" style="width:100%">
                <router-link to="general"><button href="general" style="width:20%">General</button></router-link>
                <router-link to="help"><button href="help" style="width:20%">Help</button></router-link>
                <router-link to="news"><button href="news" style="width:20%">News</button></router-link>
                <router-link to="discussion"><button href="discussion" style="width:20%">Discussion</button></router-link>
                <router-link to="offTopic"><button href="offTopic" style="width:20%">Off-Topic</button></router-link>
              </div>
            <br><br>
            </div>
</template>

<script>
export default {
    props: ['loggedUsername'],
  data() {}
}
</script>