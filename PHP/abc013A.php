<?php
$str = array("A", "B", "C", "D", "E");
$n = trim(fgets(STDIN));
$i = 1;
while( $str[$i] != $n){
  $i++;
}
$s = $i + 1;
echo $s."\n";
