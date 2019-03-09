# Big Demo

![Architecture](bigdemo.png)

## Getting started

### Docker for windows

- We assume Docker for Desktop is running and you have `kubectl` and `jq` working: `kubectl get nodes -o=json | jq .items[0].status.allocatable` should show that your machine has at least 8G and 4 cpus available for k8s.
- install helm: Download the binaries for windows and execute `helm install`
- Run the build with `skaffold build`. If the images are not present the must be pulled, this can take a while, so be patient.
- Deploy the helm chart with `skaffold deploy`. Like the step above the 3rd party images for redis and rabbitmq must be pulled, which can take a while. The redis might have some bootstrap problems with the master-slave connection, solve it with deleting the according pod. 
- Default profile which is used in the step above is configured for the Docker for Desktop local kubernetes installation. E.g. the ingress and ingress-controller which routes the incoming http requests to the services is configured for localhost and no tls. You can check this for example with  `kubectl get svc bigdemo-nginx-ingress-controller` and `kubectl get ingress bigdemo-bigdemo -o json | jq .spec.rules`. Browse to `http://localhost` and you can see the bigdemo ui!

## Development-flow

Activate the deveopment cycle with `skaffold dev --port-forward=false`, we disable port-forwarding because we are using - as we have seen before - an ingress controller.