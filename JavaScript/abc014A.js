function Main(input) {
    const a = input.split('\n');
    const n = parseInt(a[0], 10);
    const m = parseInt(a[1], 10);
    const i = parseInt(n % m, 10);
    if(i == 0){
      console.log(0);
    }else{
      console.log(m - i);
    }
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  