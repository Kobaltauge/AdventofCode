with open("day_10.txt") as f:
    adaptorlist = f.read().splitlines()

# with open("day_10_exampl.txt") as f:
#     adaptorlist = f.read().splitlines()

adaptorlist = [int(i) for i in adaptorlist]
adaptorlist.sort()

counter = 0
counter_1 = 0
counter_3 = 0
outlet = 0
for index in range(len(adaptorlist)):
    if (adaptorlist[index] - outlet) == 1:
        counter_1 += 1
        counter += 1
        outlet = adaptorlist[index]
    elif (adaptorlist[index] - outlet) == 3:
        counter_3 += 1
        counter += 1
        outlet = adaptorlist[index]
if counter == len(adaptorlist):
    print("all used")
else:
    print("remain")

print(adaptorlist[-1])
print(counter)
print(counter_1, counter_3+1)
print(counter_1 * (counter_3+1))
print("part 2")


counter = 0

class Node(object):
    def __init__(self, name, left=None, right=None):
        self.name = name
        self.left = left
        self.right = right

adaptorgraph = {}

for index in range(len(adaptorlist)):
    adaptorgraph[str(adaptorlist[index])] = []
    if str(adaptorlist[index]-1) in adaptorgraph.keys():
        adaptorgraph[str(adaptorlist[index]-1)].append(adaptorlist[index])
    if str(adaptorlist[index]-3) in adaptorgraph.keys():
        adaptorgraph[str(adaptorlist[index]-3)].append(adaptorlist[index])


def get_routes(node, route=None):
    node = str(node)
    route = route or []
    # If this node has children, clone the route, so we can process
    # both the right and left child separately.
    try:
        if adaptorgraph[node][1]:
            right_route = route.copy()
    except:
        pass

    # Process the main (left) route.  Pass the route on to the left 
    # child, which may return multiple routes.
    # route.append(NP(node, False))
    routes = get_routes(adaptorgraph[node][0], route) if (len(adaptorgraph.get(node))>0) else [route]

    # If there is a right child, process that as well.  Add the route
    # results.  Note that NP.right is set to True, to indicate a right path.
    try:
        if adaptorgraph[node][1]:
            right_routes = get_routes(adaptorgraph[node][1], right_route)
            routes.extend(right_routes)
    except:
        pass
    
    # Pass the results back
    return routes


routes = get_routes("1")

print("end")


def next_adapt(root):
    if root is None:
        return 0
    
    print(actsum, adaptorlist[index])
    if (actsum + adaptorlist[index]) == (138 + 3):
        return (counter + 1)
    elif (actsum + adaptorlist[index]) > (138 + 3):
        return counter
    else:
        newcounter = next_adapt((actsum + adaptorlist[index]), counter, (index + 1))
        if newcounter == counter:
            return counter
        else:
            return newcounter

counter = next_adapt(0,0,0)
print(counter)
print("end")