with open("day_03.txt") as f:
    map = list(f.read().split())

right = 3
down = 1
x, y = 0, 0
count = 0

while y < len(map):
    if map[y][x] == "#":
        count += 1
    x = x + right
    if x > (len(map[0]) - 1):
        x = x - (len(map[0]))
    y = y + down

print(count)

print("Part 2")

right = [1,3,5,7,1]
down = [1,1,1,1,2]
count_all = 1

for idx in range(len(right)):
    x, y = 0, 0
    count = 0
    while y < len(map):
        if map[y][x] == "#":
            count += 1
        x = x + right[idx]
        if x > (len(map[0]) - 1):
            x = x - (len(map[0]))
        y = y + down[idx]
    print(count)
    count_all = count * count_all

print(count_all)