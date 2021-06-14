function Main(input) {
    const a = input.split('\n');
    const n = parseInt(a[0], 10);
    var i = 1;
    var m = 0;
    while(i <= n){
      m += i * 10000;
      i += 1;
    }
    const s = m / n;
    console.log(s);  
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));