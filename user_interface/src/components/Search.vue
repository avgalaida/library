<template>
  <div>
    <my-input @keyup="searchBooks" v-model.trim="query" class="form-control" placeholder="Найти..."></my-input>
    <div class="mt-4">
      <Book v-for="book in books" :key="book.id" :book="book" />
    </div>
  </div>
</template>

<script>
import myInput from "@/components/UI/MyInput";
import Book from '@/components/Book';
import { mapState } from 'vuex';
export default {
  data() {
    return {
      query: '',
    };
  },
  computed: mapState({
    books: (state) => state.searchResults,
  }),
  methods: {
    searchBooks() {
      // if (this.query != this.lastQuery) {
      //   this.$store.dispatch('searchBooks', this.query);
      //   this.lastQuery = this.query;
      // }
    },
  },
  components: {
    Book, myInput
  },
};
</script>
Footer
