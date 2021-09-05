import * as ChildProcess from "child_process"

async function f(x: number) {
    console.log("In f()...")
    return x + 1;
}

function log(s: string) {
    console.log(s);
}

async function main() {
    console.log("In main()...")
    // const y = f(3);
    // console.log("Still in main()...")
    // console.log("y: " + y);
    // const q = await y;
    // console.log("y: " + y);
    // console.log("q: " + q);
    
    const p = new Promise((acc, rej) => {
        log("In Promise block, acc=" + acc + ", rej=" + rej);
        acc(123);
    });
    log("Awaiting p: " + p);
    var q = await p;
    log("Result: " + q);

    const s = ChildProcess.spawnSync("/usr/bin/ls", ["/"])
    log("" + s.error);
    log("" + s.stdout);
    // const qq = await p;
    // console.log("qq: " + qq);
}

console.log("Starting...")
const mn = main().then(() => {
    console.log("main().then()");
});
// const mnr = await mn;
console.log("mn: " + mn)
// console.log("mnr: " + mnr)
console.log("Ending...")
