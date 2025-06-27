docker-compose down
docker stop goapi
docker container rm goapi-container
docker container rm goapi_container
docker image remove eclairdb_goapi:latest 
