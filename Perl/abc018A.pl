#!/usr/bin/perl

my @n;
my @r = (1, 1, 1);
my $i = 0;
while (defined(my $line = <STDIN>)) {
  my $m = $line;
  chomp($m);
  $n[$i] = $m;
  $i = $i + 1; 
}
for ($i = 0; $i < 3; $i++){
  for (my $j = 0; $j < 3; $j++){
    if ($n[$i] < $n[$j]){
      $r[$i] = $r[$i] + 1;
    }
  }
}

print $r[0] ."\n";
print $r[1] ."\n";
print $r[2] ."\n";
