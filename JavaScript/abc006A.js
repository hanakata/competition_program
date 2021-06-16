function Main(input) {
    const a = input.split('\n');
    const n = parseInt(a[0], 10);
    const m = n % 3;
    if(n == 0){
      console.log("YES");
    }else{
      console.log("NO");
    }
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));