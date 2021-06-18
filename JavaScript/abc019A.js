function Main(input) {
    var a = input.split(' ');
    a.sort(function(a,b){
          if( parseInt(a) < parseInt(b) ) return -1;
          if( parseInt(a) > parseInt(b) ) return 1;
  });
    console.log(parseInt(a[1]));
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));