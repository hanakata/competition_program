function Main(input) {
    const p = input.split('\n');
    const a = p[0].split(' ');
    const s = p[1].split(' ');
    const n = s[0] * a[0];
    const m = s[1] * a[1];
    const o = s[0] + s[1];
    if(a[3] <= o){
      console.log(n + m - (o * a[2]));    
    }else{
      console.log(n + m);
    }
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  