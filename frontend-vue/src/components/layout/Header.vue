<template>
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link :to="{ name: 'products' }">Products</router-link>
  </nav>
</template>

<script>
export default {
  name: 'main-header'
};
</script>
