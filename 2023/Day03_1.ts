import { join } from "https://deno.land/std@0.208.0/path/join.ts";

const cwd = Deno.cwd();
const input = await Deno.readTextFile(join(cwd, "Day03_1.txt"));

const lines = input.split("\n")
const linesLength = lines.length;
for (let i = 0; i < linesLength; i++) {
	const line = lines[i];

	for (let join = 0; join < line.length; join++) {
		const letter = line.charAt(join);

		if (letter != ".") {
			if (isNaN(Number(letter))) {
				// it's a symbol
			} else {
				// it's a number
			}
		}
	}
}

const isPart = checkSurroundingFields();
