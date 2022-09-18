from itertools import chain
from functools import reduce


class Cube:

    __slots__ = ['corner_permutation', 'edge_permutation', 'corner_orientation',  'edge_orientation']

    solved_cube = {
        'corner_permutation': [0, 1, 2, 3, 4, 5, 6, 7],
        'corner_orientation': [0, 0, 0, 0, 0, 0, 0, 0],
        'edge_orientation': [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        'edge_permutation': [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11]

    }

    def __init__(self, *args, **kwargs):
        for key, value in self.solved_cube.items():
            if key in kwargs:
                value = kwargs[key]
            setattr(self, key, value.copy())

    def __str__(self):
        result = chain(*[[str(getattr(self, item)), item, '\n'] for item in self.__slots__])
        return "".join(result)

    def __hash__(self):
        return reduce(lambda x, y: x ^ y, (hash(str(getattr(self, val))) for val in self.__slots__))

    def solved(self):
        return False not in [getattr(self, val) == self.default_vals[val] for val in self.__slots__]

    def move(self, move_string):
        return self.perform_move_old(nr=int(move_string[1]), move_table=self.move_tables[move_string[0].lower()])

    move_tables = {
                    'r': [[1, 3, 7, 5], [9, 6, 10, 2], [2, 1, 2, 1], [0, 0, 0, 0]],
                    'l': [[0, 4, 6, 2], [8, 0, 11, 4], [2, 1, 2, 1], [0, 0, 0, 0]],
                    'u': [[1, 5, 4, 0], [1, 2, 3, 0], [0, 0, 0, 0], [0, 0, 0, 0]],
                    'd': [[3, 2, 6, 7], [5, 4, 7, 6], [0, 0, 0, 0], [0, 0, 0, 0]],
                    'f': [[0, 2, 3, 1], [1, 8, 5, 9], [1, 2, 1, 2], [1, 1, 1, 1]],
                    'b': [[4, 5, 7, 6], [3, 10, 7, 11], [1, 2, 1, 2], [1, 1, 1, 1]]
                  }

    def perform_move_old(self, nr, move_table):
        move_cper, move_eper, move_cor, move_eor = move_table
        c, e, ct, et = move_cper, move_eper, move_cor, move_eor

        for i in range(0, nr):
            t = self.corner_permutation[c[0]]
            self.corner_permutation[c[0]] = self.corner_permutation[c[1]]
            self.corner_permutation[c[1]] = self.corner_permutation[c[2]]
            self.corner_permutation[c[2]] = self.corner_permutation[c[3]]
            self.corner_permutation[c[3]] = t

            t = self.edge_permutation[e[0]]
            self.edge_permutation[e[0]] = self.edge_permutation[e[1]]
            self.edge_permutation[e[1]] = self.edge_permutation[e[2]]
            self.edge_permutation[e[2]] = self.edge_permutation[e[3]]
            self.edge_permutation[e[3]] = t

            t = self.corner_orientation[c[0]]
            self.corner_orientation[c[0]] = ((self.corner_orientation[c[1]] + ct[0]) % 3)
            self.corner_orientation[c[1]] = ((self.corner_orientation[c[2]] + ct[1]) % 3)
            self.corner_orientation[c[2]] = ((self.corner_orientation[c[3]] + ct[2]) % 3)
            self.corner_orientation[c[3]] = ((t + ct[3]) % 3)

            t = self.edge_orientation[e[0]]
            self.edge_orientation[e[0]] = (self.edge_orientation[e[1]] + et[0]) % 2
            self.edge_orientation[e[1]] = (self.edge_orientation[e[2]] + et[1]) % 2
            self.edge_orientation[e[2]] = (self.edge_orientation[e[3]] + et[2]) % 2
            self.edge_orientation[e[3]] = (t + et[3]) % 2
        return self

    #
    # def perform_move(self, nr, move_table):
    #     attrs = zip([getattr(self, item) for item in self.__slots__], move_table, [20, 20, 3, 2])
    #     for i in range(0, nr):
    #         for T, v, Modulo in attrs:
    #             a = [v[0], v[1], v[2], v[3]] if Modulo < 5 else [0, 0, 0, 0]
    #             print(a)
    #             print(T)
    #             print(v)
    #             print((T[v[3]] + a[2]) % Modulo)
    #             print((T[v[0]] + a[3]) % Modulo)
    #             print()
    #             T[v[0]], T[v[1]], T[v[2]], T[v[3]] = ((T[v[1]] + a[0]) % Modulo,
    #                                                   (T[v[2]] + a[1]) % Modulo,
    #                                                   (T[v[3]] + a[2]) % Modulo,
    #                                                   (T[v[0]] + a[3]) % Modulo)
    #     return self
