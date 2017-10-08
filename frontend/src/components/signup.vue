<template>
    <div>
        <div class="form-group">
            <label>Username</label>
            <input class="form-control" v-model="user.name">
        </div>
        <div class="form-group">
            <label>Password</label>
            <input class="form-control" v-model="user.password">
        </div>
        <div class="form-group">
            <label>Password again</label>
            <input class="form-control" v-model="user.password1">
        </div>
    </div>
</template>

<script>
export default {
  name : "book",
  methods: {
      create_user() {
          this.$store.dispatch('create_user',this.user)
      }
}
</script>
