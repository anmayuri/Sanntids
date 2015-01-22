# Python 3.3.3 and 2.7.6
# python oving1.py

from threading import Thread
from threading import Lock


i = 0
lck=Lock()

def add_func():
	global i
	lck.acquire()
   	for x in range (0,(1000000-1)):
		i+=1
	lck.release()	

def rest_func():
	global i
	lck.acquire()
	for x in range (0,(1000000-1)):
		i-=1
	lck.release()	



def main():
    add = Thread(target = add_func, args = (),)
    add.start()
    rest = Thread(target = rest_func, args = (),)
    rest.start()
    for x in range (0,50):
	print(i)
    add.join()
    rest.join()
    print("Done"+str(i))


main()
