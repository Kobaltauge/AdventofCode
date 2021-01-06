player1 = [9,2,6,3,1]
player2 = [5,8,4,7,10]

player1 = [3,42,4,25,14,36,32,18,33,10,35,50,16,31,34,46,9,6,41,7,15,45,30,27,49]
player2 = [8,11,47,21,17,39,29,43,23,28,13,22,5,20,44,38,26,37,2,24,48,12,19,1,40]

win = []

round = 1

while True:
    cardp1 = player1.pop(0)
    cardp2 = player2.pop(0)

    if cardp1 > cardp2:
        player1.append(cardp1)
        player1.append(cardp2)
    elif cardp1 < cardp2:
        player2.append(cardp2)
        player2.append(cardp1)
    
    # print(player1)
    # print(player2)

    if len(player1) == 0:
        win = player2
        break
    elif len(player2) == 0:
        win = player1
        break

erg = 0
for i in range(1, len(win)+1):
    erg = erg + i * win.pop()

print(erg)
print("part 2")