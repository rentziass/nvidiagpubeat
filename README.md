# Nvidiagpubeat

Welcome to Nvidiagpubeat.
Nvidiagpubeat is an elastic beat that uses NVIDIA System Management Interface (nvidia-smi) to monitor NVIDIA GPU devices and can ingest metrics into Elastic search cluster. nvidia-smi is a command line utility, based on top of the NVIDIA Management Library (NVML), intended to aid in the management and monitoring of NVIDIA GPU devices.

Nvidiagpu beat with help of nvidia-smi allows administrators to query GPU device state.  It is targeted at the TeslaTM, GRIDTM, QuadroTM and Titan X product, though limited support is also available on other NVIDIA GPUs.

NVIDIA-smi ships with NVIDIA GPU display drivers on Linux, and with 64bit Windows Server 2008 R2 and Windows 7.

Nvidiagpubeat provides ability (look at nvidiagpubeat.yml) to configure metrics that needs to be monitored and by default it queries utilization.gpu,utilization.memory,memory.total,memory.free,memory.used,temperature.gpu,pstate and can ingest them into elastic search cluster for possibly visualization using Kibana.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/deepujain`

## Getting Started with Nvidiagpubeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Nvidiagpubeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Nvidiagpubeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/deepujain/nvidiagpubeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Nvidiagpubeat run the command below. This will generate a binary
in the same directory with the name nvidiagpubeat.

```
make
```


### Run

To run Nvidiagpubeat with debugging output enabled, run:

```
./nvidiagpubeat -c nvidiagpubeat.yml -e -d "*"
```


### Test

To test Nvidiagpubeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/nvidiagpubeat.template.json and etc/nvidiagpubeat.asciidoc

```
make update
```


### Cleanup

To clean  Nvidiagpubeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Nvidiagpubeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/deepujain
cd ${GOPATH}/github.com/deepujain
git clone https://github.com/deepujain/nvidiagpubeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
