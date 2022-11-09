<script>
export default {
  name: "NewThreadView.vue",
  components: {}
}
</script>


<template>
<div id="app">
        <!-- <div v-if="Username" class="main"> -->

        <br>
        <form action="newthreadinf" method="POST">
        <tr>
            <td>Thread topic:</td>
            <td>
                <input type="text" name="thtopic" id="thtopic" minlength="2" required>
                <label for="catergory">Choose a category:</label>
                <select name="category" id="category" required>
                    <option value="general">General</option>
                    <option value="help">Help</option>
                    <option value="news">News</option>
                    <option value="discussion">Discussion</option>
                    <option value="offtopic">Off-topic</option>
                  </select>
            </td>
        </tr>   
    <textarea type="text" name="content" id="content" required></textarea>
    <td><input type="submit" name="submit" value="Post Thread"></td>
    </form>
<!-- </div>
    <div v-else>
        <h3>Please log in or create user to post a new thread.</h3>
    </div> -->
<router-view />
</div>
</template>