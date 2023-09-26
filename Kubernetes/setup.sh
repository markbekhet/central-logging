#!/bin/bash
kubectl apply -f nfs-server.yml
kubectl apply -f nfs-pv.yml
kubectl apply -f secret.yml
kubectl apply -f postgres-pvc.yml
kubectl apply -f postgres.yml

status2=$(kubectl get pods | grep postgres-2 | awk '{print $3}')
while [[ $status2 != "Running" ]]
do
    sleep 10
    status2=$(kubectl get pods | grep postgres-2 | awk '{print $3}')
done

echo "postgres-2 is up"

kubectl cp postgres.conf postgres-0:/var/lib/postgresql/data/postgres.conf -c postgres
kubectl cp postgres.conf postgres-1:/var/lib/postgresql/data/postgres.conf -c postgres
kubectl cp postgres.conf postgres-2:/var/lib/postgresql/data/postgres.conf -c postgres
kubectl cp pg_hba.conf postgres-0:/var/lib/postgresql/data/pg_hba.conf -c postgres
kubectl cp pg_hba.conf postgres-1:/var/lib/postgresql/data/pg_hba.conf -c postgres
kubectl cp pg_hba.conf postgres-2:/var/lib/postgresql/data/pg_hba.conf -c postgres

kubectl exec -t postgres-0 -- su postgres -c 'pg_ctl reload -D /var/lib/postgresql/data'
kubectl exec -t postgres-1 -- su postgres -c 'pg_ctl reload -D /var/lib/postgresql/data'
kubectl exec -t postgres-2 -- su postgres -c 'pg_ctl reload -D /var/lib/postgresql/data'

kubectl exec -t postgres-0 -- service postgresql restart
kubectl exec -t postgres-1 -- service postgresql restart
kubectl exec -t postgres-2 -- service postgresql restart

echo "DB cluster setup is done"
#kubectl apply -f postgres-replica.yml
kubectl apply -f config-map.yml
sleep 30
kubectl apply -f logging-server.yml
kubectl apply -f http-server.yml