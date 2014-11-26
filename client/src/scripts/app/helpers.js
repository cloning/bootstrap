function isFunction(functionToCheck) {
    var getType = {};
    return functionToCheck && getType.toString.call(functionToCheck) === '[object Function]';
}

function isObject(objectToCheck) {
    return objectToCheck !== null && typeof objectToCheck === 'object';
}

function splitRemoveEmpty(string, separator) {
    var split = string.split(separator);
    var result = [];
    for(var i = 0; i < split.length; i++) {
        var current = split[i];
        if(current !== "") {
            result.push(current);
        }
    }
    return result;
}

function getCookie(name) {
    var nameEQ = name + "=";
    var ca = document.cookie.split(';');
    for(var i=0;i < ca.length;i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1,c.length);
        if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length,c.length);
    }
    return null;
}

function setCookie(name, val, expires) {
    var cookieStr =  name + "=" + val;
    if(expires) {
        cookieStr += ";expires=" + new Date(expires).toGMTString();
    } 
    document.cookie = cookieStr;
}

function querystring(name) {
    name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
    var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
        results = regex.exec(location.search);
    return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
}