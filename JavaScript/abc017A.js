function Main(input) {
    var a = input.split("\n");
    var i = 0;
    var m = 0;
    while(a[i] != ""){
      var l = a[i].split(' ');
      var b = parseInt(l[0], 10);
      var c = parseInt(l[1], 10);
      var n = b * c / 10;
      m = m + n;
      i += 1;
    }
    console.log(m);  
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  