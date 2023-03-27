# whatfs

A tool to check your file system.

`man statfs` to see magic numbers.

```
$ ./whatfs
/home/me/current_dir,EXT4_SUPER_MAGIC

$ ./whatfs "/mnt/hoge/fuga"
/mnt/hoge/fuga,NFS_SUPER_MAGIC
```
