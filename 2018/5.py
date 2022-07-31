def part_1(pol):
    i = 0
    while i != len(pol) - 1:
        if abs(ord(pol[i]) - ord(pol[i + 1])) == 32:
            pol = pol[:i] + pol[i + 2 :]
            i = i - 1 if i != 0 else 0
        else:
            i += 1
    return len(pol)


def part_2():
    shortest = len(data)
    for i in range(65, 90):
        polymer = data.replace(str(chr(i)), "").replace(str(chr(i + 32)), "")
        polymer_length = part_1(polymer)
        if shortest > polymer_length:
            shortest = polymer_length
    return shortest


data = open("/mnt/d/input.txt").read().splitlines()[0]
print(f"Silver: {part_1(data)}\nGold: {part_2()}")
