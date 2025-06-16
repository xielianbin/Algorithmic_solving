
class Quene(object):
    def __init__(self):
        self.s1=list()
        self.s2=list()
    def push(self, x):
        self.s1.append(x)
    def pop(self):
        if  len(self.s2)>0:
            return self.s2.pop()
        if len(self.s1)==0:
            return "error"
        for i in range(len(self.s1)):
            self.s2.append(self.s1.pop())
        return self.s2.pop()
    def front(self):
        if len(self.s2)>0:
            return self.s2[-1]
        elif len(self.s1)>0:
            return self.s1[0]
        else:
            return "error"
    def check(self, x):
        if x=='pop':
            print(self.pop())
        elif x=='front':
            print(self.front())
        else:
            self.push(int(x.split()[-1]))
        


n=input()
tmp=Quene()
for _ in range(int(n)):
    tmp.check(input())