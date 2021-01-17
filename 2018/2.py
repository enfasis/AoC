from collections import Counter
data = [str(x) for x in open('input.txt').readlines()]
twice = 0
three = 0
for word in data:
    cnt = Counter()
    for letter in word:
        cnt[letter] += 1
    if 2 in list(cnt.values()):
        twice += 1
    if 3 in list(cnt.values()):
        three += 1

print(twice*three)

while data != []:
    word1 = data.pop()
    for word in data:
        diff = 0
        index= 0
        for i in range(0,len(word)):
            if diff > 1:
                break
            elif word[i] != word1[i]:
                diff += 1
                index = i
        if diff == 1:
            print(word[:index]+word[index+1:])            
            data = []
            break