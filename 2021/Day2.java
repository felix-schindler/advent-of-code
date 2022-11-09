import helpers.Color;

import java.nio.file.*;
import java.io.*;
import java.awt.*;

public class Day2 {
	public static void main(String[] args) {
		String[] inputArr;
		
		if (args.length > 0) {
			inputArr = args;
		} else {
			try {
				String inputStr = Files.readString(Path.of("Day2.txt"));
				inputArr = inputStr.split("\n");
			} catch (IOException ignored) {
				System.err.println("Failed to read file");
				return;
			}
		}
		
		System.out.println(solve(inputArr));
	}
	
	public static String solve(String[] input) {
		int x = 0, y = 0, aim = 0;
		
		for (String line : input) {
			int value = Integer.parseInt(String.valueOf(line.charAt(line.length() - 1)));
			
			if (line.charAt(0) == 'f') {
				x += value;
				y += aim * value;		// 2nd part
			} else if (line.charAt(0) == 'd') {
				// y += value;			// 2nd part
				aim += value;
			} else if (line.charAt(0) == 'u') {
				// y -= value;			// 2nd part
				aim -= value;
			}
		}
		
		return x + " * " + y + " = " + Color.GREEN + (x * y) + Color.RESET;
	}
}
