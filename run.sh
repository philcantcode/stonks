docker build -t stonks .
docker-compose up -d
docker exec -it stonks_server_1 /bin/bash