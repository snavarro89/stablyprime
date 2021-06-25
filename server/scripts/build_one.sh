#!/bin/bash

#Loop through all the different modules
function build(){

    base="$PWD/bin"
    exedir="$base/exe/$1/"
    deploydir="$base/deploy/$1/"
    filename="$2"
    
    echo "Compilando $1 : $2"
    #Compile for linux because we will upload this zip as a lambda function to AWS
    env GOOS=linux go build -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o $exedir$filename $PWD/aws/$1/$filename/$filename.go
    mkdir -p $deploydir
    zip -rj $deploydir$filename.zip $exedir$filename
}

build $1 $2
