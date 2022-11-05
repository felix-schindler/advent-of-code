// https://adventofcode.com/2020/day/1

let inputStr = "";
const intArr: number[] = [];

if (Deno.args.length > 0) {
	inputStr = Deno.args.join("\n");
} else {
	const decoder = new TextDecoder("utf-8");
	const data = await Deno.readFile("1.txt");
	inputStr = decoder.decode(data);
}

inputStr.split("\n").map((val: string) => {
	if (val != "")
		intArr.push(Number(val));
});

console.log("There are " + intArr.length + " input numbers");

for (const val1 of intArr)
	for (const val2 of intArr)
		for (const val3 of intArr)
			if (val1 + val2 + val3 == 2020) {
				console.log(val1, val2, val3, " ", val1 * val2 * val3);
				Deno.exit();
			}
