var bootstrap = bootstrap || {};
bootstrap.components = bootstrap.components || {};
bootstrap.components.forms = bootstrap.components.forms || {}; 

bootstrap.components.forms.register = function(options) {
    var me = this;

    me.form = options.form;
    me.app = options.app;

    me.form.on('submit', function(e) {
        e.preventDefault();

        var email = me.form.find('.email').val();
        var fullName = me.form.find('.full-name').val();
        var password = me.form.find('.password').val();

        me.app.api.register(fullName, email, password, function(err, data) {
            if(err) {
                me.form.find('.registration-errors').html(err.error).show();
                return;
            }
            me.app.setAuth(data);
            me.app.router.routeTo('/dashboard');
        });

    });

    // Highlight the first input
    me.form.find('input').first().focus();

    return me;
};