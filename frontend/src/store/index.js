import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        selectedAudio: {}
    },
    mutations: {
        setSelectedAudio(state, val) {
            state.selectedAudio = val;
        }
    },
    actions: {},
})
