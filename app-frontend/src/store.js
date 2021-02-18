import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';
import VueNativeSock from 'vue-native-websocket';

const BACKEND_URL = 'http://localhost:8080';
const PUSHER_URL = 'ws://localhost:8080/pusher';

const SET_TODOS = 'SET_TODOS';
const CREATE_TODO = 'CREATE_TODO';
const SEARCH_SUCCESS = 'SEARCH_SUCCESS';
const SEARCH_ERROR = 'SEARCH_ERROR';

const MESSAGE_TODO_CREATED = 1;

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        todos: [],
        searchResults: [],
    },
    mutations: {
        SOCKET_ONOPEN(state, event) {
        },
        SOCKET_ONCLOSE(state, event) {
        },
        SOCKET_ONERROR(state, event) {
            console.error(event);
        },
        SOCKET_ONMESSAGE(state, message) {
            switch (message.kind) {
                case MESSAGE_TODO_CREATED:
                    this.commit(CREATE_TODO, { id: message.id, body: message.body });
            }
        },
        [SET_TODOS](state, todos) {
            state.todos = todos;
        },
        [CREATE_TODO](state, todo) {
            state.todos = [todo, ...state.todos];
        },
        [SEARCH_SUCCESS](state, todos) {
            state.searchResults = todos;
        },
        [SEARCH_ERROR](state) {
            state.searchResults = [];
        },
    },
    actions: {
        getTodos({ commit }) {
            axios
                .get(`${BACKEND_URL}/todos`)
                .then(({ data }) => {
                    commit(SET_TODOS, data);
                })
                .catch((err) => console.error(err));
        },
        async createTodo({ commit }, todo) {
            const { data } = await axios.post(`${BACKEND_URL}/todos`, null, {
                params: {
                    body: todo.body,
                },
            });
        },
        async searchTodos({ commit }, query) {
            if (query.length === 0) {
                commit(SEARCH_SUCCESS, []);
                return;
            }
            axios
                .get(`${BACKEND_URL}/search`, {
                    params: { query },
                })
                .then(({ data }) => commit(SEARCH_SUCCESS, data))
                .catch((err) => {
                    console.error(err);
                    commit(SEARCH_ERROR);
                });
        },
    },
});

Vue.use(VueNativeSock, PUSHER_URL, { store, format: 'json' });

store.dispatch('getTodos');

export default store;
