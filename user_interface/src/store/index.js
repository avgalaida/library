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
const CHANGE_TITLE = 'CHANGE_TITLE';
const CHANGE_AUTHORS = 'CHANGE_AUTHORS';

const GET_VERSION_SUCCESS = 'GET_VERSION_SUCCESS';

const SEARCH_SUCCESS = 'SEARCH_SUCCESS';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    books: [
      {status: "Доступна", title: "наме", authors: "офтор",createdAt: "1", meta: "6"}
    ],
    searchResults: [],
    getVersionResult: [{title:"",authors:"",status:""}]
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
        case "книга.создана":
          this.commit(CREATE_BOOK, {
            id: message.id,
            status: message.status,
            title: message.title,
            authors: message.authors,
            createdAt: message.createdAt }
          );
          break;
        case "книга.удалена":
          this.commit(DELETE_BOOK, {
            id: message.id,
            }
          );
          break;
        case "книга.восстановлена":
          this.commit(RESTORE_BOOK, {
            id: message.id,
            status: message.status,}
          );
          break;
        case "название.изменено":
          this.commit(CHANGE_TITLE, {
            id: message.id,
            title: message.title,}
          );
          break;
        case "авторство.изменено":
          this.commit(CHANGE_AUTHORS, {
            id: message.id,
            authors: message.authors,}
          );
          break;
      }
    },
    [SET_BOOKS](state, books) {
      state.books = books;
    },
    [CREATE_BOOK](state, book) {
      book.meta = 1
      state.books = [book, ...state.books];
    },
    [DELETE_BOOK](state, b) {
      let i = findIndex(state,b.id)
      state.books.at(i).status = "Недоступна"
      state.books.at(i).meta = state.books.at(i).meta+1
    },
    [RESTORE_BOOK](state, b) {
      let i = findIndex(state,b.id)
      state.books.at(i).status = b.status
      state.books.at(i).meta = state.books.at(i).meta+1
    },
    [CHANGE_TITLE](state, b) {
      let i = findIndex(state,b.id)
      state.books.at(i).title = b.title
      state.books.at(i).meta = state.books.at(i).meta+1
    },
    [CHANGE_AUTHORS](state, b) {
      let i = findIndex(state,b.id)
      state.books.at(i).authors = b.authors
      state.books.at(i).meta = state.books.at(i).meta+1
    },
    [SEARCH_SUCCESS](state, books) {
      state.searchResults = books;
    },
    [GET_VERSION_SUCCESS](state, data) {
      state.getVersionResult = [data]
    }
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
    async getVersion({ commit }, query) {
      axios
          .get(`${BACKEND_URL}/books`, {
            params: {
              id: query.id,
              version: query.version
            },
          })
          .then(({ data }) => commit(GET_VERSION_SUCCESS, data))
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
    async changeTitle({ commit }, book) {
      await axios.post(`${BACKEND_URL}/books`, null, {
        params: {
          id: book.id,
          title: book.title
        },
      });
    },
    async changeAuthors({ commit }, book) {
      await axios.post(`${BACKEND_URL}/books`, null, {
        params: {
          id: book.id,
          authors: book.authors
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

function findIndex(state, id) {
  let index; let i;
  for (i = 0;i < state.books.length; ++i) {
    if (state.books.at(i).id === id) {
      index = i;
      break;
    }
  }
  return index;
}

Vue.use(VueNativeSock, PUSHER_URL, { store, format: 'json' });

store.dispatch('getBooks');

export default store;
