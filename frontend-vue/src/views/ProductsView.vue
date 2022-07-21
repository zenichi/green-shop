<template>
  <MainLayout>
    <h2>Products listing page</h2>
    <div v-for="(product, index) in products" :key="index">
      <span>Go to product details: </span>
      <router-link :to="{ name: 'product', params: { id: product.id } }">
        {{ product.name }}
      </router-link>
    </div>
  </MainLayout>
</template>

<script>
import { getProducts } from '@/api';
import MainLayout from '@/components/layout/Layout.vue';

export default {
  components: {
    MainLayout
  },
  created() {
    this.loadProducts();
  },
  data: () => ({ products: [] }),
  methods: {
    loadProducts() {
      getProducts('USD')
        .then(r => {
          this.products = r.data;
        })
        .catch(e => {
          console.error(e);
        });
    }
  }
};
</script>
