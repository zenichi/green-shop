<template>
  <MainLayout>
    <h1>Welcome to Green E-Commerce</h1>
    <p>This is a test project under construction.</p>
    <div>Go to <router-link :to="{ name: 'products' }">products</router-link> page</div>
  </MainLayout>
</template>

<script>
import MainLayout from '@/components/layout/Layout.vue';

export default {
  name: 'HelloWorld',
  components: {
    MainLayout
  }
};
</script>

<style lang="scss" scoped>
h3 {
  margin: 200px 0 0;
}
a {
  color: #42b983;
}
</style>
