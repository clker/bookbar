<template>
    <div>
        <div class="form-group">
            <label>Title</label>
            <input class="form-control" v-model="chapter.title">
        </div>
        <div class="form-group">
            <label>Content</label>
            <input class="form-control" v-model="chapter.content">
        </div>
        <button @click="create_or_edit_chapter(book.id)">Sumbit</button>
    </div>
</template>

<script>
export default {
  name : "chapter_edit",
  methods: {
      create_or_edit_chapter() {
          this.$store.dispatch('create_or_edit_chapter',{
              chapter:this.chapter,
              book_id:this.$route.params.id
          })
      }
  },
  computed: {
    chapter() {
        if(this.$route.params.ch_id)
            return this.$store.state.chapter
        else
            return {}
    },
    book(){
        return this.$store.state.book
    }
  }
}
</script>
