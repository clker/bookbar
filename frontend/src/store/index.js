import Vue from 'vue'
import Vuex from 'vuex'
import VueResource from 'vue-resource';
import actions from './actions'
import mutations from './mutations'


const state = {
  books: [],
  book: {},
  chapters: [],
  chapter: {},
  topics : []
};

export default new Vuex.Store({
  state,
  mutations,
  actions
})
