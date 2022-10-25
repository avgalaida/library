<template>
  <form @submit.prevent>
    <h3>Создание книги</h3>
    <my-input placeholder="Название" v-model="bookTitle" type="text"></my-input>
    <my-input placeholder="Авторы" v-model="bookAuthors" type="text"></my-input>
    <button class="createBtn" type="button" @click="createBook" style="margin-top: 20px;">Создать</button>
  </form>
</template>

<script>
import myInput from "@/components/UI/MyInput";
import book from "@/components/Book";

export default {
  data() {
    return {bookTitle: '', bookAuthors: ''}
  },
  methods: {
    createBook(){
      this.$store.dispatch('createBook', {
        title: this.bookTitle,
        authors: this.bookAuthors,
      });
      this.bookTitle = ''; this.bookAuthors= '';
    },
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