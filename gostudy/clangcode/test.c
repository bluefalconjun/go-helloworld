#include <stdio.h>
#include <stdlib.h>


#define CHAIN_SIZE  5

typedef struct __singlechain__ {
    int value;
    struct __singlechain__ * next;
} singlechain;


void init_chain(singlechain *in, int size){
    singlechain *cur = in;
    for (int i=0; i<size; i++){
        cur->value = i;
        if(i!=size-1) {
            cur->next = cur + 1; //type* + N = type* + N*sizeof(type)
            cur = cur->next;
        }else {
            cur->next = NULL;
        }
    }
}

void print_chain(singlechain *in, int size){
    singlechain *cur = in;
    printf("list is: \n");
    for (int i=0; i<size; i++){
        if (cur != NULL ) printf(" %d ", cur->value);
        if (cur->next != NULL ) cur = cur->next;
    }
    printf("\n");
}

void reverse_chain(singlechain *in, singlechain **out){
    if(in->next == NULL) *out = in;

    singlechain *prev, *cur, *next;
    prev = cur = next = NULL;

    cur = in;
    while(cur->next !=NULL){
        next = cur->next;
        cur->next = prev;
        prev = cur;
        cur=next;
    }
    cur->next = prev;   //last one is not in loop.
    *out = cur;
}

int main(int argc, char** argv){
    singlechain* in = NULL;
    singlechain* out = NULL;

    in = malloc(CHAIN_SIZE*sizeof(singlechain));
    printf("%p\n", in);
    init_chain(in, CHAIN_SIZE);
    print_chain(in, CHAIN_SIZE);
    reverse_chain(in, &out);
    print_chain(out, CHAIN_SIZE);

    free(in);

    return 0;
}


