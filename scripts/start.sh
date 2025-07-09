#!/bin/bash
docker compose -f docker-compose.yml -p mongo-go up -d

sleep 5

docker exec mongo1 /scripts/rs-init.sh
