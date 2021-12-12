
with open("day_12.txt") as f:
    paths = [x.split('-') for x in f.read().splitlines()]

mapdict = {}
for path in paths:
    if path[0] in mapdict:
        mapdict[path[0]].append(path[1])
    else:
        mapdict[path[0]] = [path[1]]

def get_ways(mapdict, root):
    x = []
    if not mapdict[root] == 'end':
        for c in mapdict[root]:
            for el in get_ways(mapdict, c):
                x.append(c[1] + el)
    else:
        x.append(mapdict[root])
    return x

# def paths(root):
#     x = []
#     if root.children:
#         for c in root.children:
#             for el in paths(c):
#                 x.append(c.data + el)
#     else:
#         x.append("")
#     return x

ways = []

for i in mapdict['start']:
    ways.append(get_ways(mapdict, 'start'))
    print("x")
