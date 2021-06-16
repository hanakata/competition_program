function Main(input) {
    const a = input.split('\n');
    const n = parseInt(a[0], 10);
    console.log( n - 1 );
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));