package dictionary

/*
以map 关键字开头，需要两种类型。第一个是键的类型，写在 [] 中。第二个是值的类型，跟在 [] 之后.

键的类型比较特别, 必须是 可比较的类型, 参见[Go语言规范]{https://golang.org/ref/spec#Comparison_operators}
值的类型比较简单, 可以是任意类型, 甚至是一个map.
*/

/*
引用类型
Map 有一个有趣的特性，不使用指针传递你就可以修改它们。这是因为 map 是引用类型。
这意味着它拥有对底层数据结构的引用，就像指针一样。它底层的数据结构是 hash table 或 hash map.

Map 作为引用类型是非常好的，因为无论 map 有多大，都只会有一个副本。
引用类型引入了 maps 可以是 nil 值。如果你尝试使用一个 nil 的 map，你会得到一个 nil 指针异常，这将导致程序终止运行。

由于 nil 指针异常，你永远不应该初始化一个空的 map 变量：
var m map[string]string

相反，你可以像我们上面那样初始化空 map，或使用 make 关键字创建 map：
dictionary = map[string]string{}
// OR
dictionary = make(map[string]string)

这两种方法都可以创建一个空的 hash map 并指向 dictionary。这确保永远不会获得 nil 指针异常
*/

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
