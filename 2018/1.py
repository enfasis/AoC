data = [int(x) for x in open('input.txt').readlines()]
print(sum(data))
frecuencies = set()
current = 0
frecuencies.add(current)
twiceNotFound = True
while twiceNotFound:
    for i in data:
        current = current+i
        if current in frecuencies:
            twiceNotFound = False
            break
        frecuencies.add(current)
print(current)