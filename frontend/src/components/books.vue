<template>
    <div>
        <a class="btn btn-primary" href="/#/add_book">Add book</a>
        <div class="" style="margin-top:20px" v-for="book in books">
            <router-link :to="{path:'/book/'+book.ID}" v-if="!book.DeletedAt">
                {{book.name}}
            </router-link>
            <span class="hover pull-right" v-on:click="delete_book(book)" >delete</span>
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
  },
  methods : {
      delete_book(book){
          this.$store.dispatch('delete_book',book)
      }
  }
}
</script>
