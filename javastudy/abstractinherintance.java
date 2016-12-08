//abstract inheritance

//if one method / members / class is defined as final.
//then it can NOT BE extends. !!

//check about Object Class !!
/*

Object clone( ) 创建一个和被复制的对象完全一样的新对象
boolean equals(Object object) 判定对象是否相等
void finalize( ) 在一个不常用的对象被使用前调用
final Class getClass( ) 获取运行时一个对象的类
int hashCode( ) 返回调用对象有关的散列值
final void notify( ) 恢复一个等待调用对象线程的执行
final void notifyAll( ) 恢复所有等待调用对象线程的执行
String toString( )
final void wait( ) 等待另一个线程的执行
void wait(long milliseconds)
void wait(long milliseconds,
int nanoseconds)

*/

// Using abstract methods and classes.
abstract class Figure {
    double dim1;
    double dim2;

    Figure(double a, double b) {
        dim1 = a;
        dim2 = b;
    }
    // area is now an abstract method
    abstract double area();
}

class Rectangle extends Figure {
    Rectangle(double a, double b) {
        super(a, b);
    }
    // override area for rectangle
    double area() {
        System.out.println("Inside Area for Rectangle.");
        return dim1 * dim2;
    }
}

class Triangle extends Figure {
    Triangle(double a, double b) {
        super(a, b);
    }
    // override area for right triangle
    double area() {
        System.out.println("Inside Area for Triangle.");
        return dim1 * dim2 / 2;
    }
}

class AbstractAreas {
    public static void main(String args[]) {
        // Figure f = new Figure(10, 10); // illegal now
        Rectangle r = new Rectangle(9, 5);
        Triangle t = new Triangle(10, 8);
        Figure figref; // this is OK, no object is created
        figref = r;
        System.out.println("Area is " + figref.area());
        figref = t;
        System.out.println("Area is " + figref.area());
    }
}
