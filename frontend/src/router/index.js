import Vue from 'vue'
import Router from 'vue-router'
import books from '@/components/books'
import book from '@/components/book'
import topics from '@/components/topics'
import book_edit from '@/components/book_edit'
import chapter_edit from '@/components/chapter_edit'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: books 
    },
    {
      path: '/books',
      component: books 
    },
    {
      path: '/book/:id',
      component: book 
    },
    {
      path: '/book/:id/edit',
      component: book_edit 
    },
    {
      path: '/book/:id/add_chapter',
      component: chapter_edit
    },
    {
      path: '/topics',
      component: topics
    },
    {
      path: '/add_book',
      component: book_edit 
    }
  ]
})
