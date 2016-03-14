#include <stdlib.h>
#include <stdio.h>
#include <string.h>

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

// for acsii code. d:48 = '0' d:49 = '1'

void proc_chucknorris(char *MsgBinary, int MsgBinary_Length,
                      char *MsgChuckNorris, int *MsgChuckNorris_Lenth) {
  int MsgBinaryIdx = 0;
  int MsgChuckNorrisProceed = 0;
  *MsgChuckNorris_Lenth = 0;

  char CurMsgBin = 48;
  char NextMsgBin = 50;
  int CurNextEqualCount = 1;

  while (MsgBinaryIdx < MsgBinary_Length) {
    CurMsgBin = *(MsgBinary + MsgBinaryIdx);
    NextMsgBin = *(MsgBinary + MsgBinaryIdx + 1);
    fprintf(stderr, "CurMsgBin/NextMsgBin:dec:%d %d\n", CurMsgBin, NextMsgBin);

    // read till different
    while (CurMsgBin == NextMsgBin) {
      MsgBinaryIdx++;
      CurNextEqualCount++;
      NextMsgBin = *(MsgBinary + MsgBinaryIdx + 1);
    }
    if (CurMsgBin == 48) {
      *(MsgChuckNorris + MsgChuckNorrisProceed) = 48;
      *(MsgChuckNorris + MsgChuckNorrisProceed + 1) = 48;
      // blank here
      MsgChuckNorrisProceed += 3;
    } else if (CurMsgBin == 49) {
      *(MsgChuckNorris + MsgChuckNorrisProceed) = 48;
      // blank here
      MsgChuckNorrisProceed += 2;
    }
    // fill the count of 1 or 0
    int i = 1;
    for (i = 1; i <= CurNextEqualCount; i++) {
      *(MsgChuckNorris + MsgChuckNorrisProceed) = 48;
      MsgChuckNorrisProceed++;
    }
    // for last blank
    CurNextEqualCount = 1;
    MsgChuckNorrisProceed++;
    MsgBinaryIdx++;
  }

  fprintf(stderr, "MsgChuckNorrisProceed:%d\n", MsgChuckNorrisProceed);
  *MsgChuckNorris_Lenth = MsgChuckNorrisProceed;

  return;
}

int chucknorris_main() {
  // char MsgRead[100] = {0};
  // fgets(MsgRead,100,stdin); //simulation from local
  // MsgRead[0] = 'C';
  // MsgRead[1] = '\0';
  char *MsgRead = "CC";
  // int MsgRead_Length = strlen(MsgRead) - 1;
  int MsgRead_Length = strlen(MsgRead);
  int MsgBinary_Length = 7 * MsgRead_Length;

  // Write an action using printf(). DON'T FORGET THE TRAILING \n
  // To debug: fprintf(stderr, "Debug messages...\n");
  fprintf(stderr, "MsgRead_Length:%d \n", MsgRead_Length);
  fprintf(stderr, "MsgRead:hex:%s\n", MsgRead);

  // change MsgRead from Charactor to binary format.
  char *MsgBinary = calloc(1, 7 * MsgRead_Length);

  // store ChuckNorris format msg.
  // the worst case of ChuckNorris is the MsgBinary is :
  //"1010.." -> "0 0 00 0 0 0 00 0.."
  // so max length of MsgChuckNorris is (5+4) * (MsgRead_Length/2+1)
  char *MsgChuckNorris = malloc(9 * (7 * MsgRead_Length / 2 + 1));
  memset(MsgChuckNorris, 0x20, 9 * (7 * MsgRead_Length / 2 + 1));

  int i = 0;
  int j = 0;
  for (j = 0; j < MsgRead_Length; j++) {
    for (i = 0; i < 7; i++) {
      *(MsgBinary + j * 7 + i) = MsgRead[j] >> (6 - i) & 1 ? 49 : 48;
      fprintf(stderr, "MsgBinary %d is:%d\n", j * 7 + i,
              *(MsgBinary + j * 7 + i));
    }
  }
  //*(MsgBinary + 7 * MsgRead_Length) = '\0';

  // fprintf(stderr, "MsgBinary is:%d\n", MsgBinary);

  int proceed_count = 0;

  proc_chucknorris(MsgBinary, MsgBinary_Length, MsgChuckNorris, &proceed_count);

  *(MsgChuckNorris + proceed_count) = '\0';

  printf("MsgChuckNorris proceed_count = %d content: %s\n", proceed_count,
         MsgChuckNorris);

  free(MsgBinary);
  free(MsgChuckNorris);

  return 0;
}
