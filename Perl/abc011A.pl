#!/usr/bin/perl

my $list = <STDIN>;
my @n = split(/ /, $list);
chomp($n[0]);
if ($n[0] == 12){
  print "1\n";
}else{
  print $n[0] + 1."\n";
}
