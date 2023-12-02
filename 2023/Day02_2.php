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
	$powers = [];

	foreach ($games as $game) {
		[$gameId, $gameContent] = explode(":", $game);
		$minVal = [
			"red" => 0,
			"green" => 0,
			"blue" => 0
		];

		foreach (explode("; ", $gameContent) as $sets) {
			foreach (explode(", ", $sets) as $set) {
				[$count, $color] = explode(" ", trim($set));
				
				if (!$count || !$color) {
					echo "ERROR: EITHER COUNT '$count' OR COLOR '$color' IS EMPTY";
				} else if (!isset($minVal[$color])) {
					echo "ERROR: THERE IS NO COUNT SET FOR COLOR '$color'";
				} else {
					if ($count > $minVal[$color]) {
						$minVal[$color] = $count;
					}
				}
			}
		}
		
		array_push($powers, array_product(array_values($minVal)));
	}
	
	echo array_sum($powers) . "\n";
}
