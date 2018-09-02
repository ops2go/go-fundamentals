Package sort
import "sort"
Overview
Index
Examples
Overview ▾
Package sort provides primitives for sorting slices and user-defined collections.

▹ Example

▹ Example (SortKeys)

▹ Example (SortMultiKeys)

▹ Example (SortWrapper)

Index ▾
func Float64s(a []float64)
func Float64sAreSorted(a []float64) bool
func Ints(a []int)
func IntsAreSorted(a []int) bool
func IsSorted(data Interface) bool
func Search(n int, f func(int) bool) int
func SearchFloat64s(a []float64, x float64) int
func SearchInts(a []int, x int) int
func SearchStrings(a []string, x string) int
func Slice(slice interface{}, less func(i, j int) bool)
func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool
func SliceStable(slice interface{}, less func(i, j int) bool)
func Sort(data Interface)
func Stable(data Interface)
func Strings(a []string)
func StringsAreSorted(a []string) bool
type Float64Slice
    func (p Float64Slice) Len() int
    func (p Float64Slice) Less(i, j int) bool
    func (p Float64Slice) Search(x float64) int
    func (p Float64Slice) Sort()
    func (p Float64Slice) Swap(i, j int)
type IntSlice
    func (p IntSlice) Len() int
    func (p IntSlice) Less(i, j int) bool
    func (p IntSlice) Search(x int) int
    func (p IntSlice) Sort()
    func (p IntSlice) Swap(i, j int)
type Interface
    func Reverse(data Interface) Interface
type StringSlice
    func (p StringSlice) Len() int
    func (p StringSlice) Less(i, j int) bool
    func (p StringSlice) Search(x string) int
    func (p StringSlice) Sort()
    func (p StringSlice) Swap(i, j int)
Examples
Package
Float64s
Float64sAreSorted
Ints
IntsAreSorted
Reverse
Search
Search (DescendingOrder)
Slice
SliceStable
Strings
Package (SortKeys)
Package (SortMultiKeys)
Package (SortWrapper)
Package files
search.go slice.go sort.go zfuncversion.go

func Float64s
func Float64s(a []float64)
Float64s sorts a slice of float64s in increasing order (not-a-number values are treated as less than other values).

▹ Example

func Float64sAreSorted
func Float64sAreSorted(a []float64) bool
Float64sAreSorted tests whether a slice of float64s is sorted in increasing order (not-a-number values are treated as less than other values).

▹ Example

func Ints
func Ints(a []int)
Ints sorts a slice of ints in increasing order.

▹ Example

func IntsAreSorted
func IntsAreSorted(a []int) bool
IntsAreSorted tests whether a slice of ints is sorted in increasing order.

▹ Example

func IsSorted
func IsSorted(data Interface) bool
IsSorted reports whether data is sorted.

func Search
func Search(n int, f func(int) bool) int
Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true, assuming that on the range [0, n), f(i) == true implies f(i+1) == true. That is, Search requires that f is false for some (possibly empty) prefix of the input range [0, n) and then true for the (possibly empty) remainder; Search returns the first true index. If there is no such index, Search returns n. (Note that the "not found" return value is not -1 as in, for instance, strings.Index.) Search calls f(i) only for i in the range [0, n).

A common use of Search is to find the index i for a value x in a sorted, indexable data structure such as an array or slice. In this case, the argument f, typically a closure, captures the value to be searched for, and how the data structure is indexed and ordered.

For instance, given a slice data sorted in ascending order, the call Search(len(data), func(i int) bool { return data[i] >= 23 }) returns the smallest index i such that data[i] >= 23. If the caller wants to find whether 23 is in the slice, it must test data[i] == 23 separately.

Searching data sorted in descending order would use the <= operator instead of the >= operator.

To complete the example above, the following code tries to find the value x in an integer slice data sorted in ascending order:

x := 23
i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
if i < len(data) && data[i] == x {
	// x is present at data[i]
} else {
	// x is not present in data,
	// but i is the index where it would be inserted.
}
As a more whimsical example, this program guesses your number:

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
▹ Example

