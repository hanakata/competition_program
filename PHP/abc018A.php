<?php
$n = [];
$r = [1,1,1];
$i = 0;
while ($line = fgets(STDIN)) {
  $a = $line;
  $m = intval($a);
  $n[$i] = $m;
  $i = $i + 1;
}
for($i = 0; $i < 3; $i++){
  for($j = 0; $j < 3; $j++){
    if($n[$i] < $n[$j]){
      $r[$i] = $r[$i] + 1;
    }
  }
}
print $r[0] ."\n";
print $r[1] ."\n";
print $r[2] ."\n";
