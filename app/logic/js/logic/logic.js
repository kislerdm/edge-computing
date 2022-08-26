function N(r, g, b) {
    let o = "",
        d0 = 50.;

    lT.some((c) => {
        const dR = c.r - r,
            dG = c.g - g,
            dB = c.b - b;

        if (dR === 0 && dG === 0 && dB === 0) {
            o = c.n;
            return true;
        }

        const d = Math.sqrt(dR * dR + dG * dG + dB * dB);
        if (d < d0) {
            o = c.n;
            d0 = d;
        }
    });

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
    return Math.exp(-o) <= 1.;
}

const start = (r, g, b) => ({name: N(r, g, b), is_warm: T(r, g, b)});
