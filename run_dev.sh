#docker compose build goapi vuefrontend-dev nginx-dev
#docker compose up goapi vuefrontend-dev nginx-dev --watch
docker compose build vuefrontend-dev nginx-dev
docker compose up vuefrontend-dev nginx-dev --watch
