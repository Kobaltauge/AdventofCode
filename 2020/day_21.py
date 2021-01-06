with open("day_21.txt") as f:
    listraw = f.read().splitlines()

# with open("day_21_exampl.txt") as f:
#     listraw = f.read().splitlines()

allergendict = {}
allergenlist = []
completelist = []

for food in listraw:
    ingredients = (food.split("(")[0].strip()).split()
    completelist = completelist + ingredients
    allergens = (food.split("(")[1].replace("contains ", "").replace(")", "").replace(" ", "")).split(",")
    for allergen in allergens:
        if allergen in allergendict.keys():
            allergendict[allergen].append(ingredients)
        else:
            allergendict[allergen] = [ingredients]
    allergenlist = allergenlist + allergens

allergens = list(set(allergenlist))

for key,value in allergendict.items():
    allergen = key
    list1 = value[0]
    for idx in range(1, len(value)):
        list1 = list(set(list1).intersection(value[idx]))
    allergendict[key] = list1

workdict = allergendict.copy()
allergeniddict = {}

while True:
    allergenraw = [[k,v] for k,v in workdict.items() if len(v) == 1]
    if len(allergenraw)==0:
        break
    allergeniddict[allergenraw[0][0]] = next(iter(allergenraw[0][1]))
    todel = next(iter(allergenraw[0][1]))
    print(todel)
    [workdict[k].remove(todel) for k,v in workdict.items() if todel in v]

for element in allergeniddict.values():
    completelist = [item for item in completelist if item != element]

print(len(completelist))
print("part 2")

sol = ""

for i in [value for (key, value) in sorted(allergeniddict.items())]:
    sol = sol + "," + i
sol = sol[1:]
print(sol)

print("wait")

