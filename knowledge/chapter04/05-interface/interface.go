package main

type File struct {
}

// 尽管 File 类并没有显式实现这些接口，甚至根本不知道这些接口的存在，
// 但是我们说 File 类实现了这些接口，因为 File 类实现了上述所有接口声明的方法。
// 当一个类的成员方法集合包含了某个接口声明的所有方法，
// 换句话说，如果一个接口的方法集合是某个类成员方法集合的子集，我们就认为该类实现了这个接口。
func (f *File) Read(buf []byte) (n int, err error)
func (f *File) Write(buf []byte) (n int, err error)
func (f *File) Seek(off int64, whence int) (pos int64, err error)
func (f *File) Close() error

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}
type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWriter interface {
	Write(buf []byte) (n int, err error)
}
type ICloser interface {
	Close() error
}
