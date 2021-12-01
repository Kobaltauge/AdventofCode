with open("day_01.txt") as f:
    numbers = list(map(int, (f.read().splitlines())))

for no in numbers:
    diff = 2020 - no
    if diff in numbers:
        print(no*diff)

print("next")

for no1 in numbers:
    for no2 in numbers:
        diff = 2020 - no1 - no2
        if diff in numbers:
            print(no1*no2*diff)
