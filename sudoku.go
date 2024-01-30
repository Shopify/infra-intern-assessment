package main

import "fmt"
import "sort"

/************ Some useful helper functions and type definitions *************/
type cell struct {
    row int 
    col int
}

//returns the number of first set bit in avail
func getNum(avail int) int {
    num := 0;
    avail >>= 1;
    for avail > 0 {
        num++;
        avail >>= 1;
    }

    return num;
}

func isPow2(n int) bool {
    return (n & (n-1)) == 0;
}

func countBits(n int) int {
    count := 0;
    for n > 0 {
        if ((n & 1) == 1) {
            count++;
        }
        n >>= 1;
    }
    return count;
}

type mapIterations func(int, int) (int, int);

//when used in a loop for i from 0 to 9, j from 0 to 9, 
//row mapper will give coordinates that traverse row
//col mapper will give coordinates that traverse columns 
//box mapper will give coordinates that traverse box by box

func byRowMapper(i, j int) (int, int) {
    return i, j;
}

func byColMapper(i, j int) (int, int) {
    return j, i;
}

func byBoxMapper(i, j int) (int, int) {
    //convert i = 0 to iterate through first square, i = 1 iterate through second square etc...
    //so for example 3, 3 will be the 4th square, 4th element (row 3, 0)
    return (i/3)*3 + j/3, (i%3)*3+j%3;
}

/************** Main Solving Functions ******************/
func SolveSudoku(input [][]int) [][]int {

	//first populate the set of possible numbers for each row, col, box
	//a set bit of 1 represents that that number is possible
    var rowAvail [9]int
    var colAvail [9]int
    var boxAvail [9]int

    for i := 0; i < 9; i++ {
        rowAvail[i] = 0b1111111110;
        colAvail[i] = 0b1111111110;
        boxAvail[i] = 0b1111111110;
    }

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if input[i][j] == 0 {
                continue;
            }

            rowAvail[i] ^= (1 << input[i][j]);
            colAvail[j] ^= (1 << input[i][j]);
            boxAvail[(i/3)*3 + j/3] ^= (1 << input[i][j]);
        }
    }

	//populate the set of possible numbers for each cell
	//this can be taken with the bitwise and of the
	//possible values for the cells row, col, and box
    available := make([][]int, 9)
    for i := 0; i < 9; i++ {
        available[i] = make([]int, 9);
        for j := 0; j < 9; j++ {
            available[i][j] = (rowAvail[i] & colAvail[j] & boxAvail[(i/3)*3+j/3]);
        }
    }

    SolveSudokuRecursive(input, available);
    return input; 
}

