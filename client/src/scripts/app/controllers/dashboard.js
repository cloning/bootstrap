var bootstrap = bootstrap || {};
bootstrap.controllers = bootstrap.controllers || {};

bootstrap.controllers.dashboard = function(options) {
    var me = this;

    me.app = options.app;

    me.index = function() {
        var view = $(templatizer.dashboard.index());
        me.app.setView(view);
    };

    return me;
};