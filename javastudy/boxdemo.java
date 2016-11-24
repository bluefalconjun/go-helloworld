//add for boxdemo
/*

*/

class Box
{
    double width;
    double height;
    double depth;

    // constructor used when no dimensions specified
    Box() {
        this.width = -1; // use -1 to indicate
        this.height = -1; // an uninitialized
        this.depth = -1; // box
    }

    // constructor used when cube is created
    Box(double len) {
        this.width = this.height = this.depth = len;
    }

    // This is the constructor for Box.
    Box(double width，double height，double depth) {
    this.width = width;
    this.height = height;
    this.depth = depth;

    // display volume of a box
    void volume() {
        System.out.print("Volume is ");
        System.out.println(width * height * depth);
    }

    protected void finalize( ){
        // finalization code here
    }

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

    mybox_reference.volume();

    }
}
