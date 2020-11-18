package leetcode

/*
 * @lc app=leetcode id=874 lang=golang
 *
 * [874] Walking Robot Simulation
 *
 * https://leetcode.com/problems/walking-robot-simulation/description/
 *
 * algorithms
 * Easy (35.48%)
 * Likes:    201
 * Dislikes: 882
 * Total Accepted:    21.8K
 * Total Submissions: 59.8K
 * Testcase Example:  '[4,-1,3]\n[]'
 *
 * A robot on an infinite grid starts at point (0, 0) and faces north.  The
 * robot can receive one of three possible types of commands:
 *
 *
 * -2: turn left 90 degrees
 * -1: turn right 90 degrees
 * 1 <= x <= 9: move forward x units
 *
 *
 * Some of the grid squares are obstacles.
 *
 * The i-th obstacle is at grid point (obstacles[i][0], obstacles[i][1])
 *
 * If the robot would try to move onto them, the robot stays on the previous
 * grid square instead (but still continues following the rest of the route.)
 *
 * Return the square of the maximum Euclidean distance that the robot will be
 * from the origin.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: commands = [4,-1,3], obstacles = []
 * Output: 25
 * Explanation: robot will go to (3, 4)
 *
 *
 *
 * Example 2:
 *
 *
 * Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
 * Output: 65
 * Explanation: robot will be stuck at (1, 4) before turning left and going to
 * (1, 8)
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= commands.length <= 10000
 * 0 <= obstacles.length <= 10000
 * -30000 <= obstacle[i][0] <= 30000
 * -30000 <= obstacle[i][1] <= 30000
 * The answer is guaranteed to be less than 2 ^ 31.
 *
 *
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	// init obstacles
	obstaclesPoint = map[point]bool{}
	for i := 0; i < len(obstacles); i++ {
		obstaclesPoint[point{x: obstacles[i][0], y: obstacles[i][1]}] = true
	}
	// loop commands
	x, y := 0, 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	direct := 0
	res := 0
	for i := 0; i < len(commands); i++ {
		if commands[i] == -1 {
			direct = (direct + 1) % 4
		} else if commands[i] == -2 {
			direct = (direct + 3) % 4
		} else {
			// caculate newx and newy
			tmpX, tmpY := x, y
			dirX, dirY := directions[direct][0], directions[direct][1]
			for j := 1; j <= commands[i]; j++ {
				newPoint := point{x: tmpX + dirX, y: tmpY + dirY}
				if _, ok := obstaclesPoint[newPoint]; ok {
					break
				} else {
					tmpX, tmpY = newPoint.x, newPoint.y
				}
			}
			x, y = tmpX, tmpY
			if res < x*x+y*y {
				res = x*x + y*y
			}
		}
	}
	return res
}

type point struct {
	x int
	y int
}

var obstaclesPoint map[point]bool

// @lc code=end
