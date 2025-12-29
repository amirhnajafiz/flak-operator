# FLAK Kubernetes Operator

A Kubernetes operator to setup and run FLAK tracers. File Logging & Access Kernel-tracer (aka FLAK) is an eBPF-based tracing tool that monitors file access patterns over regular I/O operations and memory map operations.

## TODO

- [ ] Pod creation webhook
  - Add finalizer
  - Add init container (30s delay)
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
