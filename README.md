## Introduction

checkin does parallel container launches on the client side
and logs each container start and run time stamps on the server side.

Both the server and the client must run on the same host.

## Server side

`$ make docker-run`

## Client side

```
$ make prepare INSTANCES=50
$ make run RUNTIME=cc-runtime
```

## Cleanup

`$ make clean`

