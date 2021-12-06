function Main(input) {
  const a = input.split(' ');
  if(a[1] / a[0] < a[3] / a[2]){
    console.log("AOKI")
  }else if(a[1] / a[0] > a[3] / a[2]){
    console.log("TAKAHASHI")
  }else{
    console.log("DRAW")
  }
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
