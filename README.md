# Container From Scratch

A minimal container runtime implementation built in Go. This project replicates core container functionality to demonstrate how containers work under the hood.

## Why Build This?

I created this project to deeply understand container technology by implementing it from first principles. Containers aren't magic - they're just clever uses of Linux kernel features like namespaces and cgroups!

## Implemented Features:

- [x] Process isolation using Linux namespaces (`CLONE_NEWUTS`, `CLONE_NEWPID`, `CLONE_NEWNS`)
- [x] Custom hostname within container
- [x] Filesystem isolation via `chroot`
- [x] Process filesystem (`/proc`) mounting
- [x] Resource limitations using cgroups
- [x] Cross-platform compatibility (graceful fallbacks for non-Linux systems)

## Planned Enhancements:

- [ ] Network namespace isolation
- [ ] User namespace implementation for better security
- [ ] Volume mounting support
- [ ] Custom container image format
- [ ] Better error handling and recovery
- [ ] Performance optimizations for resource usage
- [ ] Support for container metadata and labels

## Usage

### For Linux Users:

```bash
# Run a bash shell inside a container
go run main.go run /bin/bash
```

### For Non-Linux Users (via Docker):

#### 1. Install Docker

Download and install [Docker Desktop](https://www.docker.com/products/docker-desktop).

#### 2. Build the Docker Image

```bash
docker build -t container-from-scratch .
```

#### 3. Run the Container

```bash
docker run --rm -it --privileged container-from-scratch ./main run /bin/bash
```

> **Note:** The `--privileged` flag is necessary because we're running container technology inside a container.

## How It Works

This project implements containerization from scratch using:

1. **Process Isolation**: Linux namespaces create isolated process environments
2. **Filesystem Isolation**: `chroot` creates a separate root filesystem
3. **Resource Control**: cgroups limit the resources used by processes
4. **Mount Management**: Proper mounting/unmounting of the proc filesystem

Check the `learning_docs` folder for detailed explanations of the underlying concepts!

## Limitations

This is an educational project and doesn't implement all features of production container runtimes like Docker. Use for learning, not for production workloads.
