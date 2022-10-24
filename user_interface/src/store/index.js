import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';
import VueNativeSock from 'vue-native-websocket';

const BACKEND_URL = 'http://localhost:8080';
const PUSHER_URL = 'ws://localhost:8080/pusher';

const SET_BOOKS = 'SET_BOOKS';
const CREATE_BOOK = 'CREATE_BOOK';
const DELETE_BOOK = 'DELETE_BOOK';
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
            desc: message.desc,
            createdAt: message.createdAt }
          );
          break;
        case "delete_book":
          this.commit(DELETE_BOOK, {
            id: message.id,
            meta: message.meta,
            status: message.status,}
          );
      }
    },
    [SET_BOOKS](state, books) {
      state.books = books;
    },
    [CREATE_BOOK](state, book) {
      state.books = [book, ...state.books];
    },
    [DELETE_BOOK](state, b) {
      state.books.forEach((book)=>book.id===b.id?book.status=b.status:book.status=book.status)
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
      const { data } = await axios.post(`${BACKEND_URL}/books`, null, {
        params: {
          title: book.title,
          authors: book.authors,
          desc: book.desc
        },
      });
    },
    async deleteBook({ commit }, book) {
      const { data } = await axios.post(`${BACKEND_URL}/books`, null, {
        params: {
          id: book.id
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
