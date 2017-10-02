#!/bin/bash

for i in $(eval echo {1..$1});
do
    echo -n "Preparing container #$i..."
    mkdir bundle-$i
    cp config.json bundle-$i
    cd bundle-$i
    container=container-$i
    cp -rf ../bundle/rootfs rootfs
    sed -i 's/runc/'"$container"'/' config.json
    cd ..
    echo "Done"
done


for i in $(eval echo {1..$1});
do
    container=container-$i
    curl -i "127.0.0.1:9090/checkin?containerID='"$container"'&event=Starting"
    sudo runc run -b bundle-$i container-$i &
done
