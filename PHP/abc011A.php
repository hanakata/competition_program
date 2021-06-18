<?php
$a = trim(fgets(STDIN));
$n = preg_split("/ /",$a);
if ($n[0] == 12){
  echo "1\n";
}else{
  echo $n[0] + 1."\n";
}
