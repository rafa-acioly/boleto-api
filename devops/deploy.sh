#!/bin/bash

echo "  ____        _      _                     _____ ";
echo " |  _ \      | |    | |         /\        |_   _|";
echo " | |_) | ___ | | ___| |_ ___   /  \   _ __  | |  ";
echo " |  _ < / _ \| |/ _ \ __/ _ \ / /\ \ | '_ \ | |  ";
echo " | |_) | (_) | |  __/ || (_) / ____ \| |_) || |_ ";
echo " |____/ \___/|_|\___|\__\___/_/    \_\ .__/_____|";
echo "                                     | |         ";
echo "                                     |_|         ";                                                                                 

cd ..;
echo ""
echo "Creating volume folder"
mkdir -p ~/boletodb/db
mkdir -p ~/boletodb/configdb
mkdir -p ~/dump_boletodb
echo "Compiling API";
go build  -o ./devops/boletoapi;
echo "API Compiled";
cd devops
echo "Starting docker containers"
docker-compose build --no-cache
docker-compose up -d
rm boletoapi
echo "Containers started"
echo ""
echo "(•‿•) - Enjoy!"
echo ""