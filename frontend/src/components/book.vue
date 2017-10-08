<template>
    <div>
        <div>
            <router-link :to="{path: '/book/' + book.ID + '/edit'}">
                Edit
            </router-link>
            <a href="#" v-on:click="delete_book">delete</a>
        </div>
        <div v-if="book.DeletedAt">
            The book is deleted
        </div>
        <div class="row" v-else>
            <div class="col-md-3">
                <router-link :to="{path: '/book/' + book.ID}">Cover</router-link>
            </div>
            <div class="col-md-9">
                <div>{{book.name}}</div>
                <div>{{book.description}}</div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name : "book",
  created() {
      this.$store.dispatch('get_book',this.$route.params.id)
  },
  computed: {
    book() {
        return this.$store.state.book
    }
  },
  methods : {
      delete_book(){
          this.$store.dispatch('delete_book',this.book)
      }
  }
}
</script>
