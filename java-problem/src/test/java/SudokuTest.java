import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class SudokuTest {

    @Test
    public void testSolveSudoku() {
        int[][] input = {
            {5, 3, 0, 0, 7, 0, 0, 0, 0},
            {6, 0, 0, 1, 9, 5, 0, 0, 0},
            {0, 9, 8, 0, 0, 0, 0, 6, 0},
            {8, 0, 0, 0, 6, 0, 0, 0, 3},
            {4, 0, 0, 8, 0, 3, 0, 0, 1},
            {7, 0, 0, 0, 2, 0, 0, 0, 6},
            {0, 6, 0, 0, 0, 0, 2, 8, 0},
            {0, 0, 0, 4, 1, 9, 0, 0, 5},
            {0, 0, 0, 0, 8, 0, 0, 7, 9},
        };

        int[][] expected = {
            {5, 3, 4, 6, 7, 8, 9, 1, 2},
            {6, 7, 2, 1, 9, 5, 3, 4, 8},
            {1, 9, 8, 3, 4, 2, 5, 6, 7},
            {8, 5, 9, 7, 6, 1, 4, 2, 3},
            {4, 2, 6, 8, 5, 3, 7, 9, 1},
            {7, 1, 3, 9, 2, 4, 8, 5, 6},
            {9, 6, 1, 5, 3, 7, 2, 8, 4},
            {2, 8, 7, 4, 1, 9, 6, 3, 5},
            {3, 4, 5, 2, 8, 6, 1, 7, 9},
        };

        int[][] solved = Sudoku.solveSudoku(input);

        assertArrayEquals(expected, solved, "Sudoku puzzle was not solved correctly");
    }

    @Test
    public void testSolveSudokuSecondPuzzle() {
        int[][] input = {
            {0, 0, 0, 2, 6, 0, 7, 0, 1},
            {6, 8, 0, 0, 7, 0, 0, 9, 0},
            {1, 9, 0, 0, 0, 4, 5, 0, 0},
            {8, 2, 0, 1, 0, 0, 0, 4, 0},
            {0, 0, 4, 6, 0, 2, 9, 0, 0},
            {0, 5, 0, 0, 0, 3, 0, 2, 8},
            {0, 0, 9, 3, 0, 0, 0, 7, 4},
            {0, 4, 0, 0, 5, 0, 0, 3, 6},
            {7, 0, 3, 0, 1, 8, 0, 0, 0},
        };

        int[][] expected = {
            {4, 3, 5, 2, 6, 9, 7, 8, 1},
            {6, 8, 2, 5, 7, 1, 4, 9, 3},
            {1, 9, 7, 8, 3, 4, 5, 6, 2},
            {8, 2, 6, 1, 9, 5, 3, 4, 7},
            {3, 7, 4, 6, 8, 2, 9, 1, 5},
            {9, 5, 1, 7, 4, 3, 6, 2, 8},
            {5, 1, 9, 3, 2, 6, 8, 7, 4},
            {2, 4, 8, 9, 5, 7, 1, 3, 6},
            {7, 6, 3, 4, 1, 8, 2, 5, 9},
        };

        int[][] solved = Sudoku.solveSudoku(input);

        assertArrayEquals(expected, solved, "Sudoku puzzle was not solved correctly");
    }

    @Test
    public void testAlreadySolvedGridIsUnchanged() {
        int[][] input = {
            {5, 3, 4, 6, 7, 8, 9, 1, 2},
            {6, 7, 2, 1, 9, 5, 3, 4, 8},
            {1, 9, 8, 3, 4, 2, 5, 6, 7},
            {8, 5, 9, 7, 6, 1, 4, 2, 3},
            {4, 2, 6, 8, 5, 3, 7, 9, 1},
            {7, 1, 3, 9, 2, 4, 8, 5, 6},
            {9, 6, 1, 5, 3, 7, 2, 8, 4},
            {2, 8, 7, 4, 1, 9, 6, 3, 5},
            {3, 4, 5, 2, 8, 6, 1, 7, 9},
        };

        int[][] solved = Sudoku.solveSudoku(input);

        assertArrayEquals(input, solved, "An already-solved grid should be returned unchanged");
    }

    @Test
    public void testSolvedGridSatisfiesSudokuRules() {
        int[][] input = {
            {5, 3, 0, 0, 7, 0, 0, 0, 0},
            {6, 0, 0, 1, 9, 5, 0, 0, 0},
            {0, 9, 8, 0, 0, 0, 0, 6, 0},
            {8, 0, 0, 0, 6, 0, 0, 0, 3},
            {4, 0, 0, 8, 0, 3, 0, 0, 1},
            {7, 0, 0, 0, 2, 0, 0, 0, 6},
            {0, 6, 0, 0, 0, 0, 2, 8, 0},
            {0, 0, 0, 4, 1, 9, 0, 0, 5},
            {0, 0, 0, 0, 8, 0, 0, 7, 9},
        };

        int[][] solved = Sudoku.solveSudoku(input);

        assertTrue(isValidSolvedGrid(solved), "Solved grid violates Sudoku rules (row/column/box must each contain 1-9 exactly once)");
    }

    private static boolean isValidSolvedGrid(int[][] grid) {
        for (int i = 0; i < 9; i++) {
            if (!containsOneToNine(rowValues(grid, i))) return false;
            if (!containsOneToNine(columnValues(grid, i))) return false;
            if (!containsOneToNine(boxValues(grid, i))) return false;
        }
        return true;
    }

    private static int[] rowValues(int[][] grid, int row) {
        return grid[row].clone();
    }

    private static int[] columnValues(int[][] grid, int col) {
        int[] values = new int[9];
        for (int row = 0; row < 9; row++) {
            values[row] = grid[row][col];
        }
        return values;
    }

    private static int[] boxValues(int[][] grid, int boxIndex) {
        int startRow = (boxIndex / 3) * 3;
        int startCol = (boxIndex % 3) * 3;
        int[] values = new int[9];
        int k = 0;
        for (int r = startRow; r < startRow + 3; r++) {
            for (int c = startCol; c < startCol + 3; c++) {
                values[k++] = grid[r][c];
            }
        }
        return values;
    }

    private static boolean containsOneToNine(int[] values) {
        boolean[] seen = new boolean[10];
        for (int v : values) {
            if (v < 1 || v > 9 || seen[v]) return false;
            seen[v] = true;
        }
        return true;
    }
}