//Main Driver Function
func SolveSudokuRecursive(board [][]int, available [][]int) bool {
	//make a local copy of the board state in case there is no solution
	//so we must restore state. This can happen when guessing 
    boardLocal := [9][9]int{}
    availableLocal := [9][9]int{}
    onlyOneOption := make(chan cell, 81)
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            boardLocal[i][j] = board[i][j]
            availableLocal[i][j] = available[i][j];

            if(board[i][j] == 0 && isPow2(available[i][j])) {
                onlyOneOption <- cell{i, j}
            }
        }
    }

    success := true;
    successes := [9]bool{};
    made_changes := [9]bool{};
    for cond := true; cond && success ; {
        cond = false;

		//apply algorithm described in the DESIGN.md 
        if(!onlyOneOptionElimination(board, available, onlyOneOption)){ 
            success = false;
            break;
        }

        successes[0], made_changes[0] = boxRowColElimination(board, available, onlyOneOption);
        successes[1], made_changes[1] = forcedCellElimination(board, available, onlyOneOption);
        successes[2], made_changes[2] = NTupleElimination(board, available, onlyOneOption, byRowMapper);
        successes[3], made_changes[3] = NTupleElimination(board, available, onlyOneOption, byColMapper);
        successes[4], made_changes[4] = NTupleElimination(board, available, onlyOneOption, byBoxMapper);
        successes[5], made_changes[5] = HiddenPairElimination(board, available, onlyOneOption, byRowMapper);
        successes[6], made_changes[6] = HiddenPairElimination(board, available, onlyOneOption, byColMapper);
        successes[7], made_changes[7] = HiddenPairElimination(board, available, onlyOneOption, byBoxMapper);

		//if one of these functions failed, then there is no solution
		//if there were changes made in any function, then we can try to perform more reductions
        for i := 0; i <= 7; i++ {
            success = success && successes[i];
            cond = cond || made_changes[i];
        }
    }

	//No solution, restore board and state.
    if(!success) { 
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                board[i][j] = boardLocal[i][j];
                available[i][j] = availableLocal[i][j];
            }
        }
        return false;
    }

	//check for solved
    solved := true;
    for i := 0; i < 9 && solved; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == 0 {
                solved = false; 
                break;
            }
        }
    }

    //In the case where analytical techniqyes can't solve, a guess must be made
    if(!solved) {
        guessQueue := []cell{};
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                if(board[i][j] == 0){
                    guessQueue = append(guessQueue, cell{i, j});
                }
            }
        }

        //sort cells by number of available numbers to the cell
		//this increases chances that the guess is correct. 
		//(If only 2 numbers available, there is 50% change but if 5 numbers available only 20%)
        sort.Slice(guessQueue, func(i, j int) bool {
		    return countBits(available[guessQueue[i].row][guessQueue[i].col]) < countBits(available[guessQueue[j].row][guessQueue[j].col])
        });

		//perform the guesses, in sorted order
        for len(guessQueue) > 0 {
            curCellToGuess := guessQueue[0];
            row := curCellToGuess.row;
            col := curCellToGuess.col;
            for i := 1;i <= 9;i++ {
                numMask := 1 << i;

                if((available[row][col] & numMask) != 0){ 
                    oldAvail := available[row][col];

                    available[row][col] = numMask;

					//try to solve with this guess, if success then return, 
					//else try a different guess
                    if(SolveSudokuRecursive(board, available)){
                        return true;
                    }

                    available[row][col] = oldAvail;
                }

            }

            guessQueue = guessQueue[1:];
        }
    }

	//If no guesses worked, then restore board and return false.
    if(!solved) {
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                board[i][j] = boardLocal[i][j];
                available[i][j] = availableLocal[i][j];
            }
        }
        return false;
    }

    return true;
}

//Helper function that checks for valid state in the board
//and also removes bits set in mask from cell at x, y (row, col)
//also populate queue for cells with only one option.
func removeAvailableMask(x, y, curX, curY, mask int, available, board [][]int, q chan cell) (bool, bool) {
    if(x == curX && y == curY) {
        return true, false;
    }
    if(((1 << board[x][y]) & mask) != 0) { 
        return false, false;
    }
    if(board[x][y] != 0) {
        return true, false;
    }
    if((available[x][y] & mask) == 0) {
        return true, false;
    }

    available[x][y] &= ^mask; 
    if(available[x][y] == 0){
        return false, false;
    }

    if(isPow2(available[x][y])){
        q <- cell{x,y}
    }
    return true, true;
}

//This processes all cells that only have one possible option
func onlyOneOptionElimination(board, available  [][] int, q chan cell) bool {
    for len(q) > 0 {
        curCell := <- q
		row := curCell.row;
		col := curCell.col;

        num := getNum(available[row][col]); 
        board[row][col] = num;
        mask := 1 << num

		//remove the number that was set from everything else in the cell's row, box, and column
        for i := 0;i < 9; i++ {
            boxRow := (row/3)*3 + i/3
            boxCol := (col/3)*3 + i%3

            result, _ := removeAvailableMask(row, i, row, col, mask, available, board, q);
            result1, _ := removeAvailableMask(i, col, row, col, mask, available, board, q);
            result2, _ :=  removeAvailableMask(boxRow, boxCol, row, col, mask, available, board, q);
            if (!result || !result1 || !result2) {
                return false;
            }
        }
    }
    return true;
}

