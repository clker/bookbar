import VueResource from 'vue-resource';
import * as types from './mutation-types'
import Vue from 'vue'
import Vuex from 'vuex'
import router from '../router'

var api_url = 'http://localhost:8000/'
Vue.use(Vuex);
Vue.use(VueResource);


export default {

  get_books({ commit },books ) {
    Vue.http.get(api_url + 'books').then(response =>{
        books = response.body;
        commit(types.GET_BOOKS,books)
    },response => {
    });
    //commit(types.GET_BOOKS,[])
  },

  get_topics({ commit }, topics) {
    topics = [
        {'title':'DPI tutorials'},
        {'title':'VPI tutorials'}
    ]
    commit(types.GET_TOPICS,topics)
  },

  create_or_edit({ commit }, book) {
      var url = api_url + 'book'
      Vue.http.post(url,book).then(response =>{
          router.push({path:'/book/' + response.body.ID})
      },response => {
      });
  },

  get_book({commit}, url){
    Vue.http.get(api_url + url).then(response =>{
        commit(types.GET_BOOK,response.body)
    },response => {
    });
  },

  delete_book({commit},book){
    Vue.http.delete(api_url + 'book/' + book.ID).then(response =>{
        commit(types.DELETE_BOOK,book)
    },response => {
    });
  },

  create_or_edit_chapter({ commit }, chapter_and_book) {
    var book_id = chapter_and_book['book_id']
    var chapter = chapter_and_book['chapter']
    Vue.http.post(api_url + 'book/' + book_id + '/chapter', chapter).then(response =>{
        commit(types.GET_CHAPTER,response.body)
        router.push({path:'/book/'+book_id})
    },response => {
    });
  },

  delete_chapter({commit},chapter){
      Vue.http.delete(api_url + 'ch/' + chapter.ID).then(response => {
          commit(types.DELETE_CHAPTER,chapter)
          router.push({path:chapter.book_id})
      },response =>{
      });
  },

  create_user({ commit }, user) {
    Vue.http.post(api_url + 'user',user).then(response =>{
        //commit(types.ADD_BOOK,book)
        router.push({path:'/'})
    },response => {
    });
  },
}
