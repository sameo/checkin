#!/bin/bash

curl -i "$1:9090/checkin?containerID=`hostname`&event=Running"
