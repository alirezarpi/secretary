#!/bin/bash -xe
mkdir -p certs-invalid
rm certs-invalid/*
#openssl req -new -nodes -x509 -out certs-invalid/server.pem -keyout certs-invalid/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=alireza.khalili@pipi.de"
openssl req -new -nodes -x509 -out certs-invalid/client.pem -keyout certs-invalid/client.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=alireza.khalili@pipi.de"

