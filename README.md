## cmftp - A simple ftp server in go

The aim is to have a simple ftp-server that can run in userspace and have limited disk access

### Design goals

 * Handle the minimum required for FTP according to RFC-959
 * Run in userspace, listen on pårt 2121 (or other)
 * Virtual users restricted to a subdirectory and file quota
 * Configurable via config file (TOML) and web interface
 

To compile:
```
make
```

To run:

```
> ./cmftp
