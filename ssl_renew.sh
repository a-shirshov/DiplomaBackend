#!/bin/bash

/usr/bin/docker-compose -f /home/artyom/DiplomaBackend/docker-compose.yml --env-file /home/artyom/DiplomaBackend/.env run --no-deps certbot renew --dry-run \
&& /usr/bin/docker-compose -f /home/artyom/DiplomaBackend/docker-compose.yml kill -s SIGHUP nginx