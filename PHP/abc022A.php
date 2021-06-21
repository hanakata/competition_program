<?php
$a = trim(fgets(STDIN));
$n = preg_split("/ /",$a);
echo $n[1]." ".$n[0]."\n";
