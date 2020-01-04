# Generic-Daemon-operator
# Install the operator
make install
# Run the controller
make run ENABLE_WEBHOOKS=false
# Apply a test yam
kubectl apply -f config/samples/mygroup_v1beta1_genericdaemon.yaml
kubectl get genericdaemon
kubectl get daemonset
kubectl describe genericdaemons genericdaemon-sample
#Run a generic daemon on any node
kubectl label nodes mynode1 daemon=http
kubectl describe genericdaemons genericdaemon-sample
