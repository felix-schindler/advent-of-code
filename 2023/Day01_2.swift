import Foundation

// Get file path
let cwd = URL(fileURLWithPath: FileManager.default.currentDirectoryPath, isDirectory: true)
let fileName = "Day01_2.txt"
let fileURL = cwd.appendingPathComponent(fileName)

let numbersInWords = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
let wordToNumber = [
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9
]

func checkIfStringStartsWithNumber(_ someStr: String, reversed: Bool = false) -> Int {
	if let number = Int(String(reversed ? someStr.last! : someStr.first!)) {
		return number
	} else {
		for numberWord in numbersInWords {
			if (reversed ? someStr.hasSuffix(numberWord) : someStr.hasPrefix(numberWord)) {
				return wordToNumber[numberWord] ?? -1
			}
		}
	}
	
	let newSubString = (reversed ? someStr.dropLast() : someStr.dropFirst())
	return checkIfStringStartsWithNumber(String(newSubString), reversed: reversed)
}

// Read text file
do {
	let text = try String(contentsOf: fileURL, encoding: .utf8)
	let split = text.split(separator: "\n")
	var result = 0
	
	for line in split {
		var firstNumber = 0
		var lastNumber = 0

		firstNumber = checkIfStringStartsWithNumber(String(line))
		lastNumber = checkIfStringStartsWithNumber(String(line), reversed: true)

		if (firstNumber == -1 || lastNumber == -1) {
			print("Didn't find a number in the following line: " + line)
		}
		
		let newNumberString: String = "\(firstNumber)\(lastNumber)"
		if let newNumber = Int(newNumberString) {
			result += newNumber
		} else {
			print("Failed to convert string into number")
		}
	}
	
	print(result)
} catch {
	print("Failed to read contents of file")
}
