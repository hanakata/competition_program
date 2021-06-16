#!/usr/bin/perl
 
my $n = <STDIN>;
chomp($n);
my $m =  $n / 2 + $n % 2;
print $m."\n";
