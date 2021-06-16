#!/usr/bin/perl

my $list = <STDIN>;
my @n = split(/ /, $list);
chomp($n[0]);
chomp($n[1]);
$s = int $n[1] - $n[0] + 1;
print $s."\n";
