<template>
  <div>
    <input @keyup="searchTodos" v-model.trim="query" class="form-control" placeholder="Search...">
    <div class="mt-4">
      <Todo v-for="todo in todos" :key="todo.id" :todo="todo" />
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import Todo from '@/components/Todo';

export default {
  data() {
    return {
      query: '',
    };
  },
  computed: mapState({
    todos: (state) => state.searchResults,
  }),
  methods: {
    searchTodos() {
      if (this.query != this.lastQuery) {
        this.$store.dispatch('searchTodos', this.query);
        this.lastQuery = this.query;
      }
    },
  },
  components: {
    Todo,
  },
};
</script>
