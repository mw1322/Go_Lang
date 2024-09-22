package main

import "fmt"

func main() {
    arr := []int{4, 6, 2, 5, 7, 9, 1, 3}
    fmt.Println("Before sorting:", arr)
    QuickSort(arr, 0, len(arr)-1)
    fmt.Println("After sorting:", arr)
}

func QuickSort(arr []int, l, h int) {
    if l < h {
        pivot := partition(arr, l, h)
        QuickSort(arr, l, pivot-1)
        QuickSort(arr, pivot+1, h)
    }
}

func partition(arr []int, l, h int) int {
    pivot := arr[l]
    i := l
    j := h
    for i < j {
        for i <= h && arr[i] <= pivot {
            i++
        }
        for arr[j] > pivot {
            j--
        }
        if i < j {
            swap(arr, i, j)
        }
    }
    swap(arr, l, j)
    return j
}

func swap(arr []int, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}
