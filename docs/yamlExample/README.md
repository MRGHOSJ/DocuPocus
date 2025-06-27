# 📄 YAML Configuration: `devops-pi-main`

[← Back to Overview](../README.md)

## 📄 File: `backend-deployment.yaml`

> 📍 Path: `devops-pi-main\backend-deployment.yaml`

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `Deployment`
- **Description:** Kubernetes Deployment configuration for a Spring Boot application

</details>

<details>
<summary>⚙️ Configuration Example for `Deployment`</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
spec:
  replicas: 2
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
        - name: spring-app
          image: meowster/skillexchange-spring-app:latest
          ports: 
          env: 
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** Kubernetes API version
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (Deployment)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Resource metadata
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** Deployment name
</details>

<details>
<summary>`spec`</summary>

- **Type:** `map`
- **Description:** Deployment specification
</details>

<details>
<summary>`spec.replicas`</summary>

- **Type:** `scalar`
- **Description:** Number of pod replicas
</details>

<details>
<summary>`spec.selector`</summary>

- **Type:** `map`
- **Description:** Pod selection criteria
</details>

<details>
<summary>`spec.template`</summary>

- **Type:** `map`
- **Description:** Pod template
</details>

<details>
<summary>`spec.template.spec.containers`</summary>

- **Type:** `array`
- **Description:** Container specifications
</details>

<details>
<summary>`spec.template.spec.containers[].env`</summary>

- **Type:** `array`
- **Description:** Environment variables
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
spec.replicas: 2
spec.template.spec.containers[0].env[0].name: SPRING_DATASOURCE_URL
spec.template.spec.containers[0].image: meowster/skillexchange-spring-app:latest
apiVersion: apps/v1
kind: Deployment
metadata.name: backend
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **spec.replicas**: `1`
- **spec.template.spec.containers[0].ports[0].containerPort**: `8080`
</details>

<details>
<summary>🧰 Usage</summary>

Deploying a Spring Boot application with MySQL database connectivity
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Avoid using 'latest' tag in production
- Use Secrets for sensitive data like passwords
- Define resource limits for containers
- Use ConfigMaps for environment variables when possible
</details>


---

## 📄 File: `backend-service.yaml`

> 📍 Path: `devops-pi-main\backend-service.yaml`

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `Service`
- **Description:** Defines a Kubernetes Service for network access to pods

</details>

<details>
<summary>⚙️ Configuration Example for `Service`</summary>

```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend
spec:
  selector:
    app: backend
  ports:
    - port: 8084
      targetPort: 8084
  type: ClusterIP
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** Kubernetes API version
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (Service)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Service metadata
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** Service name
</details>

<details>
<summary>`spec`</summary>

- **Type:** `map`
- **Description:** Service specification
</details>

<details>
<summary>`spec.selector`</summary>

- **Type:** `map`
- **Description:** Pod selection labels
</details>

<details>
<summary>`spec.selector.app`</summary>

- **Type:** `scalar`
- **Description:** Target pod label
</details>

<details>
<summary>`spec.ports`</summary>

- **Type:** `array`
- **Description:** Exposed ports configuration
</details>

<details>
<summary>`spec.ports[0].port`</summary>

- **Type:** `scalar`
- **Description:** Service port number
</details>

<details>
<summary>`spec.ports[0].targetPort`</summary>

- **Type:** `scalar`
- **Description:** Target pod port
</details>

<details>
<summary>`spec.type`</summary>

- **Type:** `scalar`
- **Description:** Service type
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
metadata.name: backend
spec.ports[0].port: 8084
spec.ports[0].targetPort: 8084
spec.selector.app: backend
spec.type: ClusterIP
apiVersion: v1
kind: Service
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **spec.type**: `ClusterIP`
</details>

<details>
<summary>🧰 Usage</summary>

Exposes backend pods internally via a stable ClusterIP
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Use meaningful metadata.name
- Ensure selector matches pod labels
- Prefer ClusterIP for internal services
</details>


---

## 📄 File: `frontend-deployment.yaml`

> 📍 Path: `devops-pi-main\frontend-deployment.yaml`

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `Deployment`
- **Description:** Kubernetes Deployment configuration for a frontend application

</details>

<details>
<summary>⚙️ Configuration Example for `Deployment`</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
spec:
  replicas: 2
  selector:
    matchLabels:
      app: frontend
  template:
    metadata:
      labels:
        app: frontend
    spec:
      containers:
        - name: angular-frontend
          image: meowster/angular-app:latest
          ports: 
          env: 
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** Kubernetes API version
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (Deployment)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Resource metadata
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** Deployment name
</details>

<details>
<summary>`spec`</summary>

- **Type:** `map`
- **Description:** Deployment specification
</details>

<details>
<summary>`spec.replicas`</summary>

- **Type:** `scalar`
- **Description:** Number of pod replicas
</details>

<details>
<summary>`spec.selector`</summary>

- **Type:** `map`
- **Description:** Pod selection criteria
</details>

<details>
<summary>`spec.template`</summary>

- **Type:** `map`
- **Description:** Pod template
</details>

<details>
<summary>`spec.template.spec.containers`</summary>

- **Type:** `array`
- **Description:** Container specifications
</details>

<details>
<summary>`spec.template.spec.containers[].ports`</summary>

- **Type:** `array`
- **Description:** Container ports
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
spec.template.spec.containers[0].image: meowster/angular-app:latest
apiVersion: apps/v1
kind: Deployment
metadata.name: frontend
spec.replicas: 2
spec.selector.matchLabels.app: frontend
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **spec.replicas**: `1`
</details>

<details>
<summary>🧰 Usage</summary>

Defines a scalable frontend application deployment in Kubernetes
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Always specify resource limits for containers
- Use explicit image tags (avoid 'latest')
- Define liveness/readiness probes
- Match selector labels with pod labels
</details>


---

## 📄 File: `frontend-service.yaml`

> 📍 Path: `devops-pi-main\frontend-service.yaml`

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `Service`
- **Description:** Defines a Kubernetes Service to expose an application

</details>

<details>
<summary>⚙️ Configuration Example for `Service`</summary>

```yaml
apiVersion: v1
kind: Service
metadata:
  name: frontend
