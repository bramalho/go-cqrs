<template>
  <div class="container py-5">
    <div class="row mb-4">
      <h1 class="col-12">
        Todos
      </h1>
    </div>
    <div class="row">
      <Search class="col-12" />
      <List class="col-12" />
    </div>
  </div>
</template>

<script>
import List from '@/components/List';
import Search from '@/components/Search';

export default {
  name: 'App',
  components: {
    List,
    Search,
  },
};
</script>

<style lang="scss" scoped>
.container {
  max-width: 768px;
}
</style>
