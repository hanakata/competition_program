<?php
$p = trim(fgets(STDIN));
$q = trim(fgets(STDIN));
$p_list = preg_split("/ /",$p);
$q_list = preg_split("/ /",$q);
$n = $q_list[0] * $p_list[0];
$m = $q_list[1] * $p_list[1];
$o = $q_list[0] + $q_list[1];
if($p_list[3] <= $o){
  print $n + $m - $o * $p_list[2]."\n";
}else{
  print $n + $m."\n";
}
