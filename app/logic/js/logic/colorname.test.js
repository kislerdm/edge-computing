const testsN = [
    {
        name: "happy path: Aero",
        in: {r: 124, g: 185, b: 232},
        want: "Aero",
    },
    {
        name: "happy path: Black",
        in: {r: 0, g: 0, b: 0},
        want: "Black",
    },
    {
        name: "happy path: White",
        in: {r: 255, g: 255, b: 255},
        want: "White",
    },
]

testsN.forEach((test) => {
    const got = N(test.in.r, test.in.g, test.in.b);
    if (got !== test.want) {
        throw Error(`Test "${test.name}". got: ${got}, want: ${test.want}`);
    }
});
