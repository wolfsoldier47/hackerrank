import re

x = int(input())
y = []
for i in range(x):
    a = input()
    y.append(a)
is_valid = False
for i in y:
    try:
        re.compile(i)
        is_valid = True
    except re.error:
        is_valid = False

    print(is_valid)
