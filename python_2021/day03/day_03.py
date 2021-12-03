from operator import add

with open("day_03.txt") as f:
    binnumbers = f.read().splitlines()

ones_list = [0] * len(binnumbers[0])
zero_list = [0] * len(binnumbers[0])

for binnr in binnumbers:
    ones_list = list(map(add, ones_list, list(map(int, (binnr)))))
    zero_list = list(map(add, zero_list, [1 if x==0 else 0 for x in list(map(int, (binnr)))]))

    gamma_list = [1 if ones_list[i] > zero_list[i] else 0 for i in range(0, 12)]
    epsilon_list = [0 if ones_list[i] > zero_list[i] else 1 for i in range(0, 12)]
    
result = int(''.join(map(str,gamma_list)), 2) * int(''.join(map(str,epsilon_list)), 2)
print(f"Part 1 result: {result}")

oxygen_list = binnumbers

for idx in range(0,12):
    oxy_elements = [i[idx] for i in oxygen_list]
    oxy_ones = oxy_elements.count('1')
    oxy_zero = oxy_elements.count('0')
    if oxy_ones >= oxy_zero:
        oxygen_list = [subl for subl in oxygen_list if subl[idx] == '1']
    else:
        oxygen_list = [subl for subl in oxygen_list if subl[idx] == '0']
    if len(oxygen_list) == 1:
        break

co2_list = binnumbers

for idx in range(0,12):
    co2_elements = [i[idx] for i in co2_list]
    co2_ones = co2_elements.count('1')
    co2_zero = co2_elements.count('0')
    if co2_ones >= co2_zero:
        co2_list = [subl for subl in co2_list if subl[idx] == '0']
    else:
        co2_list = [subl for subl in co2_list if subl[idx] == '1']
    if len(co2_list) == 1:
        break

result = int(''.join(map(str,oxygen_list)), 2) * int(''.join(map(str,co2_list)), 2)
print(f"Part 2 result: {result}")
