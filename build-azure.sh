docker context use azure
docker compose down

docker context use default
docker build -t stocksregistry.azurecr.io/stonks:api-server .
docker push stocksregistry.azurecr.io/stonks:api-server

docker context use azure
docker-compose up -d