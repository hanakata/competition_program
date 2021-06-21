function Main(a) {
    const b = a / 2;
    console.log(b * b);   
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  