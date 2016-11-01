-----
> **Effective Go**

-----
翻译并记录golang网站中的**[Effective Go](https://golang.org/doc/effective_go.html)**文档.

-----
> **Introduction**

Go是一门从各种已存在的语言中借鉴很多特点而来的新语言. 它拥有很多不常见的属性, 使得高效率的go语言同其他语言相比有不同之处. 直接将Java或者C++程序翻译成为Go语言不见得是有效率的, - 因为Java程序是由Java语言编写的,不是Go. 也就是说, 从Go的思路来考虑解决具体编程问题将会产生一个截然不同的成功的程序. 如果需要编写好的Go语言代码, 那么需要从它的特有属性和语法中来理解. 例如: 命名, 格式, 程序构造等. 这样同时你所编写的程序也更容易被其他Go程序员所理解(**`所有的程序写出来 除了能工作 最重要的是给人读的!!`**).

该文档描述如何编写清晰, 符合惯例的Go 代码. 在阅读该文档之前, 需要先查看 **[language specification](https://golang.org/ref/spec)**, **[the Tour of Go](https://tour.golang.org/)**, 和 **[How to Write Go Code](https://golang.org/doc/code.html)**.

**Examples**

**[Go 开发包源代码](https://golang.org/src/)**不仅仅作为Go语言的核心库, 同时也作为如何使用该语言的示例. 许多以上开发包中包含有可执行的自包含的例子, 你可以直接从**[golan.org](https://golang.org/)**网站上运行来查看. 例如**[这个](https://golang.org/pkg/strings/#example_Map)**. 如何对于如何解决具体问题, 或者是某些功能是如何实现的, 关于系统核心库的代码和例子将提供答案和参考.


-----
> **Formatting**

格式问题是最有争议的, 同时也是最不重要的. 人们可以适应不同的格式化风格, 但如果他们不必这样做将会更好, 如果每个人都遵循相同的风格, 更少的时间用于选择主题.  问题是如何在不必使用一个长的规范风格指南时来处理这个乌托邦.

在Go语言中, 我们采取一种不常见的方式来处理这个问题. - 让机器来考虑所有的格式问题. 使用**`gofmt`** 程序(或者是 go fmt) 来读取Go源代码, 并将其按照变标准格式(空格, 垂直对齐, 换行或者是注释)进行输出. 如果存在某些新的代码格式是不支持或者变换后存在错误的, 请重新组织代码(或者向gofmt提交bug), 不要略过该步骤.

作为例子, 不要浪费时间对结构体中的成员注释进行对齐工作, Gofmt将处理它们:

    type T struct {
        name string // name of the object
        value int // its value
    }

gofmt将会自动进行对齐操作:

    type T struct {
        name    string // name of the object
        value   int    // its value
    }

标准核心库中所有的Go代码均由gofmt进行格式化.

简要的格式化细节:

 - 空格
使用tabs来进行缩进, gofmt自动进行替换, 仅仅需要时使用空格(**`大吐槽!!! 我要空格 我要空格!!`**)
 - 行宽
Go没有行宽的限制, 因此不要在意是否会过长, 可以通过一个额外的tab来进行换行.
 - 括号
Go相对于C和Java来讲, 需要更少的括号: 控制结构(if, fo
, switch)在对应语法中并不需要括号. 同时, 操作符的优先级将更清楚.

    x<<8 + y<<16

意味着先移位再加, 就像它的空格连接所暗示的一样. 这和其他语言有所不同.


-----
> **Commentary**

Go支持C类型的 **`/**/`** 和C++类型的 **`//`** 注释. 

godoc处理程序将从Go源代码文件中解析出该代码的相关文档. (**`相关格式参见具体代码`**)

Go的声明语法支持集合方式的声明. 该类声明的方式可以指定同一注释下的变量, 或者是相关性的变量:

    // Error codes returned by failures to parse an expression.
    var (
        ErrInternal      = errors.New("regexp: internal error")
        ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
        ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
        ...
    )

-

    var (
        countLock   sync.Mutex
        inputCount  uint32
        outputCount uint32
        errorCount  uint32
    )


-----
> **Names**

命名在Go语言中存在有重要的意义, 它还可以起到语义区分的作用: 在一个包中的元素是否被包外可见, 将有它名字的第一个字符是否大写来决定. (**`包括 函数, 变量, 类型定义?等`**)


**Package names**

当导入一个包时, 该包的包名将作为其内容的访问标识符. 

    import "bytes"

在上面的语句完成后, 导入 **`bytes`** 包的代码可以访问 **`bytes.Buffer`** . 当使用某个包的程序能够使用同样的名字来访问包对象时, 这对编程十分有帮助. 但同时这也对包名有一定的要求: 简短, 明确. 惯例情况下, 包名是全小写字符, 单个单词名; 没有必要使用下划线和混合大小写的用法. 这样可以降低使用该包的人可能产生的拼写错误. 

不需要担心包重名的问题, 包名只是在导入时使用, 它在整套的源代码中不要求唯一. 在很少出现的使用两个重名包的时候可以通过使用别名的方式来本地区分. 该情况很少发生, 因为导入时的文件名将决定会使用那个包.

另外的一个惯例, 包名一般是它的源代码目录的基准名字. 例如 **`src/encoding/base64`** 目录中的包名被导入时 是 **`"encoding/base64"`**, 但是使用的名字是 **`base64`**. 并不会加入下划线, 也不会加入其上级目录的名字.

包导入器将使用包名来指向其内容, 该包中的需要暴露出来的对象名可以使用这个机制来避免混淆. (**`不要使用 import . $pkgname 方式来导入包, 这种方式只能作为在该包外进行测试时的简写方式. 其他情况避免使用. `**)

举例来讲, bufio 包中的 buffered reader 类型被命名为 Reader, 而不是 BufReader, 因为用户使用时会看到 bufio.Reader , 这已经是一个清楚,无歧义的名字. 它也不会同 io.Reader产生混淆.

同样情况下, 尝试创建 ring.Ring 的新实例时, 本来命名wei NewRing. 但是由于 Ring 类型是 ring package中唯一暴露出来的对象. 同时该包的名字是 ring. 那么这个包方法的名字被命名为 ring.New.

另一个例子是 once.Do; once.Do(setup) 读起来非常清晰. 使用 once.DoOrWaitUntilDone(setup) 来命名并不能对其产生帮助. 长命名并不能自动的使描述更加可读. 一份有帮助的文档注释将比长命名更加有意义. (**`支持!!! 代码注释即可产生文档, 先文档 再编码!`**)


**Getters**

Go并不会自动产生 getters / setters (**`获取对象/设置对象`**) . 自行提供该功能没有任何问题, 并且是被推荐的. 

假设存在有字段为 owner (小写字符, 私有变量), 它的获取 getter 方法应当被命名为 Owner (首字符大写, public 变量). 而不是 GetOwner. 这种使用变量名首字符大写的方式来提供接口的方式会很明显的将 **`方法`** 同 **`变量`** 结合在一起. 同时, 如果需要提供方法设置该对象. 推荐使用 **`SetOwner`**.  这两个名字在使用中都能很好的被理解:

    owner := obj.Owner()
    if owner != user {
        obj.SetOwner(user)
    }


-----


