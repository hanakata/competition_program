<?php
$l = trim(fgets(STDIN));
$s = trim(fgets(STDIN));
if(strlen($l) <  strlen($s)){
  echo $s."\n";
}else{
  echo $l."\n";
}
