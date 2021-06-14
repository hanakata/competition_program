<?php
$n = trim(fgets(STDIN));
$i = 1;
$m = 0;
while( $i <= $n){
  $m += $i * 10000;
  $i++;
}
$s = $m / $n;
echo $s."\n";