//This does step 3 as described in DESIGN.md
func forcedCellElimination(board, available  [][] int, q chan cell) (bool, bool) {
    var rowOnce [9]int; var rowMore [9]int; var rowJustOnce [9]int; 
    var colOnce [9]int; var colMore [9]int; var colJustOnce [9]int;
    var boxOnce [9]int; var boxMore [9]int; var boxJustOnce [9]int;
    made_changes := false

	//for each row, column, box, figure out what numbers 
	//occur at least once, and what numbers occur more than once
    for i := 0;i < 9; i++ {
        for j := 0;j < 9; j++ {
            if(board[i][j] != 0){
                continue;
            }
            boxNum := (i/3)*3 + j/3;

            rowMore[i] |= (available[i][j] & rowOnce[i]);
            rowOnce[i] |= available[i][j];
            colMore[j] |= (available[i][j] & colOnce[j]);
            colOnce[j] |= available[i][j];

            boxMore[boxNum] |= (available[i][j] & boxOnce[boxNum]);
            boxOnce[boxNum] |= available[i][j];
        }
    }

	//then the numbers that occur just once are the numbers 
	//that occur at least once and not more than once
    for i := 0;i < 9; i++ {
        rowJustOnce[i] = rowOnce[i] & (^rowMore[i]);
        colJustOnce[i] = colOnce[i] & (^colMore[i]);
        boxJustOnce[i] = boxOnce[i] & (^boxMore[i]);
    }

	//Now these can be populated into the cells
    for i := 0;i < 9; i++ {
        for j := 0;j < 9;j++ {
            if(board[i][j] != 0) {
                continue;
            }

            if((available[i][j] & rowJustOnce[i]) != 0 && ((available[i][j] & rowJustOnce[i]) != available[i][j])){
                made_changes = true;
                available[i][j] &= rowJustOnce[i]
                q <- cell{i, j}
            }

            if((available[i][j] & colJustOnce[j]) != 0 && ((available[i][j] & colJustOnce[j]) != available[i][j])){
                made_changes = true;
                available[i][j] &= colJustOnce[j]
                q <- cell{i, j}
            }

            boxNum := (i/3)*3 + j/3;
            if((available[i][j] & boxJustOnce[boxNum]) != 0 && ((available[i][j] & boxJustOnce[boxNum]) != available[i][j])){
                made_changes = true;
                available[i][j] &= boxJustOnce[boxNum]
                q <- cell{i, j}
            }
        }
    }

    return true, made_changes;
}

//Step 4 in the described algorithm
func boxRowColElimination(board [][]int, available [][]int, oneToEliminate chan cell) (bool, bool) {
    boxInQ := [9]bool{true,true,true,true,true,true,true,true,true};
    q := make(chan int, 9)
    made_changes := false
    for i := 0;i < 9;i++ {
        q <- i
    }

    var rowAvail [3]int;
    var colAvail [3]int;
	
	//process each box in the queue. 
	//If a change was made to a cell in the box, then reprocess the box
    for len(q) > 0 {
        curBox := <- q;
        col := (curBox%3)*3; row := (curBox/3)*3;

        colAvail[0] = 0; colAvail[1] = 0; colAvail[2] = 0;
        rowAvail[0] = 0; rowAvail[1] = 0; rowAvail[2] = 0;
        boxInQ[curBox] = false;

		//for each row and column in the box, 
		//find numbers that are available to each cell in the row/column
        for i := 0;i < 3;i++ {
            for j := 0;j < 3;j++ {
                if(board[row+i][col+j] != 0){
                    continue;
                }
                colAvail[j] |= available[row+i][col+j];
                rowAvail[i] |= available[row+i][col+j];
            }
        }

        for i := 0;i < 3;i++ {

            curCol := colAvail[i];
            curRow := rowAvail[i];
			//find the numbers only available to the row/column
            for j := 0;j < 3;j++ {
                if(i == j){
                    continue;
                }

                curCol &= ^colAvail[j];
                curRow &= ^rowAvail[j];
            }

			//remove the numbers from each square outside the box in the row/column
            for k := 0; k < 9; k++ {
                if(k < col || k > col+2) {
                    res, masked := removeAvailableMask(row+i, k, row, col, curRow, available, board, oneToEliminate)
                    if(!res){
                        return false, made_changes;
                    }
                    if(masked){
                        made_changes = true;
                        box := ((row+i)/3)*3 + k/3;
                        if(!boxInQ[box]){
                            boxInQ[box] = true;
                            q <- box;
                        }
                    }
                }

                if(k > row+2 || k < row) {
                    res, masked := removeAvailableMask(k, col+i, row, col, curCol, available, board, oneToEliminate)
                    if(!res){
                        return false, made_changes;
                    }
                    if(masked){
                        made_changes = true;
                        box := (k/3)*3 + (col+i)/3;
                        if(!boxInQ[box]){
                            boxInQ[box] = true;
                            q <- box;
                        }
                    }
                }
            }
        }
    }
    return true, made_changes;
}

