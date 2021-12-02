with open("day_02.txt") as f:
    moves = [x.split() for x in f.read().splitlines()]

position = 0
depth = 0
other = 0

for move, step in moves:
    if move == "forward":
        position += int(step)
    elif move == "down":
        depth += int(step)
    elif move == "up":
        depth -= int(step)
    else:
        other += 1

print(f"position {position}\ndepth {depth}\nother {other}\nresult {position*depth}\n")

position = 0
depth = 0
aim = 0

for move, step in moves:
    if move == "forward":
        position += int(step)
        depth += int(step) * aim
    elif move == "down":
        aim += int(step)
    elif move == "up":
        aim -= int(step)

print(f"position {position}\ndepth {depth}\nresult {position*depth}")
