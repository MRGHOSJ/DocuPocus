# 📄 YAML Configuration: `devops-pi-main`

[← Back to Overview](../README.md)

## 📄 File: `backend-deployment.yaml`

> 📍 Path: `devops-pi-main\backend-deployment.yaml`

🚀 Resource Summary

- **Kind:** `Deployment`
- **Description:** Kubernetes Deployment configuration for a Spring Boot application

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
- **Description:** Pod template specification
</details>

<details>
<summary>`spec.template.spec.containers`</summary>

- **Type:** `array`
- **Description:** Container configurations
</details>

<details>
<summary>`spec.template.spec.containers[].ports`</summary>

- **Type:** `array`
- **Description:** Container port mappings
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
spec.template.spec.containers[0].image: meowster/skillexchange-spring-app:latest
spec.template.spec.containers[0].env[0].name: SPRING_DATASOURCE_URL
spec.replicas: 2
spec.selector.matchLabels.app: backend
spec.template.spec.containers[0].ports[0].containerPort: 8084
apiVersion: apps/v1
kind: Deployment
metadata.name: backend
spec.template.metadata.labels.app: backend
spec.template.spec.containers[0].name: spring-app
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **apiVersion**: `apps/v1`
- **kind**: `Deployment`
- **spec.replicas**: `1`
</details>

<details>
<summary>🧰 Usage</summary>

Defines a Kubernetes deployment for a Spring Boot application with MySQL database connectivity
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Use specific image tags instead of 'latest' in production
- Store sensitive data (passwords) in Secrets rather than environment variables
- Define resource limits for containers
- Use ConfigMaps for environment configuration when possible
- Consider liveness and readiness probes for health checks
</details>


---

## 📄 File: `backend-service.yaml`

> 📍 Path: `devops-pi-main\backend-service.yaml`

🚀 Resource Summary

- **Kind:** `Service`
- **Description:** Defines a Kubernetes Service for backend application traffic routing

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
- **Description:** Port configurations
</details>

<details>
<summary>`spec.ports[].port`</summary>

- **Type:** `scalar`
- **Description:** Service exposed port
</details>

<details>
<summary>`spec.ports[].targetPort`</summary>

- **Type:** `scalar`
- **Description:** Pod target port
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
spec.selector.app: backend
spec.ports[0].port: 8084
spec.ports[0].targetPort: 8084
spec.type: ClusterIP
apiVersion: v1
kind: Service
metadata.name: backend
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **spec.type**: `ClusterIP`
- **spec.ports[0].targetPort**: `(same as port)`
</details>

<details>
<summary>🧰 Usage</summary>

Exposes backend pods internally via ClusterIP on port 8084
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Keep targetPort identical to port when possible
- Use explicit selectors matching pod labels
- Prefer ClusterIP for internal services
</details>


---

## 📄 File: `frontend-deployment.yaml`

> 📍 Path: `devops-pi-main\frontend-deployment.yaml`

🚀 Resource Summary

- **Kind:** `Deployment`
- **Description:** Kubernetes Deployment configuration for a frontend application

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
<summary>`spec.template.metadata.labels`</summary>

- **Type:** `map`
- **Description:** Pod template labels
</details>

<details>
<summary>`spec.template.spec.containers`</summary>

- **Type:** `array`
- **Description:** Container specifications
</details>

<details>
<summary>`spec.template.spec.containers[].name`</summary>

- **Type:** `scalar`
- **Description:** Container name
</details>

<details>
<summary>`spec.template.spec.containers[].image`</summary>

- **Type:** `scalar`
- **Description:** Container image
</details>

<details>
<summary>`spec.template.spec.containers[].ports`</summary>

