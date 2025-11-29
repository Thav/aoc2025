package lists

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// 3   4
// 4   3
// 2   5
// 1   3
// 3   9
// 3   3
var listsString = []byte("3   4\n4   3\n2   5\n1   3\n3   9\n3   3")

// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9
var levelsString = []byte("7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9")

func TestMain(m *testing.M) {
	exitVal := m.Run()
	log.Println("Tests complete!")

	os.Exit(exitVal)
}

func TestImportLeftRightLists(t *testing.T) {
	left, right := ImportLeftRightLists(listsString, "   ")
	if left[3] != "1" {
		t.Errorf("Expected %v to be \"1\"", left[3])
	}
	if right[4] != "9" {
		t.Errorf("Expected %v to be \"9\"", right[4])
	}
	fmt.Println(left, right)
}

func TestImportLeftRightListsInt(t *testing.T) {
	left, right := ImportLeftRightListsInt(listsString, "   ")
	if left[3] != 1 {
		t.Errorf("Expected %v to be 1", left[3])
	}
	if right[4] != 9 {
		t.Errorf("Expected %v to be 9", right[4])
	}
	fmt.Println(left, right)
}

func TestImportRowLists(t *testing.T) {
	lists := ImportRowLists(levelsString, " ")
	if len(lists) != 6 {
		t.Errorf("Expected number of lists to be 6, got %v", len(lists))
	}
	if len(lists[0]) != 5 {
		t.Errorf("Expected length of lists to be 5, got %v", len(lists[0]))
	}
	if lists[1][3] != "8" {
		t.Errorf("Expected %v to be \"8\"", lists[1][3])
	}
	fmt.Println(lists)
}

func TestImportRowListsInt(t *testing.T) {
	lists := ImportRowListsInt(levelsString, " ")
	if len(lists) != 6 {
		t.Errorf("Expected number of lists to be 6, got %v", len(lists))
	}
	if len(lists[0]) != 5 {
		t.Errorf("Expected length of lists to be 5, got %v", len(lists[0]))
	}
	if lists[1][3] != 8 {
		t.Errorf("Expected %v to be \"8\"", lists[1][3])
	}
	fmt.Println(lists)
}
