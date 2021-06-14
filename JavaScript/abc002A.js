function Main(input) {
    const a = input.split(' ');
    const n = parseInt(a[0], 10);
    const m = parseInt(a[1], 10);
    if(n < m){
      console.log(m);
    }else{
      console.log(n);
    }
  }
  
Main(require("fs").readFileSync("/dev/stdin", "utf8"));