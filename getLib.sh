#!/bin/bash
# Gets sunvox library for your system

usage () {
    echo "$1"
    echo "Usage:"
    echo "    $0 linux {x86|x86_64}"
    echo "    $0 {osx|windows}"
}

pathFetch () {
    libPath="$1"
    echo "Connecting to github.com/ajhager/vox"
    echo "Getting library $libPath"
    curl -#O "https://raw.github.com/ajhager/vox/master/data/$libPath"
    if [ "$?" == 0 ]; then
        echo "Success. Sunvox libraries have been installed."
    else
        echo "Failure! Sunvox libraries have not been installed."
    fi
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

