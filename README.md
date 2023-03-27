# whatfs

A tool to check your file system.

**Setup**

```sh
curl -LO https://github.com/ebiiim/whatfs/releases/download/v0.1.0/whatfs_0.1.0_linux_amd64.tar.gz
tar zxvf whatfs_0.1.0_linux_amd64.tar.gz
# optional
sudo cp whatfs /usr/local/bin
```

**Usage**

```
$ ./whatfs
/home/me/current_dir,EXT4_SUPER_MAGIC

$ ./whatfs "/mnt/hoge/fuga"
/mnt/hoge/fuga,NFS_SUPER_MAGIC

$ docker run --rm ghcr.io/ebiiim/whatfs
/,OVERLAYFS_SUPER_MAGIC
```

**Note**

- `man statfs` to see magic numbers
