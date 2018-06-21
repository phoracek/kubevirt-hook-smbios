# KubeVirt SMBIOS hook sidecar

Example VM definition:

```yaml
apiVersion: kubevirt.io/v1alpha2
kind: VirtualMachineInstance
metadata:
  creationTimestamp: null
  labels:
    special: vmi-smbios
  name: vmi-smbios
  annotations:
    # Request the hook sidecar
    hooks.kubevirt.io/hookSidecars: '[{"image": "phoracek/kubevirt-hook-smbios"}]'
    # Overwrite base board manufacturer name
    smbios.vm.kubevirt.io/baseBoardManufacturer: "Radical Edward"
spec:
  domain:
    devices:
      disks:
      - disk:
          bus: virtio
        name: registrydisk
        volumeName: registryvolume
    machine:
      type: ""
    resources:
      requests:
        memory: 64M
  terminationGracePeriodSeconds: 0
  volumes:
  - name: registryvolume
    registryDisk:
      image: registry:5000/kubevirt/cirros-registry-disk-demo:devel
status: {}
```
