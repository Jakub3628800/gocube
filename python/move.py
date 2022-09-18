from cube import Cube


moves = ['r', 'l', 'u', 'd', 'f', 'b']

cube = Cube()


cache_dictionary = {}

max_moves = 3

def perform_move(cube, nr):
    if(nr == max_moves):
        return

    _hash = cube.__hash__()
    if(_hash not in cache_dictionary) or (cache_dictionary[_hash] > nr):
        cache_dictionary[_hash] = nr

    for move in moves:
        for i in range(1, 4):
            mv = move + str(i)
            cube.move(mv)
            perform_move(cube, nr+1)

perform_move(cube,0)
print(len(cache_dictionary))


a = Cube()
a.move('r1')
a.move('u1')
a.move('f2')
a.move('r3')
max_depth = 5

def search(cube, nr, moves):
    if(nr > max_depth):
        return None
    if(cube.solved()):
        print(moves, ' solvedmoves')
        print(cube)
        return True

    for move in moves:
        for i in range(1, 4):
            mv = move + str(i)
            cube.move(mv)
            search(cube, nr+1, moves + mv)

search(a,0,'')
