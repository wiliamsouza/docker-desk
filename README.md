docker-desk
===========

A Docker desktop app written in Go and qml.

It is in **ALPHA** stage.

Dependencies
------------

```
go get github.com/niemeyer/qml
```

Download
--------

```
go get github.com/wiliamsouza/docker-desk
```

Build
-----

```
cd $GOPATH/src/github.com/wiliamsouza/docker-desk
go build
```

Run
---

```
./docker-desk
```

You need run it as root if your user is not part of docker group.
