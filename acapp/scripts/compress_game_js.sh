#! /bin/bash

JS_PATH=/home/ycx/Web/acapp/game/static/js/
JS_PATH_DIST=${JS_PATH}dist/
JS_PATH_SRC=${JS_PATH}src/

find $JS_PAth_SRC -type f -name '*.js' | sort | xargs cat > ${JS_PATH_DIST}game.js