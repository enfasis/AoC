import numpy as np
import re

data = [[int(x) for x in re.findall("\d+",line)] for line in open('input.txt').readlines()]
fabric = np.zeros((1000,1000))

for e in data:    
    for i in range(0,e[3]):
        for j in range(0,e[4]):
            fabric[e[1]+i,e[2]+j] += 1

cnt = 0
for i in fabric:
    for j in i:
        if j > 1:
            cnt += 1
print(cnt)

for e in data:
    overlap = False
    if not overlap:
        for i in range(0,e[3]):
            for j in range(0,e[4]):
                if fabric[e[1]+i,e[2]+j] > 1:
                    overlap = True
                    break
            if overlap:
                break
    
    if not overlap:
        print(e[0])
        break