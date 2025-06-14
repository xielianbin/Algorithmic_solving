class Stack():
    def __init__(self):
        self.items=[]
    def isEmpty(self):
        return self.items==[]
    def pop(self):
        if self.size()>0:
            return self.items.pop()
        else:
            return "error"
    def top(self):
        if self.size()>0:
            return self.items[-1]
        else:
            return "error"
    def push(self,item):
        self.items.append(item)
    def size(self):
        return len(self.items)
s=Stack()
num=int(input())
for i in range(num):
    a=input()
    a=a.split(' ')
    b=None
    if a[0]=="push":
        s.push(a[1])
    elif a[0]=="top":
        print(s.top())
    elif a[0]=="pop":
        print(s.pop())