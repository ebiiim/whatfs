package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func main() {

	var wd string

	if len(os.Args) > 1 {
		wd = os.Args[1]
	} else {
		wd, _ = os.Getwd()
	}

	var stat unix.Statfs_t
	if err := unix.Statfs(wd, &stat); err != nil {
		panic(err)
	}

	var t string
	switch stat.Type {
	default:
		t = fmt.Sprintf("0x%x", stat.Type)
	// filesystems
	case unix.BTRFS_SUPER_MAGIC:
		t = "BTRFS_SUPER_MAGIC"
	case unix.FUSE_SUPER_MAGIC:
		t = "FUSE_SUPER_MAGIC"
	case unix.EXT4_SUPER_MAGIC:
		t = "EXT4_SUPER_MAGIC"
	// network filesystems
	case unix.NFS_SUPER_MAGIC:
		t = "NFS_SUPER_MAGIC"
	case unix.AFS_SUPER_MAGIC:
		t = "AFS_SUPER_MAGIC"
	case unix.SMB2_SUPER_MAGIC:
		t = "SMB2_SUPER_MAGIC"
	case unix.CIFS_SUPER_MAGIC:
		t = "CIFS_SUPER_MAGIC"
	// special filesystems
	case unix.TMPFS_MAGIC:
		t = "TMPFS_MAGIC"
	case unix.SOCKFS_MAGIC:
		t = "SOCKFS_MAGIC"
	case unix.CGROUP2_SUPER_MAGIC:
		t = "CGROUP2_SUPER_MAGIC"
	case unix.CGROUP_SUPER_MAGIC:
		t = "CGROUP_SUPER_MAGIC"
	case unix.BPF_FS_MAGIC:
		t = "BPF_FS_MAGIC"
	case unix.DEVPTS_SUPER_MAGIC:
		t = "DEVPTS_SUPER_MAGIC"
	case unix.PROC_SUPER_MAGIC:
		t = "PROC_SUPER_MAGIC"
	case unix.OVERLAYFS_SUPER_MAGIC:
		t = "OVERLAYFS_SUPER_MAGIC"
	case unix.PIPEFS_MAGIC:
		t = "PIPEFS_MAGIC"
	}

	fmt.Printf("%s,%s", wd, t)
}
