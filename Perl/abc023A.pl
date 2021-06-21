#!/usr/bin/perl

my $n = <STDIN>;
my $m = $n / 10 + $n % 10;
my $m_int = int $m;
print $m_int."\n";
