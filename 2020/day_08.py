import time

with open("day_08.txt") as f:
    oplist = f.read().splitlines()

with open("day_08_exampl.txt") as f:
    oplist = f.read().splitlines()

accumulator = 0
index = 0

while True:
    opraw = oplist[index]
    op = opraw[:3]
    data = opraw[4:]
    if op == "acc":
        accumulator_old = int(accumulator)
        accumulator += int(data)
        print(index, int(data), accumulator_old, accumulator)
        time.sleep(0.1)
        # if accumulator < 5:
        #     break
        index += 1
    elif op == "jmp":
        index += int(data)
    elif op == "nop":
        index += 1

print(accumulator_old)