spec:
  selector:
    app: frontend
  ports:
    - port: 80
      targetPort: 80
      nodePort: 30080
  type: NodePort
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** Kubernetes API version
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (Service)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Service metadata
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** Service name
</details>

<details>
<summary>`spec`</summary>

- **Type:** `map`
- **Description:** Service specification
</details>

<details>
<summary>`spec.selector`</summary>

- **Type:** `map`
- **Description:** Pod selector labels
</details>

<details>
<summary>`spec.selector.app`</summary>

- **Type:** `scalar`
- **Description:** Target pod label
</details>

<details>
<summary>`spec.ports`</summary>

- **Type:** `array`
- **Description:** Port mappings
</details>

<details>
<summary>`spec.ports[].port`</summary>

- **Type:** `scalar`
- **Description:** Service port
</details>

<details>
<summary>`spec.ports[].targetPort`</summary>

- **Type:** `scalar`
- **Description:** Pod port
</details>

<details>
<summary>`spec.ports[].nodePort`</summary>

- **Type:** `scalar`
- **Description:** Node port (30000-32767)
</details>

<details>
<summary>`spec.type`</summary>

- **Type:** `scalar`
- **Description:** Service type (NodePort)
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
apiVersion: v1
kind: Service
metadata: map[name:frontend]
spec: map[ports:[map[nodePort:30080 port:80 targetPort:80]] selector:map[app:frontend] type:NodePort]
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **spec.ports[].protocol**: `TCP`
- **spec.type**: `ClusterIP`
</details>

<details>
<summary>🧰 Usage</summary>

Exposes a deployment externally via static port on worker nodes
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Use meaningful metadata.name
- NodePort should be in 30000-32767 range
- Match selector.app with pod labels
</details>


---

## 📄 File: `ingress.yaml`

> 📍 Path: `devops-pi-main\ingress.yaml`

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `Ingress`
- **Description:** Defines Kubernetes Ingress rules for routing HTTP traffic

</details>

<details>
<summary>⚙️ Configuration Example for `Ingress`</summary>

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: app-ingress
  namespace: piproject
spec:
  ingressClassName: nginx
  rules:
    - http: 
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** Kubernetes API version
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (Ingress)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Resource metadata
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** Ingress resource name
</details>

<details>
<summary>`metadata.namespace`</summary>

- **Type:** `scalar`
- **Description:** Target namespace
</details>

<details>
<summary>`spec`</summary>

- **Type:** `map`
- **Description:** Ingress specification
</details>

<details>
<summary>`spec.ingressClassName`</summary>

- **Type:** `scalar`
- **Description:** Ingress controller class
</details>

<details>
<summary>`spec.rules`</summary>

- **Type:** `array`
- **Description:** Routing rules
</details>

<details>
<summary>`spec.rules[].http.paths`</summary>

- **Type:** `array`
- **Description:** Path-based routing rules
</details>

<details>
<summary>`spec.rules[].http.paths[].path`</summary>

- **Type:** `scalar`
- **Description:** URL path pattern
</details>

<details>
<summary>`spec.rules[].http.paths[].pathType`</summary>

- **Type:** `scalar`
- **Description:** Path matching strategy
</details>

<details>
<summary>`spec.rules[].http.paths[].backend.service`</summary>

- **Type:** `map`
- **Description:** Backend service reference
</details>

<details>
<summary>`spec.rules[].http.paths[].backend.service.name`</summary>

