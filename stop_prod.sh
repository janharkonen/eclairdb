docker compose down
docker stop goapicontainer
docker stop vuefrontendcontainer
docker stop nginxcontainer
docker container rm goapicontainer
docker container rm vuefrontendcontainer
docker container rm nginxcontainer
docker image remove eclairdb-goapi:latest 
docker image remove eclairdb-vuefrontend:latest 
docker image remove eclairdb-nginx:latest 
docker image remove --force janharkonen/eclairdb-goapi:latest 
docker image remove --force janharkonen/eclairdb-vuefrontend:latest 
docker image remove --force janharkonen/eclairdb-nginx:latest
docker ps
docker image ls
