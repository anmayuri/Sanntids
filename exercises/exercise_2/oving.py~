# Python 3.3.3 and 2.7.6
# python oving1.py

from threading import Thread

i = 0

def add_func():
	global i
   	for x in range (0,(1000000-1)):
		i+=1

def rest_func():
	global i
	for x in range (0,(1000000-1)):
		i-=1



def main():
    add = Thread(target = add_func, args = (),)
    add.start()
    rest = Thread(target = rest_func, args = (),)
    rest.start()
    for x in range (0,100):
	print(i)
    add.join()
    rest.join()
    print("Done"+str(i))


main()