- **Type:** `scalar`
- **Description:** Service name
</details>

<details>
<summary>`spec.rules[].http.paths[].backend.service.port.number`</summary>

- **Type:** `scalar`
- **Description:** Service port
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
spec.rules[0].http.paths[0].backend.service.port.number: 8084
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata.name: app-ingress
spec.rules[0].http.paths[0].path: /skillExchange
spec.rules[0].http.paths[0].pathType: Prefix
metadata.namespace: piproject
spec.ingressClassName: nginx
spec.rules[0].http.paths[0].backend.service.name: backend
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **pathType**: `ImplementationSpecific`
</details>

<details>
<summary>🧰 Usage</summary>

Configures external HTTP access to cluster services via path-based routing
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Always specify pathType explicitly
- Use descriptive path names matching your application structure
- Restrict namespace access where applicable
- Document ingress controller requirements
</details>


---

## 📄 File: `mysql-deployment.yaml`

> 📍 Path: `devops-pi-main\mysql-deployment.yaml`

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `StorageClass`
- **Description:** Defines a Kubernetes StorageClass for persistent volumes

</details>

<details>
<summary>⚙️ Configuration Example for `StorageClass`</summary>

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: manual-hostpath
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: Immediate
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** API version for StorageClass
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (StorageClass)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Object metadata including name
</details>

<details>
<summary>`provisioner`</summary>

- **Type:** `scalar`
- **Description:** Volume plugin for provisioning
</details>

<details>
<summary>`volumeBindingMode`</summary>

- **Type:** `scalar`
- **Description:** When volume binding occurs
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata: map[name:manual-hostpath]
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: Immediate
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **volumeBindingMode**: `Immediate`
</details>

<details>
<summary>🧰 Usage</summary>

Used to define storage provisioning behavior for PersistentVolumeClaims
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Use explicit volumeBindingMode for dynamic provisioning
- Avoid no-provisioner in production without external volume management
</details>

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `PersistentVolume`
- **Description:** Defines a PersistentVolume for cluster storage

</details>

<details>
<summary>⚙️ Configuration Example for `PersistentVolume`</summary>

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: mysql-pv
spec:
  capacity:
    storage: 2Gi
  accessModes:
    - 
  storageClassName: manual-hostpath
  hostPath:
    path: /mnt/data/mysql
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** Kubernetes API version
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (PersistentVolume)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Resource metadata
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** PV identifier
</details>

<details>
<summary>`spec.capacity`</summary>

- **Type:** `map`
- **Description:** Storage capacity definition
</details>

<details>
<summary>`spec.capacity.storage`</summary>

- **Type:** `scalar`
- **Description:** Storage size
</details>

<details>
<summary>`spec.accessModes`</summary>

- **Type:** `array`
- **Description:** Volume access permissions
</details>

<details>
<summary>`spec.storageClassName`</summary>

- **Type:** `scalar`
- **Description:** Storage class identifier
</details>

<details>
<summary>`spec.hostPath`</summary>

- **Type:** `map`
- **Description:** Host machine path binding
</details>

<details>
<summary>`spec.hostPath.path`</summary>

- **Type:** `scalar`
- **Description:** Filesystem path
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
apiVersion: v1
kind: PersistentVolume
metadata.name: mysql-pv
spec.accessModes: [ReadWriteOnce]
spec.capacity.storage: 2Gi
spec.hostPath.path: /mnt/data/mysql
spec.storageClassName: manual-hostpath
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **storageClassName**: `<nil>`
</details>

<details>
<summary>🧰 Usage</summary>

Provisioning durable storage for stateful applications like databases
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Avoid hostPath for production (use cloud volumes/NFS)
- Set reclaimPolicy explicitly
- Match accessModes with workload requirements
</details>

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `PersistentVolumeClaim`
- **Description:** Defines a PersistentVolumeClaim for storage allocation

</details>

<details>
<summary>⚙️ Configuration Example for `PersistentVolumeClaim`</summary>

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mysql-pvc
spec:
  accessModes:
    - 
  storageClassName: manual-hostpath
  resources:
    requests:
      storage: 2Gi
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** Kubernetes API version
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (PersistentVolumeClaim)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Resource metadata
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** PVC identifier
</details>

<details>
<summary>`spec`</summary>

- **Type:** `map`
- **Description:** PVC specifications
</details>

<details>
<summary>`spec.accessModes`</summary>

- **Type:** `array`
- **Description:** Volume access modes
</details>

<details>
<summary>`spec.storageClassName`</summary>

- **Type:** `scalar`
- **Description:** Storage class name
</details>

<details>
<summary>`spec.resources`</summary>

- **Type:** `map`
- **Description:** Resource requests/limits
</details>

<details>
<summary>`spec.resources.requests.storage`</summary>

