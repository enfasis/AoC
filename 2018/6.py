import numpy as np

def mdist(x1,y1,x2,y2):
    return abs(x1-x2) + abs(y1-y2)

data = [list(map(int,line.strip().split(', '))) for line in open('input.txt').readlines()]
ds = len(data)
axy= [0]*ds
s  = 0

xs = [x[0] for x in data]
ys = [x[1] for x in data]

rs = max(xs)
cs = max(ys)
ms = max(rs,cs)+1

for x in range(0,ms):
    for y in range(0,ms):
        d = [mdist(x,y,d[0],d[1]) for d in data]
        #part 2
        if sum(d) < 10000:
            s += 1
        #part 1
        md = min(d)
        idx = d.index(md)
        axy[idx] += 1
        d.remove(md)
        if md in d:
            axy[idx] -= 1
        if x == 0 or y == 0 or x == ms-1 or y == ms-1:
            axy[idx] -= ms

print(max(axy))
print(s)