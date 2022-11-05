// https://adventofcode.com/2020/day/3

const decoder = new TextDecoder("utf-8");
const data = await Deno.readFile("3.txt");
const inputStr = decoder.decode(data);

let treeCount = 0;
let count = 0;

const STEP_RIGHT = 1;
let line = 0;

inputStr.split("\n").map((val: string) => {
	if (++line & 1) {
		if (val != "") {
			let currStr = val;
			const pos = count++ * STEP_RIGHT;

			while (pos > currStr.length)
				currStr = currStr + currStr;

			if (currStr.charAt(pos) == "#")
				treeCount++;
		}
	}
});

console.log("There are %c" + treeCount + "%c trees", "color: green", "color: inherit");
console.log("Answer: %c" + (75 * 178 * 75 * 86 * 37), "color: green");
