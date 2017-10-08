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
        <button @click="create_chapter(book.id)">Sumbit</button>
    </div>
</template>

<script>
export default {
  name : "chapter_edit",
  methods: {
      create_chapter() {
          this.$store.dispatch('add_chapter',this.chapter,this.$route.params.id)
      }
  },
  computed: {
    chapter() {
        return this.$store.state.chapter
    },
    book(){
        return this.$store.state.book
    }
  }
}
</script>
