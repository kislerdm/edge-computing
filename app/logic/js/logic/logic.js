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

function T(r, g, b) {
    return false;
}

function start(r, g, b) {
    let n = N(r, g, b);
    n = n === "" ? "Not found" : n;
    const t = T(r, g, b) === true ? "Warm" : "Cool";
    document.getElementById("color_output").innerHTML = `<div><label for="output_name" id="output_label">Color Name:</label><output name="color_name" id="output_name"> ` +
        n + `</output></div><div><label for="output_type" id="output_label">Color Type:</label><output name="color_type" id="output_type"> ` + t + `</output></div>`;
    return null;
}
