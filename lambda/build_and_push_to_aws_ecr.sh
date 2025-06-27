docker image remove eclairdb-goapi-lambda
rm -rf main
####
cd ../goapi
sleep 3
GOOS=linux GOARCH=amd64 go build -o main main.go
####
mv main ../lambda/main
cd ../lambda
####
docker build -t eclairdb-goapi-lambda -f DockerfileLambda .
echo "tag to aws ecr uri"
sleep 3
docker tag eclairdb-goapi-lambda:latest 528757783961.dkr.ecr.eu-north-1.amazonaws.com/jan/eclairdb:latest
echo "push to aws ecr"
sleep 3
docker push 528757783961.dkr.ecr.eu-north-1.amazonaws.com/jan/eclairdb:latest
