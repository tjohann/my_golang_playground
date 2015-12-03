/*
 * same behaviour like os_args.go
 *
 * BSD 3-Clause
 * (c) 2015, thorsten.johannvorderbrueggen@t-online.de  
 */

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <time.h>

#define MAX_LEN 255

int main(int argc, char *argv[])
{
	int i = 0;
	char input[MAX_LEN];
	struct timespec start, stop;
	double result;

	memset(input, 0, MAX_LEN);

	clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &start);
	for (i = 1; i < argc ; i++) {
		printf("argv[%d]: %s\n", i, argv[i]);
		strcat(input, argv[i]);
		strcat(input, " ");
	}
	clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &stop);

        // result in usec
	result = (stop.tv_sec - start.tv_sec) * 1e6
		+
		(stop.tv_nsec - start.tv_nsec) / 1e3; 

	printf("Result in micro seconds: %d\n", (int) result);
	// Result -> in micro seconds: 189
	
	exit(0);
}
