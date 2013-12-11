app = angular.module('app.factories', [])

app.factory('FetchJson', ['$http', (http) ->
  return {
      go: ->
        return http.get('/data')
      doit: (name, message) ->
        return http.post('/data', JSON.stringify({Name: name, Message: message}))
    }
])
