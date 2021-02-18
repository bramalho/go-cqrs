<template>
  <div>
    <form v-on:submit.prevent="createTodo">
      <div class="input-group">
        <input v-model.trim="todoBody" type="text" class="form-control" placeholder="New todo...">
        <div class="input-group-append">
          <button class="btn btn-primary" type="submit">Add</button>
        </div>
      </div>
    </form>

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
      todoBody: '',
    };
  },
  computed: mapState({
    todos: (state) => state.todos,
  }),
  methods: {
    createTodo() {
      if (this.todoBody.length != 0) {
        this.$store.dispatch('createTodo', { body: this.todoBody });
        this.todoBody = '';
      }
    },
  },
  components: {
    Todo,
  },
};
</script>
