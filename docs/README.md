# Containers From Scratch: The Deep Dive

This document serves as my learning journal while building a container runtime from first principles. Rather than using high-level tools like Docker, I wanted to understand the core Linux kernel features that make containers possible.

## Core Container Concepts Covered

- [x] Process Isolation (Linux Namespaces)
- [x] Resource Limitation (Control Groups)
- [x] Filesystem Isolation (chroot)
- [x] Process Information (Proc Filesystem)

## What Makes a Container?

At its essence, a container is nothing magical - just a regular process with some isolation characteristics:

1. It can't see other processes on the system
2. It has its own hostname and network configuration
3. It has a restricted view of the filesystem
4. It can only use a limited amount of system resources

## Linux Namespaces: The Foundation of Process Isolation

Namespaces partition kernel resources so that different processes see different views of the system:

- [x] **UTS Namespace**: Isolates hostname and domain name

  ```go
  // Set custom hostname for the container
  syscall.Sethostname([]byte("container"))
  ```

- [x] **PID Namespace**: Isolates process IDs

  ```go
  // Child process will see itself as PID 1
  syscall.CLONE_NEWPID
  ```

- [x] **Mount Namespace**: Isolates filesystem mount points
  ```go
  // Each container gets its own view of mounted filesystems
  syscall.CLONE_NEWNS
  ```

## Filesystem Isolation

Filesystem isolation uses chroot to change the apparent root directory for a process:

```go
// Change root directory to our container filesystem
syscall.Chroot("/path/to/container/root")

// Change working directory to the new root
syscall.Chdir("/")
```

This prevents the containerized process from accessing files outside its designated root.

## Resource Management with cgroups

Control groups (cgroups) limit, account for, and isolate resource usage:

```go
// Limit container to a maximum of 20 processes
ioutil.WriteFile(filepath.Join(pids, "your_name/pids.max"), []byte("20"), 0700)

// Enable automatic cleanup when container exits
ioutil.WriteFile(filepath.Join(pids, "your_name/notify_on_release"), []byte("1"), 0700)
```

## Proc Filesystem: A Process's Window to the World

Every container needs its own /proc filesystem to function properly:

```go
// Mount proc inside the container
syscall.Mount("proc", "proc", "proc", 0, "")

// Clean up when done
syscall.Unmount("proc", 0)
```

## Advanced Topics To Explore

- [ ] Network Namespaces and Container Networking
- [ ] User Namespaces for Better Security
- [ ] Container Orchestration Principles
- [ ] Image Format and Distribution
- [ ] Storage Drivers and Optimization
- [ ] Container Security Best Practices

## Personal Learnings

Building this project taught me that containers are just clever combinations of existing kernel features. While production container runtimes like Docker add many layers of convenience, abstraction, and security, the core principles remain the same.

The most challenging part was understanding how these different isolation mechanisms interact with each other, particularly when it comes to proper cleanup and error handling.

## References and Further Reading

- [Liz Rice's Talk: Containers From Scratch](https://www.youtube.com/watch?v=8fi7uSYlOdc)
- [Linux Namespaces](https://man7.org/linux/man-pages/man7/namespaces.7.html)
- [Control Groups v2](https://www.kernel.org/doc/html/latest/admin-guide/cgroup-v2.html)
- [OCI Container Specification](https://github.com/opencontainers/runtime-spec)
