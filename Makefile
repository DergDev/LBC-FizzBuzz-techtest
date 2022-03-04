dev:
    docker-compose \
        --file docker-compose.dev.yml up \
        --detach \
        --build \
        --remove-orphans \
        --force-recreate \

prod:
    docker-compose \
        --file docker-compose.prod.yml up \
        --detach \
        --build \
        --remove-orphans \
        --force-recreate \

stop_dev:
    docker-compose -f docker-compose.dev.yml stop 

stop_prod:
    docker-compose -f docker-compose.prod.yml stop 

logs:
    docker-compose logs --follow