def parse(data):
    children, entries, *data = data
    totals = 0
    values = []

    for _ in range(children):
        total, value, data = parse(data)
        totals += total
        values.append(value)

    totals += sum(data[:entries])

    if children:
        return (
            totals,
            sum(values[i - 1] for i in data[:entries] if i > 0 and i <= len(values)),
            data[entries:],
        )
    return totals, sum(data[:entries]), data[entries:]


data = [int(number) for number in open("/mnt/d/input.txt").read().split(" ")]
part_1, part_2, _ = parse(data)

print(f"Silver: {part_1}\nGold: {part_2}")
