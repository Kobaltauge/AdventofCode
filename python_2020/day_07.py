with open("day_07.txt") as f:
    baglistraw = list(f.read().splitlines())

# with open("day_07_exampl.txt") as f:
#     baglistraw = list(f.read().splitlines())

def check_gold(bag):
    if "shiny gold" in bag:
        return True
    else:
        if any([check_gold(bagitem) for bagitem in bags[bag]]):
            return True

bags = {}

for bag in baglistraw:
    bagsplit = bag.split("bags contain")
    bagcolor = bagsplit[0].strip()
    bagcont = bagsplit[1].replace(".", "").split(",")
    bagcont = [e.split()[1] + " " + e.split()[2] for e in bagcont]
    if bagcont[0] != "other bags":
        bags[bagcolor] = bagcont
    else:
        bags[bagcolor] = []

count = 0
gesamt = 0
for bag in bags:
    if check_gold(bag):
        count += 1
    gesamt += 1
    print(gesamt)

# -1 because of the recursion the shiny gold bag is countet too
print(count -1)
