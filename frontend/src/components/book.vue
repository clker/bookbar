<template>
    <div>
        <div class="layout">
            <div class="layout-ceiling">
                <div class="layout-ceiling-main">
                    <span v-if="chapter.order">
                        <router-link :to="{path: '/book/' + book.ID + '/ch/' + chapter.order + '/edit'}">
                            Edit Chapter
                        </router-link>
                    </span>
                    <span v-else>
                        <router-link :to="{path: '/book/' + book.ID + '/edit'}">
                            Edit Book 
                        </router-link>
                    </span>
                    <router-link :to="{path: '/book/' + book.ID + '/add_chapter'}">
                        Add chapter
                    </router-link>
                    <span class="primary" v-on:click="delete_chapter()">Delete</span>
                </div>
                <div class="clearfix"></div>
            </div>
        </div>
        <div v-if="book.DeletedAt">
            The book is deleted
        </div>
        <Row v-else>
            <Col span="5">
                <Menu theme="light" active-name="0">
                    <MenuItem :name="0">
                        <div v-on:click="get_chapter(0)">
                            Cover
                        </div>
                    </MenuItem>
                    <div v-for="(chapter,i) in chapters">
                        <MenuItem :name="i+1">
                            <div v-on:click="get_chapter(i+1)">
                                {{chapter}}
                            </div>
                        </MenuItem>
                    </div>
                </Menu>
            </Col>
            <Col span="19" class="spin-articel list-style-fix">
                <div v-if="chapter.order">
                    <div>{{chapter.content}}</div>
                </div>
                <div v-else>
                    <div class="h1">{{book.name}}</div>
                    <div v-html="book_description_marked"></div>
                </div>
                <Spin size="large" fix v-if="spin_show"></Spin>
            </Col>
        </Row>
    </div>
</template>

<script>
import marked from 'marked'
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
    },
    book_description_marked() {
        if(this.book.description){
            return marked(this.book.description)
        }else{
            return ""
        }
    },
    spin_show() {
        return this.$store.state.spin_show
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
