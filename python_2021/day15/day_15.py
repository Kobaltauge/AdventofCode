from queue import PriorityQueue

with open("day_15.txt") as f:
    map = f.read().splitlines()
    map = [list(i) for i in map]

openset = [""]
closedset = []

nodes = ('A', 'B', 'C', 'D', 'E', 'F', 'G')

unvisited = {node: None for node in nodes} #using None as +inf
visited = {}
current = 'B'
currentDistance = 0
unvisited[current] = currentDistance


print("kkk")