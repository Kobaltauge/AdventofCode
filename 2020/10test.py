with open("day_10.txt") as f:
    adaptorlist = f.read().splitlines()

# with open("day_10_exampl.txt") as f:
#     adaptorlist = f.read().splitlines()

adaptorlist = [int(i) for i in adaptorlist]
adaptorlist.sort()

counter = 0
counter_1 = 0
counter_3 = 0
outlet = 0
for index in range(len(adaptorlist)):
    if (adaptorlist[index] - outlet) == 1:
        counter_1 += 1
        counter += 1
        outlet = adaptorlist[index]
    elif (adaptorlist[index] - outlet) == 3:
        counter_3 += 1
        counter += 1
        outlet = adaptorlist[index]
if counter == len(adaptorlist):
    print("all used")
else:
    print("remain")

print(adaptorlist[-1])
print(counter)
print(counter_1, counter_3+1)
print(counter_1 * (counter_3+1))
print("part 2")


counter = 0


adaptorgraph = {}
adaptmulti = {}

for index in range(len(adaptorlist)):
    adaptorgraph[str(adaptorlist[index])] = []
    adaptmulti[str(index)] = []
    if str(adaptorlist[index]-1) in adaptorgraph.keys():
        adaptorgraph[str(adaptorlist[index]-1)].append(adaptorlist[index])
        adaptmulti[str(index)].append(adaptorlist[index])
    if str(adaptorlist[index]-2) in adaptorgraph.keys():
        adaptorgraph[str(adaptorlist[index]-2)].append(adaptorlist[index])
        adaptmulti[str(index)].append(adaptorlist[index])
    if str(adaptorlist[index]-3) in adaptorgraph.keys():
        adaptorgraph[str(adaptorlist[index]-3)].append(adaptorlist[index])
        adaptmulti[str(index)].append(adaptorlist[index])



print(adaptmulti)
print("end")


