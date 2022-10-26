<template>
  <input :value="value" @input="updateInput" class="input" type="text">
</template>

<script>
export default {
  name: "my-input",
  props: {value: [String, Number]},
  methods: {
    updateInput(event) {
      this.$emit('input', event.target.value)
    }
  }
}
</script>

<style scoped>
.input{
  border: 1px solid mediumseagreen;
  padding: 10px 15px;
  margin-top: 15px;
  border-radius: 4px;
}
.input:hover,
.input:focus {
  transition: background 0.3s ease-out;
  border-color: darkslategray;
  background: darkslategrey;
  color: white;
}
.input:focus {
  outline: 1px solid #fff;
  outline-offset: -4px;
}
</style>