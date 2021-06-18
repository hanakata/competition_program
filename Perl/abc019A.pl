#!/usr/bin/perl

my $line = <STDIN>;
my @n = split(/ /, $line);
@n = sort {$b <=> $a} @n;
chomp($n[1]);
print $n[1]."\n"
