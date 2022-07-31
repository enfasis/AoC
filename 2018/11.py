import numpy as np


def get_power(x, y):
    return (((x + 10) * y + serial_number) * (x + 10) // 100) % 10 - 5


def get_identifier(initial_size=1, final_size=300):
    grid = np.zeros((grid_size + 1, grid_size + 1), dtype="int64")
    for x in range(1, grid_size + 1):
        for y in range(1, grid_size + 1):
            partial_sum = grid[y, x - 1] + grid[y - 1, x] - grid[y - 1, x - 1]
            grid[y, x] = get_power(x, y) + partial_sum

    max_power, size, pos = 0, 0, (0, 0)
    for s in range(initial_size, final_size):
        for x in range(1, grid_size - s + 1):
            for y in range(1, grid_size - s + 1):
                total = (
                    grid[y, x] - grid[y + s, x] - grid[y, x + s] + grid[y + s, x + s]
                )
                if total > max_power:
                    max_power, size, pos = total, s, (x + 1, y + 1)

    return pos, size


def part_1():
    (x, y), _ = get_identifier(3, 4)
    return x, y


def part_2():
    (x, y), s = get_identifier(3, 300)
    return x, y, s


grid_size = 300
serial_number = 9810

print(f"Silver: {part_1()}\nGold: {part_2()}")
