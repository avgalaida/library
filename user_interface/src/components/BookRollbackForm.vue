<template>
  <div @submit.prevent>
    <h3 style="color: darkslategrey; text-align: center">Откат книги</h3>
    <form class="rollback">
      <div>
        <div><strong>Название:</strong>  {{newBookTitle}}</div>
        <div><strong>Авторы:</strong> {{newBookAuthors}}</div>
        <div><strong>Статус:</strong> {{newBookStatus}}</div>
    </div>
      <div>
        <select  @change="onChange($event)" v-model="selected" style="margin-left: 20px;font-family: Georgia, serif;color: darkslategrey;" >
          <option disabled value="">Ревизия</option>
          <option v-for="option in options" v-bind:value="option.value">
            {{ option}}
          </option>
        </select>
        <button class="createBtn" type="button" @click="rollback" style="margin-left: 20px;">Изменить</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newBookTitle: this.book.title,
      newBookAuthors: this.book.authors,
      newBookStatus: this.book.status,
      selected: '',
      options: Array.from({length: this.book.meta - 1}, (_, i) => this.book.meta - i - 1)
    }
  },
  props: ['book'],
  computed: {
    id() {return this.book.id},
    meta() {return this.book.meta},
    status() {return this.book.status},
    title() {return this.book.title},
    authors() {return this.book.authors},
    createdAt() {return this.book.createdAt}
  },
  methods: {
    onChange() {
      this.$store.dispatch('getVersion', {
        id: this.id,
        version: this.selected,
      });
      this.newBookTitle = this.$store.state.getVersionResult.title;
      this.newBookAuthors = this.$store.state.getVersionResult.authors;
      this.newBookStatus = this.$store.state.getVersionResult.status;
    },

    rollback() {

    }
  }
}
</script>

<style scoped>
.rollback {
  color: darkslategrey;
  display: flex;
  /*align-items: start;*/
}
</style>