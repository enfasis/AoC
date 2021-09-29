from collections import deque, defaultdict


def play_game(max_players, last_marble):
    """
    I solve this using simple lists, but it is clearly that this
    problem can be solve with a doublly-linked list and my solution
    will take too much time to compute. I found another anwser,
    which is here and it's clear and concise using deque.
    """
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
