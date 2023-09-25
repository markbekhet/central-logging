#!/bin/bash
kubectl apply -f nfs-server.yml
kubectl apply -f nfs-pv.yml
kubectl apply -f secret.yml
kubectl apply -f postgres-pvc.yml
kubectl apply -f postgres.yml
kubectl apply -f postgres-replica.yml
kubectl apply -f config-map.yml
sleep 30
kubectl apply -f logging-server.yml
kubectl apply -f http-server.yml