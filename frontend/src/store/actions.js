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
    console.log("4")
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
  get_book({commit}, id){
    Vue.http.get(api_url + 'book/' + id).then(response =>{
        commit(types.GET_BOOK,response.body)
    },response => {
    });
  },
  delete_book({commit},book){
    console.log(book)
    Vue.http.delete(api_url + 'book/' + book.ID).then(response =>{
        commit(types.DELETE_BOOK,book)
    },response => {
    });
  },
  add_chapter({ commit }, chapter, book_id) {
    Vue.http.post(api_url + 'book/' + book_id + '/chapter', chapter).then(response =>{
        commit(types.GET_CHAPTER,response.body)
        router.push({path:'/book/'+book_id})
    },response => {
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
