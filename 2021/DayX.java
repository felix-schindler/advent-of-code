import helpers.Color;

import java.nio.file.*;
import java.io.*;
import java.awt.*;

public class DayX {
	public static void main(String[] args) {
		String[] inputArr;
		
		if (args.length > 0) {
			inputArr = args;
		} else {
			try {
				String inputStr = Files.readString(Path.of("DayX.txt"));
				inputArr = inputStr.split("\n");
			} catch (IOException ignored) {
				System.err.println("Failed to read file");
				return;
			}
		}
		
		System.out.println(solve(inputArr));
	}
	
	public static String solve(String[] input) {
		for (String line : input) {
			<#code#>
		}
		
		return Color.YELLOW + "TODO: Implement" + Color.RESET;
	}
}
