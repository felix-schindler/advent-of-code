import { fileToString } from "../utils.ts";

const reqFields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", /*"cid"*/];
const inputStr = await fileToString("4.txt");
const passports = inputStr.split("\n \n");

let count = 0;

passports.forEach((pass) => {
	pass = pass.replaceAll("\n", " ");

	const fields = pass.split(" ");
	const fieldNames: string[] = [];
	let valid = true;

	fields.map((f) => {
		if (!valid)
			return;
		const field = f.split(":")[0];
		const value = f.split(":")[1];
		fieldNames.push(field);

		let temp: any;

		switch (field) {
			case "byr":
				temp = Number(value);
				valid = (temp >= 1920 && temp <= 2002);
				break;
			case "iyr":
				temp = Number(value);
				valid = (temp >= 2010 && temp <= 2020);
				break;
			case "eyr":
				temp = Number(value);
				valid = (temp >= 2020 && temp <= 2030);
				break;
			case "hgt":
				if (value.includes("cm")) {
					temp = Number(value.split("c")[0]);
					valid = (temp >= 150 && temp <= 193);
				} else if (value.includes("in")) {
					temp = Number(value.split("i")[0]);
					valid = (temp >= 59 && temp <= 76);
				} else {
					valid = false;
				}
				break;
			case "hcl":
				valid = /^#[0-9a-f]{6}$/.test(value);
				break;
			case "ecl":
				valid = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].includes(value);
				break;
			case "pid":
				valid = value.length == 9;
				break;
			default:
				break;
		}
	});

	if (valid) {
		const allFields = fieldNames.join("_");
		valid = reqFields.every(item => allFields.includes(item));
	}

	if (valid)
		count++;
});

console.log("%c" + count + "%c of %c" + passports.length + "%c passports are valid", "color: green", "color: inherit", "color: yellow", "color: inherit");
