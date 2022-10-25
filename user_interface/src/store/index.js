import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';
import VueNativeSock from 'vue-native-websocket';

const BACKEND_URL = 'http://localhost:8080';
const PUSHER_URL = 'ws://localhost:8080/pusher';

const SET_BOOKS = 'SET_BOOKS';
const CREATE_BOOK = 'CREATE_BOOK';
const DELETE_BOOK = 'DELETE_BOOK';
const RESTORE_BOOK = 'RESTORE_BOOK';
const SEARCH_SUCCESS = 'SEARCH_SUCCESS';
const SEARCH_ERROR = 'SEARCH_ERROR';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    books: [],
    searchResults: [],
  },
  mutations: {
    SOCKET_ONOPEN(state, event) {
    },
    SOCKET_ONCLOSE(state, event) {
    },
    SOCKET_ONERROR(state, event) {
      console.error(event);
    },
    SOCKET_ONMESSAGE(state, message) {
      switch (message.type) {
        case "create_book":
          this.commit(CREATE_BOOK, {
            id: message.id,
            meta: message.meta,
            status: message.status,
            title: message.title,
            authors: message.authors,
            createdAt: message.createdAt }
          );
          break;
        case "delete_book":
          this.commit(DELETE_BOOK, {
            id: message.id,
            meta: message.meta,
            }
          );
          break;
        case "restore_book":
          this.commit(RESTORE_BOOK, {
            id: message.id,
            meta: message.meta,
            status: message.status,}
          );
          break;
      }
    },
    [SET_BOOKS](state, books) {
      state.books = books;
    },
    [CREATE_BOOK](state, book) {
      state.books = [book, ...state.books];
    },
    [DELETE_BOOK](state, b) {
      let i;
      let index;
      for (index = 0; index < state.books.length; ++index) {
        if (state.books.at(index).id === b.id) {
          i = index;
          break;
        }
      }
      state.books.at(i).status = "Недоступна"
      state.books.at(i).meta = b.meta
    },
    [RESTORE_BOOK](state, b) {
      let i;
      let index;
      for (index = 0; index < state.books.length; ++index) {
        if (state.books.at(index).id === b.id) {
          i = index;
          break;
        }
      }
      state.books.at(i).status = b.status
      state.books.at(i).meta = b.meta
    },
    [SEARCH_SUCCESS](state, books) {
      state.searchResults = books;
    },
    [SEARCH_ERROR](state) {
      state.searchResults = [];
    },
  },
  actions: {
    getBooks({ commit }) {
      axios
          .get(`${BACKEND_URL}/books`)
          .then(({ data }) => {
            commit(SET_BOOKS, data);
          })
          .catch((err) => console.error(err));
    },
    async createBook({ commit }, book) {
      await axios.post(`${BACKEND_URL}/books`, null, {
        params: {
          title: book.title,
          authors: book.authors,
        },
      });
    },
    async deleteBook({ commit }, book) {
      await axios.post(`${BACKEND_URL}/books`, null, {
        params: {
          id: book.id
        },
      });
    },
    async restoreBook({ commit }, book) {
      await axios.post(`${BACKEND_URL}/books`, null, {
        params: {
          id: book.id,
          status: book.status
        },
      });
    },
    async searchBooks({ commit }, query) {
      if (query.length === 0) {
        commit(SEARCH_SUCCESS, []);
        return;
      }
      axios
          .get(`${BACKEND_URL}/search`, {
            params: { query },
          })
          .then(({ data }) => commit(SEARCH_SUCCESS, data))
          .catch((err) => {
            console.error(err);
            commit(SEARCH_ERROR);
          });
    },
  },
});

Vue.use(VueNativeSock, PUSHER_URL, { store, format: 'json' });

store.dispatch('getBooks');

export default store;
