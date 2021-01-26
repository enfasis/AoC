import numpy as np


def manhattan_distance(x1, y1, x2, y2):
    return abs(x1 - x2) + abs(y1 - y2)


def solve():
    max_size = max(max(x, y) for x, y in data) + 1
    areas = [0] * len(data)
    total = 0
    for x, y in np.ndindex(max_size, max_size):
        distances = [manhattan_distance(x, y, a, b) for a, b in data]
        total += 1 if sum(distances) < 10000 else 0  # part 2
        min_ind, min_val = min(enumerate(distances), key=lambda v: v[1])
        areas[min_ind] += 1
        distances.pop(min_ind)
        if min_val in distances:
            areas[min_ind] -= 1
        if x == 0 or y == 0 or x == max_size - 1 or y == max_size - 1:
            areas[min_ind] -= max_size

    return max(areas), total


data = [
    tuple(map(int, line.strip().split(", "))) for line in open("input.txt").readlines()
]

part_1, part_2 = solve()

print(f"Silver: {part_1}\nGold: {part_2}")
