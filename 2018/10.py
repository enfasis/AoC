import re


def solve():
    rows, cols, m_x, m_y, m_time, m_size = 0, 0, 0, 0, 0, 100000
    for time in range(20000):
        min_x, max_x, min_y, max_y = 10000, 0, 10000, 0
        for x, y, v_x, v_y in data:
            min_x = min(min_x, x + time * v_x)
            max_x = max(max_x, x + time * v_x)
            min_y = min(min_y, y + time * v_y)
            max_y = max(max_y, y + time * v_y)

        size = (max_x - min_x) * (max_y - min_y)
        if size < m_size:
            rows, cols = max_y - min_y, max_x - min_x
            m_x, m_y, m_time, m_size = min_x, min_y, time, size

    grid = [[" "] * (cols + 1) for _ in range(rows + 1)]

    for x, y, v_x, v_y in data:
        grid[y + m_time * v_y - m_y][x + m_time * v_x - m_x] = "#"

    # Solve part 1
    print("Silver:")
    for row in grid:
        print("".join(row))

    # Solve part 2
    print("Gold:", m_time)


data = [
    list(map(int, re.findall(r"-?\d+", line)))
    for line in open("/mnt/d/input.txt").readlines()
]

solve()
