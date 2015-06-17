// Load plugins
var gulp = require('gulp'),
    gutil = require('gulp-util'),
    prefix = require('gulp-autoprefixer'),
    size = require('gulp-size'),
    rename = require('gulp-rename'),
    imagemin = require('gulp-imagemin'),
    minifyCSS = require('gulp-minify-css'),
    sass = require('gulp-sass'),
    csslint = require('gulp-csslint');


// Minify all css files in the css directory
// Run this in the root directory of the project with `gulp minify-css `
gulp.task('minify-css', function(){
  gulp.src('./public/css/style.css')
    .pipe(minifyCSS())
    .pipe(rename({ extname: '.min.css' }))
    .pipe(size({gzip:true, showFiles: true}))
    .pipe(gulp.dest('./public/css/'));
  gulp.src('./public/css/tachyons.css')
    .pipe(minifyCSS())
    .pipe(rename({ extname: '.min.css' }))
    .pipe(size({gzip:true, showFiles: true}))
    .pipe(gulp.dest('./public/css/'));
});

gulp.task('minify-img', function(){
  gulp.src('./public/img/*')
    .pipe(imagemin({
        progressive: true,
        svgoPlugins: [{removeViewBox: false}],
    }))
    .pipe(gulp.dest('./public/img/'));
});

// Task that compiles scss files down to good old css
gulp.task('pre-process', function(){
    return gulp.src("./sass/style.scss")
        .pipe(sass())
        .on('error', swallowError)
        .pipe(prefix())
        .pipe(size({gzip: false, showFiles: true}))
        .pipe(size({gzip: true, showFiles: true}))
        .pipe(gulp.dest('./public/css/'));
});

// Allows gulp to not break after a sass error.
// Spits error out to console
function swallowError(error) {
  console.log(error.toString());
  this.emit('end');
}

gulp.task('default', ['pre-process'], function(){
  gulp.start('pre-process', 'minify-css', 'minify-img');
});
