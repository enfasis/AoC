import numpy as np
import re
from collections import defaultdict
from datetime import datetime, timedelta
from operator import attrgetter


class Guard:
    def __init__(self, id):
        self.id = id
        self.sleeping_time = 0
        self.started_time = None
        self.history = []
        self.most_sleeping_minute = 0
        self.most_sleeping_frecuency = 0

    def sleep(self, time):
        self.started_time = time

    def wake(self, time):
        curr_sleeping_time = int((time - self.started_time).seconds / 60)
        self.sleeping_time += curr_sleeping_time
        self.history.append([self.started_time.minute, curr_sleeping_time])

    def calculate_sleeping_minute(self):
        watch = {i: 0 for i in range(0, 60)}
        for [minute, sleep_time] in self.history:
            for m in range(minute, minute + sleep_time):
                watch[m] += 1

        minute = max(watch.keys(), key=lambda x: watch[x])
        self.most_sleeping_minute = minute
        self.most_sleeping_frecuency = watch[minute]


def part_1():
    last_guard = None
    for record in data:
        if record["guard"]:
            last_guard = guards[record["guard"].id]
        elif record["sleep"]:
            last_guard.sleep(record["time"])
        elif record["wake"]:
            last_guard.wake(record["time"])

    guard = max(guards.values(), key=attrgetter("sleeping_time"))
    guard.calculate_sleeping_minute()
    return guard.id * guard.most_sleeping_minute


def part_2():
    for guard in guards.values():
        guard.calculate_sleeping_minute()
    guard = max(guards.values(), key=lambda x: x.most_sleeping_frecuency)
    return guard.id * guard.most_sleeping_minute


data = []
guards = defaultdict(Guard)
for x in open("/mnt/d/input.txt").readlines():
    str_datetime = re.search(r"(?<=\[).*(?=\])", x).group()
    if "Guard" in x:
        guard_id = int(re.search(r"#\d+", x).group()[1:])
        guard = Guard(guard_id)
        guards[guard_id] = guard
    else:
        guard = None
    record = {
        "time": datetime.strptime(str_datetime, "%Y-%m-%d %H:%M"),
        "guard": guard,
        "wake": "wakes" in x,
        "sleep": "falls" in x,
    }
    data.append(record)

data = sorted(data, key=lambda x: x["time"])

print(f"Silver: {part_1()}\nGold: {part_2()}")