- **Type:** `array`
- **Description:** Container ports
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
kind: Deployment
metadata.name: frontend
spec.replicas: 2
spec.selector.matchLabels.app: frontend
spec.template.spec.containers[0].name: angular-frontend
spec.template.spec.containers[0].env[0].name: DOCKER_APIURL
apiVersion: apps/v1
spec.template.metadata.labels.app: frontend
spec.template.spec.containers[0].image: meowster/angular-app:latest
spec.template.spec.containers[0].ports[0].containerPort: 80
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **apiVersion**: `apps/v1`
- **spec.replicas**: `1`
</details>

<details>
<summary>🧰 Usage</summary>

Deploying a scalable frontend application with environment variables
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Always specify resource limits for containers
- Use immutable image tags (avoid 'latest') in production
- Match selector labels with pod template labels
- Define readiness/liveness probes for containers
</details>


---

## 📄 File: `frontend-service.yaml`

> 📍 Path: `devops-pi-main\frontend-service.yaml`

🚀 Resource Summary

- **Kind:** `Service`
- **Description:** Defines a Kubernetes Service to expose an application

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
- **Description:** Port configurations
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
- **Description:** Node port (when type=NodePort)
</details>

<details>
<summary>`spec.type`</summary>

- **Type:** `scalar`
- **Description:** Service type (ClusterIP/NodePort/LoadBalancer)
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
apiVersion: v1
kind: Service
metadata.name: frontend
spec.selector.app: frontend
spec.ports[0].port: 80
spec.ports[0].targetPort: 80
spec.ports[0].nodePort: 30080
spec.type: NodePort
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **spec.type**: `ClusterIP`
- **apiVersion**: `v1`
</details>

<details>
<summary>🧰 Usage</summary>

Exposes frontend pods on port 30080 via NodePort
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Use meaningful metadata.name
- Match selector.app with pod labels
- Avoid privileged ports (<1024) for nodePort
- Prefer ClusterIP for internal services
</details>


---

## 📄 File: `ingress.yaml`

> 📍 Path: `devops-pi-main\ingress.yaml`

🚀 Resource Summary

- **Kind:** `Ingress`
- **Description:** Defines routing rules for external access to Kubernetes services

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
- **Description:** Target service definition
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata.name: app-ingress
spec.ingressClassName: nginx
spec.rules[0].http.paths[0].path: /skillExchange
spec.rules[0].http.paths[0].backend.service.name: backend
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **pathType**: `ImplementationSpecific`
</details>

<details>
<summary>🧰 Usage</summary>

Configures external HTTP access to multiple services with path-based routing
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Always specify pathType explicitly
- Use descriptive path names
- Limit wildcard paths ('/') to avoid conflicts
- Match ingressClassName with your cluster's controller
</details>


---

## 📄 File: `mysql-deployment.yaml`

> 📍 Path: `devops-pi-main\mysql-deployment.yaml`

🚀 Resource Summary

- **Kind:** `StorageClass`
- **Description:** Defines a Kubernetes StorageClass for persistent volumes

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
- **Description:** Metadata including StorageClass name
</details>

<details>
<summary>`provisioner`</summary>

- **Type:** `scalar`
- **Description:** Volume provisioner plugin
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

Defines storage provisioning behavior for PersistentVolumeClaims
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Use explicit provisioners for production environments
- Consider WaitForFirstConsumer mode for topology-aware scheduling
- Name StorageClasses meaningfully (avoid 'default')
</details>

🚀 Resource Summary

- **Kind:** `PersistentVolume`
- **Description:** Defines a PersistentVolume for cluster storage

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
- **Description:** PV name
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
- **Description:** Host machine path mapping
</details>

<details>
<summary>`spec.hostPath.path`</summary>

- **Type:** `scalar`
- **Description:** Filesystem path on host
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
spec.storageClassName: manual-hostpath
spec.hostPath.path: /mnt/data/mysql
apiVersion: v1
kind: PersistentVolume
metadata.name: mysql-pv
spec.capacity.storage: 2Gi
spec.accessModes: [ReadWriteOnce]
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **apiVersion**: `v1`
- **accessModes**: `[ReadWriteOnce]`
</details>

