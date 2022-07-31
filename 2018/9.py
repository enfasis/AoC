from collections import defaultdict, deque


def play_game(max_players, last_marble):
    scores = defaultdict(int)
    circle = deque([0])

    for marble in range(1, last_marble + 1):
        if marble % 23 == 0:
            circle.rotate(7)
            scores[marble % max_players] += marble + circle.pop()
            circle.rotate(-1)
        else:
            circle.rotate(-1)
            circle.append(marble)

    return max(scores.values())


def part_1():
    return play_game(423, 71944)


def part_2():
    return play_game(423, 71944 * 100)


print(f"Silver: {part_1()}\nGold: {part_2()}")
