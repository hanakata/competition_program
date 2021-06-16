#!/usr/bin/perl
 
my $n = <STDIN>;
chomp($n);
my $m = $n % 3;
if ($m == 0){
  print "YES\n";
}else{
  print "NO\n";
}
