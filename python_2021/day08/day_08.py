import pandas as pd

with open("day_08.txt") as f:
    display = [x.split('|') for x in f.read().splitlines()]

count=[0,0,0,0]

for output in display:
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

digitdict =  {0:('abcefg'),
              1:('cf'),
              2:('acdeg'),
              3:('acdfg'),
              4:('bcdf'),
              5:('abdfg'),
              6:('abdefg'),
              7:('acf'),
              8:('abcdefg'),
              9:('abcdfg')}

def decode(input, output):
    decodedict = {'a':'','b':'','c':'','d':'','e':'','f':'','g':''}
    complete = ''.join(input)
    occur = pd.Series(list(complete)).value_counts()
    decodedict['e'] = occur.where(occur == 4).dropna().keys()[0]
    decodedict['f'] = occur.where(occur == 9).dropna().keys()[0]
    decodedict['b'] = occur.where(occur == 6).dropna().keys()[0]
    input = [word for word in input if len(word) != 4]
    complete = ''.join(input)
    complete = complete.replace(decodedict['e'], '').replace(decodedict['f'], '').replace(decodedict['b'], '')
    occur = pd.Series(list(complete)).value_counts()
    decodedict['a'] = occur.where(occur == 8).dropna().keys()[0]
    decodedict['d'] = occur.where(occur == 6).dropna().keys()[0]
    input = [word for word in input if len(word) != 2]
    complete = ''.join(input)
    complete = complete.replace(decodedict['a'], '').replace(decodedict['d'], '').replace(decodedict['e'], '').replace(decodedict['f'], '').replace(decodedict['b'], '')
    occur = pd.Series(list(complete)).value_counts()
    decodedict['c'] = occur.where(occur == 6).dropna().keys()[0]
    decodedict['g'] = occur.where(occur == 7).dropna().keys()[0]
    result = []
    for outdigit in output:
        res = ''
        for seg in outdigit:
            res += list(decodedict.keys())[list(decodedict.values()).index(seg)]
        res = ''.join(sorted(res, key=str.lower))
        res = list(digitdict.keys())[list(digitdict.values()).index(res)]
        result.append(res)
    return result

# decode(('abcefg','cf','acdeg','acdfg','bcdf','abdfg','abdefg','acf','abcdefg','abcdfg'), ('cf', 'acf'))

result = 0

for complete in display:
    res = 0
    input, output = complete[0].split(), complete[1].split()

    res = int("".join([str(integer) for integer in decode(input, output)]))
    result += res

print(f'result part 2: {result}')
