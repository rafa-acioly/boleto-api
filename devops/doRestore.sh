#!/bin/bash
CURRENT=`pwd`
mv bkp_mongo.tar ~/
docker-compose down
cd $HOME
rm -rf $HOME/boletodb
tar -zxf bkp_mongo.tar
mv $HOME/bkp_mongo.tar $CURRENT
cd $CURRENT
docker-compose up -d