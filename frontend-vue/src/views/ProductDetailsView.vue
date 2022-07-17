<template>
  <MainLayout>
    <h2>Product page - currency: {{ currency }}</h2>
    <div>ID: {{ product.id }}</div>
    <div>Name: {{ product.name }}</div>
    <div>Description: {{ product.description }}</div>
    <div>Price: {{ priceText }}</div>
    <div>ExternalId: {{ product.externalId }}</div>
  </MainLayout>
</template>

<script>
import { getProductById } from '@/api';
import MainLayout from '@/components/layout/Layout.vue';

export default {
  components: {
    MainLayout
  },
  props: {
    id: { type: String, required: true }
  },
  computed: {
    priceText() {
      return `${Math.round(this.product.price * 100) / 100} ${this.currency}`;
    }
  },
  created() {
    this.loadProduct();
  },
  data: () => ({ product: {}, currency: 'EUR' }),
  methods: {
    loadProduct() {
      getProductById(this.id, this.currency)
        .then(r => {
          this.product = r.data;
        })
        .catch(e => {
          console.error(e);
        });
    }
  }
};
</script>
