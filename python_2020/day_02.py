with open("day_02.txt") as f:
    passwords = list(f.read().splitlines())

amount = 0

for p in passwords:
    minmax, char, password = p.split()
    min, max = minmax.split("-")
    char = char.split(":")[0]
    if (int(min) <= password.count(char) <= int(max)):
        amount +=1
print(amount)

print("next")

amount = 0

for p in passwords:
    pos, char, password = p.split()
    pos1, pos2 = map(int, pos.split("-"))
    char = char.split(":")[0]
    if password[pos1-1] == char and password[pos2-1] != char:
        amount += 1
    if password[pos1-1] != char and password[pos2-1] == char:
        amount += 1
print(amount)