with open("day_21.txt") as f:
    listraw = f.read().splitlines()

foodlistdict ={}
idx = 0
for food in listraw:
    ingredients = (food.split("(")[0].strip()).split()
    allergens = (food.split("(")[1].replace("contains ", "").replace(")", "").replace(" ", "")).split(",")
    foodlistdict[str(idx)] = [ingredients, allergens]
    idx += 1

matrix = {}
for key in foodlistdict.keys():
    for code in foodlistdict[key][0]:
        if code in matrix.keys():
            matrix[code] += foodlistdict[key][1]
        else:
            matrix[code] = foodlistdict[key][1]

print("wait")

print("part 2")