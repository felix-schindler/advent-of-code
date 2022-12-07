const input = await Deno.readTextFile("./1.txt");

let calPerElf: number[] = [];

input.split("\n\n").map(elf => {
	let cal = 0;

	elf.split("\n").map(c => {
		cal += Number(c);
	});

	calPerElf.push(cal);
});

console.log("The highest calory count per elf is", calPerElf[calPerElf.length - 1]);

calPerElf = calPerElf.sort((a, b) => b - a);
if (calPerElf.length > 3) {
	let top3 = 0;
	for (let i = 0; i < 3; i++)
		top3 += calPerElf[i];
	console.log("The top 3 elfes carry ", top3, "calories");
} else console.error("There have to be at least 3 elfes for part two of the exercise");
