with open("day_12.txt") as f:
    commands = list(f.read().split())

# with open("day_12_exampl.txt") as f:
#     commands = list(f.read().split())

class ship:
    def __init__(self):
        self.x = 0
        self.y = 0

    def move(self, direct, amount):
        if direct=="x":
            self.x += amount
        if direct=="y":
            self.y += amount

class waypoint:
    def __init__(self):
        self.x = 10
        self.y = 1
    
    def move(self, direct, amount):
        if direct=="N":
            self.y += amount
        if direct=="S":
            self.y -= amount
        if direct=="E":
            self.x += amount
        if direct=="W":
            self.x -= amount

    def rotate(self, direct, amount):
        conn = ['+y', '+x', '-y', '-x']
        conn_len = len(conn)
        if self.x > 0:
            index1 = conn.index("+x")
        elif self.x < 0:
            index1 = conn.index("-x")
        else:
            index1 = 0
        if self.y > 0:
            index2 = conn.index("+y")
        elif self.y < 0:
            index2 = conn.index("-y")
        else:
            index2 = 0
        factor = int(amount/90)
        if direct == "R":
            index1 = (index1 + factor) % conn_len
            index2 = (index2 + factor) % conn_len
        if direct == "L":
            index1 = (index1 - factor) % conn_len
            index2 = (index2 - factor) % conn_len
        
        old_x = self.x
        old_y = self.y

        if conn[index1][:1] == "+":
            new_x = abs(old_x)
        else:
            new_x = -abs(old_x)
        if conn[index2][:1] == "+":
            new_y = abs(old_y)
        else:
            new_y = -abs(old_y)

        if conn[index1][1:] == "x":
            self.x = new_x
        else:
            self.x = new_y
        if conn[index2][1:] == "y":
            self.y = new_y
        else:
            self.y = new_x
        
        if old_x == 0:
            self.x = 0
        if old_y == 0:
            self.y = 0


hms = ship()
wp = waypoint()
counter = 0

for cmd in commands:
    counter += 1
    if hms.x > 4000:
        print(counter)
    direct = cmd[0]
    amount = int(cmd[1:])
    print(hms.x, hms.y)
    print(wp.x, wp.y)
    print(direct, amount)
    if direct=="F":
        hms.move('x', wp.x * amount)
        hms.move('y', wp.y * amount)
    if direct in ("N","E","S","W"):
        wp.move(direct, amount)
    if direct=="R" or direct=="L":
        wp.rotate(direct, amount)
    print(hms.x, hms.y)
    print(wp.x, wp.y)
    print("\n")
print(abs(hms.x) + abs(hms.y))
print("part 2")
