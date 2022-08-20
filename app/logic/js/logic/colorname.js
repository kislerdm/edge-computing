import {lT} from "./colornamelookuptable.js";

export function name(r, g, b) {
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
