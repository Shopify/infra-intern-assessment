package main

import (
	"fmt"
	"sort"

	"reflect"
	"testing"
)

// The num_set struct is used to keep track of sets.
type num_set struct {
	set_map map[int]bool //set_map
	row     int          //row coordinate
	column  int          //column coordinate
	square  int          //associated subsquare (1 through 9, top-down, left-right)
	heat    int          //"closeness to solving", high heat means almost deduced. heat = 9-len(set_map)
}

func create_num_set(entries []int, row int, column int, heat int) *num_set {
	//create a num_set given its possible set and coordinates.
	//used for coordinate entries and row/column/square collections, latter having 0 for non applicable coordinate entries.
	var n_set num_set
	n_set.set_map = make(map[int]bool)
	for _, entry := range entries {
		n_set.set_map[entry] = true
	}
	n_set.column = column
	n_set.row = row
	n_set.square = coord_to_subsquare(row, column)
	n_set.heat = heat
	return &n_set
}

func coord_to_subsquare(row int, column int) int {
	//return subsquare (int from 0-8) in which the coordinate resides.
	return column/3 + (row/3)*3
}

func square_to_coord(square int) [2]int {
	//return the top left (least valued) coordinate which the square starts
	return [2]int{(square / 3) * 3, (square % 3) * 3}
}

func update_subsquare(sudoku *[][]*num_set, hot_bucket *[][]int, square int, entry int) {
	//update (remove entry from) the num_sets that are in the subsquare square.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			coord := square_to_coord(square)
			row := coord[0]
			col := coord[1]
			if (*(*sudoku)[row+i][col+j]).set_map[entry] {
				(*(*sudoku)[row+i][col+j]).heat++
				if (*sudoku)[row+i][col+j].heat == 8 {
					*hot_bucket = append(*hot_bucket, []int{i + row, j + col})
				}
			} //heat up if new info
			(*(*sudoku)[row+i][col+j]).set_map[entry] = false //redundant updates for 4 entries, might optimize later
		}
	}
}

func print_sudoku(sudoku [][]*num_set) {
	//print the sudoku board, each column separated by tab
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			fmt.Print(get_entries(sudoku[row][column]), "\t")
		}
		fmt.Print("\n")
	}
}

func print_sudoku_heatmap(sudoku [][]*num_set) {
	//print the heat of each element, each column separated by tab
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			fmt.Print(sudoku[row][column].heat)
			fmt.Print("\t")
		}
		fmt.Print("\n")

	}
}

func copy_num_set(src_set num_set) *num_set {
	return create_num_set(get_entries(&src_set), src_set.row, src_set.column, src_set.heat)
}

func get_hot_sets(sudoku [][]*num_set) []num_set {
	//return ordered slice of num_set which are still hot/unsolved (num_set.heat>=0).
	var list = []num_set{}
	for _, columns := range sudoku {
		for _, set := range columns {
			if set.heat >= 0 {
				list = append(list, *copy_num_set(*set))
			}
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].heat > list[j].heat
	})
	return list
}

func update_sudoku(sudoku *[][]*num_set, hot_bucket *[][]int, rows_slice *[]*num_set, columns_slice *[]*num_set, squares_slice *[]*num_set, row int, column int, entry int) bool {
	//update the sudoku with the newly gained information that entry belongs at coordinate [row][column]. return true if a change was made, i.e. was new info.

	invalid := (*(*rows_slice)[row]).set_map[entry] || (*(*columns_slice)[column]).set_map[entry] || (*(*rows_slice)[coord_to_subsquare(row, column)]).set_map[entry]
	//invalid if all of the subsets (row/column/subsquare) already had entry in their accounted set.

	//update subsets
	(*(*rows_slice)[row]).set_map[entry] = true
	(*(*columns_slice)[column]).set_map[entry] = true
	(*(*rows_slice)[coord_to_subsquare(row, column)]).set_map[entry] = true

	//update num_sets in common row and column
	for i := 0; i < 9; i++ {
		if (*sudoku)[row][i].set_map[entry] {
			((*sudoku)[row][i].heat)++
			if (*sudoku)[row][i].heat == 8 {
				*hot_bucket = append(*hot_bucket, []int{row, i})
			}
		} //heat up if changing true to false - we gained new info about this coordinate!
		(*sudoku)[row][i].set_map[entry] = false // all else in the same row cannot have entry
		if (*sudoku)[i][column].set_map[entry] {
			((*sudoku)[i][column].heat)++
			if (*sudoku)[i][column].heat == 8 {
				*hot_bucket = append(*hot_bucket, []int{i, column})
			}
		} //heat up if changing true to false - we gained new info about this coordinate!
		(*sudoku)[i][column].set_map[entry] = false // all else in same column cannot have entry
	}
	update_subsquare(sudoku, hot_bucket, coord_to_subsquare(row, column), entry)

	//set everything (else) to false
	for i := 1; i <= 9; i++ {
		(*sudoku)[row][column].set_map[i] = false
	}
	(*sudoku)[row][column].set_map[entry] = true
	(*sudoku)[row][column].heat = -1 //solved, no longer hot
	return invalid
}

