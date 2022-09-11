function N(r, g, b) {
    let o = "",
        d0 = 10,
        mD = 3 * d0;

    const diff = (a, b) => a > b ? a - b : b - a;

    for (c of lT) {
        let m = 0;

        const dR = diff(c.r, r);
        if (dR > d0) {
            continue;
        }
        m += dR;

        const dG = diff(c.g, g);
        if (dG > d0) {
            continue;
        }
        m += dG;

        const dB = diff(c.b, b);
        if (dB > d0) {
            continue;
        }
        m += dB;

        if (m === 0) {
            return c.n;
        }

        if (m < mD) {
            o = c.n;
            mD = m;
        }
    }

    return o;
}

function t(g, id) {
    if (g.id === id) {
        return g;
    }
    if (g.c) {
        for (let g1 of g.c) {
            const n = t(g1, id);
            if (n !== null) {
                return n;
            }
        }
    }
    return null;
}

function T(r, g, b) {
    const d = {"r": r, "g": g, "b": b};
    let o = 0;
    for (const tr of m) {
        let i = 0;
        while (true) {
            const n = t(tr, i);
            if (n === null) {
                throw Error("no node found");
            }

            if (n.c === undefined || n.c.length === 0) {
                o += n.l;
                break;
            }

            if (d[n.f] === undefined) {
                i = n.m;
            } else {
                i = d[n.f] >= n.t ? n.n : n.y;
            }
        }
    }
    return Math.exp(-o) <= 1. ? 1 : 0;
}

const start = (r, g, b) => ({name: N(r, g, b), is_warm: T(r, g, b)});
