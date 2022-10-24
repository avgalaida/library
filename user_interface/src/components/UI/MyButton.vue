<template>
  <button class="btn">
    <slot></slot>
  </button>
</template>

<script>
export default {name: "my-button"}
</script>

<style scoped>
.btn{
  background: mediumseagreen;
  transition: background 0.25s ease-in;
  border-radius: 4px;
  display: inline-block;
  border: none;
  padding: 8px 16px;
  margin: 0;
  text-decoration: none;
  color: #ffffff;
  font-size: 1rem;
  cursor: pointer;
  text-align: center;
  -webkit-appearance: none;
  -moz-appearance: none;
}
button:hover,
button:focus {
  transition: background 0.3s ease-out;
  background: darkslategrey;
}
button:focus {
  outline: 1px solid #fff;
  outline-offset: -4px;
}
</style>