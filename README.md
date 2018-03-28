# quaggad

`vtyd` is openconfigd configuration gateway to kamuee-vty.

```
$ go get github.com/coreswitch/openconfigd/openconfigd
$ go get github.com/slankdev/vtyd
```

```
$ sudo ${GOPATH}/bin/openconfigd -y ${GOPATH}/src/github.com/coreswitch/openconfigd/yang &
$ sudo /usr/lib/quagga/bgpd --config_file /dev/null --pid_file /tmp/bgpd.pid --socket /var/run/zserv.api &
$ sudo ${GOPATH}/bin/quaggad &
```
