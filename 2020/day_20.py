with open("day_20.txt") as f:
    raw = f.read().splitlines()

with open("day_20_exampl.txt") as f:
    raw = f.read().splitlines()

puzzledict = {}
for line in raw:

    if line == "":
        pass
    elif line[0] == "T":
        no = line[5:-1]
        puzzledict[str(no)] = []
    else:
        text = [line.replace(".", "0").replace("#", "1")]
        puzzledict[str(no)].append(text)
    
edges = []
puzzleedgedict = {}
for idx in puzzledict:
    top = int(puzzledict[idx][0][0], 2)
    bottom = int(puzzledict[idx][9][0], 2)
    left = int(''.join([puzzledict[idx][i][0][0] for i in range(len(puzzledict[idx]))]), 2)
    right = int(''.join([puzzledict[idx][i][0][9] for i in range(len(puzzledict[idx]))]), 2)
    puzzleedgedict[str(idx)] = (top, bottom, left, right)


toplist = [i[0] for i in puzzleedgedict.values()]
bottomlist = [i[1] for i in puzzleedgedict.values()]
leftlist = [i[2] for i in puzzleedgedict.values()]
rightlist = [i[3] for i in puzzleedgedict.values()]

tops = list(set(toplist) - set(bottomlist))
bottoms = list(set(bottomlist) - set(toplist))
lefts = list(set(leftlist) - set(rightlist))
rights = list(set(rightlist) - set(leftlist))

topkeylist = [key for ele in tops for key,val in puzzleedgedict.items() if ele == val[0]]
bottomkeylist = [key for ele in bottoms for key,val in puzzleedgedict.items() if ele == val[1]]
leftkeylist = [key for ele in lefts for key,val in puzzleedgedict.items() if ele == val[2]]
rightkeylist = [key for ele in rights for key,val in puzzleedgedict.items() if ele == val[3]]


print("wait")

print("part 2")