//Step 5 in the DESIGN.md
func NTupleElimination(board, available [][] int, onlyOneRemaining chan cell, coordMapper mapIterations) (bool, bool) {
    var group_map map[int]int
    made_changes := false;

	//traverse each group using the defined mapIterations function type
    for i := 0; i < 9; i++ {
        group_map = make(map[int]int)

		//count the frequencies of the tuples of available numbers
        for j := 0;j < 9; j++ {
            row, col := coordMapper(i, j);
            if(board[row][col] != 0) {
                continue;
            }

            group_map[available[row][col]] = group_map[available[row][col]] + 1;
        }

		//for each tuple of available numbers, if the number of numbers in the tuple
		//is equal to the number of cells with that tuple, then these available numbers
		//can only go in the cells with that tuple. 
		//Remove all number in this tuple from all other cells in the group
        for mask, num := range group_map {
            if(num != countBits(mask)) {
                continue;
            }

            for j := 0;j < 9; j++ {
                row, col := coordMapper(i, j);
                if(mask == available[row][col]){
                    continue;
                }
                success, masked := removeAvailableMask(row, col, -1, -1, mask, available, board, onlyOneRemaining);
                if(!success) {
                    return false, made_changes;
                }

                made_changes = made_changes || masked;
            }
        }

    }

    //can be done for row/col/square
    return true, made_changes;
}

//Step 6 in the algorithm
func HiddenPairElimination(board, available [][] int, onlyOneRemaining chan cell, coordMapper mapIterations) (bool, bool) {
    var freq_map map[int]int;
    made_changes := false;

    for i := 0;i < 9;i++ {
		//cellMask[i] describes which cells the number i is available in
		//eg if cellMask[1] was 0b110, then 1 is possible in squares 1 and 2.
        cellMask := [10]int{0,0,0,0,0,0,0,0,0,0}
        freq_map = make(map[int]int)
        for j := 0;j < 9;j++ {
            row, col := coordMapper(i, j);
            if(board[row][col] != 0) {
                continue;
            }

            for k := 1;k <= 9;k++ {
                if((available[row][col] & (1 << k)) != 0){
                    cellMask[k] ^= (1 << j);
                }
            }
        }

        for j := 1; j <= 9; j++ {
            freq_map[cellMask[j]]++;
        }

		//if there N numbers with the same cell mask of N cells,
		//then none of the N cells can contain anything by those N numbers
		//eg if numbers 1 and 3 can only go in cell 5, 6, 
		//then cell 5 and 6 can only have 1 and 3. 
		//If there were any other possible numbers listed for cell 5, 6 remove them
        for mask, num := range freq_map {
            if num <= 1 {
                continue;
            }
            if num != countBits(mask) {
                continue;
            }
            removalMask := 0;
            for j := 1;j <= 9;j++ {
                if cellMask[j] == mask {
                    removalMask |= (1 << j); 
                }
            }
            removalMask = 0b1111111110 & (^removalMask);

			//remove the other numbers from the cells
            for j := 0;j < 9;j++ {
                if((mask & (1 << j)) == 0) {
                    continue;
                }

                row, col := coordMapper(i,j);
                if(board[row][col] != 0) {
                    continue;
                }
                success, masked := removeAvailableMask(row, col, -1, -1, removalMask, available, board, onlyOneRemaining);
                if(!success) {
                    return false, made_changes;
                }
                made_changes = made_changes || masked;
            }
        }
    }
    return true, made_changes;
}

