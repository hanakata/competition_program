#!/usr/bin/perl

my $n = <STDIN>;
my $m = <STDIN>;
chomp($n);
chomp($m);
my $i = $n % $m;
if($i == 0){
  print "0\n";
}else{
  my $s = $m - $i;
  print $s."\n";
}
