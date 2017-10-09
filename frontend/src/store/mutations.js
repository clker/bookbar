import * as types from './mutation-types'

export default {
    [types.GET_BOOKS] (state,books){
        state.books = books
    },
    
    [types.GET_BOOK] (state,book){
        state.book = book
    },
    
    [types.DELETE_BOOK] (state,book){
        var i = state.books.indexOf(book)
        if(i != -1)
            state.books.splice(i,1)
    },
    
    [types.GET_TOPICS] (state,topics){
        state.topics = topics 
    },

    [types.GET_CHAPTER] (state,chapter){
        state.chapter = chapter
    },
}

