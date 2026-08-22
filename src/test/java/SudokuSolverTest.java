import java.util.Arrays;

public final class SudokuSolverTest {
    private static int testsRun;

    public static void main(String[] args) {
        run("solves example puzzle", SudokuSolverTest::solvesExamplePuzzle);
        run("accepts solved puzzle", SudokuSolverTest::acceptsSolvedPuzzle);
        run("rejects malformed grid", SudokuSolverTest::rejectsMalformedGrid);
        run("rejects conflicting clues", SudokuSolverTest::rejectsConflictingClues);
        run("rejects unsolvable puzzle", SudokuSolverTest::rejectsUnsolvablePuzzle);
        System.out.println("All " + testsRun + " tests passed.");
    }

    private static void solvesExamplePuzzle() {
        int[][] puzzle = {
                {5, 3, 0, 0, 7, 0, 0, 0, 0},
                {6, 0, 0, 1, 9, 5, 0, 0, 0},
                {0, 9, 8, 0, 0, 0, 0, 6, 0},
                {8, 0, 0, 0, 6, 0, 0, 0, 3},
                {4, 0, 0, 8, 0, 3, 0, 0, 1},
                {7, 0, 0, 0, 2, 0, 0, 0, 6},
                {0, 6, 0, 0, 0, 0, 2, 8, 0},
                {0, 0, 0, 4, 1, 9, 0, 0, 5},
                {0, 0, 0, 0, 8, 0, 0, 7, 9}
        };
        int[][] expected = solvedGrid();

        assertGridEquals(expected, SudokuSolver.solveSudoku(puzzle));
    }

    private static void acceptsSolvedPuzzle() {
        int[][] solved = solvedGrid();
        assertGridEquals(solvedGrid(), SudokuSolver.solveSudoku(solved));
    }

    private static void rejectsMalformedGrid() {
        assertThrows(() -> SudokuSolver.solveSudoku(new int[8][9]));
    }

    private static void rejectsConflictingClues() {
        int[][] puzzle = new int[9][9];
        puzzle[0][0] = 5;
        puzzle[0][1] = 5;
        assertThrows(() -> SudokuSolver.solveSudoku(puzzle));
    }

    private static void rejectsUnsolvablePuzzle() {
        int[][] puzzle = {
                {1, 2, 3, 4, 5, 6, 7, 8, 0},
                {0, 0, 0, 0, 0, 0, 0, 0, 9},
                {0, 0, 0, 0, 0, 0, 0, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 0, 0}
        };
        assertThrows(() -> SudokuSolver.solveSudoku(puzzle));
    }

    private static int[][] solvedGrid() {
        return new int[][] {
                {5, 3, 4, 6, 7, 8, 9, 1, 2},
                {6, 7, 2, 1, 9, 5, 3, 4, 8},
                {1, 9, 8, 3, 4, 2, 5, 6, 7},
                {8, 5, 9, 7, 6, 1, 4, 2, 3},
                {4, 2, 6, 8, 5, 3, 7, 9, 1},
                {7, 1, 3, 9, 2, 4, 8, 5, 6},
                {9, 6, 1, 5, 3, 7, 2, 8, 4},
                {2, 8, 7, 4, 1, 9, 6, 3, 5},
                {3, 4, 5, 2, 8, 6, 1, 7, 9}
        };
    }

    private static void assertGridEquals(int[][] expected, int[][] actual) {
        if (!Arrays.deepEquals(expected, actual)) {
            throw new AssertionError(
                    "Expected " + Arrays.deepToString(expected)
                            + " but got " + Arrays.deepToString(actual));
        }
    }

    private static void assertThrows(Runnable action) {
        try {
            action.run();
        } catch (IllegalArgumentException expected) {
            return;
        }
        throw new AssertionError("Expected IllegalArgumentException");
    }

    private static void run(String name, Runnable test) {
        test.run();
        testsRun++;
        System.out.println("PASS: " + name);
    }
}
