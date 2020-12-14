import re

with open("day_14.txt") as f:
    instructions = f.readlines()
instructions = [i.replace('\n','').split(" = ") for i in instructions] 

memory = dict()
actmask = ""

for instr in instructions:
    if instr[0] == "mask":
        actmask = instr[1]
    if instr[0][:3] == "mem":
        memadr = instr[0][4:-1]
        newmemdec = instr[1]
        newmembin = list(str(f"{int(instr[1]):036b}"))
        masklist = list(actmask)
        for idx in range(len(newmembin)):
            if masklist[idx] == "X":
                pass
            elif masklist[idx] == "1":
                newmembin[idx] = "1"
            elif masklist[idx] == "0":
                newmembin[idx] = "0"
        newmembin = "".join(newmembin)
        memory[memadr] = (int(newmembin, 2), newmembin)

print(sum([int(d[0]) for d in memory.values()]))
print("part 2")

# with open("day_14_exampl_2.txt") as f:
#     instructions = f.readlines()
# instructions = [i.replace('\n','').split(" = ") for i in instructions] 

memory = dict()
actmask = ""

for instr in instructions:
    if instr[0] == "mask":
        actmask = instr[1]
    if instr[0][:3] == "mem":
        memadrorg = instr[0][4:-1]
        memadr = []
        masklist = list(actmask)
        for count in range(2**actmask.count('X')):
            binstring = list(str("{0:b}".format(count)).zfill(actmask.count('X')))
            binstringindex = 0
            memadrbin = list(str(f"{int(memadrorg):036b}"))
            # print("".join(memadrbin))
            for idx in range(len(memadrbin)):
                if masklist[idx] == "X":
                    memadrbin[idx] = binstring[binstringindex]
                    binstringindex += 1
                elif masklist[idx] == "1":
                    memadrbin[idx] = "1"
                elif masklist[idx] == "0":
                    pass
            # print("".join(memadrbin))
            memadrbin = "".join(memadrbin)
            memadrbin = int(memadrbin, 2)
            memadr.append(memadrbin)
        for m in memadr:
            memory[m] = instr[1]

print(sum([int(d) for d in memory.values()]))