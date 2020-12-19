with open("day_19.txt") as f:
    raw = f.read().splitlines()

rules = raw[:134]
codes = raw[135:]

rulesdict = {}
for rule in rules:
    ruleno, ruletext = rule.split(":")[0], rule.split(":")[1].strip()
    rulesdict[str(ruleno)] = ruletext

def check(checkstring, idx):
    if checkstring[idx] == "a" or checkstring[idx] == "b":
        return True
    else:
        return check(checkstring, idx+1)

# def check(checkstring, idx, check1, check2):
#     if idx > len(checkstring):
#         return False
#     if rulesdict[str(check1)] == checkstring[idx]:
#         return True
#     elif check2:
#         if rulesdict[str(check2)] == checkstring[idx+1]:
#             return True
#         else:
#             return False
#     else:
#         return check(checkstring, idx+1, checkstring[idx], checkstring[idx+1])

counter = 0
for code in codes:
    if all(check(code, 0)):
        counter += 1

print("part 2")