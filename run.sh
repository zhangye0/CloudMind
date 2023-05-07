docker stop $(docker ps -aq)
sudo chown -R $USER:$USER data/elasticsearch
docker-compose -f deploy/docker/docker-compose-env.yml up -d
docker-compose up -d

