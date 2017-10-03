#!/bin/bash

echo $3
for i in `ls $2/bundles/`
do
    container=container-$i
    touch /var/run/checkin-clear-containers/container-$i
#    curl -i -s "127.0.0.1:9090/checkin?containerID='"$container"'&event=Starting"
    echo "Running container-$i"
    sudo $1 run -b $2/bundles/$i container-$i &
done
