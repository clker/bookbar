<template>
    <div>
        <router-link to="/add_book">Add book</router-link>
        <div class="panel panel-default" v-for="(book,idx) in books">
          <div class="panel-heading">{{book.name}}</div>
          <div class="panel-body">{{book.description}}</div>
        </div>
    </div>
</template>

<script>
export default {
  name : "books",
  created() {
      this.$store.dispatch('get_books')
  },
  computed: {
    books() {
        return this.$store.state.books
    }
  }
}
</script>
