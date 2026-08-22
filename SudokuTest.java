import java.util.Arrays;

/**
 * Dependency-free test harness for {@link Sudoku}, mirroring the Go test in this repo and adding
 * the edge cases it does not cover. Run with {@code java SudokuTest.java} (JDK 22+ compiles the
 * sibling source file automatically). Exits non-zero if any check fails.
 */
public final class SudokuTest {

    private static int failures = 0;

    public static void main(String[] args) {
        testReadmeExample();
        testInputNotMutated();
        testEmptyGridIsSolvable();
        testHardPuzzle();
        testUnsolvablePuzzleReturnsNull();
        testMalformedInputReturnsNull();
        testContradictoryInputReturnsNull();

        if (failures == 0) {
            System.out.println("\nPASS - all checks passed");
        } else {
            System.out.printf("%nFAIL - %d check(s) failed%n", failures);
            System.exit(1);
        }
    }

    /** The example from the README, identical to the Go test's fixture. */
    private static void testReadmeExample() {
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
        if (!Arrays.deepEquals(solved, expected)) {
            fail("readme example", "expected:%n%s%n%ngot:%n%s"
                    .formatted(Sudoku.format(expected), solved == null ? "null" : Sudoku.format(solved)));
            return;
        }
        pass("readme example matches the documented solution");
    }

    /** Solving must not write through to the caller's array. */
    private static void testInputNotMutated() {
        int[][] input = puzzleWithOneBlank();
        int[][] copy = deepCopy(input);

        Sudoku.solveSudoku(input);

        if (!Arrays.deepEquals(input, copy)) {
            fail("input not mutated", "the caller's grid was modified in place");
            return;
        }
        pass("input grid is left untouched");
    }

    /**
     * A completely empty board: 81 blanks, not a single clue. It has roughly 6.7e21 valid
     * solutions, so no particular grid can be asserted — instead every cell must come back filled
     * with a digit 1-9, the result must obey all three Sudoku rules, and the caller's blank grid
     * must be returned untouched. This is also the worst case for the search: with no clues to
     * propagate from, the minimum-remaining-values heuristic gets no head start.
     */
    private static void testEmptyGridIsSolvable() {
        int[][] input = {
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
        };

        long start = System.nanoTime();
        int[][] solved = Sudoku.solveSudoku(input);
        long millis = (System.nanoTime() - start) / 1_000_000;

        if (solved == null) {
            fail("empty board", "returned null for a board with no clues");
            return;
        }
        // validate() also rejects a leftover 0, but says so as an out-of-range digit; an explicit
        // blank check names the real problem when the search stops early.
        for (int r = 0; r < 9; r++) {
            for (int c = 0; c < 9; c++) {
                if (solved[r][c] == 0) {
                    fail("empty board", "cell (%d,%d) was left blank".formatted(r, c));
                    return;
                }
            }
        }
        String problem = validate(solved);
        if (problem != null) {
            fail("empty board", problem);
            return;
        }
        if (!Arrays.deepEquals(input, new int[9][9])) {
            fail("empty board", "the caller's blank grid was modified in place");
            return;
        }
        pass("empty board (81 blanks) filled validly in " + millis + " ms");
    }

    /**
     * "AI Escargot"-class puzzle designed to defeat naive left-to-right backtracking; a check
     * that the minimum-remaining-values heuristic is actually doing its job.
     */
    private static void testHardPuzzle() {
        int[][] input = {
            {8, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 3, 6, 0, 0, 0, 0, 0},
            {0, 7, 0, 0, 9, 0, 2, 0, 0},
            {0, 5, 0, 0, 0, 7, 0, 0, 0},
            {0, 0, 0, 0, 4, 5, 7, 0, 0},
            {0, 0, 0, 1, 0, 0, 0, 3, 0},
            {0, 0, 1, 0, 0, 0, 0, 6, 8},
            {0, 0, 8, 5, 0, 0, 0, 1, 0},
            {0, 9, 0, 0, 0, 0, 4, 0, 0},
        };

        long start = System.nanoTime();
        int[][] solved = Sudoku.solveSudoku(input);
        long millis = (System.nanoTime() - start) / 1_000_000;

        if (solved == null) {
            fail("hard puzzle", "returned null for a solvable puzzle");
            return;
        }
        String problem = validate(solved);
        if (problem != null) {
            fail("hard puzzle", problem);
            return;
        }
        if (!preservesClues(input, solved)) {
            fail("hard puzzle", "solution overwrote one of the given clues");
            return;
        }
        pass("hard puzzle solved in " + millis + " ms");
    }

    /** Well-formed and rule-abiding, but no completion exists. */
    private static void testUnsolvablePuzzleReturnsNull() {
        int[][] input = {
            {0, 0, 0, 0, 0, 5, 0, 8, 0},
            {0, 0, 0, 6, 0, 1, 0, 4, 3},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 1, 0, 5, 0, 0, 0, 0, 0},
            {0, 0, 0, 1, 0, 6, 0, 0, 0},
            {3, 0, 0, 0, 0, 0, 0, 0, 5},
            {5, 3, 0, 0, 0, 0, 0, 6, 1},
            {0, 0, 0, 0, 0, 0, 0, 0, 4},
            {0, 0, 0, 0, 0, 0, 0, 0, 0},
        };

