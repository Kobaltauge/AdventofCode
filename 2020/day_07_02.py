with open("day_07.txt") as f:
    baglistraw = list(f.read().splitlines())

# with open("day_07_exampl.txt") as f:
#     baglistraw = list(f.read().splitlines())

def check_gold(bag, amount):
    global cpl_amount
    print(bag, bags[bag])
    if bags[bag][0][0] == 0:
        return True
    else:
        if all ([check_gold(bagitem[1], bagitem[0]) for bagitem in bags[bag]]):
            added = sum([int(bagitem[0]) for bagitem in bags[bag]])
            cpl_amount += amount * added
        return

bags = {}

for bag in baglistraw:
    bagsplit = bag.split("bags contain")
    bagcolor = bagsplit[0].strip()
    bagcont = bagsplit[1].replace(".", "").split(",")
    bagcont = [list((([0 if e.split()[0]=="no" else int(e.split()[0])][0]), e.split()[1] + " " + e.split()[2])) for e in bagcont]
    if bagcont[0] != "other bags":
        bags[bagcolor] = bagcont
    else:
        bags[bagcolor] = []

cpl_amount = 0
check_gold("shiny gold", 0)

print(cpl_amount)
