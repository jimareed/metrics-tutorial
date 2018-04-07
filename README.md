# prometheus
This tutorial will walk through the following:
- install promethus & grafana using helm
- add metrics to your service
- monitor SLI and SLO from grafana

**Prerequisites**: install docker, Kubernetes and helm and clone repo.

![Prometheus Tutorial](./tutorial.png)


```
$ $ helm install --name prometheus stable/prometheus
NAME:   prometheus
LAST DEPLOYED: Sat Apr  7 09:40:01 2018
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1beta1/ClusterRoleBinding
NAME                           AGE
prometheus-alertmanager        1s
prometheus-kube-state-metrics  1s
prometheus-node-exporter       1s
prometheus-server              1s

==> v1/Service
NAME                           TYPE       CLUSTER-IP      EXTERNAL-IP  PORT(S)   AGE
prometheus-alertmanager        ClusterIP  10.110.178.93   <none>       80/TCP    1s
prometheus-kube-state-metrics  ClusterIP  None            <none>       80/TCP    1s
prometheus-node-exporter       ClusterIP  None            <none>       9100/TCP  1s
prometheus-pushgateway         ClusterIP  10.111.221.7    <none>       9091/TCP  1s
prometheus-server              ClusterIP  10.105.233.230  <none>       80/TCP    1s

==> v1beta1/DaemonSet
NAME                      DESIRED  CURRENT  READY  UP-TO-DATE  AVAILABLE  NODE SELECTOR  AGE
prometheus-node-exporter  1        1        0      1           0          <none>         1s

==> v1/ConfigMap
NAME                     DATA  AGE
prometheus-alertmanager  1     1s
prometheus-server        3     1s

==> v1/ServiceAccount
NAME                           SECRETS  AGE
prometheus-alertmanager        1        1s
prometheus-kube-state-metrics  1        1s
prometheus-node-exporter       1        1s
prometheus-server              1        1s

==> v1beta1/ClusterRole
NAME                           AGE
prometheus-kube-state-metrics  1s
prometheus-server              1s

==> v1beta1/Deployment
NAME                           DESIRED  CURRENT  UP-TO-DATE  AVAILABLE  AGE
prometheus-alertmanager        1        1        1           0          0s
prometheus-kube-state-metrics  1        1        1           0          0s
prometheus-pushgateway         1        1        1           0          0s
prometheus-server              1        1        1           0          0s

==> v1/Pod(related)
NAME                                            READY  STATUS             RESTARTS  AGE
prometheus-node-exporter-9h8p4                  0/1    ContainerCreating  0         0s
prometheus-alertmanager-85d944f874-g5znl        0/2    ContainerCreating  0         0s
prometheus-kube-state-metrics-786b6cbc77-bwhv4  0/1    ContainerCreating  0         0s
prometheus-pushgateway-68966b6ff7-p4rgs         0/1    ContainerCreating  0         0s
prometheus-server-6966b574d7-6tm8g              0/2    Init:0/1           0         0s

==> v1/PersistentVolumeClaim
NAME                     STATUS   VOLUME                                    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
prometheus-alertmanager  Bound    pvc-2bf79d21-3a69-11e8-a0dd-025000000001  2Gi       RWO           hostpath      1s
prometheus-server        Pending  hostpath                                  1s


NOTES:
The Prometheus server can be accessed via port 80 on the following DNS name from within your cluster:
prometheus-server.default.svc.cluster.local


Get the Prometheus server URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace default port-forward $POD_NAME 9090


The Prometheus alertmanager can be accessed via port 80 on the following DNS name from within your cluster:
prometheus-alertmanager.default.svc.cluster.local


Get the Alertmanager URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=alertmanager" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace default port-forward $POD_NAME 9093


The Prometheus PushGateway can be accessed via port 9091 on the following DNS name from within your cluster:
prometheus-pushgateway.default.svc.cluster.local


Get the PushGateway URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=pushgateway" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace default port-forward $POD_NAME 9091

For more information on running Prometheus, visit:
https://prometheus.io/

us-rem-ire-4807:elastic ire$ helm ls
NAME      	REVISION	UPDATED                 	STATUS  	CHART           	NAMESPACE
prometheus	1       	Sat Apr  7 09:40:01 2018	DEPLOYED	prometheus-6.1.1	default  
us-rem-ire-4807:elastic ire$ kubectl get pods
NAME                                             READY     STATUS    RESTARTS   AGE
prometheus-alertmanager-85d944f874-g5znl         1/2       Running   0          45s
prometheus-kube-state-metrics-786b6cbc77-bwhv4   1/1       Running   0          45s
prometheus-node-exporter-9h8p4                   1/1       Running   0          45s
prometheus-pushgateway-68966b6ff7-p4rgs          1/1       Running   0          45s
prometheus-server-6966b574d7-6tm8g               1/2       Running   0          45s

```

```
$ kubectl port-forward prometheus-server-6966b574d7-6tm8g 9090
Forwarding from 127.0.0.1:9090 -> 9090
```

http://localhost:9090/graph


```
$ helm install --name grafana stable/grafana
NAME:   grafana
LAST DEPLOYED: Sat Apr  7 09:48:01 2018
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/Secret
NAME     TYPE    DATA  AGE
grafana  Opaque  2     0s

==> v1/ConfigMap
NAME            DATA  AGE
grafana-config  1     0s
grafana-dashs   0     0s

==> v1/PersistentVolumeClaim
NAME     STATUS  VOLUME                                    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
grafana  Bound   pvc-49ed7275-3a6a-11e8-a0dd-025000000001  1Gi       RWO           hostpath      0s

==> v1/Service
NAME     TYPE       CLUSTER-IP   EXTERNAL-IP  PORT(S)  AGE
grafana  ClusterIP  10.98.70.48  <none>       80/TCP   0s

==> v1beta1/Deployment
NAME     DESIRED  CURRENT  UP-TO-DATE  AVAILABLE  AGE
grafana  1        1        1           0          0s

==> v1/Pod(related)
NAME                      READY  STATUS    RESTARTS  AGE
grafana-55b57d567b-xkvs4  0/1    Init:0/1  0         0s


NOTES:
1. Get your 'admin' user password by running:

   kubectl get secret --namespace default grafana -o jsonpath="{.data.grafana-admin-password}" | base64 --decode ; echo

2. The Grafana server can be accessed via port 80 on the following DNS name from within your cluster:

   grafana.default.svc.cluster.local

   Get the Grafana URL to visit by running these commands in the same shell:

     export POD_NAME=$(kubectl get pods --namespace default -l "app=grafana-grafana,component=grafana" -o jsonpath="{.items[0].metadata.name}")
     kubectl --namespace default port-forward $POD_NAME 3000

3. Login with the password from step 1 and the username: admin

```

```
$ kubectl get secret --namespace default grafana -o jsonpath="{.data.grafana-admin-password}" | base64 --decode ; echo
a9kjfVp6IF
```

Build the docker image locally which contains two services:
```
$ make build-docker
./build-docker.sh prometheus-tutorial latest Dockerfile
Sending build context to Docker daemon  564.7kB
Step 1/19 : FROM golang:1.10-alpine AS builder
1.10-alpine: Pulling from library/golang
ff3a5c916c92: Already exists
f32d2ea73378: Pull complete
dbfec4c268d3: Pull complete
...Successfully built 9bbd346d2575
Successfully tagged helm-chart-tutorial:latest
```
