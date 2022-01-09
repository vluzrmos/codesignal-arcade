import string

password = "iamthebest"
key = "zabcdefghijklmnopqrstuvwxy"

table = password.maketrans(string.ascii_lowercase, key)

print(password.translate(table))


def solution(t, width, precision):
    return ("{:^"+str(width)+"."+str(precision)+"f}").format(t)