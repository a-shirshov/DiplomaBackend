_#!/bin/bash

/usr/local/bin/docker-compose -f /root/DiplomaBackend/docker-compose.yml run certbot renew --dry-run \
&& /usr/local/bin/docker-compose -f /root/DiplomaBackend/docker-compose.yml kill -s SIGHUP webserver