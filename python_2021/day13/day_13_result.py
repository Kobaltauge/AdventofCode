with open("day_13_result2.txt") as f:
    newdots = [x.split() for x in f.read().splitlines()]




for x in range(0, max([x[0] for x in newdots])+1):
    if (x,y) in newdots:
        print("#")
    else:
        print(".")
print("\n")

