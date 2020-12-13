import copy

with open("day_11.txt") as f:
    seats = list(f.read().split())

width = 97
height = 96

# with open("day_11_exampl.txt") as f:
#     seats = list(f.read().split())
# width = 9
# height = 9

seats = [[char for char in r] for r in seats]

def check_adjacent(seats, seaty, seatx):
    adja = []
    if (0 <= seaty-1 <= height) and (0 <= seatx-1 <= width):
        adja.append(seats[seaty-1][seatx-1])
    if (0 <= seaty-1 <= height) and (0 <= seatx <= width):
        adja.append(seats[seaty-1][seatx])
    if (0 <= seaty-1 <= height) and (0 <= seatx+1 <= width):
        adja.append(seats[seaty-1][seatx+1])
    if (0 <= seaty <= height) and (0 <= seatx-1 <= width):
        adja.append(seats[seaty][seatx-1])
    if (0 <= seaty <= height) and (0 <= seatx+1 <= width):
        adja.append(seats[seaty][seatx+1])
    if (0 <= seaty+1 <= height) and (0 <= seatx-1 <= width):
        adja.append(seats[seaty+1][seatx-1])
    if (0 <= seaty+1 <= height) and (0 <= seatx <= width):
        adja.append(seats[seaty+1][seatx])
    if (0 <= seaty+1 <= height) and (0 <= seatx+1 <= width):
        adja.append(seats[seaty+1][seatx+1])
    # print(seats[seaty][seatx])
    seat = seats[seaty][seatx]
    if seat == ".":
        return "."
    elif (seat == "L") and (adja.count("#") == 0):
        return "#"
    elif (seat == "#") and (adja.count("#") >= 4):
        return "L"
    else:
        return seat
    
counter = 0

while True:
    newseats = copy.deepcopy(seats)
    for y in range(len(seats)):
        for x in range(len(seats[y])):
            newseats[y][x] = check_adjacent(seats, y, x)
    # print("\n")
    # [print(r) for r in seats]
    # print("\n")
    # [print(r) for r in newseats]
    counter += 1
    if not (newseats == seats):
        seats = copy.deepcopy(newseats)
#        [print(r) for r in seats]
#        print("\n")
    else:
        break
print(sum([s.count("#") for s in seats]))
print("part 2")