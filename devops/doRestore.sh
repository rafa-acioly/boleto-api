#!/bin/bash
docker-compose down
cd ..
tar -xf bkp_mongo.tar
mv boletodb ~/boletodb
docker-compose up -d