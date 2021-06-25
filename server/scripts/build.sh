#!/bin/bash

#Loop through all the function folder
function loop() { 
    for i in "$1"/*
    do
        if [ -d "$i" ]; then
            subdirectory "$i" ""
        else
            echo "$i"" - Folder Empty"
        fi
    done
}

#Loop through all the different modules
function subdirectory(){
    for i in "$1"/*
    do
        if [ -d "$i" ]; then
            subdirectory "$i" "${1##*/}"
        elif [ -e "$i" ]; then
            
            base="$PWD/bin"
            exedir="$base/exe/$2/"
            deploydir="$base/deploy/$2/"
            filename="${i##*/}"
            exe="${filename%.go}"
            
            echo "Compiling $2 : $filename"
            env GOOS=linux go build -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o $exedir$exe $PWD/aws/$2/$exe/$filename
            mkdir -p $deploydir
            zip -rj $deploydir$exe.zip $exedir$exe
        else
            echo "$i"" - Folder Empty"
        fi
    done
}

loop "$PWD/aws"
