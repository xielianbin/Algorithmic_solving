class LinkedList():
    def __init__(self):
        self.items = []

    def insert(self, x,y):
        if x in self.items:
            idx=self.items.index(x)
            self.items.insert(idx, y)
        else:
            self.items.append(y)

    def delete(self,x):
        if x in self.items:
            idx=self.items.index(x)
            self.items.pop(idx)
    def size(self):
        return len(self.items)
    def traverse(self):
        for item in self.items:
            print(item,end=" ")
s=LinkedList()
n=int(input())
for i in range(n):
    meg=input().split(" ")
    if meg[0]=='insert':
        s.insert(meg[1],meg[2])
    elif meg[0]=='delete':
        s.delete(meg[1])
if s.size()==0:
    print("NULL")
else:
    s.traverse()
 