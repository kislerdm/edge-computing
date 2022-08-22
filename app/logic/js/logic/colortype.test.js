const testsT = [
    {
        name: "happy path: cold",
        in: {r: 182, g: 221, b: 199},
        want: false,
    },
]

testsT.forEach((test) => {
    const got = T(test.in.r, test.in.g, test.in.b);
    if (got !== test.want) {
        throw Error(`Test "${test.name}". got: ${got}, want: ${test.want}`);
    }
});
