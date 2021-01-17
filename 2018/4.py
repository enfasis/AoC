import numpy as np
from datetime import datetime, timedelta
data = []
for x in open('input.txt').readlines():
    a = [x[6:17]]
    if x[19] == 'f':
        a.append(1)
    elif x[19] == 'w':
        a.append(0)
    else:
        a.append(x[26:30])
    data.append(a)

data = sorted(data, key= lambda x: datetime.strptime(x[0], '%m-%d %H:%M'))
d = {}
for x in data:
    if int(x[0][6:8]) == 23:
        key = format(datetime.strptime(x[0], '%m-%d %H:%M') +timedelta(hours=1), '%m-%d')
    else:
        key = str(x[0][0:5])    
    if isinstance(x[1],str):
        time = np.zeros(60)
        d.update({key:[x[1],time]})
    else:
        for i in range(int(x[0][9:11]),60):            
            d[key][1][i] = x[1]

guards = {}
for i in d:
    if d[i][0] in guards:
        guards[d[i][0]] = np.add(guards[d[i][0]],d[i][1])
    else:
        guards.update({d[i][0]:d[i][1]})
maxAs = 0
keyAs = ''
for key in guards:
    if sum(guards[key])>maxAs:
        maxAs = sum(guards[key])
        keyAs = key

time = guards[keyAs][0]
itime= 0
for i in range(0,60):
    if guards[keyAs][i] > time:
        time = guards[keyAs][i]
        itime = i

print(itime*int(keyAs))

guard = ''
minute= 0
maxtime= 0
for key in guards:
    for m in range(0,60):
        if guards[key][m] > maxtime:
            maxtime = guards[key][m]
            minute = m
            guard = key

print(minute*int(guard))