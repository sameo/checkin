!#/bin/bash

for i in `ls $2/bundles/`
do
    echo "Deleting container-$i"
    sudo $1 delete --force container-$i
done