//main testing function for testing purposes :)
func main() {
    input := [][][]int{
        {
            {5, 3, 0, 0, 7, 0, 0, 0, 0},
            {6, 0, 0, 1, 9, 5, 0, 0, 0},
            {0, 9, 8, 0, 0, 0, 0, 6, 0},
            {8, 0, 0, 0, 6, 0, 0, 0, 3},
            {4, 0, 0, 8, 0, 3, 0, 0, 1},
            {7, 0, 0, 0, 2, 0, 0, 0, 6},
            {0, 6, 0, 0, 0, 0, 2, 8, 0},
            {0, 0, 0, 4, 1, 9, 0, 0, 5},
            {0, 0, 0, 0, 8, 0, 0, 7, 9},
        },
        {
            {6, 0, 7, 1, 8, 0, 3, 0, 0},
            {0, 0, 0, 0, 3, 0, 0, 0, 2},
            {0, 0, 5, 0, 0, 0, 0, 0, 0},
            {0, 2, 0, 8, 0, 0, 0, 0, 0},
            {0, 5, 0, 0, 0, 0, 6, 0, 0},
            {8, 0, 6, 0, 0, 7, 0, 0, 4},
            {0, 0, 0, 0, 0, 4, 0, 9, 0},
            {0, 8, 0, 0, 0, 0, 0, 0, 0},
            {1, 0, 3, 7, 0, 0, 2, 0, 0},
        },

        {
            {0, 0, 9, 2, 0, 3, 8, 0, 0},
            {0, 0, 0, 0, 0, 9, 0, 0, 0},
            {4, 0, 8, 6, 0, 5, 1, 0, 3},
            {1, 0, 2, 0, 0, 0, 9, 0, 4},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {8, 0, 3, 0, 0, 0, 5, 0, 2},
            {9, 0, 6, 5, 0, 2, 3, 0, 7},
            {0, 0, 1, 0, 0, 0, 0, 0, 0},
            {0, 0, 5, 4, 0, 8, 6, 0, 0},
        },

        {
            {0, 0, 5, 3, 0, 7, 0, 0, 0},
            {9, 0, 2, 5, 0, 0, 0, 0, 0},
            {0, 7, 0, 0, 2, 0, 0, 1, 6},
            {0, 3, 0, 0, 0, 0, 0, 4, 0},
            {6, 0, 0, 0, 0, 0, 2, 0, 0},
            {0, 0, 0, 9, 0, 5, 0, 0, 0},
            {0, 9, 0, 0, 0, 0, 0, 0, 3},
            {0, 1, 0, 0, 0, 0, 7, 0, 4},
            {0, 0, 0, 4, 8, 2, 0, 0, 0},
        },
        {
            {0, 0, 0, 2, 0, 7, 0, 0, 4},
            {9, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 8, 2, 6, 0, 0, 1, 0, 0},
            {0, 0, 0, 0, 2, 0, 0, 1, 0},
            {0, 3, 6, 8, 0, 0, 2, 0, 0},
            {0, 0, 5, 0, 0, 0, 0, 0, 0},
            {5, 0, 0, 0, 0, 0, 3, 0, 0},
            {1, 0, 0, 9, 0, 0, 0, 0, 0},
            {0, 9, 3, 0, 0, 6, 0, 7, 0},
        },
    }


    for in := 0; in < len(input); in++ {
        fmt.Println("================ ORIGINAL =================")
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                fmt.Printf("%2d ", input[in][i][j]);
            }
            fmt.Println();
        }

        res := SolveSudoku(input[in]);

		//verifying sudoku output
        rowMasks := [9]int{0,0,0,0,0,0,0,0,0}
        colMasks := [9]int{0,0,0,0,0,0,0,0,0}
        squareMasks := [9]int{0,0,0,0,0,0,0,0,0}
        failed := false;
        for i := 0; i < 9 && !failed; i++ {
            for j := 0; j < 9 && !failed; j++ {
                shifted := 1 << res[i][j]
                square := (i/3)*3 + j/3;
                if((rowMasks[i] & shifted) != 0 || (colMasks[j] & shifted) != 0 || (squareMasks[square] & shifted) != 0){
                    failed = true;
                    break;
                }

                rowMasks[i] |= shifted;
                colMasks[j] |= shifted;
                squareMasks[square] |= shifted;
            }
        }

        if(failed) {
            fmt.Println("================ FAILED =================")
        } else {
            fmt.Println("================ SOLVED =================")
        }
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                fmt.Printf("%2d ", res[i][j]);
            }
            fmt.Println();
        }
    }
}
