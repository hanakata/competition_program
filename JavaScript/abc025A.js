function Main(input) {
    const d = input.split('\n');
    const a = d[0].split("");
    const b = d[1];
    const p = Math.floor((b - 1) / 5);
    const q = b % 5;
    if(q == 0){
      s = 4;
    }else{
      s = q - 1;
    }
    console.log(a[p] + a[s]);   
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  