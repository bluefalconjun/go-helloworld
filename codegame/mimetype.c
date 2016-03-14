#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <time.h>

#define MIMETYPE_SIMULATION 1

typedef struct { char content[501]; } OneContent;

clock_t cur_time = 0;
clock_t next_time = 0;

OneContent *kCurExt;
OneContent *kCurMIME;
OneContent *kCurFileName;

#if MIMETYPE_SIMULATION
int N = 4;
int Q = 7;
/*
kCurExt[0]:png kCurMIME[0]:image/png
kCurExt[1]:TIFF kCurMIME[1]:image/TIFF
kCurExt[2]:css kCurMIME[2]:text/css
kCurExt[3]:TXT kCurMIME[3]:text/plain
*********************************
kCurFileName[0]:example.TXT
kCurFileName[1]:referecnce.txt
kCurFileName[2]:strangename.tiff
kCurFileName[3]:resolv.CSS
kCurFileName[4]:matrix.TiFF
kCurFileName[5]:lanDsCape.Png
kCurFileName[6]:extract.cSs
*/

OneContent kTestExt[4] = {
    "png", "TIFF", "css", "TXT",
};

OneContent kTestMIME[4] = {
    "image/png", "image/TIFF", "text/css", "text/plain",
};
OneContent kTestFileName[7] = {
    "example.TXT", "referecnce.txt", "strangename.tiff", "resolv.CSS",
    "matrix.TiFF", "lanDsCape.Png",  "extract.cSs"};

#else

int N;
int Q;

#endif

void process_onefilename(char *filename) {
  int filenameLength = 0;
  filenameLength = strlen(filename) - 1;

  int i = filenameLength;
  int filematch = 0;
  int dotmatch = 0;
  char *dotpositioninfile = NULL;
  int j = 0;

  // optimize
  dotpositioninfile = strrchr(filename, '.');

  if (NULL == dotpositioninfile) {
    printf("UNKNOWN\n");
    return;
  } else {
    for (j = 0; j < N; j++) {
      if (strncmp(dotpositioninfile + 1, (kCurExt + j)->content,
                  strlen((kCurExt + j)->content) - 1) == 0) {
        filematch = 1;
        // printf("%s\n", (kCurMIME + j)->content);
      }
    }
    if (filematch == 0) {
      printf("UNKNOWN\n");
    }
  }

  return;
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

#if MIMETYPE_SIMULATION
int mimetype_main() {
#else
int main() {
#endif

  int i = 0;
  int j = 0;

#if MIMETYPE_SIMULATION

#else
  scanf("%d", &N);
  scanf("%d", &Q);
#endif

  fprintf(stderr, "N:[%d] Q:[%d]\n", N, Q);

  kCurExt = calloc(N, sizeof(OneContent));
  kCurMIME = calloc(N, sizeof(OneContent));
  kCurFileName = calloc(Q, sizeof(OneContent));

  for (i = 0; i < N; i++) {
    cur_time = clock();
#if MIMETYPE_SIMULATION
    memcpy((kCurExt + i)->content, (kTestExt + i)->content,
           strlen((kTestExt + i)->content));
    memcpy((kCurMIME + i)->content, (kTestMIME + i)->content,
           strlen((kTestMIME + i)->content));
#else
    scanf("%s%s", (kCurExt + i)->content, (kCurMIME + i)->content);
#endif
    // chage all ext into lower case
    for (j = 0; j < strlen((kCurExt + i)->content); j++) {
      // A=dex:65 Z=dex:90 a=dex:97 z=dex:122
      if ((kCurExt + i)->content[j] < 91) {
        (kCurExt + i)->content[j] = (kCurExt + i)->content[j] + 32;
      }
    }
    next_time = clock();
// fprintf(stderr, "ext cost :%ld\n", next_time - cur_time);

#if MIMETYPE_SIMULATION
#else
    fgetc(stdin);
#endif
  }

  for (i = 0; i < Q; i++) {
    cur_time = clock();
#if MIMETYPE_SIMULATION
    memcpy((kCurFileName + i)->content, (kTestFileName + i)->content,
           strlen((kTestFileName + i)->content));
#else
    fgets((kCurFileName + i)->content, 501, stdin);  // One file name per line.
#endif
    // change all filename into lower case
    for (j = 0; j < strlen((kCurFileName + i)->content); j++) {
      // A=dex:65 Z=dex:90 a=dex:97 z=dex:122
      if ((kCurFileName + i)->content[j] < 91 &&
          (kCurFileName + i)->content[j] > 64) {
        (kCurFileName + i)->content[j] = (kCurFileName + i)->content[j] + 32;
      }
    }
    next_time = clock();
    fprintf(stderr, "name change :%ld\n", next_time - cur_time);
    cur_time = clock();
    // process filename
    process_onefilename((kCurFileName + i)->content);
    next_time = clock();
    fprintf(stderr, "process :%ld\n", next_time - cur_time);
  }

  free(kCurExt);
  free(kCurMIME);
  free(kCurFileName);

  return 0;
}
