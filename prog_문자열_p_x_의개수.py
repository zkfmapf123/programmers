def solution(s):
    return s.lower().count('p') == s.lower().count('y')


if __name__ == "__main__":
    print(solution("pPoooyY"))
    print(solution("Pyy"))
