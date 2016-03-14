#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <limits.h>
#include <stdint.h>

#define HORSERACE_SIMULATION 1

int N = 0;
uint64_t *kCurValue;
uint64_t MinDiff = 0;

uint64_t SelectSortAndDiff(uint64_t *array, int array_size) {
  int i, j, k, temp;

  for (i = 0; i < array_size; i++) {
    k = i;
    for (j = i + 1; j < array_size; j++) {
      if (array[k] > array[j])
        k = j;
    }
    if (k != i) {
      temp = array[i];
      array[i] = array[k];
      array[k] = temp;
    }
  }

  MinDiff = *array;

  for (i = 0; i < array_size - 1; i++) {
    if (MinDiff > abs(*(array + i) - *(array + i + 1))) MinDiff = abs(*(array + i) - *(array + i + 1);
    if (MinDiff == 0)
      return MinDiff;
  }

  return MinDiff;
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
int main() {
  int i;
  scanf("%d", &N);
  fprintf(stderr, "N:%d\n", N);

  kCurValue = calloc(sizeof(uint64_t), N);

  for (i = 0; i < N; i++) {
    int Pi;
    scanf("%d", kCurValue + i);
    fprintf(stderr, "kCurValue:%d\n", *(kCurValue + i));
  }

  // Write an action using printf(). DON'T FORGET THE TRAILING \n
  // To debug: fprintf(stderr, "Debug messages...\n");
  MinDiff = SelectSortAndDiff(kCurValue, N);

  printf("%d\n", MinDiff);
  free(kCurValue);

  return 0;
}
