<?php
$a = trim(fgets(STDIN));
$b = trim(fgets(STDIN));
$a_list = preg_split("//",$a);
$p = floor(($b - 1) / 5);
$p = $p + 1;
$q = $b % 5;
if($q == 0){
  $s = 5;
}else{
  $s = $q;
}
print $a_list[$p].$a_list[$s]."\n";
