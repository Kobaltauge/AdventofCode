doorpublic = 5764801
cardpublic = 17807724

doorpublic = 3418282
cardpublic = 8719412

loop = 1
subjectdoor = 7
subjectcard = 7
valuedoor = 1
valuecard = 1

while True:
    valuedoor = valuedoor * subjectdoor
    valuedoor = valuedoor % 20201227
    valuecard = valuecard * subjectcard
    valuecard = valuecard % 20201227
    if valuecard == doorpublic:
        loopcard = loop
    if valuedoor == cardpublic:
        loopdoor = loop
    loop += 1
    if 'loopcard' in locals() and 'loopdoor' in locals():
        break

print(loopcard)
print(loopdoor)

loop = 1
valuedoor = 1
valuecard = 1

while True:
    valuedoor = valuedoor * cardpublic
    valuedoor = valuedoor % 20201227
    valuecard = valuecard * doorpublic
    valuecard = valuecard % 20201227
    if loop == loopcard:
        encrcard = valuedoor
    if loop == loopdoor:
        encrdoor = valuecard
    loop += 1
    if 'encrcard' in locals() and 'encrdoor' in locals():
        break

print(encrcard)
print(encrdoor)

