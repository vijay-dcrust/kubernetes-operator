# Generic-Daemon-operator
Easy to deploy any kind of daemon on any node

Demonstrates the kubernetes operator concepts

# Install the operator
make install

# Run the controller
make run ENABLE_WEBHOOKS=false

# Apply a test yaml to create GenericDaemon

kubectl apply -f test.yaml

kubectl get genericdaemon

kubectl get daemonset

kubectl describe genericdaemons genericdaemon-sample

#Run a generic daemon on any node

kubectl label nodes mynode1 daemon=http

kubectl describe genericdaemons genericdaemon-sample
