# https://www.redblobgames.com/grids/hexagons/
# https://www.techwithtim.net/tutorials/game-development-with-python/pygame-tutorial/pygame-tutorial-movement/

import pygame
from math import sin, cos, pi
import os
position = 50, 50
os.environ['SDL_VIDEO_WINDOW_POS'] = str(position[0]) + "," + str(position[1])

with open("day_24.txt") as f:
    directionsrawarray = f.read().splitlines()

# with open("day_24_exampl.txt") as f:
#     directionsrawarray = f.read().splitlines()

directionsarray = []
for directionsraw in directionsrawarray:
    idx = 0
    directions = []
    while idx < len(directionsraw):
        getdir = directionsraw[idx]
        if getdir == "e" or getdir == "w":
            directions.append(str(getdir))
            idx += 1
        elif getdir == "n" or getdir == "s":
            directions.append(str(getdir) + str(directionsraw[idx+1]))
            idx += 2
    directionsarray.append(directions)

# directions = ("se","se","nw","ne","ne","ne","w","se","e","sw","w","sw","sw","w","ne","ne","w","se","w","sw")
# directions = ("nw","nw","nw","nw","nw","nw","nw")
# directionsarray = [("nw","w","sw","e","e")]
# directionsarray = [("e","se","w")]

class title:
    def __init__(self, x, y):
        self.x = x
        self.y = y
        self.color = "white"
    
    def create(self, x, y):
        return self.__class__(x, y)
    
    def toggle(self):
        if self.color == "white":
            self.color = "black"
        else:
            self.color = "white"
    
    def newpos(self, direc):
        actpos = [self.x, self.y]
        if direc == "e":
            actpos[0] = actpos[0] + 1
        elif direc == "w":
            actpos[0] = actpos[0] - 1
        elif direc == "se":
            if actpos[1] % 2 == 0:
                actpos[1] = actpos[1] + 1
            else:
                actpos[0] = actpos[0] + 1
                actpos[1] = actpos[1] + 1
        elif direc == "sw":
            if actpos[1] % 2 == 0:
                actpos[0] = actpos[0] - 1
                actpos[1] = actpos[1] + 1
            else:
                actpos[1] = actpos[1] + 1
        elif direc == "ne":
            if actpos[1] % 2 == 0:
                actpos[1] = actpos[1] - 1
            else:
                actpos[0] = actpos[0] + 1
                actpos[1] = actpos[1] - 1
        elif direc == "nw":
            if actpos[1] % 2 == 0:
                actpos[0] = actpos[0] - 1
                actpos[1] = actpos[1] - 1
            else:
                actpos[1] = actpos[1] - 1
        return actpos

    def adjacent(self):
        neighbours = []
        for adj in ("e", "se", "sw", "w", "nw", "ne"):
            check = self.newpos(adj)
            if str(str(check[0]) + '_' +  str(check[1])) in placed.keys():
                neighbours.append(placed[str(str(check[0]) + '_' +  str(check[1]))].color)
            else:
                # newtitle = self.create(check[0], check[1])
                placed[str(str(check[0]) + '_' +  str(check[1]))] = self.create(check[0], check[1])
                neighbours.append(placed[str(str(check[0]) + '_' +  str(check[1]))].color)
        return neighbours


WIDTH, HEIGHT = 800, 800
CENTER = [400,400]
color_black = (0, 0, 0)
color_white = (255, 255, 255)
color_grey = (105,105,105)
color_red = (255, 0,0)

pygame.init()
screen = pygame.display.set_mode((WIDTH, HEIGHT), 0, 32)

def draw_hex(surface, color, position):
    r = 6
    a, b = position

    if b % 2 == 0:
        x = a * r + CENTER[0] + (a * r)
    else:
        x = (a * r + CENTER[0]) + (a * r) + r
    y = int(b * (r * 1.5) + CENTER[1])

    pygame.draw.polygon(surface, color, [
        (x + r * cos(pi / 180 * (60 * i - 30)), y + r * sin(pi / 180 * (60 * i - 30)))
        for i in range(6)
    ])

