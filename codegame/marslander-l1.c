#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include <stdbool.h>

#ifndef MARLANDER_DEBUG
#define MARLANDER_DEBUG 1
#define MARLANDER_SIMULATION 1
#endif

// mar area
#define MAX_AREA_X 7000
#define MAX_AREA_Y 3000

#define MAX_MAX_LAND_HS 20
#define MAX_LAND_VS 40
#define LAND_ANGLE 0

#define LAND_WIDE 1000

// current ship position positive with landarea
typedef enum {
  kAREA_INIT = 0,
  kAREA_1 = 1, // straight above the landarea
  kAREA_2, // side abobe the landarea
  kAREA_3, // side above and need to climb up mountain
  kAREA_4 // side below the landarea
} ShipArea;

// current ship action
typedef struct {
  int angle;
  int power;
} ShipAction;

// surface point.
typedef struct {
  int x;
  int y; // or x1;
} LandSurfacePoint;

typedef struct {
  int x; // X coordinate
  int y; // Y coordinate
  int hs; // the horizontal speed (in m/s), can be negative.
  int vs; // the vertical speed (in m/s), can be negative.
  int f;  // the quantity of remaining fuel in liters.
  int r;  // the rotation angle in degrees (-90 to 90).
  int p;  // the thrust power (0 to 4).
} ShipStatus;

// g array
int g_gravity[7] = {4, 4, 4, 3, 4, 4, 3};
int g_gravityidx = 0;

// variable
LandSurfacePoint *gcur_surface = NULL; // all surface record.
LandSurfacePoint gdest_surface[2]; // dest surface.
LandSurfacePoint gland_point; // landing point at center of dest surface.
ShipArea gcur_area; // ship is in which area.
ShipStatus gcur_status; // current ship status
ShipStatus gnext_status; // current ship status
ShipAction gcur_action; // ship action.

int g_highY; // highest surface in land
int g_checkY; // check whether Y is too high

int simulation_land[6][2] = {{0, 100},    {1000, 500},  {1500, 100},
                             {3000, 100}, {5000, 1500}, {6999, 1000}};

// simulation function:
void simulation_scanf1(int *N) {
  *N = 6;
  return;
}

void simulation_scanf2(int *x, int *y) {
  static int i = 0;
  *x = simulation_land[i][0];
  *y = simulation_land[i][1];
  i++;
  // if (MARLANDER_DEBUG) fprintf(stderr, "simulation_scanf2: x[%d] y[%d]\n",
  // *x, *y);
  return;
}

void simulation_scanf3(ShipStatus *x) { return; }

void simulation_printf(int angle, int power) { return; }

// function
// init variables
void init_var(void) {
  memset(gdest_surface, 0, sizeof(LandSurfacePoint) * 2);
  memset(&gland_point, 0, sizeof(LandSurfacePoint));

  gcur_area = kAREA_INIT;

  memset(&gcur_status, 0, sizeof(ShipStatus));
  memset(&gnext_status, 0, sizeof(ShipStatus));
  memset(&gcur_action, 0, sizeof(ShipAction));

  g_highY = 0;
  g_checkY = 0;

  g_gravityidx = 6; // g_idx must start from 0

  return;
}

// get gravity change
int get_gVSchange(void) {
  g_gravityidx++;
  if (g_gravityidx == 7)
    g_gravityidx = 0;
  return g_gravity[g_gravityidx];
}

// expect next point
void expect_point(void) {
  int direction = 0;
  direction = gcur_status.r < 0 ? -1 : 1;
  if (gcur_status.r == 0)
    direction = 0;

  gnext_status.x = gcur_status.x + gcur_status.hs;
  gnext_status.y = gcur_status.y + gcur_status.vs - 2;
  gnext_status.hs = gcur_status.hs + gcur_status.p * direction;
  gnext_status.vs =
      gcur_status.vs + gcur_status.p * direction - get_gVSchange();
}

