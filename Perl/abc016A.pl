#!/usr/bin/perl
 
my $list = <STDIN>;
my @n = split(/ /, $list);
chomp($n[0]);
chomp($n[1]);
my $m = $n[0] % $n[1];
if ($m == 0){
  print "YES\n";
}else{
  print "NO\n";
}
