n = int(input())
c = 1
while c <= n:
    print('*' * c)
    c += 1

m = int(input())
stars = '*'
while len(stars) <= m:
    print(stars)
    stars += '*'


i = 0
while i < 5:
    print('*')
    if i % 2 == 0:
        print('**')
    if i > 2:
        print('***')
    i = i + 1


#looks nice
