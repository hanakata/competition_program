function Main(a) {
    console.log(Math.floor(a/10 + a%10));
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  