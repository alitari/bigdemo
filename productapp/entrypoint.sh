#!/bin/bash
printf "%s %s" $BIGDEMO_KEYCLOAK_HTTP_SERVICE_HOST "bigdemo.com" >> /etc/hosts
java -Djava.security.egd=file:/dev/./urandom -jar /app.jar
