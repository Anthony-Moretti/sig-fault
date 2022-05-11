#include <stdio.h>
#include <signal.h>

int *p;

void Fault() {
    printf("Raising segmentation fault...\n");
    *p=1;
    printf("Completed segmentation fault!\n");
}