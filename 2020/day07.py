import re

with open("day_07.txt") as f:
    baglistraw = list(f.read().splitlines())

with open("day_07_exampl.txt") as f:
    baglistraw = list(f.read().splitlines())


baglist = {}

# (\w+ \w+) bags contain (((\d+)|(\bno\b)) (\w+ \w+) (bags|bag)[,|.]\s*)*$

for bag in baglistraw:
    bagreg = re.match(r"(\w+ \w+) bags contain (([\d+]) (\w+ \w+) (bags|bag)[,|.]\s*)*$", bag)
    print(bagreg[0])
    print(bagreg[1]) 

    print("ende")