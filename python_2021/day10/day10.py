with open("day_10.txt") as f:
    chunkes = f.read().splitlines()

chunkes_2 = chunkes.copy() 

bracketlist = []
illegalchar = []
for chunk in chunkes:
    for char in chunk:
        if char == '(' or char == '{' or char == '<' or char == '[':
            bracketlist.append(char)
        elif char == ')' and bracketlist[-1] == '(':
            del bracketlist[-1]
        elif char == '}' and bracketlist[-1] == '{':
            del bracketlist[-1]
        elif char == '>' and bracketlist[-1] == '<':
            del bracketlist[-1]
        elif char == ']' and bracketlist[-1] == '[':
            del bracketlist[-1]
        else:
            illegalchar.append(char)
            chunkes_2.remove(chunk)
            break

replacements = {')':3, ']':57, '}':1197, '>':25137}
replacer = replacements.get  # For faster gets.
result = sum([replacer(n, n) for n in illegalchar])

print(f'result part 1: {result}')

# Part 2

resultlist = []
for chunk in chunkes_2:
    bracketlist = []
    missingchar = []
    for char in chunk:
        if char == '(' or char == '{' or char == '<' or char == '[':
            bracketlist.append(char)
        elif char == ')' and bracketlist[-1] == '(':
            del bracketlist[-1]
        elif char == '}' and bracketlist[-1] == '{':
            del bracketlist[-1]
        elif char == '>' and bracketlist[-1] == '<':
            del bracketlist[-1]
        elif char == ']' and bracketlist[-1] == '[':
            del bracketlist[-1]
    missingchar = bracketlist.copy()
    missingchar.reverse()

    replacements = {'(':1, '[':2, '{':3, '<':4}
    replacer = replacements.get  # For faster gets.
    misspoints = [replacer(n, n) for n in missingchar]

    score = 0
    for mp in misspoints:
        score *= 5
        score += mp
    resultlist.append(score)

resultlist.sort()
result = resultlist[int(((len(resultlist)+1) / 2))-1]

print(f'result part 2: {result}')

