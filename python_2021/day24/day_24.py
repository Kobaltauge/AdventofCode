import time

with open("day_24.txt") as f:
    instructions = [x.split() for x in f.read().splitlines()]
    instructions = [[int(i) if i.lstrip('-').isnumeric() else i for i in j] for j in instructions]

def check(modelnr):
    modelnrlist = list(map(int, list(str(modelnr))))
    idx = 0
    variables = {'w':0,'x':0,'y':0,'z':0}
    for inst in instructions:
        # print(inst)
        if inst[0] == "inp":
            variables[inst[1]] = modelnrlist[idx]
            idx += 1
        elif inst[0] == "add":
            a = [inst[1] if type(inst[1])==int else variables[inst[1]]][0]
            b = [inst[2] if type(inst[2])==int else variables[inst[2]]][0]
            variables[inst[1]] = a + b
        elif inst[0] == "mul":
            a = [inst[1] if type(inst[1])==int else variables[inst[1]]][0]
            b = [inst[2] if type(inst[2])==int else variables[inst[2]]][0]
            variables[inst[1]] = a * b
        elif inst[0] == "div":
            a = [inst[1] if type(inst[1])==int else variables[inst[1]]][0]
            b = [inst[2] if type(inst[2])==int else variables[inst[2]]][0]
            if not b == 0:
                variables[inst[1]] = int(a / b)
        elif inst[0] == "mod":
            a = [inst[1] if type(inst[1])==int else variables[inst[1]]][0]
            b = [inst[2] if type(inst[2])==int else variables[inst[2]]][0]
            if not (a < 0) or (b <= 0):
                variables[inst[1]] = a % b
        elif inst[0] == "eql":
            a = [inst[1] if type(inst[1])==int else variables[inst[1]]][0]
            b = [inst[2] if type(inst[2])==int else variables[inst[2]]][0]
            if a == b:
                variables[inst[1]] = 1
            else:
                variables[inst[1]] = 0
        # print(variables)
    return variables


modelnr   = 11111111111111
modelnr   = 99999999999999
# modelnr = 13579246899999

while True:
    # print(modelnr)
    while "0" in list(str(modelnr)):
        modelnr -= 1
    variables = check(modelnr)
    print(modelnr)
    print(variables)
    time.sleep(0.4)
    if variables['z'] == 0:
        break
    modelnr -= 1

print(f'result part 1: {modelnr}')