#include <stdio.h>
#include "test.h"

#ifdef __TEST__

void hello(){
	printf("Hello world!\n");
}

int main(int args, char *argv[]){
	hello();
	return 0;
}

#endif




