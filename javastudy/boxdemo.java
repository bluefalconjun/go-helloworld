//add for boxdemo
/*

*/

class Box
{
    double width;
    double height;
    double depth;
} class boxdemo


{
    // Your program begins with a call to main().
  public
    static void main(String args[])
    {
	Box mybox = new Box();

	//note that mybox_reference is not a copy of Box but only the reference of the mybox
	Box mybox_reference = mybox;

	double vol;
	// assign values to mybox's instance variables
	mybox.width = 10;
	mybox.height = 20;
	mybox.depth = 15;
	// compute volume of box
	vol = mybox.width * mybox.height * mybox.depth;
	System.out.println("Volume is " + vol);
    }
}
