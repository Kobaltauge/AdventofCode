import re

buses = "19,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,523,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,x,23,x,x,x,x,x,29,x,547,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,17"
time = 1000053

# buses = "7,13,x,x,59,x,31,19"
# time = 939

buses = buses.replace(r',x','').split(",")
buses = [int(e) for e in buses]

timetables = []
findingid = []
findingtime = []
for bus in buses:
    times = [e for e in range(bus, time+bus, bus)]
    #timetables.append(times)
    findingid.append(times[0])
    findingtime.append(times[-1])

closest_time = min(findingtime)
min_wait = closest_time - time
bus_id = findingid[findingtime.index(closest_time)]
print(min_wait * bus_id)

print("part 2")

buses = "19,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,523,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,x,23,x,x,x,x,x,29,x,547,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,17"
# buses = "7,13,x,x,59,x,31,19"
# buses = "17,x,13,19"
# buses = "67,7,59,61"
# buses = "67,x,7,59,61"
# buses = "67,7,x,59,61"
# buses = "1789,37,47,1889"

buses = buses.split(",")
buses = [int(e) if not e=="x" else str(e) for e in buses]
buses = [[e, buses[e]] for e in range(len(buses)) if (isinstance(buses[e], int))]

multi = buses[0][1]
multi = 99999999986
while True:
    # print(multi)
    # if ((multi % 7) == 0) and \
    #     ((multi+1) % 13 == 0) and \
    #     ((multi+4) % 59 == 0) and \
    #     ((multi+6) % 31 == 0) and \
    #     ((multi+7) % 19 == 0):
    #     print(multi)
    #     break

    if all([((multi + e[0]) % e[1] == 0) for e in buses]):
        print(multi)
        break
    multi += buses[0][1]