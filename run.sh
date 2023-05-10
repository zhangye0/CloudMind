docker stop $(docker ps -aq)
docker-compose -f  docker-compose-env.yml up -d
docker stop $(docker ps -aq)
sudo chown -R $USER:$USER data/elasticsearch
docker-compose -f  docker-compose-env.yml up -d
docker-compose up -d

