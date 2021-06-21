#!/usr/bin/perl

my $a = <STDIN>;
my $b = <STDIN>;
my $s = 0;
my @a_list = split(//, $a);
my $p = int(($b - 1) / 5);
my $q = $b % 5;
if($q == 0){
  $s = 4;
}else{
  $s = $q - 1;
}
print $a_list[$p].$a_list[$s]."\n";
