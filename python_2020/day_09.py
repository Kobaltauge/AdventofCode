with open("day_09.txt") as f:
    xmas = f.read().splitlines()
    xmas = [int(i) for i in xmas]

# with open("day_09_exampl.txt") as f:
#     xmas = f.read().splitlines()
#     xmas = [int(i) for i in xmas]

preamblelenght = 25

for i in range(0, len(xmas)-preamblelenght):
    switch = True
    preamble = xmas[i:i+preamblelenght]
    checkno = xmas[i+preamblelenght]
    for c in preamble:
        search = checkno - c
        if not search == c:
            if search in preamble:
                switch = True
                break
            else:
                switch = False
    if not switch:
        print(checkno)
        break

print("done")
print("part 2")

start = 0
stop = 1
checkslice = []

# 144381670

while True:
    if sum(checkslice) == 144381670:
        print("got it")
        break
    elif sum(checkslice) > 144381670:
        start += 1
        stop = start + 1
    elif sum(checkslice) < 144381670:
        stop += 1
    checkslice = xmas[start:stop]

checkslice.sort()
code = checkslice[0] + checkslice[-1]
print(checkslice)
print(code)