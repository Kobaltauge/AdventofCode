import re

with open("day_14.txt") as f:
    instructions = f.readlines()
instructions = [i.replace('\n','').split(" = ") for i in instructions] 

memory = dict()

for instr in instructions:
    actmask = ""
    if instr[0] == "mask":
        actmask = instr[1]
    if instr[0][:3] == "mem":
        newmem = instr[0][4:-2]
        if newmem in memory.keys():
            pass
        else:
            # memory[newmem] = instr[1]
            memory[newmem] = str(f"{int(instr[1]):036b}")
print("part 2")