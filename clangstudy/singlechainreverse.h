#ifndef __SINGLECHAINREVERSE_H__
#define __SINGLECHAINREVERSE_H__

#include <stdio.h>
#include <stdlib.h>
#include <string.h>


#define CHAIN_SIZE  5

typedef struct __singlechain__ {
    int value;
    struct __singlechain__ * next;
} singlechain;

int test_singlechainreverse(void);

#endif