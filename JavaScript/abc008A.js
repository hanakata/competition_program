function Main(input) {
    const a = input.split(' ');
    const n = parseInt(a[0], 10);
    const m = parseInt(a[1], 10);
    console.log( m - n + 1 );
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));