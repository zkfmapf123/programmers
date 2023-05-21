def solution(left, right):

    num = 0
    for i in range(left, right+1):
        gcd_len = list(gcd(i)).__len__()

        if gcd_len % 2 == 0:
            num += i
        else:
            num -= i

    return num


def gcd(num):
    for i in range(1, num+1):
        if num % i == 0:
            yield i


if __name__ == "__main__":
    print(solution(13, 17))
    print(solution(24, 27))
