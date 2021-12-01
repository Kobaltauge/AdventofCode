from types import prepare_class


with open("day_01.txt") as f:
    numbers = list(map(int, (f.read().splitlines())))

prev = 0
count_more = 0
count_less = 0
count_equal = 0

for nr in numbers:
    if prev == 0:
        pass
    elif nr > prev:
        count_more += 1
    elif nr < prev:
        count_less += 1
    else:
        count_equal += 1
    prev = nr
print(f"more {count_more}\nless {count_less}\nequal {count_equal}\n")

prev = 0
index = 0
count_more = 0
count_less = 0
count_equal = 0

for idx in range(0, len(numbers)-2):
    sum_windows = numbers[idx] + numbers[idx+1] + numbers[idx+2]
    if prev == 0:
        pass
    elif sum_windows > prev:
        count_more += 1
    elif sum_windows < prev:
        count_less += 1
    else:
        count_equal += 1
    prev = sum_windows
print(f"more {count_more}\nless {count_less}\nequal {count_equal}\n")
