with open("day_07.txt") as f:
    baglistraw = list(f.read().splitlines())

# with open("day_07_exampl.txt") as f:
    # baglistraw = list(f.read().splitlines())

def check_gold(bag):
    print(bag, bags[bag])
    if bags[bag][0][0] == 0:
        return 0
    else:
        inneradded = 0
        for bagitem in bags[bag]:
            innersum = check_gold(bagitem[1]) * bagitem[0] + bagitem[0]
            inneradded = inneradded + innersum
            print("{} contains {}".format(bagitem, inneradded))
        return inneradded

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

cpl_amount = check_gold("shiny gold")

print(cpl_amount)
