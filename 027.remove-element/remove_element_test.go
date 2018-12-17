package main

import (
    "testing"
)

func TestRemoveElement(t *testing.T) {
    arr := []int{3, 2, 2, 3}
    if  removeElement(arr, 3) != 2 {
        t.Error("length should be 2")
    }

    if arr[0] != 2 || arr[1] != 2 {
        t.Error("now, arr should be [2, 2, 2, 3]")
    }

    arr1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
    if  removeElement(arr1, 2) != 5 {
        t.Error("length should be 5")
    }

    if arr1[0] != 0 || arr1[1] !=  1 || arr1[2] !=  3 || arr1[3] !=  0 || arr1[4] != 4 {
        t.Error("now, arr should be [0 1 3 0 4 0 4 2]")
    }
}
