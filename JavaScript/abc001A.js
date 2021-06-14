function Main(input) {
    const a = input.split('\n');
    const n = parseInt(a[0], 10);
    const m = parseInt(a[1], 10);
  
    console.log(n - m);
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));