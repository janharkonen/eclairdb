docker compose down
docker stop goapicontainer
docker stop vuefrontendcontainer
docker container rm goapicontainer
docker container rm vuefrontendcontainer
docker image remove eclairdb-goapi:latest 
docker image remove eclairdb-vuefrontend:latest 
docker image remove --force janharkonen/eclairdb-goapi:latest 
docker ps
docker image ls