def draw_dot(surface, color, position):
    r = 10
    a, b = position

    if b % 2 == 0:
        x = a * r + CENTER[0] + (a * r)
    else:
        x = (a * r + CENTER[0]) + (a * r) + r
    y = int(b * (r * 1.5) + CENTER[1])

    pygame.draw.circle(surface, color, (x,y), 3)


screen.fill(color_grey)
run = True

placed = {}
# for i in range(-18,18,1):
#     for j in range(-18,19,1):
#         actpos = [i,j]
#         placed[str(str(actpos[0]) + str(actpos[1]))] = title(actpos[0], actpos[1])
#         draw_hex(screen, color_white, actpos)
#         pygame.display.update()

while run:
    # pygame.time.delay(1000)
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            run = False

    actpos = [0,0]
    placed[str(str(actpos[0]) + '_' + str(actpos[1]))] = title(actpos[0], actpos[1])
    draw_hex(screen, color_white, actpos)
    pygame.display.update()

    for directions in directionsarray:
        actpos = [0,0]
        # placed[str(str(actpos[0]) + str(actpos[1]))] = title(actpos[0], actpos[1])
        # draw_hex(screen, color_white, actpos)
        for i, direc in enumerate(directions):
            # print(direc)
            # print(actpos)
            actpos = placed[str(str(actpos[0]) + '_' +  str(actpos[1]))].newpos(direc)
            # print(actpos)
            if i == len(directions) - 1:
                if str(str(actpos[0]) + '_' +  str(actpos[1])) in placed.keys():
                    placed[str(str(actpos[0]) + '_' +  str(actpos[1]))].toggle()
                else:
                    placed[str(str(actpos[0]) + '_' +  str(actpos[1]))] = title(actpos[0], actpos[1])
                    placed[str(str(actpos[0]) + '_' +  str(actpos[1]))].toggle()
            else:
                if not str(str(actpos[0]) + '_' +  str(actpos[1])) in placed.keys():
                    placed[str(str(actpos[0]) + '_' +  str(actpos[1]))] = title(actpos[0], actpos[1])
            draw_hex(screen, placed[str(str(actpos[0]) + '_' +  str(actpos[1]))].color, actpos)
            pygame.display.update()
        pygame.display.update()
        run = False

print(len([i for i in placed.keys() if placed[i].color == "black"]))
print("part 2")


for day in range(1,101):
    xxx = list(dict.fromkeys(placed))
    for i in range(min([int(i.split("_")[0])-1 for i in xxx]),max([int(i.split("_")[0]) for i in xxx])+1):
        for j in range(min([int(i.split("_")[1])-1 for i in xxx]),max([int(i.split("_")[1]) for i in xxx])+1):
            actpos = [i,j]
            if not str(str(actpos[0]) + '_' +  str(actpos[1])) in placed.keys():
                placed[str(str(actpos[0]) + '_' +  str(actpos[1]))] = title(actpos[0], actpos[1])
                draw_hex(screen, color_white, actpos)
            # pygame.display.update()

    to_toggle = []
    checkdict = dict.fromkeys(placed)
    for chk_title in checkdict:
        titlecolor = placed[chk_title].color
        neighbours = placed[chk_title].adjacent()
        if titlecolor == "black" and (neighbours.count('black') == 0 or neighbours.count('black') > 2):
            to_toggle.append(chk_title)
            # draw_dot(screen, color_red, (int(chk_title.split("_")[0]), int(chk_title.split("_")[1])))
        if titlecolor == "white" and neighbours.count('black') == 2:
            to_toggle.append(chk_title)
            # draw_dot(screen, color_red, (int(chk_title.split("_")[0]), int(chk_title.split("_")[1])))
    pygame.display.update()

    for chk_title in to_toggle:
        placed[chk_title].toggle()
        draw_hex(screen, placed[chk_title].color, (int(chk_title.split("_")[0]), int(chk_title.split("_")[1])))

    pygame.display.update()
    print("{}:{}".format(day, len([i for i in placed.keys() if placed[i].color == "black"])))

pygame.quit()
# exit()
