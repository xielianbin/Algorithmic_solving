class CirQue():
    def __init__(self):
        self.items=[]

    def push(self, x):
        return self.items.append(x)

    def front(self):
        return self.items[0]

    def pop(self):
        return self.items.pop(0)
    def size(self):
        return len(self.items)
meg=input()
megsplit=meg.split(" ")
lgt=int(megsplit[0])
opt=int(megsplit[1])
sq=CirQue()
for i in range(opt):
    opo= input()
    opsplit=opo.split(" ")
    if opsplit[0]=="push":
        if sq.size() >= lgt:
            print("full")
        else:
            sq.push(int(opsplit[1]))
    elif opsplit[0]=="pop":
        if sq.size()==0:
            print("empty")
        else:
            print(sq.pop())
    elif opsplit[0]=="front":
        if sq.size()==0:
            print("empty")
        else:
            print(sq.front())