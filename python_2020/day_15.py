# line = [0,3,6]
# line = [1,3,2]
# line = [2,1,3]
# line = [1,2,3]
# line = [2,3,1]
# line = [3,2,1]
# line = [3,1,2]
# turn = 4

line = [14,3,1,0,9,5]
turn = 7

while turn < 2021:
    newline = line.copy()
    lastno = line[-1]
    checkline = line[:-1]
    if checkline.count(lastno) == 0:
        newline.append(0)
    else:
        lastindex = (len(line) - 1 - line[::-1].index(lastno)) + 1
        lastindex2 = (len(checkline) - 1 - checkline[::-1].index(lastno)) + 1
        newline.append(lastindex - lastindex2)
    # print(newline)
    line = newline.copy()
    turn += 1
print(line[-1])

print("part 2")

# line = {0:[1],3:[2],6:[3]}
# turn = 4
# lastno = 6

# line = {1:[1],3:[2],2:[3]}
# turn = 4
# lastno = 2

line = {14:[1],3:[2],1:[3],0:[4],9:[5],5:[6]}
turn = 7
lastno = 5

while turn < 30000001:
    if lastno in line.keys():
        indexes = line[lastno]
        if len(indexes) == 1:
            lastno = 0
            if lastno in line.keys():
                line[lastno].append(turn)
                line[lastno] = line[lastno][-2:]
            else:
                line[lastno] = [turn]
        else:
            lastindex = indexes[-1]
            lastindex2 = indexes[-2]
            lastno = (lastindex - lastindex2)
            if lastno in line.keys():
                line[lastno].append(turn)
                line[lastno] = line[lastno][-2:]
            else:
                line[lastno] = [turn]
                lastno = (lastindex - lastindex2)
    else:
        line[lastno] = [turn]
        lastno = turn
    turn += 1

print(lastno)
