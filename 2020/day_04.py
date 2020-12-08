with open("day_04.txt") as f:
    passportsraw = f.read()

passportsraw = passportsraw.replace("\n\n", ";").replace("\n", ",").replace(" ", ",").split(";")
passportsraw = [p.split(",") for p in passportsraw]
passports = []
for p in passportsraw:
    dicttmp = {}
    for e in p:
        e = e.split(":")
        dicttmp[e[0]] = e[1]
    passports.append(dicttmp)

mandatory = [byr, iyr, eyr, ion, hgt, hcl, ecl, pid]

valid = 0

for p in range(len(passports)):
    if all()

print("ende")