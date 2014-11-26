var bootstrap = bootstrap || {};
bootstrap.controllers = bootstrap.controllers || {};

bootstrap.controllers.auth = function(options) {
    var me = this;

    me.app = options.app;
    
    me.whenNotLoggedIn = function(callback) {
        me.app.verifyAuth(function(err, valid) {
            if(valid === false) {
                callback();
            } else {
                // TODO: Can we eraze /auth/login from history
                // before redirecting? 
                // Otherwise, this prevents back button 
                // when auto-login
                me.app.routeAfterLogin();
            }
        });
    };

    me.login = function() {
        me.whenNotLoggedIn(function() {
            var view = $(templatizer.auth.login());
            
            var loginForm = new bootstrap.components.forms.login({
                app : me.app,
                form : view.find('form')
            });

            me.app.setView(view);
        });
    };

    me.register = function() {
        var view = $(templatizer.auth.register());

        var registerForm = new bootstrap.components.forms.register({
            app : me.app,
            form : view.find('form')
        });
        
        me.app.setView(view);
    };

    return me;
};