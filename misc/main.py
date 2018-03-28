#!/usr/bin/env python
import sys, os
import telnetlib

if len(sys.argv) < 2:
    print('Usage: {} cmd'.format(sys.argv[0]))
    quit()

cmd = sys.argv[1]

tn = telnetlib.Telnet('localhost', 2605)
tn.read_until(b'Password: ')
tn.write(b'zebra\n')
tn.expect(['.*> '])
tn.write(cmd + '\n')
res = tn.expect(['.*> '])[2]
tn.write(b'exit\n')

tL = res.split('\n')[1:-1]
for line in tL:
    print line
