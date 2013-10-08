#!/bin/bash
# Gets library for your system

usage () {
    echo "$1"
    echo "Usage:"
    echo "    $0 linux {x86|x86_64}"
    echo "    $0 {osx|windows}"
}

pathFetch () {
    libPath="$1"
    curl -#O "https://github.com/ajhager/vox/raw/master/data/$libPath"
}

case "$1" in
    osx)
        pathFetch "osx/lib_x86_64/sunvox.dylib"
        ;;
    windows)
        pathFetch "windows/lib_x86/sunvox.dll"
        ;;
    linux)
        if [ "$2" == 'x86_64' ]  || [ "$2" == '64' ]; then
            pathFetch "linux/lib_x86_64/sunvox.so"
        elif [ "$2" == 'x86' ] || [ "$2" == '32' ] || [ "$2" == '86' ]; then
            pathFetch "linux/lib_x86/sunvox.so"
        else
            usage "Linux options need an architecture!"
        fi
        ;;
    *)
        if [ "$*" == "" ]; then
            usage "No arguments given!"
        else
            usage "Unknown arguments \"$*\""
        fi
esac

