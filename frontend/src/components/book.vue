<template>
    <div>
        <div v-if="book.DeletedAt">
            The book is deleted
        </div>
        <div class="row" v-else>
            <div class="col-md-3">
                <div>
                    <router-link :to="{path: '/book/' + book.ID + '/edit'}">
                        Edit
                    </router-link>
                    <router-link :to="{path: '/book/' + book.ID + '/add_chapter'}">
                        Add chapter
                    </router-link>
                </div>
                <ul class="list-group">
                    <li class="hover list-group-item" v-on:click="get_chapter(0)">Home</li>
                    <li v-for="(chapter,i) in chapters" class="hover list-group-item" v-on:click="get_chapter(i + 1)">{{chapter}}</li>
                </ul>
            </div>
            <div class="col-md-9" v-if="chapter.order">
                <div class="pull-right">
                    <router-link 
                        :to="{path:'/book/'+book.ID+'/ch/'+chapter.order+'/edit'}">
                        Edit
                    </router-link>
                    <span v-on:click="delete_chapter()">Delete</span>
                </div>
                <div class="clearfix"></div>
                <div>{{chapter.content}}</div>
            </div>
            <div class="col-md-9" v-else>
                <div class="h1">{{book.name}}</div>
                <div>{{book.description}}</div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name : "book",
  created() {
      var url = 'book/'+this.$route.params.id
      this.$store.dispatch('get_book',url)
  },
  computed: {
    book() {
        return this.$store.state.book
    },
    chapters() {
        return this.$store.state.chapters
    },
    chapter() {
        return this.$store.state.chapter
    }
  },
  methods: {
    get_chapter(ch_id) {
      var url
      if(ch_id == 0)
        url = 'book/'+this.$route.params.id
      else
        url = 'book/'+this.$route.params.id + '/ch/' + ch_id
      this.$store.dispatch('get_book',url)
    },
    delete_chapter(){
        this.$store.dispatch('delete_chapter',this.chapter)
    }
  }
}
</script>
