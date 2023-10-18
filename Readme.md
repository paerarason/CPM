## This is the repo that manages to solve few Problems based on Operatinal Research in GOlang 
   - CPM 
   - Assignment Model 
   - Transpotation Model 


### Transpotation Model
   The Hungarian Algorithm is a method for solving the Assignment Problem, which is a specific case of linear sum assignment problem. The goal is to find the optimal assignment of a set of tasks to a set of agents, where each agent has a cost associated with performing each task, and the objective is to minimize the total cost. The Hungarian Algorithm can handle various scenarios, including exception cases. Here are the steps for the Hungarian Algorithm:

   #### 1. Create the Cost Matrix:
      - Create an n x n cost matrix, where n is the number of agents (rows) and the number of tasks     (columns).
      - Populate the matrix with the cost associated with each agent-task pair.

   #### 2. Subtract Row Minima:
      - For each row, subtract the minimum value in that row from every element in that row. This step ensures that at least one zero appears in every row.

   #### 3. Subtract Column Minima:
      - For each column, subtract the minimum value in that column from every element in that column. This step ensures that at least one zero appears in every column.

   #### 4. Mark Zeros with Lines:
      - Starting with the first row, draw lines to cover as many zeros as possible while minimizing the total number of lines required. The lines can be horizontal or vertical, and you need to use them strategically to maximize the coverage.

   #### 5. Check for Optimality:
      - Check if you have exactly n lines (n is the number of rows or columns). If you do, go to step 7. If not, proceed to step 6.

   #### 6. Find Minimum Uncovered Element:
      - Find the smallest uncovered element in the cost matrix, denoted as "minval."

   #### 7. Modify the Matrix:
      - Subtract minval from all uncovered elements and add it to the elements at the intersections of the lines (double-covered elements). This step reduces the matrix to one with more zeros.

   #### 8. Cover and Uncover Rows and Columns:
      - Cover every row that contains a zero after the modifications in step 7.
      - Uncover every column.
      - Go back to step 5 to check for optimality again.

   #### 9. Identify the Optimal Assignment:
      - Once you have exactly n lines covering rows and columns and the matrix is fully optimized, you can determine the optimal assignment. The assignment is based on the original cost matrix before any modifications.

   #### 10. Handling Exception Cases:
      - Exception cases may arise when the algorithm encounters difficulties in reaching optimality due to a tied situation, degenerate cases, or non-square matrices. In such situations, you may need to apply heuristics or additional rules to resolve the problem.

   #### 11. Repeat the Process (Optional):
      - In some cases, you may want to explore multiple solutions or find all possible optimal assignments. You can repeat the algorithm by relaxing some constraints or exploring alternative solutions.

   #### 12. Terminate:
      - Continue the algorithm until you reach an optimal solution or determine that no further solutions are possible.
