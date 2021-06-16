function Main(input) {
    const a = input.split('\n');
    console.log(a[0] + "pp");
  }
   
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));