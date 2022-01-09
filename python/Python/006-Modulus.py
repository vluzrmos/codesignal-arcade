def solution(n) :
    if (not callable(getattr(n, 'is_integer', None))) or n.is_integer():
        return n%2

    return -1

m = 23

print(solution(m))
