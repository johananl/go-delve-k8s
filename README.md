# Debugging Go k8s services with Delve

This guide demonstrates how to debug a Go microservice running in a pod on k8s using VS Code.

## Requirements

* VS Code
* Go
* Docker
* kubectl
* A running k8s cluster

## Instructions

Build and push the sample service to a Docker repository:

```
IMAGE=quay.io/foo/bar
docker build -t $IMAGE .
docker push $IMAGE
```

>NOTE: If you're using [kind](https://kind.sigs.k8s.io/) as your k8s cluster, instead of pushing
>the image you can use `kind load docker-image $IMAGE` to load the image into your cluster locally.

Deploy the sample service:

```
sed -i "s|<image_placeholder>|$IMAGE|" deployment.yaml
kubectl apply -f deployment.yaml
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
    "host": "127.0.0.1"
}
```

Start a debugging session in VS Code by clicking the |> button or hitting F5. VS Code should attach
to the remove Delve process inside the pod, which in turn should make Delve start the Go service
and begin debugging.

## Notes and caveats

**Go service should be built with debug information.** By default, Go optimizes binaries by
stripping away debug information. This can render certain information unavailable to Delve when
debugging. To avoid stripping debug information, use the following flags when building the Go
binary: `-gcflags "all=-N -l"`


**Local code must match remote code.** Delve may behave unexpectedly when the code revision in VS
Code differs from the one running on k8s.
