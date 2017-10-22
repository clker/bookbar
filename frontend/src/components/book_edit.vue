<template>
    <div class="container">
        <Form :label-width="80" class="mt-md">
            <FormItem label="Title">
                <Input v-model="book.name"></Input>
            </FormItem>
            <FormItem label="Content">
                <Tabs value="edit" :animated="false">
                    <TabPane label="Edit" name="edit">
                        <Input v-model="book.description" 
                        type="textarea" :autosize="{minRows: 18}"></Input>
                    </TabPane>
                    <TabPane label="Preview" name="preview"><div class="list-style-fix" v-html="marked_text"></div></TabPane>
                </Tabs>
            </FormItem>
            <FormItem>
                <Button type="primary" @click="create_or_edit_book()">Sumbit</Button>
            </FormItem>
        </Form>
    </div>
</template>

<script>
import marked from 'marked'
export default {
  name : "book_edit",
  methods: {
      create_or_edit_book() {
          this.$store.dispatch('create_or_edit',this.book)
      }
  },
  //filters : {
  //    marked : this.$marked
  //},
  computed: {
    book() {
        if(this.$route.params.id)
            return this.$store.state.book
        else
            return {name:'',description:''}
    },
    marked_text() {
    //    console.log(this.$marked)
        return marked(this.book.description)
    }
  }
}
</script>
