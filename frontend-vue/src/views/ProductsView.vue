<template>
  <MainLayout>
    <h2>Products listing page</h2>
    <div>
      Go to
      <router-link :to="{ name: 'product', params: { id: '2' } }"> product </router-link>
      details
    </div>
  </MainLayout>
</template>

<script>
import MainLayout from '@/components/layout/Layout.vue';

export default {
  components: {
    MainLayout
  }
};
</script>
