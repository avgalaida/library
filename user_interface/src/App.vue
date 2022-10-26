<template>
  <div class="container py-5">
    <div class="row mb-4">
      <h1 class="col-12" style="text-align: center">БИБЛИОТЕКА</h1>
    </div>
<!--    <h4>{{$store.state.books}}</h4>-->
<!--    <h4>{{$store.state.getVersionResult}}</h4>-->
    <button class="createBtn" type="button" @click="showCreateDialog" style="font-size: 1rem;">Создать книгу</button>
    <h4 style="padding: 20px 0; text-align: center;" >Список книг</h4>
    <my-dialog v-model="createDialogVisible"><create-book-form/></my-dialog>
    <div class="row">
      <book-list  class="col"/>
<!--      <Search class="col" />-->
    </div>
  </div>
</template>

<script>

import bookList from "@/components/BookList";
import myDialog from "@/components/UI/MyDialog";
import createBookForm from "@/components/CreateBookForm";
import book from "@/components/Book";
export default {
  data(){
    return {createDialogVisible: false}
  },
  components: {
    book,
    bookList,
    myDialog,
    createBookForm,
  },
  methods: {
    showCreateDialog() {this.createDialogVisible = true},
  },
};

</script>

<style>
.container {
  color: darkslategrey;
  font-family: Georgia, serif;
    max-width: 720px;
}
.createBtn{
  background: mediumseagreen;
  color: white;
}
/*.bookBtns{*/
/*  background: white;*/
/*  color: mediumseagreen;*/
/*}*/
button{
  background: white;
  color: mediumseagreen;
  transition: background 0.25s ease-in;
  border-radius: 4px;
  border: none;
  padding: 8px 16px;
  font-size: 1rem;
  cursor: pointer;
  text-align: center;
}
button:hover,
button:focus {
  transition: background 0.3s ease-out;
  background: darkslategrey;
  color: white;
}
button:focus {
  outline: 1px solid #fff;
  outline-offset: -4px;
}
</style>