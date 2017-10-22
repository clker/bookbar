<template>
    <div class="container">
        <Form :label-width="80" class="mt-md">
            <FormItem label="Title">
                <Input v-model="chapter.title"></Input>
            </FormItem>
            <FormItem label="Content">
                <Tabs value="edit" :animated="false">
                    <TabPane label="Edit" name="edit">
                        <Input v-model="chapter.content" 
                        type="textarea" :autosize="{minRows: 18}"></Input>
                    </TabPane>
                    <TabPane label="Preview" name="preview"><div class="list-style-fix" v-html="marked_text"></div></TabPane>
                </Tabs>
            </FormItem>
            <FormItem>
                <Button type="primary" @click="create_or_edit_chapter()">Sumbit</Button>
            </FormItem>
        </Form>
    </div>
</template>

<script>
import marked from 'marked'
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
            return {title:"",content:""}
    },
    book(){
        return this.$store.state.book
    },
    marked_text() {
    //    console.log(this.$marked)
        return marked(this.chapter.content)
    }
  }
}
</script>
