function Main(input) {
    const a = input.split(' ');
    const n = parseInt(a[0], 10);
    if(n == 12){
      console.log(1);
    }else{
      console.log(n + 1);
    }
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  