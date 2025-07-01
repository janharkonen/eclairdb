docker compose down
docker stop goapicontainer
docker container rm goapicontainer
docker container rm goapicontainer
docker image remove eclairdb-goapi:latest 
docker image remove --force janharkonen/eclairdb-goapi:latest 
docker ps
docker image ls
