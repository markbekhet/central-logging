#!/bin/bash

kubectl apply -f secret.yml
kubectl apply -f postgres.yml
kubectl apply -f config-map.yml
sleep 30
kubectl apply -f logging-server.yml
kubectl apply -f http-server.yml