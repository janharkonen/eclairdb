docker stop eclairdb-mono-container
docker rm eclairdb-mono-container
docker build -t eclairdb_mono .
docker run -p 8080:80 --name eclairdb-mono-container eclairdb_mono
