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
        <button @click="create_book()">Sumbit</button>
    </div>
</template>

<script>
export default {
  name : "book_edit",
  computed : {
      book(){
          if(this.$store.state.add_book_done) this.$router.go(-1)
          else return {
              name : '',
              description : '',
          }
      }
  },
  methods: {
      create_book() {
          this.$store.dispatch('add_book',this.book)
      }
  }
}
</script>
