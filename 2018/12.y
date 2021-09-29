def part_1():
    return sum(data)


def part_2():
    frecuencies = set()
    frecuency = 0
    again = True
    while again:
        for i in data:
            frecuency += i
            if frecuency in frecuencies:
                again = False
                break
            frecuencies.add(frecuency)
    return frecuency


data = [int(x) for x in open("input.txt").readlines()]


print(f"Silver: {part_1()}\nGold: {part_2()}")
