<?php

enum Color : string {
	case GREEN = "\e[32m";
	case YELLOW = "\e[33m";
	case RED = "\e[31m";
	case RESET = "\e[0m";
}

$inputStr = file_get_contents("6.txt");
$groups = explode("\n\n", $inputStr);

$count = 0;

foreach ($groups as $value) {
	$personCount = 0;
	$answers = [];
	$realAnswers = [];
	$answersByPersons = explode("\n", $value);

	foreach ($answersByPersons as $answersByPerson) {
		$personCount++;
		if ($answersByPerson != "") {
			$answerArr = str_split($answersByPerson);
			foreach ($answerArr as $letter) {
				if (!in_array($letter, $answers)) {
					array_push($answers, $letter);
				}
			}
		}
	}

	foreach ($answers as $answer) {
		$letters = str_split($answer);
		$letterCount = [];

		foreach ($letters as $letter) {
			if (isset($letterCount[$letter])) {
				$letterCount[$letter]++;
			} else {
				$letterCount[$letter] = 0;
			}
		}

		print_r($letterCount);

		foreach ($letterCount as $letter => $letCount) {
			if ($letCount == $personCount)
				array_push($realAnswers, $letter);
		}
	}

	// $count += count($answers);
	$count += count($realAnswers);
}

echo "The sum of all answers is " . Color::GREEN->value . $count . Color::RESET->value . "\n";
