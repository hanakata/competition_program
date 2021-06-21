function Main(input) {
    const a = input.split(' ');
    const n = a[0].split('\n');
    const m = a[1].split('\n');
    console.log(m[0] + " " + n[0]);
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  