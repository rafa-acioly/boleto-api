#!/bin/bash
CURRENT=`pwd`
docker exec -i -t mongodb mongodump
DATE=`date +%Y-%m-%d`
BASE_DIR=$HOME/backups
mkdir -p $HOME/backups
BCK_NAME=backup-boleto-api-$DATE.tar
cd $HOME
tar -cf $BCK_NAME dump_boletodb/
mv $BCK_NAME $BASE_DIR
cd $CURRENT