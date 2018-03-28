#!/usr/bin/env python
import sys, os
import telnetlib
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('-c', action='append')
ret = parser.parse_args(sys.argv[1:])
commands = ret.c

tn = telnetlib.Telnet('localhost', 2605)

# entering password
tn.read_until(b'Password: ')
tn.write(b'zebra\n')
tn.expect(['.*> '])

# set terminal length 0
tn.write('terminal length 0\n')
tn.expect(['.*> '])

# insert enable-mode
tn.write('enable\n')
tn.read_until(b'Password: ')
tn.write(b'zebra\n')
tn.expect(['.*# '])

# execute command
allline = []
for cmd in commands:
    tn.write(cmd + '\n')
    res = tn.expect(['.*# '])[2]
    tL = res.split('\n')[1:-1]
    for line in tL:
        allline.append(line)

# exit
tn.write(b'exit\n')

for line in allline:
    print line


