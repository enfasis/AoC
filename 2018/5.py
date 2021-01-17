def reaction(data):
    i=0
    end=len(data)
    while i != end-1:
        react = False
        if data[i].isupper():
            if data[i+1].islower():
                react = data[i] == data[i+1].upper()
        else:
            if data[i+1].isupper():
                react = data[i].upper() == data[i+1] 
        if react:
            data = data[:i]+data[i+2:]
            end -= 2
            if i != 0:
                i -= 1
        else:
            i += 1
    return end


data = open('input.txt').readlines()[0]
print(reaction(data))

shortest = len(data)
for i in range(65,90):
     polymer = data.replace(str(chr(i)),"").replace(str(chr(i+32)),"")
     polymerLength = reaction(polymer)
     if shortest > polymerLength:
         shortest = polymerLength

print(shortest)