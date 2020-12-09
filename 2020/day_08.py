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

