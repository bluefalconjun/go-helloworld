#include "singlechainreverse.h"


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

void print_chain(singlechain *in){
    singlechain *cur = in;
    printf("list is: \n");
    while(cur->next != NULL){
        printf(" %d ", cur->value);
        cur = cur->next;
    }
    printf(" %d ", cur->value); // last one is not in loop
    printf("\n");
}

int reverse_chain(singlechain *in, singlechain **out){
    if(in->next == NULL) *out = in;

    singlechain *prev, *cur, *next;
    int i = 0;
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
    //memcpy(out, cur, sizeof(singlechain)); //if use struct but not struct pointer in func call. should cpy.
}

int test_singlechainreverse(void){
    singlechain* in = NULL;
    singlechain* out;

    printf("test_singlechainreverse start: \n");

    in = calloc(CHAIN_SIZE,sizeof(singlechain));

    init_chain(in, CHAIN_SIZE);
    print_chain(in);
    reverse_chain(in, &out);
    print_chain(out);

    free(in);

    printf("test_singlechainreverse end. \n");

    return 0;
}


