<?php

function solution($prices, $notes, $x)
{
    $sum = 0;
    
    foreach ($notes as $i => $note) {
        if ($note == "Same as in-store") {
            continue;
        }
        
        $price = $prices[$i];
        
        sscanf($note, "%f%s%s", $percent, $_, $type);
        $percent /= 100;
        
        if ($percent == 0.0) {
            continue;
        }

        if ($type == 'higher') {
            $store = $price / (1 + $percent);
        } else {
            $store = $price / (1 - $percent);
        }

        $overpay = $price - $store;
        $sum += $overpay;
    }

    if (round($sum, 6) <= round($x, 6)) {
        return true;
    }

    return false;
}

$input= [
    'prices' => [48, 165],
    'notes' => [
        "20.00% lower than in-store",
        "10.00% higher than in-store",
    ],
    'x' => 2
];

$input= [
    'prices' => [110, 95, 70],
    'notes' => [
        "10.0% higher than in-store",
        "5.0% lower than in-store",
        "Same as in-store"
    ],
    'x' => 5
];

function boolstr($bool) {
    if (boolval($bool)) {
        return "true";
    }

    return "false";
}

echo boolstr(solution(...$input)).PHP_EOL;
