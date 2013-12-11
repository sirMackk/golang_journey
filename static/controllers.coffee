app = angular.module('app.controllers', [])

app.controller('exampleCtrl', ['$scope', 'FetchJson',  (sc, fetcher) ->
  sc.Items = [1, 2, 3, 4]
  fetcher.go().then((result) ->
    sc.Bombel = result.data["Name"]
  )

  sc.Name = ""
  sc.Message = ""

  sc.sendForm = =>
    console.log(fetcher)
    console.log("data submited")
    fetcher.doit(sc.Name, sc.Message).success((data, status) ->
      console.log("success")
      console.log(status)
    )


])
