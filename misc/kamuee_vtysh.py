#!/usr/bin/env python
import sys, os
import telnetlib
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('-c', action='append')
ret = parser.parse_args(sys.argv[1:])
commands = ret.c

tn = telnetlib.Telnet('localhost', 9077)

# set terminal length 0
tn.expect(['.*> '])
tn.write('terminal length 0\n')
tn.expect(['.*> '])

# execute commands
allline = []
for cmd in commands:
    tn.write(cmd + '\n')
    res = tn.expect(['.*> '])[2]
    tL = res.split('\n')[1:-1]
    for line in tL:
        allline.append(line)

# exit telnet
tn.write(b'exit\n')

for line in allline:
    print line


