import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

var apiUrl = 'http://localhost:3000'

export default new Vuex.Store({
  state: {
    currentUser: {}
  },
  mutations: {

  },
  actions: {
    createUser: function (context, payload) {
      axios.post(
        apiUrl + '/users/add', payload, {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          }
        }
      ).then((response) => {
        console.log(response)
      }).catch((error) => {
        console.error(error)
      })
    },
    authUser: function (context, payload) {
      axios.post(
        apiUrl + '/users/auth', payload, {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          }
        }
      ).then((response) => {
        console.log(response)
      }).catch((error) => {
        console.error(error)
      })
    }
  }
})
