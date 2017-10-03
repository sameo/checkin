#!/bin/bash

for i in $(eval echo {1..$1});
do
    echo -n "Preparing container #$i..."
    mkdir -p $2/bundles/$i
    cp config.json $2/bundles/$i
    sudo -E cp -rf $2/bundle/rootfs $2/bundles/$i/rootfs
    container=container-$i
    sed -i 's/runc/'"$container"'/' $2/bundles/$i/config.json
    echo "Done"
done
