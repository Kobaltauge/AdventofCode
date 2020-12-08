with open("day_05.txt") as f:
    seatlist = f.read().split()

ids = []
for seat in seatlist:
    row = int(seat[:-3].replace("B", "1").replace("F", "0"), 2)
    column = int(seat[-3:].replace("R", "1").replace("L", "0"), 2)
    id = row * 8 + column
    ids.append(id)

ids.sort(reverse=True)

print(ids[0])

print("part 2")

for i in range(ids[-2], ids[1]):
    if (i+1 in ids) and (i-1 in ids) and (i not in ids):
        print(i)
