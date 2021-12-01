import math

values = {"departure_location":[[31,201],[227,951]], \
"departure_station":[[49,885],[892,961]], \
"departure_platform":[[36,248],[258,974]], \
"departure_track":[[37,507],[527,965]], \
"departure_date":[[37,331],[351,970]], \
"departure_time":[[38,370],[382,970]], \
"arrival_location":[[33,686],[711,960]], \
"arrival_station":[[46,753],[775,953]], \
"arrival_platform":[[34,138],[154,959]], \
"arrival_track":[[26,167],[181,961]], \
"classes":[[43,664],[675,968]], \
"duration":[[47,603],[620,954]], \
"price":[[40,290],[313,972]], \
"route":[[37,792],[799,972]], \
"row":[[32,97],[115,954]], \
"seat":[[25,916],[942,966]], \
"train":[[39,572],[587,966]], \
"typees":[[25,834],[858,953]], \
"wagon":[[48,534],[544,959]], \
"zone":[[47,442],[463,969]]}

my_ticket = [127,83,79,197,157,67,71,131,97,193,181,191,163,61,53,89,59,137,73,167]

with open("day_16.txt") as f:
    nearbys = list(f.read().split())

nearbys = [e.split(",") for e in nearbys]
nearbys = [[int(i[e]) for e in range(len(i))] for i in nearbys]

a = c = 1000
b = d = 0

for [[i1,i2],[j1,j2]] in values.values():
    if a > i1:
        a = i1
    if b < i2:
        b = i1
    if c > j1:
        c = j1
    if d < j2:
        d = j2  

scan_err = 0
right_nearbys = []
for nearby in nearbys:
    valid = True
    for no in nearby:
        scan_err_single = [no for [[a,b],[c,d]] in values.values() if (not ((a <= no <= b) or (c <= no <= d)))]
        if len(scan_err_single) == 20:
            scan_err += scan_err_single[0]
            valid = False
    if (valid):
        right_nearbys.append(nearby)
print(len(right_nearbys))
print(scan_err)

print("part 2")

reorder = {}
for e in range(20):
    reorder[str(e)] = []

for nearby in right_nearbys:
    for field in range(20):
        reorder[str(field)].append(nearby[field])

sorter = {}
for key in values.keys():
    sorter[key] = [i for i in range(20)]    

counter1 = 0

for key in values.keys():
    [[a,b],[c,d]] = values[key]
    for i in range(20):
        check = [i for no in reorder[str(i)] if (((a <= no <= b) or (c <= no <= d)))]
        # print(len(check))
        if len(check) == 190:
            # print("{}:{}".format(i, key))
            sorter[key].remove(i)
            counter1 += 1

fields = [k for k in sorted(sorter, key=lambda k: len(sorter[k]))]

answerdict = {}
answerdict[fields[0]] = 20
answerlist = []
counter = 1

while counter < 20:
    x = [x for x in sorter if (len(sorter[x])) == counter][0]
    list1 = list(set(sorter[x]) - set(answerlist))
    answerdict[x] = list1[0]
    answerlist.append(list1[0])
    counter += 1

indices = [val for key, val in answerdict.items() if 'departure' in key] 

answerlist = []
for i in indices:
    answerlist.append(my_ticket[i-1])

answer = math.prod(answerlist)
print(answer)

print("end")