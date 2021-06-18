function Main(input) {
    const a = input.split('\n');
    var str = ["A", "B", "C", "D", "E"];
    var i = 0;
    while(str[i] != a[0]){
      i += 1;
    }
    console.log(i + 1);  
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));
  