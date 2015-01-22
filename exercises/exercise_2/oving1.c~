// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o oving1 oving1.c -lpthread

#include <pthread.h>
#include <stdio.h>

int i=0;

void* add_func(){
	int x;
	for(x=0;x<1000000;x++)
	{
        	i++;
	}
	return NULL;
}

void* rest_func(){
	int x;
	for( x=0;x<1000000;x++)
	{
        	i--;
	}
	return NULL;
}

int main(){
    pthread_t add;
    pthread_t rest;
    pthread_create(&add, NULL, add_func, NULL);
    pthread_create(&rest, NULL, rest_func, NULL);
    for(int x=0;x<50;x++){
	printf("%i\n",i);
    }
    pthread_join(add, NULL);
    pthread_join(rest, NULL);  
    printf("Done:%i\n",i);
    return 0;
    
}

