from cube import Cube


import random
def scramble():
    a = Cube()
    k = ['r','l','u','d','f','b']
    kk = [1,2,3]
    for i in range(0,2):
        c = random.choice(k)
        cc = random.choice(kk)
        a.move((c,cc))
    return a

def search(cube,depth=0):
    if(cube.solved):
        print('solved',depth)
        print(cube)
        return cube, True

    if(depth > 3):
        return cube, False

    available_moves = ['r','l','u','d','f','b']
    available_nr = [1,2,3]
    for move in available_moves:
        for nr in available_nr:
            cube.move((move,nr))
            k = search(cube,depth= depth+1)
            if(k[1]):
                return cube, True

    return cube, False


o = scramble()
g=0
print(o)
search(o)
