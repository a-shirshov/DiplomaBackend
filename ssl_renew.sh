#!/bin/bash

/usr/bin/docker-compose -f /root/DiplomaBackend/docker-compose.yml run certbot renew --dry-run \
&& /usr/bin/docker-compose -f /root/DiplomaBackend/docker-compose.yml kill -s SIGHUP nginx