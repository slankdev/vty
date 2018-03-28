#!/usr/bin/env python
import telnetlib
from telnetlib import Telnet

tn = Telnet('localhost', 2605)
# tn.read_until(b'Password: ')
# tn.write(b'zebra\n')
tn.read_until('router> ')
tn.write('show memory\n')
res = tn.read_until('router> ')
tn.write(b'exit\n')

tL = res.split('\n')[1:-1]
for line in tL:
    print line
