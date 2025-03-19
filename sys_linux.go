//go:build linux

package main

import (
	"os/exec"
	"syscall"
)

// `CLONE_NEWUTS": Isolates the UTS (UNIX Timesharing System) namespace,
// which means the new process can have its own hostname.
// `CLONE_NEWPID": Isolates the PID namespace, allowing the child process to
// have a new process ID space, making it seem as if it's the only process
// running inside this namespace.
// `CLONE_NEWNS": Isolates the mount namespace, so that the child process can have
// its own view of the filesystem.
func setSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
}

// Sets the hostname of the isolated environment
func setHostName() error {
	return syscall.Sethostname([]byte("container"))
}

// Placeholder implementations - to be implemented in future commits
func isolateFilesystem() error {
	return nil
}

func mountProc() error {
	return nil
}

func unmountProc() error {
	return nil
}

func cg() error {
	return nil
}