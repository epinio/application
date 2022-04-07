
The Epinio application CRD was created to replace the not-maintained [App CRD](https://github.com/kubernetes-sigs/application).

Epinio doesn't run a controller for this CRD. It has it's own API which updates the application resource in an imperative way.
The Application Custom resource is used only as a "placeholder" when an application is created without a workload. It's also
used as a parent for the various other resources. This allows automatic cascading deletion of resources.

## Building

This project was created with [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder). You should read the docs of that project if you are planning to
do some heavy work (e.g. create a controller etc). The simplest thing you may want to do is to add some field to the `AppSpec` struct in [app_types.go](api/v1/app_types.go).

After you do so, you will have to generate the CRD yaml to be used in Epinio.
And possibly the C code as well.

You can do that with a single command:

```
make generate manifests
```

## Making more CRDS

Requires `kubebuilder` and `controller-gen`. To install:

```
cd /path/to/bin/directory
wget https://github.com/kubernetes-sigs/kubebuilder/releases/download/v3.3.0/kubebuilder_linux_amd64
mv kubebuilder_linux_amd64 kubebuilder
chmod u+x ./kubebuilder
./kubebuilder --version
```

and

```
wget https://github.com/kubernetes-sigs/controller-tools/archive/refs/tags/v0.4.1.tar.gz
tar xf v0.4.1.tar.gz 
cd controller-tools-0.4.1/
go install ./cmd/controller-gen/
```

For each new CRD `MY_CRD` then run

```
kubebuilder create api --group application --version v1 --kind MY_CRD --resource true --force
```

where `MY_CRD` is the Kind of the new CRC.
