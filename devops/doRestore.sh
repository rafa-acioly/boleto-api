#!/bin/bash
CURRENT=`pwd`
BASE_DIR=$HOME/backups
if [[ "$1" != "" ]]; then
    BCK_DATE=$1
    #clen actual dump folder
    rm -rf $HOME/dump_boletodb/*
    cd $BASE_DIR
    tar -xvf backup-boleto-api-$BCK_DATE.tar
    ls -la $HOME/dump_boletodb
    cp -r dump_boletodb/* $HOME/dump_boletodb/
    echo "----------------------------------------"
    ls -la $HOME/dump_boletodb
    sudo docker exec -i -t mongodb mongorestore
    rm -rf dump_boletodb/
    cd $CURRENT
else
    echo "Input backup date formatted like YYYY-MM-DD"
fi