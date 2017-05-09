#!/bin/bash

if [[ "$1" != "" ]]; then

echo "  ____        _      _                     _____ ";
echo " |  _ \      | |    | |         /\        |_   _|";
echo " | |_) | ___ | | ___| |_ ___   /  \   _ __  | |  ";
echo " |  _ < / _ \| |/ _ \ __/ _ \ / /\ \ | '_ \ | |  ";
echo " | |_) | (_) | |  __/ || (_) / ____ \| |_) || |_ ";
echo " |____/ \___/|_|\___|\__\___/_/    \_\ .__/_____|";
echo "                                     | |         ";
echo "                                     |_|         ";                                                                                 

echo ""

cd "$1"
echo "Creating volume folder"
mkdir -p ~/boletodb/upMongo
mkdir -p ~/boletodb/db
mkdir -p ~/boletodb/configdb
mkdir -p ~/dump_boletodb
echo "Compiling API";
go build  -o ./devops/boletoapi;
echo "API Compiled";
cd devops
echo "Starting docker containers"
docker-compose build --no-cache
if [ "$2" == 'local' ]; then
    docker-compose up -d
else
    docker-compose -f ./docker-compose.release.yml up -d
fi
rm boletoapi
echo "Containers started"
echo ""
echo "(•‿•) - Enjoy!"
echo ""
else
    echo "[ERROR] Expecting build directory as argument"
fi



