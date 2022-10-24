<template>
  <div v-if="status === 'Доступна'" class="book" @submit.prevent>
    <div>
      <div><strong>Название:</strong> {{title}}</div>
      <div><strong>Авторы:</strong> {{authors}}</div>
      <div><strong>Описание:</strong> {{desc}}</div>
      <div><strong>Ревизия:</strong> {{meta}}</div>
    </div>
    <div>
      <button @click="changeBook" style="margin-right: 10px;">Изменить</button>
      <button @click="deleteBook">Удалить</button>
    </div>
  </div>

  <div v-else class="deletedBook" @submit.prevent>
    <div>
      <div><strong>Название:</strong> {{title}}</div>
      <div><strong>Авторы:</strong> {{authors}}</div>
      <div><strong>Описание:</strong> {{desc}}</div>
      <div><strong>Ревизия:</strong> {{meta}}</div>
    </div>
    <div class="book__btn">
      <button class="restoreBtn" @click="restoreBook">Восстановить</button>
    </div>
  </div>
</template>

<script>

export default {
  props: ['book'],
  data(){
    return {isAvailable: true}
  },
  computed: {
    id() {return this.book.id},
    meta() {return this.book.meta},
    status() {return this.book.status},
    title() {return this.book.title},
    authors() {return this.book.authors},
    desc() {return this.book.desc},
    createdAt() {return this.book.createdAt}
  },
  methods:{
    deleteBook(){
      this.$store.dispatch('deleteBook', {id: this.id});
    },
    restoreBook(){
      // this.$store.dispatch('deleteBook', {id: "2"});
    },
    changeBook(){

    }
  },
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