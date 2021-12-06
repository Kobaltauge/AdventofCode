with open("day_06.txt") as f:
    fishes = list(map(int, (f.read().split(','))))

# fishes = [3,4,3,1,2]

for day in range(0,80):
    fishes = [fish-1 for fish in fishes]
    breed = fishes.count(-1)
    for new in range(0,breed):
        fishes.append(8)
    fishes = [6 if fish == -1 else fish for fish in fishes]
print(f"result part 1: {len(fishes)}")

# Part 2

with open("day_06.txt") as f:
    fishes = list(map(int, (f.read().split(','))))

fishes = [3,4,3,1,2]

fishesarray =[0] * 9
for fish in fishes:
    fishesarray[fish+1] += 1

for day in range(0,256):
    breed = fishesarray.pop(0)
    fishesarray.append(breed)
    fishesarray[6] += breed
    # print(fishesarray)
    if day == 79:
        print(fishesarray)
        print(f"check part 1: {sum(fishesarray)}")
print(f"result part 2: {sum(fishesarray)}")