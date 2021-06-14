#!/usr/bin/perl

my $i = 1;
my $m = 0;
my $n = <STDIN>;
chomp($n);
while ( $i <= $n ){
  $m = $m + $i * 10000;
  $i++;
}
my $s = $m / $n;
print $s."\n";