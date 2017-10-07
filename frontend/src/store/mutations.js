import * as types from './mutation-types'

export default {
    [types.GET_BOOKS] (state,books){
        state.books = books
    },
    
    [types.GET_TOPICS] (state,topics){
        state.topics = topics 
    },

    [types.ADD_BOOK] (state, add_book_done){
        state.add_book_done = add_book_done
    }

}

