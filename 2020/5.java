import java.nio.file.Files;
import java.nio.file.Path;
import java.io.IOException;

class Main {
	enum Color {
		GREEN("\u001B[32m"),
		YELLOW("\u001B[33m"),
		RED("\u001B[31m"),
		RESET("\u001B[0m");

		final String val;

		Color(String val) {
			this.val = val;
		}

		public String toString() {
			return this.val;
		}
	}

	public static void main(String[] args) {
		try {
			String[] inputArr;
			var highest = 0;    // the answer

			if (args.length > 0) {
				inputArr = args;
			} else {
				var inputStr = Files.readString(Path.of("5.txt"));
				inputArr = inputStr.split("\n");
			}

			for (String space : inputArr) {
				int row = 0, column = 0; // min: 0

				var rowMax = 127;  // back
				var rowMin = 0;    // front

				var colMax = 7;    // right
				var colMin = 0;    // left

				for (int i = 0; i < 6; i++) {
					if (space.charAt(i) == 'F') {
						rowMax = (int) Math.floor((rowMin + rowMax) * 0.5);
					} else if (space.charAt(i) == 'B') {
						rowMin = (int) Math.floor((rowMin + rowMax) * 0.5 + 1);
					}
				}
				row = space.charAt(6) == 'B' ? rowMax : rowMin;

				for (int i = 7; i < 9; i++) {
					if (space.charAt(i) == 'L') {
						colMax = (int) Math.floor((colMin + colMax) * 0.5);
					} else if (space.charAt(i) == 'R') {
						colMin = (int) Math.floor((colMin + colMax) * 0.5 + 1);
					}
				}
				column = space.charAt(9) == 'R' ? colMax : colMin;

				var seatId = row * 8 + column;

				if (seatId > highest)
					highest = seatId;

				System.out.println("Seat " + Color.YELLOW + space + Color.RESET + " is at row " + Color.GREEN + row + Color.RESET + ", column " + Color.GREEN + column + Color.RESET + ", seat ID " + Color.GREEN + seatId + Color.RESET);
			}

			System.out.println("The highest seat ID is " + Color.GREEN + highest + Color.RESET);
		} catch (IOException ignored) {
			System.err.println(Color.RED + "Failed to open file" + Color.RESET);
		}
	}
}
