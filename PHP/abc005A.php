<?php
$a = trim(fgets(STDIN));
$n = preg_split("/ /",$a);
echo intval($n[1] / $n[0])."\n";
