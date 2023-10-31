## Kubernetes API Gateway example

1. Create Pods
```
kubectl run mipod1 --image=nginx --restart=Never
kubectl run mipod2 --image=nginx --restart=Never
```
2. Create Services
```
kubectl expose pod mipod1 --port=80 --name=mipod1-svc
kubectl expose pod mipod2 --port=80 --name=mipod2-svc
```
3. Install Gateway Feature from K8s
```
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v0.8.1/standard-install.yaml
```
4. Install NGINX Gateway
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v0.8.1/standard-install.yaml
kubectl apply -f https://github.com/nginxinc/nginx-gateway-fabric/releases/download/v1.0.0/crds.yaml
kubectl apply -f https://github.com/nginxinc/nginx-gateway-fabric/releases/download/v1.0.0/nginx-gateway.yaml
kubectl get pods -n nginx-gateway
kubectl apply -f https://raw.githubusercontent.com/nginxinc/nginx-gateway-fabric/v1.0.0/deploy/manifests/service/loadbalancer.yaml
5. Get the public API Assigned to the gateway
```
kubectl get svc -n nginx-gateway
```
6. Create 2 pods and test the traffic split 90% mipod1-svc and 10% mipod2-svc
```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: mipod1
  name: mipod1
spec:
  containers:
  - image: nginx
    name: mipod1
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    run: mipod1
  name: mipod1-svc
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    run: mipod1
status:
  loadBalancer: {}
---
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: mipod2
  name: mipod2
spec:
  containers:
  - image: nginx
    name: mipod2
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    run: mipod2
  name: mipod2-svc
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    run: mipod2
status:
  loadBalancer: {}
---
apiVersion: gateway.networking.k8s.io/v1beta1
kind: Gateway
metadata:
  name: prod-web
spec:
  gatewayClassName: nginx
  listeners:
  - protocol: HTTP
    port: 80
    name: prod-web-gw
    allowedRoutes:
      namespaces:
        from: Same
---
apiVersion: gateway.networking.k8s.io/v1beta1
kind: HTTPRoute
metadata:
  name: simple-split
spec:
  rules:
  - backendRefs:
    - name: mipod1-svc
      port: 80
      weight: 90
    - name: mipod2-svc
      port: 80
      weight: 10
EOF
```

# References
- https://gateway-api.sigs.k8s.io/guides/#installing-a-gateway-controller
- https://github.com/nginxinc/nginx-gateway-fabric/blob/main/docs/installation.md
- https://gateway-api.sigs.k8s.io/guides/simple-gateway
- https://gateway-api.sigs.k8s.io/guides/traffic-splitting
