with open("day_12.txt") as f:
    commands = list(f.read().split())

#with open("day_12_exampl.txt") as f:
    #commands = list(f.read().split())

class ship:
    def __init__(self):
        self.x = 0
        self.y = 0
        self.face = "E"

    def move(self, direct, amount):
        if direct=="N":
            self.x += amount
        if direct=="S":
            self.x -= amount
        if direct=="E":
            self.y += amount
        if direct=="W":
            self.y -= amount
        if direct=="F":
            self.move(self.face, amount)
        if direct=="B":
            self.move(self.face, (-1*amount))
        if direct=="R" or direct=="L":
            self.rotate(direct, amount)

    def rotate(self, direct, amount):
        conn = ['E', 'S', 'W', 'N']
        conn_len = len(conn)
        index = conn.index(self.face)
        factor = amount/90
        print(self.face)
        if direct == "R":
            index = (index + factor) % conn_len
        if direct == "L":
            index = (index - factor) % conn_len
        self.face = conn[int(index)]
        print(self.face)

hms = ship()

for cmd in commands:
    direct = cmd[0]
    amount = int(cmd[1:])
    print(direct, amount)
    hms.move(direct, amount)
    print(hms.x, hms.y)
print(abs(hms.x) + abs(hms.y))
print("part 2")
