import itertools

with open("day_14.txt") as f:
    start = f.readline().strip('\n')
    _ = f.readline()
    input = [x.split('->') for x in f.readlines()]
    input = [(i[0].strip(),i[1].strip('\n').strip()) for i in input]
    input = {input[i][0]: input[i][1] for i in range(0, len(input))}

output = start
for step in range(0,10):
    insertlist = []
    for i in range(0, len(output)-1):
        check = output[i]+output[i+1]
        if check in input:
            insertlist.append([i+1,input[check]])
    insertlist.reverse()
    for insert in insertlist:
        output = output[:insert[0]] + insert[1] + output[insert[0]:]
        # output.insert(insert[0], insert[1])

resultdict ={}
for l in output:
    resultdict[l] = output.count(l)
x = max(resultdict.values())
y = min(resultdict.values())
result = x-y
print(result)

output = start
for step in range(0,40):
    insertlist = []
    for i in range(0, len(output)-1):
        check = output[i]+output[i+1]
        if check in input:
            insertlist.append(input[check])
    lastone = output[-1]
    output = list(itertools.chain.from_iterable(zip(output,insertlist)))
    output.append(lastone)
    print(step)

resultdict ={}
for l in output:
    resultdict[l] = output.count(l)
x = max(resultdict.values())
y = min(resultdict.values())
result = x-y
print(result)
print("end")
 