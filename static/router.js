app = angular.module('app.router', ['ngRoute'])

app.config(['$routeProvider', function($routeProvider) {
  console.log("configing")
//app.config(function($routeProvider) {
  $routeProvider
    .when('/about', {
      templateUrl: 'static/about.html',
      controller: 'exampleCtrl'
    });
}]);
