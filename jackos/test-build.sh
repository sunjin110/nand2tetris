#!/bin/sh


# ./build.sh

# jackosを各テストディレクトリに配置
for dir in `ls testdata/`; do
    # echo $dir
    cp ./src/*.jack ./testdata/$dir/
done

# compile
for dir in `ls testdata/`; do
    ./compiler ./testdata/$dir/
done
