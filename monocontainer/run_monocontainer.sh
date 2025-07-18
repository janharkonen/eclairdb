cd $(pwd)/../
docker build -t eclairdb_mono . -f ./monocontainer/Dockerfile
docker run -p 8080:80 --name eclairdb-mono-container eclairdb_mono
cd $(pwd)/monocontainer

