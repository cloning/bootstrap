var bootstrap = bootstrap || {};
bootstrap.core = bootstrap.core || {};

bootstrap.core.api = function() {
    var me = this;

    me.baseUrl = "http://localhost:4000/";

    me.eventHandlers = {};

    me.on = function(eventName, eventHandler) {
        if(!me.eventHandlers.hasOwnProperty(eventName)) {
            me.eventHandlers[eventName] = [];
        }

        me.eventHandlers[eventName].push(eventHandler);
    };

    me.trigger = function(eventName) {
        if(!me.eventHandlers.hasOwnProperty(eventName)) return;
        for(var i = 0; i < me.eventHandlers[eventName].length; i++) {
            me.eventHandlers[eventName][i]();
        }
    };

    me.validateToken = function(callback) {
        me.get("/auth/token/validate", function(err, result) {
            if(err !== null) {
                callback(err, false);
            } else {
                callback(err, true);
            }
        });
    };

    me.login = function(email, password, callback) {
        me.post("/auth/login", {
                email : email,
                password : password
            }, callback);
    };

    me.register = function(fullName, email, password, callback) {
        me.post("/auth/register", {
            fullName : fullName,
            email : email,
            password : password
        }, callback);
    };
    
    me.apiUrl = function(path) {
        if(path.indexOf("/") === 0) {
            path = path.slice(1);
        }
        return me.baseUrl + path;
    };

    me.request = function(method, path, data, callback) {
        $.ajax({
            type : method,
            url : me.apiUrl(path),
            data : data !== null ? JSON.stringify(data) : null,
            contentType : 'application/json',
            success : function(data) {
                callback(null, data);
            },
            headers: {
                'Authorization' : me.tokenProvider()
            }
        }).fail(function(err) {
            if(err.status === 403) {
                me.trigger('logged-out');
                callback(err, null);
                return;
            }
            if(err.responseText !== "") {
                try {
                    // Check if the error is parseable
                    err = JSON.parse(err.responseText);    
                }
                catch (ex) {
                    err = { error : err.responseText };
                }
            } else {
                err = {error : err};
            }
            callback(err, null);
        });
    };

    me.del = function(path, callback) {
        me.request('DELETE', path, null, callback);
    };  

    me.get = function(path, callback) {
        me.request("GET", path, null, callback);
    };

    me.post = function(path, data, callback) {
        me.request("POST", path, data, callback);
    };

    me.put = function(path, data, callback) {
        me.request('PUT', path, data, callback);
    };

    return me;
};