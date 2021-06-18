#!/usr/bin/perl

my @str = ("A", "B", "C", "D", "E");
my $i = 0;
my $n = <STDIN>;
chomp($n);
while ( $str[$i] ne $n ){
  $i++;
}
my $s = $i + 1;
print $s."\n";
