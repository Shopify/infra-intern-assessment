import java.util.StringJoiner;

/**
 * Sudoku solver.
 *
 * <p>The digits already used in each row, column and 3x3 box are tracked as 9-bit masks. That
 * reduces "is this placement legal?" to a single bitwise AND, and makes counting a cell's
 * candidates cheap enough to drive a minimum-remaining-values heuristic: the most constrained
 * empty cell is always filled first, which prunes far more of the search tree than scanning
 * cells in reading order.
 */
public final class Sudoku {

    /** Side length of the grid. */
    private static final int SIZE = 9;
    /** Side length of a sub-grid. */
    private static final int BOX_SIZE = 3;
    /** Bits 1..9 set; bit d means "digit d is taken". */
    private static final int FULL = 0b1111111110;

    private final int[][] grid = new int[SIZE][SIZE];
    private final int[] rows = new int[SIZE];
    private final int[] cols = new int[SIZE];
    private final int[] boxes = new int[SIZE];

    /** Cells still to fill. */
    private int empty = SIZE * SIZE;

    private Sudoku() {}

    /**
     * Solves the given 9x9 puzzle, in which 0 marks an empty cell, and returns the completed
     * grid. The input grid is left untouched.
     *
     * @return the solved grid, or null if the input is not a well-formed 9x9 puzzle or has no
     *     solution
     */
    public static int[][] solveSudoku(int[][] puzzle) {
        Sudoku state = load(puzzle);
        if (state == null || !state.solve()) {
            return null;
        }

        int[][] solved = new int[SIZE][];
        for (int r = 0; r < SIZE; r++) {
            solved[r] = state.grid[r].clone();
        }
        return solved;
    }

    /**
     * Loads a puzzle into a fresh state, returning null if the grid is not 9x9, holds a value
     * outside 0..9, or already breaks a Sudoku rule.
     */
    private static Sudoku load(int[][] puzzle) {
        if (puzzle == null || puzzle.length != SIZE) {
            return null;
        }

        Sudoku state = new Sudoku();
        for (int r = 0; r < SIZE; r++) {
            if (puzzle[r] == null || puzzle[r].length != SIZE) {
                return null;
            }
            for (int c = 0; c < SIZE; c++) {
                int d = puzzle[r][c];
                if (d == 0) {
                    continue;
                }
                if (d < 1 || d > SIZE) {
                    return null;
                }
                if ((state.candidates(r, c) & (1 << d)) == 0) {
                    return null; // digit repeats in its row, column or box
                }
                state.place(r, c, d);
            }
        }
        return state;
    }

    /** Maps a cell coordinate to the index of its 3x3 sub-grid. */
    private static int boxIndex(int r, int c) {
        return (r / BOX_SIZE) * BOX_SIZE + c / BOX_SIZE;
    }

    /** Writes digit d into (r, c) and marks it used in all three masks. */
    private void place(int r, int c, int d) {
        int bit = 1 << d;
        grid[r][c] = d;
        rows[r] |= bit;
        cols[c] |= bit;
        boxes[boxIndex(r, c)] |= bit;
        empty--;
    }

    /** Undoes {@link #place}, emptying the cell again. */
    private void remove(int r, int c, int d) {
        int bit = 1 << d;
        grid[r][c] = 0;
        rows[r] &= ~bit;
        cols[c] &= ~bit;
        boxes[boxIndex(r, c)] &= ~bit;
        empty++;
    }

    /** Returns the mask of digits that may legally go in (r, c). */
    private int candidates(int r, int c) {
        return FULL & ~(rows[r] | cols[c] | boxes[boxIndex(r, c)]);
    }

    /** Fills every remaining cell by backtracking, reporting whether the puzzle completed. */
    private boolean solve() {
        if (empty == 0) {
            return true;
        }

        // Minimum remaining values: branch on the empty cell with the fewest legal digits. One
        // candidate means the cell is forced, so this doubles as constraint propagation; zero
        // means this branch is already dead.
        int bestRow = -1;
        int bestCol = -1;
        int bestMask = 0;
        int bestCount = SIZE + 1;
        outer:
        for (int r = 0; r < SIZE; r++) {
            for (int c = 0; c < SIZE; c++) {
                if (grid[r][c] != 0) {
                    continue;
                }
                int mask = candidates(r, c);
                int count = Integer.bitCount(mask);
                if (count < bestCount) {
                    bestRow = r;
                    bestCol = c;
                    bestMask = mask;
                    bestCount = count;
                    if (count <= 1) {
                        break outer; // no other cell can be a better branch point
                    }
                }
            }
        }

        if (bestCount == 0) {
            return false; // some empty cell can take no digit at all
        }

        // Iterate the set bits of bestMask, lowest first; mask &= mask - 1 clears it.
        for (int mask = bestMask; mask != 0; mask &= mask - 1) {
            int d = Integer.numberOfTrailingZeros(mask);
            place(bestRow, bestCol, d);
            if (solve()) {
                return true;
            }
            remove(bestRow, bestCol, d);
        }
        return false;
    }

    /** Renders a grid in the bracketed form used by the README. */
    public static String format(int[][] grid) {
        StringJoiner rows = new StringJoiner(",\n  ", "[\n  ", "\n]");
        for (int[] row : grid) {
            StringJoiner cells = new StringJoiner(", ", "[", "]");
            for (int d : row) {
                cells.add(Integer.toString(d));
            }
            rows.add(cells.toString());
        }
        return rows.toString();
    }

    public static void main(String[] args) {
        int[][] puzzle = {
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

        int[][] solved = solveSudoku(puzzle);
        System.out.println(solved == null ? "puzzle has no solution" : format(solved));
    }
}
