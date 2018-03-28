# Vty

`vtyd` is openconfigd configuration gateway to kamuee-vty.

install
```
$ go get github.com/coreswitch/openconfigd/openconfigd
$ go get github.com/slankdev/vty/vtyd
$ sudo cp $GOPATH/src/github.com/slankdev/vty/misc/kamuee_vtysh.py /usr/local/bin
```

start
```
$ sudo ${GOPATH}/bin/openconfigd &
$ sudo /usr/lib/quagga/bgpd --config_file /dev/null --pid_file /tmp/bgpd.pid --socket /var/run/zserv.api &
$ sudo ${GOPATH}/bin/vtyd &
```

enter cli
```
$ cli
hscr-dev> show author
Hiroki Shirokura aka slankdev
```


