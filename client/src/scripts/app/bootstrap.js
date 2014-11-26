
/* 
    Main entry point for application
*/

$(document).ready(function() {
    var mainContainer = $('#container');
    var appInstance = new bootstrap.app(mainContainer);
    appInstance.init();
});