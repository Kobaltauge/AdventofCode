from collections import defaultdict 
import copy

with open("day_12_test.txt") as f:
    paths = [x.split('-') for x in f.read().splitlines()]

mapdict = {}
for path in paths:
    if path[0] in mapdict:
        mapdict[path[0]].append(path[1])
    else:
        mapdict[path[0]] = [path[1]]

#This class shows a directed graph 
class Graph: 
    def __init__(self): 
        self.graph = defaultdict(list)
        self.allpaths = []

    #this function adds an edge to the graph 
    def addEdge(self,u,v): 
        self.graph[u].append(v) 
   
    def printAllPathsUtil(self, u, d, visited, path): 
        # checking only lowercase as visited
        if u.islower():
            visited.append(u)
        path.append(u) 
  
        if u == d: 
            print(path)
            self.allpaths.append(copy.deepcopy(path))
        else: 
            for i in self.graph[u]: 
                if not i in visited: 
                    self.printAllPathsUtil(i, d, visited, path) 

        path.pop()
        if u.islower():
            visited.remove(u)

    # Printing all paths from sourse to destination
    def printAllPaths(self,s, d): 
        visited = [] 
        #Array for paths and all paths 
        path = []
        #Calling a recursive function for printing all paths 
        paths = self.printAllPathsUtil(s, d,visited, path)
  
#Creating the graph
g = Graph()
for k,v in mapdict.items():
    for v in mapdict[k]:
        if v.islower():
            g.addEdge(k, v)
        else:
            g.addEdge(k, v)
            g.addEdge(v, k) 
   
s = 'start' ; d = 'end'
g.printAllPaths(s, d)
print("ende")
