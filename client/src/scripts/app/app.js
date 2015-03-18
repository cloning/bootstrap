ns('marginalen.core').router = function(options) {
    var me = this;

    me.routeMap = options.routeMap;

    me.init = function(routeMap) {
        me.route();
        
        // Listen to link clicks
        $('body').on('click', 'a', function(e) {
            var link = $(this);
            
            if(link.attr('target') === '_blank') return;
            // Check if link is external
            var thishost = window.location.hostname + (window.location.port ? ':' + window.location.port : '');
            if(thishost !== link.context.host) return;
            
            e.preventDefault();

            me.routeTo(link.attr('href'));
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
        
        var getVariableEndpoint = function(list) {
            for(var key in list) {
                if(key.indexOf(':') === 0) {
                    return list[key];
                }
            }
            return null;
        };
        
        var split = splitRemoveEmpty(path, "/");

        if(split.length === 0) {
            me.routeMap.__home();
        }
        var variables = [];
        var prev = me.routeMap;
        
        for(var i = 0; i < split.length; i++) {
            var current = split[i];
            var currentEndpoint = prev[current];
            var variableEndpoint = getVariableEndpoint(prev);
            // we are at a defined endpoitn
            if(isFunction(currentEndpoint)) {
                currentEndpoint.apply(null, variables);
            }
            // Special case: Allow for index
            // endpoints
            else if (
                    (i === split.length - 1) && 
                    (typeof(currentEndpoint) !== 'undefined') && 
                    (isFunction(currentEndpoint.index))) {
                currentEndpoint.index.apply(null, variables);
            }

            // We are at a object, go deeper
            else if (isObject(currentEndpoint)) {
                prev = currentEndpoint;
            } 

            else if (variableEndpoint !== null) {
                variables.push(current);
                if(i === split.length - 1) {
                    if(isFunction(variableEndpoint.index)) {
                        variableEndpoint.index.apply(null, variables);
                        break;
                    } 
                }
                else {
                     prev = variableEndpoint;
                }
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