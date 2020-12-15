line = [0,3,6]
line = [1,3,2]
line = [2,1,3]
line = [1,2,3]
line = [2,3,1]
line = [3,2,1]
line = [3,1,2]
line = [14,3,1,0,9,5]

turn = 4
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
print(line)    
print(line[-1])
print("part 2")

