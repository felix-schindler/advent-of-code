import Foundation

// Get file path
let cwd = URL(fileURLWithPath: FileManager.default.currentDirectoryPath, isDirectory: true)
let fileName = "Day01_1.txt"
let fileURL = cwd.appendingPathComponent(fileName)

// Read text file
do {
	let text = try String(contentsOf: fileURL, encoding: .utf8)
	let split = text.split(separator: "\n")
	var result = 0
	
	for line in split {
		let chars = Array(line)
		
		var firstNumber = 0
		var lastNumber = 0
		
		for char in chars {
			if let number = Int(String(char)) {
				firstNumber = number
				break
			}
		}
		
		for char in chars.reversed() {
			if let number = Int(String(char)) {
				lastNumber = number
				break
			}
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
