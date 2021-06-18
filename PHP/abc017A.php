<?php
$m = 0;
while ($line = fgets(STDIN)) {
  $a = preg_split("/ /",$line);
  $n = intval($a[0] * $a[0] / 10);
  $m = $m + $n;
}
echo $m."\n";
