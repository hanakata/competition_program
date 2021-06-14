<?php
$a = trim(fgets(STDIN));
$n = preg_split("/ /",$a);
if ($n[0] < $n[1]){
  echo $n[1]."\n";
}else{
  echo $n[0]."\n";
}