func init_sudoku(sudoku *[][]*num_set) {
	//initialize sudoku board of *num_sets with a full board of all possibilities and no progress.
	*sudoku = make([][]*num_set, 9)
	full_possibility := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for row := 0; row < 9; row++ {
		cur_row := make([]*num_set, 9)
		for column := 0; column < 9; column++ {
			cur_row[column] = create_num_set(full_possibility[:], row, column, 0)
		}
		(*sudoku)[row] = cur_row
	}
}

func copy_map(dest *num_set, src *num_set) {
	for i := 1; i <= 9; i++ {
		(*dest).set_map[i] = (*src).set_map[i]
	}
}

func copy_sudoku(dest *[][]*num_set, src *[][]*num_set) {
	*dest = make([][]*num_set, 9)
	for row := 0; row < 9; row++ {
		cur_row := make([]*num_set, 9)
		for column := 0; column < 9; column++ {
			cur_row[column] = copy_num_set((*(*src)[row][column]))
		}
		(*dest)[row] = cur_row
	}
}

func SolveSudoku(sudoku [][]int) [][]int {
	//returns a sudoku in the form of slice of slice of num_set struct

	//create slice of num_sets. each num_set represents that subset's used numbers.
	//update: original intention was to use for heatmapping, but now serves more as double-check
	var rows_slice = []*num_set{}
	var columns_slice = []*num_set{}
	var squares_slice = []*num_set{}

	//hot bucket is the slice of coordinates at which have only 1 possible remaining number in their num_set
	var hot_bucket = [][]int{}
	//initialize [][]*num_set sudoku. Every coordinate has a num_set with its set_map as {1: true, 2:true, ... 9:true}
	var out_sudoku = [][]*num_set{}
	init_sudoku(&out_sudoku)

	//create subset slices: an array with each entry representing the set of already used numbers for that row/column/subsquare
	empty_array := [0]int{}
	for i := 0; i < 9; i++ {
		rows_slice = append(rows_slice, create_num_set(empty_array[:], i, 0, 0))
		columns_slice = append(columns_slice, create_num_set(empty_array[:], 0, i, 0))
		squares_slice = append(squares_slice, create_num_set(empty_array[:], 0, 0, i))
	}

	//look through every entry in the sudoku and update the num_sets of each coordinate and subsets along the way

	//tester := 0
	for row, row_entries := range sudoku {
		for column, entry := range row_entries {
			if entry != 0 {
				update_sudoku(&out_sudoku, &hot_bucket, &rows_slice, &columns_slice, &squares_slice, row, column, entry)
				// if tester < 20 {
				// 	print_sudoku(out_sudoku)
				// 	print_sudoku_heatmap(out_sudoku)
				// }
				// tester++
			}
		}
	}
	//print_sudoku(out_sudoku)
	//print_sudoku_heatmap(out_sudoku)

	//at this point, out_sudoku should have "grid" of num_sets which represent possible entries
	//start by updating hot_buckets, since they are straightforward fill and update.

	//made_change := false //checking for the loop in which nothing was updated, meaning
	for len(hot_bucket) > 0 {
		//fmt.Println("\n\n + Hotbucket", hot_bucket)

		cur_row := hot_bucket[0][0]
		cur_col := hot_bucket[0][1]
		//fmt.Println(cur_row, cur_col)
		hot_bucket = hot_bucket[1:] //pop first item
		//fmt.Println("\n\n + Hotbucket", hot_bucket)

		set_slice := get_entries(out_sudoku[cur_row][cur_col])
		update_sudoku(&out_sudoku, &hot_bucket, &rows_slice, &columns_slice, &squares_slice, cur_row, cur_col, set_slice[0])
		//print_sudoku(out_sudoku)

		// if tester > 3 {
		// 	break
		// }
		// tester++
	}
	//hypothesis testing /backtracking needed from this point.
	hit_lists := [][]num_set{get_hot_sets(out_sudoku)} //hit_lists is the list of lists of coordinates to hit up.
	sudokus := [][][]*num_set{out_sudoku}              //sudokus is the list of sudokus that have been saved for reverting.
	entries := []int{}
	depth := 0
	//tester := 0
	broken := false
	// hit_list := hit_lists[0]
	// entries := get_entries(&hit_list[0])
	//fmt.Println(hit_lists[depth])
	for len(hit_lists[0]) > 0 {
		// tester++
		// if tester > 20 {
		// 	break
		// }
		entries = get_entries(&(hit_lists[depth][0]))
		//fmt.Println("\nentries\t", entries) //????
		if len(entries) == 0 {
			//when empty, know that its time to go out one depth and cross out this path as a possibility
			copy_sudoku(&out_sudoku, &(sudokus[depth-1]))
			sudokus = sudokus[:depth]
			hit_lists = hit_lists[:depth]
			depth--
			(hit_lists[depth][0]).set_map[get_entries(&(hit_lists[depth][0]))[0]] = false
			continue
		}

		// fmt.Println("saved:")
		// print_sudoku(out_sudoku)
		// fmt.Print("\n")

		// hit_list := hit_lists[depth]
		// fmt.Println(hit_list)
		//fmt.Println("depth:", depth)
		row := hit_lists[depth][0].row
		col := hit_lists[depth][0].column
		entries := get_entries(&hit_lists[depth][0])
		broken = false
		//fmt.Println("row:", row, "col:", col, "entry:", entries[0])

		update_sudoku(&out_sudoku, &hot_bucket, &rows_slice, &columns_slice, &squares_slice, row, col, entries[0])

		//fmt.Println("\n\n + Hotbucket", hot_bucket)
		//print_sudoku(out_sudoku)

		//solve as much as possible, see if it leads to having to make another guess or if it leads to an invalid sudoku
		for len(hot_bucket) > 0 {
			//fmt.Println("\n\n + Hotbucket", hot_bucket)

			cur_row := hot_bucket[0][0]
			cur_col := hot_bucket[0][1]
			//fmt.Println(cur_row, cur_col)
			hot_bucket = hot_bucket[1:] //pop first item

			set_slice := get_entries(out_sudoku[cur_row][cur_col])
			if len(set_slice) == 0 {
				broken = true
				break
			}
			update_sudoku(&out_sudoku, &hot_bucket, &rows_slice, &columns_slice, &squares_slice, cur_row, cur_col, set_slice[0])
			//print_sudoku(out_sudoku)
			//fmt.Print("\n----", broken)
		}

		if len(get_hot_sets(out_sudoku)) == 0 {
			//solved = true
			break //if sudoku is no longer hot at any point, solved.
		}
		if broken { // if broken, meaning guess was incorrect, restore to last saved sudoku, try the next possible.
			//fmt.Println("broken")
			hot_bucket = [][]int{} // clear hot bucket
			copy_sudoku(&out_sudoku, &(sudokus[depth]))

			//hit_list = hit_lists[depth]
			(hit_lists[depth][0]).set_map[get_entries(&(hit_lists[depth][0]))[0]] = false
			//fmt.Println("restore:")
			//print_sudoku(out_sudoku)
			//fmt.Println(hit_lists[depth])
		} else { //if not broken, but not solved, save current sudoku, go deeper in exploration.
			new_sudoku := [][]*num_set{}
			copy_sudoku(&new_sudoku, &out_sudoku)
			sudokus = append(sudokus, new_sudoku)
			//print_sudoku(out_sudoku)
			//fmt.Println("\n----")
			//fmt.Println(hit_lists[depth])
			hit_lists = append(hit_lists, get_hot_sets(out_sudoku))
			depth++
		}

	}
	return convert_board(out_sudoku) //convert to [][]int format and return.
}

