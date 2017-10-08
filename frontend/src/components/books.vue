<template>
    <div>
        <a class="btn btn-primary" href="/#/add_book">Add book</a>
        <div class="" style="margin-top:20px" v-for="book in books">
            <router-link :to="{path:'/book/'+book.ID}" v-if="!book.DeletedAt">{{book.description}}</router-link>
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
