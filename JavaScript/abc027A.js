  function Main(input) {
      const a = input[0].split(' ');
      if(a[0] == a[1]){
        console.log(a[2]);
      }else if(a[0] == a[2]){
        console.log(a[1]);
      }else{
        console.log(a[0]);
      }
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  