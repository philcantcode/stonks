docker context use default
docker-compose down

docker build -t stocksregistry.azurecr.io/stonks:api-server server/
docker build -t stocksregistry.azurecr.io/stonks-web:web-server web/

docker-compose up -d
docker-compose logs -f server