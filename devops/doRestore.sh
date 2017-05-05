#!/bin/bash
CURRENT=`pwd`
BASE_DIR=$HOME/backups
#limpa a pasta de dump atual
sudo rm -rf $HOME/dump_boletodb/*
echo "Informe a data do arquivo de backup que será carregado:YYYY-MM-DD"
read BCK_DATE
cd $BASE_DIR
tar -xvf bck_boletoapi-$BCK_DATE.tar
ls -la $HOME/dump_boletodb
cp -r dump_boletodb/* $HOME/dump_boletodb/
echo "----------------------------------------"
ls -la $HOME/dump_boletodb
sudo docker exec -i -t mongodb mongorestore
rm -rf dump_boletodb
cd $CURRENT