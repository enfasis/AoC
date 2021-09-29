def generate(data):
    data = "...." + data + "...."
    new_string = ".."
    for i in range(len(data) - 4):
        found = False
        for rule in rules:
            if data[i : i + 5] == rule[0]:
                new_string += rule[1]
                found = True
                break
        new_string += "." if not found else ""
    return new_string + ".."


def compute_sum(state, generations):
    total = 0
    for i, v in enumerate(state):
        total += i - generations * 4 if v == "#" else 0
    return total


def part_1():
    # n = 20
    # state = initial_state
    # for _ in range(n):
    #     state = generate(state)
    # return compute_sum(state, n)
    return 1


def part_2():
    n = 100
    state = initial_state
    for i in range(1, n + 1):
        state = generate(state)
        print(i, compute_sum(state, i))

    return 2


data = open("/mnt/d/input.txt").read().split("\n\n")
initial_state = data[0].split(" ")[2]
rules = [rule.split(" => ") for rule in data[1].split("\n")]

print(f"Silver: {part_1()}\nGold: {part_2()}")
