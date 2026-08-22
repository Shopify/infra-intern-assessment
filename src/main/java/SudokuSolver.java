public final class SudokuSolver {
    private static final int SIZE = 9;
    private static final int ALL_DIGITS = (1 << SIZE) - 1;

    private SudokuSolver() {
    }

    /**
     * Solves the supplied Sudoku grid in place and returns the same grid.
     *
     * @throws IllegalArgumentException if the grid is malformed, contains
     *         conflicting clues, or has no solution
     */
    public static int[][] solveSudoku(int[][] grid) {
        validateShape(grid);

        int[] rows = new int[SIZE];
        int[] columns = new int[SIZE];
        int[] boxes = new int[SIZE];

        for (int row = 0; row < SIZE; row++) {
            for (int column = 0; column < SIZE; column++) {
                int value = grid[row][column];
                if (value == 0) {
                    continue;
                }
                if (value < 1 || value > SIZE) {
                    throw new IllegalArgumentException("Cell values must be between 0 and 9");
                }

                int bit = 1 << (value - 1);
                int box = boxIndex(row, column);
                if ((rows[row] & bit) != 0
                        || (columns[column] & bit) != 0
                        || (boxes[box] & bit) != 0) {
                    throw new IllegalArgumentException("Grid contains conflicting clues");
                }
                rows[row] |= bit;
                columns[column] |= bit;
                boxes[box] |= bit;
            }
        }

        if (!solve(grid, rows, columns, boxes)) {
            throw new IllegalArgumentException("Grid has no solution");
        }
        return grid;
    }

    private static boolean solve(int[][] grid, int[] rows, int[] columns, int[] boxes) {
        int bestRow = -1;
        int bestColumn = -1;
        int bestCandidates = 0;
        int fewestCandidates = SIZE + 1;

        for (int row = 0; row < SIZE; row++) {
            for (int column = 0; column < SIZE; column++) {
                if (grid[row][column] != 0) {
                    continue;
                }

                int candidates = ALL_DIGITS
                        & ~(rows[row] | columns[column] | boxes[boxIndex(row, column)]);
                int candidateCount = Integer.bitCount(candidates);
                if (candidateCount == 0) {
                    return false;
                }
                if (candidateCount < fewestCandidates) {
                    bestRow = row;
                    bestColumn = column;
                    bestCandidates = candidates;
                    fewestCandidates = candidateCount;
                }
            }
        }

        if (bestRow == -1) {
            return true;
        }

        int box = boxIndex(bestRow, bestColumn);
        for (int candidates = bestCandidates; candidates != 0; candidates &= candidates - 1) {
            int bit = candidates & -candidates;
            grid[bestRow][bestColumn] = Integer.numberOfTrailingZeros(bit) + 1;
            rows[bestRow] |= bit;
            columns[bestColumn] |= bit;
            boxes[box] |= bit;

            if (solve(grid, rows, columns, boxes)) {
                return true;
            }

            grid[bestRow][bestColumn] = 0;
            rows[bestRow] &= ~bit;
            columns[bestColumn] &= ~bit;
            boxes[box] &= ~bit;
        }
        return false;
    }

    private static void validateShape(int[][] grid) {
        if (grid == null || grid.length != SIZE) {
            throw new IllegalArgumentException("Grid must contain exactly 9 rows");
        }
        for (int[] row : grid) {
            if (row == null || row.length != SIZE) {
                throw new IllegalArgumentException("Each row must contain exactly 9 cells");
            }
        }
    }

    private static int boxIndex(int row, int column) {
        return row / 3 * 3 + column / 3;
    }
}
