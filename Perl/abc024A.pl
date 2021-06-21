#!/usr/bin/perl

my $p = <STDIN>;
my $q = <STDIN>;
my @p_list = split(/ /, $p);
my @q_list = split(/ /, $q);
my $n = $q_list[0] * $p_list[0];
my $m = $q_list[1] * $p_list[1];
my $o = $q_list[0] + $q_list[1];
if($p_list[3] <= $o){
  print $n + $m - $o * $p_list[2]."\n";
}else{
  print $n + $m."\n";
}
