<template>
    <div>
        <ul class="list-group">
            <li class="list-group-item" v-for="(topic,idx) in topics">
                {{topic.title}}
            </li>
        </ul>
    </div>
</template>

<script>
export default {
  name : "topics",
  computed: {
    topics() {
        this.$store.dispatch('get_topics')
        return this.$store.state.topics
    }
  }
}
</script>
