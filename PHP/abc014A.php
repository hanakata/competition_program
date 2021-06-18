<?php
$n = trim(fgets(STDIN));
$m = trim(fgets(STDIN));
$i = $n % $m;
if($i == 0){
  echo "0\n";
}else{
  $s = $m - $i;
  echo $s."\n";
}
