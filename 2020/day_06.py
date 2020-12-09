with open("day_06.txt") as f:
    answersraw = f.read()

#with open("day_06_exampl.txt") as f:
    #answersraw = f.read()

answers = answersraw.replace("\n\n", ";").replace("\n", "").split(";")
count = 0

for answer in answers:
    count += len(set(answer))

print(count)
print("part 2")

answersraw = answersraw.replace("\n\n", ";").replace("\n", ",").split(";")
answers = [p.split(",") for p in answersraw]
count = 0

for answer in answers:
    anslist = list(answer[0])
    for i in range(1,len(answer)):
        anslist = list(set(anslist).intersection(list(answer[i])))
    count += len(anslist)
print(count)
