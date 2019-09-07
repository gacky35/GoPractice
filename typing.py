prob = ["golang", "hands-on", "in", "kagawa"]

for term in prob:
    print(term)
    # print("input : ", end='')
    ans = input('input : ')
    if term == ans:
        print("OK")
    else:
        print("NG")
