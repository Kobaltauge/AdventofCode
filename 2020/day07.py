import re

with open("day_07.txt") as f:
    baglistraw = list(f.read().splitlines())

with open("day_07_exampl.txt") as f:
    baglistraw = list(f.read().splitlines())

def check_white_yellow(bag, bagcont, count):
    # print(bag,bagcont)
    if "bright white" in bag:
        # print(bag,bagcont)
        return True
    elif "muted yellow" in bag:
        # print(bag,bagcont)
        return True
    elif ("dark orange" in bag) and (any("bright white" in c for c in bagcont) or any("muted yellow" in c for c in bagcont)):
        # print(bag,bagcont)
        return True
    elif ("light red" in bag) and (any("bright white" in c for c in bagcont) or any("muted yellow" in c for c in bagcont)):
        # print(bag,bagcont)
        return True
    else:
        for bagitemraw in bagcont:
            bagitem = bagitemraw.split()[1] + " " + bagitemraw.split()[2]
            if not bagitem == "other bags":
                if bagitem in bags:
                    if check_white_yellow(bagitem, bags[bagitem], count):
                        return True
            else:
                return False
        return False

    # if any("bright white" in c for c in bagcont):
    #     count +=1
    #     return count
    # elif bag in bags:
    #     check_white_yellow(bag, bagcont, count)
    #     return count

bags = {}

for bag in baglistraw:
    bagsplit = bag.split("bags contain")
    bagcolor = bagsplit[0].strip()
    bagcont = bagsplit[1].replace(".", "").split(",")
    bags[bagcolor] = bagcont

count = 0
gesamt = 0
valid = set()
for bag in bags:
    bagcont = bags[bag]
    if check_white_yellow(bag, bagcont, count):
        valid.add(bag)
    gesamt += 1
    print(gesamt)

print(len(valid))
