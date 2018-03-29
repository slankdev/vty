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

## Support Commands

normal mode
```
$ show port
$ show port WORD
$ show port all
$ show vport
$ show vport WORD
$ show vport all
$ show memory mempool
$ show memory mempool all
$ show memory mempool arp
$ show memory mempool WORD
$ show acl port WORD
$ show running status
$ show version
```

configure mode
```
# set/delete interfaces interface 0 mirror 3
```

