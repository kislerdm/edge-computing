// function NN(t, i) {
//     console.log(i, t);
//     if (t.id === i) {
//         return t;
//     }
//
//     if (t.c === undefined) {
//         return null
//     }
//
//
//     for (const o of t.c) {
//         if (o.id === i) {
//             return o
//         }
//         const n = NN(o, i);
//         if (n !== null) {
//             return n
//         }
//     }
// }

//
// let p = 0.;
// const v = {"r": r, "g": g, "b": b};
//
// m.forEach((t) => {
//     let i = 0;
//     let n = t;
//     while (true) {
//         if (n.c === undefined) {
//             p += n.l;
//             break;
//         }
//
//         n = NN(n, i);
//         if (n === null) {
//             throw Error("node not found")
//         }
//
//
//         const f = v[n.f];
//         if (f === undefined) {
//             i = n.m;
//             continue;
//         }
//
//         i = f >= n.t ? n.n : n.y;
//     }
// })
//
// return Math.exp(-p) <= 1.;

function T(r, g, b) {
    return false;
}

