import Vue from 'vue'
import Router from 'vue-router'
import books from '@/components/books'
import topics from '@/components/topics'
import book_edit from '@/components/book_edit'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: books 
    },
    {
      path: '/books',
      name: 'books',
      component: books 
    },
    {
      path: '/topics',
      name: 'topics',
      component: topics
    },
    {
      path: '/add_book',
      name: 'book_edit',
      component: book_edit 
    }
  ]
})
