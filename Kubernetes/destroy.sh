kubectl delete statefulset --all
kubectl delete deployment --all
sleep 30
kubectl delete service --all
kubectl delete persistentVolumeClaim --all
kubectl delete persistentVolume --all
kubectl delete secret --all
kubectl delete configMap --all
