function Main(input) {
    var a = input.split("\n");
    var n = parseInt(a[0],10);
    var c = [ 8, 4, 2, 1 ]
    var l = [];
    var i = 0;
    var j = 0;
    while(n != 0){
      if(c[j] <= n){
        n = n - c[j];
        l[i] = c[j];
        i++;
      }
      j++;
    }
    j = 0;
    console.log(i);
    while(j < i){
      console.log(l[j]);
      j++;
    }
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  