import { fileToString } from "../utils.ts";

let count = 0;

(await fileToString("7.txt")).split("\n").map((line) => {
	console.log(line);
	count++;
});

console.log("There are %c" + count + "%c shiny gold bags", "color: green", "color: inherit");
