//add for init

//basictypes in java
/*
(integer): byte[8] short[16] int[32] long[64]
(float): float double
(bool): boolean
(char): char[16]

array in java must be dynamicly new and delete.

*/

class study
{
    // Your program begins with a call to main().
    // static keywords makes your class can be used without using NEW*
  public
    static void main(String args[])
    {
	System.out.println("This is a simple Java program.");

	int num;   // this declares a variable called num
	num = 100; // this assigns num the value 100
	System.out.println("This is num: " + num);
	num = num * 2;
	System.out.print("The value of num * 2 is ");
	System.out.println(num);

	int lightspeed;
	long days;
	long seconds;
	long distance;
	// approximate speed of light in miles per second
	lightspeed = 186000;
	days = 1000;                     // specify number of days here
	seconds = days * 24 * 60 * 60;   // convert to seconds
	distance = lightspeed * seconds; // compute distance
	System.out.print("In " + days);
	System.out.print(" days light will travel about ");
	System.out.println(distance + " miles.");

	//2D array
	int twoD[][] = new int[4][5];
	int i, j, k = 0;
	for (i = 0; i < 4; i++)
	    for (j = 0; j < 5; j++) {
		twoD[i][j] = k;
		k++;
	    }
	for (i = 0; i < 4; i++) {
	    for (j = 0; j < 5; j++)
		System.out.print(twoD[i][j] + " ");
	    System.out.println();
	}

	//using comma
	int a, b;
	for (a = 1, b = 4; a < b; a++, b--) {
	    System.out.println("a = " + a);
	    System.out.println("b = " + b);
	}
    }
}
