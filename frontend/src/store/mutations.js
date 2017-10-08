import * as types from './mutation-types'

export default {
    [types.GET_BOOKS] (state,books){
        state.books = books
    },
    
    [types.GET_BOOK] (state,book){
        state.book = book
    },
    
    [types.GET_TOPICS] (state,topics){
        state.topics = topics 
    },

    [types.GET_CHAPTER] (state,chapter){
        state.chapter = chapter
    },
}

