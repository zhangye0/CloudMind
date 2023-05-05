docker stop $(docker ps -aq)

sudo chown -R $USER:$USER ../CloudMind
sudo rm -rf data/mysql/data

docker-compose -f docker-compose-env.yml up -d

docker-compose up -d

docker exec -it cloudmind_golang_1 bash

