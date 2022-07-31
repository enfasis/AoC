from collections import defaultdict


def part_1():
    """Lexicographical topological sort with khan's algorithm"""
    adj = defaultdict(list)
    for (b, a) in data:
        adj[b].append(a)

    in_degree = defaultdict(int)
    for _, adj_list in adj.items():
        for var in adj_list:
            in_degree[var] += 1

    multi_set = [key for key in adj.keys() if in_degree[key] == 0]

    top_order = list()

    while multi_set:
        multi_set.sort()
        vertex = multi_set[0]
        top_order.append(vertex)

        t_set = []

        for val in adj[vertex]:
            i_d = in_degree[val] - 1
            in_degree[val] = i_d
            if not i_d and val not in t_set:
                t_set.append(val)
        multi_set = multi_set[1:] + t_set

    return "".join(top_order)


def part_2():
    """Dumb solution"""
    childs = defaultdict(list)
    parents = defaultdict(list)
    for (b, a) in data:
        childs[b].append(a)
        parents[a].append(b)

    queue = [key for key in childs.keys() if key not in parents.keys()]

    working = [0 for _ in range(5)]
    work = [None for _ in range(5)]

    time = 0

    while queue or any(working):
        queue.sort()

        for i in range(5):
            working[i] = max(working[i] - 1, 0)
            if not working[i]:
                if work[i]:
                    for ch in childs[work[i]]:
                        if work[i] in parents[ch]:
                            parents[ch].remove(work[i])
                    queue += [
                        val
                        for val in childs[work[i]]
                        if not parents[val] and val not in work and val not in queue
                    ]
                if queue:
                    work_name = queue[0]
                    working[i] = 60 + ord(work_name) - ord("A") + 1
                    work[i] = work_name
                    queue = queue[1:]
        time += 1

    return time - 1


data = [(line[5], line[36]) for line in open("input.txt").readlines()]

print(f"Silver: {part_1()}\nGold: {part_2()}")
