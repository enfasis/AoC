from collections import defaultdict
from re import findall


def iterate(curr_state):
    curr_state = f"....{curr_state}...."
    new_state = ""
    for i in range(2, len(curr_state) - 2):
        new_state += mapping.get(curr_state[i - 2 : i + 3], ".")
    return f"..{new_state}.."


def compute_sum(state, n):
    return sum(i - n * 4 for i, v in enumerate(state) if v == "#")


def part_1():
    n = 20
    state = initial
    for _ in range(n):
        state = iterate(state)
    return compute_sum(state, n)


def part_2():
    N = int(50e9)
    state = initial
    prev_sum = 0
    increment = defaultdict(lambda: 0)
    for n in range(1, 1000):
        state = iterate(state)
        total = compute_sum(state, n)
        diff = total - prev_sum

        if increment[diff] > 10:
            return total + diff * (N - n)

        prev_sum = total
        increment[diff] += 1


initial, *pairs = findall(r"[.#]+", open("input.txt").read())
mapping = dict(zip(pairs[::2], pairs[1::2]))

print(f"Silver: {part_1()}\nGold: {part_2()}")
