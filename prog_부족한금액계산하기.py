def solution(price, money, count):
    answer = 0

    for i in range(1, count+1):
        answer += price * i

    if money - answer < 0:
        return abs(money - answer)

    return 0


if __name__ == "__main__":
    print(solution(3, 20, 4))
