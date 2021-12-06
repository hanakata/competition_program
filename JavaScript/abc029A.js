function Main(input) {
  const w = input.split('\n');
  console.log(w[0] + "s");
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
