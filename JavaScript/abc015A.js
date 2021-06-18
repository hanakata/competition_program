Javascript(Node.js)
function Main(input) {
  const s = input.split('\n');
  if(s[0].length < s[1].length){
    console.log(s[1]);
  }else{
    console.log(s[0]);
  }
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
