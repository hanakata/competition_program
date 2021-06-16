function Main(input) {
    const a = input.split('\n');
    const n = parseInt(a[0], 10);
    console.log(parseInt(n / 2, 10) + n % 2);
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));