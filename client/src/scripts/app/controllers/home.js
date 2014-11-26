var bootstrap = bootstrap || {};
bootstrap.controllers = bootstrap.controllers || {};

bootstrap.controllers.home = function(options) {
    var me = this;

    me.app = options.app;

    me.index = function() {
        var view = $(templatizer.home.index());
        
        view.find('.header').children().hide();
        
        me.app.setView(view);
        
        view.find('.header').children().fadeIn('slow');
    };

    me.error404 = function() {
        var view = $(templatizer.home.error404());
        me.app.setView(view);
    };

    return me;
};