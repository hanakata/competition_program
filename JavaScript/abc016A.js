function Main(input) {
    const a = input.split(' ');
    const n = parseInt(a[0], 10);
    const m = parseInt(a[1], 10);
    const s = n % m;
    if(s == 0){
      console.log("YES");
    }else{
      console.log("NO");
    }
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  