- **Type:** `scalar`
- **Description:** Storage size request
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata.name: mysql-pvc
spec.accessModes: [ReadWriteOnce]
spec.resources.requests.storage: 2Gi
spec.storageClassName: manual-hostpath
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **apiVersion**: `v1`
- **storageClassName**: `default`
</details>

<details>
<summary>🧰 Usage</summary>

Requesting persistent storage for stateful applications like databases
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Always specify accessModes matching application requirements
- Avoid empty storageClassName (may trigger default provisioner)
- Set realistic storage requests to prevent over-provisioning
</details>

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `Deployment`
- **Description:** Kubernetes Deployment configuration for MySQL with persistent storage

</details>

<details>
<summary>⚙️ Configuration Example for `Deployment`</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
        - name: mysql
          image: mysql:8.0
          args: 
          env: 
          ports: 
          volumeMounts: 
      volumes:
        - name: mysql-storage
          persistentVolumeClaim: 
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** Kubernetes API version
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (Deployment)
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** Deployment name
</details>

<details>
<summary>`spec.replicas`</summary>

- **Type:** `scalar`
- **Description:** Number of pod replicas
</details>

<details>
<summary>`spec.selector.matchLabels`</summary>

- **Type:** `map`
- **Description:** Pod selection labels
</details>

<details>
<summary>`spec.template.spec.containers`</summary>

- **Type:** `array`
- **Description:** Container specifications
</details>

<details>
<summary>`spec.template.spec.containers[].env`</summary>

- **Type:** `array`
- **Description:** Environment variables
</details>

<details>
<summary>`spec.template.spec.containers[].ports`</summary>

- **Type:** `array`
- **Description:** Container ports
</details>

<details>
<summary>`spec.template.spec.volumes`</summary>

- **Type:** `array`
- **Description:** Storage volumes
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
kind: Deployment
metadata.name: mysql
spec.replicas: 1
spec.selector.matchLabels.app: mysql
spec.template.spec.containers[0].env[0].name: MYSQL_ROOT_PASSWORD
spec.template.spec.containers[0].image: mysql:8.0
spec.template.spec.volumes[0].persistentVolumeClaim.claimName: mysql-pvc
apiVersion: apps/v1
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **apiVersion**: `apps/v1`
- **replicas**: `1`
</details>

<details>
<summary>🧰 Usage</summary>

Deploying a stateful MySQL database with persistent storage in Kubernetes
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Always set resource limits for database containers
- Use secrets for sensitive environment variables like passwords
- Ensure PVC exists before deployment
- Consider using ConfigMaps for non-sensitive configuration
- Use specific image tags (not 'latest') for production
</details>


---

## 📄 File: `mysql-service.yaml`

> 📍 Path: `devops-pi-main\mysql-service.yaml`

<details>
<summary>🚀 Resource Summary</summary>

- **Kind:** `Service`
- **Description:** Defines a Kubernetes Service for MySQL with headless configuration

</details>

<details>
<summary>⚙️ Configuration Example for `Service`</summary>

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql
spec:
  selector:
    app: mysql
  ports:
    - port: 3306
      targetPort: 3306
  clusterIP: None
```
</details>

<details>
<summary>📑 Field Reference</summary>

<details>
<summary>`apiVersion`</summary>

- **Type:** `scalar`
- **Description:** Kubernetes API version
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (Service)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Service metadata
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** Service name
</details>

<details>
<summary>`spec`</summary>

- **Type:** `map`
- **Description:** Service specification
</details>

<details>
<summary>`spec.selector`</summary>

- **Type:** `map`
- **Description:** Pod selection labels
</details>

<details>
<summary>`spec.selector.app`</summary>

- **Type:** `scalar`
- **Description:** Label selector for MySQL pods
</details>

<details>
<summary>`spec.ports`</summary>

- **Type:** `array`
- **Description:** Exposed ports configuration
</details>

<details>
<summary>`spec.ports[0].port`</summary>

- **Type:** `scalar`
- **Description:** Service port number
</details>

<details>
<summary>`spec.ports[0].targetPort`</summary>

- **Type:** `scalar`
- **Description:** Target pod port number
</details>

<details>
<summary>`spec.clusterIP`</summary>

- **Type:** `scalar`
- **Description:** Cluster IP assignment
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
apiVersion: v1
kind: Service
metadata.name: mysql
spec.clusterIP: None
spec.ports[0].port: 3306
spec.ports[0].targetPort: 3306
spec.selector.app: mysql
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **apiVersion**: `v1`
- **spec.clusterIP**: `Automatically assigned if not specified`
</details>

<details>
<summary>🧰 Usage</summary>

Creates a headless Service for MySQL pods in Kubernetes (DNS resolution without load balancing)
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Use headless Services (clusterIP: None) for stateful applications
- Ensure port numbers match container exposed ports
- Selector labels must match pod labels exactly
</details>

