## Introduction

checkin does parallel container launches on the client side
and logs each container start and run time stamps on the server side.

Both the server and the client must run on the same host.

## Server side

```
$ cd server
$ make docker-run
```

## Client side

```
$ cd client
$ make prepare INSTANCES=50
$ make run RUNTIME=cc-runtime
```

### Cleanup

To clean all client bundles up:

```
$ make clean
```

