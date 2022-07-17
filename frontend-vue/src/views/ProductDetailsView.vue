<template>
  <MainLayout>
    <h2>Product page</h2>
    <div>ID: {{ id }}</div>
  </MainLayout>
</template>

<script>
import MainLayout from '@/components/layout/Layout.vue';

export default {
  components: {
    MainLayout
  },
  props: {
    id: { type: String, required: true }
  }
};
</script>
