import * as types from './mutation-types'

export default {
    [types.GET_BOOKS] (state,books){
        state.books = books
    },
    
    [types.GET_BOOK] (state,book_chapters){
        state.book = book_chapters['book']
        state.chapters = book_chapters['chapters']
        state.chapter = book_chapters['chapter']
    },
    
    [types.DELETE_BOOK] (state,book){
        var i = state.books.indexOf(book)
        if(i != -1)
            state.books.splice(i,1)
    },
    
    [types.DELETE_CHAPTER] (state,chapter){
        var i = state.chapters.indexOf(chapter.title)
        if(i != -1)
            state.chapters.splice(i,1)
        state.chapter = {}
    },
    
    [types.GET_TOPICS] (state,topics){
        state.topics = topics 
    },

    [types.GET_CHAPTER] (state,chapter){
        state.chapter = chapter
    },

    [types.SHOW_SPIN] (state){
        state.spin_show = true
    },

    [types.HIDE_SPIN] (state){
        state.spin_show = false
    },
}

