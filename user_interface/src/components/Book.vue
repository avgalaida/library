<template>
  <div v-if="status === 'Доступна'" class="book" @submit.prevent>
    <div>
      <div><strong>Название:</strong> {{title}}</div>
      <div><strong>Авторы:</strong> {{authors}}</div>
      <div><strong>Ревизия:</strong> {{meta}}</div>
    </div>
    <my-dialog v-model="this.$store.state.changeDialogVisible"><change-book-form :book="this.book"/></my-dialog>
    <my-dialog v-model="this.$store.state.bookRollbackDialogVisible"><book-rollback-form :book="this.book"/></my-dialog>
    <div>
      <button @click="showRollbackDialog" type="button" style="margin-right: 10px;">Откат</button>
      <button @click="showChangeDialog" type="button" style="margin-right: 10px;">Изменить</button>
      <button @click="deleteBook">Удалить</button>
    </div>
  </div>

  <div v-else-if="status === 'Недоступна'" class="deletedBook" @submit.prevent>
    <div>
      <div><strong>Название:</strong> {{title}}</div>
      <div><strong>Авторы:</strong> {{authors}}</div>
      <div><strong>Ревизия:</strong> {{meta}}</div>
    </div>
    <div class="book__btn">
      <button class="restoreBtn" @click="restoreBook">Восстановить</button>
    </div>
  </div>
</template>

<script>
import changeBookForm from "@/components/ChangeBookForm";
import bookRollbackForm from "@/components/BookRollbackForm";
import myDialog from "@/components/UI/MyDialog";
export default {
  props: ['book'],
  computed: {
    id() {return this.book.id},
    meta() {return this.book.meta},
    status() {return this.book.status},
    title() {return this.book.title},
    authors() {return this.book.authors},
    createdAt() {return this.book.createdAt}
  },
  methods:{
    deleteBook(){
      this.$store.dispatch('deleteBook', {id: this.id});
    },
    restoreBook(){
      this.$store.dispatch('restoreBook', {id: this.id, status: "Доступна"});
    },
    showChangeDialog(){this.$store.state.changeDialogVisible = true;},
    showRollbackDialog(){this.$store.state.bookRollbackDialogVisible = true;}
  },
  components: {changeBookForm, bookRollbackForm, myDialog}
}
</script>

<style scoped>
.restoreBtn{
  color: darkslategrey;
  border: white;
}
.restoreBtn:hover,
.restoreBtn:focus {
  transition: background 0.3s ease-out;
  background: mediumseagreen;
  color: white;
}
.restoreBtn:focus {
  outline: 1px solid #fff;
  outline-offset: -4px;
}
.book{
  background: mediumseagreen;
  color: white;
  padding: 15px;
  border: 1px solid mediumseagreen;
  margin-top: 15px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-radius: 4px;
}
.deletedBook{
  background: darkslategrey;
  color: white;
  padding: 15px;
  border: 1px solid gray;
  margin-top: 15px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-radius: 4px;
}
</style>