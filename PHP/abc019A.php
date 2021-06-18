<?php
$line = fgets(STDIN);
$a = preg_split("/ /",$line);
for($i = 0;$i < 2; $i++){
  for($j = 2; $j > $i; $j--){
    if(intval($a[$j]) < intval($a[$j-1])){
      $t = $a[$j-1];
      $a[$j-1] = $a[$j];
      $a[$j] = $t;
    }
  }
}
echo (int)$a[1]."\n";
