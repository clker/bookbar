<template>
    <div>
        <div>
            <router-link :to="{path: '/book/' + book.ID + '/edit'}">
                Edit
            </router-link>
            <router-link :to="{path: '/book/' + book.ID + '/add_chapter'}">
                Edit
            </router-link>
        </div>
        <div v-if="book.DeletedAt">
            The book is deleted
        </div>
        <div class="row" v-else>
            <div class="col-md-3">
                <router-link :to="{path: '/book/' + book.ID}">Cover</router-link>
                <div v-for="chapter in chapters">
                    <router-link :to="{path: '/book/' + book.ID + '/' + chapter.ID}">{{chapter.title}}</router-link>
                </div>
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
    },
    chapters() {
        return this.$store.state.chapters
    },
  },
}
</script>