        if (Sudoku.solveSudoku(input) != null) {
            fail("unsolvable puzzle", "expected null, got a grid");
            return;
        }
        pass("unsolvable puzzle returns null");
    }

    /** Wrong shapes and out-of-range digits are rejected rather than crashing. */
    private static void testMalformedInputReturnsNull() {
        record Case(String name, int[][] grid) {}

        int[][] shortRow = new int[9][9];
        shortRow[4] = new int[8];

        int[][] outOfRange = new int[9][9];
        outOfRange[0][0] = 10;

        int[][] negative = new int[9][9];
        negative[8][8] = -1;

        Case[] cases = {
            new Case("null grid", null),
            new Case("too few rows", new int[8][9]),
            new Case("too many rows", new int[10][9]),
            new Case("ragged row", shortRow),
            new Case("digit above 9", outOfRange),
            new Case("negative digit", negative),
        };

        for (Case c : cases) {
            if (Sudoku.solveSudoku(c.grid()) != null) {
                fail("malformed input", c.name() + ": expected null");
                return;
            }
        }
        pass("malformed input returns null (" + cases.length + " cases)");
    }

    /** A grid that already violates a Sudoku rule is rejected up front. */
    private static void testContradictoryInputReturnsNull() {
        int[][] duplicateRow = new int[9][9];
        duplicateRow[0][0] = 5;
        duplicateRow[0][8] = 5;

        int[][] duplicateCol = new int[9][9];
        duplicateCol[0][3] = 7;
        duplicateCol[8][3] = 7;

        int[][] duplicateBox = new int[9][9];
        duplicateBox[4][4] = 2;
        duplicateBox[5][5] = 2;

        if (Sudoku.solveSudoku(duplicateRow) != null
                || Sudoku.solveSudoku(duplicateCol) != null
                || Sudoku.solveSudoku(duplicateBox) != null) {
            fail("contradictory input", "expected null for a duplicate in a row, column or box");
            return;
        }
        pass("duplicate clues in a row, column or box return null");
    }

    // --- helpers ---------------------------------------------------------

    /** Returns a description of the first rule broken by grid, or null if it is a valid solution. */
    private static String validate(int[][] grid) {
        for (int i = 0; i < 9; i++) {
            int rowSeen = 0;
            int colSeen = 0;
            int boxSeen = 0;
            for (int j = 0; j < 9; j++) {
                int rowDigit = grid[i][j];
                int colDigit = grid[j][i];
                // Walk the i-th box cell by cell.
                int boxDigit = grid[(i / 3) * 3 + j / 3][(i % 3) * 3 + j % 3];

                if (rowDigit < 1 || rowDigit > 9) {
                    return "cell (%d,%d) holds %d, not a digit 1-9".formatted(i, j, rowDigit);
                }
                if ((rowSeen & (1 << rowDigit)) != 0) {
                    return "row %d repeats %d".formatted(i, rowDigit);
                }
                if ((colSeen & (1 << colDigit)) != 0) {
                    return "column %d repeats %d".formatted(i, colDigit);
                }
                if ((boxSeen & (1 << boxDigit)) != 0) {
                    return "box %d repeats %d".formatted(i, boxDigit);
                }
                rowSeen |= 1 << rowDigit;
                colSeen |= 1 << colDigit;
                boxSeen |= 1 << boxDigit;
            }
        }
        return null;
    }

    /** Reports whether every non-zero clue in puzzle survived into solved. */
    private static boolean preservesClues(int[][] puzzle, int[][] solved) {
        for (int r = 0; r < 9; r++) {
            for (int c = 0; c < 9; c++) {
                if (puzzle[r][c] != 0 && puzzle[r][c] != solved[r][c]) {
                    return false;
                }
            }
        }
        return true;
    }

    private static int[][] puzzleWithOneBlank() {
        int[][] grid = {
            {5, 3, 4, 6, 7, 8, 9, 1, 2},
            {6, 7, 2, 1, 9, 5, 3, 4, 8},
            {1, 9, 8, 3, 4, 2, 5, 6, 7},
            {8, 5, 9, 7, 6, 1, 4, 2, 3},
            {4, 2, 6, 8, 5, 3, 7, 9, 1},
            {7, 1, 3, 9, 2, 4, 8, 5, 6},
            {9, 6, 1, 5, 3, 7, 2, 8, 4},
            {2, 8, 7, 4, 1, 9, 6, 3, 5},
            {3, 4, 5, 2, 8, 6, 1, 7, 0},
        };
        return grid;
    }

    private static int[][] deepCopy(int[][] grid) {
        int[][] copy = new int[grid.length][];
        for (int r = 0; r < grid.length; r++) {
            copy[r] = grid[r].clone();
        }
        return copy;
    }

    private static void pass(String what) {
        System.out.println("ok   - " + what);
    }

    private static void fail(String what, String detail) {
        failures++;
        System.out.printf("FAIL - %s: %s%n", what, detail);
    }
}