// reorg the surface from non-rect to rect
// can be done in loop
void reorg_surface(int N) { return; }

// check ship area
void check_area(void) {
  if (gcur_status.x >= gdest_surface[0].x &&
      gcur_status.x <= gdest_surface[1].x) {
    if (gcur_status.y >= gdest_surface[0].y)
      gcur_area = kAREA_1;
    else
      gcur_area = kAREA_2;
  } else
    gcur_area = kAREA_3;

  if (MARLANDER_DEBUG)
    fprintf(stderr, "Current Area:[%d]\n", gcur_area);

  return;
}

// refine H/V status
void refine_ship(void) {
  check_area();

  if (gcur_area == kAREA_1) {
    if (gcur_status.hs == 0 && gcur_status.r == 0)
      gcur_action.angle = 0;
    if (gcur_status.vs < 0)
      gcur_action.power = 4;
    else
      gcur_action.power = 3;
  }
  return;
}

// action
void atcion(void) {
  if (MARLANDER_SIMULATION)
    simulation_printf(gcur_action.angle, gcur_action.power);
  else
    printf("%d %d\n", gcur_action.angle, gcur_action.power);
  return;
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
int marlander_main() {
  init_var();

  int N; // the number of points used to draw the surface of Mars.

  if (MARLANDER_SIMULATION)
    simulation_scanf1(&N);
  else
    scanf("%d", &N);

  gcur_surface = malloc(sizeof(LandSurfacePoint) * N);
  memset(gcur_surface, 0, sizeof(LandSurfacePoint) * N);

  int i = 0;
  for (i = 0; i < N; i++) {
    if (MARLANDER_SIMULATION)
      simulation_scanf2(&(gcur_surface[i].x), &(gcur_surface[i].y));
    else
      scanf("%d%d", &(gcur_surface[i].x), &(gcur_surface[i].y));
    if (MARLANDER_DEBUG)
      fprintf(stderr, "LAND_X:[%d] LAND_Y:[%d]\n", gcur_surface[i].x,
              gcur_surface[i].y);

    // check HighY
    if (g_highY < gcur_surface[i].y)
      g_highY = gcur_surface[i].y;

    // check dest surface
    if (gcur_surface[i - 1].y == gcur_surface[i].y &&
        (gcur_surface[i].x - gcur_surface[i - 1].x >= LAND_WIDE)) {
      gdest_surface[0].x = gcur_surface[i - 1].x;
      gdest_surface[0].y = gcur_surface[i - 1].y;
      gdest_surface[1].x = gcur_surface[i].x;
      gdest_surface[1].y = gcur_surface[i].y;
      gland_point.x =
          (gdest_surface[1].x - gdest_surface[0].x) / 2 + gdest_surface[0].x;
      gland_point.y = gdest_surface[1].y;
      if (MARLANDER_DEBUG) {
        fprintf(stderr, "dest area:[%d][%d]\n", gdest_surface[0].x,
                gdest_surface[1].x);
        fprintf(stderr, "land point:[%d][%d]\n", gland_point.x, gland_point.y);
      }
    }
  }

  // game loop
  int game_loop = 0;
  while (game_loop++) {
    // check current -1
    if (MARLANDER_SIMULATION)
      simulation_scanf3(&gcur_status);
    else
      scanf("%d%d%d%d%d%d%d", &gcur_status.x, &gcur_status.y, &gcur_status.hs,
            &gcur_status.vs, &gcur_status.f, &gcur_status.r, &gcur_status.p);

    if (MARLANDER_DEBUG)
      fprintf(stderr, "Current X:[%d] Y:[%d] HS:[%d] VS:[%d]\n", gcur_status.x,
              gcur_status.y, gcur_status.hs, gcur_status.vs);

    // expect next -2
    expect_point();
    refine_ship();
    atcion();
    if (game_loop >= 200)
      break;
  }

  free(gcur_surface);

  return 0;
}