▹ Example (DescendingOrder)

func SearchFloat64s
func SearchFloat64s(a []float64, x float64) int
SearchFloat64s searches for x in a sorted slice of float64s and returns the index as specified by Search. The return value is the index to insert x if x is not present (it could be len(a)). The slice must be sorted in ascending order.

func SearchInts
func SearchInts(a []int, x int) int
SearchInts searches for x in a sorted slice of ints and returns the index as specified by Search. The return value is the index to insert x if x is not present (it could be len(a)). The slice must be sorted in ascending order.

func SearchStrings
func SearchStrings(a []string, x string) int
SearchStrings searches for x in a sorted slice of strings and returns the index as specified by Search. The return value is the index to insert x if x is not present (it could be len(a)). The slice must be sorted in ascending order.

func Slice
func Slice(slice interface{}, less func(i, j int) bool)
Slice sorts the provided slice given the provided less function.

The sort is not guaranteed to be stable. For a stable sort, use SliceStable.

The function panics if the provided interface is not a slice.

▹ Example

func SliceIsSorted
func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool
SliceIsSorted tests whether a slice is sorted.

The function panics if the provided interface is not a slice.

func SliceStable
func SliceStable(slice interface{}, less func(i, j int) bool)
SliceStable sorts the provided slice given the provided less function while keeping the original order of equal elements.

The function panics if the provided interface is not a slice.

▹ Example

func Sort
func Sort(data Interface)
Sort sorts data. It makes one call to data.Len to determine n, and O(n*log(n)) calls to data.Less and data.Swap. The sort is not guaranteed to be stable.

func Stable
func Stable(data Interface)
Stable sorts data while keeping the original order of equal elements.

It makes one call to data.Len to determine n, O(n*log(n)) calls to data.Less and O(n*log(n)*log(n)) calls to data.Swap.

func Strings
func Strings(a []string)
Strings sorts a slice of strings in increasing order.

▹ Example

func StringsAreSorted
func StringsAreSorted(a []string) bool
StringsAreSorted tests whether a slice of strings is sorted in increasing order.

type Float64Slice
Float64Slice attaches the methods of Interface to []float64, sorting in increasing order (not-a-number values are treated as less than other values).

type Float64Slice []float64
func (Float64Slice) Len
func (p Float64Slice) Len() int
func (Float64Slice) Less
func (p Float64Slice) Less(i, j int) bool
func (Float64Slice) Search
func (p Float64Slice) Search(x float64) int
Search returns the result of applying SearchFloat64s to the receiver and x.

func (Float64Slice) Sort
func (p Float64Slice) Sort()
Sort is a convenience method.

func (Float64Slice) Swap
func (p Float64Slice) Swap(i, j int)
type IntSlice
IntSlice attaches the methods of Interface to []int, sorting in increasing order.

type IntSlice []int
func (IntSlice) Len
func (p IntSlice) Len() int
func (IntSlice) Less
func (p IntSlice) Less(i, j int) bool
func (IntSlice) Search
func (p IntSlice) Search(x int) int
Search returns the result of applying SearchInts to the receiver and x.

func (IntSlice) Sort
func (p IntSlice) Sort()
Sort is a convenience method.

func (IntSlice) Swap
func (p IntSlice) Swap(i, j int)
type Interface
A type, typically a collection, that satisfies sort.Interface can be sorted by the routines in this package. The methods require that the elements of the collection be enumerated by an integer index.

type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}
func Reverse
func Reverse(data Interface) Interface
Reverse returns the reverse order for data.

▹ Example

type StringSlice
StringSlice attaches the methods of Interface to []string, sorting in increasing order.

type StringSlice []string
func (StringSlice) Len
func (p StringSlice) Len() int
func (StringSlice) Less
func (p StringSlice) Less(i, j int) bool
func (StringSlice) Search
func (p StringSlice) Search(x string) int
Search returns the result of applying SearchStrings to the receiver and x.

func (StringSlice) Sort
func (p StringSlice) Sort()
Sort is a convenience method.

func (StringSlice) Swap
func (p StringSlice) Swap(i, j int)