<template>
  <form @submit.prevent>
    <h3 style="color: darkslategrey">Изменение книги</h3>
    <my-input placeholder="Название" v-model="bookTitle" type="text"></my-input>
    <my-input placeholder="Авторы" v-model="bookAuthors" type="text"></my-input>
    <button class="createBtn" type="button" @click="changeBook" style="margin-top: 20px;">Изменить</button>
  </form>
</template>

<script>
import myInput from "@/components/UI/MyInput";
import book from "@/components/Book";

export default {
  data() {
    return {bookTitle: this.book.title, bookAuthors: this.book.authors}
  },
  props: ['book'],
  computed:{
    bookTitle() {return this.book.title},
    bookAuthors() {return this.book.authors}
  },
  methods: {
    changeBook(){
      if (this.bookTitle !== this.book.title()) {
        this.$store.dispatch('changeTitle', {
              title: this.bookTitle,
        });
      }
      if (this.bookAuthors !== this.book.authors()) {
        this.$store.dispatch('changeAuthors', {
          authors: this.book.authors,
        });
      }
    }
  },
  components: {book, myInput}
}
</script>

<style scoped>
form {
  display: flex;
  flex-direction: column;
}
</style>