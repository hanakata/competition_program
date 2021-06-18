function Main(input) {
    var a = input.split('\n');
    if( parseInt(a) == 1 ) {
      console.log("ABC");
    }else{
      console.log("chokudai");
    }
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  