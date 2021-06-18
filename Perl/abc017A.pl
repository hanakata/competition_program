my  $i = 0;
while (defined(my $line = <STDIN>)) {
  my @n = split(/ /, $line);
  chomp($n[0]);
  chomp($n[1]);
  my $m = $n[0] * $n[1] / 10;
  $i = $i + $m; 
}
print $i."\n"
