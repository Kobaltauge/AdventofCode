with open("day_19.txt") as f:
    raw = f.read().splitlines()

rules = raw[:134]
codes = raw[135:]

rulesdict = {}
for rule in rules:
    ruleno, ruletext = rule.split(":")[0], rule.split(":")[1].strip()
    rulesdict[str(ruleno)] = ruletext

def check(checkstring, idx, rule, erg):
    text = rulesdict[str(rule)]
    if checkstring[idx] == text.replace("\"",""):
        return True
    else:
        if len(text.split()) == 1:
            return erg.append(check(checkstring, idx+1, str(text).strip()))
        elif len(text) == 2:
            text1, text2 = checkstring[idx].split()
            return erg.append(all((check(checkstring, idx+1, str(text1).strip()) and checkstring(checkstring, idx+2), str(text2).strip())))
        else:
            text1, text2 = text.split("|")
            elem1, elem2 = text1.split()
            elem3, elem4 = text2.split()
            return erg.append(all((check(checkstring, idx+1, str(elem1).strip()) and checkstring(checkstring, idx+2), str(elem2).strip()), \
                    (check(checkstring, idx+1, str(elem3).strip()) and checkstring(checkstring, idx+2), str(elem4).strip())))

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
    if all(check(code, 0, 8, []), code(code, 0, 11, [])):
        counter += 1

print("part 2")