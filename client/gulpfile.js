var gulp        = require('gulp');
var stylus      = require('gulp-stylus');
var bower       = require('gulp-bower');
var uglify      = require('gulp-uglify');
var concat      = require('gulp-concat');
var vendor      = require('gulp-concat-vendor');
var jshint      = require('gulp-jshint');
var connect     = require('connect');
var serveStatic = require('serve-static');
var fs          = require('fs');
var templatizer = require('templatizer');

gulp.task('css', function() {
    gulp.src('src/css/main.styl')
        .pipe(stylus({
            compress: true
        }))
        .pipe(gulp.dest('dist/css'));
});

gulp.task('images', function() {
    gulp.src('src/images/*')
        .pipe(gulp.dest('dist/images'));
});

gulp.task('html', function() {
    gulp.src('src/*.html')
        .pipe(gulp.dest('dist/'));
});

gulp.task('scripts:vendor', function() {
    bower({cwd : 'src/scripts/vendor/'}).on('end', function() {
        gulp.src('src/scripts/vendor/bower_components/*')
            .pipe(vendor('vendor.js'))
            .pipe(uglify())
            .pipe(gulp.dest('dist/js'));
    });    
});

gulp.task('fonts', function() {
    gulp.src('src/fonts/**')
        .pipe(gulp.dest('dist/fonts'));
});

gulp.task('scripts:templates', function() {
    templatizer('src/templates/', 'dist/js/templates.js');
});

gulp.task('scripts:app', function() {
    gulp.src('src/scripts/app/**')
        .pipe(jshint())
        .pipe(jshint.reporter('jshint-stylish'))
        .pipe(concat('app.js'))
        //.pipe(uglify())
        .pipe(gulp.dest('dist/js'));
});

gulp.task('watch', ['default'], function() {
    gulp.watch('src/images/*',                          ['images']);
    gulp.watch('src/*.html',                            ['html']);
    gulp.watch('src/scripts/vendor/bower.json',         ['scripts:bower', 'scripts:vendor']);
    gulp.watch('src/**/*.js',                           ['scripts:app']);
    gulp.watch('src/**/*.styl',                         ['css']);
    gulp.watch('src/templates/**',                      ['scripts:templates']);
});

gulp.task('develop', ['default', 'watch'], function() {
    connect()
        .use(serveStatic(__dirname + '/dist'))
        .use(function(req, res) {
            // Aways serve index.html
            res.end(
                fs.readFileSync(
                    'dist/index.html'));
        })
        .listen(3000);
    console.log("Listening at localhost:3000/");
});

gulp.task('default', ['css', 'html', 'images', 'scripts:vendor', 'scripts:app', 'scripts:templates', 'fonts']);