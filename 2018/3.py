import numpy as np
import re


def part_1():
    for e in data:
        for x, y in np.ndindex(e[3], e[4]):
            fabric[e[1] + x, e[2] + y] += 1
    return sum(1 for e in np.nditer(fabric) if e > 1)


def part_2():
    for e in data:
        overlap = False

        for x, y in np.ndindex(e[3], e[4]):
            if fabric[e[1] + x, e[2] + y] > 1:
                overlap = True
                break

        if not overlap:
            return e[0]


data = [
    [int(x) for x in re.findall("\d+", line)]
    for line in open("/mnt/d/input.txt").readlines()
]
fabric = np.zeros((1000, 1000))

print(f"Silver: {part_1()}\nGold: {part_2()}")
