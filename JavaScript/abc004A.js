function Main(input) {
    const a = input
    const n = parseInt(a, 10);
    const m = n * 2;
    console.log(n);
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));