#!/bin/bash

/usr/bin/docker-compose -f /root/DiplomaBackend/docker-compose.yml run certbot renew \
&& /usr/bin/docker-compose -f /root/DiplomaBackend/docker-compose.yml kill -s SIGHUP nginx