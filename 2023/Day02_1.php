<?php

//$input = <<<INPUT
//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
//Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
//Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
//Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
//Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
//INPUT;
	
$input = file_get_contents(__DIR__ . DIRECTORY_SEPARATOR . "Day02_1.txt");

if ($input !== false) {
	$games = explode("\n", $input);
	$ids = [];

	foreach ($games as $game) {
		[$gameId, $gameContent] = explode(":", $game);

		$setsPossible = [];
		foreach (explode("; ", $gameContent) as $sets) {
			// NOTE: This was wrong at first.
			// At first, I only put the cubes
			// back into the bag after each
			// game. Instead you have to put
			// the cubes back after each set.
			$leftOnes = [
				"red" => 12,
				"green" => 13,
				"blue" => 14
			];

			foreach (explode(", ", $sets) as $set) {
				[$count, $color] = explode(" ", trim($set));
				
				if (!$count || !$color) {
					echo "ERROR: EITHER COUNT '$count' OR COLOR '$color' IS EMPTY";
				} else if (!isset($leftOnes[$color])) {
					echo "ERROR: THERE IS NO COUNT SET FOR COLOR '$color'";
				} else {
					$leftOnes[$color] -= $count;
				}
			}

			foreach (array_values($leftOnes) as $count) {
				$setPossible = true;
				if ($count < 0) {
					$setPossible = false;
				}
				array_push($setsPossible, $setPossible);
			}
		}
		
		if (!in_array(false, $setsPossible)) {
			[$_, $id] = explode(" ", $gameId);
			array_push($ids, $id);
		}
	}
	
	echo array_sum($ids) . "\n";
}
