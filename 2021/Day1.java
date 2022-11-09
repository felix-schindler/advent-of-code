import helpers.Color;

import java.nio.file.*;
import java.io.*;
import java.awt.*;

public class Day1 {
	public static void main(String[] args) {
		String[] inputArr;
		
		if (args.length > 0) {
			inputArr = args;
		} else {
			try {
				String inputStr = Files.readString(Path.of("Day1.txt"));
				inputArr = inputStr.split("\n");
			} catch (IOException ignored) {
				System.err.println("Failed to read file");
				return;
			}
		}
		
		System.out.println(solve(inputArr));
	}
	
	public static String solve(String[] input) {
		int last = -1;
		int countIncreases = 0;
		
		for (int i = 0; i < input.length - 2; i++) {
			int[] window = new int[3];
			
			try {
				int sum = Integer.parseInt(input[i]) + Integer.parseInt(input[i + 1]) + Integer.parseInt(input[i + 2]);
				
				if (last != -1 && sum > last)
					countIncreases++;
				last = sum;
			} catch (NumberFormatException e) {
				System.err.println(e);
			}
		}
	
//	for (String line : input) {
//		try {
//			int intVal = Integer.parseInt(line);
//			if (last != -1 && intVal > last)
//				countIncreases++;
//			last = intVal;
//		} catch (NumberFormatException ignored) {
//			System.err.println(line + " is not a number");
//		}
//	}
		
		return "There are " + Color.GREEN + countIncreases + Color.RESET + " increases";
	}
}
