#!/bin/bash

/usr/bin/docker-compose -f /home/artyom/DiplomaBackend/docker-compose.yml restart certbot renew --dry-run \
&& /usr/bin/docker-compose -f /home/artyom/DiplomaBackend/docker-compose.yml kill -s SIGHUP nginx