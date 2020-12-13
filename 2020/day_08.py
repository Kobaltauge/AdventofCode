with open("day_08.txt") as f:
    oplist = f.read().splitlines()

#with open("day_08_exampl.txt") as f:
    #oplist = f.read().splitlines()

accumulator = 0
accumulator_old = 0
index = 0
stack = []

while True:
    opraw = oplist[index]
    op = opraw[:3]
    data = opraw[4:]
    # print(index, op, int(data), accumulator_old, accumulator)
    stack.append(index)
    if op == "acc":
        accumulator_old = int(accumulator)
        accumulator += int(data)
        index += 1
    elif op == "jmp":
        index += int(data)
        if index in stack:
            break
    elif op == "nop":
        index += 1

print(accumulator)
print("part 2")

with open("day_08_new.txt") as f:
    oplist = f.read().splitlines()


def check(oplist):
    accumulator = 0
    index = 0
    stack = []

    while True:
        if index == 611:
            print("end")
            return accumulator
        opraw = oplist[index]
        op = opraw[:3]
        data = opraw[4:]
        # print(index, op, int(data), accumulator_old, accumulator)
        stack.append(index)
        if op == "acc":
            accumulator_old = int(accumulator)
            accumulator += int(data)
            index += 1
        elif op == "jmp":
            index += int(data)
            if index in stack:
                return False
        elif op == "nop":
            index += 1

for i in range(0,len(oplist)):
    oplistmod = oplist.copy()
    opraw = oplist[i]
    op = opraw[:3]
    data = opraw[4:]
    if op == "jmp":
        oplistmod[i] = str("nop "+data)
        solution = check(oplistmod)
        if solution:
            print(solution)
    elif op == "nop":
        oplistmod[i] = str("jmp "+data)
        solution = check(oplistmod)
        if solution:
            print(solution)
        
    