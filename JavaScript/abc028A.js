function Main(a) {
    if(a <= 59){
      console.log("Bad");
    }else if(60 <= a && a <= 89){
      console.log("Good");
    }else if(90 <= a && a <= 99){
      console.log("Great");
    }else{
      console.log("Perfect");
    }
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
