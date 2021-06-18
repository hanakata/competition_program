<?php
$a = trim(fgets(STDIN));
$n = preg_split("/ /",$a);
$m = $n[0] % $n[1];
if ($m == 0){
  echo "YES\n";
}else{
  echo "NO\n";
}
