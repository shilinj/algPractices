package lsyntax

import "fmt"

// 切片的扩容
func SliceCapAppend() {
	s := make([]int, 3)
	fmt.Println(cap(s))

	s = append(s, 1, 2)
	fmt.Println(cap(s))

	s = append(s, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(cap(s))

	s = append(s, 7)
	fmt.Println(cap(s))

	fmt.Println(s)
}

/*
source file: src/runtime/slice.go
func growslice(et *_type, old slice, cap int) slice {
    ...省略部分
    newcap := old.cap
    doublecap := newcap + newcap
    if cap > doublecap {
        newcap = cap
    } else {
        if old.len < 1024 {
            newcap = doublecap
        } else {
            // Check 0 < newcap to detect overflow
            // and prevent an infinite loop.
            for 0 < newcap && newcap < cap {
                newcap += newcap / 4
            }
            // Set newcap to the requested cap when
            // the newcap calculation overflowed.
            if newcap <= 0 {
                newcap = cap
            }
        }
    }
    ...省略部分
}
当需要的容量超过原切片容量的两倍时，会使用需要的容量作为新容量。

当原切片长度小于1024时，新切片的容量会直接翻倍。而当原切片的容量大于等于1024时，会反复地增加25%，直到新容量超过所需要的容量。

结论
GoLang中的切片扩容机制，与切片的数据类型、原本切片的容量、所需要的容量都有关系，比较复杂。对于常见数据类型，在元素数量较少时，大致可以认为扩容是按照翻倍进行的,但具体情况需要具体分析。
*/
