#!/bin/bash
docker-compose down
#Altere o arquivo para zipar a pasta que foi mapeada no arquivo docker-compose.yml
tar -cvf bkp_mongo.tar ~/boletodb

docker-compose up -d