## cmftp - A simple ftp server in go

The aim is to have a simple ftp-server that can run in userspace and have limited disk access

### Design goals

 * Handle the minimum required for FTP according to [RFC-959](http://tools.ietf.org/html/rfc959)
 * Run in userspace, listen on port 2121 (or other port over 1024)
 * Virtual users, restricted to subdirectories with a per user and total file quota.
 * Configurable via config file (probably [TOML](https://github.com/toml-lang/toml)) and web interface

#### To compile:
```
make
```

To run:

```
> ./cmftp
