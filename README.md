# FLAK Kubernetes Operator

A Kubernetes operator to setup and run FLAK tracers.

## TODO

- [ ] Pod creation webhook
  - Add finalizer
  - Add init container
  - Create FLAK tracer pod
- [ ] Pod update webhook
  - Check for existing FLAK tracer pod
  - Recreate if not exist or failed
- [ ] Pod deletion webhook
  - Remove FLAK tracer pod
  - Remove finalizer
- [ ] Prometheus metrics
- [ ] FLAK tuning parameters
  - PVC and StorageClass
  - Service
  - Service Account
  - Role
  - Role Binding
