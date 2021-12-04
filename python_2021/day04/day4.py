with open("day_04.txt") as f:
    numbers = list(map(int, f.readline().split(',')))
    f.readline()
    input = f.read()
    bingocards = [[list(map(int, j.split())) for j in i.split('\n')] for i in input.split('\n\n')]

def check_card(card, drawnnumbers,idx):
    for i in range(0,5):
        row = card[i]
        if len(set(row)-set(drawnnumbers)) == 0:
            win_list.append(idx)
            return True
        column = [j[i] for j in card]
        if len(set(column)-set(drawnnumbers)) == 0:
            win_list.append(idx)
            return True
    return False

gotcha_first = 0
win_list = []
for drawn in range(5, len(numbers)):
    drawnnumbers = numbers[0:drawn]
    for idx, card in enumerate(bingocards):
        gotcha = 0
        if idx not in win_list:
            if (check_card(card,drawnnumbers,idx)):
                gotcha = idx
        if gotcha != 0:
            if gotcha_first == 0:
                gotcha_first = gotcha
                gotcha_first_card = bingocards[gotcha]
                gotcha_first_drawnumbers = drawnnumbers
                gotcha_first_lastno = numbers[0:drawn][-1]
            else:
                gotcha_last_card = bingocards[gotcha]
                gotcha_last_drawnumbers = drawnnumbers
                gotcha_last_lastno = drawnnumbers[-1]

print(f"result part 1: {sum((set([item for sublist in gotcha_first_card for item in sublist]) - set(gotcha_first_drawnumbers))) * gotcha_first_lastno}")
print(f"result part 2: {sum((set([item for sublist in gotcha_last_card for item in sublist]) - set(gotcha_last_drawnumbers))) * gotcha_last_lastno}")




