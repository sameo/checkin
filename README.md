## Introduction

checkin does parallel container launches on the client side
and logs each container start and run time stamps on the server side.

Both the server and the client must run on the same host.

## Server side

As root:

```
$ cd server
$ make
$ ./checkin-server
```

## Client side

As root:

```
$ cd client
$ export INSTANCES=50
$ export RUNTIME=cc-runtime
$ make prepare
$ make run
```

### Cleanup

To clean all client bundles up:

```
$ make clean
```

