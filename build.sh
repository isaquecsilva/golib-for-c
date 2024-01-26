#!/bin/bash

create_build_folder() {
	mkdir -p build/lib
}

if [ ! -d build ]; then
	create_build_folder
else
	rm -rf build
	create_build_folder
fi

echo 'building go library...'
cd golib
go build -buildmode=c-archive -o ../build/lib/gobase64.a .
cd ..

echo 'building c app...'

gcc -o build/app capp/capp.c -L./build/lib -l:gobase64.a


