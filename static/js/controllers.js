var controllers = angular.module('controllers', []);

controllers.controller('homeController', function($window, $scope){
    $scope.name = "test";
    

    //Event Handler
    $scope.nameBob = function(){
        console.log("Clicked!");
        $scope.name = "Bob";
    }


})