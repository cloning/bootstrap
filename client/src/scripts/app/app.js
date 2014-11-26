var bootstrap = bootstrap || {};

/*
    Main application class
*/
bootstrap.app = function() {

    var me = this;


    me.init = function() {    
        
        // Main framework element
        me.$container = $('#container');
        me.initApi();
        me.initControllers();
        me.initRouting();

    };

    me.initApi = function() {
        // Initialize api and set the handlers for
        // events
        me.api = new bootstrap.core.api();
        
        // Action to be taken when the api detects an
        // invalid token
        me.api.on('logged-out', function() {
                        
            var currentUrl = window.location.pathname;

            // Only redirect if we aren't currently on the login page
            if(currentUrl.indexOf("/auth/login") === -1) {
                me.router.routeTo("/auth/login?return=" + currentUrl);   
            }
        });

        // This tells the api where to look for any current tokens
        me.api.tokenProvider = me.getAuth;

    };

    me.initControllers = function() {
        // Initialize controllers
        me.homeController = new bootstrap.controllers.home({app : me});
        me.authController = new bootstrap.controllers.auth({app : me});
        me.dashboardController = new bootstrap.controllers.dashboard({app : me});
    };

    me.initRouting = function() {
        // Routing
        var routes = {
            "__home" : me.homeController.index,
            "__404" : me.homeController.error404,
            "auth" : {
                "login" : me.authController.login,
                "register" : me.authController.register
            },
            "dashboard" : {
                "index" : me.dashboardController.index
            }
        };

        // Create and initialize the router
        me.router = new bootstrap.core.router({
            routeMap : routes
        });
        
        me.router.init();
    };

    /*
        Replaces the visible content with the supplied view
        @view contains the html to be presented.
    */
    me.setView = function(view) {
        me.$container.html('');
        me.$container.append(view);
    };


    /*
        Sets the current authentication cookie
    */ 
    me.setAuth = function(token) {
        setCookie("token", token.token, token.expires);
    };

    /*
        Gets the current authentication token from cookie
    */
    me.getAuth = function(token) {
        return getCookie("token");
    };

    /*
        Verifies the existing token by 
        sending it to serverside.

        callback
            error
            isValid (bool)
    */
    me.verifyAuth = function(callback) {
        
        // Retrieve token from cookie and check 
        // that it exists.
        var token = me.getAuth();
        
        if(token === null)
        {
            callback(null, false);
        }

        // Execute server-side validation
        me.api.validateToken(callback);
    };

    /*
        Determines where to send the 
        user after a successful login
    */
    me.routeAfterLogin = function() {
        // Check for the return querystring and redirect here if present
        var returnUrl = querystring('return');
        if(returnUrl !== '') {
            me.router.routeTo(returnUrl);
        } else {
            me.router.routeTo('/dashboard');
        }    
    };

    return me;
};
