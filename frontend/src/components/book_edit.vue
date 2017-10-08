<template>
    <div>
        <div class="form-group">
            <label>Name</label>
            <input class="form-control" v-model="book.name">
        </div>
        <div class="form-group">
            <label>Description</label>
            <input class="form-control" v-model="book.description">
        </div>
        <button @click="create_or_edit_book()">Sumbit</button>
    </div>
</template>

<script>
export default {
  name : "book_edit",
  methods: {
      create_or_edit_book() {
          this.$store.dispatch('create_or_edit',this.book)
      }
  },
  computed: {
    book() {
        if(this.$route.params.id)
            return this.$store.state.book
        else
            return {}
    }
  }
}
</script>
