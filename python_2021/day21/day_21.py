class player:
    def __init__(self, pos):
        self.pos = pos
        self.score = 0
    def addscore(self, amount):
        self.score += amount
    def move(self, amount):
        self.pos += amount
        if self.pos > 10:
            self.pos = self.pos%10
        if self.pos == 0:
            self.pos = 10

class dice100:
    def __init__(self):
        self.value = 1
        self.rolled = 0

    def throw(self, amount):
        self.rolled += 1
        self.value += amount
        if self.value > 100:
            self.value -= 100

player1 = player(6)
player2 = player(2)

# player1 = player(4)
# player2 = player(8)

round = 1
turn = 1
dice = dice100()

while True:
    dicescore = 0 
    for i in range(0,3):
        dicescore += dice.value
        dice.throw(1)
    if turn == 1:
        player1.move(dicescore)
        player1.addscore(player1.pos)
        print(f"player1 {player1.pos} {player1.score}")
        turn = 2
    else:
        player2.move(dicescore)
        player2.addscore(player2.pos)
        print(f"player2 {player2.pos} {player2.score}")
        turn = 1
    if (player1.score >= 1000):
        winner = 1
        break
    elif (player2.score >= 1000):
        winner = 2
        break
    round += 1
if winner == 1:
    print(f"result: {dice.rolled * player2.score}")
else:
    print(f"result: {dice.rolled * player1.score}")