import { fileToString } from "../utils.ts";

let count = 0;

(await fileToString("7.txt")).split("\n").map((line) => {
	if (line.includes("shiny gold")) {
		if (line.startsWith("shiny gold"))
			count++;
		else {
			line.split(", ").map((l) => {
				if (l.includes("shiny gold")) {
					const temp = Number(l.charAt(0));
					if (!isNaN(temp))
						count += temp;
					else {
						console.warn(l.charAt(0), "is not a number!", l);

						l.split("contain ").map((l) => {
							if (l.includes("shiny gold")) {
								const temp = Number(l.charAt(0));
								if (!isNaN(temp))
									count += temp;
								else console.warn(l.charAt(0), "is not a number!", l);
							}
						});

						l.split("contains ").map((l) => {
							if (l.includes("shiny gold")) {
								const temp = Number(l.charAt(0));
								if (!isNaN(temp))
									count += temp;
								else console.warn(l.charAt(0), "is not a number!", l);
							}
						});
					}
				}
			});
		}
	}
});

console.log("There are %c" + count + "%c shiny gold bags", "color: green", "color: inherit");
