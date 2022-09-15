#!/bin/bash

/usr/bin/docker-compose -f /home/artyom/DiplomaBackend/docker-compose.yml run certbot renew --dry-run --no-deps \
&& /usr/bin/docker-compose -f /home/artyom/DiplomaBackend/docker-compose.yml kill -s SIGHUP nginx