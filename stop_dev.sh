docker compose down
docker stop goapicontainer
docker stop vuefrontendcontainer-dev
docker container rm goapicontainer
docker container rm vuefrontendcontainer-dev
docker image remove eclairdb-goapi:latest 
docker image remove eclairdb-vuefrontend-dev:latest 
docker image remove --force janharkonen/eclairdb-goapi:latest 
docker image remove --force janharkonen/eclairdb-vuefrontend-dev:latest 
docker ps
docker image ls
