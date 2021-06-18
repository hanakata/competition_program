#!/usr/bin/perl

my $l = <STDIN>;
my $s = <STDIN>;
chomp($l);
chomp($s);
if(length($l) < length($s)){
  print $s."\n";
}else{
  print $l."\n";
}
