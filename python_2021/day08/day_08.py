with open("day_08.txt") as f:
    moves = [x.split('|') for x in f.read().splitlines()]

count=[0,0,0,0]

for output in moves:
    output = output[1].split()
    for outdigit in output:
        if len(outdigit) == 2:
            count[0] += 1
        elif len(outdigit) == 3:
            count[2] +=1
        elif len(outdigit) == 4:
            count[1] += 1
        elif len(outdigit) == 7:
            count[3] += 1
result = sum(count)
print(f'result part 1: {result}')

digitdict =  {0:('a','b','c','e','f','g'),
                1:('c','f'),
                2:('a','c','d','e','g'),
                3:('a','c','d','f','g'),
                4:('b','c','d','f'),
                5:('a','b','d','f','g'),
                6:('a','b','d','e','f','g'),
                7:('a','c','f'),
                8:('a','b','c','d','e','f','g'),
                9:('a','b','c','d','f','g')}

decodedict = {'a':'','b':'','c':'','d':'','e':'','f':'','g':''}

def decode(input, output):
    while not any(decodedict.values()):
        decodedict = [word for word in words if len(word) == size]
        pass

for move in moves:
    result = decode(move[0].split(), move[1].split())
print(f'result part 2: {result}')
