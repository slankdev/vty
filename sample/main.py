#!/usr/bin/env python
import telnetlib
from telnetlib import Telnet

tn = Telnet('localhost', 2605)
tn.read_until(b'Password: ')
tn.write(b'zebra\n')
tn.expect(['.*> '])
tn.write('show memory\n')
res = tn.expect(['.*> '])[2]
tn.write(b'exit\n')

tL = res.split('\n')[1:-1]
for line in tL:
    print line
