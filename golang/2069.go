//go:build ignore

package main

// A width x height grid is on an XY-plane with the bottom-left cell at (0, 0) and the top-right cell at (width - 1, height - 1). The grid is aligned with the four cardinal directions ("North", "East", "South", and "West"). A robot is initially at cell (0, 0) facing direction "East".
//
// The robot can be instructed to move for a specific number of steps. For each step, it does the following.
//
// Attempts to move forward one cell in the direction it is facing.
// If the cell the robot is moving to is out of bounds, the robot instead turns 90 degrees counterclockwise and retries the step.
// After the robot finishes moving the number of steps required, it stops and awaits the next instruction.
//
// Implement the Robot class:
//
// Robot(int width, int height) Initializes the width x height grid with the robot at (0, 0) facing "East".
// void step(int num) Instructs the robot to move forward num steps.
// int[] getPos() Returns the current cell the robot is at, as an array of length 2, [x, y].
// String getDir() Returns the current direction of the robot, "North", "East", "South", or "West".
//
//
// Example 1:
//
// example-1
// Input
// ["Robot", "step", "step", "getPos", "getDir", "step", "step", "step", "getPos", "getDir"]
// [[6, 3], [2], [2], [], [], [2], [1], [4], [], []]
// Output
// [null, null, null, [4, 0], "East", null, null, null, [1, 2], "West"]
//
// Explanation
// Robot robot = new Robot(6, 3); // Initialize the grid and the robot at (0, 0) facing East.
// robot.step(2);  // It moves two steps East to (2, 0), and faces East.
// robot.step(2);  // It moves two steps East to (4, 0), and faces East.
// robot.getPos(); // return [4, 0]
// robot.getDir(); // return "East"
// robot.step(2);  // It moves one step East to (5, 0), and faces East.
//                 // Moving the next step East would be out of bounds, so it turns and faces North.
//                 // Then, it moves one step North to (5, 1), and faces North.
// robot.step(1);  // It moves one step North to (5, 2), and faces North (not West).
// robot.step(4);  // Moving the next step North would be out of bounds, so it turns and faces West.
//                 // Then, it moves four steps West to (1, 2), and faces West.
// robot.getPos(); // return [1, 2]
// robot.getDir(); // return "West"
//
//
//
// Constraints:
//
// 2 <= width, height <= 100
// 1 <= num <= 105
// At most 104 calls in total will be made to step, getPos, and getDir.

type Robot struct {
	x         int
	y         int
	d         int
	direction [][2]int
	w         int
	h         int
}

func Constructor(width int, height int) Robot {
	return Robot{
		x:         0,
		y:         0,
		d:         0,
		direction: [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}, // E, N, W, S
		w:         width,
		h:         height,
	}
}

func (this *Robot) Step(num int) {
	chuvi := ((this.w + this.h - 2) * 2)
	num %= chuvi
	if num == 0 {
		num = chuvi
	}
	for range num {
		for this.direction[this.d%4][0]+this.x >= this.w || this.direction[this.d%4][1]+this.y >= this.h || this.direction[this.d%4][0]+this.x < 0 || this.direction[this.d%4][1]+this.y < 0 {
			this.d += 1
		}

		this.x = this.direction[this.d%4][0] + this.x
		this.y = this.direction[this.d%4][1] + this.y
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
	if this.d%4 == 0 {
		return "East"
	} else if this.d%4 == 1 {
		return "North"
	} else if this.d%4 == 2 {
		return "West"
	}
	return "South"
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */

func main() {
}
