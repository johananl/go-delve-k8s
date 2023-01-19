# Debugging Go k8s services with Delve

**WARNING! Attaching a debugger to a process may cause the process to hang, which can disrupt
prduction services.**

## Assumptions

* We have a Go service we want to debug running inside a Docker container on a k8s cluster.
* The `dlv` binary exists in the container image of the Go service. Alternatively, you can copy a
  `dlv` binary for the correct architecture to the pod using `kubectl cp`.
* You are OK with disrupting the remote process while debugging.

## Instructions

Build and deploy the sample service:

```
kind create cluster
docker build -t sample-app .
kind load docker-image sample-app
kubectl apply -f deployment.yaml
```

Start a remote debugger:

```
kubectl exec -it $(kubectl get pods -l app=sample-app -oname) -- dlv attach 1 --listen :2345 --headless --accept-multiclient
```

Enable port forwarding for both the sample service and Delve:

```
kubectl port-forward $(kubectl get pods -l app=sample-app -oname) 8080:8080 2345:2345
```

>NOTE: We need port forwarding for the sample service only so that we can send dummy traffic to it.
>If the code path we want to debug is already being executed, all we need is port forwarding for
>Delve (port 2345).

Wire VS Code to the remote debugger using the following debug config:

```json
{
    "name": "Debug remote Go process",
    "type": "go",
    "request": "attach",
    "mode": "remote",
    "remotePath": "",
    "port": 2345,
    "host": "127.0.0.1",
    "showLog": true,
    "trace": "log",
    "logOutput": "rpc"
}
```

Start a debugging session in VS Code by clicking the |> button or hitting F5.

## Caveats

* Go service should be built with debug information.
* Local code must match remote code.
