import re

with open("day_04.txt") as f:
    passportsraw = f.read()

# with open("day_04_exampl.txt") as f:
#     passportsraw = f.read()

passportsraw = passportsraw.replace("\n\n", ";").replace("\n", ",").replace(" ", ",").split(";")
passportsraw = [p.split(",") for p in passportsraw]
passports = []
for p in passportsraw:
    dicttmp = {}
    for e in p:
        e = e.split(":")
        dicttmp[e[0]] = e[1]
    passports.append(dicttmp)

mandatory = ["byr" , "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]

def check(field, data):
    if field == "byr" and 1920 <= int(data) <= 2002:
        return True
    elif field == "iyr" and 2010 <= int(data) <= 2020:
        return True
    elif field == "eyr" and 2020 <= int(data) <= 2030:
        return True
    elif field == "hgt":
        unit = data[-2:]
        if data[-2:] == "cm" and 150 <= int(data[:-2]) <= 193:
            return True
        elif data[-2:] == "in" and 59 <= int(data[:-2]) <= 76:
            return True
        else:
            return False
    elif field == "hcl" and re.match(r'^#[0-9a-f]{6}$',data):
        return True
    elif field == "ecl" and data in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]:
        return True
    elif field == "pid" and len(data) == 9:
        return True
    else:
        return False

validcount = 0

for p in range(len(passports)):
    switch = True
    for m in mandatory:
        if not m in passports[p].keys():
            switch = False
            break
    if switch:
        validcount += 1

print(validcount)
print("Part 2")

validcount2 = 0

for p in range(len(passports)):
    switch = True
    for m in mandatory:
        if not m in passports[p].keys():
            switch = False
            break
        elif not check(m, passports[p][m]):
            switch = False
            break
    if switch:
        validcount2 += 1

print(validcount2)
