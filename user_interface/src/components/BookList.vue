<template>
  <div v-if="books.length > 0 " >
    <transition-group name="book-list">
      <book-item
          v-for="book in books"
          :key="book.createdAt"
          :book="book"
      />
    </transition-group>
  </div>
  <h4 v-else style="padding: 0 15px;">Книг нет</h4>
</template>

<script>
import { mapState } from 'vuex';
import bookItem from "@/components/Book";
export default {
  components: {bookItem},
  computed: mapState({books: state => state.books}),
}
</script>
