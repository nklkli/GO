package main

// Echo4 prints its command-line arguments. Book page 34
import (
	"flag"
	"fmt"
	"io"
	"os"

	// "slices"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("sep", " ", "separator")

func main() {

	s := []rune{'Z', 'F', 'G', '4', 'A', 'W', 'X', 'B'}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	for _, x := range s {
		fmt.Println(string(x))
	}

	os.Exit(0)

	// slices.SortFunc(s, func(a, b int) int {
	// 	return b - a
	// })

	// for _, x := range s {
	// 	fmt.Println(x)
	// }

	// os.Exit(0)

	var funcs []func()

	for i := range 5 {
		funcs = append(funcs, func() {
			fmt.Printf("func %d  %p\n", i, &i)
		})
	}

	for i := range 5 {

		funcs[i]()
	}

	os.Exit(0)

	for _, f := range os.Args[1:] {
		file, err := os.Open(f)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		// fmt.Println(f)
		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(2)
		}

	}

	os.Exit(0)

	arr1 := [...]int{1, 2, 3}
	arr2 := arr1

	arr2[0] = 34234234

	fmt.Println("arr1[0]=", arr1[0], "arr2[0]=", arr2[0])

	m := make(map[string]int)

	if f, ok := m["hello"]; ok {
		fmt.Println("ok", f)
	} else {
		fmt.Println("not ok=", f)
	}

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func delta(old, new int) int { return new - old }
