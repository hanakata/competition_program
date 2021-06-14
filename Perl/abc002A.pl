#!/usr/bin/perl

my $list = <STDIN>;
my @n = split(/ /, $list);
chomp($n[0]);
chomp($n[1]);
if ($n[0] < $n[1]){
  print $n[1]."\n";
}else{
  print $n[0]."\n";
}