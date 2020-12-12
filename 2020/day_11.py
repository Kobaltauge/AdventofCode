with open("day_11.txt") as f:
    map = list(f.read().split())

with open("day_11_exampl.txt") as f:
    map = list(f.read().split())

counter = 0

def check_adjacent(map, seaty, seatx):
    adja = []
    try:
        adja.append(map[seaty-1][seatx-1])
    except:
        pass
    try:
        adja.append(map[seaty-1][seatx])
    except:
        pass
    try:
        adja.append(map[seaty-1][seatx+1])
    except:
        pass
    try:
        adja.append(map[seaty][seatx-1])
    except:
        pass
    try:
        adja.append(map[seaty][seatx+1])
    except:
        pass
    try:
        adja.append(map[seaty+1][seatx-1])
    except:
        pass
    try:
        adja.append(map[seaty+1][seatx])
    except:
        pass
    try:
        adja.append(map[seaty+1][seatx+1])
    except:
        pass
    print(adja)
    return 






while True:
    for y in range(len(map)):
        for x in range(len(map[y])):
            check_adjacent(map, y, x)


            print(map[y][x])
print("part 2")