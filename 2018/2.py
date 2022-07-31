from collections import defaultdict


def part_1():
    twice = thrice = 0
    for word in data:
        cnt = defaultdict(int)
        for letter in word:
            cnt[letter] += 1
        twice += int(2 in cnt.values())
        thrice += int(3 in cnt.values())
    return twice * thrice


def part_2():
    while data:
        removed = data.pop()
        for word in data:
            diff = index = 0
            for key, val in enumerate(word):
                if diff > 1:
                    break
                if val != removed[key]:
                    diff += 1
                    index = key
            if diff == 1:
                return word[:index] + word[index + 1 :]


data = [str(x) for x in open("input.txt").readlines()]

print(f"Silver: {part_1()}\nGold: {part_2()}")
