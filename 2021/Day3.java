import helpers.Color;

import java.nio.file.*;
import java.io.*;
import java.awt.*;

public class Day3 {
	public static void main(String[] args) {
		String[] inputArr;
		
		if (args.length > 0) {
			inputArr = args;
		} else {
			try {
				String inputStr = Files.readString(Path.of("Day3.txt"));
				inputArr = inputStr.split("\n");
			} catch (IOException ignored) {
				System.err.println("Failed to read file");
				return;
			}
		}
		
		System.out.println(solve(inputArr));
	}
	
	public static String solve(String[] input) {
		int power = 0, gamma = 0, epsilon = 0;
		
		for (String line : input) {
		}
		
		return "The power consumption is: " + Color.GREEN + power + Color.RESET;
	}
}
