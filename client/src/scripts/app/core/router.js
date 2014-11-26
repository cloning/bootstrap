var bootstrap = bootstrap || {};
bootstrap.core = bootstrap.core || {};

bootstrap.core.router = function(options) {
    var me = this;

    me.routeMap = options.routeMap;

    me.init = function(routeMap) {
        me.route();
        
        // Listen to link clicks
        $('body').on('click', 'a', function(e) {
            e.preventDefault();
            me.routeTo($(this).attr('href'));
        });

        $(window).on('popstate', function() {
            me.route();
        });
    };

    me.routeTo = function(href) {
        history.pushState(null, null, href);
        me.route();
    };

    me.route = function() {        

        var path = window.location.pathname;        
        
        var split = splitRemoveEmpty(path, "/");

        if(split.length === 0) {
            me.routeMap.__home();
        }

        var prev = me.routeMap;
        
        for(var i = 0; i < split.length; i++) {
            var current = split[i];
            var currentEndpoint = prev[current];
            
            // we are at a defined endpoitn
            if(isFunction(currentEndpoint)) {
                currentEndpoint();
            }
            // Special case: Allow for index
            // endpoints
            else if (
                    (i === split.length - 1) && 
                    (typeof(currentEndpoint) !== 'undefined') && 
                    (isFunction(currentEndpoint.index))) {
                currentEndpoint.index();
            }

            // We are at a object, go deeper
            else if (isObject(currentEndpoint)) {
                prev = currentEndpoint;
            } 

            // Page cannot be found
            else {
                me.routeMap.__404();
                break;
            }
        }
    };

    return me;
};