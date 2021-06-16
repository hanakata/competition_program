<?php
$n = trim(fgets(STDIN));
$m = $n % 3;
if ($m == 0){
  echo "YES\n";
}else{
  echo "NO\n";
}
