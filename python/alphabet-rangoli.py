def check_value(a,b):
    if b == a*3:
        return True
    elif b != a*3:
        return False

def print_carpet(a,b):
    half = False
    count = 1
    for i in range(1,a+1):
        if i == (int((a/2)+1)):
            half = True
            print("-"*((b - 7)//2),
                    "WELCOME",
                    "-"*((b - 7)//2),sep=''
                    )
            count -= 2

        else:
            if half:
                print("-"*(int(b/2 + 2)-(3*(a-i+1)**1)),\
                        ".|."*count,\
                        "-"*((int(b/2 + 2)-(3*(a-i+1)**1))),sep='')
                count -= 2
            else:
                print("-"*(int(b/2 + 2)-(3*i)),\
                        ".|."*count,\
                        "-"*((int(b/2 + 2)-(3*i))),sep='')
                count += 2



if __name__ == '__main__':
    a,b = input().split(" ")
    a = int(a)
    b = int(b)
    check = check_value(a,b)
    if check:
        print_carpet(a,b)
    else:
        print("second value should be equal to first multiply by 3")
        exit()
