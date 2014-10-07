var gulp        = require('gulp');
var stylus      = require('gulp-stylus');
var bower       = require('gulp-bower');
var uglify      = require('gulp-uglify');
var concat      = require('gulp-concat');
var vendor      = require('gulp-concat-vendor');
var jshint      = require('gulp-jshint');

gulp.task('css', function() {
    gulp.src('src/css/main.styl')
        .pipe(stylus({
            compress: true
        }))
        .pipe(gulp.dest('dist/css'));
});

gulp.task('scripts:bower', function() {
    bower({cwd : 'src/scripts/vendor/'});
});

gulp.task('scripts:vendor', function() {
    gulp.src('src/scripts/vendor/bower_components/*')
        .pipe(vendor('vendor.js'))
        .pipe(uglify())
        .pipe(gulp.dest('dist/js'));
});

gulp.task('scripts:app', function() {
    gulp.src('src/scripts/app/*')
        .pipe(jshint())
        .pipe(jshint.reporter('jshint-stylish'))
        .pipe(concat('app.js'))
        .pipe(uglify())
        .pipe(gulp.dest('dist/js'));
});


gulp.task('watch', ['default'], function() {
    gulp.watch('src/scripts/vendor/bower.json',         ['scripts:bower', 'scripts:vendor']);
    gulp.watch('src/**/*.js',                           ['scripts:app']);
    gulp.watch('src/**/*.styl',                         ['css']);
});

gulp.task('default', ['css', 'scripts:bower', 'scripts:vendor', 'scripts:app']);