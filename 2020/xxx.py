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
topkeylist = [key for ele in toplist for key,val in puzzleedgedict.items() if ele in val]

bottomlist = [i[1] for i in puzzleedgedict.values()]
bottomkeylist = [key for ele in bottomlist for key,val in puzzleedgedict.items() if ele in val]

leftlist = [i[2] for i in puzzleedgedict.values()]
leftkeylist = [key for ele in leftlist for key,val in puzzleedgedict.items() if ele in val]

rightlist = [i[3] for i in puzzleedgedict.values()]
rightkeylist = [key for ele in rightlist for key,val in puzzleedgedict.items() if ele in val]

tops = list(set(topkeylist) - set(bottomkeylist))
bottoms = list(set(bottomkeylist) - set(topkeylist))
lefts = list(set(leftkeylist) - set(rightkeylist))
rights = list(set(rightkeylist) - set(leftkeylist))

print("wait")

print("part 2")