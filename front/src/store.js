import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import Methods from './methods'

Vue.use(Vuex)

export const STORE = new Vuex.Store({
  strict: true,
  state: {
    authStatus: false,
    currentUser: {},
    artists: []
  },

  mutations: {
    setCurrentUser: (state, data) => {
      state.currentUser = data
    },

    setAuthStatus: (state, status) => {
      state.authStatus = status

      if (!status) {
        state.currentUser = {}
        document.cookie = deleteCookie('auth_key')
      }
    },

    setArtists: (state, data) => {
      state.artists = data
    }
  },

  getters: {
    getCurrentUser: (state) => {
      console.log('get user', state)
      return state.currentUser
    },

    getAuthStatus: (state) => {
      return state.authStatus
    },

    getArtists: (state) => {
      console.log(state)
      return state.artists
    }
  },

  actions: {
    createUser: function (context, payload) {
      payload.url = '/users/add'
      payload.method = 'POST'
      payload.action = (res) => {
        console.log(res)
        // TODO: add success message
      }
      payload.error = (err) => {
        console.error(err)
        // TODO: add error message
      }
    },

    authUser: function (context, payload) {
      payload.url = '/users/auth'
      payload.method = 'POST'

      payload.action = function(res) {
        if (typeof res.Session !== 'undefined' && res.Session.length > 10) {
          context.commit(Methods.SET_CURRENT_USER, res)
          context.commit(Methods.SET_AUTH_STATUS, true)
          document.cookie = 'auth_key=' + res.Session

          if (typeof res.DateCreated !== 'undefined' && typeof res.UserId !== 'undefined') {
            localStorage.setItem('userId', res.UserId)
            localStorage.setItem('dateCreated', res.DateCreated)
          }

          // TODO: rerouting
        } else {
          console.error('Error' + res)
        }
      }

      payload.error = (err) => {
        console.error(err)
      }

      this.dispatch('request', payload)
    },

    logout: (context) => {
      context.commit(Methods.LOGOUT, false)
    },

    loadArtists: (context, payload = {}) => {
      payload.url = '/artists/list'
      payload.method = 'GET'
      payload.action = (res) => {
        console.log(res)
        context.commit(Methods.SET_ARTISTS, res)
      }
      payload.error = (err) => {
        console.error(err)
      }
      context.dispatch('request', payload)
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
      const error = payload.error
      const url = payload.url
      delete payload.action
      delete payload.url

      switch (payload.method) {
        case 'POST':
          axios.post(
            Methods.API_URL + url, payload, {
              headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
              }
            }
            ).then((response) => {
              if (typeof response.data.test !== 'undefined') {
                console.log('Test:', response)
              }
              else if (typeof response.logout !== 'undefined' && response.logout === true) {
                context.commit(Methods.LOGOUT);
              }
              else if (typeof response.data.result !== 'undefined') {
                func(response.data.result)
              }
              else if (typeof response.data.error !== 'undefined') {
                error(response.data.error)
              }
              else {
                error('Answer error')
              }

              // console.log('Response:', response)
              console.log('Data:', response.data)
            }).catch((err) => {
              error(err)
            })
          break

        case 'GET':
          axios.get(Methods.API_URL + url, payload/*, {
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded'
            }
          }*/)
            .then((response) => {
              func(response.data.result)
            })
            .catch((err) => {
              error(err)
            })

          break
      }

    }
  }
})

function getCookie(key) {
  let value = "; " + document.cookie
  let parts = value.split("; " + key + "=")
  return (parts.length === 2) ? parts.pop().split(";").shift() : null
}
