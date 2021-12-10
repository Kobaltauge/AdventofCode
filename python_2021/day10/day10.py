with open("day_10.txt") as f:
    chunkes = f.read().splitlines()

brace_o, brace_c = 0,0
curl_o, curl_c = 0,0
arrow_o, arrow_c = 0,0
square_o, square_c = 0,0

brace = 0
curl = 0
arrow = 0
square = 0


for chunk in chunkes:
    for char in chunk:


    brace_o += chunk.count('(')
    brace_c += chunk.count(')')
    curl_o += chunk.count('{')
    curl_c += chunk.count('}')
    arrow_o += chunk.count('<')
    arrow_c += chunk.count('>')
    square_o += chunk.count('[')
    square_c += chunk.count(']') 
result = ""
print(f'result part 1: {result}')