<details>
<summary>🧰 Usage</summary>

Provisioning persistent storage for applications like databases
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Avoid hostPath in production (use cloud volumes/NFS)
- Explicitly set reclaimPolicy
- Match accessModes with workload requirements
</details>

🚀 Resource Summary

- **Kind:** `PersistentVolumeClaim`
- **Description:** Defines a PersistentVolumeClaim for storage allocation

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
- **Description:** API version (v1)
</details>

<details>
<summary>`kind`</summary>

- **Type:** `scalar`
- **Description:** Resource type (PersistentVolumeClaim)
</details>

<details>
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Object metadata
</details>

<details>
<summary>`metadata.name`</summary>

- **Type:** `scalar`
- **Description:** PVC identifier
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
<summary>`spec.resources.requests.storage`</summary>

- **Type:** `scalar`
- **Description:** Storage capacity request
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
metadata.name: mysql-pvc
spec.accessModes: [ReadWriteOnce]
spec.storageClassName: manual-hostpath
spec.resources.requests.storage: 2Gi
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **apiVersion**: `v1`
- **kind**: `PersistentVolumeClaim`
</details>

<details>
<summary>🧰 Usage</summary>

Requesting persistent storage for applications like databases
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Use explicit storageClassName (avoid default)
- Align accessModes with workload requirements
- Set realistic storage requests
</details>

🚀 Resource Summary

- **Kind:** `Deployment`
- **Description:** Kubernetes Deployment configuration for MySQL with persistent storage

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
<summary>`metadata`</summary>

- **Type:** `map`
- **Description:** Resource metadata including name
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
- **Description:** Pod template specification
</details>

<details>
<summary>`spec.template.spec.containers`</summary>

- **Type:** `array`
- **Description:** Container configurations
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
metadata.name: mysql
spec.replicas: 1
spec.template.spec.containers[0].image: mysql:8.0
spec.template.spec.containers[0].env[0].name: MYSQL_ROOT_PASSWORD
spec.template.spec.containers[0].ports[0].containerPort: 3306
apiVersion: apps/v1
kind: Deployment
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

- Never store sensitive data like passwords directly in YAML (use Secrets)
- Specify resource limits for containers
- Use specific image tags (avoid 'latest')
- Consider readiness/liveness probes for database containers
- Use proper storage class for production databases
</details>


---

## 📄 File: `mysql-service.yaml`

> 📍 Path: `devops-pi-main\mysql-service.yaml`

🚀 Resource Summary

- **Kind:** `Service`
- **Description:** Defines a Kubernetes Service for MySQL with headless configuration

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
- **Description:** Exposed ports
</details>

<details>
<summary>`spec.ports[0].port`</summary>

- **Type:** `scalar`
- **Description:** Service port
</details>

<details>
<summary>`spec.ports[0].targetPort`</summary>

- **Type:** `scalar`
- **Description:** Pod port
</details>

<details>
<summary>`spec.clusterIP`</summary>

- **Type:** `scalar`
- **Description:** Cluster IP setting
</details>

</details>

<details>
<summary>🔍 Examples</summary>

```yaml
kind: Service
metadata.name: mysql
spec.selector.app: mysql
spec.ports[0].port: 3306
spec.ports[0].targetPort: 3306
spec.clusterIP: None
apiVersion: v1
```
</details>

<details>
<summary>🌐 Defaults</summary>

- **apiVersion**: `v1`
- **clusterIP**: `Automatically assigned if not 'None'`
</details>

<details>
<summary>🧰 Usage</summary>

Creates a headless Service for MySQL pods in Kubernetes
</details>

<details>
<summary>⚠️ Edge Cases</summary>

- Use headless Services (clusterIP: None) for stateful applications
- Ensure selector matches pod labels exactly
- Explicitly define targetPort for clarity
</details>

