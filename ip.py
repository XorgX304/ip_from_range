#!/usr/bin/python
#-*- coding: utf-8 -*-
# Author: GH0st3rs

def parse_ip(ipRange):
    d = {'o1':[], 'o2':[], 'o3':[], 'o4':[]}
    ipRange = ipRange.split('.')
    mask = ''
    i = 1
    for item in ipRange:
        if '-' in item: d['o%d' %(i)] = [int(x) for x in item.split('-')]
        else: d['o%d' %(i)] = [int(item), int(item)]
        item = '%s'
        if len(mask)>0:
            if mask[-1] == '.': mask += item
            else: mask += '.' + item
        else: mask += item + '.'
        i += 1
    ipList = []
    for o1 in range(d['o1'][0], d['o1'][1]+1):
        for o2 in range(d['o2'][0], d['o2'][1]+1):
            for o3 in range(d['o3'][0], d['o3'][1]+1):
                for o4 in range(d['o4'][0], d['o4'][1]+1):
                    ipList.append(mask %(o1, o2, o3, o4))
    return ipList

def usage():
    print('Usage ./ip.py 192.168.1.0-255\n')

if __name__ == "__main__":
    import sys
    if len(sys.argv) > 0:
        lst = parse_ip(sys.argv[1])
        for item in lst: print(item)
    else: usage()
