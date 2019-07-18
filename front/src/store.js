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
      // console.log(sendRequest('/users/auth', payload))
      payload.url = '/users/auth'
      payload.action = function(res) {
        if (typeof res.Session !== 'undefined' && res.Session.length > 10) {
          document.cookie = 'auth_key=' + res.Session
        } else {
          console.log('WTF:' + res)
        }
      }

      this.dispatch('request', payload)
    },
    request: function (context, payload) {
      if (typeof payload.url === 'undefined') {
        console.error('Url not found')
        return;
      }
      let key = getCookie('auth_key')
      if (key != null && key.length > 0) {
        payload.auth = key
      }
      const func = payload.action
      const url = payload.url
      delete payload.action
      delete payload.url

      axios.post(
        apiUrl + url, payload, {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          }
        }
      ).then((response) => {
        if (typeof response.data.test !== 'undefined') {
          console.log('Test:', response)
        } else if (typeof response.logout !== 'undefined' && response.logout === true) {
          logout()
        } else if (typeof response.data.result !== 'undefined') {
          func(response.data.result)
        } else if (typeof response.data.error !== 'undefined') {
          console.error(response.data.error)
        } else {
          console.error('Answer error')
        }

        console.log('Response:', response)
      }).catch((error) => {
        console.error(error)
      })
    }
  }
})

function getCookie(name) {
  var value = "; " + document.cookie
  var parts = value.split("; " + name + "=")
  return (parts.length === 2) ? parts.pop().split(";").shift() : null
}
