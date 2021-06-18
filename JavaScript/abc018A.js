function Main(input) {
    var a = input.split("\n");
    var r = [1,1,1]
    var n = [];
    var i = 0;
    while(a[i] != undefined){
      n[i] = parseInt(a[i],10);
      i += 1;
    }
    for(i = 0;i < 3; i++){
      for (var j = 0; j < 3; j++){
        if(n[i] < n[j]){
          r[i] = r[i] + 1;
        }
      }
    }
    console.log(r[0]);
    console.log(r[1]);
    console.log(r[2]);
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  