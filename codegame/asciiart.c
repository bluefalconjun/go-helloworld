#include <stdlib.h>
#include <stdio.h>
#include <string.h>

/*
 #  ##   ## ##  ### ###  ## # # ###  ## # # #   # # ###  #  ##   #  ##   ## ###
# # # # # # # # # # ### ###
# # # # #   # # #   #   #   # #  #    # # # #   ### # # # # # # # # # # #    #
# # # # # # # # # #   #   #
### ##  #   # # ##  ##  # # ###  #    # ##  #   ### # # # # ##  # # ##   #   #
# # # # ###  #   #   #   ##
# # # # #   # # #   #   # # # #  #  # # # # #   # # # # # # #    ## # #   #  #
# # # # ### # #  #  #
# # ##   ## ##  ### #    ## # # ###  #  # # ### # # # #  #  #     # # # ##   #
###  #  # # # #  #  ###  #



*/

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

/*
    65:A 90:Z 97:a 122:z 63:?
*/

int asciiart_main() {
  int width;
  scanf("%d", &width);
  fgetc(stdin);
  fprintf(stderr, "W:%d\n", width);

  int height;
  scanf("%d", &height);
  fgetc(stdin);
  fprintf(stderr, "H:%d\n", height);

  // what the hell is this for ??? to ignore the first blank ?
  fgetc(stdin);

  char textout[256];
  int ascidx[256];
  fgets(textout, 256, stdin);
  int textlenth = strlen(textout) - 1;
  fprintf(stderr, "Number of text_out:%d\n", textlenth);

  // get the string of ASCII Art
  char *asciiart = malloc(width * height * 27);
  int j = 0;
  for (j = 0; j < height; j++) {
    memset(asciiart, 0, width * 27);
    fgets(asciiart, width * 27, stdin);
    fprintf(stderr, "%s", asciiart);
  }

  int i = 0;
  for (i = 0; i < textlenth; i++) {
    fprintf(stderr, "c:%d %c\n", textout[i], textout[i]);
    if (97 <= textout[i] && textout[i] <= 122)
      textout[i] -= 32;
    if (65 > textout[i] || textout[i] > 90)
      textout[i] = 63;
    if (textout[i] == 63) {
      ascidx[i] = 27;
    } else
      ascidx[i] = textout[i] - 65;
  }

  fprintf(stderr, "%d", ascidx[255]);

  free(asciiart);

  return 0;
}
