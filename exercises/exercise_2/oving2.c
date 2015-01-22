// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o oving2 oving2.c -lpthread

#include <pthread.h>
#include <stdio.h>

int i=0;

pthread_mutex_t mtx;

void* add_func(){
	int x;
	pthread_mutex_lock(&mtx);
	for(x=0;x<1000000;x++)
	{
        	i++;
	}
	pthread_mutex_unlock(&mtx);
	return NULL;
}

void* rest_func(){
	int x;
	pthread_mutex_lock(&mtx);
	for( x=0;x<1000000;x++)
	{
        	i--;
	}
	pthread_mutex_unlock(&mtx);
	return NULL;
}

int main(){
    pthread_mutex_init(&mtx, NULL);
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
    pthread_mutex_destroy(&mtx);
    
    return 0;
    
}

