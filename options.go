// Copyright © 2013, 2014, The Go-LXC Authors. All rights reserved.
// Use of this source code is governed by a LGPLv2.1
// license that can be found in the LICENSE file.

// +build linux

package lxc

import (
	"os"
)

// AttachOptions type is used for defining various attach options.
type AttachOptions struct {

	// Specify the namespaces to attach to, as OR'ed list of clone flags (syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS ...).
	Namespaces int

	// Specify the architecture which the kernel should appear to be running as to the command executed.
	Arch Personality

	// Cwd specifies the working directory of the command.
	Cwd string

	// UID specifies the user id to run as.
	UID int

	// GID specifies the group id to run as.
	GID int

	// If ClearEnv is true the environment is cleared before running the command.
	ClearEnv bool

	// Env specifies the environment of the process.
	Env []string

	// EnvToKeep specifies the environment of the process when ClearEnv is true.
	EnvToKeep []string

	// StdinFd specifies the fd to read input from.
	StdinFd uintptr

	// StdoutFd specifies the fd to write output to.
	StdoutFd uintptr

	// StderrFd specifies the fd to write error output to.
	StderrFd uintptr
}

// DefaultAttachOptions is a convenient set of options to be used.
var DefaultAttachOptions = AttachOptions{
	Namespaces: -1,
	Arch:       -1,
	Cwd:        "/",
	UID:        -1,
	GID:        -1,
	ClearEnv:   false,
	Env:        nil,
	EnvToKeep:  nil,
	StdinFd:    os.Stdin.Fd(),
	StdoutFd:   os.Stdout.Fd(),
	StderrFd:   os.Stderr.Fd(),
}

// TemplateOptions type is used for defining various template options.
type TemplateOptions struct {
	// Template specifies the name of the template.
	Template string

	// Backend specifies the type of the backend.
	Backend BackendStore

	// Distro specifies the name of the distribution.
	Distro string

	// Release specifies the name/version of the distribution.
	Release string

	// Arch specified the architecture of the container.
	Arch string

	// Variant specifies the variant of the image (default: "default").
	Variant string

	// Image server (default: "images.linuxcontainers.org").
	Server string

	// GPG keyid (default: 0x...).
	KeyID string

	// GPG keyserver to use.
	KeyServer string

	// Disable GPG validation (not recommended).
	DisableGPGValidation bool

	// Flush the local copy (if present).
	FlushCache bool

	// Force the use of the local copy even if expired.
	ForceCache bool

	// ExtraArgs provides a way to specify template specific args.
	ExtraArgs []string
}

// DownloadTemplateOptions is a convenient set of options for "download" template.
var DownloadTemplateOptions = TemplateOptions{
	Template: "download",
	Distro:   "ubuntu",
	Release:  "trusty",
	Arch:     "amd64",
}

// BusyboxTemplateOptions is a convenient set of options for "busybox" template.
var BusyboxTemplateOptions = TemplateOptions{
	Template: "busybox",
}

// UbuntuTemplateOptions is a convenient set of options for "ubuntu" template.
var UbuntuTemplateOptions = TemplateOptions{
	Template: "ubuntu",
}