func get_entries(n_set *num_set) []int {
	//return an integer slice of the set.
	set_slice := []int{}
	for i := 1; i <= 9; i++ {
		if (*n_set).set_map[i] {
			set_slice = append(set_slice, i)
		}
	}
	return set_slice
}

func convert_board(sudoku [][]*num_set) [][]int {
	//return a [][]int type sudoku board given a [][]*num_set sudoku board. This function expects [][]*num_set type sudoku board to be solved.
	var sudoku_arr = make([][]int, 9)
	for i := range sudoku_arr {
		sudoku_arr[i] = make([]int, 9)
	}
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if len(get_entries(sudoku[row][column])) != 1 {
				sudoku_arr[row][column] = -1 //error -1 if more than 1 possible element in the set
			} else {
				sudoku_arr[row][column] = get_entries(sudoku[row][column])[0]
			}
		}
	}
	return sudoku_arr
}

func TestSolveSudoku(t *testing.T) {
	input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	expected := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	solved := SolveSudoku(input)
	fmt.Print(solved)
	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}

// func main() {
// 	var sudoku = make([][]*num_set, 9)
// 	init_sudoku(&sudoku)
// 	//print_sudoku(sudoku)
// 	fmt.Println(quote.Go())
// 	input := [][]int{
// 		{5, 3, 0, 0, 7, 0, 0, 0, 0},
// 		{6, 0, 0, 1, 9, 5, 0, 0, 0},
// 		{0, 9, 8, 0, 0, 0, 0, 6, 0},
// 		{8, 0, 0, 0, 6, 0, 0, 0, 3},
// 		{4, 0, 0, 8, 0, 3, 0, 0, 1},
// 		{7, 0, 0, 0, 2, 0, 0, 0, 6},
// 		{0, 6, 0, 0, 0, 0, 2, 8, 0},
// 		{0, 0, 0, 4, 1, 9, 0, 0, 5},
// 		{0, 0, 0, 0, 8, 0, 0, 7, 9},
// 	}
// 	// input := [][]int{
// 	// 	{0, 6, 9, 8, 0, 0, 0, 0, 0},
// 	// 	{3, 0, 4, 7, 0, 6, 5, 0, 0},
// 	// 	{2, 0, 0, 5, 9, 0, 6, 0, 0},
// 	// 	{0, 2, 0, 0, 0, 0, 0, 1, 0},
// 	// 	{9, 0, 0, 1, 0, 7, 0, 0, 8},
// 	// 	{0, 3, 0, 0, 0, 0, 0, 4, 0},
// 	// 	{0, 0, 2, 0, 5, 4, 0, 0, 7},
// 	// 	{0, 0, 3, 2, 0, 1, 4, 0, 5},
// 	// 	{0, 0, 0, 0, 0, 9, 2, 6, 0},
// 	// }
// 	// input := [][]int{
// 	// 	{0, 1, 0, 0, 6, 0, 0, 0, 4},
// 	// 	{0, 0, 4, 0, 1, 5, 0, 6, 0},
// 	// 	{0, 6, 0, 0, 3, 0, 8, 0, 0},
// 	// 	{5, 0, 9, 7, 0, 0, 0, 0, 0},
// 	// 	{1, 0, 3, 0, 5, 0, 0, 0, 0},
// 	// 	{7, 0, 0, 2, 0, 0, 0, 0, 0},
// 	// 	{0, 0, 0, 0, 0, 0, 0, 3, 1},
// 	// 	{9, 0, 0, 0, 0, 0, 0, 0, 0},
// 	// 	{0, 0, 8, 0, 0, 0, 0, 4, 0},
// 	// }
// 	solved := SolveSudoku(input)
// 	//fmt.Print(input)
// 	fmt.Print(solved)
// 	//fmt.Println("test")
// }
