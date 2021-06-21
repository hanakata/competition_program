<?php
$l = [];
$c = [8,4,2,1];
$a = trim(fgets(STDIN));
$n = intval($a);
$i = 0;
$j = 0;
while ($n != 0) {
  if($c[$j] <= $n){
    $n = $n - $c[$j];
    $l[$i] = $c[$j];
    $i = $i + 1;
  }
  $j = $j + 1;
}
$j = 0;
print $i."\n";
while ($j < $i) {
  print $l[$j]."\n";
  $j = $j + 1;
}

