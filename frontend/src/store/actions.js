import VueResource from 'vue-resource';
import * as types from './mutation-types'
import Vue from 'vue'
import Vuex from 'vuex'

var api_url = 'http://localhost:8000/'
Vue.use(Vuex);
Vue.use(VueResource);


export default {
  get_books({ commit },books ) {
    Vue.http.get(api_url + 'books',{'headers':{'Origin':'*'}}).then(response =>{
            books = response.body;
            commit(types.GET_BOOKS,books)
    },response => {
        console.log(response)
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
  add_book({ commit }, book) {
    console.log(book)
    Vue.http.post(api_url + 'book',book).then(response =>{
        commit(types.ADD_BOOK,1)
    },response => {
    });
  }
}
